package card

import (
	"fmt"
	"go/types"
)

func ExamplePaymentSource_possitive(){
	cards := types.Card{balance: 999_99, Active: true}
	Payment(&cards, 500_00)
	fmt.Println(Number)

	//Output: 
	//50000
}