package chainsdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNeo3AddrToHash160(t *testing.T) {
	addrStr, _ := Neo3AddrToHash160("NVGUQ1qyL4SdSm7sVmGVkXetjEsvw2L3NT")
	fmt.Println(addrStr)
	fmt.Println(Hash160StrToNeo3Addr(addrStr.String()))
	reversed, _ := Neo3AddrToReverseHash160("NVGUQ1qyL4SdSm7sVmGVkXetjEsvw2L3NT")
	fmt.Println(reversed)
	fmt.Println(ReversedHash160ToNeo3Addr(reversed.String()))
}

func TestNeo3Sdk_Nep11OwnerOf(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11OwnerOf("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "278"))
}

func TestNeo3Sdk_Nep11Properties(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11PropertiesByInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "278"))
}

func TestNeo3Sdk_GetUrl(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11TokenUri("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "278"))
}

func TestNeo3Sdk_Nep11BalanceOf(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.Nep11BalanceOf("0xb13b57056775529e9461418a0a66b6dd97640ef8", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9"))

}

func TestNeo3Sdk_Nep11TokensOf(t *testing.T) {
	//assetHash := "0xafc8bb5d26f43980212da297d3316cf663917bab"
	//owner := "NbtRKN9wPV7WMAQFhVGsecDfNcHZcWkKoY"

	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	tokenLists, err := sdk.Nep11TokensOf("9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9", 0, 12)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range tokenLists {
		fmt.Println(v)
	}
	tokenLists, err = sdk.Nep11TokensOfWithBatchInvoke("9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range tokenLists {
		fmt.Println(v)
	}

	//p, _ := sdk.Nep11PropertiesByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", tokenLists)
	//a, _ := json.MarshalIndent(p, "", "	")
	//fmt.Println(string(a))

}
func TestNeo3Sdk_Nep11PropertiesByInvoke(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	res, _ := sdk.Nep11PropertiesByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", []string{"278", "279"})
	a, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(a))
}

func TestNeo3Sdk_Nep11UriByBatchInvoke(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	res, _ := sdk.Nep11UriByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", []string{"279"})
	a, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(a))
}