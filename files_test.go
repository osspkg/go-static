/*
 *  Copyright (c) 2021-2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package static_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.osspkg.com/static"
)

func TestUnit_ReadDir(t *testing.T) {
	c := static.New()

	require.NoError(t, c.FromDir("./cmd"))
	require.Equal(t, c.List(), []string{"/static/main.go", "/static/main_test.go"})
}
