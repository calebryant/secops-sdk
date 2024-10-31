package chronicle

import (
	"encoding/base64"
	"fmt"
)

const (
	instancespath         string = "projects/%s/locations/%s/instances"
	logtypespath          string = instancespath + "/%s/logTypes"
	logspath              string = logtypespath + "/%s/logs"
	parserspath           string = logtypespath + "/%s/parsers"
	parsersextensionspath string = logtypespath + "/%s/parserExtensions"
)

// Resource: Instances

type Instance struct {
	Name string `json:"name,omitempty"`
}

// Resource: LogTypes

type LogType struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Golden      bool   `json:"golden,omitempty"`
}

// Method: list

type ListLogTypesRequest struct {
	Path      string
	PageSize  int
	PageToken string
}

func NewListLogTypesRequest(project, location, instance string) *ListLogTypesRequest {
	return &ListLogTypesRequest{
		Path: fmt.Sprintf(logtypespath, project, location, instance),
	}
}

type ListLogTypesResult struct {
	LogTypes []LogType `json:"logTypes"`
}

// Method: runParser

type RunParserRequest struct {
	Path             string
	Cbn              string
	CbnSnippet       string
	Log              []string
	StatedumpAllowed bool
}

func NewRunParserRequest(project, location, instance, logtype string, cbn, cbnsnippet []byte, logs []string, statedumpallowed bool) *RunParserRequest {
	encodedlogs := []string{}
	for _, log := range logs {
		if log == "" {
			continue
		}
		encodedlogs = append(encodedlogs, base64.StdEncoding.EncodeToString([]byte(log)))
	}
	return &RunParserRequest{
		Path:             fmt.Sprintf(logtypespath, project, location, instance) + "/" + logtype + ":runParser",
		Cbn:              base64.StdEncoding.EncodeToString(cbn),
		CbnSnippet:       base64.StdEncoding.EncodeToString(cbnsnippet),
		Log:              encodedlogs,
		StatedumpAllowed: statedumpallowed,
	}
}

type RunParserResult struct {
	RunParserResults []ParserLogResult `json:"runParserResults"`
}

type ParserLogResult struct {
	Log              string            `json:"log,omitempty"`
	StatedumpResults []StatedumpResult `json:"statedumpResults,omitempty"`
	ParsedEvents     ParsedEvents      `json:"parsedEvents,omitempty"`
	Error            Status            `json:",omitempty"`
}

type StatedumpResult struct {
	Label           string `json:"label,omitempty"`
	StatedumpResult string `json:"statedumpResult,omitempty"`
}

type ParsedEvents struct {
	Events []ParsedEvent `json:"events,omitempty"`
}

type ParsedEvent struct {
	Event  map[string]any `json:"event,omitempty"`
	Entity map[string]any `json:"entity,omitempty"`
}

type Status struct {
	Code    int                 `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Details []map[string]string `json:"details,omitempty"`
}

// Resource: Parsers

type Parser struct {
	Name                 string            `json:"name,omitempty"`
	Creator              map[string]string `json:"creator,omitempty"`
	CreateTime           string            `json:"createTime,omitempty"`
	Changelogs           map[string]any    `json:"changelogs,omitempty"`
	ParserExtension      string            `json:"parserExtension,omitempty"`
	Type                 string            `json:"type,omitempty"`
	State                string            `json:"state,omitempty"`
	ValidationReport     string            `json:"validationReport,omitempty"`
	ValidatedOnEmptyLogs bool              `json:"validatedOnEmptyLogs,omitempty"`
	Cbn                  string            `json:"cbn,omitempty"`
	LowCode              map[string]any    `json:"lowCode,omitempty"`
	ReleaseStage         string            `json:"releaseStage,omitempty"`
	ValidationStage      string            `json:"validationStage,omitempty"`
}

// Method: activate

type ActivateParserRequest struct {
	Path string
}

func NewActivateParserRequest(project, location, instance, logtype, parser string) *ActivateParserRequest {
	return &ActivateParserRequest{
		Path: fmt.Sprintf(parserspath, project, location, instance, logtype) + "/" + parser + ":activate",
	}
}

// Method: create

type CreateParserRequest struct {
	Path                 string
	Cbn                  string
	ValidatedOnEmptyLogs bool
}

func NewCreateParserRequest(project, location, instance, logtype string, cbn []byte, validatedonemptylogs bool) *CreateParserRequest {
	return &CreateParserRequest{
		Path:                 fmt.Sprintf(parserspath, project, location, instance, logtype),
		Cbn:                  base64.StdEncoding.EncodeToString(cbn),
		ValidatedOnEmptyLogs: validatedonemptylogs,
	}
}

// Method: deactivate

type DeactivateParserRequest struct {
	Path string
}

func NewDeactivateParserRequest(project, location, instance, logtype, parser string) *ActivateParserRequest {
	return &ActivateParserRequest{
		Path: fmt.Sprintf(parserspath, project, location, instance, logtype) + "/" + parser + ":deactivate",
	}
}

// Method: delete

type DeleteParserRequest struct {
	Path  string
	Force bool
}

func NewDeleteParserRequest(project, location, instance, logtype, parser string, force bool) *DeleteParserRequest {
	return &DeleteParserRequest{
		Path:  fmt.Sprintf(parserspath, project, location, instance, logtype) + "/" + parser,
		Force: force,
	}
}

// Method: get

type GetParserRequest struct {
	Path string
}

func NewGetParserRequest(project, location, instance, logtype, parser string) *GetParserRequest {
	return &GetParserRequest{
		Path: fmt.Sprintf(parserspath, project, location, instance, logtype) + "/" + parser,
	}
}

// Method: list

type ListParsersRequest struct {
	Path      string
	PageSize  int
	PageToken string
	Filter    string
}

func NewListParsersRequest(project, location, instance, logtype string) *ListParsersRequest {
	return &ListParsersRequest{
		Path: fmt.Sprintf(parserspath, project, location, instance, logtype),
	}
}

type ListParsersResult struct {
	Parsers []Parser `json:"parsers"`
}

// Resource: parserExtension

type ParserExtension struct {
	CbnSnippet string `json:"cbnSnippet,omitempty"`
}

// Method: activate

type ActivateParserExtensionRequest struct {
	Path string
}

func NewActivateParserExtensionRequest(project, location, instance, logtype, extension string) *ActivateParserRequest {
	return &ActivateParserRequest{
		Path: fmt.Sprintf(parsersextensionspath, project, location, instance, logtype) + "/" + extension + ":activate",
	}
}

// Method: create

type CreateParserExtensionRequest struct {
	Path       string
	CbnSnippet string
}

func NewCreateParserExtensionRequest(project, location, instance, logtype, cbnsnippet string) *CreateParserExtensionRequest {
	return &CreateParserExtensionRequest{
		Path:       fmt.Sprintf(parsersextensionspath, project, location, instance, logtype),
		CbnSnippet: cbnsnippet,
	}
}

// Method: Delete

type DeleteParserExtensionRequest struct {
	Path string
}

func NewDeleteParserExtensionRequest(project, location, instance, logtype, parser string, force bool) *DeleteParserExtensionRequest {
	return &DeleteParserExtensionRequest{
		Path: fmt.Sprintf(parserspath, project, location, instance, logtype) + "/" + parser,
	}
}
