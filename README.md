## Details on the structure of the service

### Authentication (`auth`)
Handles user registration, login, and session management. Utilizes JWT for secure authentication.

**Endpoints:**
- `POST /signup`: User registration.
- `POST /signin`: User login.
- `GET  /verification/:token`: Verification registration.

### User Profile (`profile`)
Manages user profiles, including viewing and editing personal information.

**Endpoints:**
- `GET /me`: Retrieve user profile.
- `POST /me`: Update user profile.

### Swipe Data (`swipe`)
Implements the swiping feature, allowing users to get profile.

**Endpoints:**
- `GET /swipe`: Load one data profile.

### Swiping (`swipe`)
Implements the swiping feature, allowing users to like or pass on other users.

**Endpoints:**
- `POST /swipe`: Like/Pass a user.

### Matchmaking (`match`)
Handles the logic for determining matches between users.
And send email notification


## Instructions on how to run the service

### Prerequisites
Ensure you have the following installed:
- **Go** (1.2 or later)
- **PostgreSQL** (v13 or later)
- **Git**
- **Docker** (optional, for message queue)

## Installation

### Clone the Repository
```bash
git clone https://github.com/devdirga/tinder-backend.git
cd tinder-service
```
---

### Environment Configuration
Create a `.config.json` file in the project root and populate it with the following variables:
```bash
{
  "DB": "host=localhost user=postgres password=mysecretpassword dbname=tinder port=5432 sslmode=disable",
  "IsDebug": true,
  "IsQueue": false,
  "IsConcurrent":true,
  "Secret": "secret",
  "GoogleSmtpKey": "azdg rkiv wnqe vuil ",
  "URL": "http://localhost:5000/"
}
```

### Backend Setup
Install Go dependencies:
```bash
go mod tidy
```

Run the backend service:
```bash
go run main.go
```
The service will run on `http://localhost:5000` by default.

## Deployment

CI/CD using Github Actions

Create new repository secrets to add your server access
```bash
VPS_HOST=
VPS_PRIVATE_KEY=
VPS_USER=
```

Setup your path in your server by change in file ```bash .github/workflows/deploy.yml  ```

Every time you commit to the repository, GitHub Actions will build and deploy to your server, using SSH to copy the Golang binary file.

## Details on the structure of the service
