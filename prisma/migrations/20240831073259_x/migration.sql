-- AlterTable
ALTER TABLE "cards" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "deployment_logs" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "deployments" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "github_account_connections" ADD COLUMN     "github_installation_id" INTEGER,
ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "invoices" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "machine_plans" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "machines" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "networks" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "repo_connections" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "social_connections" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "transactions" ALTER COLUMN "deleted_at" SET DEFAULT null;

-- AlterTable
ALTER TABLE "users" ALTER COLUMN "deleted_at" SET DEFAULT null;
