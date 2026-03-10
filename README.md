# Command Line REPL — In-Memory Key-Value Store

A command-line REPL for an in-memory key-value store that supports nested transactions, rollback/commit, and basic commands.
## Installation

```bash
go mod download
```

## How to Run

```bash
go run .
```

## How to Run Tests

```bash
go test ./...
```

## Commands

| Command | Description | Output |
|---|---|---|
| `SET <key> <value>` | Sets key to value | Silent |
| `GET <key>` | Prints the current value | Value or `NULL` |
| `DELETE <key>` | Removes key if it exists | Silent |
| `COUNT <value>` | Prints how many keys equal value | Count |
| `BEGIN` | Starts a new transaction | Silent |
| `ROLLBACK` | Reverts the most recent transaction | Silent or `NO TRANSACTION` |
| `COMMIT` | Permanently applies all open transactions | Silent or `NO TRANSACTION` |
| `END` | Exits the program | Silent |