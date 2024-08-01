
BEGIN;

CREATE TABLE article (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content VARCHAR(1000),
    view INT DEFAULT 0,
    account_id INTEGER NOT NULL REFERENCES account (id) ON DELETE CASCADE
);

END;
