/* functions for manipulating the database */

package utils

import (
	"database/sql"
	"os"
)


func GetStatusQuo(db *sql.DB) (map[string]DatabasePrice, error) {

	// get map of station-uuid and DatabasePrice pairs

	// init empty map with space for 15 items before reallocation
	prices := make(map[string]DatabasePrice, 15)

	// load SQL query from file and convert to string
	// NOTE: relative to `main.go`!
	queryBytes, err := os.ReadFile("./queries/crawler-latest.sql") 
	if err != nil {
		return nil, err
	}
	query := string(queryBytes)

	// execute query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		
		var k string // use `station uuid` as key
		var p DatabasePrice

		if err := rows.Scan(&p.Station_id, &k, &p.Is_Open, &p.E5, &p.E10, &p.Diesel); err != nil {
			return nil, err
		}

		// add new key-value pair to `prices` map
		prices[k] = p
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// return map with prices	
	return prices, nil

}


// func UpdatePriceTimestamp(db *sql.DB, ts time.Time, fuel string, station_id int) error {
func UpdatePriceTimestamp(db *sql.DB, ts int64, fuel string, station_id int) error {

	// update a `price` row's timestamp; this is used when the latest price of a 
	// given `fuel & station_id` combination stays the same between iterations
	
	// convert time stamp to string
	// strTs := TsToString(ts)
	
	// load SQL query from file and convert to string
	// NOTE: relative to `main.go`!
	queryBytes, err := os.ReadFile("./queries/crawler-update-timestamp.sql") 
	if err != nil {
		return err
	}
	query := string(queryBytes)
	
	// execute query
	// rows, err := db.Query(query, strTs, fuel, station_id)
	rows, err := db.Query(query, ts, fuel, station_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	return nil
	
}


// func InsertPrice(db *sql.DB, ts time.Time, price float64, fuel string, station_id int) error {
func InsertPrice(db *sql.DB, ts int64, price float64, fuel string, station_id int) error {

	// insert a new `price` row for a given `fuel` and `station_id` combination

	// convert time stamp to string
	// strTs := TsToString(ts)
	
	// load SQL query from file and convert to string
	// NOTE: relative to `main.go`!
	queryBytes, err := os.ReadFile("./queries/crawler-insert-price.sql")
	if err != nil {
		return err
	}
	query := string(queryBytes)

	// execute query
	// rows, err := db.Query(query, strTs, price, fuel, station_id)
	rows, err := db.Query(query, ts, price, fuel, station_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil

}


func UpdateStationStatus(db *sql.DB, id int, is_open bool) error {

	// update a station's `is_open` status via its `id`

	/* NOTE

	SQLite does not have a separate Boolean storage class. Instead, 
	Boolean values are stored as integers 0 (false) and 1 (true).
		[https://www.sqlite.org/draft/datatype3.html]

	In Go, is no built-in way of casting from bool to int!
		[https://stackoverflow.com/questions/8393933/]	
	*/ 

	// manually convert from bool to int
	status := 0
	if is_open {
		status = 1
	}
	
	// load SQL query from file and convert to string
	// NOTE: relative to `main.go`!
	queryBytes, err := os.ReadFile("./queries/crawler-update-status.sql")
	if err != nil {
		return err
	}
	query := string(queryBytes)

	// execute query: insert values for `is_open` and `id` (note the order!)
	rows, err := db.Query(query, status, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil

}
