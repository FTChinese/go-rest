package enum

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseGender(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Gender
		wantErr bool
	}{
		{
			name:    "Parse Male",
			args:    args{name: "M"},
			want:    GenderMale,
			wantErr: false,
		},
		{
			name:    "Parse Female",
			args:    args{name: "F"},
			want:    GenderFemale,
			wantErr: false,
		},
		{
			name:    "Parse Null",
			args:    args{name: `null`},
			want:    InvalidGender,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGender(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseGender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGender_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Gender Gender `json:"gender"`
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
			name: "Unmarshal Male",
			args: args{
				b: []byte(`{"gender": "M"}`),
			},
			want: fields{
				Gender: GenderMale,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Female",
			args: args{
				b: []byte(`{"gender": "F"}`),
			},
			want: fields{
				Gender: GenderFemale,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Null",
			args: args{
				b: []byte(`{"gender": null}`),
			},
			want: fields{
				Gender: InvalidGender,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got fields
			if err := json.Unmarshal(tt.args.b, &got); (err != nil) != tt.wantErr {
				t.Errorf("Gender.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gender.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
