package functionality

import "fmt"

type ShowBalanceAll struct {
}

func (c *ShowBalanceAll) ParseArgs(args string) error {
	return nil
}

func (c *ShowBalanceAll) Execute() error {
	spliWiseIns := GetSplitwiseAppIns()
	fnd := spliWiseIns.ShowAllExpense()

	for _, ele := range fnd {
		user := ele.ProvidedUSer
		for us, expse := range ele.HisExpense {
			fmt.Printf("%s -> %s = %v\n", user, us, expse)
		}
	}

	return nil
}
