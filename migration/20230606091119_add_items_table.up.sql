CREATE TABLE IF NOT EXISTS ITEMS (
    id bigserial PRIMARY KEY,
    name varchar(50) not null,
    description varchar(255),
    notes varchar(100),
    status varchar(20) not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null
);