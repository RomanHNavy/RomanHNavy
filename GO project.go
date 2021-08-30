package main

import (
	"fmt"
)

func main() {
	// This slice/list will be used to count the number of correct answers.
	var mySlice []int

	//Question 1
	fmt.Println("")
	fmt.Println("1) Explain the 'The Disposition Effect' in behavioural finance.")
	fmt.Println("A - A reluctance to take losses and change behaviour.")
	fmt.Println("B - The tendency for people to mimic the financial behaviours of the majority of the herd.")
	fmt.Println("Type A or B")

	var response1 string
	fmt.Scan(&response1)
	if response1 == "A" || response1 == "a" {
		mySlice = append(mySlice, 1)
	}

	//Question 2
	fmt.Println("")
	fmt.Println("2) What is not included in GDP, out of the following?")
	fmt.Println("A - A newly built home.")
	fmt.Println("B - Trading of financial securities (Stocks & Bonds).")
	fmt.Println("Type A or B")

	var response2 string
	fmt.Scan(&response2)
	if response2 == "B" || response2 == "b" {
		mySlice = append(mySlice, 1)
	}

	//Question 3
	fmt.Println("")
	fmt.Println("3) If the variance of a stocks returns is large. Does this mean the returns are predicatbale or unpredicatable?")
	fmt.Println("A - Predictable")
	fmt.Println("B - Unpredictable")
	fmt.Println("Type A or B")

	var response3 string
	fmt.Scan(&response3)
	if response3 == "B" || response3 == "b" {
		mySlice = append(mySlice, 1)
	}

	//Question 4
	fmt.Println("")
	fmt.Println("4) The P/B (Price-To-Book) multiple is not a good ratio to use for what type of company?")
	fmt.Println("A - A fast growing company i.e Tech industry.")
	fmt.Println("B - A company with fewer assest i.e Service industries.")
	fmt.Println("Type A or B")

	var response4 string
	fmt.Scan(&response4)
	if response4 == "B" || response4 == "b" {
		mySlice = append(mySlice, 1)
	}

	// Question 5
	fmt.Println("")
	fmt.Println("5) The type of risk that can be driven towards zero as a fund invests in more and more securities is known as...")
	fmt.Println("A - Specific")
	fmt.Println("B - Systematic")
	fmt.Println("Type A or B")

	var response5 string
	fmt.Scan(&response5)
	if response5 == "A" || response5 == "a" {
		mySlice = append(mySlice, 1)
	}

	fmt.Println("")
	fmt.Println("The number of questions you answered correctly was:", len(mySlice), "out of 5.")

}
