package main

import (
	"fmt"
	"gotcmb"
)

func main() {
	usd, err := gotcmb.Kur("USD")
	if err != nil {
		fmt.Println("Hata: Kur cekilemedi")
		return
	}
	fmt.Printf("1 Amerikan Doları %v Türk Lirasıdır\n", usd)
}
