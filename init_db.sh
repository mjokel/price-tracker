#!/bin/bash

# SQLite database name
DATABASE="data.db"

# check if sqlite3 is installed
if ! command -v sqlite3 &> /dev/null; then
    echo "sqlite3 could not be found. Please install it and try again."
    exit 1
fi

# initialize table and insert dummy values
sqlite3 $DATABASE <<EOF
CREATE TABLE IF NOT EXISTS stations (
    id          INTEGER PRIMARY KEY, 
    uuid        TEXT,
    name        TEXT,
    brand       TEXT,
    zip         TEXT,
    city        TEXT,
    address     TEXT,
    is_open     BOOLEAN NOT NULL CHECK (is_open IN (0, 1))
);

INSERT INTO stations (uuid, name, brand, zip, city, address, is_open) 
VALUES ("41c5eac2-7a17-46d8-90fe-d0a8c56db5fc", "BK-Tankstelle Isarparkhaus", "BK", 80469, "München", "Baader Str. 6", 1);


CREATE TABLE IF NOT EXISTS prices (
    id          INTEGER PRIMARY KEY, 
    timestamp   TIMESTAMP,
    price       FLOAT,
    fuel        TEXT,
    station_id  INTEGER,
    FOREIGN KEY (station_id) REFERENCES stations(id)
);

.quit
EOF

echo "Initialized tables 'stations' and 'prices'!"
