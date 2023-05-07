CREATE TABLE "activity" (
  "ID" SERIAL PRIMARY KEY,
  "days" VARCHAR(255) NULL,
  "description" VARCHAR(255) NULL,
  "created_at" TIMESTAMP NULL DEFAULT (now()),
  "update_at" TIMESTAMP NULL DEFAULT (now())
);
