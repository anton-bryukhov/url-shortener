package converter_test

import (
	"fmt"
	. "github.com/anton-bryukhov/url-shortener/pkg/converter"
	"testing"
)

func TestConvertToURL(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{64, "ab"},
		{512, "hp"},
		{1024, "pF"},
		{2048, "Gb"},
		{4096, "add"},
		{8192, "bhh"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Converts '%d' to '%s'", test.in, test.out), func(t *testing.T) {
			got := ConvertToURL(test.in)

			if got != test.out {
				t.Fatalf("got: '%s', want: '%s'", got, test.out)
			}
		})
	}
}

func TestConvertToID(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"ab", 64},
		{"hp", 512},
		{"pF", 1024},
		{"Gb", 2048},
		{"add", 4096},
		{"bhh", 8192},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Converts '%s' to '%d'", test.in, test.out), func(t *testing.T) {
			got := ConvertToID(test.in)

			if got != test.out {
				t.Fatalf("got: '%d', want: '%d'", got, test.out)
			}
		})
	}
}
