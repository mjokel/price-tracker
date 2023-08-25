/* custom structs used within the scraper routine */

package utils

import "encoding/json"

/*
	Represents a row of the `crawler-latest.sql` query. Note that the struct
	does not feature a `station_uuid` field, as this value is used as access
	key in the map.
*/ 

type DatabasePrice struct {
	Station_id 		int
	Is_Open			bool
	E5  			float64
	E10	  			float64
	Diesel	  		float64
}


/*
	The Tankerkönig API response is represented by the `ApiResponse` struct, which may include multiple `ApiPrice` structs. Both feature JSON mapping annotations.
*/ 

/* 	temporary struct that can handle prices being numeric AND boolean; this is
	necessary, as the API might not be able to return a certain price 		  */

type ApiPriceHelper struct {
	Status  	string 	`json:"status"`
	E5  		json.RawMessage	`json:"e5"`
	E10	  		json.RawMessage `json:"e10"`
	Diesel	  	json.RawMessage `json:"diesel"`
}

type ApiPrice struct {
	Status  	string 	`json:"status"`
	E5  		float64 `json:"e5"`
	E10	  		float64 `json:"e10"`
	Diesel	  	float64 `json:"diesel"`
}

type ApiResponse struct {
	Ok  		bool   `json:"ok"`
	License  	string `json:"license"`
	Data	  	string `json:"data"`

	// maps are used to represent key-value pairs!
	// note that it is of type `helper`!
	Prices		map[string]ApiPriceHelper `json:"prices"`
}