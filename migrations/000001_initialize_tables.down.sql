ALTER TABLE "expense" DROP CONSTRAINT "fk_event_id_expense";
ALTER TABLE "expense" DROP CONSTRAINT "fk_category_id_expense";

ALTER TABLE "event_membership" DROP CONSTRAINT "fk_user_id_event_membership";
ALTER TABLE "event_membership" DROP CONSTRAINT "fk_event_id_event_membership";

ALTER TABLE "sessions" DROP CONSTRAINT "fk_event_id_sessions";


-- Drop tables in reverse order

DROP TABLE "category";
DROP TABLE "expense";
DROP TABLE "sessions";
DROP TABLE "event_membership";
DROP TABLE "event";
DROP TABLE "user";

-- Drop the RoleEnum type

DROP TYPE RoleEnum;
