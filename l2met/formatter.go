package l2met

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Formatter is usable as a logrus formatter a la:
//
//   log := logrus.New()
//   log.Formatter = new(l2met.Formatter)
//
type Formatter struct {
	prefix string
}

// NewFormatter returns a *Formatter yey!
func NewFormatter(prefix string) *Formatter {
	return &Formatter{
		prefix: prefix,
	}
}

// Format formats a *logrus.Entry whee hoo
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	formatted := []string{}

	if len(f.prefix) > 0 {
		formatted = []string{f.prefix}
	}

	_, filename, _, _ := runtime.Caller(4)
	if len(filename) > 0 {
		entry.Data["filename"] = filename
	}

	for key, value := range entry.Data {
		formatted = append(formatted, makeKeyValue(key, fmt.Sprintf("%v", value)))
	}

	return []byte(strings.Join(formatted, " ") + "\n"), nil
}

func makeKeyValue(key, value string) string {
	if strings.ContainsRune(value, ' ') {
		value = "\"" + value + "\""
	}

	return key + "=" + value
}
