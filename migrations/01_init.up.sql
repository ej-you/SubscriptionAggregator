DROP TABLE IF EXISTS subs;

CREATE TABLE subs (
    id SERIAL PRIMARY KEY,
    service VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NULL
);