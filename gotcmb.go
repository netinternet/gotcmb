package gotcmb

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/httplib"
)

type Tarih_Date struct {
	XMLName   xml.Name `xml:"Tarih_Date"`
	Tarih     string   `xml:"Tarih,attr"`
	Date      string   `xml:"Date,attr"`
	Bulten_No string   `xml:"Bulten_No,attr"`
	Currency  []Currency
}

type Currency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}

func Kur(code string) (float64, error) {
	kurData, err := httplib.Get("http://www.tcmb.gov.tr/kurlar/today.xml").Bytes()
	if err != nil {
		fmt.Println("unmarshal")
		return 0, err
	}
	t := Tarih_Date{}
	if err := xml.Unmarshal(kurData, &t); err != nil {
		fmt.Println("unmarshal")
		return 0, err
	}
	var kur string
	Req := strings.ToUpper(code)
	for i, v := range t.Currency {
		if t.Currency[i].Kod == Req {
			kur = v.BanknoteSelling
		}
	}
	format, err := strconv.ParseFloat(kur, 64)
	return format, err
}
