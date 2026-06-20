-- +goose Up
CREATE TABLE IF NOT EXISTS utgift (
    id INTEGER PRIMARY KEY,
    beløp INTEGER NOT NULL,
    dato DATE NOT NULL,
    gjentagende BOOLEAN NOT NULL DEFAULT TRUE,
    gjentagendeDato DATE,
    type TEXT,
    beskrivelse TEXT
);

-- +goose Down
DROP TABLE utgift;