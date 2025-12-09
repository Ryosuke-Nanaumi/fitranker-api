CREATE TABLE IF NOT EXISTS training_records
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    exercise_id INTEGER NOT NULL REFERENCES exercises (id) ON DELETE CASCADE,
    date DATE NOT NULL,
    amount INTEGER NOT NULL
);

INSERT INTO training_records
    (user_id, exercise_id, date, amount)
VALUES
    (1, 1, NOW()::date, 20),
    (1, 2, NOW()::date, 30),
    (2, 1, NOW()::date, 10);
