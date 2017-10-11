package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func hex2Bin(hexSlice []byte) []byte {
	binSlice, err := hex.DecodeString(string(hexSlice))
	if err != nil {
		panic(err)
	}
	return binSlice
}

func bin2Hex(binSlice []byte) []byte {
	return []byte(hex.EncodeToString(binSlice))
}

func bin2Base64(binSlice []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(binSlice))
}

func xorFixedLength(binSlice []byte, xorSlice []byte) ([]byte, error) {
	binLen := len(binSlice)
	xorLen := len(xorSlice)
	if len(binSlice) != len(xorSlice) {
		return nil, fmt.Errorf("error: Expected bin slice length (%d) to match xor length (%d)", binLen, xorLen)
	}
	var outputSlice []byte
	for i, b := range binSlice {
		outputSlice = append(outputSlice, b^xorSlice[i])
	}
	return outputSlice, nil
}

func xorSingleByte(binSlice []byte, xorByte byte) []byte {
	var outputSlice []byte
	for _, b := range binSlice {
		outputSlice = append(outputSlice, b^xorByte)
	}
	return outputSlice
}

// Set1Challenge1 converts a hex string to base64
func Set1Challenge1(input string) (string, error) {

	binSlice := hex2Bin([]byte(input))
	base64Slice := bin2Base64(binSlice)

	return string(base64Slice), nil
}

// Set1Challenge2 performs an XOR operation on two equal length HEX strings
func Set1Challenge2(input string, xor string) (string, error) {

	inputBin := hex2Bin([]byte(input))
	xorBin := hex2Bin([]byte(xor))
	binSlice, err := xorFixedLength(inputBin, xorBin)

	if err != nil {
		return "", err
	}

	outputString := string(bin2Hex(binSlice))

	return string(outputString), nil
}
