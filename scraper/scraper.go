package main

/*	COLLECTION OF USEFUL RESOURCES

	https://go.dev/doc/tutorial/database-access
	https://go.dev/doc/database/prepared-statements
	https://gosamples.dev/sqlite-intro/
	https://gobyexample.com/string-formatting -> Printing

*/

import (
	"log"
	"math/rand"
	"os"
	"time"

	U "scraper/utils"

	"github.com/robfig/cron"
	_ "github.com/robfig/cron/v3"
	_ "modernc.org/sqlite"
)

func wrapper() {

	// get API key from environment variable
    ak := os.Getenv("API_KEY")
    if ak == "" {
        log.Fatal("API_KEY is not set or is empty")
    } else {
		log.Println("API_KEY:", ak)
	}

	// get path to database file from environment variable
	dp := os.Getenv("DB_PATH")
    if dp == "" {
        log.Fatal("DB_PATH is not set or is empty")
    } else {
		log.Println("DB_PATH:", dp)
	}


	// delay program execution randomly by up to 30 seconds
	secs := time.Duration(rand.Intn(30)+1)
	log.Printf("delay = %d seconds\n", secs)
	time.Sleep(secs * time.Second)


	// init new database connection
	db, err := U.OpenDB(dp)
	if err != nil {
		log.Fatal(err)
		return 
	}
	defer db.Close() // close when done
	
	
	// get all stations with latest prices from database
	pricesDB, err := U.GetStatusQuo(db)
	if err != nil {
		log.Fatal(err)
	}

	// for key, item := range pricesDB {
	// 	fmt.Printf("%s: %+v\n", key, item)
	// }


	// get all keys of `pricesDB` and partition into lists with 10 items each
	var keys []string
	for key := range pricesDB {
		keys = append(keys, key)
	}
	partitions := U.PartitionList(keys, 10)


	// init map for storing price data from API
	// pricesAPI := make(map[string]*U.ApiPrice)
	pricesAPI := make(map[string]U.ApiPrice)
	
	// get current time as UNIX epoch as integer
	ts := time.Now().In(time.UTC).Unix()


	// iterate partitions and call API
	for _, partition := range partitions {

		// for each partition, query API for latest prices
		// note: pass pointer to map!
		err := U.CallAPI(ak, &pricesAPI, partition)
		if err != nil {
			log.Fatal(err)
		}

	}
	
	// fmt.Println("Prices Map:")
	// for key, item := range pricesAPI {
	// 	fmt.Printf("%s: %+v\n", key, item)
	// }


	// finally, compare prices and update database accordingly
	// pass pricesDB and pricesAPI as pointers, i.e. their memory addresses
	U.Compare(db, ts, &pricesDB, &pricesAPI)

}


func main() {

	// set up a logger, based on https://stackoverflow.com/a/19966217
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)
	log.Println("Initialized logger!")
	

	// set up "crontab" and call wrapper routine every 5 minutes
	c := cron.New()
	c.AddFunc("@every 5m", wrapper)
	c.Start()
	defer c.Stop()

	select {}

}
