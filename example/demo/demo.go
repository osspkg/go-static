/*
 *  Copyright (c) 2021-2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package demo

import (
	"fmt"

	"go.osspkg.com/static"
)

//go:generate static . ui

var ui static.Reader

func run() {
	fmt.Println(ui.List())
}
