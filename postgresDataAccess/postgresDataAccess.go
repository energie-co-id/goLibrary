package postgresDataAccess

import (
	"database/sql"
	"fmt"
)

func Select(db *sql.DB, organization string, table string, filter string, selectQuery string, values []any) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s.%s %s", selectQuery, organization, table, filter)

	if organization == "dev" {
		fmt.Printf("query %s %v\n", query, values)
	}

	return db.Query(query, values...)
}

func Insert(db *sql.DB, key string, params string, value, organization string, table string, unique string, returning string, values []any) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", organization, table, key, params)

	if unique != "" {
		query += fmt.Sprintf(" ON CONFLICT (%s) DO NOTHING", unique)
	}

	if returning != "" {
		query += fmt.Sprintf(" RETURNING %s", returning)
	}

	if organization == "dev" {
		fmt.Printf("insertQuery %s %v\n", query, values)
	}

	return db.Exec(query, values...)
}

func Update(db *sql.DB, key, filter, organization, table, returning string, values []any) (sql.Result, error) {
	query := fmt.Sprintf("UPDATE %s.%s SET %s %s", organization, table, key, filter)

	if returning != "" {
		query += fmt.Sprintf(" RETURNING %s", returning)
	}

	if organization == "dev" {
		fmt.Printf("updateQuery %s %v\n", query, values)
	}

	return db.Exec(query, values...)
}

func Delete(db *sql.DB, organization, table, filter, returning string, values []any) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s.%s %s", organization, table, filter)

	if returning != "" {
		query += fmt.Sprintf(" RETURNING %s", returning)
	}

	if organization == "dev" {
		fmt.Printf("delete query %s %v\n", query, values)
	}

	return db.Exec(query, values...)
}

func InsertMuchDatas(db *sql.DB, key, params, value, organization, table, unique, returning string, values []any) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES %s", organization, table, key, params)

	if unique != "" {
		query += fmt.Sprintf(" ON CONFLICT (%s) DO NOTHING", unique)
	}

	if returning != "" {
		query += fmt.Sprintf(" RETURNING %s", returning)
	}

	if organization == "dev" {
		fmt.Printf("insertQuery %s %v\n", query, values)
	}

	return db.Exec(query, values...)
}

// func ComposeUpdateParams(tableFields []string, keyValList map[string]interface{}, keyParams []string, updateKeyParams map[string]interface{}) (string, []interface{}, string) {
// 	var fields string
// 	var values []interface{}
// 	var index int = 1
// 	var keyParameter string

// 	for key, value := range keyValList {
// 		if contains(tableFields, key) {
// 			if contains(keyParams, key) {
// 				if keyValList[key] == nil {
// 					keyParameter += fmt.Sprintf("%s is NULL", key)
// 				} else {
// 					keyParameter += fmt.Sprintf("%s = $%d", key, index)
// 					values = append(values, value)
// 					index++
// 				}
// 			} else {
// 				if fields != "" {
// 					fields += ", "
// 				}
// 				fields += fmt.Sprintf("%s = $%d", key, index)
// 				values = append(values, value)
// 				index++
// 			}
// 		}
// 	}

// 	if updateKeyParams != nil {
// 		for key, value := range updateKeyParams {
// 			if contains(tableFields, key) {
// 				if fields != "" {
// 					fields += ", "
// 				}
// 				fields += fmt.Sprintf("%s = $%d", key, index)
// 				values = append(values, value)
// 				index++
// 			}
// 		}
// 	}

// 	return fields, values, keyParameter
// }

// func ComposeInsertParams(tableFields []string, keyValList map[string]interface{}, notUUID bool) (string, []interface{}, string) {
// 	var fields string
// 	var valuesTemplate string
// 	var values []interface{}
// 	var index int = 1

// 	for key, value := range keyValList {
// 		if contains(tableFields, key) {
// 			if fields != "" {
// 				fields += ", "
// 				valuesTemplate += ", "
// 			}

// 			fields += key

// 			if key == "id" && !notUUID {
// 				valuesTemplate += "uuid_generate_v4()"
// 			} else {
// 				valuesTemplate += fmt.Sprintf("$%d", index)
// 				values = append(values, value)
// 				index++
// 			}
// 		}
// 	}

// 	return fields, values, valuesTemplate
// }
