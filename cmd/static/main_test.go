package main

import (
	"os"
	"testing"
)

func TestUnit_Regexp(t *testing.T) {
	out := validRex.ReplaceAllString(`aaaa-bbb*25`, `_`)
	if out != `aaaa_bbb_25` {
		t.Fatal(out)
	}
}

func TestUnit_extractFromFiles(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	pack, err := extractFromFiles(path)
	if err != nil {
		t.Fatal(err)
	}
	if pack != "main" {
		t.Fatal(pack)
	}
}
