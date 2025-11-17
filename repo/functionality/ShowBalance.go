package functionality

import (
	"fmt"
	"strings"
)

type ShowBalance struct {
	ReqUser string
}

func (c *ShowBalance) ParseArgs(args string) error {
	c.ReqUser = strings.TrimSpace(args)
	return nil
}

func (c *ShowBalance) Execute() error {
	spliWiseIns := GetSplitWiseAppIns()
	expenses := spliWiseIns.ShowExpense(c.ReqUser)

	user := c.ReqUser
	for borrower, expense := range expenses {
		if expense < 0 {
			fmt.Printf("%s -> %s = %v\n", user, borrower, expense*-1)
		} else {
			fmt.Printf("%s -> %s = %v\n", borrower, user, expense)
		}

	}

	return nil
}
