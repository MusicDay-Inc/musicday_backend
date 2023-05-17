CREATE TABLE user_appid
(
    user_id UUID        not null PRIMARY KEY,
    app_id  UUID UNIQUE not null,
    constraint user_fk foreign key (user_id) references users (id)
);