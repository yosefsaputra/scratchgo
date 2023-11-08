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

func SPt(str string) string {
	lines := []string{}
	if prefix != "" {
		lines = append(lines, prefix, ": ")
	}
	lines = append(lines, str)
	return strings.Join(lines, "")
}

// Print Title
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

// Print Variable
func PtV(key string, value any) {
	line := ""
	if prefix != "" {
		line = prefix + ": "
	}
	line = line + fmt.Sprintf("%s=%s", key, value)
	line = strings.ReplaceAll(line, "\n", " ")
	fmt.Println(line)
}

func SPtV(key string, value any) string {
	sb := strings.Builder{}
	if prefix != "" {
		sb.WriteString(prefix)
		sb.WriteString(": ")
	}
	if key != "" {
		sb.WriteString(key)
		sb.WriteString("=")
	}
	sb.WriteString(fmt.Sprintf("\"%s\"", value))
	return sb.String()
}

// Print Any
func PtA(obj any) {
	fmt.Println(SPtA(obj))
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

func Ptln() {
	fmt.Println()
}
