-- +migrate Up
alter table main.character add column temporary_hit_points integer;

-- +migrate Down
alter table main.character drop column temporary_hit_points;