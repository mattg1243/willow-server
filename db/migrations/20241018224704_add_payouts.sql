-- migrate:up
create table if not exists users (
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

create table if not exists user_contact_info (
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

create table if not exists clients (
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


create table if not exists event_types (
	id uuid primary key not null,
	user_id uuid,
	foreign key (user_id) references users (id) on delete cascade,
	source varchar(50), -- default or custom
	name varchar(255) not null,
	charge boolean not null,
	created_at timestamp not null,
	updated_at timestamp
);

create table if not exists events (
	id uuid primary key not null,
	user_id uuid not null,
	foreign key (user_id) references users (id) on delete cascade,
  client_id uuid not null,
 	foreign key (client_id) references clients (id) on delete cascade,
	date timestamp not null,
	duration decimal default 0,
	event_type_id uuid not null,
	foreign key (event_type_id) references event_types (id),
	"detail" text,
	rate int not null,
	amount int not null,
	running_balance int not null
);

create INDEX if not exists idx_events_client_user_date
on events (client_id, user_id, date);

create table if not exists payouts (
	id uuid primary key not null,
	user_id uuid not null,
	foreign key (user_id) references users (id) on delete cascade,
	client_id uuid,
	foreign key (client_id) references clients (id) on delete cascade,
	date timestamp not null,
	amount int not null
);

create table if not exists payout_events (
	payout_id uuid not null,
	foreign key (payout_id) references payouts (id) on delete cascade,
	event_id uuid not null,
	foreign key (event_id) references events (id) on delete cascade,
	primary key (payout_id, event_id)
)

-- migrate:down
DROP TABLE events;
DROP TABLE clients;
DROP TABLE event_types;
DROP TABLE user_contact_info;
DROP TABLE users;
DROP TABLE payouts;
DROP TABLE payout_events;