SELECT * FROM users;

SELECT name, email, password
FROM users
WHERE email = "sabbir185@gmail.com";

SELECT name, city
FROM cities
WHERE area NOT IN(500, 1000) AND name = "Dhaka";

SELECT name, city
FROM cities
WHERE area BETWEEN 500 AND 1000;