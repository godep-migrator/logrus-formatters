package l2met

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

var formatTests = []struct {
	in  logrus.Fields
	out string
}{
	{logrus.Fields{"foo": "bar", "fizz": "buzz"}, "fizz=buzz foo=bar\n"},
	{logrus.Fields{"ham": 1, "derp": 99.9}, "derp=99.9 ham=1\n"},
	{logrus.Fields{"wat": "the butt"}, "wat=\"the butt\"\n"},
}

func TestNewFormatter(t *testing.T) {
	f := NewFormatter("foo")
	if f.prefix != "foo" {
		t.Fail()
	}
}

func TestFormat(t *testing.T) {
	formatter := NewFormatter("")
	entry := logrus.NewEntry(nil)

	for i, tt := range formatTests {
		entry.Data = tt.in
		fb, err := formatter.Format(entry)

		if err != nil {
			t.Error(err)
			continue
		}

		actual := string(fb)
		if actual != tt.out {
			t.Errorf("%d. formatter.Format(%v) => %q, want %q",
				i, entry, actual, tt.out)
		}
	}
}

func TestPrefixedFormat(t *testing.T) {
	formatter := NewFormatter("app")
	entry := logrus.NewEntry(nil)

	for i, tt := range formatTests {
		entry.Data = tt.in
		fb, err := formatter.Format(entry)

		if err != nil {
			t.Error(err)
			continue
		}

		expected := "app " + tt.out

		actual := string(fb)
		if actual != expected {
			t.Errorf("%d. formatter.Format(%v) => %q, want %q",
				i, entry, actual, expected)
		}
	}
}
