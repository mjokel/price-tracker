-- insert new fuel price
BEGIN TRANSACTION;

INSERT INTO prices (timestamp, price, fuel, station_id)
VALUES (?, ?, ?, ?);

COMMIT;