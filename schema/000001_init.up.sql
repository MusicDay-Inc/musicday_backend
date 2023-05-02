CREATE TABLE users
(
    id            uuid PRIMARY KEY default gen_random_uuid(),
    nickname      varchar(50)  not null,
    username      varchar(50)  not null unique,
    gmail         varchar(255) not null unique,
    is_registered boolean,
    has_picture   boolean
);

CREATE TABLE songs
(
    id            uuid PRIMARY KEY default gen_random_uuid(),
    song_name     varchar(510) not null,
    song_author   varchar(255) not null,
    song_date     date         not null,
    song_duration time         not null
);

CREATE TABLE albums
(
    id           uuid PRIMARY KEY default gen_random_uuid(),
    album_name   varchar(510) not null,
    album_author varchar(255) not null,
    album_date   date         not null,
    song_amount  int,
    album_duration     time         not null
);

CREATE TABLE album_songs
(
    id       uuid PRIMARY KEY default gen_random_uuid(),
    album_id uuid        not null,
    song_id  uuid unique not null,
    constraint album_fk foreign key (album_id) references albums (id),
    constraint song_fk foreign key (song_id) references songs (id),
    constraint album_song_unique UNIQUE (album_id, song_id)
);


CREATE TABLE single_releases
(
    song_id uuid unique not null PRIMARY KEY,
    constraint song_fk foreign key (song_id) references songs (id)
);

CREATE TABLE song_reviews
(
    id           uuid PRIMARY KEY default gen_random_uuid(),
    user_id      uuid          not null,
    song_id      uuid          not null,
    published_at timestamp     not null,
    score        int8          not null,
    review_text  varchar(2000) not null,
    constraint song_fk foreign key (song_id) references songs (id),
    constraint user_fk foreign key (user_id) references users (id)
);

CREATE TABLE album_reviews
(
    id           uuid PRIMARY KEY default gen_random_uuid(),
    user_id      uuid          not null,
    album_id     uuid          not null,
    published_at timestamp     not null,
    score        int8          not null,
    review_text  varchar(2000) not null,
    constraint album_fk foreign key (album_id) references albums (id),
    constraint user_fk foreign key (user_id) references users (id)
);

CREATE TABLE reviews
(
    id               uuid PRIMARY KEY default gen_random_uuid(),
    user_id          uuid          not null,
    is_song_reviewed bool          not null,
    song_or_album_id uuid          not null,
    published_at     timestamp     not null,
    score            int8          not null,
    review_text      varchar(2000) not null,
    --     constraint song_fk foreign key (song_or_album_id) references songs (id)
    constraint user_fk foreign key (user_id) references users (id)
);

CREATE TABLE stories
(
    id             uuid PRIMARY KEY default gen_random_uuid(),
    has_background bool       not null,
    user_id        uuid       not null,
    items          uuid array not null,
    published_at   timestamp  not null,
    story_text     varchar(1200),
    constraint user_fk foreign key (user_id) references users (id)
);
