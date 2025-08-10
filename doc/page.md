## Pages

### 1. Add Bulk Page `POST api/comics/{comicID}/chapters/{chapterID}/pages`

- Header

    `Authorization: Bearer token`
    `Content-Type: multipart/form-data`

- Body

    | Field | Type | Ext        | Required |
    |-------|------|------------|----------|
    |pages  | file | .jpg, .png | ✅       |

- Response Code

    `201` `400` `404`

### 2. Delete Page `DELETE api/comics/{comicID}/chapters/{chapterID}/pages/{pageID}`

- Header

    `Authorization: Bearer token`

- Response Code

    `200` `404`