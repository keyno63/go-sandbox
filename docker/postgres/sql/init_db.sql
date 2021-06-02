-- CREATE DATABASE IF NOT EXISTS development;

CREATE TABLE IF NOT EXISTS m_users(
    id VARCHAR(64),
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    primary key(id)
);

INSERT INTO m_users (id, first_name, last_name) VALUES ('1000', 'taro', 'yamada');
INSERT INTO m_users (id, first_name, last_name) VALUES ('1001', 'jiro', 'yamada');
INSERT INTO m_users (id, first_name, last_name) VALUES ('1002', 'hanako', 'yamada');
INSERT INTO m_users (id, first_name, last_name) VALUES ('2000', 'ichiro', 'tanaka');
