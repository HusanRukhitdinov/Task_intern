create table users
(
    id         uuid         not null primary key default gen_random_uuid(),
    status     varchar(50)  not null, -- [active, deleted]
    name       varchar(255) not null,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp    default CURRENT_TIMESTAMP,
    created_by uuid
);

create table otp
(
    id         uuid         not null primary key default gen_random_uuid(),
    email      varchar(255) not null,
    status     varchar(50)  not null, -- [unconfirmed, confirmed]
    code       varchar(50)  not null,
    expires_at timestamp,
    created_at timestamp    default CURRENT_TIMESTAMP
);

create table roles
(
    id         uuid         not null primary key default gen_random_uuid(),
    status     varchar(50)  not null, -- [active, deleted]
    name       varchar(255) not null unique,
    created_at timestamp    default CURRENT_TIMESTAMP,
    created_by uuid
);

create table sysusers
(
    id         uuid         not null primary key default gen_random_uuid(),
    status     varchar(50)  not null, -- [active, deleted]
    name       varchar(255) not null,
    phone      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp    default CURRENT_TIMESTAMP,
    created_by uuid
);

create table sysuser_roles
(
    id          uuid not null primary key default gen_random_uuid(),
    sysuser_id  uuid not null references sysusers(id),
    role_id     uuid not null references roles(id)
);
