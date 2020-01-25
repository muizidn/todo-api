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

type SqlPair map[string]interface{}

// returns (column1,column2,...) (?,?,...) []interface{}
func sqlStatementBuildWithValues(pair SqlPair) (string, string, []interface{}) {
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

// returns column1=value, column2=value
func sqlUpdateStatementBuildWithValues(pair SqlPair) string {
	updates := make([]string, 0)
	for k, v := range pair {
		update := fmt.Sprintf(`%s='%v'`, k, v)
		updates = append(updates, update)
	}
	return strings.Join(updates, ",")
}
