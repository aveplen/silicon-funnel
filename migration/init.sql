create table if not exists chats (
    chat_id bigserial unique,
    tg_chat_id bigint
);

create table if not exists mailboxes (
    mailbox_id bigserial unique,
    chat_id bigint,

    name varchar,
    host varchar,
    port int,
    username varchar,
    password varchar,
    "offset" int,

    unique(chat_id, name, host, port, username),

    foreign key (chat_id) references chats(chat_id)
);
