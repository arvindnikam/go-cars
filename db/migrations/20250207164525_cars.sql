-- migrate:up
create table cars (
    id int NOT NULL AUTO_INCREMENT,
    make varchar(255) NOT NULL,
    car_model varchar(255) NOT NULL,
    year int,
    body_type varchar(255) NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (`id`)
);

-- migrate:down
drop table cars;
