# FailureDetails

Failure details is non-null only for 'declined', 'error', 'timedout', 'chargedback'


## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `Code`                     | *int64*                    | :heavy_check_mark:         | The failure code           |
| `Description`              | *string*                   | :heavy_check_mark:         | Description of the failure |