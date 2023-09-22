package debug

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	prefix = "ys"
)

func Pt(value string) {
	fmt.Println(SPt(value))
}

func PtT(value string) {
	fmt.Println(SPtT(value))
}

func SPtT(s string) string {
	sb := strings.Builder{}
	if prefix != "" {
		sb.WriteString(prefix)
		sb.WriteString(": ")
	}
	sb.WriteString("======= ")
	width := 20
	centeredStr := fmt.Sprintf("%*s", -width, fmt.Sprintf("%*s", (width+len(s))/2, s))
	sb.WriteString(centeredStr)
	sb.WriteString(" =======")
	return sb.String()
}

func PtV(key string, value any) {
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

func SPt(str string) string {
	lines := []string{}
	if prefix != "" {
		lines = append(lines, prefix, ": ")
	}
	lines = append(lines, str)
	return strings.Join(lines, "")
}

func SPtV(key string, value any) string {
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
			lines = append(lines, SPtV(k, SPtA(v)))
		}
	case int:
		lines = append(lines, strconv.Itoa(typedobj))
	case string:
		lines = append(lines, typedobj)
	}
	return strings.Join(lines, "\n")
}
