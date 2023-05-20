package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Merge struct {
	Path     string   `yaml:"path"`
	Resource string   `yaml:"resource"`
	Props    []string `yaml:"props"`
}

type Config struct {
	Merges []Merge `yaml:"merges"`
}

func main() {
	// YAMLファイルの読み込み
	data, err := ioutil.ReadFile("MergeInfo.yml")
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
	for _, merge := range config.Merges {
		fmt.Printf("Path: %s\n", merge.Path)
		fmt.Printf("Resource: %s\n", merge.Resource)
		fmt.Println("Props:")
		for _, prop := range merge.Props {
			fmt.Printf("  - %s\n", prop)
		}
		fmt.Println()
	}
}

