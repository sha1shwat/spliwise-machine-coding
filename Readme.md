Problem Statement
Design and implement a program that facilitates expense sharing among a group of people,
accurately calculates balances, and facilitates settlements.

Requirements
Add Expenses:
Users should be able to add expenses.
Split Strategies:
Users should be able to split the expenses with a set of people using one of the following strategies:
Equal split: The amount is equally split among the set of people.
Example: 3 people, amount INR 300 → each owes INR 100.
Exact amount: The amount owed by each person is specified explicitly.
Example: 3 people, amount INR 300 → User1 owes 50, User2 owes 100, User3 owes 150.
Percentages: The percentage owed by each person is specified explicitly.
Example: 3 people, amount INR 300 → User1 owes 20% (INR 60), User2 owes 30% (INR 90), User3 owes 50% (INR 150).
Check Balances:
Users should be able to check how much they owe others and how much others owe them.
Check Net Balances:
Users should be able to see their net balance (positive means others owe them; negative means they owe others).
Settle Debt Optimally:
The program should suggest the minimum number of transactions required to settle debts.
Input and Output
The program will use following commands to interact with the user -

ADD_EXPENSE - A command used to add expenses with the following order
● Command - ADD_EXPENSE
● Name of the user adding expenses
● Total expense amount
● Number of friends to split the amount followed by names of the friends
● Split strategy - EQUAL, EXACT, PERCENT
● If the split strategy is EXACT or PERCENT, the split strategy will be followed by the
exact amount or percentage for all friends, respectively

SHOW_BALANCE_ALL
Display all balances showing who owes whom and how much.

SHOW_BALANCE <UserName>
Display balances related to the specified user.

SHOW_NET_BALANCE - A command to display net balance of all the users along with their debt and receivables.

Please check the format in the sample below. Positive means others owe
them, negative means they owe others.
e.g. With following balances,
> Bob owes Alice: 300
> Charlie owes Alice: 300
> Alice owes Charlie: 960
> The net balance for Alice is = 300 + 300 - 960 = -360
> The net balance for Bob is = -300
> The net balance for Charlie is = -300 + 960 = 660

SETTLE_DEBT_OPTIMIZED - Show minimum number of transactions to settle the debt between all users.

Each line of the output represents 1 transaction in the format - UserName1 pays UserName2 Amount
Sample Input
ADD_EXPENSE Alice 1200 4 Alice Bob Charlie David EQUAL
ADD_EXPENSE Bob 1500 3 Bob Eve Frank EXACT 0 500 1000
ADD_EXPENSE Charlie 2400 3 Alice David Charlie PERCENT 40 20 40
ADD_EXPENSE Eve 900 3 Eve Alice Bob EQUAL
ADD_EXPENSE Frank 1800 2 Frank David EXACT 0 1800

Sample Output
SHOW_BALANCE Bob
Bob → Alice: 300
Eve → Bob: 500
Frank → Bob: 1000
Bob → Eve: 300

SHOW_BALANCE ALL
Bob → Alice: 300
Charlie → Alice: 300
David → Alice: 300
Eve → Bob: 500
Frank → Bob: 1000
Alice → Charlie: 960
David → Charlie: 480
Alice → Eve: 300
Bob → Eve: 300
David → Frank: 1800

SHOW_NET_BALANCE
name | debt | receivables | balance
Alice | 1260 | 900 | -360
Bob | 600 | 1500 | 900
Charlie | 300 | 1440 | 1140
David | 2580 | 0 | -2580
Eve | 500 | 600 | 100
Frank | 1000 | 1800 | 800

SETTLE_DEBT_OPTIMIZED
David pays Charlie 1140
David pays Bob 900
David pays Frank 540
Alice pays Frank 260
Alice pays Eve 100
