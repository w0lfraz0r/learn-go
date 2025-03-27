package main

import "fmt"

/*
Calculate Balance
We need to be able to calculate the cost of entire batches of messages at once.

Assignment
Using the given variables, write conditional statements to calculate and update the variables.

First, set finalCost to the bulkMessageCost. If the user is premium, then the discountRate should be applied to the finalCost. For example, a discountRate of 0.10 means the discounted price per message would be 90% of the original price.

Next, if the user has enough money in their accountBalance:

Process the payment by deducting the finalCost from their accountBalance
Print the purchaseSuccessMessage
Otherwise, just print the insufficientFundMessage.
*/

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line
	finalCost = bulkMessageCost
	if isPremiumUser {
		finalCost *= 1 - discountRate
	}
	var message = insufficientFundMessage
	if finalCost <= accountBalance {
		accountBalance -= finalCost
		message = purchaseSuccessMessage
	}
	fmt.Println(message)
	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}
