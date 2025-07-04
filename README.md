# Go Blog

## Railway Deployment

1. Push your code to GitHub.
2. Create a new Railway project and link your repo.
3. Add the PostgreSQL plugin.
4. In your Go service's Variables tab, set:
   - `DATABASE_URL` (copy from PostgreSQL plugin)
   - (Optional) `PORT` (Railway sets this automatically)
5. Deploy! Check logs for `Connected to PostgreSQL database.`
6. For local dev, use `.env` with `DB_PATH` and `PORT`.

## Local Development

- Create a `.env` file (not committed to git):
  ```
  PORT=8080
  DB_PATH=blog.db
  # DATABASE_URL=postgres://user:pass@host:port/dbname (for local Postgres)
  ```
- Run your app with `go run cmd/server/main.go`.

## 6. Notes
- The app will use PostgreSQL if `DATABASE_URL` is set, otherwise it falls back to SQLite.
- For production, always use PostgreSQL. 