package chrono

import (
	"testing"
	"time"
)

func TestParseDateTime(t *testing.T) {
	type args struct {
		str string
		loc *time.Location
	}

	now := time.Now()

	result, err := ParseDateTime(now.In(time.UTC).Format(SQLDateTime), time.UTC)

	if err != nil {
		t.Error(err)
	}

	t.Logf("ParseDateTime %v", result)
}
