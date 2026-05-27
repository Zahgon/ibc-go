package mock

type TestAddressCodec struct{}

func (t TestAddressCodec) StringToBytes(text string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TestAddressCodec) BytesToString(bz []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
