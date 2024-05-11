-- +goose Up
create table train_programs
(
    id           serial primary key,
    program_name varchar(50)  not null,
    description  varchar(500) not null,
    status       integer      not null default 0
);

create table train_days
(
    id          serial primary key,
    day_name    varchar(50)  not null,
    description varchar(500) not null
);

create table exercises
(
    id            serial primary key,
    exercise_name varchar(50)  not null,
    pictures      varchar(500) not null,
    description   varchar(500) not null
);

create table sets
(
    id       serial primary key,
    quantity int              not null,
    weight   double precision not null
);

create table statistics
(
    id            serial primary key,
    time          timestamp        not null default now(),
    quantity      int              not null,
    weight        double precision not null,
    set_id        int              not null,
    exercise_id   int              not null,
    trains_day_id int              not null,
    program_id    int              not null
);

create table users_programs
(
    id         serial primary key,
    user_id    int not null,
    program_id int not null,
    foreign key (program_id) references train_programs (id) on delete cascade
);

create table trainers_programs
(
    id         serial primary key,
    trainer_id int not null,
    program_id int not null,
    foreign key (program_id) references train_programs (id) on delete cascade
);

create table users_trainers_programs
(
    id               serial primary key,
    user_id          int  not null,
    trainer_id       int  not null,
    program_id       int  not null default 0,
    is_user_agree    boolean not null default false,
    is_trainer_agree boolean not null default false,
    foreign key (program_id) references train_programs (id) on delete cascade
);

create table days_programs
(
    id            serial primary key,
    program_id    int not null,
    trains_day_id int not null,
    foreign key (program_id) references train_programs (id) on delete cascade,
    foreign key (trains_day_id) references train_days (id) on delete cascade
);

create table exercises_days
(
    id            serial primary key,
    trains_day_id int not null,
    exercise_id   int not null,
    foreign key (trains_day_id) references train_days (id) on delete cascade,
    foreign key (exercise_id) references exercises (id) on delete cascade
);

create table sets_exercises
(
    id          serial primary key,
    exercise_id int not null,
    set_id      int not null,
    foreign key (exercise_id) references exercises (id) on delete cascade,
    foreign key (set_id) references sets (id) on delete cascade
);

create table statistics_sets
(
    id           serial primary key,
    set_id       int not null,
    statistic_id int not null,
    foreign key (set_id) references sets (id) on delete cascade,
    foreign key (statistic_id) references statistics (id) on delete cascade
);

-- insert into train_programs (id, program_name, description, status)
-- VALUES (1, 'test program1', 'program description', 1),
--        (2, 'test program2', 'program description', 2),
--        (3, 'test program3', 'program description', 3),
--        (4, 'test program4', 'program description', 2),
--        (5, 'test program5', 'program description', 3);
--
-- insert into train_days (id, day_name, description)
-- VALUES (1, 'test day1', 'day description'),
--        (2, 'test day2', 'day description'),
--        (3, 'test day3', 'day description'),
--        (4, 'test day4', 'day description'),
--        (5, 'test day5', 'day description');
--
-- insert into exercises (id, exercise_name, description, pictures)
-- VALUES (1, 'test exercise1', 'exercise description', ' '),
--        (2, 'test exercise2', 'exercise description', ' '),
--        (3, 'test exercise3', 'exercise description', ' '),
--        (4, 'test exercise4', 'exercise description', ' '),
--        (5, 'test exercise5', 'exercise description', ' ');
--
-- insert into sets (id, quantity, weight)
-- VALUES (1, 5, 11.1),
--        (2, 11, 105.1),
--        (3, 7, 23.1),
--        (4, 9, 15.5),
--        (5, 15, 55.0);
--
-- insert into users_programs (id, user_id, program_id)
-- VALUES (1, 4, 1),
--        (2, 4, 2),
--        (3, 4, 3),
--        (4, 4, 4),
--        (5, 4, 5);
--
-- insert into trainers_programs (id, trainer_id, program_id)
-- VALUES (1, 4, 1),
--        (2, 4, 2),
--        (3, 4, 3),
--        (4, 4, 4),
--        (5, 4, 5);
--
-- insert into days_programs (id, program_id, trains_day_id)
-- VALUES (1, 1, 1),
--        (2, 1, 2),
--        (3, 3, 3),
--        (4, 3, 4),
--        (5, 5, 5);
--
-- insert into exercises_days (id, trains_day_id, exercise_id)
-- VALUES (1, 1, 1),
--        (2, 1, 2),
--        (3, 3, 3),
--        (4, 3, 4),
--        (5, 5, 5);
--
-- insert into sets_exercises (id, exercise_id, set_id)
-- VALUES (1, 1, 1),
--        (2, 1, 2),
--        (3, 3, 3),
--        (4, 3, 4),
--        (5, 5, 5);

-- +goose Down
drop table statistics_sets;
drop table sets_exercises;
drop table exercises_days;
drop table days_programs;
drop table users_programs;
drop table trainers_programs;
drop table users_trainers_programs;
drop table statistics;
drop table sets;
drop table exercises;
drop table train_days;
drop table train_programs;