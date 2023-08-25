-- update timestamp only
UPDATE prices
SET timestamp = ?
WHERE id = (
    SELECT id
    FROM prices
	WHERE fuel = ? AND station_id = ?
    ORDER BY id DESC
    LIMIT 1
);

-- UPDATE prices
-- SET timestamp = CURRENT_TIMESTAMP
-- WHERE id = (
--     SELECT id
--     FROM prices
-- 	WHERE fuel = "Diesel" AND station_id = 3
--     ORDER BY id DESC
--     LIMIT 1
-- );