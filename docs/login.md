# POST /api/login

Authenticates a user and returns a JWT token.

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
```json
{
  "id": "6823a8ce-2e0a-44ef-8087-15b4389c11d5",
  "email": "jhon.doe@mail.com",
  "is_chirpy_red": false,
  "created_at": "2025-04-22T19:18:10.908091Z",
  "updated_at": "2025-04-22T19:18:10.908091Z",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHkiLCJzdWIiOiI2ODIzYThjZS0yZTBhLTQ0ZWYtODA4Ny0xNWI0Mzg5YzExZDUiLCJleHAiOjE3NDUzNTMwOTAsImlhdCI6MTc0NTM0OTQ5MH0.uL9-OjefZ3RygWvPvR6fhBZTarEfu9bTYqGz9NgcE5E",
  "refresh_token": "5d5436af33dd684fc87ed76259ff53a084de0588aa3394148219cf862d7c61ab"
}
```
