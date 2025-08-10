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

### 6. Get User `GET api/users/{userID}`

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`

### 7. Add Admin `POST api/users/admins`

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

    `200` `400` `404`

### 8. Delete Admin `DELETE api/users/{userID}/admins`

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`

### 9. Logout User `POST api/users/{userID}/logout`

- Header
    `Authorization: Bearer token`

- Response Code

    `200` `400` `404`