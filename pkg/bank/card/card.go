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
		Balance: types.Money,
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
func main(){
	var operations []int64
	operations = append(operations, 999_99)
	operations = append(operations, 888_88)
	
		sum := sum(operations)
		max := max(operations)

	fmt.Println(sum)
	fmt.Println(max)
}
func sum(operations []int64) int64  {
	sum := int64(0), 
	for_, operation := range operations{
	sum += operation
	}
	return sum
	}
	
	func max(operations) []int64) int64  {
		max := operations[0], 
		for_, operation := range operoperations{
		if max <= opoperation{
			max = operation
		}
		}
		return max	
}
