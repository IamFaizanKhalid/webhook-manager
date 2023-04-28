package main

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/hook"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Name      string `yaml:"name"`
	Command   string `yaml:"command"`
	Directory string `yaml:"directory"`
	Branch    string `yaml:"branch"`
}

const webhookSecret = "my_secret"

func main() {
	// Read configs from YAML file
	configsFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("failed to read configs file: %v", err)
	}

	var configs []Config
	err = yaml.Unmarshal(configsFile, &configs)
	if err != nil {
		log.Fatalf("failed to unmarshal configs: %v", err)
	}

	//configs := []Config{
	//	{Name: "test-api", Command: "/usr/local/test-api/scripts/deploy.sh", Directory: "/usr/local/test-api", Branch: "master"},
	//	{Name: "other-api", Command: "/usr/local/other-api/scripts/deploy.sh", Directory: "/usr/local/other-api", Branch: "develop"},
	//}

	hooks := make([]*hook.Hook, len(configs))
	for i, config := range configs {
		hooks[i] = WebhookFromConfig(&config)
	}

	data, err := yaml.Marshal(hooks)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("hooks.yml", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("hooks.yml file created successfully")
}
