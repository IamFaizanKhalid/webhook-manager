package model

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
	Source       string `yaml:"source,omitempty" json:"source"`
	Name         string `yaml:"name,omitempty" json:"name"`
	EnvName      string `yaml:"envname,omitempty" json:"env_name"`
	Base64Decode bool   `yaml:"base64decode,omitempty" json:"base64_decode"`
}

// Header is a structure containing header name, and it's value
type Header struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
}

// ResponseHeaders is a slice of Header objects
type ResponseHeaders []Header

// Hook type is a structure containing details for a single hook
type Hook struct {
	ID                                  string          `yaml:"id,omitempty" json:"id"`
	ExecuteCommand                      string          `yaml:"execute-command,omitempty" json:"execute_command"`
	CommandWorkingDirectory             string          `yaml:"command-working-directory,omitempty" json:"command_working_directory"`
	ResponseMessage                     string          `yaml:"response-message,omitempty" json:"response_message"`
	ResponseHeaders                     ResponseHeaders `yaml:"response-headers,omitempty" json:"response_headers"`
	CaptureCommandOutput                bool            `yaml:"include-command-output-in-response,omitempty" json:"include_command_output_in_response"`
	CaptureCommandOutputOnError         bool            `yaml:"include-command-output-in-response-on-error,omitempty" json:"include_command_output_in_response_on_error"`
	PassEnvironmentToCommand            []Argument      `yaml:"pass-environment-to-command,omitempty" json:"pass_environment_to_command"`
	PassArgumentsToCommand              []Argument      `yaml:"pass-arguments-to-command,omitempty" json:"pass_arguments_to_command"`
	PassFileToCommand                   []Argument      `yaml:"pass-file-to-command,omitempty" json:"pass_file_to_command"`
	JSONStringParameters                []Argument      `yaml:"parse-parameters-as-json,omitempty" json:"parse_parameters_as_json"`
	TriggerRule                         *Rules          `yaml:"trigger-rule,omitempty" json:"trigger_rule"`
	TriggerRuleMismatchHttpResponseCode int             `yaml:"trigger-rule-mismatch-http-response-code,omitempty" json:"trigger_rule_mismatch_http_response_code"`
	TriggerSignatureSoftFailures        bool            `yaml:"trigger-signature-soft-failures,omitempty" json:"trigger_signature_soft_failures"`
	IncomingPayloadContentType          string          `yaml:"incoming-payload-content-type,omitempty" json:"incoming_payload_content_type"`
	SuccessHttpResponseCode             int             `yaml:"success-http-response-code,omitempty" json:"success_http_response_code"`
	HTTPMethods                         []string        `yaml:"http-methods,omitempty" json:"http_methods"`
}

// FileParameter describes a pass-file-to-command instance to be stored as file
type FileParameter struct {
	File    *os.File
	EnvName string
	Data    []byte
}

// Rules is a structure that contains one of the valid rule types
type Rules struct {
	And   *AndRule   `yaml:"and,omitempty" json:"and"`
	Or    *OrRule    `yaml:"or,omitempty" json:"or"`
	Not   *NotRule   `yaml:"not,omitempty" json:"not"`
	Match *MatchRule `yaml:"match,omitempty" json:"match"`
}

// AndRule will evaluate to true if and only if all the ChildRules evaluate to true
type AndRule []Rules

// OrRule will evaluate to true if any of the ChildRules evaluate to true
type OrRule []Rules

// NotRule will evaluate to true if any and only if the ChildRule evaluates to false
type NotRule Rules

// MatchRule will evaluate to true based on the type
type MatchRule struct {
	Type      string   `yaml:"type,omitempty" json:"type"`
	Regex     string   `yaml:"regex,omitempty" json:"regex"`
	Secret    string   `yaml:"secret,omitempty" json:"secret"`
	Value     string   `yaml:"value,omitempty" json:"value"`
	Parameter Argument `yaml:"parameter,omitempty" json:"parameter"`
	IPRange   string   `yaml:"ip-range,omitempty" json:"ip_range"`
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
