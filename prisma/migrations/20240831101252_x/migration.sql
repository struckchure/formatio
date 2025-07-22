/*
  Warnings:

  - You are about to drop the column `github_installation_id` on the `repo_connections` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "cards" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "deployment_logs" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "deployments" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "github_account_connections" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "invoices" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "machine_plans" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "machines" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "networks" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "repo_connections" DROP COLUMN "github_installation_id",
ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "social_connections" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "transactions" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "users" ALTER COLUMN "deleted_at" SET DEFAULT null;
