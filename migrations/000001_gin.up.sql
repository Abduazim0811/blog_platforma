CREATE TABLE IF NOT EXISTS users_blog(
    id  serial PRIMARY KEY,
    user_name VARCHAR(100),
    email VARCHAR(100),
    passwoord VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS posts(
    id serial PRIMARY KEY,
    user_id int,
    title  VARCHAR(200),
    content VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
)