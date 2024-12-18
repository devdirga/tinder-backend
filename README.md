## Project Structure

```
.

├── config
│   └── config.go        # Configuration settings (e.g., database, environment variables).
├── internal
│   ├── auth             # Authentication module (e.g., login, registration).
│   ├── profile          # User profile module.
│   ├── swipe            # Swipe functionality.
│   ├── match            # Matchmaking logic.
│   └── middleware       # Common middleware (e.g., authentication, logging).
├── pkg
│   └── utils            # Utility functions and helpers.
├── migrations
│   └── sql              # Database migration files.
├── test
│   └── unittest         # Integration tests.
├── go.mod               # Go module file.
├── go.sum               # Dependency lock file.
└── README.md            # Project documentation.
```

## Key Modules

### Authentication (`auth`)
Handles user registration, login, and session management. Utilizes JWT for secure authentication.

**Endpoints:**
- `POST /signup`: User registration.
- `POST /signin`: User login.
- `GET  /me`: Retrieve logged-in user details.
