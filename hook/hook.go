package hook

import (
	"os"
)

// Constants used to specify the parameter source
const (
	SourceHeader         string = "header"
	SourceQuery          string = "url"
	SourceQueryAlias     string = "query"
	SourcePayload        string = "payload"
	SourceRawRequestBody string = "raw-request-body"
	SourceRequest        string = "request"
	SourceString         string = "string"
	SourceEntirePayload  string = "entire-payload"
	SourceEntireQuery    string = "entire-query"
	SourceEntireHeaders  string = "entire-headers"
)

const (
	// EnvNamespace is the prefix used for passing arguments into the command
	// environment.
	EnvNamespace string = "HOOK_"
)

// Argument type specifies the parameter key name and the source it should
// be extracted from
type Argument struct {
	Source       string `yaml:"source,omitempty"`
	Name         string `yaml:"name,omitempty"`
	EnvName      string `yaml:"envname,omitempty"`
	Base64Decode bool   `yaml:"base64decode,omitempty"`
}

// Header is a structure containing header name, and it's value
type Header struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// ResponseHeaders is a slice of Header objects
type ResponseHeaders []Header

// Hook type is a structure containing details for a single hook
type Hook struct {
	ID                                  string          `yaml:"id,omitempty"`
	ExecuteCommand                      string          `yaml:"execute-command,omitempty"`
	CommandWorkingDirectory             string          `yaml:"command-working-directory,omitempty"`
	ResponseMessage                     string          `yaml:"response-message,omitempty"`
	ResponseHeaders                     ResponseHeaders `yaml:"response-headers,omitempty"`
	CaptureCommandOutput                bool            `yaml:"include-command-output-in-response,omitempty"`
	CaptureCommandOutputOnError         bool            `yaml:"include-command-output-in-response-on-error,omitempty"`
	PassEnvironmentToCommand            []Argument      `yaml:"pass-environment-to-command,omitempty"`
	PassArgumentsToCommand              []Argument      `yaml:"pass-arguments-to-command,omitempty"`
	PassFileToCommand                   []Argument      `yaml:"pass-file-to-command,omitempty"`
	JSONStringParameters                []Argument      `yaml:"parse-parameters-as-json,omitempty"`
	TriggerRule                         *Rules          `yaml:"trigger-rule,omitempty"`
	TriggerRuleMismatchHttpResponseCode int             `yaml:"trigger-rule-mismatch-http-response-code,omitempty"`
	TriggerSignatureSoftFailures        bool            `yaml:"trigger-signature-soft-failures,omitempty"`
	IncomingPayloadContentType          string          `yaml:"incoming-payload-content-type,omitempty"`
	SuccessHttpResponseCode             int             `yaml:"success-http-response-code,omitempty"`
	HTTPMethods                         []string        `yaml:"http-methods,omitempty"` // omitempty added later
}

// FileParameter describes a pass-file-to-command instance to be stored as file
type FileParameter struct {
	File    *os.File
	EnvName string
	Data    []byte
}

// Hooks is an array of Hook objects
type Hooks []Hook

// Rules is a structure that contains one of the valid rule types
type Rules struct {
	And   *AndRule   `yaml:"and,omitempty"`
	Or    *OrRule    `yaml:"or,omitempty"`
	Not   *NotRule   `yaml:"not,omitempty"`
	Match *MatchRule `yaml:"match,omitempty"`
}

// AndRule will evaluate to true if and only if all the ChildRules evaluate to true
type AndRule []Rules

// OrRule will evaluate to true if any of the ChildRules evaluate to true
type OrRule []Rules

// NotRule will evaluate to true if any and only if the ChildRule evaluates to false
type NotRule Rules

// MatchRule will evaluate to true based on the type
type MatchRule struct {
	Type      string   `yaml:"type,omitempty"`
	Regex     string   `yaml:"regex,omitempty"`
	Secret    string   `yaml:"secret,omitempty"`
	Value     string   `yaml:"value,omitempty"`
	Parameter Argument `yaml:"parameter,omitempty"`
	IPRange   string   `yaml:"ip-range,omitempty"`
}

// Constants for the MatchRule type
const (
	MatchValue      string = "value"
	MatchRegex      string = "regex"
	MatchHMACSHA1   string = "payload-hmac-sha1"
	MatchHMACSHA256 string = "payload-hmac-sha256"
	MatchHMACSHA512 string = "payload-hmac-sha512"
	MatchHashSHA1   string = "payload-hash-sha1"
	MatchHashSHA256 string = "payload-hash-sha256"
	MatchHashSHA512 string = "payload-hash-sha512"
	IPWhitelist     string = "ip-whitelist"
	ScalrSignature  string = "scalr-signature"
)
