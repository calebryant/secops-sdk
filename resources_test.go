package chronicle_test

import (
	"testing"

	chronicle "github.com/calebryant/secops-sdk"
	"github.com/stretchr/testify/assert"
)

const (
	project  string = "testproject"
	location string = "us"
	instance string = "12345"
	logtype  string = "WINEVTLOG"
	parser   string = "12345"
)

func TestPaths(t *testing.T) {
	tt := []struct {
		name     string
		expected string
		actual   string
	}{
		{
			name:     "list log types",
			expected: "projects/testproject/locations/us/instances/12345/logTypes",
			actual:   chronicle.NewListLogTypesRequest(project, location, instance).Path,
		},
		{
			name:     "run parser",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG:runParser",
			actual:   chronicle.NewRunParserRequest(project, location, instance, logtype, nil, nil, nil, false).Path,
		},
		{
			name:     "activate parser",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parsers/12345:activate",
			actual:   chronicle.NewActivateParserRequest(project, location, instance, logtype, parser).Path,
		},
		{
			name:     "create parser",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parsers",
			actual:   chronicle.NewCreateParserRequest(project, location, instance, logtype, nil, false).Path,
		},
		{
			name:     "deactivate parser",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parsers/12345:deactivate",
			actual:   chronicle.NewDeactivateParserRequest(project, location, instance, logtype, parser).Path,
		},
		{
			name:     "delete parser",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parsers/12345",
			actual:   chronicle.NewDeleteParserRequest(project, location, instance, logtype, parser, false).Path,
		},
		{
			name:     "list parsers",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parsers",
			actual:   chronicle.NewListParsersRequest(project, location, instance, logtype).Path,
		},
		{
			name:     "get parser",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parsers/12345",
			actual:   chronicle.NewGetParserRequest(project, location, instance, logtype, parser).Path,
		},
		{
			name:     "activate parser extension",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parserExtensions/12345:activate",
			actual:   chronicle.NewActivateParserExtensionRequest(project, location, instance, logtype, parser).Path,
		},
		{
			name:     "create parser extension",
			expected: "projects/testproject/locations/us/instances/12345/logTypes/WINEVTLOG/parserExtensions",
			actual:   chronicle.NewCreateParserExtensionRequest(project, location, instance, logtype, parser).Path,
		},
	}
	for _, testcase := range tt {
		assert.Equal(t, testcase.expected, testcase.actual, testcase.name)
	}
}

func TestRunParser(t *testing.T) {
	tt := []struct {
		name               string
		expectedcbn        string
		expectedcbnsnippet string
		expectedlogs       []string
		request            *chronicle.RunParserRequest
	}{
		{
			name:               "empty Extension empty logs",
			expectedcbn:        "bXkgY2JuIGZpbGU=",
			expectedcbnsnippet: "",
			expectedlogs:       []string{},
			request:            chronicle.NewRunParserRequest(project, location, instance, logtype, []byte("my cbn file"), []byte(""), []string{""}, false),
		},
		{
			name:               "nil Extension nil logs",
			expectedcbn:        "bXkgY2JuIGZpbGU=",
			expectedcbnsnippet: "",
			expectedlogs:       []string{},
			request:            chronicle.NewRunParserRequest(project, location, instance, logtype, []byte("my cbn file"), nil, nil, false),
		},
		{
			name:               "With Extension",
			expectedcbn:        "bXkgY2JuIGZpbGU=",
			expectedcbnsnippet: "bXkgY2JuIHNuaXBwZXQgZmlsZQ==",
			expectedlogs:       []string{"bXkgbG9nIG51bWJlciAx", "bG9nIG51bWJlciAy"},
			request:            chronicle.NewRunParserRequest(project, location, instance, logtype, []byte("my cbn file"), []byte("my cbn snippet file"), []string{"my log number 1", "log number 2"}, false),
		},
	}
	for _, testcase := range tt {
		assert.Equal(t, testcase.expectedcbn, testcase.request.Cbn, testcase.name)
		assert.Equal(t, testcase.expectedcbnsnippet, testcase.request.CbnSnippet, testcase.name)
		assert.Equal(t, testcase.expectedlogs, testcase.request.Log, testcase.name)
	}
}

func TestCreateParser(t *testing.T) {
	tt := []struct {
		name        string
		expectedcbn string
		request     *chronicle.CreateParserRequest
	}{
		{
			name:        "empty Extension empty logs",
			expectedcbn: "bXkgY2JuIGZpbGU=",
			request:     chronicle.NewCreateParserRequest(project, location, instance, logtype, []byte("my cbn file"), false),
		},
	}
	for _, testcase := range tt {
		assert.Equal(t, testcase.expectedcbn, testcase.request.Cbn, testcase.name)
	}
}
