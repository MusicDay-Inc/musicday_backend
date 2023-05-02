insert into authors (id, author_name)
VALUES ('11f3417e-0d40-4dd7-b0a8-44ab13c4163b', 'The Police');
-- 11f3417e-0d40-4dd7-b0a8-44ab13c4163b
-- The Police

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Don''t Stand So Close to Me', '1980-09-03', '00:04:04', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Driven to Tears', '1980-09-03', '00:03:20', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('When the World Is Running Down, You Make the Best of What''s Still Around', '1980-09-03', '00:03:38',
        '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Canary in a Coalmine', '1980-09-03', '00:02:26', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Voices Inside My Head', '1980-09-03', '00:03:53', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Bombs Away', '1980-09-03', '00:03:06', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('De Do Do Do, De Da Da Da', '1980-09-03', '00:04:09', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Behind My Camel', '1980-09-03', '00:02:54', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Man in a Suitcase', '1980-09-03', '00:02:19', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Shadows in the Rain', '1980-09-03', '00:05:04', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('The Other Way of Stopping', '1980-09-03', '00:03:22', '11f3417e-0d40-4dd7-b0a8-44ab13c4163b');

insert into authors (id, author_name)
VALUES ('3f69e4b8-2163-435e-8c48-28a867454ea2', 'King Crimson');
-- 3f69e4b8-2163-435e-8c48-28a867454ea2

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('21st Century Schizoid Man', '1969-10-10', '00:07:20', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('I Talk to the Wind', '1969-10-10', '00:06:05', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Epitaph', '1969-10-10', '00:08:47', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('Moonchild', '1969-10-10', '00:12:11', '3f69e4b8-2163-435e-8c48-28a867454ea2');

INSERT INTO songs (song_name, song_date, song_duration, song_author_id)
VALUES ('The Court of the Crimson King', '1969-10-10', '00:09:22', '3f69e4b8-2163-435e-8c48-28a867454ea2');

UPDATE songs
SET song_author = (SELECT author_name
                   FROM authors
                   WHERE authors.id = songs.song_author_id);
