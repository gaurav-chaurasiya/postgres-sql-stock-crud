-- Adminer 4.8.1 PostgreSQL 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1) dump

DROP TABLE IF EXISTS "stocks";
DROP SEQUENCE IF EXISTS stocks_stockid_seq;
CREATE SEQUENCE stocks_stockid_seq INCREMENT 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;

CREATE TABLE "public"."stocks" (
    "stockid" smallint DEFAULT nextval('stocks_stockid_seq') NOT NULL,
    "name" character(100),
    "price" smallint,
    "company" character(100),
    CONSTRAINT "stocks_pkey" PRIMARY KEY ("stockid")
) WITH (oids = false);

INSERT INTO "stocks" ("stockid", "name", "price", "company") VALUES
(16,	'n1 name                                                                                             ',	213,	'c new dwew                                                                                          '),
(15,	'n1 name 123                                                                                         ',	2133,	'c2 new dwew                                                                                         ');

-- 2024-03-10 19:43:49.12502+05:30
