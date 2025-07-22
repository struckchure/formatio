-- CreateEnum
CREATE TYPE "currency" AS ENUM ('NGN');

-- CreateEnum
CREATE TYPE "deployment_status" AS ENUM ('IDLE', 'IN_PROGRESS', 'SUCCESSFUL', 'FAILED');

-- CreateEnum
CREATE TYPE "invoice_status" AS ENUM ('PAID', 'UNPAID');

-- CreateEnum
CREATE TYPE "machine_status" AS ENUM ('CREATING', 'RUNNING', 'REBOOTING', 'SHUTTING_DOWN', 'SHUTDOWN');

-- CreateEnum
CREATE TYPE "transaction_status" AS ENUM ('PENDING', 'SUCCESS', 'FAILED');

-- CreateEnum
CREATE TYPE "transaction_type" AS ENUM ('DEBIT', 'CREDIT');

-- CreateTable
CREATE TABLE "users" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "first_name" TEXT,
    "last_name" TEXT,
    "email" TEXT,
    "password" TEXT,
    "roles" TEXT[],

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "cards" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "user_id" TEXT NOT NULL,
    "is_default" BOOLEAN NOT NULL DEFAULT false,
    "card_type" TEXT,
    "last_digits" TEXT,
    "auth_token" TEXT,
    "is_approved" BOOLEAN NOT NULL DEFAULT false,
    "expiry_month" TEXT,
    "expiry_year" TEXT,

    CONSTRAINT "cards_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "deployment_logs" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "deployment_id" TEXT NOT NULL,
    "job_id" TEXT,
    "message" TEXT,

    CONSTRAINT "deployment_logs_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "deployments" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "machine_id" TEXT NOT NULL,
    "repo_connection_id" TEXT NOT NULL,
    "commit_hash" TEXT,
    "commit_message" TEXT,
    "actor" TEXT,
    "status" "deployment_status" NOT NULL DEFAULT 'IDLE',

    CONSTRAINT "deployments_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "github_account_connections" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "user_id" TEXT NOT NULL,
    "github_id" TEXT,
    "github_username" TEXT,
    "github_email" TEXT,

    CONSTRAINT "github_account_connections_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "invoices" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "user_id" TEXT NOT NULL,
    "product_id" TEXT,
    "reference" TEXT,
    "description" TEXT,
    "from" TIMESTAMPTZ(6),
    "to" TIMESTAMPTZ(6),
    "quantity" INTEGER,
    "unit_price" MONEY NOT NULL,
    "total_price" MONEY NOT NULL,
    "currency" "currency" NOT NULL DEFAULT 'NGN',
    "status" "invoice_status" NOT NULL DEFAULT 'UNPAID',

    CONSTRAINT "invoices_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "machine_plans" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "name" TEXT,
    "currency" "currency" NOT NULL DEFAULT 'NGN',
    "monthly_rate" MONEY,
    "hourly_rate" MONEY,
    "cpu" TEXT NOT NULL,
    "memory" TEXT NOT NULL,

    CONSTRAINT "machine_plans_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "machines" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "owner_id" TEXT NOT NULL,
    "plan_id" TEXT NOT NULL,
    "container_id" TEXT,
    "machine_name" TEXT,
    "machine_image" TEXT,
    "machine_status" "machine_status" NOT NULL DEFAULT 'CREATING',

    CONSTRAINT "machines_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "networks" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "machine_id" TEXT NOT NULL,
    "host_name" TEXT,
    "protocol" TEXT,
    "listening_port" INTEGER NOT NULL,
    "destination_port" INTEGER NOT NULL,
    "service_id" TEXT,
    "ingress_id" TEXT,

    CONSTRAINT "networks_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "repo_connections" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "machine_id" TEXT NOT NULL,
    "repo_id" TEXT,
    "repoName" TEXT,

    CONSTRAINT "repo_connections_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "social_connections" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "user_id" TEXT NOT NULL,
    "connection_id" TEXT,
    "connection_type" TEXT,

    CONSTRAINT "social_connections_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "transactions" (
    "id" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT null,
    "user_id" TEXT NOT NULL,
    "amount" MONEY NOT NULL,
    "currency" "currency" NOT NULL DEFAULT 'NGN',
    "status" "transaction_status" NOT NULL,
    "type" "transaction_type" NOT NULL,
    "reference" TEXT NOT NULL,
    "description" TEXT,

    CONSTRAINT "transactions_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

-- CreateIndex
CREATE UNIQUE INDEX "machine_plans_name_key" ON "machine_plans"("name");

-- CreateIndex
CREATE UNIQUE INDEX "machines_owner_id_machine_name_key" ON "machines"("owner_id", "machine_name");

-- CreateIndex
CREATE UNIQUE INDEX "transactions_reference_key" ON "transactions"("reference");

-- AddForeignKey
ALTER TABLE "cards" ADD CONSTRAINT "cards_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "deployment_logs" ADD CONSTRAINT "deployment_logs_deployment_id_fkey" FOREIGN KEY ("deployment_id") REFERENCES "deployments"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "deployments" ADD CONSTRAINT "deployments_machine_id_fkey" FOREIGN KEY ("machine_id") REFERENCES "machines"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "deployments" ADD CONSTRAINT "deployments_repo_connection_id_fkey" FOREIGN KEY ("repo_connection_id") REFERENCES "repo_connections"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "github_account_connections" ADD CONSTRAINT "github_account_connections_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "invoices" ADD CONSTRAINT "invoices_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE NO ACTION;

-- AddForeignKey
ALTER TABLE "machines" ADD CONSTRAINT "machines_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "machines" ADD CONSTRAINT "machines_plan_id_fkey" FOREIGN KEY ("plan_id") REFERENCES "machine_plans"("id") ON DELETE NO ACTION ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "networks" ADD CONSTRAINT "networks_machine_id_fkey" FOREIGN KEY ("machine_id") REFERENCES "machines"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "repo_connections" ADD CONSTRAINT "repo_connections_machine_id_fkey" FOREIGN KEY ("machine_id") REFERENCES "machines"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "social_connections" ADD CONSTRAINT "social_connections_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "transactions" ADD CONSTRAINT "transactions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;
