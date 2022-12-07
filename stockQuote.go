package main

import (
	"flag"
	"fmt"

	"github.com/piquette/finance-go/quote"
	log "github.com/sirupsen/logrus"
)

func main() {

	// parse the ticker from the user
	flag.Parse()

	// log if no args is found
	if len(flag.Args()) == 0 {
		log.Fatalf("Input symbol for security to lookup")
	}

	// get info on the security
	ticker := flag.Args()[0]
	t, _ := quote.Get(ticker)
	fmt.Printf("-- %v --\n", t.ShortName)

	// extract current market data for security
	fmt.Printf("Ticker: %v \n", t.Symbol)
	fmt.Printf("Current Price: $%v \n", t.Ask)
	fmt.Printf("52 week high: %v \n", t.FiftyTwoWeekHigh)
	fmt.Printf("52 week low: %v \n", t.FiftyTwoWeekLow)

}
