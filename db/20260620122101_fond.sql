-- +goose Up
CREATE TABLE
  IF NOT EXISTS fond (
    id INTEGER PRIMARY KEY NOT NULL,
    navn TEXT NOT NULL,
    beløp_start INTEGER NOT NULL,
    startdato DATE NOT NULL,
    forventet_årlig_avkastning_prosent INTEGER NULL,
    pris_prosent INTEGER NULL,
    ISIN TEXT NULL,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

-- +goose Down
DROP TABLE IF EXISTS fond;