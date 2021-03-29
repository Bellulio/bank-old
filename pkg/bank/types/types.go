package types

//Money представляет собой денежную сумму в миимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

//Currency представляет код валюты
type Currency string

//Payment представляет информацию о платеже.
type Payment struct{
    ID int
    Amount Money
}

// Коды валют
const(
TJS Currency = "TJS"
RUB Currency = "RUB"
USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платежной карте
type Card struct {
	ID          int
	PAN         string
	Balance     int
	Currency    string
	Color       string
	Name        string
	Active      bool    
	}

type PaymentSource struct{
    Type string // 'card'
    Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money //баланс в дирамах
   
}
