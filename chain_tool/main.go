/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	log "github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/poly/native/service/header_sync/bsc"
	polyutils "github.com/polynetwork/poly/native/service/utils"
	"github.com/urfave/cli"
	"math/big"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	nftwrap "poly-bridge/go_abi/nft_wrap_abi"
	pabi "poly-bridge/utils/abi"
	xecdsa "poly-bridge/utils/ecdsa"
	"poly-bridge/utils/files"
	"poly-bridge/utils/leveldb"
	"poly-bridge/utils/math"
	"poly-bridge/utils/wallet"
	"runtime"
	"strings"
	"time"
)

var (
	cfgPath  string
	cfg      = new(Config)
	cc       *ChainConfig
	storage  *leveldb.LevelDBImpl
	sdk      *chainsdk.EthereumSdk
	adm      *ecdsa.PrivateKey
	keystore string
)

const defaultAccPwd = "111111"

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "poly nftbridge deploy tool"
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2020 The Ontology Authors"
	app.Flags = []cli.Flag{
		LogLevelFlag,
		//LogDirFlag,
		ConfigPathFlag,
		ChainIDFlag,
		NFTNameFlag,
		NFTSymbolFlag,
		DstChainFlag,
		AssetFlag,
		DstAssetFlag,
		SrcAccountFlag,
		DstAccountFlag,
		//FeeTokenFlag,
		//ERC20TokenFlag,
		//NativeTokenFlag,
		AmountFlag,
		TokenIdFlag,
		MethodCodeFlag,
		OwnerAccountFlag,
		AdminIndexFlag,
		TxHashFlag,
		GasValueFlag,
	}
	app.Commands = []cli.Command{
		CmdSample,
		CmdDeployECCDContract,
		CmdDeployECCMContract,
		CmdDeployCCMPContract,
		CmdDeployUnMintableNFTContract,
		CmdSetNFTLockProxy,
		CmdDeployMintableNFTContract,
		CmdDeployFeeContract,
		CmdDeployLockProxyContract,
		CmdDeployNFTWrapContract,
		CmdDeployNFTQueryContract,
		CmdLockProxySetCCMP,
		CmdBindLockProxy,
		CmdGetBoundLockProxy,
		CmdBindNFTAsset,
		CmdBindERC20Asset,
		CmdTransferECCDOwnership,
		CmdTransferECCMOwnership,
		CmdRegisterSideChain,
		CmdApproveSideChain,
		CmdSyncSideChainGenesis2Poly,
		CmdSyncPolyGenesis2SideChain,
		CmdNFTWrapSetFeeCollector,
		CmdNFTWrapSetLockProxy,
		CmdNFTMint,
		CmdNFTApprove,
		CmdNFTOwner,
		CmdNFTWrapLock,
		CmdMintFee,
		CmdTransferFee,
		CmdGetFeeBalance,
		CmdNativeBalance,
		CmdApproveFee,
		CmdWrapAllowance,
		CmdNativeTransfer,
		CmdTransferNFT,
		CmdTokenUrls,
		CmdNFTBalance,
		CmdParseLockParams,
		CmdEnv,
		CmdAddGas,
	}

	app.Before = beforeCommands
	app.After = afterCommond
	return app
}

