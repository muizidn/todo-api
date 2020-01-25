package app

import "strings"

import "fmt"

// returns (column1,column2,...) (?,?,...)
func sqlStatementBuild(columns ...string) (string, string) {
	placeholders := []string{}
	for range columns {
		placeholders = append(placeholders, "?")
	}
	cs := fmt.Sprintf("(%s)", strings.Join(columns, ","))
	ps := fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
	return cs, ps
}

// returns (column1,column2,...) (?,?,...) []interface{}
func sqlStatementBuildWithValues(pair map[string]interface{}) (string, string, []interface{}) {
	placeholders := []string{}
	columns := []string{}
	values := make([]interface{}, 0)
	for k, v := range pair {
		columns = append(columns, k)
		placeholders = append(placeholders, "?")
		values = append(values, v)
	}
	cs := fmt.Sprintf("(%s)", strings.Join(columns, ","))
	ps := fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
	return cs, ps, values
}
