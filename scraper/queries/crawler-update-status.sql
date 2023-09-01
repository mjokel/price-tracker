-- update station is_open flag
BEGIN TRANSACTION;

UPDATE stations
SET is_open = ?
WHERE id = ?;

COMMIT;