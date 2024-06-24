SET allow_experimental_object_type = 1;

CREATE TABLE events (
    timestamp DateTime64 DEFAULT current_timestamp(),
    payload JSON
) ENGINE = MergeTree() ORDER BY timestamp;
