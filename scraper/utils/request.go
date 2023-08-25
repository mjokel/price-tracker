/* functions for API interaction */

package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)


func ParseResponse(m *map[string]ApiPrice, r ApiResponse) {

	/*

	helper function for translating the API response object into Go structs;
	modularized to facilitate testing;
	
	Arguments
		m	pointer to map, that the parsed prices need to be inserted
		r	ApiResponse instance
	
	*/


	// iterate the structure holding the prices
	for key, val := range r.Prices {

		/* note: `val` is of type `ApiPriceHelper`
			-> 	need to check if `E5`, `E10` and `Diesel` fields are actually
				`float64` or `bool`
			-> we therefore try to unmarshal the fields to `float64`
			-> if this fails, we simply set the price to 0
		*/

		// init new variable of type `ApiPrice`, where prices are float64
		var price ApiPrice

		// first, transfer `status`; then consider the price fields
		price.Status = val.Status

		// in case of error, ignore it and set price to 0
		if err := json.Unmarshal(val.E5, &price.E5); err != nil {
			price.E5 = 0
		}
		if err := json.Unmarshal(val.E10, &price.E10); err != nil {
			price.E10 = 0
		}
		if err := json.Unmarshal(val.Diesel, &price.Diesel); err != nil {
			price.Diesel = 0
		}

		// finally, add `ApiPrice` instance to map
		(*m)[key] = price

	}

}



func CallAPI(k string, m *map[string]ApiPrice, ids []string) error {

	/*

	Queries Tankerkönig API for latest prices of given stations by their ids
	
	Arguments
		k 		API key as string
		m		map of key-value pairs for inserting the prices per station;
				note that the `value` is actually a pointer to a `ApiPrice` instance!
		ids		list of relevant station uuids

	Returns
		error
		implicitly returns map `m` with prices per station
	
	*/ 

	// programmatically create URL for API call
	u, err := url.Parse("https://creativecommons.tankerkoenig.de/json/prices.php")
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Add("ids", strings.Join(ids, ",")) // concatenate list items to single string
	params.Add("apikey", k)

	// set query parameters
	u.RawQuery = params.Encode()

	// get final URL as a string
	finalURL := u.String() 

	log.Println("CallAPI(): calling", finalURL)
	

	// init new client and configure request headers
	client := http.Client{}

	request, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", "Go Fuel Price Monitor")
	request.Header.Set("Content-Type", "application/json")

	// make request
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// check the response status code
	if response.StatusCode != http.StatusOK {
		return err
	}

	// process response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// fmt.Println(string(body))

	// init struct to hold the parsed JSON data
	var r ApiResponse

	// unmarshal JSON response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	// translate `ApiPriceHelper` to `ApiPrice`
	ParseResponse(m, r)

	return nil

}
