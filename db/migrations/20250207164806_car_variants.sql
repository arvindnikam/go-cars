-- migrate:up
create table car_variants (
    id int NOT NULL AUTO_INCREMENT,
    car_id integer,
    variant_code varchar(255) NOT NULL,
    variant_name varchar(255) NOT NULL,
    transmission varchar(255) NOT NULL,
    color varchar(255) NOT NULL,
    engine varchar(255) NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (`id`)
);


-- migrate:down
drop table car_variants;
