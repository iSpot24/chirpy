# POST /api/revoke

Revoke refresh token for user.

## Request 

## Headers:
```yaml
Authorization: Bearer refresh_token
```

## Response:

#### 204 No content

#### 400 Bad Request
```json
{
  "error": "Something went wrong"
}
```

#### 401 Unauthorized
```json
{
  "error": "Missing authorization"
}
```
