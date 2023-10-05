# CreateIntentRequest


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Amount`                                      | *int64*                                       | :heavy_check_mark:                            | amount (in cents) for the transaction         |
| `Currency`                                    | [Currency](../../models/shared/currency.md)   | :heavy_check_mark:                            | Currency associated with the amount           |
| `Direction`                                   | [Direction](../../models/shared/direction.md) | :heavy_check_mark:                            | Direction of the money                        |
| `UserID`                                      | *string*                                      | :heavy_check_mark:                            | Push's unique identifier for the user         |