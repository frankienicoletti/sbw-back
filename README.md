# Small Business Week Backend
==============================

### Setup for Front End

Create and start a postgres docker container:
```
$ docker pull postgres
$ docker run --name sbwdb -p 5432:5432 -e POSTGRES_USER=postgres -d postgres
```

Build docker container for this service from the binary:
```
$
```

### API

#### /accounts
Returns the account balances for all three types of accounts.

Returns:
```json
{
    "id": 1,
    "merchant": "Visa",
    "name": "Double Rewards",
    "type": "credit_card",
    "balance": -743.03,
    "last_updated": "2017-09-05T03:03:18.788217Z",
}
{
    "id": 2,
    "merchant": "Wells Fargo",
    "name": "checking", 
    "type": "checking",
    "balance": 10888.85,
    "last_updated": "2017-10-05T03:03:18.788217Z",
}
```

#### /bills
Returns the upcoming bills.

Returns:
```json
{
    "id": 1,
    "name": "Rent",
    "type": "Rent",
    "amount": 1500.00,
    "paid": false,
    "due_date": "2017-09-01T00:00:00.000000Z",
    "frequency": "monthly",
}
{
    "id": 2,
    "name": "Baking supplier",
    "type": "cogs",
    "balance": 600.00,
    "paid": true,
    "due_date": "2017-10-05T03:03:18.788217Z",
    "frequency": "ad_hoc",
}
```

#### /{account_id}/transactions
Returns the transactions for the account from the last 7 days.  

Accepts an id in the url and returns:
```json
{
    "id": 1,
    "description": "Staples",
    "type": "debit",
    "amount": 100.84,
    "date": "2017-11-05T03:03:18.788217Z",
}
{
    "id": 2,
    "description": "Payroll",
    "type": "debit",
    "amount": 6000.23,
    "date": "2017-10-05T03:03:18.788217Z",
}
```
