package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func PaymentSource(cards []types.Card)[]types.PaymentSource{
	var payments []types.PaymentSource
	for _, card := range cards{
		if !card.Active || card.Balance <= 0{
			continue
		}
		return payments
	}
}
func ExamplePaymentSource_possitive(){
	card:= []types.Card{
    {
       Balance : 999_99,
	   PAN : "5058 xxxx xxxx 9999",
	   Active: true,
    },
	fmt.Println(PaymentSouce(card)),
	//Output:"5058 xxxx xxxx 9999" 
}
}
	func ExamplePaymentSource_positive(){
		
		card:= []types.Card{
		{
		   Balance : 888_88,
		   PAN : "5058 xxxx xxxx 8888",
		   Active: true,
		},
		fmt.Println(PaymentSouce(card)),
		//Output:"5058 xxxx xxxx 8888" 
	}
 
      

Number == card.PAN
fmt.Println(PaymentSource.ID)
