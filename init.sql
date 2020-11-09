create table users (
  user_id SERIAL primary key,
  name varchar(100) not null
);

create table room (
  room_id SERIAL primary key,
  name varchar(50) not null,
  private boolean
);

create table participant (
  participant_id SERIAL primary key,
  user_id int references users (user_id),
  room_id int references room (room_id)
);

create table message (
  message_id SERIAL primary key,
  user_id int references users (user_id),
  room_id int references room (room_id),
  message varchar(500)
)