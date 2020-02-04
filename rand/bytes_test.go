package rand

import "testing"

func TestHex(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "GorillaCSRF token",
			args:    args{len: 32},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hex(tt.args.len)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Log(got)
		})
	}
}
