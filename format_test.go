package timef

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {

	for _, c := range []struct {
		str    string
		format string
	}{
		{
			str:    "2006-01-02T15:04:05+07:00",
			format: time.RFC3339,
		},
		{
			str:    "2006-01-02T15:04:05-07:00",
			format: time.RFC3339,
		},
	} {
		t.Run(fmt.Sprintf("converts '%s'", c.str), func(t *testing.T) {
			n, err := time.Parse(c.format, c.str)
			assert.NoError(t, err)

			f, err := Format(c.format, n)
			assert.NoError(t, err)

			parsed, err := time.Parse(c.format, f)
			assert.Nil(t, err)
			assert.Equal(t, n.Unix(), parsed.Unix())
		})
	}
}

func Benchmark_Format(b *testing.B) {
	t := time.Now()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Format(time.RFC3339, t)
	}
}

func Benchmark_FormatBytes(b *testing.B) {
	b.ResetTimer()

	t := time.Now()

	for n := 0; n < b.N; n++ {
		FormatBytes(time.RFC3339, t)
	}
}

func Benchmark_FormatRFC3339(b *testing.B) {
	t := time.Now()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		FormatRFC3339(t)
	}
}

func Benchmark_WriteRFC3339(b *testing.B) {
	t := time.Now()
	bs := make([]byte, RFC3339BufSize, RFC3339BufCap)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		WriteRFC3339(bs, t)
	}
}

func Benchmark_TimeFormat(b *testing.B) {
	t := time.Now()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		t.Format(time.RFC3339)
	}
}
