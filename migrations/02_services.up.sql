DROP TABLE IF EXISTS services;

CREATE TABLE services (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

DROP TABLE IF EXISTS subs;

CREATE TABLE subs (
    id UUID PRIMARY KEY,
    service_id UUID NOT NULL,
    price INT NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NULL
);

ALTER TABLE subs
ADD CONSTRAINT fk_subs_service FOREIGN KEY (service_id) REFERENCES services (id) ON DELETE CASCADE ON UPDATE CASCADE;