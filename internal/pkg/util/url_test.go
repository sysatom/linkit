package util

import "testing"

func TestFillScheme(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "http-1",
			args: args{
				host: "127.0.0.1:6060",
			},
			want: "http://127.0.0.1:6060",
		},
		{
			name: "http-2",
			args: args{
				host: "192.168.0.1:6060",
			},
			want: "http://192.168.0.1:6060",
		},
		{
			name: "https",
			args: args{
				host: "github.com",
			},
			want: "https://github.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FillScheme(tt.args.host); got != tt.want {
				t.Errorf("FillScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}
