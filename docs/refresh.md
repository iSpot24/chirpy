# POST /api/refresh

Refresh JWT token of user.

## Request 

## Headers:
```yaml
Authorization: Bearer refresh_token
```

## Response:

#### 200 OK
```json
{
  "token": "5d5436af33dd684fc87ed76259ff53a084de0588aa3394148219cf862d7c61ab"
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
  "error": "Missing authorization"
}
```
