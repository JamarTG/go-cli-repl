Summary
Build a small command-line REPL for an in-memory key–value store that supports nested transactions, rollback/commit, and basic commands.

Any language is allowed. 

Input/Output
Read commands from stdin, one per line.

Write responses to stdout.

Exit on END.

Commands
SET <key> <value>
Sets key to value.

GET <key>
Prints the current value, or NULL if unset.

DELETE <key>
Removes key if it exists. If it doesn’t exist, do nothing.

BEGIN
Starts a new transaction.

ROLLBACK
Reverts all changes in the most recent transaction.
If no transaction is open, print  NO TRANSACTION.

COMMIT
Permanently applies changes from all open transactions and closes them.
If no transaction is open, print  NO TRANSACTION.

END
Exits the program.

Output rules
Only these produce output:

GET → prints value or NULL

ROLLBACK when invalid → prints NO TRANSACTION

COMMIT when invalid → prints NO TRANSACTION

Everything else is silent.

Example
Input:

GET a
SET a foo
GET a
BEGIN
SET a bar
GET a
BEGIN
DELETE a
GET a
ROLLBACK
GET a
ROLLBACK
GET a
COMMIT
END
Output:

NULL
foo
bar
NULL
bar
foo
NO TRANSACTION
Deliverables
Source code

README with:

how to run

brief explanation of your transaction approach (a few paragraphs is fine)

any assumptions/tradeoffs

Bonus (optional)
Implement COUNT <value> → prints how many keys currently equal <value> (including changes inside transactions).