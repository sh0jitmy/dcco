package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Transaction struct {
	ID       string   `yaml:"id"`
	Path     string   `yaml:"path"`
	PropsID  []string `yaml:"propsid"`
	Services []string `yaml:"services"`
}

type Config struct {
	Transactions []Transaction `yaml:"transactions"`
}

func main() {
	// YAMLファイルの読み込み
	data, err := ioutil.ReadFile("MultiTransaction.yml")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// YAMLのパース
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// パース結果の表示
	for _, transaction := range config.Transactions {
		fmt.Printf("ID: %s\n", transaction.ID)
		fmt.Printf("Path: %s\n", transaction.Path)
		fmt.Println("PropsID:")
		for _, propID := range transaction.PropsID {
			fmt.Printf("  - %s\n", propID)
		}
		fmt.Println("Services:")
		for _, service := range transaction.Services {
			fmt.Printf("  - %s\n", service)
		}
		fmt.Println()
	}
}

