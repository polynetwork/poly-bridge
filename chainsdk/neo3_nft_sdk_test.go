package chainsdk

import (
	"fmt"
	"testing"
)

func TestNeo3Sdk_GetNFTTokenUri(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetNFTTokenUri("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "1601"))
}

func TestNeo3Sdk_GetAndCheckNFTUri(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetAndCheckNFTUri("", "0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9", "1601"))
}

func TestNeo3Sdk_GetNFTBalance(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetNFTBalance("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9"))
}

func TestNeo3Sdk_GetOwnerNFTsByIndex(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetOwnerNFTsByIndex("", "0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "NSAcLL5B85z6R1QdjsDxPw49Dw51zbq8J9", 4, 5))
}
