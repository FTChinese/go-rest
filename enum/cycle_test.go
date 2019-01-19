package enum

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCycleJSON(t *testing.T) {
	data := struct {
		Invalid Cycle
		Month   Cycle
		Year    Cycle
	}{
		Invalid: InvalidCycle,
		Month:   CycleMonth,
		Year:    CycleYear,
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s", b)
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
