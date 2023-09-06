-- Adminer 4.8.1 PostgreSQL 15.4 (Debian 15.4-1.pgdg120+1) dump

CREATE TABLE "server"."etc" (
    "user_id" integer NOT NULL,
    "key" text NOT NULL,
    "value" text NOT NULL,
    CONSTRAINT "etc_user_id" PRIMARY KEY ("user_id")
) WITH (oids = false);


CREATE TABLE "server"."profile" (
    "user_id" integer NOT NULL,
    "nickname" text NOT NULL,
    CONSTRAINT "profile_nickname" PRIMARY KEY ("nickname"),
    CONSTRAINT "profile_user_id" UNIQUE ("user_id")
) WITH (oids = false);


CREATE TABLE "server"."recovery" (
    "user_id" integer NOT NULL,
    "code" text NOT NULL,
    "expired" integer NOT NULL,
    CONSTRAINT "recovery_code" PRIMARY KEY ("code")
) WITH (oids = false);


CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "server"."user" (
    "id" integer DEFAULT nextval('user_id_seq') NOT NULL,
    "status" integer NOT NULL,
    "uuid" uuid NOT NULL,
    "created" integer NOT NULL,
    "updated" integer NOT NULL,
    CONSTRAINT "user_id" PRIMARY KEY ("id"),
    CONSTRAINT "user_uuid" UNIQUE ("uuid")
) WITH (oids = false);


CREATE SEQUENCE social_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "server"."social" (
    "id" integer DEFAULT nextval('social_id_seq') NOT NULL,
    "user_id" integer NOT NULL,
    "target_id" integer NOT NULL,
    "type" text NOT NULL,
    "vars" text NOT NULL,
    "created" integer NOT NULL,
    "updated" integer NOT NULL,
    CONSTRAINT "social_id" PRIMARY KEY ("id")
) WITH (oids = false);


-- 2023-08-22 14:23:40.34976+00