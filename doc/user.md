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
        "password": "example" // optional
    }
    ```

- Response Code

    `200` `400` `404`

### 5. Change Password `POST api/users/{userID}`

- Header
    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "old_password": "example",
        "new_password": "example" 
    }
    ```

- Response Code

    `200` `400` `404`