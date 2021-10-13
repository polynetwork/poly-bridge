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
	"strings"

	"poly-bridge/basedef"

	"github.com/urfave/cli"
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: 1,
	}

	LogDirFlag = cli.StringFlag{
		Name:  "logdir",
		Usage: "log directory",
		Value: "./logs",
	}

	ConfigPathFlag = cli.StringFlag{
		Name:  "cliconfig",
		Usage: "Server config file `<path>`",
		Value: "./config.json",
	}

	ChainIDFlag = cli.Uint64Flag{
		Name:  "chain",
		Usage: "select chainID",
		Value: basedef.ETHEREUM_CROSSCHAIN_ID,
	}

	NFTNameFlag = cli.StringFlag{
		Name:  "name",
		Usage: "set nft name for deploy nft contract, etc.",
		Value: "",
	}

	NFTSymbolFlag = cli.StringFlag{
		Name:  "symbol",
		Usage: "set nft symbol for deploy nft contract, etc.",
		Value: "",
	}

	DstChainFlag = cli.Uint64Flag{
		Name:  "dstChain",
		Usage: "set dest chain for cross chain",
		Value: 0,
	}

	AssetFlag = cli.StringFlag{
		Name:  "asset",
		Usage: "set asset for cross chain or mint nft",
	}

	DstAssetFlag = cli.StringFlag{
		Name:  "dstAsset",
		Usage: "set dest asset for cross chain",
	}

	OwnerAccountFlag = cli.StringFlag{
		Name:  "owner",
		Usage: "set `owner` account",
	}
	SrcAccountFlag = cli.StringFlag{
		Name:  "from",
		Usage: "set `from` account, or approve `sender` account",
	}
	DstAccountFlag = cli.StringFlag{
		Name:  "to",
		Usage: "set `to` account, or approve `spender` account",
	}

	//FeeTokenFlag = cli.BoolTFlag{
	//	Name:  "feeToken",
	//	Usage: "choose erc20 token to be fee token",
	//}

	//NativeTokenFlag = cli.BoolFlag{
	//	Name:  "nativeToken",
	//	Usage: "choose native token as wrapper fee token",
	//}

	//ERC20TokenFlag = cli.BoolFlag{
	//	Name:  "erc20Token",
	//	Usage: "choose erc20 token to be fee token",
	//}

	AmountFlag = cli.StringFlag{
		Name:  "amount",
		Usage: "transfer amount or fee amount, can also used as approve amount",
		Value: "",
	}

	TokenIdFlag = cli.Uint64Flag{
		Name:  "tokenId",
		Usage: "set token id while mint nft",
	}

	LockIdFlag = cli.Uint64Flag{
		Name:  "lockId",
		Usage: "wrap lock nft item id",
	}

	StartFlag = cli.Uint64Flag{
		Name:  "start",
		Usage: "batch get user tokens info with index start",
	}

	LengthFlag = cli.Uint64Flag{
		Name:  "length",
		Usage: "batch get user tokens info with length",
	}

	MethodCodeFlag = cli.StringFlag{
		Name:  "code",
		Usage: "decode method code to params, and code format MUST be hex string",
	}

	AdminIndexFlag = cli.IntFlag{
		Name:  "admin",
		Usage: "admin index in keystore, default value is 0",
		Value: 0,
	}

	GasValueFlag = cli.Uint64Flag{
		Name:  "gasValue",
		Usage: "new gas price if the estimated gas price is not enough, the value should be nGwei, e.g: 4 denotes add 4000000000wei",
		Value: 0,
	}

	EpochFlag = cli.Uint64Flag{
		Name:  "epoch",
		Usage: "set okex epoch",
		Value: 0,
	}

	TxHashFlag = cli.StringFlag{
		Name:  "hash",
		Usage: "set tx hash",
	}
)

