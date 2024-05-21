CREATE TYPE userRole AS ENUM ('admin', 'user', 'root');

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email VARCHAR(50) UNIQUE,
                       hashedPassword VARCHAR(50),
                       role userRole
);