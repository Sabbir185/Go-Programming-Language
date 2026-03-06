UPDATE users
SET name = "Sabbir Ahmmed"
WHERE id = 1;

UPDATE users
SET email = "sabbir185@gmail.com"
WHERE id = 1
RETURNING *;