-- update station is_open flag
UPDATE stations
SET is_open = ?
WHERE id = ?;