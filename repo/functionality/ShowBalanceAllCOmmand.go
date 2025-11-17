package functionality

import "fmt"

type ShowBalanceAll struct {
}

func (c *ShowBalanceAll) ParseArgs(args string) error {
	return nil
}

func (c *ShowBalanceAll) Execute() error {
	spliWiseIns := GetSplitWiseAppIns()
	fnd := spliWiseIns.ShowAllExpense()

	for _, ele := range fnd {
		payee := ele.Payee
		for borrower, expse := range ele.Borrower {
			if expse < 0.0 {
				continue
			}
			fmt.Printf("%s -> %s = %v\n", borrower, payee, expse)
		}
	}

	return nil
}
