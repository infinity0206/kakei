CREATE TABLE asignee (
  id bigint auto_increment primary key,
  name varchar(255) not null,
  created_at datetime,
  updated_at datetime
);

CREATE TABLE task (
  id bigint auto_increment primary key,
  name varchar(255) not null,
  type integer not null,
  asignee_id integer,
  start_date datetime,
  end_date datetime,
  deadline datetime,
  created_at datetime,
  updated_at datetime
);