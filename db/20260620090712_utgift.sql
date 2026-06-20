-- +goose Up
CREATE TABLE
    IF NOT EXISTS utgift (
        id INTEGER PRIMARY KEY,
        navn TEXT NOT NULL,
        beløp INTEGER NOT NULL,
        dato DATE NOT NULL,
        type TEXT,
        beskrivelse TEXT dato DATE NOT NULL,
        fast BOOLEAN NOT NULL DEFAULT TRUE,
        fast_dato DATE,
        fast_årlig BOOLEAN NOT NULL DEFAULT FALSE,
        eier_id INTEGER,
        FOREIGN KEY (eier_id) REFERENCES eier (id)
    );

-- +goose Down
DROP TABLE utgift;