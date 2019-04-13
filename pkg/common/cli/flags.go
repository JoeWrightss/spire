package cli

import (
	"strings"
	"time"
)

// CommaStringsFlag facilitates parsing flags representing a comma separated list of strings
type CommaStringsFlag []string

func (f CommaStringsFlag) String() string {
	return strings.Join(f, ",")
}

func (f *CommaStringsFlag) Set(v string) error {
	*f = strings.Split(v, ",")
	return nil
}

// DurationFlag facilitates parsing flags representing a time.Duration
type DurationFlag time.Duration

func (f DurationFlag) String() string {
	return time.Duration(f).String()
}

func (f *DurationFlag) Set(v string) error {
	d, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*f = DurationFlag(d)
	return nil
}
