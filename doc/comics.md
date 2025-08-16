## COMIC

### 1. Add Comic `POST api/comics`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "title": "example",
        "synopsis": "example", // optional
        "author": "axample",
        "artist": "axample",
        "type": "axample", // enum(manga,manhua,manhwa)
        "status": "axample", // enum(completed,hiatus,ongoing)
    }
    ```

- Response Code

    `201` `400` `404`

### 2. Update Comic `PATCH api/comics/{comicId}`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "title": "example", // optional
        "synopsis": "example", // optional
        "author": "axample", // optional
        "artist": "axample", // optional
        "type": "axample", // enum(manga,manhua,manhwa), optional
        "status": "axample", // enum(completed,hiatus,ongoing), optional
    }
    ```

- Response Code

    `200` `400` `404`

### 3. Delete Comic `DELETE api/comics/{comicId}`

- Header

    `Authorization: Bearer token`

- Body

    cover required

- Response Code

    `200` `404`

### 4. Upload Cover `POST api/comics/{comicId}/cover`

- Header

    `Authorization: Bearer token`
    `Content-Type: multipart/form-data`

- Body

    | Field | Type | Ext       | Required |
    |-------|------|-----------|----------|
    |cover  | file | .jpg, .png| ✅       |

- Response Code

    `200` `400` `404`

### 5. Get Comic By Slug `GET api/comics/{slug}`

- Response Code

    `200` `404`

### 6. Get Comic Recent `GET api/comics/recent?page={page}&size={size}`

- Response Code

    `200` `404`

### 7. Get Comic Popular `GET api/comics/popular?by={by}&page={page}&size={size}` <!--- SKIP DULU -->

> by: enum(daily, weekly, monthly, all-time)

- Response Code

    `200` `404`

### 8. Get Total Comic `GET api/comics/total`

- Header

    `Authorization: Bearer token`

- Response Code

    `200` `404`

### 9. Search Comic `GET api/comics/search?keyword={keyword}&page={page}&size={size}`
    
- Response Code

    `200` `404`

### 10. Get Comic by Type & Status `GET api/comics?type={type}&status={status}&page={page}&size={size}`

> type: enum(manga, manhua, manhwa), status: enum(completed, hiatus, ongoing)
    
- Response Code

    `200` `404`

### 11. Get Comic Related `GET api/comics/{slug}/related`
    
- Response Code

    `200` `404`

### 11. Get Comic New `GET api/comics/new`
    
- Response Code

    `200` `404`

### 12. Comic Add Genre `POST api/comics/:comicID/genre`

- Header

    `Authorization: Bearer token`
    `Content-Type: application/json`

- Body

    ```json
    {
        "genre_id": 1,
    }
    ```

- Response Code

    `201` `400` `404`