package functionality

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AddExpenseCommand struct {
	UserDPayment  string
	TotalAmt      float64
	DebtUsersList []string
	Strategy      string
}

// ADD_EXPENSE Alice 1200 4 Alice Bob Charlie David EQUAL

// Alice 1200 4 Alice Bob Charlie David EQUAL

func (c *AddExpenseCommand) ParseArgs(args string) error {

	var (
		err error
	)

	firstArgs := strings.SplitN(args, " ", 4)
	if len(firstArgs) != 4 {
		return errors.New("Invalid arguments")
	}

	c.UserDPayment = firstArgs[0]
	c.TotalAmt, err = strconv.ParseFloat(firstArgs[1], 64)
	if err != nil {
		fmt.Print("Parsing error: ", err)
		return err
	}

	numberOfUsers, err := strconv.ParseInt(firstArgs[2], 10, 64)
	if err != nil {
		return err
	}

	firstArgs = strings.Split(firstArgs[3], " ")

	n := len(firstArgs)

	if n != int(numberOfUsers+1) {
		return errors.New("Invalid command")
	}

	for i := 0; i < int(numberOfUsers); i++ {
		c.DebtUsersList = append(c.DebtUsersList, firstArgs[i])
	}

	c.Strategy = firstArgs[n-1]

	return nil

}

func (c *AddExpenseCommand) Execute() error {

	splitWiseIns := GetSplitwiseAppIns()
	splitWiseIns.AddExpense(c.UserDPayment, c.DebtUsersList, c.TotalAmt)
	fmt.Print("Expnse Added Successfully: ")
	return nil
}
