-- get latest price per station for E5, E10 and Diesel (transposed!)
SELECT 
	id AS station_id,
	station_uuid,
	is_open,
	IFNULL(MAX(CASE WHEN fuel = 'E5' THEN price END), 0) AS E5,
	IFNULL(MAX(CASE WHEN fuel = 'E10' THEN price END), 0) AS E10,
	IFNULL(MAX(CASE WHEN fuel = 'Diesel' THEN price END), 0) AS Diesel
FROM (
	SELECT
		id,
		uuid AS station_uuid,
		is_open,
		brand,
		timestamp,
		price,
		fuel,
		station_id
	FROM 
		stations
	LEFT JOIN (
		SELECT max(prices.id) AS price_id, timestamp, price, fuel, station_id
		FROM prices
		GROUP BY station_id, fuel) 
	ON id = station_id
	)
GROUP BY id;

-- -- get latest price per station for E5, E10 and Diesel (transposed!)
-- SELECT 
-- 	station_id,
-- 	station_uuid,
-- 	is_open,
-- 	IFNULL(MAX(CASE WHEN fuel = 'E5' THEN price END), 0) AS E5,
-- 	IFNULL(MAX(CASE WHEN fuel = 'E10' THEN price END), 0) AS E10,
-- 	IFNULL(MAX(CASE WHEN fuel = 'Diesel' THEN price END), 0) AS Diesel
-- FROM (
-- 	SELECT 
--         max(prices.id) AS id, 
--         timestamp, 
--         price, 
--         fuel, 
--         station_id, 
--         stations.uuid AS station_uuid,
-- 		stations.is_open
-- 	FROM prices
-- 	LEFT JOIN stations ON station_id = stations.id
-- 	GROUP BY station_id, fuel)
-- GROUP BY station_id;