-- Disable foreign keys support
PRAGMA foreign_keys = off;

BEGIN TRANSACTION;

-- User schema
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  id  INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

-- Message schema
DROP TABLE IF EXISTS "messages";
CREATE TABLE "messages" (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  "sender_id" INTEGER REFERENCES "users"(id),
  "recipient_id" INTEGER REFERENCES "users"(id),
  "content_type" TEXT NOT NULL,
  "created_at"  TIMESTAMP
);

-- Text content schema
DROP TABLE IF EXISTS "texts";
CREATE TABLE "texts" (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  "message_id" INTEGER REFERENCES "messages"(id),
  "text" TEXT
);

-- Image content schema
DROP TABLE IF EXISTS "images";
CREATE TABLE "images"(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  "message_id" INTEGER REFERENCES "messages"(id),
  url TEXT NOT NULL,
  height INTEGER NOT NULL,
  width INTEGER NOT NULL
);

-- Video content schema
DROP TABLE IF EXISTS "videos";
CREATE TABLE "videos"(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  "message_id" INTEGER REFERENCES "messages"(id),
  url TEXT NOT NULL,
  source TEXT NOT NULL
);

COMMIT TRANSACTION;

-- Enable foreigh keys support
PRAGMA foreign_keys=on;