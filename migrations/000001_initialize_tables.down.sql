ALTER TABLE "Expense" DROP CONSTRAINT "Expense_category_id_fkey";
ALTER TABLE "Expense" DROP CONSTRAINT "Expense_event_id_fkey";

ALTER TABLE "event_membership" DROP CONSTRAINT "fk_user_id";
ALTER TABLE "event_membership" DROP CONSTRAINT "fk_event_id";

-- Drop tables in reverse order

DROP TABLE "Category";
DROP TABLE "Expense";
DROP TABLE "event_membership";
DROP TABLE "event";
DROP TABLE "user";

-- Drop the RoleEnum type

DROP TYPE RoleEnum;
