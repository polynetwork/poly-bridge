package abi

//func UnpackMethod(ab abi.ABI, method string, data interface{}, payload []byte) error {
//	if len(payload) < 4 || len(payload[4:])%32 != 0 {
//		return fmt.Errorf("invalid payload")
//	}
//	if reflect.TypeOf(data).Kind() != reflect.Ptr {
//		return fmt.Errorf("interface should be pointer")
//	}
//
//	mth, ok := ab.Methods[method]
//	if !ok {
//		return fmt.Errorf("method [%s] not exist", method)
//	}
//	return mth.Inputs.Unpack(data, payload[4:])
//}
