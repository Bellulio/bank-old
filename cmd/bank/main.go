package main

import  ("bank/pkg/bank/types"
		 "fmt")

func main (){
var card = types.Card {
	ID:          22345,
	PAN:         "5058 xxxx xxxx 9999",
	Balance:     999_99,
	Currency:    "TJS",
	Color:       "white",
	Name:        "Infinity",
	Active:      true,    
	}
var card= types.Card {	
		
	ID:          22345,
	PAN:         "5058 xxxx xxxx 8888",
	Balance:     888_88,
	Currency:    "TJS",
	Color:       "white",
	Name:        "Infinity",
	Active:      true,    	
    }
		
}

fmt.Printf("%+v", card)
handle(card)

	
}