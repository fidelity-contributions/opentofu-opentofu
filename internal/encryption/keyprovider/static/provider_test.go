package static_test

import (
	"testing"

	"github.com/opentofu/opentofu/internal/encryption/keyprovider/static"
)

func TestEmpty(t *testing.T) {
	descriptor := static.New()
	config := descriptor.ConfigStruct()
	keyProvider, err := config.Build()
	if err != nil {
		panic(err)
	}
	data, err := keyProvider.Provide()
	if err != nil {
		t.Fatalf("unespected error: %v", err)
	}
	if len(data) != 0 {
		t.Fatalf("unexpected key output: %v", data)
	}
}

// TestInvalidInput tests if an error is throw with an invalid input.
func TestInvalidInput(t *testing.T) {
	descriptor := static.New()
	config := descriptor.ConfigStruct().(*static.Config)
	config.Key = "G"
	_, err := config.Build()
	if err == nil {
		t.Fatalf("unexpected success")
	}
}

func TestSuccess(t *testing.T) {
	descriptor := static.New()
	config := descriptor.ConfigStruct().(*static.Config)
	config.Key = "48656c6c6f20776f726c6421"
	keyProvider, err := config.Build()
	if err != nil {
		panic(err)
	}
	data, err := keyProvider.Provide()
	if err != nil {
		t.Fatalf("unespected error: %v", err)
	}
	if string(data) != "Hello world!" {
		t.Fatalf("unexpected key output: %v", data)
	}
}