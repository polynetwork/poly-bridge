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
	fmt.Println(sdk.Nep11BalanceOf("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NbtRKN9wPV7WMAQFhVGsecDfNcHZcWkKoY"))

}

func TestNeo3Sdk_Nep11TokensOf(t *testing.T) {
	//assetHash := "0xafc8bb5d26f43980212da297d3316cf663917bab"
	//owner := "NbtRKN9wPV7WMAQFhVGsecDfNcHZcWkKoY"

	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	tokenLists, _ := sdk.Nep11TokensOf("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NVGUQ1qyL4SdSm7sVmGVkXetjEsvw2L3NT")
	for _, v := range tokenLists {
		fmt.Println(v)
	}
	p, _ := sdk.Nep11PropertiesByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", tokenLists)
	a, _ := json.MarshalIndent(p, "", "	")
	fmt.Println(string(a))

}
func TestNeo3Sdk_Nep11PropertiesByInvoke(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	res, _ := sdk.Nep11PropertiesByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", []string{"278", "279"})
	a, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(a))
}

func TestNeo3Sdk_Nep11UriByBatchInvoke(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	res, _ := sdk.Nep11UriByBatchInvoke("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", []string{"278", "279"})
	a, _ := json.MarshalIndent(res, "", "	")
	fmt.Println(string(a))
}
