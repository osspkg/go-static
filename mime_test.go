/*
 *  Copyright (c) 2021-2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package static

import "testing"

func TestDetectContentType(t *testing.T) {
	type args struct {
		filename string
		b        []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "stype.css",
			args: args{
				filename: "/style.css",
				b:        []byte(`import ...`),
			},
			want: "text/css",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectContentType(tt.args.filename, tt.args.b); got != tt.want {
				t.Errorf("DetectContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}
