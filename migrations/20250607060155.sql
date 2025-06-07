-- Create "organizations" table
CREATE TABLE "todo"."organizations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "todo"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "organization_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_organization" FOREIGN KEY ("organization_id") REFERENCES "todo"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "organization_users" table
CREATE TABLE "todo"."organization_users" (
  "organization_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  PRIMARY KEY ("organization_id", "user_id"),
  CONSTRAINT "fk_organization_users_organization" FOREIGN KEY ("organization_id") REFERENCES "todo"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_organization_users_user" FOREIGN KEY ("user_id") REFERENCES "todo"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "todos" table
CREATE TABLE "todo"."todos" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "due_date" timestamptz NULL,
  "title" character varying(255) NOT NULL,
  "description" text NULL,
  "status" character varying(50) NULL,
  "visibility" character varying(50) NOT NULL,
  "created_by_user_id" uuid NOT NULL,
  "organization_id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_todos_organization" FOREIGN KEY ("organization_id") REFERENCES "todo"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_todos" FOREIGN KEY ("created_by_user_id") REFERENCES "todo"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "todo_assignees" table
CREATE TABLE "todo"."todo_assignees" (
  "user_id" uuid NOT NULL,
  "todo_id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "created_by_user_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "todo_id"),
  CONSTRAINT "fk_todo_assignees_todo" FOREIGN KEY ("todo_id") REFERENCES "todo"."todos" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_todo_assignees_user" FOREIGN KEY ("user_id") REFERENCES "todo"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "roles" table
CREATE TABLE "todo"."roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(50) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_roles_name" UNIQUE ("name")
);
-- Create "user_roles" table
CREATE TABLE "todo"."user_roles" (
  "user_id" uuid NOT NULL,
  "role_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "role_id"),
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "todo"."roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "todo"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
