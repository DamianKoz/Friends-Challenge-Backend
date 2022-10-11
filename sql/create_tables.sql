CREATE TABLE user(
	id int PRIMARY KEY AUTO_INCREMENT, 
  user_level INT DEFAULT 1, 
  exp int DEFAULT 0,
  firstName varchar(50),
  lastName varchar(50),
  age int,
  gender text,
  email text,
  updatedAt timestamp,
  created_at timestamp 
);

CREATE TABLE task (
  id int PRIMARY KEY AUTO_INCREMENT,
  title  text,
   amount  int,
   amount_unit  text,
   duration  text,
   user_id  int
);

ALTER TABLE  task  ADD FOREIGN KEY ( user_id ) REFERENCES  user  ( id );