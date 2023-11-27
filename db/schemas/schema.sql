create table users (
  id serial primary key,
  username varchar(128),
  email varchar(128),
  balance numeric(5,2)
);

create table artists (
  id serial primary key,
  "name" varchar(128) unique not null,
  birthday date not null
);

create table albums (
  id serial primary key,
  title varchar(128),
  artist varchar(128),
  price numeric(5,2),
  foreign key (artist) references artists (name)
);

create table purchases (
  id serial primary key,
  "user" integer,
  album integer,
  "date" date not null,
  foreign key ("user") references users (id),
  foreign key (album) references albums (id)
);