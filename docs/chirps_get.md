# GET /api/chirps?author_id=uuid&sort=desc

List all chirps. Queries `author_id` and `sort` are optional. Default sorting is `created_at asc`.

## Response:

#### 200 Ok
```json
[
  {
    "id": "6823a8ce-2e0a-44ef-8087-15b4389c11d5",
    "email": "john.doe@mail.com",
    "user_id": "1234a8ce-2e0a-44ef-8087-15b4389c11d5",
    "created_at": "2025-04-22T19:18:10.908091Z",
    "updated_at": "2025-04-22T19:18:10.908091Z",
  },
  {
    "id": "a5d23a8ce-2e0a-44ef-8087-15b4389c11e5",
    "email": "john.doe@mail.com",
    "user_id": "1234a8ce-2e0a-44ef-8087-15b438ewqe1d5",
    "created_at": "2025-04-22T19:18:10.908091Z",
    "updated_at": "2025-04-22T19:18:10.908091Z",
  },
]
```

#### 400 Bad Request
```json
{
  "error": "Something went wrong"
}
```
