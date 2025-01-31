delete from users_events_actives where id is not null;

alter table users_events_actives
drop column minutes;

alter table users_events_actives
  rename column event_time to start;

alter table users_events_actives
alter column start type timestamp using start::timestamp;

alter table users_events_actives
  alter column start_time set default CURRENT_TIMESTAMP;

alter table users_events_actives
  rename column created_at to "end";

alter table users_events_actives
  alter column "end_time" set default current_timestamp;

drop index users_events_actives_user_id_event_id_event_time_key;

alter table users_events_actives
drop column event_id;

alter table users_events_actives
  rename to users_events_activities;


create or replace view users_activities_times_sums as
SELECT
  date_trunc('day', start_time) as time_day,
  sum((uea.end_time - uea.start_time)::interval)::varchar as time_sum,
  user_id
FROM users_events_activities uea
GROUP BY uea.user_id, date_trunc('day', start_time);


create or replace view unique_users_by_hour as
SELECT date_trunc('hour', start_time), count(distinct user_id) FROM users_events_activities
GROUP BY  date_trunc('hour', start_time);
