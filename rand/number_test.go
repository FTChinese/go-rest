package rand

import "testing"

func TestIntRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "100000 - 999999",
			args: args{
				min: 100000,
				max: 999999,
			},
		},
		{
			name: "1 - 3",
			args: args{
				min: 1,
				max: 3,
			},
		},
		{
			name: "10 - 15",
			args: args{
				min: 10,
				max: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntRange(tt.args.min, tt.args.max)

			t.Logf("A random number: %d", got)
		})
	}
}
