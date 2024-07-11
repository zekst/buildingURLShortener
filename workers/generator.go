package workers

import (

	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"github.com/itchyny/base58-go"
)


func utilSha256(inputString string) [] byte{

	algorithm := sha256.New()

	algorithm.Write([]byte(inputString))

	return  algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {

	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)

	if err!=nil{
		fmt.Printf("Something happened %s", err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func generateShortURL(initialLink string, userID string) string{

	urlHash := utilSha256(initialLink+userID)

	generatedNumber := new(big.Int).SetBytes(urlHash).Uint64()

	finalURLString := base58Encoded([] byte(fmt.Sprintf("%d", generatedNumber)))

	return finalURLString
}
