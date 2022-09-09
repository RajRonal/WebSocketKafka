create table chat_app
(
    sender_id  text,
    receiver_id text,
    message text,
    sent_at  timestamp with time zone default now()
)