CREATE TABLE branches(
                           "id" SERIAL PRIMARY KEY,
                           "name" VARCHAR NOT NULL,
                           "maxUsers" INT NOT NULL,
                           "currentUsers" INT NOT NULL DEFAULT 0,
                           "groupIds" INT ARRAY NOT NULL DEFAULT array[]::int[]
);