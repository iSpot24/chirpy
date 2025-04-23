# POST /api/chirps

Create a chirp.

## Request
```json
{
  "body": "This is a chirp"
}
```

## Headers:
```yaml
Content-Type: application/json
Authorization: Bearer jwt
```

## Response:

#### 201 Created
```json
{
  "id": "6823a8ce-2e0a-44ef-8087-15b4389c11d5",
  "email": "john.doe@mail.com",
  "user_id": "1234a8ce-2e0a-44ef-8087-15b4389c11d5",
  "created_at": "2025-04-22T19:18:10.908091Z",
  "updated_at": "2025-04-22T19:18:10.908091Z",
}
```

#### 400 Bad Request
```json
{
  "error": "Chirp is too long"
}
```

#### 401 Unauthorized
```json
{
  "error": "Missing authorization"
}
```
