# Transfer


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Amount`                                                              | *int64*                                                               | :heavy_check_mark:                                                    | Amount of the transfer                                                |
| `ArrivalDate`                                                         | [types.Date](../../../types/date.md)                                  | :heavy_check_mark:                                                    | Date the funds will arrive, formatted "YYYY-MM-DD"                    |
| `Currency`                                                            | [shared.Currency](../../../pkg/models/shared/currency.md)             | :heavy_check_mark:                                                    | Currency associated with the amount                                   |
| `Direction`                                                           | [shared.Direction](../../../pkg/models/shared/direction.md)           | :heavy_check_mark:                                                    | Direction of the money                                                |
| `ID`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | The unique identifier assigned by Push, prefix is "transfer_"         |
| `Status`                                                              | [shared.TransferStatus](../../../pkg/models/shared/transferstatus.md) | :heavy_check_mark:                                                    | N/A                                                                   |