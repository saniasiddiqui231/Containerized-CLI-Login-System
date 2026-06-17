# Containerized CLI Login System with Optional TOTP-Based 2FA

## Overview

This project is a secure command-line authentication system built in Go. It provides user registration, login, session management, account lockout protection, and optional TOTP-based Multi-Factor Authentication (MFA) compatible with Google Authenticator.

The application uses SQLite for persistent storage and can be run locally or inside Docker containers.

---

## Features

### Authentication

* User registration with username and password
* Secure password hashing using bcrypt
* Login with credential verification
* Last login tracking
* Username uniqueness enforcement

### Account Lockout Protection

* Failed login attempts are tracked
* Account lockout after multiple failed login attempts
* Configurable lockout duration
* Failed-attempt counter reset after successful authentication

### Multi-Factor Authentication (MFA)

* Optional TOTP-based MFA
* Compatible with Google Authenticator
* MFA enable/disable functionality
* TOTP verification during login

### Session Management

* Session creation after successful authentication
* Session persistence in database
* Configurable session timeout
* Session expiration handling
* Logout support

### Interactive CLI

* Interactive command-based interface
* Command history support
* Tab completion support
* Context-aware help command
* Clear success and error messages

### Persistence

* SQLite database storage
* Database migrations included
* Docker volume persistence
* Data survives container restarts

### Security Features

* bcrypt password hashing
* TOTP-based MFA
* Account lockout protection
* Session expiration
* Input validation
* No plaintext password storage

---

## Technology Stack

### Backend

* Go

### Database

* SQLite
* modernc.org/sqlite driver

### Security

* golang.org/x/crypto/bcrypt
* github.com/pquerna/otp

### Containerization

* Docker
* Docker Compose

---

## Project Structure

```text
.
├── cmd/
│   └── app/
│
├── internals/
│   ├── auth/
│   ├── cli/
│   ├── database/
│   ├── session/
│   ├── mfa/
│   ├── models/
│   └── config/
│
├── migrations/
│   └── 001_initial_schema.sql
│
├── data/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## Database Schema

### users

| Column          | Description            |
| --------------- | ---------------------- |
| id              | Primary key            |
| username        | Unique username        |
| password_hash   | bcrypt hashed password |
| created_at      | Registration timestamp |
| last_login      | Last successful login  |
| mfa_enabled     | MFA status             |
| totp_secret     | TOTP secret            |
| failed_attempts | Failed login counter   |
| locked_until    | Account lockout expiry |

### sessions

| Column     | Description                  |
| ---------- | ---------------------------- |
| id         | Session identifier           |
| user_id    | User reference               |
| token      | Session token                |
| created_at | Session creation timestamp   |
| expires_at | Session expiration timestamp |
| active     | Session status               |

---

## Configuration

The application uses configurable security settings:

| Setting                 | Default Value |
| ----------------------- | ------------- |
| Maximum Failed Attempts | 5             |
| Lockout Duration        | 15 Minutes    |
| Session Timeout         | 30 Minutes    |

---

## Running Locally

### Install Dependencies

```bash
go mod tidy
```

### Run Application

```bash
go run ./cmd/app
```

---

## Running with Docker

### Build and Start

```bash
docker compose up --build
```

### Stop Containers

```bash
docker compose down
```

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

## Example Workflow

### Register

```text
> register
Username: john
Password: ********

Registration successful
```

### Login

```text
> login
Username: john
Password: ********

Login successful
```

### Enable MFA

```text
> enable-2fa
```

Add the generated secret to Google Authenticator and save the account.

### Login with MFA

```text
> login
Username: john
Password: ********
TOTP Code: 123456

Login successful
```

### User Information

```text
> whoami
```

Displays:

* Username
* Registration date
* MFA status
* Session expiration time
* Last login time

---

## Security Considerations

* Passwords are never stored in plaintext.
* Passwords are hashed using bcrypt.
* Accounts are locked after repeated failed login attempts.
* TOTP secrets are used exclusively for MFA verification.
* Sessions automatically expire after the configured timeout.
* Authentication requires both password and TOTP code when MFA is enabled.

---

## Persistence

SQLite database files are stored in the `data/` directory.

When running with Docker, the directory is mounted as a volume, ensuring that user accounts and session data persist across container restarts and rebuilds.

---

## Future Improvements

* Unit and integration tests
* Audit logging
* Password reset functionality
* Role-based access control
* Refresh-token support
* Account recovery workflows
* Enhanced MFA enrollment experience

```
```
