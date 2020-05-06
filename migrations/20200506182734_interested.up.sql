-- flattrackio.interested definition

begin;

create table if not exists interested (
  id text default md5(random()::text || clock_timestamp()::text)::uuid not null,
  email text not null,
  creationTimestamp int not null default date_part('epoch',CURRENT_TIMESTAMP)::int,
  modificationTimestamp int not null default date_part('epoch',CURRENT_TIMESTAMP)::int,
  deletionTimestamp int not null default 0,

  primary key (id)
);

comment on table interested is 'The interested table is used for storing emails people who are interested in FlatTrack once it is ready';

commit;
