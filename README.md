# stay.tene.life MVP

Go + MariaDB + server-side templates app for mobile-first stay pages.

## Run locally
1. Create MariaDB database `stay_tene_life`.
2. Copy `.env.example` to `.env` and fill values.
3. Run:
   ```bash
   go mod tidy
   go run ./cmd/server
   ```

## Routes
- `GET /c/{token}` public guest card (noindex)
- `GET /login` login page (OAuth placeholders)
- `GET /admin` card dashboard
- `GET /admin/cards/new` new card form

## Notes
- Tokens are random URL-safe strings.
- After `valid_until`, public page returns neutral expiration page.
- Cleanup goroutine deletes records where `delete_after` has passed.
- OAuth for Google/Apple prepared as placeholders via `.env` credentials.
