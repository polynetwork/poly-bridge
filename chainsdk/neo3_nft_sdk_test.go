package chainsdk

import (
	"fmt"
	"testing"
)

func TestConvertTokenIdFromHexStr2IntStr(t *testing.T) {
	fmt.Println(ConvertTokenIdFromHexStr2IntStr("1601"))
}

func TestConvertTokenIdFromIntStr2HexStr(t *testing.T) {
	fmt.Println(ConvertTokenIdFromIntStr2HexStr("278"))
}
func TestNeo3Sdk_GetNFTTokenUri(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetNFTTokenUri("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "278"))
}

func TestNeo3Sdk_GetAndCheckNFTUri(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetAndCheckNFTUri("", "0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "4495966d7cd6726cd04d50996a421ec2dd6ee3fe", "278"))
}

func TestNeo3Sdk_GetNFTBalance(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetNFTBalance("0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "4495966d7cd6726cd04d50996a421ec2dd6ee3fe"))
}

func TestNeo3Sdk_GetOwnerNFTsByIndex(t *testing.T) {
	sdk := NewNeo3Sdk("http://seed1t5.neo.org:20332")
	fmt.Println(sdk.GetOwnerNFTsByIndex("", "0x9f344fe24c963d70f5dcf0cfdeb536dc9c0acb3a", "4495966d7cd6726cd04d50996a421ec2dd6ee3fe", 0, 5))
}
