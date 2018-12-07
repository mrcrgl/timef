package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/mrcrgl/timef"
)

func main() {
	toString()
	toBytes()
	writeAtBytes()
}

func toString() {
	f, _ := timef.Format(time.RFC3339, time.Now())

	fmt.Println(f)
	// Output: "2018-12-07T14:21:09+01:00"
}

func toBytes() {
	b := new(bytes.Buffer)
	b.WriteString("T: ")

	n, _ := b.Write(timef.FormatRFC3339(time.Now()))

	b.WriteString(fmt.Sprintf(" is %d bytes long", n))

	fmt.Println(b.String())
	// Output: "T: 2018-12-07T14:21:09+01:00 is 25 bytes long"
}

func writeAtBytes() {
	b := make([]byte, 64)

	var pos int
	off := copy(b, []byte("T: "))
	pos += off

	n, _ := timef.WriteRFC3339At(time.Now(), b, int64(off))
	pos += n

	n = copy(b[n+off:], []byte(fmt.Sprintf(" is %d bytes long", n)))
	pos += n

	fmt.Println(string(b[0:pos]))
	// Output: "T: 2018-12-07T14:31:10+01:00 is 25 bytes long"
}
