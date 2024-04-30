SELECT eventType
FROM events
GROUP BY eventType
HAVING COUNT(eventType) > 1000;


---------------------

SELECT eventID
FROM events
WHERE toDayOfMonth(eventTime) = 1;

---------------------

SELECT userID
FROM events
GROUP BY userID
HAVING COUNT(eventType) > 3;