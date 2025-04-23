# POST /api/users

Create a user record.

## Request
```json
{
  "email": "john.doe@mail.com",
  "password": "password",
}
```

## Headers:
```yaml
Content-Type: application/json
```

## Response:

#### 201 Created
```json
{
  "id": "6823a8ce-2e0a-44ef-8087-15b4389c11d5",
  "email": "john.doe@mail.com",
  "is_chirpy_red": false,
  "created_at": "2025-04-22T19:18:10.908091Z",
  "updated_at": "2025-04-22T19:18:10.908091Z",
}
```

#### 400 Bad Request
```json
{
  "error": "Something went wrong"
}
```

#### 401 Unauthorized
```json
{
  "error": "Incorrect email or password"
}
```
