/* establish SQLite database connection */

package utils

import (
	"database/sql"
	"errors"
	"log"
)


func OpenDB(p string) (*sql.DB, error) {

	/* 
	
	init and return new database connection; ensure to call `defer db.Close()`!

	Argument
		p	path to database file as string

	*/

	// check if file exists
	if !FileExists(p) {
		log.Println("File", p, "does NOT exist!")
		return nil, errors.New("File does not exist!")
	} 
	
	db, err := sql.Open("sqlite", p) // NOTE: relative to `main.go`!
	if err != nil {
		return nil, err
	}

	log.Println("openDB(): connected to database at", p)

	return db, nil

}
