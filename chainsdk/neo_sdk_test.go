package chainsdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNeo3Sdk_Nep11OwnerOf(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11OwnerOf("0x4fb2f93b37ff47c0c5d14cfc52087e3ca338bc56", "4d65746150616e616365612023302d3031"))
}

func TestNeo3Sdk_Nep11Properties(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11PropertiesByInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "ed00"))
}

func TestNeo3Sdk_GetUrl(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11TokenUri("0x4fb2f93b37ff47c0c5d14cfc52087e3ca338bc56", "4d65746150616e616365612023302d3031"))
}

func TestNeo3Sdk_Nep11BalanceOf(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11BalanceOf("0xafc8bb5d26f43980212da297d3316cf663917bab", "NbtRKN9wPV7WMAQFhVGsecDfNcHZcWkKoY"))

}

func TestNeo3Sdk_Nep11TokensOf(t *testing.T) {
	//assetHash := "0xafc8bb5d26f43980212da297d3316cf663917bab"
	//owner := "NbtRKN9wPV7WMAQFhVGsecDfNcHZcWkKoY"

	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	tokenLists, _ := sdk.Nep11TokensOf("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9")
	for _, v := range tokenLists {
		fmt.Println(v)
	}
	p, _ := sdk.Nep11PropertiesByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", tokenLists)
	a, _ := json.MarshalIndent(p, "", "	")
	fmt.Println(string(a))

}
func TestNeo3Sdk_Nep11PropertiesByInvoke(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	res, _ := sdk.Nep11PropertiesByBatchInvoke("0x4fb2f93b37ff47c0c5d14cfc52087e3ca338bc56", []string{"4d65746150616e616365612023302d3031", "4d65746150616e61636561202332312d3032"})
	a, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(a))
}

func TestNeo3Sdk_Nep11UriByBatchInvoke(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	res, _ := sdk.Nep11UriByBatchInvoke("0x4fb2f93b37ff47c0c5d14cfc52087e3ca338bc56", []string{"4d65746150616e616365612023302d3031", "4d65746150616e61636561202332312d3032"})
	a, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(a))
}
