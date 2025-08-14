## GENRE

### 1. Add Genre `POST api/genres`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "name": "example",
    }
    ```

- Response Code

    `201` `400` `404`

### 2. Update Genre `POST api/genres/{genreID}`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "name": "example",
    }
    ```

- Response Code

    `200` `400` `404`

### 3. Delete Genre `DELETE api/genres/{genreID}`

- Header

    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`

### 4. Get All Genre `GET api/genres`

- Response Code

    `200` `404`

### 5. Get All Comic by genre Name `GET api/genres/{name}?page={page}&size={size}`

- Response Code

    `200` `404`