func main() {

	app := setupApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// action execute after commands
func beforeCommands(ctx *cli.Context) (err error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load config instance
	cfgPath = ctx.GlobalString(getFlagName(ConfigPathFlag))
	if err = files.ReadJsonFile(cfgPath, cfg); err != nil {
		return fmt.Errorf("read config json file, err: %v", err)
	}

	//logDir := ctx.GlobalString(getFlagName(LogDirFlag))
	//logFormat := fmt.Sprintf(`{"filename":"%s/deploy.log", "perm": "0777"}`, logDir)
	loglevel := ctx.GlobalUint64(getFlagName(LogLevelFlag))
	logFormat := fmt.Sprintf(`{"level:":"%d"}`, loglevel)
	if err := log.SetLogger("console", logFormat); err != nil {
		return fmt.Errorf("set logger failed, err: %v", err)
	}

	// prepare storage for persist account passphrase
	storage = leveldb.NewLevelDBInstance(cfg.LevelDB)

	// select src chainID and prepare config and accounts
	chainID := ctx.GlobalUint64(getFlagName(ChainIDFlag))
	selectChainConfig(chainID)

	if _, err := os.Stat(cfg.Keystore); os.IsNotExist(err) {
		return fmt.Errorf("keystore dir %s is not exist", cfg.Keystore)
	}
	keystore = cfg.Keystore
	admIndex := ctx.GlobalInt(getFlagName(AdminIndexFlag))
	if admIndex < 0 || admIndex >= len(cfg.AdminAccountList) {
		return fmt.Errorf("admin index out of range")
	}
	admAddr := cfg.AdminAccountList[admIndex]
	if adm, err = wallet.LoadEthAccount(storage, keystore, admAddr, defaultAccPwd); err != nil {
		return fmt.Errorf("load eth account for chain %d faild, err: %v", cc.SideChainID, err)
	}

	if sdk, err = chainsdk.NewEthereumSdk(cc.RPC); err != nil {
		return fmt.Errorf("generate sdk for chain %d faild, err: %v", cc.SideChainID, err)
	}

	return nil
}

func afterCommond(ctx *cli.Context) error {
	log.Info("\r\n" +
		"\r\n" +
		"\r\n")
	return nil
}

func handleSample(ctx *cli.Context) error {
	log.Info("start to debug sample...")
	//feeToken := ctx.BoolT(getFlagName(FeeTokenFlag))
	//nativeToken := ctx.Bool(getFlagName(NativeTokenFlag))
	//log.Info("feeToken %v, nativeToken %v", feeToken, nativeToken)
	//getFeeTokenOrERC20Asset(ctx)
	return nil
}

func handleCmdDeployECCDContract(ctx *cli.Context) error {
	log.Info("start to deploy eccd contract...")

	addr, err := sdk.DeployECCDContract(adm)
	if err != nil {
		return fmt.Errorf("deploy eccd for chain %d failed, err: %v", cc.SideChainID, err)
	}

	cc.ECCD = addr.Hex()
	log.Info("deploy eccd for chain %d success %s", cc.SideChainID, addr.Hex())
	return updateConfig()
}

//func handleCmdDeployECCMContract(ctx *cli.Context) error {
//	log.Info("start to deploy eccm contract...")
//
//	eccd := common.HexToAddress(cc.ECCD)
//	addr, err := sdk.DeployECCMContract(adm, eccd, cc.SideChainID)
//	if err != nil {
//		return fmt.Errorf("deploy eccm for chain %d failed, err: %v", cc.SideChainID, err)
//	}
//	cc.ECCM = addr.Hex()
//	log.Info("deploy eccm for chain %d success %s", cc.SideChainID, addr.Hex())
//	return updateConfig()
//}

func handleCmdDeployCCMPContract(ctx *cli.Context) error {
	log.Info("start to deploy ccmp contract...")

	eccm := common.HexToAddress(cc.ECCM)
	addr, err := sdk.DeployECCMPContract(adm, eccm)
	if err != nil {
		return fmt.Errorf("deploy ccmp for chain %d failed, err: %v", cc.SideChainID, err)
	}
	cc.CCMP = addr.Hex()
	log.Info("deploy ccmp for chain %d success %s", cc.SideChainID, addr.Hex())
	return updateConfig()
}

func handleCmdDeployMintableNFTContract(ctx *cli.Context) error {
	log.Info("start to deploy nft contract...")

	name := flag2string(ctx, NFTNameFlag)
	symbol := flag2string(ctx, NFTSymbolFlag)
	owner := xecdsa.Key2address(adm)
	if addr, err := sdk.DeployMintableNFT(adm, name, symbol); err != nil {
		return fmt.Errorf("deploy nft contract for owner %s on chain %d failed, err: %v", owner.Hex(), cc.SideChainID, err)
	} else {
		log.Info("deploy nft contract %s for user %s on chain %d success!", addr.Hex(), owner.Hex(), cc.SideChainID)
	}
	return nil
}

func handleCmdDeployUnMintableNFTContract(ctx *cli.Context) error {
	log.Info("start to deploy nft contract...")

	name := flag2string(ctx, NFTNameFlag)
	symbol := flag2string(ctx, NFTSymbolFlag)
	owner := xecdsa.Key2address(adm)

	addr, err := sdk.DeployUnMintableNFT(adm, name, symbol)
	if err != nil {
		return fmt.Errorf("deploy nft contract for owner %s on chain %d failed, err: %v", owner.Hex(), cc.SideChainID, err)
	} else {
		log.Info("deploy nft contract %s for user %s on chain %d success!", addr.Hex(), owner.Hex(), cc.SideChainID)
	}

	return nil
}

func handleCmdSetNFTLockProxy(ctx *cli.Context) error {
	log.Info("start to deploy nft contract...")

	asset := flag2address(ctx, AssetFlag)
	proxy := common.HexToAddress(cc.NFTLockProxy)

	if _, err := sdk.SetNFTLockProxy(adm, asset, proxy); err != nil {
		return fmt.Errorf("set nft asset lock proxy manager failed, err: %s", err)
	}
	return nil
}

//func handleCmdDeployFeeContract(ctx *cli.Context) error {
//	log.Info("start to deploy erc20 token......")
//
//	addr, err := sdk.DeployERC20(adm)
//	if err != nil {
//		return fmt.Errorf("deploy erc20 token failed, err: %v", err)
//	}
//
//	log.Info("deploy erc20 %s success", addr.Hex())
//	cc.FeeToken = addr.Hex()
//	return updateConfig()
//}

func handleCmdDeployLockProxyContract(ctx *cli.Context) error {
	log.Info("start to deploy nft lock proxy contract...")

	addr, err := sdk.DeployNFTLockProxy(adm)
	if err != nil {
		return fmt.Errorf("deploy nft lock proxy for chain %d failed, err: %v", cc.SideChainID, err)
	}
	cc.NFTLockProxy = addr.Hex()
	log.Info("deploy nft lock proxy for chain %d success %s", cc.SideChainID, addr.Hex())
	return updateConfig()
}

func handleCmdDeployNFTWrapContract(ctx *cli.Context) error {
	log.Info("start to deploy nft wrap contract...")

	addr, err := sdk.DeployWrapContract(adm, cc.SideChainID)
	if err != nil {
		return err
	}

	cc.NFTWrap = addr.Hex()
	log.Info("deploy wrap contract %s success!", addr.Hex())
	return updateConfig()
}

func handleCmdDeployNFTQueryContract(ctx *cli.Context) error {
	log.Info("start to deploy nft query contract...")

	var queryLimit uint64 = 24
	addr, err := sdk.DeployNFTQueryContract(adm, queryLimit)
	if err != nil {
		return err
	}

	cc.NFTQuery = addr.Hex()
	log.Info("deploy query contract %s success!", addr.Hex())
	return updateConfig()
}

func handleCmdLockProxySetCCMP(ctx *cli.Context) error {
	log.Info("start to set ccmp for lock proxy contract...")

	proxy := common.HexToAddress(cc.NFTLockProxy)
	ccmp := common.HexToAddress(cc.CCMP)
	hash, err := sdk.NFTLockProxySetCCMP(adm, proxy, ccmp)
	if err != nil {
		return fmt.Errorf("nft lock proxy set ccmp for chain %d failed, err: %v", cc.SideChainID, err)
	}
	log.Info("nft lock proxy set ccmp for chain %d success! hash %s", cc.SideChainID, hash.Hex())
	return nil
}

func handleCmdBindLockProxy(ctx *cli.Context) error {
	log.Info("start to bind lock proxy...")

	dstChainId := flag2Uint64(ctx, DstChainFlag)
	dstChainCfg := customSelectChainConfig(dstChainId)
	proxy := common.HexToAddress(cc.NFTLockProxy)
	dstProxy := common.HexToAddress(dstChainCfg.NFTLockProxy)

	hash, err := sdk.BindLockProxy(adm, proxy, dstProxy, dstChainId)
	if err != nil {
		return fmt.Errorf("bind lock proxy (src proxy %s, dst proxy %s, src chain id %d, dst chain id %d)",
			cc.NFTLockProxy, dstChainCfg.NFTLockProxy, cc.SideChainID, dstChainId)
	}

	log.Info("bind lock proxy (src proxy %s, dst proxy %s, src chain id %d, dst chain id %d), txhash %s",
		cc.NFTLockProxy, dstChainCfg.NFTLockProxy, cc.SideChainID, dstChainId, hash.Hex())
	return nil
}

func handleCmdGetBoundLockProxy(ctx *cli.Context) error {
	log.Info("start to get bound lock proxy contract...")

	dstChainId := ctx.GlobalUint64(getFlagName(DstChainFlag))
	proxy := common.HexToAddress(cc.NFTLockProxy)

	if addr, err := sdk.GetBoundNFTProxy(proxy, dstChainId); err != nil {
		return fmt.Errorf("check bound nft lock proxy err: %v", err)
	} else {
		log.Info("proxy %s bound to %s in with target chain id of %d", cc.NFTLockProxy, addr.Hex(), dstChainId)
	}

	return nil
}

func handleCmdBindNFTAsset(ctx *cli.Context) error {
	log.Info("start to bind nft asset...")

	srcAsset := flag2address(ctx, AssetFlag)
	dstAsset := flag2address(ctx, DstAssetFlag)
	dstChainId := flag2Uint64(ctx, DstChainFlag)
	dstChainCfg := customSelectChainConfig(dstChainId)
	owner := xecdsa.Key2address(adm)
	proxy := common.HexToAddress(cc.NFTLockProxy)

	hash, err := sdk.BindNFTAsset(
		adm,
		proxy,
		srcAsset,
		dstAsset,
		dstChainId,
	)
	if err != nil {
		return fmt.Errorf("bind nft asset (src chain id %d, src asset %s, src proxy %s) - "+
			"(dst chain id %d, dst asset %s, dst proxy %s)"+
			" for user %s failed, err: %v",
			cc.SideChainID, srcAsset.Hex(), cc.NFTLockProxy,
			dstChainId, dstAsset.Hex(), dstChainCfg.NFTLockProxy,
			owner.Hex(), err)
	}

	log.Info("bind nft asset (src chain id %d, src asset %s, src proxy %s) - "+
		"(dst chain id %d, dst asset %s, dst proxy %s)"+
		" for user %s success! txhash %s",
		cc.SideChainID, srcAsset.Hex(), cc.NFTLockProxy,
		dstChainId, dstAsset.Hex(), dstChainCfg.NFTLockProxy,
		owner.Hex(), hash.Hex())
	return nil
}

func handleCmdBindERC20Asset(ctx *cli.Context) error {
	log.Info("start to bind nft asset...")

	srcAsset := flag2address(ctx, AssetFlag)
	dstAsset := flag2address(ctx, DstAssetFlag)
	dstChainId := flag2Uint64(ctx, DstChainFlag)
	dstChainCfg := customSelectChainConfig(dstChainId)
	owner := xecdsa.Key2address(adm)
	proxy := common.HexToAddress(cc.LockProxy)

	hash, err := sdk.BindERC20Asset(
		adm,
		proxy,
		srcAsset,
		dstAsset,
		dstChainId,
	)
	if err != nil {
		return fmt.Errorf("bind erc20 asset (src chain id %d, src asset %s, src proxy %s) - "+
			"(dst chain id %d, dst asset %s, dst proxy %s)"+
			" for user %s failed, err: %v",
			cc.SideChainID, srcAsset.Hex(), cc.LockProxy,
			dstChainId, dstAsset.Hex(), dstChainCfg.LockProxy,
			owner.Hex(), err)
	}

	log.Info("bind erc20 asset (src chain id %d, src asset %s, src proxy %s) - "+
		"(dst chain id %d, dst asset %s, dst proxy %s)"+
		" for user %s success! txhash %s",
		cc.SideChainID, srcAsset.Hex(), cc.LockProxy,
		dstChainId, dstAsset.Hex(), dstChainCfg.LockProxy,
		owner.Hex(), hash.Hex())
	return nil
}

func handleCmdTransferECCDOwnership(ctx *cli.Context) error {
	log.Info("start to transfer eccd ownership...")

	eccd := common.HexToAddress(cc.ECCD)
	eccm := common.HexToAddress(cc.ECCM)

	if hash, err := sdk.TransferECCDOwnership(adm, eccd, eccm); err != nil {
		return fmt.Errorf("transfer eccd %s ownership to eccm %s on chain %d failed, err: %v",
			cc.ECCD, cc.ECCM, cc.SideChainID, err)
	} else {
		log.Info("transfer eccd %s ownership to eccm %s on chain %d success, txhash: %s",
			cc.ECCD, cc.ECCM, cc.SideChainID, hash.Hex())
	}
	return nil
}

func handleCmdTransferECCMOwnership(ctx *cli.Context) error {
	log.Info("start to transfer eccm ownership...")

	eccm := common.HexToAddress(cc.ECCM)
	ccmp := common.HexToAddress(cc.CCMP)

	if hash, err := sdk.TransferECCMOwnership(adm, eccm, ccmp); err != nil {
		return fmt.Errorf("transfer eccm %s ownership to ccmp %s on chain %d failed, err: %v",
			cc.ECCM, cc.CCMP, cc.SideChainID, err)
	} else {
		log.Info("transfer eccm %s ownership to ccmp %s on chain %d success, txhash: %s",
			cc.ECCM, cc.CCMP, cc.SideChainID, hash.Hex())
	}

	return nil
}

func handleCmdRegisterSideChain(ctx *cli.Context) error {
	validators, err := wallet.LoadPolyAccountList(cfg.Poly.Keystore, cfg.Poly.Passphrase)
	if err != nil {
		return err
	}
	polySdk, err := chainsdk.NewPolySdkAndSetChainID(cfg.Poly.RPC)
	if err != nil {
		return err
	}

	// todo: 验证heco的注册方式
	eccd := common.HexToAddress(cc.ECCD)
	chainID := cc.SideChainID
	switch chainID {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		router := polyutils.ETH_ROUTER
		err = polySdk.RegisterSideChain(validators[0], chainID, 1, router, eccd, cc.SideChainName)

	case basedef.BSC_CROSSCHAIN_ID:
		router := polyutils.BSC_ROUTER
		ext := bsc.ExtraInfo{
			ChainID: new(big.Int).SetUint64(chainID),
		}
		extEnc, _ := json.Marshal(ext)
		err = polySdk.RegisterSideChainExt(validators[0], chainID, 1, router, eccd, cc.SideChainName, extEnc)

	case basedef.HECO_CROSSCHAIN_ID:
		router := polyutils.HECO_ROUTER
		err = polySdk.RegisterSideChain(validators[0], chainID, 1, router, eccd, cc.SideChainName)

	default:
		err = fmt.Errorf("chain id %d invalid", chainID)
	}

	if err != nil {
		return err
	}

	log.Info("register side chain %d eccd %s success", chainID, eccd.Hex())
	return nil
}

func handleCmdApproveSideChain(ctx *cli.Context) error {
	validators, err := wallet.LoadPolyAccountList(cfg.Poly.Keystore, cfg.Poly.Passphrase)
	if err != nil {
		return err
	}
	polySdk, err := chainsdk.NewPolySdkAndSetChainID(cfg.Poly.RPC)
	if err != nil {
		return err
	}
	if err := polySdk.ApproveRegisterSideChain(cc.SideChainID, validators); err != nil {
		return fmt.Errorf("failed to approve register side chain, err: %s", err)
	}

	log.Info("approve register side chain %d success", cc.SideChainID)
	return nil
}

func handleCmdSyncSideChainGenesis2Poly(ctx *cli.Context) error {
	log.Info("start to sync side chain %s genesis header to poly chain...", cc.SideChainName)

	polySdk, err := chainsdk.NewPolySdkAndSetChainID(cfg.Poly.RPC)
	if err != nil {
		return err
	}
	validators, err := wallet.LoadPolyAccountList(cfg.Poly.Keystore, cfg.Poly.Passphrase)
	if err != nil {
		return err
	}

	switch cc.SideChainID {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		err = SyncEthGenesisHeader2Poly(cc.SideChainID, sdk, polySdk, validators)
	case basedef.BSC_CROSSCHAIN_ID:
		err = SyncBscGenesisHeader2Poly(cc.SideChainID, sdk, polySdk, validators)
	case basedef.HECO_CROSSCHAIN_ID:
		err = SyncHecoGenesisHeader2Poly(cc.SideChainID, sdk, polySdk, validators)
	default:
		err = fmt.Errorf("chain id %d invalid", cc.SideChainID)
	}
	if err != nil {
		return fmt.Errorf("sync side chain %d genesis header to poly failed, err: %v", cc.SideChainID, err)
	} else {
		log.Info("sync side chain %d genesis header to poly success!", cc.SideChainID)
	}
	return nil
}

func handleCmdSyncPolyGenesis2SideChain(ctx *cli.Context) error {
	log.Info("start to sync poly chain genesis header to side chain...")

	polySdk, err := chainsdk.NewPolySdkAndSetChainID(cfg.Poly.RPC)
	if err != nil {
		return err
	}
	eccm := common.HexToAddress(cc.ECCM)

	if err := SyncPolyGenesisHeader2Eth(
		polySdk,
		adm,
		sdk,
		eccm,
	); err != nil {
		return fmt.Errorf("sync poly chain genesis header to side chain %d failed, err: %v", cc.SideChainID, err)
	}
	log.Info("sync poly chain genesis header to side chain %d success!", cc.SideChainID)
	return nil
}

func handleCmdNFTWrapSetFeeCollector(ctx *cli.Context) error {
	log.Info("start to set fee collector for wrap contract...")

	wrapper := common.HexToAddress(cc.NFTWrap)
	feeCollector := common.HexToAddress(cc.FeeCollector)

	if tx, err := sdk.SetWrapFeeCollector(adm, wrapper, feeCollector); err != nil {
		return fmt.Errorf("set fee collector failed, err: %v", err)
	} else {
		log.Info("set fee collector success, hash %s", tx.Hex())
	}
	return nil
}

func handleCmdNFTWrapSetLockProxy(ctx *cli.Context) error {
	log.Info("start to set lock proxy for wrap contract...")

	wrapper := common.HexToAddress(cc.NFTWrap)
	proxy := common.HexToAddress(cc.NFTLockProxy)

	if tx, err := sdk.SetWrapLockProxy(adm, wrapper, proxy); err != nil {
		return fmt.Errorf("set lock proxy for wrap failed, err: %v", err)
	} else {
		log.Info("set wrap lock proxy success, hash %s", tx.Hex())
	}
	return nil
}

func handleCmdNFTMint(ctx *cli.Context) error {
	log.Info("start to mint nft...")

	asset := flag2address(ctx, AssetFlag)
	to := flag2address(ctx, DstAccountFlag)
	tokenID := flag2big(ctx, TokenIdFlag)
	uri := cfg.OSS + tokenID.String()
	tx, err := sdk.MintNFT(adm, asset, to, tokenID, uri)
	if err != nil {
		return err
	}
	log.Info("mint nft %d to %s success! txhash %s", tokenID.Uint64(), to.Hex(), tx.Hex())
	return nil
}

func handleCmdNFTApprove(ctx *cli.Context) error {
	log.Info("start to approve nft owner...")

	asset := flag2address(ctx, AssetFlag)
	tokenID := flag2big(ctx, TokenIdFlag)
	owner := flag2address(ctx, SrcAccountFlag)
	key, err := wallet.LoadEthAccount(storage, keystore, owner.Hex(), defaultAccPwd)
	if err != nil {
		return err
	}
	spender := common.HexToAddress(cc.NFTWrap)
	tx, err := sdk.NFTApprove(key, asset, spender, tokenID)
	if err != nil {
		return err
	}

	log.Info("nft %d owner %s approved to %s success, txhash %s",
		tokenID.Uint64(), owner.Hex(), spender.Hex(), tx.Hex())
	return nil
}

func handleCmdNFTOwner(ctx *cli.Context) error {
	log.Info("start to get nft owner...")

	asset := flag2address(ctx, AssetFlag)
	tokenID := flag2big(ctx, TokenIdFlag)
	owner, err := sdk.GetNFTOwner(asset, tokenID)
	if err != nil {
		return err
	}

	approvedTo, err := sdk.GetNFTApproved(asset, tokenID)
	if err != nil {
		return err
	}

	uri, err := sdk.GetNFTTokenUri(asset, tokenID)
	if err != nil {
		return err
	}

	log.Info("nft %d owner %s, uri %s, approved to %s", tokenID.Uint64(), owner.Hex(), uri, approvedTo.Hex())
	return nil
}

func handleCmdNFTWrapLock(ctx *cli.Context) error {
	log.Info("start to lock nft...")

	from := flag2address(ctx, SrcAccountFlag)
	key, err := wallet.LoadEthAccount(storage, keystore, from.Hex(), defaultAccPwd)
	if err != nil {
		return err
	}

	asset := flag2address(ctx, AssetFlag)
	to := flag2address(ctx, DstAccountFlag)
	dstChainID := flag2Uint64(ctx, DstChainFlag)
	tokenID := flag2big(ctx, TokenIdFlag)
	fee := flag2big(ctx, AmountFlag)
	id := new(big.Int).SetUint64(flag2Uint64(ctx, LockIdFlag))
	wrapper := common.HexToAddress(cc.NFTWrap)

	log.Info("wrap lock nft, [assset:%s, to:%s, dstChainID:%d, tokenID:%s, feeToken:%s, fee:%s, id: %s]",
		asset.Hex(), to.Hex(), dstChainID, tokenID.String(), cc.FeeToken, fee.String(), id.String())

	{
		log.Info("approve nft token")
		approved, err := sdk.GetNFTApproved(asset, tokenID)
		if err != nil {
			return err
		}
		if approved != wrapper {
			if _, err := sdk.NFTApprove(key, asset, wrapper, tokenID); err != nil {
				return err
			}
			time.Sleep(20 * time.Second)
		}
	}

	// wrapper lock with native fee token
	if cc.FeeToken == "" {
		if tx, err := sdk.WrapLockWithNativeFeeToken(key, wrapper, asset, to, dstChainID, tokenID, fee, id); err != nil {
			return err
		} else {
			log.Info("wrap lock with native fee token success, tx %s", tx.Hex())
			return nil
		}
	}

	//{
	//	feeToken := common.HexToAddress(cc.FeeToken)
	//	if _, err := sdk.ApproveERC20Token(key, feeToken, wrapper, fee); err != nil {
	//		return err
	//	}
	//	allowance, err := sdk.GetERC20Allowance(feeToken, from, wrapper)
	//	if err != nil {
	//		return err
	//	}
	//	if allowance.Cmp(fee) < 0 {
	//		return fmt.Errorf("approve fee token to wrapper contract failed")
	//	}
	//	tx, err := sdk.WrapLockWithErc20FeeToken(key, wrapper, asset, to, dstChainID, tokenID, feeToken, fee, id)
	//	if err != nil {
	//		return err
	//	}
	//	log.Info("wrap lock with erc20 feeToken success, feeToken %s, tx %s", feeToken.Hex(), tx.Hex())
	//}

	return nil
}

func handleCmdMintFee(ctx *cli.Context) error {
	log.Info("start to mint fee token...")

	//asset := getFeeTokenOrERC20Asset(ctx)
	asset := common.HexToAddress(cc.FeeToken)
	to := flag2address(ctx, DstAccountFlag)
	amount := flag2big(ctx, AmountFlag)
	log.Debug("mint to %s %s", to.Hex(), amount.String())

	tx, err := sdk.MintERC20Token(adm, asset, to, amount)
	if err != nil {
		return err
	}
	log.Info("mint %s to %s success, hash %s", amount.String(), to.Hex(), tx.Hex())
	return nil
}

func handleCmdTransferFee(ctx *cli.Context) error {
	log.Info("start to transfer fee token...")

	asset := common.HexToAddress(cc.FeeToken) //getFeeTokenOrERC20Asset(ctx)
	from := flag2address(ctx, SrcAccountFlag)
	key, err := wallet.LoadEthAccount(storage, keystore, from.Hex(), defaultAccPwd)
	if err != nil {
		return err
	}

	to := flag2address(ctx, DstAccountFlag)
	amount := flag2big(ctx, AmountFlag)
	tx, err := sdk.TransferERC20Token(key, asset, to, amount)
	if err != nil {
		return err
	}

	log.Info("%s transfer %s to %s success, tx %s", from.Hex(), amount.String(), to.Hex(), tx.Hex())
	return nil
}

func handleGetFeeBalance(ctx *cli.Context) error {
	owner := flag2address(ctx, SrcAccountFlag)
	asset := common.HexToAddress(cc.FeeToken) //getFeeTokenOrERC20Asset(ctx)

	balance, err := sdk.GetERC20Balance(asset, owner)
	if err != nil {
		return fmt.Errorf("get balance failed, err: %v", err)
	}
	log.Info("%s balance of asset %s is %s", owner.Hex(), asset.Hex(), balance.String())
	return nil
}

func handleGetNativeBalance(ctx *cli.Context) error {
	owner := flag2address(ctx, SrcAccountFlag)
	balance, err := sdk.GetNativeBalance(owner)
	if err != nil {
		return fmt.Errorf("get native balance faild, err: %v", err)
	}
	log.Info("%s native balance is %s", owner.Hex(), balance.String())
	return nil
}

func handleCmdApprove(ctx *cli.Context) error {
	asset := common.HexToAddress(cc.FeeToken) //getFeeTokenOrERC20Asset(ctx)
	sender := flag2address(ctx, SrcAccountFlag)
	key, err := wallet.LoadEthAccount(storage, keystore, sender.Hex(), defaultAccPwd)
	if err != nil {
		return err
	}
	spender := common.HexToAddress(cc.NFTWrap) //flag2address(ctx, DstAccountFlag)
	amount := flag2big(ctx, AmountFlag)
	if tx, err := sdk.ApproveERC20Token(key, asset, spender, amount); err != nil {
		return err
	} else {
		log.Info("sender %s approve spender %s %s on asset of %s, txhash %s",
			sender.Hex(), spender.Hex(), amount.String(), asset.Hex(), tx.Hex())
	}
	return nil
}

func handleCmdAllowance(ctx *cli.Context) error {
	asset := common.HexToAddress(cc.FeeToken) //getFeeTokenOrERC20Asset(ctx)
	owner := flag2address(ctx, SrcAccountFlag)
	spender := common.HexToAddress(cc.NFTWrap) //flag2address(ctx, DstAccountFlag)
	if amount, err := sdk.GetERC20Allowance(asset, owner, spender); err != nil {
		return err
	} else {
		log.Info("sender %s already approved spender %s %s on asset of %s",
			owner.Hex(), spender.Hex(), amount.String(), asset.Hex())
	}
	return nil
}

func handleCmdNativeTransfer(ctx *cli.Context) error {
	log.Info("start to transfer native token on chain %s...", cc.SideChainName)

	from := flag2address(ctx, SrcAccountFlag)
	key, err := wallet.LoadEthAccount(storage, keystore, from.Hex(), defaultAccPwd)
	if err != nil {
		return err
	}

	to := flag2address(ctx, DstAccountFlag)
	amount := flag2big(ctx, AmountFlag)
	tx, err := sdk.TransferNative(key, to, amount)
	if err != nil {
		return err
	}
	log.Info("%s transfer %s to %s success, txhash %s", from.Hex(), amount.String(), to.Hex(), tx.Hex())
	return nil
}

func handleCmdBatchGetTokenUrls(ctx *cli.Context) error {
	log.Info("start to get user's NFT token urls...")

	asset := flag2address(ctx, AssetFlag)
	owner := flag2address(ctx, SrcAccountFlag)
	start := flag2Uint64(ctx, StartFlag)
	length := flag2Uint64(ctx, LengthFlag)
	wrapper := common.HexToAddress(cc.NFTWrap)

	data, err := sdk.GetOwnerNFTsByIndex(wrapper, asset, owner, int(start), int(length))
	if err != nil {
		return err
	}

	for tokenId, url := range data {
		log.Info("user NFT token %d url is %s", tokenId, url)
	}

	return nil
}

func handleCmdNFTBalance(ctx *cli.Context) error {
	log.Info("start to get user's NFT balance")

	asset := flag2address(ctx, AssetFlag)
	owner := flag2address(ctx, SrcAccountFlag)

	data, err := sdk.GetNFTBalance(asset, owner)
	if err != nil {
		return err
	}

	log.Info("user %s balance on NFT Asset %s is %d", owner.Hex(), asset.Hex(), data.Uint64())
	return nil
}

func handleCmdTransferNFT(ctx *cli.Context) error {
	log.Info("start to transfer NFT token...")

	asset := flag2address(ctx, AssetFlag)
	from := flag2address(ctx, SrcAccountFlag)
	to := flag2address(ctx, DstAccountFlag)
	tokenId := flag2big(ctx, TokenIdFlag)
	key, err := wallet.LoadEthAccount(storage, keystore, from.Hex(), defaultAccPwd)
	if err != nil {
		return err
	}

	tx, err := sdk.NFTSafeTransferTo(key, asset, from, to, tokenId)
	if err != nil {
		return err
	}

	log.Info("%s transfer NFT %d to %s success, tx %s", from.Hex(), tokenId.Uint64(), to.Hex(), tx.Hex())
	return nil
}

func handleCmdDecodeWrapLock(ctx *cli.Context) error {
	log.Info("start to decode wrapper lock method params")

	code := flag2string(ctx, MethodCodeFlag)
	code = strings.TrimPrefix(code, "0x")

	abiStr := strings.NewReader(nftwrap.PolyNFTWrapperABI)
	wrapperABI, err := abi.JSON(abiStr)
	if err != nil {
		return err
	}

	enc, err := hex.DecodeString(code)
	if err != nil {
		return err
	}

	data := new(chainsdk.WrapLockMethod)
	err = pabi.UnpackMethod(wrapperABI, "lock", data, enc[:])
	if err != nil {
		return err
	}

	log.Info("data: {\r\n toChainId %d\r\n tokenId %d\r\n fromAsset %s\r\n toAddress %s\r\n feeToken %s\r\n fee %s\r\n dataId %d\r\n}",
		data.ToChainId, data.TokenId.Uint64(), data.FromAsset.Hex(), data.ToAddress.Hex(), data.FeeToken.Hex(), data.Fee.String(), data.Id)

	return nil
}

func handleCmdAddGas(ctx *cli.Context) error {
	log.Info("add gas price for transaction")

	num := flag2Uint64(ctx, GasValueFlag)
	if num == 0 {
		return fmt.Errorf("added gasValue is 0")
	}
	unit := new(big.Int).SetUint64(num)
	gw := big.NewInt(1000000000)
	newGasPrice := new(big.Int).Mul(unit, gw)

	hashstr := flag2string(ctx, TxHashFlag)
	hash := common.HexToHash(hashstr)
	newhash, err := sdk.AddGas(adm, hash, newGasPrice)
	if err != nil {
		return err
	}

	log.Info("add gas for %s --> %s", hashstr, newhash.Hex())
	return nil
}

func handleCmdEnv(ctx *cli.Context) error {
	currentInfo := fmt.Sprintf("current env: side chain name %s, side chain id %d\r\n", cc.SideChainName, cc.SideChainID)

	polyInfo := fmt.Sprintf("poly side chain id - %d\r\n", basedef.POLY_CROSSCHAIN_ID)
	ethInfo := fmt.Sprintf("eth side chain id - %d\r\n", basedef.ETHEREUM_CROSSCHAIN_ID)
	ontInfo := fmt.Sprintf("ont side chain id - %d\r\n", basedef.ONT_CROSSCHAIN_ID)
	neoInfo := fmt.Sprintf("neo side chain id - %d\r\n", basedef.NEO_CROSSCHAIN_ID)
	bscInfo := fmt.Sprintf("bsc side chain id - %d\r\n", basedef.BSC_CROSSCHAIN_ID)
	hecoInfo := fmt.Sprintf("heco side chain id - %d\r\n", basedef.HECO_CROSSCHAIN_ID)
	o3Info := fmt.Sprintf("o3 side chain id - %d\r\n", basedef.O3_CROSSCHAIN_ID)

	log.Info(currentInfo, polyInfo, ethInfo, ontInfo, neoInfo, bscInfo, hecoInfo, o3Info)

	owner := flag2address(ctx, OwnerAccountFlag)
	addr := owner.Hex()
	log.Info("check your owner address %s in dir %s", keystore, addr)
	_, err := wallet.LoadEthAccount(storage, keystore, addr, defaultAccPwd)
	if err != nil {
		return err
	}
	//enc := crypto.FromECDSA(key)
	//log.Info(hex.EncodeToString(enc))
	return nil
}

// getFeeTokenOrERC20Asset return feeToken if `feeToken` is true
//func getFeeTokenOrERC20Asset(ctx *cli.Context) common.Address {
//	if ctx.Bool(getFlagName(NativeTokenFlag)) {
//		return nativeToken
//	}
//	if ctx.Bool(getFlagName(ERC20TokenFlag)) {
//		return flag2address(ctx, ERC20TokenFlag)
//	}
//	return common.HexToAddress(cc.FeeToken)
//}

func updateConfig() error {
	if err := files.WriteJsonFile(cfgPath, cfg, true); err != nil {
		return err
	}
	log.Info("update config success!", cfgPath)
	return nil
}

func selectChainConfig(chainID uint64) {
	cc = customSelectChainConfig(chainID)
}

func flag2string(ctx *cli.Context, f cli.Flag) string {
	fn := getFlagName(f)
	data := ctx.String(fn)
	return data
}

func flag2address(ctx *cli.Context, f cli.Flag) common.Address {
	data := flag2string(ctx, f)
	return common.HexToAddress(data)
}

func flag2big(ctx *cli.Context, f cli.Flag) *big.Int {
	fn := getFlagName(f)
	data := ctx.String(fn)
	return math.String2BigInt(data)
}

func flag2Uint64(ctx *cli.Context, f cli.Flag) uint64 {
	fn := getFlagName(f)
	data := ctx.Uint64(fn)
	return data
}

func customSelectChainConfig(chainID uint64) *ChainConfig {
	switch chainID {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return cfg.Ethereum
	case basedef.BSC_CROSSCHAIN_ID:
		return cfg.Bsc
	case basedef.HECO_CROSSCHAIN_ID:
		return cfg.Heco
	}
	panic(fmt.Sprintf("invalid chain id %d", chainID))
}
