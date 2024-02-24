package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const word = "Hello "

func BenchmarkConcat(t *testing.B) {
	t.ResetTimer()
	var str string
	for n := 0; n < t.N; n++ {
		str += word
	}
	t.StopTimer()
}

func BenchmarkSprintF(t *testing.B) {
	t.ResetTimer()
	var str string
	for n := 0; n < t.N; n++ {
		str = fmt.Sprintf("%v%v", str, word)
	}
	t.StopTimer()
}

func BenchmarkBuffer(t *testing.B) {
	t.ResetTimer()
	var buffer bytes.Buffer
	for n := 0; n < t.N; n++ {
		buffer.WriteString(word)
	}
	t.StopTimer()
}

func BenchmarkAppendStringArray(t *testing.B) {
	t.ResetTimer()
	data := make([]string, 0)
	for n := 0; n < t.N; n++ {
		data = append(data, word)
	}
	_ = strings.Join(data, "")
	t.StopTimer()
}

func BenchmarkCopy(t *testing.B) {
	t.ResetTimer()
	byteString := make([]byte, t.N)
	byteLength := 0

	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		byteLength += copy(byteString[byteLength:], word)
	}
	t.StopTimer()
}

func BenchmarkStringBuilder(t *testing.B) {
	t.ResetTimer()
	var strBuilder strings.Builder
	for n := 0; n < t.N; n++ {
		strBuilder.WriteString(word)
	}
	t.StopTimer()
}

func BenchmarkStringBuilderAndconvertToString(t *testing.B) {
	t.ResetTimer()
	var strBuilder strings.Builder
	for n := 0; n < t.N; n++ {
		strBuilder.WriteString(word)
	}
	_ = strBuilder.String()
	t.StopTimer()
}
