# Command Line REPL — In-Memory Key-Value Store

A command-line REPL for an in-memory key-value store that supports nested transactions, rollback/commit, and basic commands.

## Installation

```bash
npm install
```

## How to Run

```bash
npm start
```

## How to Run Tests

```bash
npm test
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

## Transaction Approach

Transactions are managed as a stack of `Transaction` objects. Each call to `BEGIN` pushes a new `Transaction` onto the stack. Each `Transaction` wraps a `Map` that holds pending key-value changes for that layer.

Changes made inside a transaction always take priority over what is in the base store. Nothing is written to the base store until `COMMIT` is called.

When reading a value with `GET`, the stack is searched from the top (most recent transaction) down to the base store. This is a bottom-to-top traversal — the most recent change always wins and the search stops as soon as the key is found.

`COUNT` works differently. It builds a merged view of all keys by going from the bottom of the stack to the top, so newer transaction values overwrite older ones. The base store fills in any keys not covered by a transaction. This top-to-bottom merge ensures every key is accounted for exactly once with the correct value.

`COMMIT` applies all open transaction layers directly to the base store at once and clears the entire stack.

## Assumptions and Tradeoffs

### Transaction values take priority over the store
Any value set inside a transaction will shadow the value in the base store until the transaction is rolled back or committed. This is by design — transactions represent a pending set of changes that have not yet been made permanent.

### GET traverses bottom-to-top, COUNT merges top-to-bottom
`GET` searches from the newest transaction down because it only needs the single most recent value for a key. Stopping early makes it efficient. `COUNT` needs a complete picture of all keys so it builds a full merged map instead, with newer values overwriting older ones.

**Tradeoff:** Two different traversal strategies are used depending on the operation. This keeps each operation efficient for what it needs to do but means the two methods reason about the stack in opposite directions.

### COMMIT flattens all transactions
`COMMIT` applies all open transactions to the base store at once. This means you cannot partially commit — once `COMMIT` is called, all nested transactions are gone and cannot be rolled back.

**Tradeoff:** This is simpler to implement and reason about but it means nested transactions do not behave like true savepoints. If you have three transactions open and call `COMMIT`, all three are applied and none can be undone.

### NULL as a deletion sentinel
Inside a transaction, `DELETE` stores `null` as the value rather than removing the entry. This is how the store tracks that a key was deleted inside a transaction without touching the base store yet.

**Tradeoff:** This requires carefully distinguishing between a key that was explicitly deleted (`null`) and a key that was never set (`undefined`). A key set to `null` means deleted. A key returning `undefined` means it was never touched in that transaction layer.

### No data types
All keys and values are stored as strings. There is no support for numbers, booleans, or objects.

### No key validation