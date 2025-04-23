# POST /api/polka/webhooks

Simulate external service webhook.

## Request
```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "6823a8ce-2e0a-44ef-8087-15b4389c11d5"
  }
}
```

## Headers:
```yaml
Content-Type: application/json
Authorization: ApiKey key
```

## Response:

#### 204 No Content

#### 400 Bad Request
```json
{
  "error": "Something went wrong"
}
```

#### 400 Bad Request

#### 401 Unauthorized

#### 404 Not Found
