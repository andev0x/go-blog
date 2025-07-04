# Go Blog Railway Deployment Guide

## 1. Prerequisites
- Push your code to GitHub.
- Sign up at https://railway.app/.

## 2. Add PostgreSQL to Your Railway Project
- In Railway, create a new project and link your GitHub repo.
- Add the PostgreSQL plugin (Database > Add Plugin > PostgreSQL).
- Copy the `DATABASE_URL` provided by Railway.

## 3. Set Environment Variables
- In Railway dashboard, go to your project > Variables.
- Set the following:
  - `DATABASE_URL` (paste from Railway PostgreSQL plugin)
  - `PORT` (Railway sets this automatically, but you can set it to `8080` for local dev)

## 4. Deploy
- Railway will auto-detect the Procfile and deploy your app.
- Logs and deployments are visible in the Railway dashboard.

## 5. Local Development
- For local dev, use SQLite by setting `DB_PATH` in a `.env` file:
  ```
  PORT=8080
  DB_PATH=local.db
  ```
- For PostgreSQL local dev, set `DATABASE_URL` in your `.env` file instead of `DB_PATH`.

## 6. Notes
- The app will use PostgreSQL if `DATABASE_URL` is set, otherwise it falls back to SQLite.
- For production, always use PostgreSQL. 