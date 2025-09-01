package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zevinto/expense-tracker/internal/service"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: expense-tracker <command> [arguments]\n")
		fmt.Println("Commands: add, update, delete, list, summary, set-budget, export")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		flag.Usage()
	}

	command := os.Args[1]
	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Description of the expense")
		amount := addCmd.String("amount", "", "Amount of the expense")
		addCmd.Parse(os.Args[2:])
		if *description == "" || *amount == "" {
			fmt.Println("Description and amount are required")
			addCmd.Usage()
			os.Exit(1)
		}
		service.AddExpense(*description, *amount)
	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.String("id", "", "ID of the expense to update")
		description := updateCmd.String("description", "", "New desciption")
		amount := updateCmd.String("amount", "", "New amount")
		updateCmd.Parse(os.Args[2:])
		if *id == "" {
			fmt.Println("ID is required")
			updateCmd.Usage()
			os.Exit(1)
		}
		service.UpdateExpense(*id, *description, *amount)
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.String("id", "", "ID of the expense to delete")
		deleteCmd.Parse(os.Args[2:])
		if *id == "" {
			fmt.Println("ID is required")
			deleteCmd.Usage()
			os.Exit(1)
		}
		service.DeleteExpense(*id)
	case "list":
		service.ListExpenses()
	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.String("month", "", "Month for summary(1-12)")
		summaryCmd.Parse(os.Args[2:])
		service.Summary(*month)
	case "set-budget":
		budgetCmd := flag.NewFlagSet("set-budget", flag.ExitOnError)
		month := budgetCmd.String("month", "", "Month for budget (1-12)")
		amount := budgetCmd.String("amount", "", "Budget amount")
		budgetCmd.Parse(os.Args[2:])
		if *month == "" || *amount == "" {
			fmt.Println("Month and amount are required")
			budgetCmd.Usage()
			os.Exit(1)
		}
		service.SetBudget(*month, *amount)
	case "export":
		service.ExportCSV()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Commands: add, update, delete, list, summary")
		os.Exit(1)
	}
}
