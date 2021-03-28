package main


import  ("bank/pkg/bank/types")

func main (){

card := []types.PaymentSource{
    {
        ID: "5058 xxxx xxxx 9999",
        Amount: 999_99,
    },
    {
        ID: "5058 xxxx xxxx 8888",
        Amount: 888_88,
    },

}
}

Card := PaymentSouce(card)
fmt.Println(Card)
}