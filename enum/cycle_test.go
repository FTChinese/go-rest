package enum

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func TestParseCycle(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Cycle
		wantErr bool
	}{
		{
			name: "Parse Cycle Month",
			args: args{
				name: "month",
			},
			want:    CycleMonth,
			wantErr: false,
		},
		{
			name: "Parse Cycle Year",
			args: args{
				name: "year",
			},
			want:    CycleYear,
			wantErr: false,
		},
		{
			name: "Parse Invalid Cycle",
			args: args{
				name: `null`,
			},
			want:    InvalidCycle,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCycle(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCycle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCycle_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Cycle Cycle `json:"cycle"`
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    fields
		wantErr bool
	}{
		{
			name: "Unmarshal Cycle Month",
			args: args{
				b: []byte(`{"cycle": "month"}`),
			},
			want: fields{
				Cycle: CycleMonth,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Cycle Year",
			args: args{
				b: []byte(`{"cycle": "year"}`),
			},
			want: fields{
				Cycle: CycleYear,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Null",
			args: args{
				b: []byte(`{"cycle": null}`),
			},
			want: fields{
				Cycle: InvalidCycle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got fields
			if err := json.Unmarshal(tt.args.b, &got); (err != nil) != tt.wantErr {
				t.Errorf("Cycle.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			log.Printf("got %+v", got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cycle.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCycle_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		c       Cycle
		want    []byte
		wantErr bool
	}{
		{
			name:    "Cycle Year",
			c:       CycleYear,
			want:    []byte(`"year"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cycle.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cycle.MarshalJSON() = %v, want %v", got, tt.want)
			}

			t.Logf("%s", got)
		})
	}
}

func TestCycle_Scan(t *testing.T) {
	type args struct {
		src interface{}
	}

	tests := []struct {
		name    string
		want    Cycle
		args    args
		wantErr bool
	}{
		{
			name:    "Scan Month",
			want:    CycleMonth,
			args:    args{src: []byte("month")},
			wantErr: false,
		},
		{
			name:    "Scan Nil",
			want:    InvalidCycle,
			args:    args{src: nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Cycle
			if err := c.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Cycle.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Cycle.Scan() = %v, want %v", c, tt.want)
			}
		})
	}
}
