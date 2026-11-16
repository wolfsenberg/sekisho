package vault

func (v *Vault) SealSetting(name string, plaintext []byte) ([]byte, error) {
	if name == "" {
		return nil, ErrDecrypt
	}
	return v.seal("sekisho:setting:"+name, plaintext)
}

func (v *Vault) OpenSetting(name string, encrypted []byte) ([]byte, error) {
	if name == "" {
		return nil, ErrDecrypt
	}
	return v.open("sekisho:setting:"+name, encrypted)
}
