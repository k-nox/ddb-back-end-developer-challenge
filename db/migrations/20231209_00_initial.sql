-- +migrate Up
create table damage_type(
    type text unique primary key not null
);

create table defense_type(
    type text unique primary key not null,
    modifier real not null
);

create table character(
    character_id integer primary key autoincrement,
    name text not null unique,
    max_hit_points integer not null,
    current_hit_points integer not null,
    level integer not null,
    strength integer not null,
    dexterity integer not null,
    constitution integer not null,
    intelligence integer not null,
    wisdom integer not null,
    charisma integer not null
);

create table character_defense(
    character_defense_id integer primary key autoincrement,
    character_id integer not null,
    damage_type text not null,
    defense_type text not null,
    foreign key(character_id) references character(character_id),
    foreign key(damage_type) references damage_type(type),
    foreign key(defense_type) references defense_type(type)
);

-- +migrate Down
drop table character_defense;
drop table character;
drop table defense_type;
drop table damage_type;