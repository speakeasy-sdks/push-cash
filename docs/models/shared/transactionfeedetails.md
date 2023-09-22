# TransactionFeeDetails


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Amount`                                                                      | *int64*                                                                       | :heavy_check_mark:                                                            | Push User ID associated with payment credential                               |
| `Currency`                                                                    | *string*                                                                      | :heavy_check_mark:                                                            | Currency associated with the amount                                           |
| `Description`                                                                 | *string*                                                                      | :heavy_check_mark:                                                            | Description of the fee                                                        |
| `Type`                                                                        | [TransactionFeeDetailsType](../../models/shared/transactionfeedetailstype.md) | :heavy_check_mark:                                                            | Type of the fee                                                               |