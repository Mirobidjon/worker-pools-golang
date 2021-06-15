CREATE TABLE IF NOT EXISTS exceldb (
    id uuid primary key not null,
    name varchar(255) not null,
    phone varchar(50) not null,
    parent_id uuid not null
);