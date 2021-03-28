package card
import (
	"bank/pkg/bank/types"

)


func IssueCard(currency types.Currency, color string, name string) types.Card   {
	card := types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
	return card	
}

func PaymentSours(cards[]types.Card)[]types.PaymentSource {
	//TODO:

	if card.Active == false{
		return 
	}
	if amount < 0{
		return 
	}
	card.Balance -= amount
		
}

}
