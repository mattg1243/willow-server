create table users (
	id uuid primary key not null,
	fname varchar (50) not null,
	lname varchar (50) not null,
	email varchar (255) unique not null,
	"hash" varchar (255) not null,
	nameForHeader varchar (255) not null, 
	license varchar (255),
	created_at timestamp not null,
	updated_at timestamp
);

create table user_contact_info (
	id uuid primary key not null,
	user_id uuid not null,
	foreign key (user_id) references users (id) on delete cascade,
	phone varchar(255),
	city varchar (255),
	"state" varchar (255),
	street varchar (255),
	zip varchar (255),
	paymentInfo json,
	created_at timestamp not null,
	updated_at timestamp
);

create table clients (
	id uuid primary key not null,
  user_id uuid not null,
	foreign key(user_id) references users (id) on delete cascade,
	fname varchar(255) not null,
	lname varchar(255),
	email varchar(255),
	phone varchar(255),
	balance int default 0 not null,
	balanceNotifyThreshold int default 0 not null,
	rate int not null,
	isArchived boolean default false,
	created_at timestamp not null,
	updated_at timestamp
);

create table events (
	id uuid primary key not null,
  client_id uuid not null,
 	foreign key (client_id) references clients (id) on delete cascade,
	date timestamp not null,
	duration decimal default 0,
	"type" varchar(255),
	"detail" text,
	rate int not null,
	amount decimal not null,
	newBalance decimal not null
);
