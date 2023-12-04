package meta_data

import "fmt"

type Table struct {
	SchemaName string `json:"schemaname"`
	Name       string `json:"tablename"`
	Owner      string `json:"tableowner"`
}

func PrintMetaTables(tables []Table) {
	for _, table := range tables {
		fmt.Printf("SchemaName='%s', Name='%s', Owner='%s'\n", table.SchemaName, table.Name, table.Owner)
	}
}
