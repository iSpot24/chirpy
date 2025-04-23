# DELETE /api/chirps/{chirpID}

Delete a chirpy by its ID if authenticated user is the owner.

## Headers:
```yaml
Authorization: Bearer jwt
```

## Response:

#### 204 No Content

#### 400 Bad Request

#### 401 Unauthorized

#### 403 Forbidden