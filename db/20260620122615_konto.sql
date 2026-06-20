-- +goose Up
CREATE TABLE
  IF NOT EXISTS konto (
    id INTEGER PRIMARY KEY NOT NULL,
    navn TEXT NOT NULL,
    beløp_start INTEGER NOT NULL,
    startdato DATE NOT NULL,
    rente INTEGER NULL,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

-- +goose Down
DROP TABLE IF EXISTS konto;