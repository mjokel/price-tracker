package tests

import (
	U "crawler/utils"
	"testing"
)


func TestListPartitioning(t *testing.T) {

	/* case: partition list into sub-lists of given length */

	i := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "jackfruit", "kiwi", "lemon", "mango", "nectarine"}


	// chunk size = 6
	p := U.PartitionList(i, 6)

    if len(p) != 3 {
       t.Errorf("Bad partitioning, got size %d, expected: %d.", len(p), 3)
    }

	// chunk size = 10
	p = U.PartitionList(i, 10)

    if len(p) != 2 {
       t.Errorf("Bad partitioning, got size %d, expected: %d.", len(p), 2)
    }
	

}


func TestApiFalse(t *testing.T) {

	/* case: station is `open`, yet features no price data */

	a := U.ApiPriceHelper{
		Status: "open",
		E5: 	[]byte(`false`), // resembles "raw json response"
		E10: 	[]byte(`false`),
		Diesel: []byte(`false`),
	}

	r := U.ApiResponse{
		Ok:		true,
		License: "",
		Data: 	"",
		Prices: map[string]U.ApiPriceHelper{
			"a": a,
		},
	}

	// init map for bundling data
	p := make(map[string]U.ApiPrice)

	// translate `ApiPriceHelper` to `ApiPrice`
	U.ParseResponse(&p, r)

	// check if `false` prices were converted to `0`
	if p["a"].E5 != 0 {
		t.Errorf("Incorrect handling of bad price, got %f, expected: %f.", p["a"].E5, 0.0)
	}
	if p["a"].E10 != 0 {
		t.Errorf("Incorrect handling of bad price, got %f, expected: %f.", p["a"].E10, 0.0)
	}
	if p["a"].Diesel != 0 {
		t.Errorf("Incorrect handling of bad price, got %f, expected: %f.", p["a"].Diesel, 0.0)
	}

}


func TestApiClosed(t *testing.T) {

	/* case: station is `closed` */

	a := U.ApiPriceHelper{
		Status: "closed",
	}

	r := U.ApiResponse{
		Ok:		true,
		License: "",
		Data: 	"",
		Prices: map[string]U.ApiPriceHelper{
			"a": a,
		},
	}

	// init map for bundling data
	p := make(map[string]U.ApiPrice)

	// translate `ApiPriceHelper` to `ApiPrice`
	U.ParseResponse(&p, r)

	// check if `false` prices were converted to `0`
	if p["a"].E5 != 0 {
		t.Errorf("Incorrect handling of bad price, got %f, expected: %f.", p["a"].E5, 0.0)
	}
	if p["a"].E10 != 0 {
		t.Errorf("Incorrect handling of bad price, got %f, expected: %f.", p["a"].E10, 0.0)
	}
	if p["a"].Diesel != 0 {
		t.Errorf("Incorrect handling of bad price, got %f, expected: %f.", p["a"].Diesel, 0.0)
	}

}
