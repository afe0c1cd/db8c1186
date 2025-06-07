schema "public" {
  comment = "Todo application schema"
}

table "users" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("uuid_generate_v4()")
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  column "created_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "organization_id" {
    type = uuid
    null = false
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "users_organization_id_fkey" {
    columns = [column.organization_id]
    ref_columns = [table.organizations.column.id]
  }
}

table "organizations" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("uuid_generate_v4()")
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  primary_key {
    columns = [column.id]
  }
}

table "roles" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("uuid_generate_v4()")
  }
  column "name" {
    type = varchar(50)
    null = false
    unique = true
  }
  primary_key {
    columns = [column.id]
  }
}

table "user_roles" {
  schema = schema.public
  column "user_id" {
    type = uuid
  }
  column "role_id" {
    type = uuid
  }
  primary_key {
    columns = [column.user_id, column.role_id]
  }
  foreign_key {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
  }
  foreign_key {
    columns = [column.role_id]
    ref_columns = [table.roles.column.id]
  }
}

table "todos" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("uuid_generate_v4()")
  }
  column "due_date" {
    type = timestamp
    null = true
  }
  column "title" {
    type = varchar(255)
    null = false
  }
  column "description" {
    type = text
    null = true
  }
  column "status" {
    type = varchar(50)
    null = true
  }
  column "visibility" {
    type = varchar(50)
    null = false
  }
  column "created_by_user_id" {
    type = uuid
    null = false
  }
  column "organization_id" {
    type = uuid
    null = false
  }
  column "created_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key {
    columns = [column.created_by_user_id]
    ref_columns = [table.users.column.id]
  }
  foreign_key {
    columns = [column.organization_id]
    ref_columns = [table.organizations.column.id]
  }
}

table "todo_assignees" {
  schema = schema.public
  column "user_id" {
    type = uuid
  }
  column "todo_id" {
    type = uuid
  }
  column "created_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "created_by_user_id" {
    type = uuid
    null = false
  }
  primary_key {
    columns = [column.user_id, column.todo_id]
  }
  foreign_key {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
  }
  foreign_key {
    columns = [column.todo_id]
    ref_columns = [table.todos.column.id]
  }
  foreign_key {
    columns = [column.created_by_user_id]
    ref_columns = [table.users.column.id]
  }
} 
