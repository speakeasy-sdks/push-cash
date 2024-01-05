# PaymentCredential

The array will be empty until the user completes their first transaction. Additional transactions will utilize stored payment credentials


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Account`                                                         | [shared.Account](../../../pkg/models/shared/account.md)           | :heavy_check_mark:                                                | N/A                                                               |                                                                   |
| `BankName`                                                        | *string*                                                          | :heavy_check_mark:                                                | the name of the bank associated with the credential               | Space Coast Credit Union                                          |
| `Card`                                                            | [shared.Card](../../../pkg/models/shared/card.md)                 | :heavy_check_mark:                                                | N/A                                                               |                                                                   |
| `CreatedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | Datetime for when the payment credential was created for the user |                                                                   |