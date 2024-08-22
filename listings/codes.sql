SELECT payload.code as code, COUNT(payload.code) as count
FROM logs
GROUP BY payload.code;
