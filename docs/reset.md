# POST /admin/reset

Delete all users with all associated records. In a real case scenario this would be behind a `superadmin` role.

## Response:

#### 200 Ok

#### 400 Bad Request
```json
{
  "error": "Something went wrong"
}
```

#### 403 Forbidden
