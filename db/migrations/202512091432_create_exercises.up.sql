CREATE TABLE IF NOT EXISTS exercises
(
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    display_name VARCHAR (50) NOT NULL,
    point INTEGER NOT NULL,
    base_unit VARCHAR (255) NOT NULL,
    training_intensity FLOAT NOT NULL
);

INSERT INTO exercises
    (name, display_name, point, base_unit, training_intensity)
VALUES
    ('push_up','腕立て伏せ',5, '回', 1.0),
    ('squat',  'スクワット', 5, '回', 1.0),
    ('sit_up', '腹筋',     5, '回', 1.0),
    ('plank',  'プランク',  10, '秒', 1.2);
