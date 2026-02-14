# AGENTS.md - Coding Guidelines for Go + React Scaffold

## Project Structure

```
.
├── frontend/          # React + TypeScript + Vite + TailwindCSS
│   ├── src/
│   │   ├── App.tsx
│   │   ├── main.tsx
│   │   ├── index.css  # Tailwind directives
│   │   └── vite-env.d.ts
│   ├── package.json
│   ├── tsconfig*.json
│   └── eslint.config.js
│
├── backend/           # Go + Echo + MongoDB
│   ├── main.go
│   ├── auth/          # Auth handlers & JWT
│   ├── users/         # User models
│   ├── configs/       # DB & env config
│   ├── integrations/  # External APIs
│   ├── utils/         # Utilities
│   ├── Makefile
│   └── go.mod
│
└── README.md
```

## Build & Development Commands

### Frontend (from `frontend/` directory)

```bash
# Development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Lint code
npm run lint

# Install dependencies
npm install
```

### Backend (from `backend/` directory)

```bash
# Build the application
make build

# Run the application
make run

# Run all tests
make test

# Run a specific test file
# NOTE: No tests exist yet. When adding tests, use:
# go test ./path/to/package -run TestFunctionName -v

# Live reload (requires `air`)
make watch

# Docker compose for dependencies
make docker-run
make docker-down

# Clean build artifacts
make clean
```

### Root Level

```bash
# No root-level commands defined. Operate from frontend/ or backend/ subdirectories.
```

## Code Style Guidelines

### TypeScript / React (Frontend)

**Imports:**
- Use ES modules (`import` syntax)
- Group imports: React/core libraries first, then local imports
- Use path aliases if configured (none currently)

**Formatting:**
- No Prettier config - use consistent 2-space indentation
- Single quotes for strings
- Semicolons required

**Types:**
- Strict TypeScript enabled (`strict: true` in tsconfig)
- Enable all strict options: `noUnusedLocals`, `noUnusedParameters`, `noFallthroughCasesInSwitch`
- Use explicit return types for functions when not obvious

**Naming:**
- Components: PascalCase (e.g., `UserProfile.tsx`)
- Functions/variables: camelCase
- Constants: UPPER_SNAKE_CASE for true constants
- Files: PascalCase for components, camelCase for utilities

**React Patterns:**
- Use functional components with hooks
- Destructure props in function parameters
- Prefer `const` over `let`
- Use `useState`, `useEffect` from React (no class components)

**TailwindCSS:**
- Utility-first approach - use Tailwind classes directly
- Custom classes in `App.css` for component-specific styles
- Use `@apply` sparingly

**Error Handling:**
- Always check for null/undefined before accessing properties
- Use optional chaining (`?.`) when appropriate
- Handle fetch errors with try/catch

### Go (Backend)

**Imports:**
- Group imports: standard library first, then external packages, then local packages
- Use blank imports only when necessary (e.g., `godotenv/autoload`)
- Import path: `backend/<package>` for local packages

**Formatting:**
- Standard Go formatting (`gofmt` / `goimports`)
- Use `goimports` to organize imports
- Run `gofmt -w .` before committing

**Naming:**
- Packages: lowercase, single word (e.g., `auth`, `configs`)
- Exported: PascalCase (e.g., `ConnectDB`, `User`)
- Unexported: camelCase (e.g., `processUserLogin`)
- Acronyms: all caps (e.g., `MI`, `ID`, `URI`)
- Files: descriptive, lowercase with underscores if needed

**Types:**
- Use structs with tags for JSON/BSON/validation
- Tag format: `` `json:"field" bson:"field"` ``
- Prefer explicit types over `any`/`interface{}`

**Error Handling:**
- Always check errors: `if err != nil { return err }`
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Use `log.Fatalf()` only in main/init, not in libraries
- Return errors to caller, don't log and swallow

**Structure:**
- One package per directory
- Separate concerns: handlers in `routes.go`, logic in `controller.go`, models in `model.go`
- Main package in `main.go` at root

**Database (MongoDB):**
- Use context with timeouts: `context.TODO()` for now, `context.WithTimeout()` preferred
- Handle connection errors at startup
- Use `bson.M{}` for queries
- Close connections properly (deferred disconnect)

**HTTP (Echo):**
- Use middleware for CORS, logging, recovery
- Handler signature: `func(c echo.Context) error`
- Return JSON responses: `c.JSON(status, data)`
- Validate inputs before processing

## Testing (Not Yet Implemented)

When adding tests:

**Go:**
- Test files: `*_test.go`
- Run: `go test ./... -v` or `go test ./package -run TestName -v`
- Use testify for assertions (not yet included)
- Mock external dependencies

**TypeScript/React:**
- Use Vitest (aligned with Vite) or Jest
- Test files: `*.test.ts` or `*.test.tsx`
- Run single test: `npm test -- TestName`
- Use React Testing Library for component tests

## Environment Variables

Backend uses `.env` file (loaded automatically via `godotenv/autoload`):

```bash
APP_ENV=development
MONGODB_URI=mongodb://localhost:27017
PORT=8080
SESSION_KEY=your-secret-key
REDIS_URL=redis://localhost:6379
PAYPACK_CLIENT_ID=
PAYPACK_CLIENT_SECRET=
USE_PLUNK=
TELEGRAM_BOT_ID=
TELEGRAM_CHAT_ID=
```

Copy from `backend/.env.example` and fill in values.

## Git Workflow

1. Never commit `.env` files
2. Run `make test` and `npm run lint` before committing
3. Keep commits focused and atomic
4. Use conventional commit messages if possible

## Security Notes

- NEVER log or expose secrets, API keys, or passwords
- Hash passwords with bcrypt (cost 14 currently used)
- Validate all user inputs
- Use HTTPS in production
- Set secure session cookies
- Sanitize MongoDB queries to prevent injection
