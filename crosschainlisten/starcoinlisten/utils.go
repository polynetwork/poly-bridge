package starcoinlisten

import (
	"encoding/hex"
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"math/big"
	"strings"
)

func HexToBytes(str string) ([]byte, error) {
	if !strings.HasPrefix(str, "0x") {
		return hex.DecodeString(str[:])
	}
	return hex.DecodeString(str[2:])
}

func GetTokenCodeString(tc *TokenCode) string {
	return "0x" + hex.EncodeToString(tc.Address[:]) + "::" + tc.Module + "::" + tc.Name
}

func Uint128ToBigInt(u *serde.Uint128) *big.Int {
	h := new(big.Int).SetUint64(u.High)
	l := new(big.Int).SetUint64(u.Low)
	return new(big.Int).SetBytes(append(h.Bytes(), l.Bytes()...))
}

func BcsDeserializeCrossChainEvent(input []byte) (CrossChainEvent, error) {
	if input == nil {
		var obj CrossChainEvent
		return obj, fmt.Errorf("cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeCrossChainEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("some input bytes were not read")
	}
	return obj, err
}

func DeserializeCrossChainEvent(deserializer serde.Deserializer) (CrossChainEvent, error) {
	var obj CrossChainEvent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Sender = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.TxId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ProxyOrAssetContract = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ToChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToContract = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.RawData = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeLockEvent(input []byte) (LockEvent, error) {
	if input == nil {
		var obj LockEvent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeLockEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

func DeserializeLockEvent(deserializer serde.Deserializer) (LockEvent, error) {
	var obj LockEvent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeTokenCode(deserializer); err == nil {
		obj.FromAssetHash = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.FromAddress = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ToChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToAssetHash = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToAddress = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj.Amount = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func DeserializeTokenCode(deserializer serde.Deserializer) (TokenCode, error) {
	var obj TokenCode
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Address = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Module = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func DeserializeAccountAddress(deserializer serde.Deserializer) (AccountAddress, error) {
	var obj [16]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (AccountAddress)(obj), err
	}
	if val, err := deserialize_array16_u8_array(deserializer); err == nil {
		obj = val
	} else {
		return ((AccountAddress)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (AccountAddress)(obj), nil
}

func deserialize_array16_u8_array(deserializer serde.Deserializer) ([16]uint8, error) {
	var obj [16]uint8
	for i := range obj {
		if val, err := deserializer.DeserializeU8(); err == nil {
			obj[i] = val
		} else {
			return obj, err
		}
	}
	return obj, nil
}

func BcsDeserializeCrossChainFeeLockEvent(input []byte) (CrossChainFeeLockEvent, error) {
	if input == nil {
		var obj CrossChainFeeLockEvent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeCrossChainFeeLockEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

func DeserializeCrossChainFeeLockEvent(deserializer serde.Deserializer) (CrossChainFeeLockEvent, error) {
	var obj CrossChainFeeLockEvent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeTokenCode(deserializer); err == nil {
		obj.FromAssetHash = val
	} else {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Sender = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ToChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToAddress = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj.Net = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj.Fee = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj.Id = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeVerifyHeaderAndExecuteTxEvent(input []byte) (VerifyHeaderAndExecuteTxEvent, error) {
	if input == nil {
		var obj VerifyHeaderAndExecuteTxEvent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeVerifyHeaderAndExecuteTxEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

func DeserializeVerifyHeaderAndExecuteTxEvent(deserializer serde.Deserializer) (VerifyHeaderAndExecuteTxEvent, error) {
	var obj VerifyHeaderAndExecuteTxEvent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.FromChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToContract = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.CrossChainTxHash = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.FromChainTxHash = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeUnlockEvent(input []byte) (UnlockEvent, error) {
	if input == nil {
		var obj UnlockEvent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeUnlockEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

func DeserializeUnlockEvent(deserializer serde.Deserializer) (UnlockEvent, error) {
	var obj UnlockEvent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToAssetHash = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToAddress = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj.Amount = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}
