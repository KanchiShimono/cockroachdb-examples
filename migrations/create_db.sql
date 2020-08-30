CREATE USER IF NOT EXISTS maxroach;
CREATE DATABASE review;
GRANT ALL ON DATABASE review TO maxroach;


CREATE TABLE person (
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE post (
    id UUID NOT NULL,
    person_id UUID NOT NULL,
    text VARCHAR NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (person_id)
    REFERENCES person (id)
);
