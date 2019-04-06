package miniq

import (
	"testing"
)

func TestTokenHasClass(t *testing.T) {
	table := []struct {
		input    string
		expected bool
	}{
		{
			input:    "p",
			expected: false,
		},
		{
			input:    "p#id",
			expected: false,
		},
		{
			input:    "p.class",
			expected: true,
		},
	}
	for _, row := range table {
		tk := &token{raw: row.input}
		received := tk.hasClass()
		if received != row.expected {
			t.Errorf("Received %t, expected %t", received, row.expected)
		}
	}
}

func TestTokenGetClassVal(t *testing.T) {
	table := []struct {
		input    string
		expected string
	}{
		{
			input:    "p.hello",
			expected: "hello",
		},
		{
			input:    "p.world",
			expected: "world",
		},
	}
	for _, row := range table {
		tk := &token{raw: row.input}
		received := tk.getClassVal()
		if received != row.expected {
			t.Errorf("Received %s, expected %s", received, row.expected)
		}
	}
}

func TestTokenGetData(t *testing.T) {
	table := []struct {
		input    string
		expected string
	}{
		{
			input:    "p",
			expected: "p",
		},
		{
			input:    "p.class",
			expected: "p",
		},
		{
			input:    "p#id",
			expected: "p",
		},
	}
	for _, row := range table {
		tk := &token{raw: row.input}
		received := tk.getData()
		if received != row.expected {
			t.Errorf("Received %s, expected %s", received, row.expected)
		}
	}
}

func TestTokenise(t *testing.T) {
	input := "ol.repo-list li.d-block"
	tokens := tokenise(input)

	if len(tokens) != 2 {
		t.Errorf("Length expected %d, %d received", 2, len(tokens))
	}
	if tokens[0].raw != "ol.repo-list" {
		t.Errorf("Received %s, expected %s", tokens[0].raw, "ol.repo-list")
	}
	if tokens[1].raw != "li.d-block" {
		t.Errorf("Received %s, expected %s", tokens[1].raw, "li.d-block")
	}
}
