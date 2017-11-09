# gotcmb
Exchange Rates Central Bank of Turkey Edit

## Usage

```
go get -u github.com/netinternet/gotcmb
```

## Exemple

```go
package main

import (
	"fmt"
	"gotcmb"
)

func main() {
	usd, err := gotcmb.Kur("USD")
	if err != nil {
		fmt.Println("Error: Could not get TCMB data")
		return
	}
	fmt.Printf("1 Dollar %v Turkey Lira\n", usd)
	
	euro, _ := gotcmb.Kur("EUR")
	fmt.Printf("1 Euro %v Turkey Lira\n", euro)
}


```

## Currency Code List

| Currency Code | Description |
| --- | --- |
| USD | 1 Dollar to Turkish Lira |
| EUR | 1 Euro to Turkish Lira |
| GBP | 1 Pound to Turkish Lira |
| AUD | 1 Australia Dollar to Turkish Lira |
| DKK | 1 Denmark Krone to Turkish Lira |
| CHF | 1 Switzerland Franc to Turkish Lira |
| SEK | 1 Sweden Krona to Turkish Lira |
| CAD | 1 Canada Dollar to Turkish Lira |
| KWD | 1 Kuwait Dinar to Turkish Lira |
| NOK | 1 Norway Krone to Turkish Lira |
| SAR | 1 Saudi Arabia Riyal to Turkish Lira |
| JPY | 1 Japanese Yen to Turkish Lira |
| BGN | 1 Bulgaria Lev to Turkish Lira |
| RON | 1 Romania Leu to Turkish Lira |
| RUB | 1 Russia Ruble to Turkish Lira |
| IRR | 1 Iran Rial to Turkish Lira |
| CNY | 1 China Yuan Renminbi to Turkish Lira |
| PKR | 1 Pakistan Rupee to Turkish Lira |