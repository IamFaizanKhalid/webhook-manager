package main

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/hook"
)

func WebhookFromConfig(config *Config) *hook.Hook {
	return &hook.Hook{
		ID:                      config.Name,
		ExecuteCommand:          config.Command,
		CommandWorkingDirectory: config.Directory,
		TriggerRule: &hook.Rules{
			And: &hook.AndRule{
				{
					Match: &hook.MatchRule{
						Type:  hook.MatchHashSHA256,
						Value: webhookSecret,
						Parameter: hook.Argument{
							Source: hook.SourceHeader,
							Name:   "X-Hub-Signature-256",
						},
					},
				},
				{
					Match: &hook.MatchRule{
						Type:  hook.MatchValue,
						Value: fmt.Sprintf("refs/heads/%s", config.Branch),
						Parameter: hook.Argument{
							Source: hook.SourcePayload,
							Name:   "ref",
						},
					},
				},
			},
		},
	}
}
