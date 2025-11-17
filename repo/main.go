package main

import (
	"bufio"
	"fmt"
	"kredivo/repo/functionality"
	"os"
)

func main() {
	cmd := functionality.InitCommand()

	var (
		err error
	)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		err = cmd.RunCommnd(text)
		if err != nil {
			fmt.Println(err)
		}

	}

	//err = cmd.RunCommnd("ADD_EXPENSE Alice 1200 4 Alice Bob Charlie David EQUAL")
	//if err != nil {
	//	fmt.Println(err)
	//}
	////err = cmd.RunCommnd("SHOW_BALANCE_ALL All")
	////if err != nil {
	////	fmt.Println(err)
	////}
	//err = cmd.RunCommnd("SHOW_BALANCE Bob")
	//if err != nil {
	//	fmt.Println(err)
	//}

}
