package dcco 


// ItemのSort
func SortItem (im IndexMap, items *[]Item) {
        // リスト順に文字列をソート
        sort.SliceStable(items, func(i, j int) bool {
                return im[items.ID[i] < im[items.ID[j]]
        })
        // ソートされた結果を出力
        fmt.Println(items)
	return items
}


// インデックスマップを作成
func CreateIndexMap(path string) (IndexMap,error) {
        // YAMLのパース
        var yamlList []string
        indexMap := make(IndexMap)
	
	err := CommonYamlUnmarshal(path,&yamlList)
        if err != nil {
        	return indexMap,err
        }

        for i, item := range yamlList {
                indexMap[item] = i
        }
        return indexMap,nil
}

// Itemの差分抽出
func FindDiffItem(list1 []Item, list2 []Item) []Item {
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

func CreateTransactionConfig (path string) (TranScGenConfig,error) {
        var config TranScGenConfig
	err := CommonYamlUnmarshal(path,&config)
        // パース結果の表示
        for _, transaction := range config.Transactions {
                fmt.Printf("ID: %s\n", transaction.ID)
                fmt.Printf("Path: %s\n", transaction.Path)
                fmt.Println("ItemID:")
                for _, propID := range transaction.ItemIDs {
                        fmt.Printf("  - %s\n", propID)
                }
                fmt.Println("Services:")
                for _, service := range transaction.Services {
                        fmt.Printf("  - %s\n", service)
                }
                fmt.Println()
        }
	return config,nil
}

func CommonYamlUnmarshal(path string,schema interface{}) error {
        data, err := ioutil.ReadFile(path)
        if err != nil {
                log.Fatalf("Failed to read file: %v", err)
        	return config err
        }

        // YAMLのパース
        err = yaml.Unmarshal(data, &schema)
	return schema,err
}


func CreateMergeConfig (path string) (MergeConfig,error) {
        var config MergeConfig
	err := CommonYamlUnmarshal(path,&config)
        // YAMLファイルの読み込み
        if err != nil {
                log.Fatalf("Failed to parse YAML: %v", err)
        	return config err
        }

        // パース結果の表示
        for _, merge := range config.Merges {
                fmt.Printf("Path: %s\n", merge.Path)
                fmt.Printf("Resource: %s\n", merge.Resource)
                fmt.Println("Item:")
                for _, prop := range merge.Item {
                        fmt.Printf("  - %s\n", prop)
                }
                fmt.Println()
        }
	return config,nil
}

func CreateResourceMap(path string)(ResouceMap,error) {
        var config ResourceConfig
    	rmap := map[string]ResourceEntry
	 
	err := CommonYamlUnmarshal(path,&config)
        // YAMLファイルの読み込み
        if err != nil {
                log.Fatalf("Failed to parse YAML: %v", err)
        	return rmap err
        }
	return rmap,nil
        for _, rf := range config.FindList {
		rmap[Path+Cmd] = rf.Entry
	}	
        return rmap nil 
	
}

func ExtGetResourceIDs(cmd string,path string,rm *ResourceMap)([]string) {
	empty := []string
	if val,ok := rm[path+cmd]; ok {
		return val.SourceIDS	
	} 
	return  empty
}

func ExtGetResourceType(cmd string,path string,rm *ResourceMap)(string) {
	if val,ok := rm[path+cmd]; ok {
		return val.Type	
	} 
	return "" 
}

