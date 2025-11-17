package functionality

import (
	"errors"
	"fmt"
	"strings"
)

type MenuCommand interface {
	ParseArgs(args string) error
	Execute() error
}

type Commands struct {
	args         []string
	MenuCommands map[string]MenuCommand
}

func InitCommand() *Commands {
	commands := new(Commands)
	commands.MenuCommands = make(map[string]MenuCommand)
	commands.MenuCommands["ADD_EXPENSE"] = new(AddExpenseCommand)
	commands.MenuCommands["SHOW_BALANCE_ALL"] = new(ShowBalanceAll)
	commands.MenuCommands["SHOW_BALANCE"] = new(ShowBalance)
	commands.MenuCommands["SETTLE_DEBT_OPTIMIZED"] = new(SettleDebtCommand)
	commands.MenuCommands["SHOW_NET_BALANCE"] = new(ShowNetBalance)

	return commands
}

func (commands *Commands) RunCommnd(args string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	arg := strings.SplitN(args, " ", 2)
	cmd, ok := commands.MenuCommands[arg[0]]

	if cmd == nil || !ok {
		return errors.New("Unknown command: " + arg[0])
	}

	fmt.Println("Executing command: " + arg[0])

	err := cmd.ParseArgs(arg[1])
	if err != nil {
		fmt.Println("Error parsing command: " + err.Error())
		return err
	}
	return cmd.Execute()

}
