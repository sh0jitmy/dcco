package dcco 

type Item struct {
	ID    string
	Value string
}

type IndexMap map[string]int


type ResourceEntry struct {
	Type string   `yaml:"type"`
	SourceIDS  []string `yaml:"sourceids"`
}

type ResourceFinds struct {
	Path  string `yaml:"path"`
	Cmd   string `yaml:"cmd"`
	ResourceEntry Entry  `yaml:"entry"`
}

type ResourceConfig struct {
	FindsList []ResourceFinds `yaml:"resources"`
}

type ResourceMap map[string] ResourceEntry

type Merge struct {
        Path     string   `yaml:"path"`
        Resource string   `yaml:"resource"`
        ItemIDs    []string `yaml:"props"`
}

type MergeConfig struct {
        Merges []Merge `yaml:"merges"`
}

type TransactionScenario struct {
        ID       string   `yaml:"id"`
        Path     string   `yaml:"path"`
        ItemIDs  []string `yaml:"propsid"`
        Services []string `yaml:"services"`
}

type TranScGenConfig struct {
         [] TransactionScenario `yaml:"transactions"`
}


