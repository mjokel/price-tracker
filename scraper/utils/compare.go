/* controller that handles price comparisons and database manipulation */

package utils

import (
	"database/sql"
	"log"
)


func Compare(db *sql.DB, ts int64, oldP *map[string]DatabasePrice, newP *map[string]ApiPrice) {

	/*

	Handles comparison of old and new prices and initiates respective database manipulations
	
	Arguments
		db		pointer to an open database connection
		ts		current time stamp in Unix epoch as integer
		oldP	pointer to map containing old prices, from database
		newP	pointer to map containing new prices, from API
	
	Returns
		error	error object
	
	*/

	// iterate old (DB) prices map
	for oKey, oItems := range *oldP {

		// if there is a matching key in the new (API) prices map
		if nItems, ok := (*newP)[oKey]; ok {

			log.Println("Compare(): key", oKey)
			log.Printf("Compare(): %+v\n", oItems)
			log.Printf("Compare(): %+v\n", nItems)

			
			// check if station is closed or open and update database accordingly
			if nItems.Status == "closed" {

				// update `is_open` flag, if necessary
				if oItems.Is_Open {
					log.Println("Compare(): updating status to CLOSED")
					UpdateStationStatus(db, oItems.Station_id, false)
				}

			} else if nItems.Status == "open" {

				// update `is_open` flag, if necessary
				if !oItems.Is_Open {
					log.Println("Compare(): updating status to OPEN")
					UpdateStationStatus(db, oItems.Station_id, true)
				}


				/* compare prices ------------------------------------------- */

				if oItems.E5 != nItems.E5 {
					log.Println("ΔE5 : from", oItems.E5, "to", nItems.E5)
					InsertPrice(db, ts, nItems.E5, "E5", oItems.Station_id)
				} else {
					UpdatePriceTimestamp(db, ts, "E5", oItems.Station_id)
				}

				if oItems.E10 != nItems.E10 {
					log.Println("ΔE10: from", oItems.E10, "to", nItems.E10)
					InsertPrice(db, ts, nItems.E10, "E10", oItems.Station_id)
				} else {
					UpdatePriceTimestamp(db, ts, "E10", oItems.Station_id)
				}

				if oItems.Diesel != nItems.Diesel {
					log.Println("ΔDsl: from", oItems.Diesel, "to", nItems.Diesel)
					InsertPrice(db, ts, nItems.Diesel, "Diesel", oItems.Station_id)
				} else {
					UpdatePriceTimestamp(db, ts, "Diesel", oItems.Station_id)
				}

			} else {
				log.Println("Compare(): unhandled API `status`", nItems.Status)
			}

		} else {
			log.Println("Compare(): key", oKey, "from DB is missing in API map")
		}

	} // end of loop over old (DB) map

}