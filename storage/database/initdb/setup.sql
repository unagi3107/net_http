CREATE USER user IDENTIFIED BY 'password';
GRANT all privileges ON gwp.* TO user@'%';
USE gwp;
CREATE TABLE posts (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  content text,
  author varchar(255)
);