# Expense Tracker CLI

A simple command-line interface (CLI) application written in Go to manage personal expenses. It allows users to add, update, delete, and view expenses, as well as generate summaries of expenses, including monthly summaries for the current year.

Source: https://roadmap.sh/projects/expense-tracker

## Features

- **Add Expense**: Record a new expense with a description and amount.
- **Update Expense**: Modify the description or amount of an existing expense.
- **Delete Expense**: Remove an expense by its ID.
- **List Expenses**: View all recorded expenses in a tabulated format.
- **Summary**: Display the total amount of all expenses or expenses for a specific month in the current year.

## Prerequisites

- Go 1.18 or higher
- No external library dependencies; uses only the Go standard library

## Installation

1. Clone or download the repository to your local machine.

2. Navigate to the project directory:

   ```bash
   cd expense-tracker
   ```

3. Build the application:

   ```bash
   go build -o expense-tracker
   ```

## Usage

Run the application using the `./expense-tracker` command followed by a subcommand. The application stores data in a file named `expenses.json` in the same directory.

### Commands

1. **Add an expense**:

   ```bash
   ./expense-tracker add --description "Lunch" --amount 20
   ```

   Output:

   ```
   Expense added successfully (ID: 1)
   ```

2. **Update an expense**:

   ```bash
   ./expense-tracker update --id 1 --description "Lunch at Cafe" --amount 25
   ```

   Output:

   ```
   Expense updated successfully
   ```

3. **Delete an expense**:

   ```bash
   ./expense-tracker delete --id 1
   ```

   Output:

   ```
   Expense deleted successfully
   ```

4. **List all expenses**:

   ```bash
   ./expense-tracker list
   ```

   Output:

   ```
   ID  Date       Description  Amount
   1   2025-09-01  Lunch        $20.00
   2   2025-09-01  Dinner       $10.00
   ```

5. **View total expenses**:

   ```bash
   ./expense-tracker summary
   ```

   Output:

   ```
   Total expenses: $30.00
   ```

6. **View expenses for a specific month**:

   ```bash
   ./expense-tracker summary --month 9
   ```

   Output:

   ```
   Total expenses for September: $30.00
   ```

## Data Storage

Expenses are stored in a JSON file (`expenses.json`) in the same directory as the executable. Each expense entry includes:

- `id`: Unique identifier for the expense.
- `date`: Date of the expense (format: YYYY-MM-DD).
- `description`: Description of the expense.
- `amount`: Amount spent.

Example `expenses.json`:

```json
[
  {
    "id": 1,
    "date": "2025-09-01",
    "description": "Lunch",
    "amount": 20
  },
  {
    "id": 2,
    "date": "2025-09-01",
    "description": "Dinner",
    "amount": 10
  }
]
```

## Error Handling

The application includes error handling for:

- Invalid amounts (e.g., negative numbers or non-numeric input).
- Invalid expense IDs (e.g., non-existent or non-numeric IDs).
- Invalid months (e.g., values outside 1-12).
- Missing required arguments for commands.

## Running the Application

Ensure the executable is in your PATH or run it directly from the project directory. For example:

```bash
./expense-tracker add --description "Coffee" --amount 5
```

## Notes

- The application assumes all expenses are in the current year for monthly summaries.
- The `expenses.json` file is created automatically when the first expense is added.
- Ensure write permissions in the directory where `expenses.json` is stored.

## Future Enhancements

- Add support for expense categories and filtering by category.
- Implement budget settings with warnings for exceeding budgets.
- Add functionality to export expenses to a CSV file.

## License

This project is licensed under the MIT License.
