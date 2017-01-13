package gopy


type Cmd struct {
	Args    []string `json:"args"`
	Extra   map[string]interface{} `json:"extra"`
	Machine string
}

