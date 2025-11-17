package main

import (
	"fmt"
	"kredivo/repo/functionality"
)

func main() {
	cmd := functionality.InitCommand()

	var (
		err error
	)

	err = cmd.RunCommnd("ADD_EXPENSE Alice 1200 4 Alice Bob Charlie David EQUAL")
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.RunCommnd("SHOW_BALANCE_ALL All")
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.RunCommnd("SHOW_BALANCE Alice")
	if err != nil {
		fmt.Println(err)
	}

}
