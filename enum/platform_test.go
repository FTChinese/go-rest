package enum

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParsePlatform(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Platform
		wantErr bool
	}{
		{
			name:    "Parse Android",
			args:    args{"android"},
			want:    PlatformAndroid,
			wantErr: false,
		},
		{
			name:    "Parse iOS",
			args:    args{"ios"},
			want:    PlatformIOS,
			wantErr: false,
		},
		{
			name:    "Parse Web",
			args:    args{"web"},
			want:    PlatformWeb,
			wantErr: false,
		},
		{
			name:    "Parse Null",
			args:    args{`null`},
			want:    PlatformNull,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePlatform(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePlatform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParsePlatform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlatform_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Platform Platform `json:"platform"`
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
			name:    "Unmarshal Android",
			args:    args{b: []byte(`{"platform": "android"}`)},
			want:    fields{Platform: PlatformAndroid},
			wantErr: false,
		},
		{
			name:    "Unmarshal iOS",
			args:    args{b: []byte(`{"platform": "ios"}`)},
			want:    fields{Platform: PlatformIOS},
			wantErr: false,
		},
		{
			name:    "Unmarshal Web",
			args:    args{b: []byte(`{"platform": "web"}`)},
			want:    fields{Platform: PlatformWeb},
			wantErr: false,
		},
		{
			name:    "Unmarshal Null",
			args:    args{b: []byte(`{"platform": null}`)},
			want:    fields{Platform: PlatformNull},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got fields
			if err := json.Unmarshal(tt.args.b, &got); (err != nil) != tt.wantErr {
				t.Errorf("Platform.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Platform.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlatform_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		x       Platform
		want    []byte
		wantErr bool
	}{
		{
			name:    "Marshal Android",
			x:       PlatformAndroid,
			want:    []byte(`"android"`),
			wantErr: false,
		},
		{
			name:    "Marshal iOS",
			x:       PlatformIOS,
			want:    []byte(`"ios"`),
			wantErr: false,
		},
		{
			name:    "Marshal Web",
			x:       PlatformWeb,
			want:    []byte(`"web"`),
			wantErr: false,
		},
		{
			name:    "Marshal Null",
			x:       PlatformNull,
			want:    []byte(`null`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.x.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Platform.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Platform.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
