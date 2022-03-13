package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var messageTests = []struct {
	lang     string
	key      string
	expected string
}{
	{"id", "empty_state", "Data kosong"},
	{"id", "empty_notfound", "Data tidak ditemukan"},
	{"en", "empty_state", "Data is empty"},
	{"en", "empty_notfound", "Data is not found"},
}

var messageFormattedTests = []struct {
	lang        string
	key         string
	replacement string
	expected    string
}{
	{"id", "empty_state_formatted", "Task", "Task kosong"},
	{"id", "empty_notfound_formatted", "Task", "Task tidak ditemukan"},
	{"en", "empty_state_formatted", "Task", "Task is empty"},
	{"en", "empty_notfound_formatted", "Task", "Task is not found"},
}

func TestMessage(t *testing.T) {
	for _, test := range messageTests {
		os.Setenv("APP_LANGUAGE", test.lang)
		if err := assert.Equal(t, test.expected, GetMessage(test.key)); !err {
			t.Errorf("Expected %q, result %q", test.expected, GetMessage(test.key))
		}
	}
}

func TestFormattedMessage(t *testing.T) {
	for _, test := range messageFormattedTests {
		os.Setenv("APP_LANGUAGE", test.lang)
		if err := assert.Equal(t, test.expected, GetFormattedMessage(test.key, test.replacement)); !err {
			t.Errorf("Expected %q, result %q", test.expected, GetFormattedMessage(test.key, test.replacement))
		}
	}
}
