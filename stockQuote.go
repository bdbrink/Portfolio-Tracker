package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/piquette/finance-go/quote"
	log "github.com/sirupsen/logrus"
)

func currentMarketData(ticker string) {

	t, _ := quote.Get(ticker)
	fmt.Printf("-- %v --\n", t.ShortName)

	// extract current market data for security
	fmt.Printf("Ticker: %v \n", t.Symbol)
	fmt.Printf("Current Price: $%v \n", t.Ask)
	fmt.Printf("52 week high: %v \n", t.FiftyTwoWeekHigh)
	fmt.Printf("52 week low: %v \n", t.FiftyTwoWeekLow)

	// get the upside from the current 52 week low
	upside := math.Trunc(t.FiftyTwoWeekLow / t.FiftyTwoWeekHigh * 100)
	fmt.Printf("Percent Upside: %v %% \n", upside)

}

func addToPortfolio(ticker string) {

	fmt.Println("Is this in your Portfolio ? yes or no")

	var answer string

	// ask for input
	fmt.Scanln(&answer)

	// sort depending on user input
	if answer == "yes" {
		fmt.Println("if you held for 5 years with 10k you would have 10k")
	} else if answer == "no" {
		fmt.Println("Would you like to purchase ?")
	} else {
		fmt.Println("I have no idea how to respond to that")
	}
}

func main() {

	// parse the ticker from the user
	flag.Parse()

	// log if no args is found
	if len(flag.Args()) == 0 {
		log.Fatalf("Input symbol for security to lookup")
	}

	// get info on the security
	ticker := flag.Args()[0]

	currentMarketData(ticker)
	addToPortfolio(ticker)

}
