create schema go;

create sequence go.campaign_id_seq;

create table if not exists go.campaign
(
    id         int primary key default nextval('go.campaign_id_seq'),
    name       text,
    created_at timestamp default now(),
    content    text,
    email      text,
    contact_ID int,
    status     text
);

create sequence go.contact_id_seq;

create table if not exists go.contact
(
    id          int primary key default nextval('go.contact_id_seq'),
    email       text,
    campaign_id int,
    foreign key (campaign_id) references go.campaign (id)
);
