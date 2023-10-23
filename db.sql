-- Adminer 4.8.1 PostgreSQL 15.4 (Debian 15.4-1.pgdg120+1) dump

CREATE TABLE "account" (
    "id" uuid NOT NULL,
    "token" text NOT NULL,
    CONSTRAINT "account_id" UNIQUE ("id")
) WITH (oids = false);

-- 2023-08-22 14:23:40.34976+00