CREATE TABLE IF NOT EXISTS users (
    id UUID,
    nickname TEXT PRIMARY KEY,
    email TEXT,
    password TEXT,
    points INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS questions (
    id UUID PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date DATE NOT NULL,
    level TEXT NOT NULL,
    params1 TEXT NOT NULL,
    response1 TEXT NOT NULL,
    params2 TEXT NOT NULL,
    response2 TEXT NOT NULL,
    params3 TEXT NOT NULL,
    response3 TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS answers (
    id UUID PRIMARY KEY NOT NULL,
    nickname TEXT NOT NULL,
    questionid UUID NOT NULL,
    status TEXT NOT NULL,
    created_at TEXT NOT NULL,
    FOREIGN KEY (questionid) REFERENCES questions(id)
);

CREATE TABLE IF NOT EXISTS cognito (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

