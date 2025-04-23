# **Chirpy** 🐦

Chirpy is a lightweight Twitter-like microblogging API built in Go. It features user authentication, chirp (tweet) creation, and administrative endpoints. The backend provides a clean RESTful interface designed for learning or extending into a full-featured social app.

---

## 🚀 Features

🗣️ Post and manage Chirps

🔐 JWT-based user authentication & session management

👤 Create, update, and revoke users

🩺 Health check and metrics for monitoring

🧹 Admin endpoint to reset all users

🌐 Static file server with metrics middleware

🧪 Webhook integration example (Polka)

## 📦 Endpoints
### 🔧 System & Admin

* GET	/api/healthz   -> Health check
* GET	/admin/metrics ->	Application metrics (TODO, right now it's just a dummy endpoint)
* POST /admin/reset  ->	Delete all users (admin)

### 👤 Authentication & Users

* POST	[/api/login](docs/login.md) ->	 Log in and get JWT tokens
* POST	[/api/refresh](docs/refresh.md) ->	 Refresh JWT token
* POST	[/api/revoke](docs/revoke.md)	 ->  Revoke access token
* POST	[/api/users](docs/users_post.md)	 ->  Create new user
* PUT	  [/api/users](docs/users_put.md)	 ->  Update current user info
  
### 🐦 Chirps

* POST	[/api/chirps](docs/chirps_post.md)	-> Create a new chirp
* GET	[/api/chirps](docs/chirps_get.md)	-> List all chirps
* GET	[/api/chirps/{chirpID}](docs/chirp_get.md) -> Get chirp by ID
* DELETE	[/api/chirps/{chirpID}](docs/chirp_delete.md) -> Delete chirp by ID

### 📡 Webhooks

* POST	[/api/polka/webhooks](docs/polka.md)	-> Mark user as Chirpy Red (Polka)

### 🖼️ Static Files

* / ->	Serves static files
* /assets/	-> Serves static assets
All static routes are served via http.FileServer and wrapped with a metrics-incrementing middleware.

## 🛠 Requirements

Before you begin, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

Make sure your PostgreSQL server is running and accessible.

## 🚀 Installation

Install the `chirpy` using `go install`:

```bash
go install github.com/iSpot24/chirpy
```

