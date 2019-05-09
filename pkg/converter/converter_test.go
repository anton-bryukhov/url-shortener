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
		{64, "bc"},
		{512, "iq"},
		{1024, "qG"},
		{10000, "cLs"},
		{100000, "Aa4"},
		{1000000, "emjc"},
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
		{"bc", 64},
		{"iq", 512},
		{"qG", 1024},
		{"cLs", 10000},
		{"Aa4", 100000},
		{"emjc", 1000000},
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
