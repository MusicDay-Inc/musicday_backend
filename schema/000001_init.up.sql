CREATE TABLE users
(
    id              UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    gmail           VARCHAR(255) not null UNIQUE,
    username        VARCHAR(30) UNIQUE,
    nickname        VARCHAR(30)  not null DEFAULT '',
    is_registered   BOOLEAN      not null DEFAULT false,
    has_picture     BOOLEAN      not null DEFAULT false,
    subscribers_c   INT          not null DEFAULT 0,
    subscriptions_c INT          not null DEFAULT 0

);

CREATE TABLE subscriptions
(
    subscriber_id   UUID not null,
    subscription_id UUID not null,
    constraint subscriber_fk foreign key (subscriber_id) references users (id),
    constraint subscription_fk foreign key (subscription_id) references users (id),
    constraint subscription_pk PRIMARY KEY (subscriber_id, subscription_id)
);

CREATE TABLE user_bios
(
    user_id UUID PRIMARY KEY not null,
    bio     VARCHAR(100),
    constraint user_bio_fk foreign key (user_id) references users (id)
);

CREATE TABLE authors
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) not null
);

CREATE TABLE songs
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author    VARCHAR(255),
    name      VARCHAR(510) not null,
    date      TIMESTAMP    not null,
    duration  TIME         not null,
    author_id UUID         not null,
    constraint author_fk foreign key (author_id) references authors (id)
);

CREATE TABLE albums
(
    id          UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    name        VARCHAR(510) not null,
    author      VARCHAR(255),
    date        DATE         not null,
    song_amount INT,
    duration    TIME         not null DEFAULT '0:0:0',
    author_id   UUID         not null,
    constraint author_fk foreign key (author_id) references authors (id)
);

CREATE TABLE album_songs
(
    album_id UUID        not null,
    song_id  UUID UNIQUE not null,
    constraint album_fk foreign key (album_id) references albums (id),
    constraint song_fk foreign key (song_id) references songs (id),
    constraint album_song_unique UNIQUE (album_id, song_id)
);

CREATE TABLE single_releases
(
    song_id UUID UNIQUE not null PRIMARY KEY,
    constraint song_fk foreign key (song_id) references songs (id)
);

CREATE TABLE reviews
(
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id          UUID          not null,
    is_song_reviewed BOOL          not null,
    release_id       UUID          not null,
    published_at     TIMESTAMP     not null,
    score            INT           not null,
    review_text      VARCHAR(2000) not null,
    constraint user_fk foreign key (user_id) references users (id),
    constraint user_review_unique UNIQUE (user_id, release_id)
);

-- CREATE TABLE stories
-- (
--     id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     user_id          UUID      not null,
--     background_color INT       not null,
--     song_id          UUID      not null,
--     published_at     TIMESTAMP not null,
--     story_text       VARCHAR(1200),
--     constraint user_fk foreign key (user_id) references users (id)
-- );

-- CREATE TABLE user_likes
-- (
--     user_id  UUID PRIMARY KEY not null,
--     story_id UUID             not null,
--     constraint user_fk foreign key (user_id) references users (id),
--     constraint story_fk foreign key (story_id) references users (id),
--     constraint user_story_unique UNIQUE (user_id, story_id)
--
-- );

insert into authors (id, name)
VALUES ('11f3417e-0d40-4dd7-b0a8-44ab13c4163b', 'The Police');

