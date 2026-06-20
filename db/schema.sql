CREATE TABLE goose_db_version (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		version_id INTEGER NOT NULL,
		is_applied INTEGER NOT NULL,
		tstamp TIMESTAMP DEFAULT (datetime('now'))
	);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE inntekt (
  id INTEGER PRIMARY KEY,
  navn TEXT NOT NULL,
  beløp INTEGER NOT NULL,
  beskrivelse TEXT,
  dato DATE NOT NULL,
  gjentagende BOOLEAN NOT NULL DEFAULT TRUE,
  gjentagende_dato DATE,
  gjentagende_årlig BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE utgift (
    id INTEGER PRIMARY KEY,
    navn TEXT NOT NULL,
    beløp INTEGER NOT NULL,
    type TEXT,
    beskrivelse TEXT dato DATE NOT NULL,
    gjentagende BOOLEAN NOT NULL DEFAULT TRUE,
    gjentagende_dato DATE,
    gjentagende_årlig BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE "lån" (
  id INTEGER PRIMARY KEY,
  navn TEXT NOT NULL,
  beløp_start INTEGER NOT NULL,
  rente INTEGER NOT NULL,
  start_dato DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  nedbetaling_planlagt_avsluttet_dato DATE
);
