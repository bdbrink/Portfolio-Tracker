package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/piquette/finance-go/quote"
	googlesearch "github.com/rocketlaunchr/google-search"
	log "github.com/sirupsen/logrus"
)

func currentMarketData(ticker string) (string, float64, string) {

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

	return t.Symbol, t.Ask, t.ShortName

}

func addToPortfolio(symbol string, price float64) bool {

	fmt.Println("Is this in your Portfolio ? yes (y) or no (n)")

	var answer string

	// ask for input
	fmt.Scanln(&answer)

	// sort depending on user input
	if answer == "yes" || answer == "y" || answer == "Yes" {
		fmt.Println("Adding to spreadsheet")

		// ask for number of shares owned
		fmt.Println("How many shares do you own ?")
		var sharesOwned string
		fmt.Scanln(&sharesOwned)

		file, err := os.Create("stocks.csv")

		defer file.Close()

		if err != nil {
			log.Fatalln("failed to open file", err)
		}

		// get the current date and convert for the spreadsheet
		currentTime := time.Now()
		convertTime := fmt.Sprintf("%v", currentTime.Format("01-02-2006"))

		// format data to write to file, convert float to string
		convertPrice := fmt.Sprintf("%.2f", price)
		data := []string{convertTime, symbol, convertPrice, sharesOwned}

		// write data to the csv file
		w := csv.NewWriter(file)
		defer w.Flush()

		w.Write(data)

		return true

	} else if answer == "no" || answer == "n" || answer == "No" {
		fmt.Println("Okie Dokie, Warren Buffet wouldn't own it anyways")
	} else {
		fmt.Println("I have no idea how to respond to that")
	}
	return false
}

func readFile() {

	file, err := os.Open("stocks.csv")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(file)

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
		}
	}
}

func checkForFile() [][]string {

	file, err := os.Open("stocks.csv")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}
	return records
}

func howDoTheyDoIt(name string) {

	company := fmt.Sprintf("how does %v make money", name)
	searchResults, _ := googlesearch.Search(nil, company)

	// get top result from the search
	resultTop := searchResults[0]

	fmt.Printf("the top result on how %v, makes money: %v", name, resultTop)
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

	symbol, price, name := currentMarketData(ticker)
	howDoTheyDoIt(name)
	portfolio := addToPortfolio(symbol, price)

	if portfolio == true {
		checkForFile()
		readFile()
	}

}
