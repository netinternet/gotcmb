package gotcmb

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
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

func getTarihDate() (*Tarih_Date, error) {
	// The server does not respond on IPv6. Request it over IPv4 only.

	addr, err := net.ResolveIPAddr("ip4", "www.tcmb.gov.tr")
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, fmt.Errorf("no addresssed for www.tcmb.gov.tr")
	}
	ip := addr.String()

	transport := *http.DefaultTransport.(*http.Transport)
	defaultDialContext := transport.DialContext
	transport.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		return defaultDialContext(ctx, "tcp4", ip+":443")
	}

	client := &http.Client{Transport: &transport}

	res, err := client.Get("https://www.tcmb.gov.tr/kurlar/today.xml")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	kurData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	t := Tarih_Date{}
	if err := xml.Unmarshal(kurData, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func Kur(code string) (float64, error) {
	t, err := getTarihDate()
	if err != nil {
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

func ForexSelling() (map[string]float64, error) {
	t, err := getTarihDate()
	if err != nil {
		return nil, err
	}

	m := make(map[string]float64, len(t.Currency))
	for _, v := range t.Currency {
		value, err := strconv.ParseFloat(v.ForexSelling, 64)
		if err != nil || value == 0 {
			continue
		}
		m[v.Kod] = value
	}
	return m, err
}
