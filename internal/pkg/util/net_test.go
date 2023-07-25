package util

import "testing"

func TestPortAvailable(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check 8000",
			args: args{
				port: "8000",
			},
			want: false,
		},
		{
			name: "check 62256",
			args: args{
				port: "8000",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PortAvailable(tt.args.port); got != tt.want {
				t.Errorf("PortAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
