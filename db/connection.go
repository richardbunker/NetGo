package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DatabaseConnection *sql.DB = connection()

// OpenDB opens a connection to the database
func connection() *sql.DB {
	db, err := sql.Open("sqlite3", "./Database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Query(sql string, params ...any) []map[string]interface{} {
	// Query the database
	rows, err := DatabaseConnection.Query(sql, params...)
	if err != nil {
		log.Fatal("Error with Query: ", err)
	}

	defer rows.Close()

	// Get column names from the result set
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error getting columns: ", err)
	}

	// Prepare a slice of maps to hold the result
	data := []map[string]interface{}{}

	// Iterate over rows
	for rows.Next() {
		// Create a slice of interface{} to hold each column's data
		columnsData := make([]interface{}, len(columns))
		columnsDataPtrs := make([]interface{}, len(columns))

		for i := range columnsData {
			columnsDataPtrs[i] = &columnsData[i]
		}

		// Scan the row into the slice of interfaces
		err := rows.Scan(columnsDataPtrs...)
		if err != nil {
			log.Fatal("Error scanning row: ", err)
		}

		// Create a map to store the row with column names as keys
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnsData[i]
			// Handle `[]byte` (which is common for database strings) by converting to string
			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		// Add the row map to the result slice
		data = append(data, rowMap)
	}

	// Check for errors after iteration
	err = rows.Err()
	if err != nil {
		log.Fatal("Error after rows.Next: ", err)
	}

	return data
}

// // Query the database
// rows, err := db.Query("select * from users")
// if err != nil {
// 	log.Fatal("Error with Query: ", err)
// }
//
// // Iterate over the rows
// names := []string{}
// defer rows.Close()
// for rows.Next() {
// 	var id int
// 	var name string
// 	var email string
// 	err = rows.Scan(&id, &name, &email)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	names = append(names, name)
// 	fmt.Println(id, name, email)
// }
// err = rows.Err()
// if err != nil {
// 	log.Fatal(err)
// }
//
// fmt.Println(names)
