/*
 *  Copyright (c) 2021-2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

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
