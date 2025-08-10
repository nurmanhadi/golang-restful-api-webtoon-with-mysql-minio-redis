## COMIC-GENRE

### 1. Add Comic Genre `POST api/comic-genres`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "comic_id": 1,
        "genre_id": 1,
    }
    ```

- Response Code

    `201` `400` `404`

### 2. Delete Comic Genre `DELETE api/comic-genres/{comicGenreID}`

- Header

    `Authorization: Bearer token`

- Response Code

    `200` `404`