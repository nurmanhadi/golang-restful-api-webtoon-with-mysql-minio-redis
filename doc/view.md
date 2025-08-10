## VIEW

### 1. Add View `POST api/views`

- Header
    `Content-Type: application/json`

- Body

    ```json
    {
        "comic_id": 1,
    }
    ```

- Response Code

    `201` `400` `404`

### 1. Get Total View `GET api/views?by={by}`

> by: enum(daily, weekly, monthly, all-time)

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`