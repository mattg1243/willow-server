alter table payout_events
drop constraint if exists payout_events_event_id_fkey;

alter table payout_events
add constraint payout_events_event_id_fkey
foreign key (event_id) references events(id) on delete restrict;