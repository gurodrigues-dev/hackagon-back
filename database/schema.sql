CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    point INTEGER
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE questions (
    question_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date DATE,
    level INTEGER,
    params1 TEXT,
    response1 TEXT,
    params2 TEXT,
    response2 TEXT,
    params3 TEXT,
    response3 TEXT
    );

CREATE TABLE answers (
    answers_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nickname VARCHAR(50) NOT NULL,
    questionid UUID NOT NULL,
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

