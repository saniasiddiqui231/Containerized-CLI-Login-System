# Containerized CLI Login System with 2FA

## Overview

A secure command-line authentication system built with Go. The application supports user registration, login, session management, account lockout protection, and optional TOTP-based Multi-Factor Authentication (MFA) compatible with Google Authenticator.

The project is containerized using Docker and stores data in SQLite with persistent storage.

---

## Features

### Authentication

* User registration
* User login
* Secure password hashing using bcrypt
* Password verification using bcrypt
* Last login tracking
* Input validation

### Account Security

* Account lockout after multiple failed login attempts
* Configurable lockout duration
* Protection against brute-force attacks

### Multi-Factor Authentication (MFA)

* Enable MFA
* Disable MFA
* TOTP-based authentication
* Compatible with Google Authenticator
* Compatible with Microsoft Authenticator
* TOTP verification during login

### Session Management

* Session creation after successful login
* Session expiration
* Session timeout configuration
* Logout support
* Session validation for protected commands

### Interactive CLI

* Interactive command prompt
* Command history support
* Tab completion
* Dynamic command suggestions
* Context-aware help command
* Clear success and error messages

### Docker Support

* Dockerized application
* Docker Compose support
* Persistent SQLite database storage
* Easy setup and execution

---

## Tech Stack

* Go
* SQLite
* Docker
* Docker Compose

### Go Libraries

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
│
├── internals/
│   ├── auth/
│   ├── cli/
│   ├── config/
│   ├── database/
│   ├── mfa/
│   ├── models/
│   └── session/
│
├── migrations/
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

| Column          | Description                  |
| --------------- | ---------------------------- |
| id              | Primary key                  |
| username        | Unique username              |
| password_hash   | bcrypt hashed password       |
| created_at      | Registration timestamp       |
| last_login      | Last successful login        |
| mfa_enabled     | MFA enabled flag             |
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

## Configuration

Application settings are centralized in the configuration package.

Examples:

* Maximum failed login attempts
* Lockout duration
* Session timeout duration

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

### Build

```bash
docker compose build
```

### Run Interactive CLI

```bash
docker compose run --rm app
```

### Stop Containers

```bash
docker compose down
```

### View Running Containers

```bash
docker ps
```

### View Logs

```bash
docker compose logs -f
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

## User Information Displayed After Login

After successful authentication the application automatically displays:

* Username
* Registration date
* MFA status
* Session expiration time
* Last login time

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

### Register

```text
> register

Username: john
Password: password123

Registration successful
```

### Login

```text
> login

Username: john
Password: password123

Login successful
```

### Enable MFA

```text
> enable-2fa
```

The application generates:

* Secret key
* Provisioning URL

Add the generated secret to Google Authenticator.

### Login With MFA

```text
> login

Username: john
Password: password123
TOTP Code: 123456
```

### View User Information

```text
> whoami
```

### Logout

```text
> logout
```

---

## Security Features

### Password Security

* Passwords are never stored in plaintext
* bcrypt hashing is used
* Password verification uses bcrypt comparison

### Account Lockout

* Failed login attempts are tracked
* Accounts are temporarily locked after repeated failures

### Multi-Factor Authentication

* TOTP secrets stored securely
* Google Authenticator compatible
* Additional authentication factor during login

### Session Security

* Session tokens generated securely
* Session expiration enforced
* Logout invalidates active session

---

## Persistence

SQLite database files are stored in a persistent Docker volume.

Data survives:

* Container restarts
* Docker Compose restarts
* Application rebuilds

---

## Testing Checklist

### Registration

* Create new user
* Prevent duplicate usernames

### Login

* Correct password succeeds
* Incorrect password fails

### Lockout

* Trigger account lockout
* Verify lockout expiration

### MFA

* Enable MFA
* Login with valid TOTP
* Reject invalid TOTP
* Disable MFA

### Sessions

* Create session
* Validate session
* Expire session
* Logout session

### CLI

* Command history
* Tab completion
* Help command
* Dynamic command visibility

### Docker

* Build successfully
* Run successfully
* Persist data across restarts


