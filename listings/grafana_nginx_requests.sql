SELECT toDateTime(payload.'@timestamp') as time,
       COUNT(*) as requests FROM logs
GROUP by time
ORDER BY time;
