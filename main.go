package main

import (
	"deploy-automation/config"
	"deploy-automation/deploy"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ServerConfig struct {
	Name            string `json:"name"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	KeyPath         string `json:"key_path"`
	DeployPath      string `json:"deploy_path"`
	RestartCommands []string `json:"restart_commands"`
}

type Config struct {
	Servers []ServerConfig `json:"servers"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: deploy-automation <file_to_deploy>")
	}

	filePath := os.Args[1]

	configFile, err := os.ReadFile("config/servers.json")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var cfg Config
	if err := json.Unmarshal(configFile, &cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	for _, server := range cfg.Servers {
		fmt.Printf("\n=== Deploying to %s (%s) ===\n", server.Name, server.Host)

		client, err := deploy.NewSSHClient(server.Host, server.Port, server.User, server.KeyPath)
		if err != nil {
			log.Printf("Failed to connect to %s: %v", server.Name, err)
			continue
		}

		deployer := deploy.NewDeployer(client)

		if err := deployer.UploadFile(filePath, server.DeployPath); err != nil {
			log.Printf("Failed to upload to %s: %v", server.Name, err)
			client.Close()
			continue
		}

		fmt.Printf("File uploaded to %s\n", server.DeployPath)

		for _, cmd := range server.RestartCommands {
			fmt.Printf("Executing: %s\n", cmd)
			output, err := deployer.RunCommand(cmd)
			if err != nil {
				log.Printf("Command failed: %v", err)
			} else {
				fmt.Printf("Output: %s\n", output)
			}
		}

		client.Close()
		fmt.Printf("Deployment to %s completed\n", server.Name)
	}

	fmt.Println("\n=== All deployments completed ===")
}