INSERT INTO albums (name, date, author_id)
VALUES ('Zenyatta Mondatta', '1980-09-03', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Don''t Stand So Close to Me', '1980-09-03', '00:04:04', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Driven to Tears', '1980-09-03', '00:03:20', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('When the World Is Running Down, You Make the Best of What''s Still Around', '1980-09-03', '00:03:38', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Canary in a Coalmine', '1980-09-03', '00:02:26', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Voices Inside My Head', '1980-09-03', '00:03:53', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Bombs Away', '1980-09-03', '00:03:06', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('De Do Do Do, De Da Da Da', '1980-09-03', '00:04:09', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Behind My Camel', '1980-09-03', '00:02:54', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Man in a Suitcase', '1980-09-03', '00:02:19', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Shadows in the Rain', '1980-09-03', '00:05:04', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');
INSERT INTO songs (name, date, duration, author_id) VALUES ('The Other Way of Stopping', '1980-09-03', '00:03:22', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

insert into authors (id, name)
VALUES ('3f69e4b8-2163-435e-8c48-28a867454ea2', 'King Crimson');

INSERT INTO albums (name, date, author_id)
VALUES ('In the Court of the Crimson King', '1969-10-10', '3f69e4b8-2163-435e-8c48-28a867454ea2');
INSERT INTO songs (name, date, duration, author_id) VALUES ('21st Century Schizoid Man', '1969-10-10', '00:07:20', '3f69e4b8-2163-435e-8c48-28a867454ea2');
INSERT INTO songs (name, date, duration, author_id) VALUES ('I Talk to the Wind', '1969-10-10', '00:06:05', '3f69e4b8-2163-435e-8c48-28a867454ea2');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Epitaph', '1969-10-10', '00:08:47', '3f69e4b8-2163-435e-8c48-28a867454ea2');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Moonchild', '1969-10-10', '00:12:11', '3f69e4b8-2163-435e-8c48-28a867454ea2');
INSERT INTO songs (name, date, duration, author_id) VALUES ('The Court of the Crimson King', '1969-10-10', '00:09:22', '3f69e4b8-2163-435e-8c48-28a867454ea2');

insert into authors (id, name)
VALUES ('d4650ed8-a74c-4340-8fe2-5d33eb723db7', 'Tame Impala');

INSERT INTO albums (name, date, author_id)
VALUES ('Currents', '2015-07-17', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Let It Happen', '2015-07-17', '00:07:46', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Nangs', '2015-07-17', '00:01:47', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('The Moment', '2015-07-17', '00:04:15', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Yes I''m Changing', '2015-07-17', '00:04:30', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Eventually', '2015-07-17', '00:05:19', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Gossip', '2015-07-17', '00:00:55', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('The Less I Know the Better', '2015-07-17', '00:03:38', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Past Life', '2015-07-17', '00:03:47', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Disciples', '2015-07-17', '00:01:48', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Cause I''m a Man', '2015-07-17', '00:04:02', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Reality in Motion', '2015-07-17', '00:04:12', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Love/Paranoia', '2015-07-17', '00:03:10', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');
INSERT INTO songs (name, date, duration, author_id) VALUES ('New Person, Same Old Mistakes', '2015-07-17', '00:06:03', 'd4650ed8-a74c-4340-8fe2-5d33eb723db7');


insert into authors (id, name)
VALUES ('075bf4b5-d8ce-4c00-967c-6593f575e025', 'Aphex Twin');

INSERT INTO albums (name, date, author_id)
VALUES ('Drukqs', '2001-10-22', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Jynweythek', '2001-10-22', '00:02:14', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Vordhosbn', '2001-10-22', '00:04:42', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Kladfvgbung Micshk', '2001-10-22', '00:02:00', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Omgyjya-Switch7', '2001-10-22', '00:04:46', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Strotha Tynhe', '2001-10-22', '00:02:03', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Gwely Mernans', '2001-10-22', '00:05:00', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Bbydhyonchord', '2001-10-22', '00:02:21', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Cock/Ver10', '2001-10-22', '00:05:17', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Avril 14th', '2001-10-22', '00:02:05', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Mt Saint Michel + Saint Michaels Mount', '2001-10-22', '00:08:02', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Gwarek2', '2001-10-22', '00:06:38', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Orban Eq Trx 4', '2001-10-22', '00:01:27','075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Aussois', '2001-10-22', '00:00:07', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Hy a Scullyas Lyf Adhagrow', '2001-10-22', '00:02:10', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Kesson Daslef', '2001-10-22', '00:01:18', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('54 Cymru Beats', '2001-10-22', '00:05:59', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Btoum-Roumada', '2001-10-22', '00:01:56', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Lornaderek', '2001-10-22', '00:00:30', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('QKThr', '2001-10-22', '00:01:27', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Meltphace 6', '2001-10-22', '00:06:14', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Bit 4', '2001-10-22', '00:00:18', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Prep Gwarlek 3b', '2001-10-22', '00:01:13', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Father', '2001-10-22', '00:00:51', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Taking Control', '2001-10-22', '00:07:08', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Petiatil Cx Htdui', '2001-10-22', '00:02:05', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Ruglen Holon', '2001-10-22', '00:01:45', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Afx237 v.7', '2001-10-22', '00:04:15', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Ziggomatic 17', '2001-10-22', '00:08:28', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Beskhu3epnm', '2001-10-22', '00:01:58', '075bf4b5-d8ce-4c00-967c-6593f575e025');
INSERT INTO songs (name, date, duration, author_id) VALUES ('Nanou2', '2001-10-22', '00:03:22', '075bf4b5-d8ce-4c00-967c-6593f575e025');


insert into authors (id, name)
VALUES ('075bf4b5-d8ce-4c00-967c-6593f575e025', 'Elvis Presley');


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
                                                      INNER JOIN albums a on s.author_id = a.author_id);

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
