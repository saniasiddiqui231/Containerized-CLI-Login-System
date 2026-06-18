# Containerized CLI Login System with 2FA

## Overview

A secure command-line authentication system built with Go. The application supports user registration, login, session management, account lockout protection, and optional TOTP-based Multi-Factor Authentication (MFA) compatible with Google Authenticator.

The project is containerized using Docker and uses SQLite for persistent data storage.

---

## Features

### Authentication

* User registration and login
* bcrypt password hashing
* Last login tracking
* Input validation

### Security

* Account lockout after repeated failed login attempts
* Configurable lockout duration
* Protection against brute-force attacks

### Multi-Factor Authentication (MFA)

* Enable/Disable MFA
* TOTP-based authentication
* Compatible with Google Authenticator and Microsoft Authenticator
* TOTP verification during login

### Session Management

* Session creation after login
* Session expiration and timeout handling
* Logout support
* Session validation for protected commands

### Interactive CLI

* Interactive command prompt
* Command history
* Tab completion
* Context-aware help command
* Clear success and error messages

### Docker Support

* Dockerized application
* Docker Compose support
* Persistent SQLite storage

---

## Tech Stack

* Go
* SQLite
* Docker
* Docker Compose

### Libraries

* golang.org/x/crypto/bcrypt
* github.com/pquerna/otp
* github.com/c-bata/go-prompt
* modernc.org/sqlite

---

## Project Structure

```text
.
├── cmd/
│   └── app/
├── internals/
│   ├── auth/
│   ├── cli/
│   ├── config/
│   ├── database/
│   ├── mfa/
│   ├── models/
│   └── session/
├── migrations/
├── data/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## Database Schema

### users

| Column          | Description                  |
| --------------- | ---------------------------- |
| id              | Primary key                  |
| username        | Unique username              |
| password_hash   | bcrypt hashed password       |
| created_at      | Registration timestamp       |
| last_login      | Last successful login        |
| mfa_enabled     | MFA status                   |
| totp_secret     | TOTP secret                  |
| failed_attempts | Failed login counter         |
| locked_until    | Lockout expiration timestamp |

### sessions

| Column     | Description             |
| ---------- | ----------------------- |
| id         | Session identifier      |
| user_id    | Associated user         |
| token      | Session token           |
| created_at | Session creation time   |
| expires_at | Session expiration time |
| active     | Session status          |

---

## Local Development

### Clone Repository

```bash
git clone https://github.com/saniasiddiqui231/Containerized-CLI-Login-System.git
cd Containerized-CLI-Login-System
```

### Install Dependencies

```bash
go mod tidy
```

### Run Application

```bash
go run ./cmd/app
```

### Run Tests

```bash
go test ./... -v
```

Authentication tests only:

```bash
go test ./internals/auth -v
```

---

## Docker Setup

### Prerequisites

Ensure Docker Desktop is installed and running.

Verify Docker:

```bash
docker version
```

### Build Image

```bash
docker compose build
```

### Run Application

```bash
docker compose run --rm app
```

### Stop Containers

```bash
docker compose down
```

### Persistence Test

1. Register a user
2. Exit the application
3. Run the application again
4. Login with the same account

User data remains available because SQLite is stored in a persistent volume.

### Useful Docker Commands

```bash
docker compose build
docker compose run --rm app
docker compose down
docker ps
docker compose logs
```

**Note:** Tab completion and command history are implemented using `go-prompt`. These features work best when running the application directly in a terminal. Some Docker terminal environments may have limitations with advanced TTY features.

---

## Available Commands

### Guest Commands

```text
register
login
help
exit
```

### Authenticated Commands

```text
whoami
enable-2fa
disable-2fa
logout
help
exit
```

---

## User Information Displayed After Login

After successful authentication:

* Username
* Registration date
* MFA status
* Last login time
* Session expiration time

Example:

```text
Login successful

Username: john
Registration Date: 2026-06-17 18:16:15
MFA Status: Enabled
Last Login: 2026-06-17 18:16:31
Session Expiration: 2026-06-17 18:46:31
```

---

## Example Workflow

```text
register
↓
login
↓
enable-2fa
↓
logout
↓
login with TOTP
↓
whoami
↓
logout
```

---

## Testing

Unit tests cover:

* Authentication
* Password hashing and verification
* Account lockout logic
* MFA generation
* Session token generation
* Configuration validation

Run all tests:

```bash
go test ./... -v
```

Packages without tests will display:

```text
[no test files]
```

which is expected.

### Test Screenshot

![Tests](Screenshot%20\(159\).png)

---

## Security Features

* bcrypt password hashing
* Account lockout protection
* TOTP-based MFA
* Secure session tokens
* Session expiration enforcement
* Input validation

---

## Persistence

Application data is stored in SQLite and persists across:

* Container restarts
* Docker Compose restarts
* Application rebuilds

---

## Verification Checklist

* User registration
* Duplicate username prevention
* Successful login
* Failed login handling
* Account lockout
* MFA enable/disable
* MFA login verification
* Session creation and expiration
* Logout functionality
* Command history
* Tab completion
* Docker build and execution
* Persistent data storage

```
```
