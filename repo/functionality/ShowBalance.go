package functionality

import "fmt"

type ShowBalance struct {
	ReqUser string
}

func (c *ShowBalance) ParseArgs(args string) error {
	c.ReqUser = args
	return nil
}

func (c *ShowBalance) Execute() error {
	spliWiseIns := GetSplitwiseAppIns()
	fnd := spliWiseIns.ShowExpense(c.ReqUser)

	user := c.ReqUser
	for us, expse := range fnd {
		fmt.Printf("%s -> %s = %v\n", user, us, expse)
	}

	return nil
}
