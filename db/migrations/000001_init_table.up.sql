create table users
(
    id         uuid         not null
        primary key,
    status     varchar(50)  not null, -- [active, deleted]
    name       varchar(255) not null,
    email      varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamp,
    created_by uuid
);
create table otp
(
    id         uuid         not null
        primary key,
    email      varchar(255) not null,
    status     varchar(50)  not null, -- [unconfirmed, confirmed]
    code       varchar(50)  not null,
    expires_at timestamp
);

create table sysusers
(
    id         uuid         not null
        primary key,
    status     varchar(50)  not null, -- [active, deleted]
    name       varchar(255) not null,
    phone      varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamp,
    created_by uuid
);

create table sysuser_roles
(
    id         uuid not null
        primary key,
    sysuser_id uuid not null,
    role_id    uuid not null
);

create table roles
(
    id         uuid         not null
        primary key,
    status     varchar(50)  not null, -- [active, deleted]
    name       varchar(255) not null,
    created_at timestamp,
    created_by uuid
);
