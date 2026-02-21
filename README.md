# Go + React Project Scaffold

A full-stack template repository with a React + TypeScript frontend and Go + Echo backend.

## Tech Stack

| Layer    | Technology                              |
| -------- | --------------------------------------- |
| Frontend | React 18, TypeScript, Vite, TailwindCSS |
| Backend  | Go 1.21+, Echo v4, MongoDB              |
| Auth     | JWT tokens with bcrypt password hashing |

## Features

- **Frontend**: Modern React SPA with TypeScript, Vite for fast builds, and TailwindCSS for styling
- **Backend**: RESTful API built with Echo framework
- **Database**: MongoDB for persistent storage
- **Authentication**: JWT-based auth with secure password hashing (bcrypt cost 14)
- **Middleware**: CORS, logging, and recovery middleware configured

## Prerequisites

- **Frontend**: Node.js 18+, npm
- **Backend**: Go 1.21+, Docker (for MongoDB)
- **Optional**: [`air`](https://github.com/cosmtrek/air) for live reload

## Quick Start

### 1. Clone and Setup

```bash
# Clone the repository
git clone <your-repo-url>
cd go-react-scaffold
```

### 2. Backend Setup

```bash
cd backend

# Copy environment file and configure
cp .env.example .env

# Start MongoDB via Docker
make docker-run

# Run the backend
make run
```

The backend will start on `http://localhost:8080`.

### 3. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will start on `http://localhost:5173` (default Vite port).

## Project Structure

```
.
├── frontend/                     # React + TypeScript frontend
│   ├── src/
│   │   ├── App.tsx              # Main app component
│   │   ├── main.tsx             # Entry point
│   │   └── index.css            # Tailwind directives
│   ├── package.json
│   ├── tailwind.config.js
│   ├── tsconfig.json
│   └── vite.config.ts
│
├── backend/                      # Go + Echo backend
│   ├── main.go                  # Application entry point
│   ├── auth/                    # Authentication handlers & JWT
│   │   ├── auth.go              # JWT token handling
│   │   ├── configs.go          # Auth configuration
│   │   ├── controller.go       # Auth business logic
│   │   └── routes.go           # Auth route handlers
│   ├── users/
│   │   └── model.go             # User data model
│   ├── configs/
│   │   ├── db.go                # MongoDB connection
│   │   └── env.go               # Environment variables
│   ├── integrations/             # External API integrations
│   │   ├── paypack.go
│   │   ├── telegram-bot.go
│   │   └── useplunk.go
│   ├── utils/
│   │   └── emails.go            # Email utilities
│   ├── Makefile
│   ├── go.mod
│   └── .env.example
│
├── AGENTS.md                     # Coding guidelines
└── README.md                     # This file
```

## Environment Variables

Configure the backend by editing `backend/.env`:

| Variable                | Description                        | Default                     |
| ----------------------- | ---------------------------------- | --------------------------- |
| `APP_ENV`               | Application environment            | `development`               |
| `PORT`                  | Server port                        | `8080`                      |
| `MONGODB_URI`           | MongoDB connection string          | `mongodb://localhost:27017` |
| `SESSION_KEY`           | Secret key for JWT signing         | Required                    |
| `REDIS_URL`             | Redis connection string            | Optional                    |
| `PAYPACK_CLIENT_ID`     | Paypack API client ID              | Optional                    |
| `PAYPACK_CLIENT_SECRET` | Paypack API secret                 | Optional                    |
| `USE_PLUNK`             | Enable Plunk email integration     | Optional                    |
| `TELEGRAM_BOT_ID`       | Telegram bot token                 | Optional                    |
| `TELEGRAM_CHAT_ID`      | Telegram chat ID for notifications | Optional                    |

## Development Commands

### Frontend

```bash
cd frontend

# Start development server with HMR
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Run linter
npm run lint
```

### Backend

```bash
cd backend

# Build the application
make build

# Run the application
make run

# Run tests
make test

# Live reload (requires air)
make watch

# Start MongoDB container
make docker-run

# Stop MongoDB container
make docker-down

# Clean build artifacts
make clean
```

## API Endpoints

| Method | Endpoint    | Description                         |
| ------ | ----------- | ----------------------------------- |
| `POST` | `/register` | Register a new user                 |
| `POST` | `/login`    | Authenticate user and get JWT token |

### Register User

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "yourpassword"}'
```

### Login User

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "yourpassword"}'
```

Response includes a JWT token for authenticated requests.

## Contributing

1. **Fork** the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Follow the coding guidelines in [AGENTS.md](AGENTS.md)
4. Run tests and lint before committing:

   ```bash
   # Backend
   cd backend && make test

   # Frontend
   cd frontend && npm run lint
   ```

5. Commit with a clear message
6. Push to your fork and submit a Pull Request

### Coding Standards

- **Frontend**: Strict TypeScript, ESLint configured, 2-space indentation
- **Backend**: Standard Go formatting (gofmt), proper error handling, context-based DB operations
- **Both**: Write tests for new features, validate inputs, never expose secrets

## Security

- Never commit `.env` files or secrets
- Passwords are hashed with bcrypt (cost 14)
- JWT tokens used for authentication
- CORS configured for development (adjust for production)
- Input validation on all endpoints

## License

MIT
