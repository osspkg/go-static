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
