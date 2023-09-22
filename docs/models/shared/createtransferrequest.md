# CreateTransferRequest


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Amount`                                      | *int64*                                       | :heavy_check_mark:                            | amount (in cents) for the transaction         |
| `Currency`                                    | *string*                                      | :heavy_check_mark:                            | Currency associated with the amount           |
| `Direction`                                   | [Direction](../../models/shared/direction.md) | :heavy_check_mark:                            | Direction of the money                        |