
This folder not important

Just place for save something need when code

Add test data
```sh
INSERT INTO account
    (name, email, bio)
VALUES
    ('win', 'win@gmail.com', ''),
    ('nguyen', 'nguyen@gmail.com', 'hello they');

INSERT INTO article
    (title, content, account_id)
VALUES
    ('welcome', 'Welcome everybody', 1),
    ('internet', 'this article show how the internet work', 2),
    ('virtual', 'all is virtual', 2)
RETURNING
    id, title;
```