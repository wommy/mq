package main

import "testing"

func TestParseMethodCall(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantMethod string
		wantArg    string
		wantOk     bool
	}{
		// Basic cases
		{"bare method", ".tree", "tree", "", true},
		{"method no arg", ".search", "search", "", true},

		// Double quotes
		{"double quotes", `.tree("full")`, "tree", "full", true},
		{"double quotes preview", `.tree("preview")`, "tree", "preview", true},
		{"double quotes search", `.search("term")`, "search", "term", true},

		// Single quotes (Windows-friendly)
		{"single quotes", `.tree('full')`, "tree", "full", true},
		{"single quotes search", `.search('test query')`, "search", "test query", true},

		// No quotes (Windows CMD may strip them)
		{"no quotes", ".tree(full)", "tree", "full", true},
		{"no quotes search", ".search(term)", "search", "term", true},

		// Edge cases
		{"empty arg with quotes", `.tree("")`, "tree", "", true},
		{"empty arg no quotes", ".tree()", "tree", "", true},
		{"arg with spaces", `.search("hello world")`, "search", "hello world", true},

		// Invalid cases
		{"no dot prefix", "tree", "", "", false},
		{"missing close paren", ".tree(full", "", "", false},
		{"random text", "invalid", "", "", false},
		{"empty string", "", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMethod, gotArg, gotOk := parseMethodCall(tt.input)
			if gotMethod != tt.wantMethod {
				t.Errorf("parseMethodCall(%q) method = %q, want %q", tt.input, gotMethod, tt.wantMethod)
			}
			if gotArg != tt.wantArg {
				t.Errorf("parseMethodCall(%q) arg = %q, want %q", tt.input, gotArg, tt.wantArg)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseMethodCall(%q) ok = %v, want %v", tt.input, gotOk, tt.wantOk)
			}
		})
	}
}