var (
	CmdSample = cli.Command{
		Name:   "sample",
		Usage:  "only used to debug this tool.",
		Action: handleSample,
		Flags: []cli.Flag{
			LogLevelFlag,
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
			AmountFlag,
			TokenIdFlag,
		},
	}

	CmdDeployECCDContract = cli.Command{
		Name:   "deployECCD",
		Usage:  "admin account deploy ethereum cross chain data contract.",
		Action: handleCmdDeployECCDContract,
	}

	//CmdDeployECCMContract = cli.Command{
	//	Name:   "deployECCM",
	//	Usage:  "admin account deploy ethereum cross chain manage contract.",
	//	Action: handleCmdDeployECCMContract,
	//}

	CmdDeployCCMPContract = cli.Command{
		Name:   "deployCCMP",
		Usage:  "admin account deploy ethereum cross chain manager proxy contract.",
		Action: handleCmdDeployCCMPContract,
	}

	CmdDeployMintableNFTContract = cli.Command{
		Name:   "deployMintableNFT",
		Usage:  "admin account deploy new mintable nft asset with mapping contract.",
		Action: handleCmdDeployMintableNFTContract,
		Flags: []cli.Flag{
			NFTNameFlag,
			NFTSymbolFlag,
		},
	}

	CmdDeployUnMintableNFTContract = cli.Command{
		Name:   "deployNFT",
		Usage:  "admin account deploy new unmintable nft asset with mapping contract.",
		Action: handleCmdDeployUnMintableNFTContract,
		Flags: []cli.Flag{
			NFTNameFlag,
			NFTSymbolFlag,
		},
	}

	CmdSetNFTLockProxy = cli.Command{
		Name:   "setNFTLockProxy",
		Usage:  "admin set lock proxy for nft asset contract.",
		Action: handleCmdSetNFTLockProxy,
		Flags: []cli.Flag{
			AssetFlag,
		},
	}

	//CmdDeployFeeContract = cli.Command{
	//	Name:   "deployFee",
	//	Usage:  "admin account deploy new mintable erc20 contract.",
	//	Action: handleCmdDeployFeeContract,
	//	//Flags: []cli.Flag{
	//	//	FeeTokenFlag,
	//	//},
	//}

	CmdDeployLockProxyContract = cli.Command{
		Name:   "deployNFTLockProxy",
		Usage:  "admin account deploy nft lock proxy contract.",
		Action: handleCmdDeployLockProxyContract,
	}

	CmdDeployNFTWrapContract = cli.Command{
		Name:   "deployNFTWrapper",
		Usage:  "admin account deploy nft wrapper contract.",
		Action: handleCmdDeployNFTWrapContract,
	}

	CmdDeployNFTQueryContract = cli.Command{
		Name:   "deployNFTQuery",
		Usage:  "admin account deploy nft query contract.",
		Action: handleCmdDeployNFTQueryContract,
	}

	CmdLockProxySetCCMP = cli.Command{
		Name:   "proxySetCCMP",
		Usage:  "admin account set cross chain manager proxy address for lock proxy contract.",
		Action: handleCmdLockProxySetCCMP,
	}

	CmdBindLockProxy = cli.Command{
		Name:   "bindProxy",
		Usage:  "admin  account bind lock proxy contract with another side chain's lock proxy contract.",
		Action: handleCmdBindLockProxy,
		Flags: []cli.Flag{
			DstChainFlag,
		},
	}

	CmdGetBoundLockProxy = cli.Command{
		Name:   "getBoundProxy",
		Usage:  "get bound lock proxy contract.",
		Action: handleCmdGetBoundLockProxy,
		Flags: []cli.Flag{
			DstChainFlag,
		},
	}

	CmdBindNFTAsset = cli.Command{
		Name:   "bindNFT",
		Usage:  "admin account bind nft asset to side chain.",
		Action: handleCmdBindNFTAsset,
		Flags: []cli.Flag{
			AssetFlag,
			DstChainFlag,
			DstAssetFlag,
		},
	}

	CmdBindERC20Asset = cli.Command{
		Name:   "bindToken",
		Usage:  "admin account bind erc20 asset to side chain.",
		Action: handleCmdBindERC20Asset,
		Flags: []cli.Flag{
			AssetFlag,
			DstChainFlag,
			DstAssetFlag,
		},
	}

	CmdTransferECCDOwnership = cli.Command{
		Name:   "transferECCDOwnership",
		Usage:  "admin account transfer ethereum cross chain data contract ownership eccm contract.",
		Action: handleCmdTransferECCDOwnership,
	}

	CmdTransferECCMOwnership = cli.Command{
		Name:   "transferECCMOwnership",
		Usage:  "admin account transfer ethereum cross chain manager contract ownership to ccmp contract.",
		Action: handleCmdTransferECCMOwnership,
	}

	CmdRegisterSideChain = cli.Command{
		Name:   "registerSideChain",
		Usage:  "register side chain in poly.",
		Action: handleCmdRegisterSideChain,
	}

	CmdApproveSideChain = cli.Command{
		Name:   "approveSideChain",
		Usage:  "register side chain in poly.",
		Action: handleCmdApproveSideChain,
	}

	CmdSyncSideChainGenesis2Poly = cli.Command{
		Name:   "syncSideGenesis",
		Usage:  "sync side chain genesis header to poly chain.",
		Action: handleCmdSyncSideChainGenesis2Poly,
		Flags: []cli.Flag{
			EpochFlag,
		},
	}

	CmdSyncPolyGenesis2SideChain = cli.Command{
		Name:   "syncPolyGenesis",
		Usage:  "sync poly genesis header to side chain.",
		Action: handleCmdSyncPolyGenesis2SideChain,
	}

	CmdNFTWrapSetFeeCollector = cli.Command{
		Name:   "setFeeCollector",
		Usage:  "admin account set nft fee collecotr for wrap contract",
		Action: handleCmdNFTWrapSetFeeCollector,
	}

	CmdNFTWrapSetLockProxy = cli.Command{
		Name:   "setWrapLockProxy",
		Usage:  "admin account set nft lock proxy for wrap contract.",
		Action: handleCmdNFTWrapSetLockProxy,
	}

	CmdNFTMint = cli.Command{
		Name:   "mintNFT",
		Usage:  "admin account mint nft token.",
		Action: handleCmdNFTMint,
		Flags: []cli.Flag{
			AssetFlag,
			DstAccountFlag,
			TokenIdFlag,
		},
	}

	CmdNFTApprove = cli.Command{
		Name:   "nftApprove",
		Usage:  "approve nft.",
		Action: handleCmdNFTApprove,
		Flags: []cli.Flag{
			SrcAccountFlag,
			AssetFlag,
			TokenIdFlag,
		},
	}

	CmdNFTOwner = cli.Command{
		Name:   "owner",
		Usage:  "check nft token owner.",
		Action: handleCmdNFTOwner,
		Flags: []cli.Flag{
			AssetFlag,
			TokenIdFlag,
		},
	}

	CmdNFTWrapLock = cli.Command{
		Name:   "lockNFT",
		Usage:  "lock nft token on wrap contract.",
		Action: handleCmdNFTWrapLock,
		Flags: []cli.Flag{
			SrcAccountFlag,
			AssetFlag,
			DstChainFlag,
			DstAccountFlag,
			TokenIdFlag,
			//NativeTokenFlag,
			AmountFlag,
			LockIdFlag,
		},
	}

	CmdMintFee = cli.Command{
		Name:   "mintFee",
		Usage:  "admin account mint fee token.",
		Action: handleCmdMintFee,
		Flags: []cli.Flag{
			//FeeTokenFlag,
			//ERC20TokenFlag,
			DstAccountFlag,
			AmountFlag,
		},
	}

	CmdApproveFee = cli.Command{
		Name:   "approve",
		Usage:  "approve nft wrap contract.",
		Action: handleCmdApprove,
		Flags: []cli.Flag{
			//FeeTokenFlag,
			//ERC20TokenFlag,
			SrcAccountFlag,
			//DstAccountFlag,
			AmountFlag,
		},
	}

	CmdWrapAllowance = cli.Command{
		Name:   "allowance",
		Usage:  "get wrap allowance.",
		Action: handleCmdAllowance,
		Flags: []cli.Flag{
			//FeeTokenFlag,
			//ERC20TokenFlag,
			SrcAccountFlag,
			DstAccountFlag,
		},
	}

	CmdTransferFee = cli.Command{
		Name:   "transferFee",
		Usage:  "transfer fee token.",
		Action: handleCmdTransferFee,
		Flags: []cli.Flag{
			//FeeTokenFlag,
			//ERC20TokenFlag,
			SrcAccountFlag,
			DstAccountFlag,
			AmountFlag,
		},
	}

	CmdGetFeeBalance = cli.Command{
		Name:   "feeBalance",
		Usage:  "get fee balance.",
		Action: handleGetFeeBalance,
		Flags: []cli.Flag{
			//FeeTokenFlag,
			//ERC20TokenFlag,
			SrcAccountFlag,
		},
	}

	CmdNativeTransfer = cli.Command{
		Name:   "transferNative",
		Usage:  "transfer native token.",
		Action: handleCmdNativeTransfer,
		Flags: []cli.Flag{
			SrcAccountFlag,
			DstAccountFlag,
			AmountFlag,
		},
	}

	CmdNativeBalance = cli.Command{
		Name:   "nativeBalance",
		Usage:  "get native balance.",
		Action: handleGetNativeBalance,
		Flags: []cli.Flag{
			SrcAccountFlag,
		},
	}

	CmdTokenUrls = cli.Command{
		Name:   "urls",
		Usage:  "batch get users token url",
		Action: handleCmdBatchGetTokenUrls,
		Flags: []cli.Flag{
			AssetFlag,
			SrcAccountFlag,
			StartFlag,
			LengthFlag,
		},
	}

	CmdNFTBalance = cli.Command{
		Name:   "nftBalance",
		Usage:  "get NFT balance",
		Action: handleCmdNFTBalance,
		Flags: []cli.Flag{
			AssetFlag,
			SrcAccountFlag,
		},
	}

	CmdTransferNFT = cli.Command{
		Name:   "transferNFT",
		Usage:  "transfer NFT token.",
		Action: handleCmdTransferNFT,
		Flags: []cli.Flag{
			AssetFlag,
			SrcAccountFlag,
			DstAccountFlag,
			TokenIdFlag,
		},
	}

	CmdParseLockParams = cli.Command{
		Name:   "decodeLock",
		Usage:  "decode NFT Wrapper Lock method",
		Action: handleCmdDecodeWrapLock,
		Flags: []cli.Flag{
			MethodCodeFlag,
		},
	}

	CmdEnv = cli.Command{
		Name:   "env",
		Usage:  "ensure your environment is correct",
		Action: handleCmdEnv,
		Flags: []cli.Flag{
			OwnerAccountFlag,
		},
	}

	CmdAddGas = cli.Command{
		Name:   "addGas",
		Usage:  "add gas price",
		Action: handleCmdAddGas,
		Flags: []cli.Flag{
			TxHashFlag,
			GasValueFlag,
		},
	}
)

//getFlagName deal with short flag, and return the flag name whether flag name have short name
func getFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
