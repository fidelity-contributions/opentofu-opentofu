package static

import (
	"encoding/hex"
	"fmt"
	"github.com/opentofu/opentofu/internal/encryption/keyprovider"
)

type Config struct {
	Key string `hcl:"key"`
}

func (c Config) Build() (keyprovider.KeyProvider, error) {
	decodedData, err := hex.DecodeString(c.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to hex-decode the provided key (%w)", err)
	}
	return &staticKeyProvider{decodedData}, nil
}