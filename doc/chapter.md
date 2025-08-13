## CHAPTER

### 1. Add Chapter `POST api/comics/{comicID}/chapters`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "number": 1,
    }
    ```

- Response Code

    `201` `400` `404`

### 2. Update Chapter `PATCH api/comics/{comicID}/chapters/{chapterID}`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "number": 1, // optional
        "publish": true, // optional
    }
    ```

- Response Code

    `200` `400` `404`

### 3. Delete Chapter `DELETE api/comics/{comicID}/chapters/{chapterID}`

- Header

    `Authorization: Bearer token`

- Response Code

    `200` `404`

### 4. Get Chapter by Slug & Number `GET api/comics/{slug}/chapters/{number}`

- Response Code

    `200` `404`