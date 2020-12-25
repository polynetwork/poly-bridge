package txstatuslisten

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

type FChainTxResp struct {
	ChainId    uint32    `json:"chainid"`
	ChainName  string    `json:"chainname"`
	TxHash     string    `json:"txhash"`
	State      byte      `json:"state"`
	TT         uint32    `json:"timestamp"`
	Fee        string    `json:"fee"`
	Height     uint32    `json:"blockheight"`
	User       string    `json:"user"`
	TChainId   uint32    `json:"tchainid"`
	TChainName string    `json:"tchainname"`
	Contract   string    `json:"contract"`
	Key        string    `json:"key"`
	Param      string    `json:"param"`
	Transfer   *FChainTransferResp `json:"transfer"`
}

type FChainTransferResp struct {
	TokenHash    string `json:"tokenhash"`
	TokenName    string `json:"tokenname"`
	TokenType    string `json:"tokentype"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	ToChain      uint32 `json:"tchainid"`
	ToChainName  string `json:"tchainname"`
	ToTokenHash  string `json:"totokenhash"`
	ToTokenName  string `json:"totokenname"`
	ToTokenType  string `json:"totokentype"`
	ToUser       string `json:"tuser"`
}

type MChainTxResp struct {
	ChainId    uint32 `json:"chainid"`
	ChainName  string `json:"chainname"`
	TxHash     string `json:"txhash"`
	State      byte   `json:"state"`
	TT         uint32 `json:"timestamp"`
	Fee        string `json:"fee"`
	Height     uint32 `json:"blockheight"`
	FChainId   uint32 `json:"fchainid"`
	FChainName string `json:"fchainname"`
	FTxHash    string `json:"ftxhash"`
	TChainId   uint32 `json:"tchainid"`
	TChainName string `json:"tchainname"`
	Key        string `json:"key"`
}

type TChainTxResp struct {
	ChainId    uint32    `json:"chainid"`
	ChainName  string    `json:"chainname"`
	TxHash     string    `json:"txhash"`
	State      byte      `json:"state"`
	TT         uint32    `json:"timestamp"`
	Fee        string    `json:"fee"`
	Height     uint32    `json:"blockheight"`
	FChainId   uint32    `json:"fchainid"`
	FChainName string    `json:"fchainname"`
	Contract   string    `json:"contract"`
	RTxHash    string    `json:"mtxhash"`
	Transfer   *TChainTransferResp `json:"transfer"`
}

type TChainTransferResp struct {
	TokenHash    string `json:"tokenhash"`
	TokenName    string `json:"tokenname"`
	TokenType    string `json:"tokentype"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
}

type CrossTransferResp struct {
	CrossTxType  uint32 `json:"crosstxtype"`
	CrossTxName  string `json:"crosstxname"`
	TT           uint32    `json:"timestamp"`
	FromChainId  uint32 `json:"fromchainid"`
	FromChain    string `json:"fromchainname"`
	FromAddress  string `json:"fromaddress"`
	ToChainId    uint32 `json:"tochainid"`
	ToChain      string `json:"tochainname"`
	ToAddress    string `json:"toaddress"`
	TokenHash    string `json:"tokenhash"`
	TokenName    string `json:"tokenname"`
	TokenType    string `json:"tokentype"`
	Amount       string `json:"amount"`
}

type CrossTxResp struct {
	Transfer       *CrossTransferResp `json:"crosstransfer"`
	Fchaintx       *FChainTxResp      `json:"fchaintx"`
	FchainHeight    uint32               `json:"fchainheight"`
	Fchaintx_valid bool               `json:"fchaintx_valid"`
	Mchaintx       *MChainTxResp      `json:"mchaintx"`
	Mchaintx_valid bool               `json:"mchaintx_valid"`
	MchainHeight    uint32               `json:"mchainheight"`
	Tchaintx       *TChainTxResp      `json:"tchaintx"`
	Tchaintx_valid bool               `json:"tchaintx_valid"`
	TchainHeight    uint32               `json:"tchainheight"`
}

// getcrosstx response
// swagger:response CrossTxResponse
type CrossTxResponse struct {
	// response body
	// in: body
	Body struct {
		Code          int                    `json:"code"`
		Action        string                 `json:"action"`
		Desc          string                 `json:"desc"`
		Version       string                 `json:"version"`
		Result        CrossTxResp            `json:"result"`
	}
}

type PolyExplorerSdk struct {
	client *http.Client
	urls    []string
}

func DefaultPolyExplorerSdk() *PolyExplorerSdk {
	//return NewCoinMarketCapSdk("https://api.coinmarketcap.com/v2")
	return NewPolyExplorerSdk([]string{"https://explorer.poly.network/api/v1/"})
}

func NewPolyExplorerSdk(urls []string) *PolyExplorerSdk {
	client := &http.Client{}
	sdk := &PolyExplorerSdk{
		client: client,
		urls:    urls,
	}
	return sdk
}

func (sdk *PolyExplorerSdk) TxStatus(txHash string) (*CrossTxResponse, error) {
	for i := 0; i < len(sdk.urls); i++ {
		txStatus, err := sdk.txStatus(i, txHash)
		if err != nil {
			logs.Error("poly explorer status err: %s", err.Error())
			continue
		} else {
			return txStatus, nil
		}
	}
	return nil, fmt.Errorf("Cannot get poly explorer status!")
}

func (sdk *PolyExplorerSdk) txStatus(node int, txHash string) (*CrossTxResponse, error) {
	req, err := http.NewRequest("GET", sdk.urls[node]+"getcrosstx", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("txhash", txHash)

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("resp body: %s\n", string(respBody))
	tx := new(CrossTxResponse)
	err = json.Unmarshal(respBody, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
