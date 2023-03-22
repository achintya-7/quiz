CREATE TABLE "question" (
  "id" bigserial PRIMARY KEY,
  "quiz_id" bigserial NOT NULL,
  "question" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "answer" varchar NOT NULL
);

CREATE TABLE "quiz" (
  "id" bigserial PRIMARY KEY,
  "created_by" varchar NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "question" ADD FOREIGN KEY ("quiz_id") REFERENCES "quiz" ("id");
