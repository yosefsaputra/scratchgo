package debug

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	prefix = "ys"
)

func Pt(key string, value any) {
	line := ""
	if prefix != "" {
		line = prefix + ": "
	}
	line = line + fmt.Sprintf("%s=%s", key, value)
	fmt.Println(line)
}

func PtA(obj any) {
	fmt.Println(SPtA(obj))
}

func SPt(key string, value any) string {
	lines := []string{}
	if prefix != "" {
		lines = append(lines, prefix, ": ")
	}
	if key != "" {
		lines = append(lines, key, "=")
	}
	lines = append(lines, fmt.Sprintf("\"%s\"", value))
	return strings.Join(lines, "")
}

func SPtA(obj any) string {
	lines := []string{}
	switch typedobj := obj.(type) {
	case map[string]any:
		for k, v := range typedobj {
			lines = append(lines, SPt(k, SPtA(v)))
		}
	case int:
		lines = append(lines, strconv.Itoa(typedobj))
	case string:
		lines = append(lines, typedobj)
	}
	return strings.Join(lines, "\n")
}
