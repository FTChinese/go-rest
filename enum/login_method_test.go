package enum

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func TestParseLoginMethod(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    LoginMethod
		wantErr bool
	}{
		{
			name: "Login with Email",
			args: args{
				name: "email",
			},
			want:    LoginMethodEmail,
			wantErr: false,
		},
		{
			name: "Login with Wechat",
			args: args{
				name: "wechat",
			},
			want:    LoginMethodWx,
			wantErr: false,
		},
		{
			name: "Login Method Unknown",
			args: args{
				name: "unknown",
			},
			want:    LoginMethodNull,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLoginMethod(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLoginMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseLoginMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginMethod_UnmarshalJSON(t *testing.T) {
	type fields struct {
		LoginMethod LoginMethod `json:"loginMethod"`
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    fields
		wantErr bool
	}{
		{
			name: "Unmarshal Wechat Login",
			args: args{
				data: `{"loginMethod": "wechat"}`,
			},
			want: fields{
				LoginMethod: LoginMethodWx,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Email Login",
			args: args{
				data: `{"loginMethod": "email"}`,
			},
			want: fields{
				LoginMethod: LoginMethodEmail,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Unknown",
			args: args{
				data: `{"loginMethod": "unknown"}`,
			},
			want: fields{
				LoginMethod: LoginMethodNull,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got fields
			err := json.Unmarshal([]byte(tt.args.data), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("LoginMethod.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			log.Printf("got %+v", got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginMethod.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginMethod_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		x       LoginMethod
		want    []byte
		wantErr bool
	}{
		{
			name:    "Marshal LoginMethodWx",
			x:       LoginMethodWx,
			want:    []byte(`"wechat"`),
			wantErr: false,
		},
		{
			name:    "Marshal LoginMethodEmail",
			x:       LoginMethodEmail,
			want:    []byte(`"email"`),
			wantErr: false,
		},
		{
			name:    "Marshal Invalid",
			x:       LoginMethodNull,
			want:    []byte(`null`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.x.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginMethod.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginMethod.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
