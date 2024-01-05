create table users (
	id uuid primary key,
	username varchar (255) unique not null,
	fname varchar (50) not null,
	lname varchar (50) not null,
	email varchar (255) unique not null,
	salt text not null,
	hash text not null,
	city varchar (255),
	nameForHeader varchar (255), 
	phone varchar (255),
	"state" varchar (255),
	street varchar (255),
	zip varchar (255),
	license varchar (255),
	paymentInfo json,
	created_at timestamp not null,
	updated_at timestamp
);

create table clients (
	id uuid primary key,
	foreign key("user") references users (id),
	fname varchar(255) not null,
	lname varchar(255),
	email varchar(255),
	balance int default 0,
	balanceNotifyThreshold int default 0,
	rate int not null,
	isArchived boolean default false,
	created_at timestamp not null,
	update_at timestamp
);

create table events (
	id uuid primary key,
 	foreign key (client) references clients (id),
	date timestamp not null,
	duration decimal default 0,
	"type" varchar(255),
	"detail" text,
	rate int not null,
	amount decimal not null,
	newBalance decimal not null
);
