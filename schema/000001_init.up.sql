CREATE TABLE users
(
    id             uuid PRIMARY KEY      default gen_random_uuid(),
    gmail          varchar(255) not null unique,
    username       varchar(30) unique,
    nickname       varchar(30)  not null default '',
    is_registered  boolean      not null DEFAULT false,
    has_picture    boolean      not null DEFAULT false,
    subscribers_c   int          not null default 0,
    subscriptions_c int          not null default 0

);

CREATE TABLE subscriptions
(
    subscriber_id   uuid PRIMARY KEY not null,
    subscription_id uuid             not null,
    constraint subscriber_fk foreign key (subscriber_id) references users (id),
    constraint subscription_fk foreign key (subscription_id) references users (id),
    constraint subscription_unique UNIQUE (subscriber_id, subscription_id)
);

CREATE TABLE authors
(
    id   uuid PRIMARY KEY default gen_random_uuid(),
    name varchar(255) not null
);

CREATE TABLE songs
(
    id        uuid PRIMARY KEY default gen_random_uuid(),
    author    varchar(255),
    name      varchar(510) not null,
    date      date         not null,
    duration  time         not null,
    author_id uuid         not null,
    constraint author_fk foreign key (author_id) references authors (id)
);

CREATE TABLE albums
(
    id          uuid PRIMARY KEY      default gen_random_uuid(),
    name        varchar(510) not null,
    author      varchar(255),
    date        date         not null,
    song_amount int,
    duration    time         not null default '0:0:0',
    author_id   uuid         not null,
    constraint author_fk foreign key (author_id) references authors (id)
);

CREATE TABLE album_songs
(
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
    user_id        uuid       not null,
    has_background bool       not null,
    items          uuid array not null,
    published_at   timestamp  not null,
    story_text     varchar(1200),
    likes_amount   int,
    constraint user_fk foreign key (user_id) references users (id)
);

CREATE TABLE user_likes
(
    user_id  uuid PRIMARY KEY not null,
    story_id uuid             not null,
    constraint user_fk foreign key (user_id) references users (id),
    constraint story_fk foreign key (story_id) references users (id),
    constraint user_story_unique UNIQUE (user_id, story_id)

);

insert into authors (id, name)
VALUES ('11f3417e-0d40-4dd7-b0a8-44ab13c4163b', 'The Police');
-- 11f3417e-0d40-4dd7-b0a8-44ab13c4163b

INSERT INTO albums (name, date, author_id)
VALUES ('Zenyatta Mondatta', '1980-09-03',
        '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Don''t Stand So Close to Me', '1980-09-03', '00:04:04', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Driven to Tears', '1980-09-03', '00:03:20', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('When the World Is Running Down, You Make the Best of What''s Still Around', '1980-09-03', '00:03:38',
        '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Canary in a Coalmine', '1980-09-03', '00:02:26', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Voices Inside My Head', '1980-09-03', '00:03:53', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Bombs Away', '1980-09-03', '00:03:06', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('De Do Do Do, De Da Da Da', '1980-09-03', '00:04:09', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Behind My Camel', '1980-09-03', '00:02:54', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Man in a Suitcase', '1980-09-03', '00:02:19', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Shadows in the Rain', '1980-09-03', '00:05:04', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('The Other Way of Stopping', '1980-09-03', '00:03:22', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

insert into authors (id, name)
VALUES ('3f69e4b8-2163-435e-8c48-28a867454ea2', 'King Crimson');
-- 3f69e4b8-2163-435e-8c48-28a867454ea2

INSERT INTO albums (name, date, author_id)
VALUES ('In the Court of the Crimson King', '1969-10-10',
        '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('21st Century Schizoid Man', '1969-10-10', '00:07:20', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('I Talk to the Wind', '1969-10-10', '00:06:05', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Epitaph', '1969-10-10', '00:08:47', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('Moonchild', '1969-10-10', '00:12:11', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (name, date, duration, author_id)
VALUES ('The Court of the Crimson King', '1969-10-10', '00:09:22', '3f69e4b8-2163-435e-8c48-28a867454ea2');


UPDATE songs
SET author = (SELECT name
              FROM authors
              WHERE authors.id = songs.author_id);

UPDATE albums
SET author = (SELECT name
              FROM authors
              WHERE authors.id = albums.author_id);

INSERT INTO album_songs (album_id, song_id) (SELECT a.id, s.id
                                             FROM songs s
                                                      JOIN albums a on s.author_id = a.author_id);

UPDATE albums a
SET song_amount =
        (SELECT count(*)
         FROM album_songs
         WHERE a.id = album_songs.album_id);

-- SELECT s.*, SUM(s.duration)
-- FROM album_songs a_s
--          LEFT JOIN songs s on a_s.song_id = s.id
-- GROUP BY s.id;
UPDATE albums a
SET duration = (SELECT sq.album_duration
                FROM (SELECT SUM(s.duration) album_duration, s.author_id
                      FROM songs s
                      GROUP BY s.author_id) sq
                WHERE a.author_id = sq.author_id);
