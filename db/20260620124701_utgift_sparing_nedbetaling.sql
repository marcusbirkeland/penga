-- +goose Up
CREATE TABLE
  IF NOT EXISTS fond_sparing (
    id INTEGER PRIMARY KEY,
    fond_id INTEGER NOT NULL,
    utgift_id INTEGER NOT NULL,
    FOREIGN KEY (fond_id) REFERENCES fond (id),
    FOREIGN KEY (utgift_id) REFERENCES utgift (id)
  );

CREATE TABLE
  IF NOT EXISTS konto_sparing (
    id INTEGER PRIMARY KEY,
    konto_id INTEGER NOT NULL,
    utgift_id INTEGER NOT NULL,
    FOREIGN KEY (konto_id) REFERENCES konto (id),
    FOREIGN KEY (utgift_id) REFERENCES utgift (id)
  );

CREATE TABLE
  IF NOT EXISTS lån_nedbetaling (
    id INTEGER PRIMARY KEY,
    lån_id INTEGER NOT NULL,
    utgift_id INTEGER NOT NULL,
    FOREIGN KEY (lån_id) REFERENCES "lån" (id),
    FOREIGN KEY (utgift_id) REFERENCES utgift (id)
  );

-- +goose Down
DROP TABLE IF EXISTS lån_nedbetaling;

DROP TABLE IF EXISTS konto_sparing;

DROP TABLE IF EXISTS fond_sparing;