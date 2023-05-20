package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"gopkg.in/yaml.v2"
)

type IndexMap map[string]int

func main() {
	// YAMLファイルの読み込み
	yamlData, err := ioutil.ReadFile("SortProp.yml")
	if err != nil {
		log.Fatal(err)
	}

	// YAMLのパース
	var yamlList []string
	err = yaml.Unmarshal(yamlData, &yamlList)
	if err != nil {
		log.Fatal(err)
	}

	// インデックスマップを作成
	indexMap := createIndexMap(yamlList)

	// ソート対象の文字列スライス
	strSlice := []string{"orange", "grape", "apple", "banana"}

	// リスト順に文字列をソート
	sort.SliceStable(strSlice, func(i, j int) bool {
		return indexMap[strSlice[i]] < indexMap[strSlice[j]]
	})

	// ソートされた結果を出力
	fmt.Println(strSlice)
}

// インデックスマップを作成
func createIndexMap(yamlList []string) IndexMap {
	indexMap := make(IndexMap)

	for i, item := range yamlList {
		indexMap[item] = i
	}

	return indexMap
}


type Item struct {
    ID    string
    Value string
}

func FindDifference(list1, list2 []Item) []Item {
    difference := []Item{}
    seen := map[string]string{}

    for _, item := range list2 {
        seen[item.ID] = item.Value
    }

    for _, item := range list1 {
        if value, ok := seen[item.ID]; !ok || value != item.Value {
            difference = append(difference, item)
        }
    }

    return difference
}

