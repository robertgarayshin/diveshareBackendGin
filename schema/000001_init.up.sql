create type role as enum ('SELLER', 'RENTER');
create type status as enum ('FREE', 'BOOKED', 'RENTED');
CREATE TABLE IF NOT EXISTS users
(
    id                 serial PRIMARY KEY,
    user_role          role,
    username           varchar NOT NULL,
    password_hash      varchar NOT NULL,
    name               varchar NOT NULL,
    surname            varchar,
    confirmation_token varchar,
    created_at         timestamp,
    updated_at         timestamp,
    deleted_at         timestamp,
    email_is_confirmed bool
);

CREATE TABLE if not exists cars
(
    id            serial PRIMARY KEY,
    category      varchar,
    brand         varchar,
    model         varchar,
    owner_id      int,
    constraint owner_fk foreign key (owner_id) references users (id),
    price         varchar,
    produced_year varchar,
    status        status,
    car_rating    real,
    created_at    timestamp,
    updated_at    timestamp,
    deleted_at    timestamp
);



CREATE TABLE if not exists rents
(
    id         serial PRIMARY KEY,
    rent_begin timestamp,
    rent_end   timestamp,
    seller     int,
    constraint seller_fk foreign key (seller) references users (id),
    renter     int,
    constraint renter_kf foreign key (renter) references users (id)
);

CREATE TABLE if not exists reviews
(
    id              serial PRIMARY KEY,
    "from"          int,
    constraint from_fk foreign key ("from") references users (id),
    "to"            int,
    constraint to_fk foreign key ("to") references cars (id),
    review_datetime timestamp,
    review_rating   real,
    review_text     text
);