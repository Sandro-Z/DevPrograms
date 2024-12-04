package qqguild

import (
	"strings"
	"time"
)

type Milliseconds float64

func DurationToMilliseconds(dura time.Duration) Milliseconds {
	return Milliseconds(dura.Milliseconds())
}

func (ms Milliseconds) String() string {
	return ms.Duration().String()
}

func (ms Milliseconds) Duration() time.Duration {
	const f64ms = Milliseconds(time.Millisecond)
	return time.Duration(ms * f64ms)
}

type Timestamp time.Time

const TimestampFormat = time.RFC3339 // same as ISO8601

func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(t)
}

func NowTimestamp() Timestamp {
	return NewTimestamp(time.Now())
}

func (t *Timestamp) UnmarshalJSON(v []byte) error {
	str := strings.Trim(string(v), `"`)
	if str == "null" {
		*t = Timestamp{}
		return nil
	}

	r, err := time.Parse(TimestampFormat, str)
	if err != nil {
		return err
	}

	*t = Timestamp(r)
	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	if !t.IsValid() {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Format(TimestampFormat) + `"`), nil
}

func (t Timestamp) IsValid() bool {
	return !t.Time().IsZero()
}

func (t Timestamp) Format(fmt string) string {
	return t.Time().Format(fmt)
}

func (t Timestamp) Time() time.Time {
	return time.Time(t)
}
