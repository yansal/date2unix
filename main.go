package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(time.Now().Unix())
		return
	}

	layouts := []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02 15:04:05.999999999 -0700",
		"2006-01-02 15:04:05.999999999",
		"2006-01-02 15:04:05",
	}

	date := strings.Join(os.Args[1:], " ")
	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err != nil {
			continue
		}
		fmt.Println(t.Unix())
		return
	}

	fmt.Fprintf(os.Stderr, "can't parse %s\n", date)
	os.Exit(1)
}
