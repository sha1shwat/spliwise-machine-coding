package functionality

type SplitwiseApp struct {
	txnNo       int
	Expenses    map[int]Expense
	UniqueUsers []string
}

type Expense struct {
	PayD     string
	UserList []string
	Amts     []float64
}

var (
	SplitwiseAppIns *SplitwiseApp
)

func newSplitwiseApp() *SplitwiseApp {
	SplitwiseAppIns = &SplitwiseApp{
		txnNo:    0,
		Expenses: make(map[int]Expense),
	}
	return SplitwiseAppIns
}

func GetSplitwiseAppIns() *SplitwiseApp {
	if SplitwiseAppIns == nil {
		SplitwiseAppIns = newSplitwiseApp()
	}
	return SplitwiseAppIns
}

func (app *SplitwiseApp) getAmtList(payee string, amt float64, userList []string, strategy string) []float64 {
	total := len(userList) + 1
	flt := float64(total)
	eaamt := amt / flt

	amtSplit := make([]float64, 0)
	for i := 0; i < total; i++ {
		amtSplit = append(amtSplit, eaamt)
	}

	return amtSplit
}

func (app *SplitwiseApp) AddExpense(payee string, userList []string, amt float64) {

	app.Expenses[app.txnNo] = Expense{
		UserList: userList,
		Amts:     app.getAmtList(payee, amt, userList, "EQ"),
		PayD:     payee,
	}
	app.UniqueUsers = append(app.UniqueUsers, payee)
	app.UniqueUsers = append(app.UniqueUsers, userList...)

	app.txnNo++
}

func (app *SplitwiseApp) ShowExpense(userID string) map[string]float64 {

	expenses := make(map[string]float64)
	for _, expense := range app.Expenses {

		if userID == expense.PayD {
			for i, user := range expense.UserList {
				_, ok := expenses[user]
				if !ok {
					expenses[user] = 0
				}
				expenses[user] += expense.Amts[i]
			}
		} else {
			fnd := false
			ind := -1

			for i, user := range expense.UserList {
				if userID == user {
					fnd = true
					ind = i

					break
				}
			}

			if fnd {
				expenses[expense.PayD] -= expense.Amts[ind]
			}
		}
	}

	return expenses

}

type ExoenseALl struct {
	ProvidedUSer string
	HisExpense   map[string]float64
}

func (app *SplitwiseApp) ShowAllExpense() []ExoenseALl {
	xoenseALlList := make([]ExoenseALl, 0)
	vis := make(map[string]string, 0)
	for _, uniqueUsers := range app.UniqueUsers {
		_, ok := vis[uniqueUsers]
		if ok {
			continue
		}
		vis[uniqueUsers] = "vis"
		expense := app.ShowExpense(uniqueUsers)
		exoenseALl := ExoenseALl{
			ProvidedUSer: uniqueUsers,
			HisExpense:   expense,
		}

		xoenseALlList = append(xoenseALlList, exoenseALl)
	}

	return xoenseALlList

}
