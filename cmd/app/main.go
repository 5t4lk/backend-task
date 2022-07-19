package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

const (
	sourceLatestN   = "https://www.brentfordfc.com/api/incrowd/getnewlistinformation?count="
	sourceArticleID = "https://www.brentfordfc.com/api/incrowd/getnewsarticleinformation?id="
)

func main() {
	userChoiceAndLogic()
	askRepeat()
}

func userChoiceAndLogic() error {
	var choice int // creating variable for user's choice

	enter("Make your choice:\n1. Get a list of the latest [n] news articles\n2. Retrieve a single article by ID\n", &choice)

	if choice == 1 {
		var listLength string // user selected option #1, so we need to create variable for the length of news articles

		enter("Enter the length of the list: \n", &listLength)

		bytesData, _ := reqLatestArticles(listLength) // the body of response in var. bytesData

		var newListInformation NewListInformation

		err := xml.Unmarshal(bytesData, &newListInformation)
		if err != nil {
			return err
		}

		resultData, err := json.Marshal(newListInformation)
		if err != nil {
			return err
		}

		fmt.Print(string(resultData) + "\n") // printing in console the result of it.
	} else if choice == 2 {
		var articleID string // user selected option #2, so we need to create variable for the ID

		enter("Just enter ID of article: ", &articleID)

		bytesData, err := reqArticleByID(articleID) // the body of response in var. bytesData
		if err != nil {
			return err
		}

		var newListInformation MainData

		err = xml.Unmarshal(bytesData, &newListInformation)
		if err != nil {
			return err
		}

		resultDate, err := json.Marshal(newListInformation)
		if err != nil {
			return err
		}
		if newListInformation.NewsArticle.Title == "" {
			fmt.Print("Incorrect ID! Try again in 5 seconds!\n")
			time.Sleep(5 * time.Second)
			userChoiceAndLogic()
		} else {
			fmt.Print(string(resultDate) + "\n") // printing in console the result of it.
		}
	} else {
		fmt.Print("Incorrect input. Try again in 5 seconds!\n")
		time.Sleep(5 * time.Second)
		userChoiceAndLogic()
	}

	return nil
}

func enter(text, value interface{}) { // makes it easier to read the input.
	fmt.Print(text)
	fmt.Scan(value)
}

func askRepeat() { // for ease of use, the software asks the user if he wants to repeat the process
	var repeat int
	enter("Do you want to use the program one more time?\nWrite 1 if YES, else 0\n", &repeat)
	if repeat == 1 {
		main()
	}
	fmt.Print("Thanks for using.\n")
	os.Exit(1)
}
