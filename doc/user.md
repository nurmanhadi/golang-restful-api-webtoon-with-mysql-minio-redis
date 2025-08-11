## USER

### 1. Register User `POST api/users/register`

- Header
    `Content-Type: application/json`

- Body

    ```json
    {
        "username": "example",
        "password": "example"
    }
    ```

- Response Code

    `201` `400` `404`

### 2. Login User `POST api/users/login`

- Header
    `Content-Type: application/json`

- Body

    ```json
    {
        "username": "example",
        "password": "example"
    }
    ```

- Response Code

    `200` `400` `404`

### 3. Upload Avatar `POST api/users/{userID}/avatar`

- Header
    `Authorization: Bearer token`
    `Content-Type: multipart/form-data`

- Body

    | Field  | Type | Ext       | Required |
    |--------|------|-----------|----------|
    | avatar | file | .jpg, .png| ✅       |

- Response Code

    `200` `400` `404`

### 4. Update User `PATCH api/users/{userID}`

- Header
    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "username": "example", // optional
        "old_password": "example", // optional, old_password & new_password most be match
        "new_password": "example" // optional, old_password & new_password most be match
    }
    ```

- Response Code

    `200` `400` `404`

### 5. Get User `GET api/users/{userID}`

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`

### 6. Add Admin `POST api/users/admins`

- Header
    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "username": "example",
        "password": "example"
    }
    ```

- Response Code

    `201` `400` `404`

### 7. Delete User `DELETE api/users/{userID}`

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`

### 8. Logout User `POST api/users/{userID}/logout`

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`

### 9. Get Total User `GET api/users/total?by={by}`
> by: enum(daily, weekly, monthly, all-time)
- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`