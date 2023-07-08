-- get_schema_create
create table users (
  user_id text not null,
  email varchar(255) not null,
  password varchar(255) not null,
  first_name varchar(255) not null,
  last_name varchar(255) not null,
  active bool default true,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_users primary key (user_id),
  constraint email_unique UNIQUE (email)
);
create table roles (
  role_id text not null,
  name varchar(255) not null,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_roles primary key (role_id)
);
create table user_roles (
  user_id text not null,
  role_id text not null,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_user_roles primary key (user_id, role_id)
);

create table products (
  product_id text not null,
  sku varchar(255) not null,
  name varchar(255) not null,
  description text,
  regular_price integer default 0,
  discount_price integer default 0,
  quantity integer default 0,
  taxable bool default false,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_products primary key (product_id)
);

-- get_associations_create
alter table user_roles
add constraint fk_roles_user_roles foreign key (role_id) references roles (role_id);
alter table user_roles
add constraint fk_users_user_roles foreign key (user_id) references users (user_id);

