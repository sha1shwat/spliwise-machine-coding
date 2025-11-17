package functionality

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type SplitWiseApp struct {
	Transactions []Transaction
	UniqueUsers  mapset.Set[string]
}

var (
	txnNo           int
	SplitWiseAppIns *SplitWiseApp
)

type Transaction struct {
	txnId   string
	payee   string
	amount  float64
	paidFor string
}

func newSplitWiseApp() *SplitWiseApp {
	SplitWiseAppIns = &SplitWiseApp{
		Transactions: make([]Transaction, 0),
		UniqueUsers:  mapset.NewSet[string](),
	}
	txnNo = 0
	return SplitWiseAppIns
}

func GetSplitWiseAppIns() *SplitWiseApp {
	if SplitWiseAppIns == nil {
		SplitWiseAppIns = newSplitWiseApp()
	}
	return SplitWiseAppIns
}

func (app *SplitWiseApp) getAmtList(payee string, amt float64, userList []string, strategy string) []float64 {
	total := len(userList)
	flt := float64(total)
	eaamt := amt / flt

	amtSplit := make([]float64, 0)
	for i := 0; i < total; i++ {
		amtSplit = append(amtSplit, eaamt)
	}

	return amtSplit
}

func (app *SplitWiseApp) AddExpense(payee string, userList []string, amt float64) {

	dividedExpenses := app.getAmtList(payee, amt, userList, "EQ")

	app.UniqueUsers.Add(payee)

	for ind, user := range userList {
		txn := Transaction{
			payee:   payee,
			amount:  dividedExpenses[ind],
			paidFor: user,
			txnId:   fmt.Sprintf("%d_%d", txnNo, ind),
		}
		app.UniqueUsers.Add(user)
		app.Transactions = append(app.Transactions, txn)
	}

	txnNo++
}

func (app *SplitWiseApp) ShowExpense(userID string) map[string]float64 {

	expenses := make(map[string]float64)
	for _, txn := range app.Transactions {

		if txn.payee == txn.paidFor {
			continue
		}

		if txn.payee != userID && txn.paidFor != userID {
			continue
		}

		if txn.payee == userID {
			expenses[txn.paidFor] += txn.amount
		} else {
			expenses[txn.payee] -= txn.amount
		}

	}

	return expenses

}

type AllExpenseForUser struct {
	Payee    string
	Borrower map[string]float64
}

func (app *SplitWiseApp) ShowAllExpense() []AllExpenseForUser {
	allExpenseList := make([]AllExpenseForUser, 0)

	for uniqueUser := range app.UniqueUsers.Iter() {

		expense := app.ShowExpense(uniqueUser)
		allExpenseForUser := AllExpenseForUser{
			Payee:    uniqueUser,
			Borrower: expense,
		}

		allExpenseList = append(allExpenseList, allExpenseForUser)
	}

	return allExpenseList

}

func (app *SplitWiseApp) SettleDebtOptimised() {

}
