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
git clone https://github.com/yourusername/tinder-service.git
cd tinder-service
```
---
