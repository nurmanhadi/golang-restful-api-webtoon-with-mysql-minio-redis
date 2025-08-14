## Pages

### 1. Add Bulk Page `POST api/pages`

- Header

    `Authorization: Bearer token`
    `Content-Type: multipart/form-data`

- Body

    | Field      | Type | Ext        | Required |
    |------------|------|------------|----------|
    | chapter_id | text |            | ✅       |
    | pages      | file | .jpg, .png | ✅       |

- Response Code

    `201` `400` `404`

### 2. Delete Page `DELETE api/pages/{pageID}`

- Header

    `Authorization: Bearer token`

- Response Code

    `200` `404`