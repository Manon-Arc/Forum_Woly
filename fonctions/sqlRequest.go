package pagesfonctions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Insert(table string, entries map[string]interface{}, toReturn string) (int64, error) {

	// ?   To use this function :
	// ?   'table' is the name of the table where we want to put data
	// ?   'entries' is the map which contains the name of the column for key and the value to put in as value (column:value)
	// ?   To called this function :
	// ?   Make a map[string]interface{}   =>    <map_name> := make(map[string]interface{})
	// ?   Then put some value in the map  =>    <map_name>["<column_name>"] = <value>
	// ?   Finaly call the function        =>    Insert(<table_name>, <map_name>)

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}

	ctx := context.Background()

	// Check if database is alive.
	err := database.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}

	tsql := `INSERT INTO ` + table + ` (`
	tsqlEnd := `VALUES (`

	i := 0

	for key, values := range entries {
		tsql += `'` + key + `'`
		tsqlEnd += `'` + fmt.Sprint(values) + `'`
		i++
		if i != len(entries) {
			tsql += `,`
			tsqlEnd += `,`
		} else {
			tsql += `) `
			tsqlEnd += `)`
		}
	}
	tsql += tsqlEnd

	if toReturn != "" {
		tsql += `RETURNING ` + toReturn
	}

	stmt, err := database.Prepare(tsql)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx)

	var newID int64
	err = row.Scan(&newID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return newID, nil

}

func Update(table string, changeValue, where map[string]interface{}) {

	// ?   To use this function :
	// ?   'table' is the name of the table where we want to put data
	// ?   'changeValue' is the map which contains the name of the column to update for key and the value to change in as value (column:value)
	// ?   To called this function :
	// ?   Make tow map[string]interface{}   =>    <map_name> := make(map[string]interface{})
	// ?   Then put some value in the map  =>    The first one : <map_name>["<column_to_update>"] = <new_value> | The second one : <map_name>["<column_requeried>"] = <value_to_check>
	// ?   Finaly call the function        =>    Update(<table_name>, <map_name>, <map_name>)

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}

	ctx := context.Background()

	// Check if database is alive.
	err := database.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}

	tsql := `UPDATE ` + table
	tsqlSet := ` SET `
	tsqlEnd := ` WHERE `

	i := 0

	for key, value := range changeValue {
		tsqlSet += `'` + key + `'` + ` = "` + fmt.Sprint(value) + `"`
		i++
		if i != len(changeValue) {
			tsqlSet += `,`
		} else {
			tsqlSet += ` `
		}
	}

	if len(where) != 0 {
		i = 0
		for key, value := range where {
			tsqlEnd += key + ` = '` + fmt.Sprint(value) + `'`
			i++
			if i != len(where) {
				tsqlEnd += ` AND `
			}
		}
	} else {
		tsqlEnd = ``
	}

	tsql += tsqlSet + tsqlEnd

	stmt, err := database.Prepare(tsql)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	rows := stmt.QueryRowContext(ctx)

	var newID int64
	err = rows.Scan(&newID)
	if err != nil && !errors.Is(err, sql.ErrNoRows){
		fmt.Println(err)
	}

}
