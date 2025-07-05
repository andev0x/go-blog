-- Fix PostgreSQL schema for comments table
-- This script fixes the auto-incrementing issue

-- Drop the existing table if it exists (WARNING: This will lose existing data)
-- DROP TABLE IF EXISTS comments;

-- Create the table with proper PostgreSQL syntax
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    content TEXT NOT NULL,
    rating INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Alternative: If you want to keep existing data, you can alter the table
-- ALTER TABLE comments ALTER COLUMN id SET DEFAULT nextval('comments_id_seq');
-- ALTER TABLE comments ALTER COLUMN id SET NOT NULL; 