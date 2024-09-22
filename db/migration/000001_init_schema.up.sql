CREATE TABLE "Branches"(
                           "Id" SERIAL PRIMARY KEY,
                           "Name" VARCHAR NOT NULL,
                           "MaxUsers" INT NOT NULL,
                           "CurrentUsers" INT NOT NULL DEFAULT 0,
                           "GroupIds" INT ARRAY NOT NULL DEFAULT array[]::int[]
);