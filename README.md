# Containerized CLI Login System with 2FA

## Overview

A secure command-line authentication system built with Go. The application supports user registration, login, session management, account lockout protection, and optional TOTP-based Multi-Factor Authentication (MFA) compatible with Google Authenticator.

The project is containerized using Docker and stores data in SQLite with persistent storage.

---

## Features
### Tests Screenshot

![Tests](Screenshot%20\(159\).png)


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

## Running with Docker

### Prerequisites

Before running the application, make sure the following are installed:

* Docker Desktop
* Docker Compose (included with modern Docker Desktop versions)

### Start Docker Desktop

Before running any Docker commands, ensure Docker Desktop is running.

You can verify Docker is available by running:

```bash
docker version
```

You should see both **Client** and **Server** information.

If you receive an error such as:

```text
failed to connect to the docker API
dockerDesktopLinuxEngine
```

Docker Desktop is not running. Start Docker Desktop and wait until it shows that the Docker Engine is running.

---

### Build the Application

From the project root directory:

```bash
docker compose build
```

This will:

* Download dependencies
* Build the Go application
* Create the Docker image

---

### Run the Interactive CLI

Because this is an interactive command-line application, use:

```bash
docker compose run --rm app
```

You should see:

```text
CLI Login System
Type 'help' to see available commands.
>
```

---

### Stop the Application

If the application is running, press:

```text
Ctrl + C
```

To remove containers:

```bash
docker compose down
```

---

### Persistence

User data is stored in a persistent SQLite database volume.

Data remains available after:

```bash
docker compose down
docker compose run --rm app
```

For example:

1. Register a user
2. Exit the application
3. Start the application again
4. Login using the same account

The user account should still exist.

---

### Common Commands

Build image:

```bash
docker compose build
```

Run application:

```bash
docker compose run --rm app
```

View running containers:

```bash
docker ps
```

Stop and remove containers:

```bash
docker compose down
```

View container logs:

```bash
docker compose logs
```

Run tests:

```bash
go test ./... -v
```


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
## Testing

The project includes unit tests for the following components:

* Authentication
* Account lockout logic
* Password hashing and verification
* MFA setup generation
* Session token generation
* Configuration validation

### Run All Available Tests

```bash
go test ./...
```

### Run Tests with Verbose Output

```bash
go test ./... -v
```

### Run Authentication Tests Only

```bash
go test ./internals/auth -v
```

**Note:** Some packages (such as `cmd/app`, `database`, `models`, and `cli`) currently do not contain unit tests and will appear as `[no test files]`.
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


