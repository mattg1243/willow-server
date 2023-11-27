create table users (
  id serial primary key,
  username varchar(128) not null,
  email varchar(128) not null,
  balance integer not null
);

create table artists (
  id serial primary key,
  "name" varchar(128) unique not null,
  birthday date not null
);

create table albums (
  id serial primary key,
  title varchar(128) not null,
  artist varchar(128) not null,
  price integer not null,
  foreign key (artist) references artists (name)
);

create table purchases (
  id serial primary key,
  "user" integer not null,
  album integer not null,
  "date" date not null,
  foreign key ("user") references users (id),
  foreign key (album) references albums (id)
);