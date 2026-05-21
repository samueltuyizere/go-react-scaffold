# Dead Code Analysis Report

**Generated:** 2026-05-21  
**Scope:** Full monorepo (frontend + backend)  
**Tools:** knip, depcheck, manual grep analysis

---

## Executive Summary

| Area      | Dead Files | Dead Functions | Unused Deps | Risk Level |
| --------- | ---------- | -------------- | ----------- | ---------- |
| Frontend  | 2          | 0              | 1 (eslint)  | LOW        |
| Backend   | 0          | 22             | 0           | MEDIUM     |
| **Total** | **2**      | **22**         | **1**       |            |

---

## Frontend Findings

### SAFE to Delete

| File                | Reason                                                               | Severity |
| ------------------- | -------------------------------------------------------------------- | -------- |
| `.prettierrc.mjs`   | Not referenced in any script or config; Prettier not in package.json | SAFE     |
| `src/vite-env.d.ts` | Ambient type declarations; not required for this minimal app         | SAFE     |

### Note: False Positives from depcheck

`depcheck` flagged `autoprefixer`, `postcss`, `tailwindcss`, and `eslint` as unused. These are **false positives** — they are used indirectly:

- `tailwindcss` + `autoprefixer` + `postcss` — configured in `postcss.config.js`, used via `@tailwind` directives in `src/index.css`
- `eslint` — configured in `eslint.config.js`, referenced in `lint` script

Only `eslint` itself shows as unused by knip because it's invoked via `npx` in scripts rather than as a direct import.

### Not Flagged (Active Code)

- `src/App.css` — empty file (0 bytes) but imported by `App.tsx`. Removing the import would be safe.
- `src/assets/react.svg` — used in `App.tsx`
- `public/vite.svg` — used in `App.tsx`

---

## Backend Findings

### DANGER — Unused Exported Functions

These functions are defined but **never called** anywhere in the codebase:

#### configs/env.go (8 functions)

| Function             | Status                    |
| -------------------- | ------------------------- |
| `AppEnv()`           | UNUSED                    |
| `EnvIsProd()`        | UNUSED                    |
| `GetRedisUrl()`      | UNUSED                    |
| `GetSessionKey()`    | USED by `auth/configs.go` |
| `EnvMongoURI()`      | USED by `configs/db.go`   |
| `EnvPort()`          | USED by `main.go`         |
| `GetPaypackSecret()` | UNUSED                    |
| `GetPaypackId()`     | UNUSED                    |
| `GetPlunkKey()`      | UNUSED                    |
| `TelegramBotId()`    | UNUSED                    |
| `TelegramChatID()`   | UNUSED                    |

#### configs/db.go (2 functions)

| Function              | Status            |
| --------------------- | ----------------- |
| `ConnectDB()`         | USED by `main.go` |
| `StoreRequestInDb()`  | UNUSED            |
| `UpdateRequestInDb()` | UNUSED            |

#### auth/auth.go (5 functions)

| Function                        | Status |
| ------------------------------- | ------ |
| `GetJWTSecret()`                | UNUSED |
| `GenerateTokensAndSetCookies()` | UNUSED |
| `generateAccessToken()`         | UNUSED |
| `generateToken()`               | UNUSED |
| `setTokenCookie()`              | UNUSED |
| `setUserCookie()`               | UNUSED |
| `JWTErrorChecker()`             | UNUSED |

#### auth/configs.go (1 variable)

| Variable       | Status |
| -------------- | ------ |
| `SessionStore` | UNUSED |

#### auth/controller.go (2 functions)

| Function             | Status                                         |
| -------------------- | ---------------------------------------------- |
| `processUserLogin()` | UNUSED (called by routes but routes not wired) |
| `createNewUser()`    | UNUSED (called by routes but routes not wired) |

#### auth/routes.go (2 functions)

| Function                   | Status            |
| -------------------------- | ----------------- |
| `HandleUserLogin()`        | USED by `main.go` |
| `HandleUserRegistration()` | USED by `main.go` |

Note: `HandleUserLogin` and `HandleUserRegistration` are wired in `main.go`, but internally call `processUserLogin` and `createNewUser` which call further unused functions. The entire auth flow appears incomplete/unused at runtime.

#### integrations/paypack.go (5 functions)

| Function                      | Status |
| ----------------------------- | ------ |
| `Authenticate()`              | UNUSED |
| `PaypackCashIn()`             | UNUSED |
| `PaypackCashOut()`            | UNUSED |
| `PollTransactionStatus()`     | UNUSED |
| `TestPollTransactionStatus()` | UNUSED |

#### integrations/telegram-bot.go (1 function)

| Function                | Status |
| ----------------------- | ------ |
| `SendTelegramMessage()` | UNUSED |

#### integrations/useplunk.go (1 function)

| Function               | Status |
| ---------------------- | ------ |
| `SendEmailWithPlunk()` | UNUSED |

#### users/model.go (3 functions)

| Function           | Status |
| ------------------ | ------ |
| `CreateNewUser()`  | UNUSED |
| `GetUserById()`    | UNUSED |
| `GetUserByPhone()` | UNUSED |
| `GetUserByEmail()` | UNUSED |

#### utils/emails.go (1 function)

| Function                     | Status |
| ---------------------------- | ------ |
| `SendOtpVerificationEmail()` | UNUSED |

---

## Recommendations

### Phase 1: SAFE Deletions (No Risk)

1. **Delete `.prettierrc.mjs`** — Not used by any script or config
2. **Delete `src/vite-env.d.ts`** — Not required for this minimal app
3. **Remove empty `src/App.css` import** from `App.tsx` (file is 0 bytes)

### Phase 2: CAUTION — Backend Unused Functions

The entire backend appears to be a scaffold with incomplete integration. The routes are wired but internal logic is not connected. Before deleting:

1. **Verify with user** if the integrations (Paypack, Telegram, Plunk) are planned for future use
2. **If scaffold is WIP**: Keep unused functions as they represent planned features
3. **If cleanup is desired**: Remove unused env helpers, integration functions, and DB utility functions

### Phase 3: DANGER — Not Recommended

- Do NOT delete `configs/env.go` or `configs/db.go` — core infrastructure
- Do NOT delete package files (`auth/`, `integrations/`, `utils/`) — they contain the scaffold structure
- Do NOT delete `users/model.go` — core domain model

---

## Test Verification Required

Before any deletion:

1. Run `cd frontend && npm run build` to verify frontend builds
2. Run `cd backend && go build ./...` to verify backend compiles
3. Re-run knip/depcheck after changes

---

## Files Analyzed

### Frontend (6 files)

- `src/App.tsx`
- `src/main.tsx`
- `src/App.css` (empty)
- `src/index.css`
- `src/vite-env.d.ts`
- `src/assets/react.svg`

### Backend (12 Go files)

- `main.go`
- `auth/auth.go`, `auth/configs.go`, `auth/controller.go`, `auth/routes.go`
- `configs/db.go`, `configs/env.go`
- `integrations/paypack.go`, `integrations/telegram-bot.go`, `integrations/useplunk.go`
- `users/model.go`
- `utils/emails.go`
