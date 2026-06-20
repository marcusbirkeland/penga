-- +goose Up
CREATE TABLE
  IF NOT EXISTS "lån" (
    id INTEGER PRIMARY KEY,
    navn TEXT NOT NULL,
    beløp_start INTEGER NOT NULL,
    rente INTEGER NOT NULL,
    startdato DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    nedbetaling_planlagt_avsluttet_dato DATE,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

-- +goose Down
DROP TABLE IF EXISTS "lån";