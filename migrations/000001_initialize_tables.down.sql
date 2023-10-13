ALTER TABLE "Expense" DROP CONSTRAINT "Expense_category_id_fkey";
ALTER TABLE "Expense" DROP CONSTRAINT "Expense_event_id_fkey";

ALTER TABLE "event_membership" DROP CONSTRAINT "event_membership_event_id_fkey";
ALTER TABLE "event_membership" DROP CONSTRAINT "event_membership_user_id_fkey";

-- Drop tables in reverse order

DROP TABLE "Category";
DROP TABLE "Expense";
DROP TABLE "event_membership";
DROP TABLE "event";
DROP TABLE "user";

-- Drop the RoleEnum type

DROP TYPE RoleEnum;
