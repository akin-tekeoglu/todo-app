CREATE TABLE todo (
    id integer primary key,
    title varchar(255),
    description text,
    completed boolean,
    event_date text,
    user_id integer
);