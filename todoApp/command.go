package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo.")
	flag.StringVar(&cf.Add, "a", "", "Add a new todo.")
	flag.IntVar(&cf.Del, "del", -1, "Delete a existent todo by ID.")
	flag.IntVar(&cf.Del, "d", -1, "Delete a existent todo by ID.")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a existent todo by the ID. Syntax: ID:new_description Example: 1:\"Do something \"")
	flag.StringVar(&cf.Edit, "e", "", "Edit a existent todo by the ID. Syntax: ID:new_description Example: 1:\"Do something \"")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a existent todo.")
	flag.IntVar(&cf.Toggle, "t", -1, "Toggle a existent todo.")
	flag.BoolVar(&cf.List, "list", false, "List all todos.")
	flag.BoolVar(&cf.List, "l", false, "List all todos.")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use Syntax: ID:new_description. Example: 1:\"Do something \"")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Del != -1:
		todos.delete(cf.Del)

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	}

	if !cf.List {
		todos.print()
	}

}
