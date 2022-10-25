CREATE TABLE user(
	id          int PRIMARY KEY AUTO_INCREMENT, 
  first_name  VARCHAR(50) NOT NULL DEFAULT '',
  last_name   VARCHAR(50) NOT NULL DEFAULT '',
  user_level  INT DEFAULT 1, 
  exp int     DEFAULT 0,
  age int     NOT NULL DEFAULT 0,
  gender      text NOT NULL,
  email       text NOT NULL,
  updated_at  timestamp,
  created_at  timestamp 
);

-- Example Task:
-- {title: 'Pushup Challenge'}
-- {amount: 5}
-- {activity: pushups}
-- {duration: 5 Days}
-- {end_date: 2022-09-23}
-- Example: Read 5 Books per Week // Do 5 PushUps per Day

CREATE TABLE task (
  id          int NOT NULL AUTO_INCREMENT,
  title       varchar(350) NOT NULL DEFAULT NULL,
  verb VARCHAR(50) NOT NULL,
  amount      int NOT NULL DEFAULT '0',
  activity    varchar(300) NOT NULL DEFAULT NULL,
  duration    varchar(300) DEFAULT NULL,
  user_id     int NOT NULL DEFAULT NULL,
  created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  end_date    date DEFAULT NULL,
  time_unit   enum('day','week','month','year') NOT NULL DEFAULT 'day'
);

ALTER TABLE  task  ADD FOREIGN KEY ( user_id ) REFERENCES  user  ( id );