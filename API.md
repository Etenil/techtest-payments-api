# API Design Document

_Guillaume Pasquet, 2019_

This API allows interation with payment resources. Payment resource comply
with the following JSON schema:

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Payment resource",
  "description": "A payment resource of the payments API",
  "type": "object",

  "properties": {
    "id": {
      "description": "Unique identifier of payment resource",
      "type": "integer"
    },
    "beneficiary": {
      "description": "Unique ID of the payment's beneficiary",
      "type": "integer"
    },
    "debtor": {
      "description": "Unique ID of the payment's originator",
      "type": "integer"
    },
    "amount": {
      "description": "Amount transfered from debtor to beneficiary",
      "type": "number",
      "minimum": 0,
      "exclusiveMinimum": true
    },
    "currency": {
      "description": "The 3-letter identifier of the payment currency",
      "type": "string"
    }
  }
}
```

_Note: For the purpose of this exercise, beneficiary and debtor ids are imaginary._

## Payments endpoint

### GET /payments

Retrieves the list of payments stored by the API.

No parameters are required.

Example response:

```json
[
  {
    "id": 1,
    "beneficiary": 1,
    "debtor": 2,
    "amount": 9.99,
    "currency": "USD"
  },
  {
    "id": 2,
    "beneficiary": 2,
    "debtor": 1,
    "amount": 2.87,
    "currency": "GBP"
  }
]
```

### POST /payments

Creates a new payment resource.

Body paramters:
* _beneficiary_: (required) integer, The beneficiary ID of the payment
* _debtor_: (required) integer, ID of the payment's initiator
* _amount_: (required) floating-point number, the amount being paid

Example request:
```
POST /payments

{
  "beneficiary": 1,
  "debtor": 2,
  "amount": 22.30
}
```

Example response:

```
{
  "id": 6,
  "beneficiary": 1,
  "debtor": 2,
  "amount": 22.30,
  "currency": "GBP"
}
```

### GET /payments/:id

Retrieves information about a single payment resource.

Parameter: The payment resource ID. Must be an integer and exist in the API's storage.

Example response:

```json
{
  "id": 2,
  "beneficiary": 2,
  "debtor": 1,
  "amount": 2.87,
  "currency": "GBP"
}
```


### PUT /payments/:id

Updates a payment resource.

Parameter: the payments resource ID to update.

Body paramters:
* _beneficiary_: (optional) integer, The beneficiary ID of the payment
* _debtor_: (optional) integer, ID of the payment's initiator
* _amount_: (optional) floating-point number, the amount being paid

Example request:
```
PUT /payments

{
  "amount": 12.79
}
```

Example response:

```
{
  "id": 6,
  "beneficiary": 1,
  "debtor": 2,
  "amount": 12.79,
  "currency": "GBP"
}
```

### DELETE /payments/:id

Deletes a payments resource.

Parameter: the payments resource ID to delete.

If successful, the API will respond with status 201 and an empty body.
