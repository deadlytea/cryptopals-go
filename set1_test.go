package cryptopals

import (
	"strings"
	"testing"
)

func TestSet1Challenge1(t *testing.T) {

	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	output, err := Set1Challenge1(input)

	if err != nil {
		t.Error(err)
	}

	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expected: %s, got: %s", expectedOutput, output)
		t.Fail()
	}
}

func TestSet1Challenge2(t *testing.T) {

	input := "1c0111001f010100061a024b53535009181c"
	xorModifier := "686974207468652062756c6c277320657965"

	output, err := Set1Challenge2(input, xorModifier)
	if err != nil {
		t.Error(err)
	}

	expectedOutput := "746865206b696420646f6e277420706c6179"
	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expected: %s, got: %s", expectedOutput, output)
		t.Fail()
	}
}
