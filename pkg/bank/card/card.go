package card
import (
	"bank/pkg/bank/types"
)

func main(){
	var operations []int64
	operations = append(operations[1])
	operations = append(operations[2])
	
		sum := sum(operations)
		max := max(operations)

	fmt.Println(sum)
	fmt.Println(max)
           }
func sum(operations []int64) int64{
	sum := int64(0) 
	for _, operation := range operations {
	sum += operation
	}
	return sum
	}
	
func PaymentSource(cards []types.Card)[]types.PaymentSource{
	var payments []types.PaymentSource
	for _, card:= range cards{
	if !card.Active || card.Balance<=0 {
		continue
	}
	return payments

	fmt.Println(payment[0].Number)
	fmt.Println(payment[1].Number)
	fmt.Println(payment[2].Number)
	fmt.Println(payment[n].Number)

		
}