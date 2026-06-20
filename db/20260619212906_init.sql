-- +goose Up
CREATE TABLE IF NOT EXISTS inntekt (
  id INTEGER PRIMARY KEY,
  beløp INTEGER NOT NULL,
  gjentagende BOOLEAN NOT NULL DEFAULT TRUE,
  gjentagendeDato DATE,
  dato DATE NOT NULL,
  beskrivelse TEXT
);

-- +goose Down
DROP TABLE inntekt;