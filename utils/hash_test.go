package utils

import (
	"fmt"
	"testing"
)

func TestBcryptHash(t *testing.T) {
	fmt.Println(BcryptHash("123456"))
}
