CREATE TABLE
  sqlite_sequence (name, seq);

CREATE TABLE
  eier (id INTEGER PRIMARY KEY, navn TEXT NOT NULL);

CREATE TABLE
  inntekt (
    id INTEGER PRIMARY KEY,
    navn TEXT NOT NULL,
    "beløp" INTEGER NOT NULL,
    beskrivelse TEXT,
    dato DATE NOT NULL,
    fast BOOLEAN NOT NULL DEFAULT TRUE,
    fast_dato DATE,
    "fast_årlig" BOOLEAN NOT NULL DEFAULT FALSE,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

CREATE TABLE
  utgift (
    id INTEGER PRIMARY KEY,
    navn TEXT NOT NULL,
    "beløp" INTEGER NOT NULL,
    dato DATE NOT NULL,
    type TEXT,
    beskrivelse TEXT,
    fast BOOLEAN NOT NULL DEFAULT TRUE,
    fast_dato DATE,
    "fast_årlig" BOOLEAN NOT NULL DEFAULT FALSE,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

CREATE TABLE
  IF NOT EXISTS "lån" (
    id INTEGER PRIMARY KEY,
    navn TEXT NOT NULL,
    "beløp_start" INTEGER NOT NULL,
    rente INTEGER NOT NULL,
    startdato DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    nedbetaling_planlagt_avsluttet_dato DATE,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

CREATE TABLE
  fond (
    id INTEGER PRIMARY KEY NOT NULL,
    navn TEXT NOT NULL,
    "beløp_start" INTEGER NOT NULL,
    startdato DATE NOT NULL,
    "forventet_årlig_avkastning_prosent" INTEGER NULL,
    pris_prosent INTEGER NULL,
    ISIN TEXT NULL,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

CREATE TABLE
  konto (
    id INTEGER PRIMARY KEY NOT NULL,
    navn TEXT NOT NULL,
    "beløp_start" INTEGER NOT NULL,
    startdato DATE NOT NULL,
    rente INTEGER NULL,
    eier_id INTEGER,
    FOREIGN KEY (eier_id) REFERENCES eier (id)
  );

CREATE TABLE
  fond_sparing (
    id INTEGER PRIMARY KEY,
    fond_id INTEGER NOT NULL,
    utgift_id INTEGER NOT NULL,
    FOREIGN KEY (fond_id) REFERENCES fond (id),
    FOREIGN KEY (utgift_id) REFERENCES utgift (id)
  );

CREATE TABLE
  konto_sparing (
    id INTEGER PRIMARY KEY,
    konto_id INTEGER NOT NULL,
    utgift_id INTEGER NOT NULL,
    FOREIGN KEY (konto_id) REFERENCES konto (id),
    FOREIGN KEY (utgift_id) REFERENCES utgift (id)
  );

CREATE TABLE
  IF NOT EXISTS "lån_nedbetaling" (
    id INTEGER PRIMARY KEY,
    "lån_id" INTEGER NOT NULL,
    utgift_id INTEGER NOT NULL,
    FOREIGN KEY ("lån_id") REFERENCES "lån" (id),
    FOREIGN KEY (utgift_id) REFERENCES utgift (id)
  );