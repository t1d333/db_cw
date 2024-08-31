SELECT quantilesExactExclusive(0.5, 0.95, 0.99)(size)
FROM (SELECT toInt64(payload.size) AS size FROM logs)
