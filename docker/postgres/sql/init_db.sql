CREATE DATABASE IF NOT EXISTS development;

CREATE TABLE IF NOT EXISTS m_users(
    id VARCHAR(64),
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    primary key(id)
);
