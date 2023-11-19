package oqs

import (
	"fmt"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func GenerateKey() {
	fmt.Println("Generate Key")
	var liboqsVersion = oqs.LiboqsVersion()
	fmt.Println("liboqs version: ", liboqsVersion)
	// oqs.EnabledKEMs()
}
