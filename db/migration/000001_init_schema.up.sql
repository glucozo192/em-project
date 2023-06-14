-- get_schema_create
create table users (
  id text not null,
  email varchar(255) not null,
  password varchar(255) not null,
  first_name varchar(255) not null,
  last_name varchar(255) not null,
  active bool default true,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_users primary key (id),
  constraint email_unique UNIQUE (email)
);
create table roles (
  id text not null,
  name varchar(255) not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_roles primary key (id)
);
create table user_roles (
  user_id text not null,
  role_id text not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_user_roles primary key (user_id, role_id)
);
create table categories (
  id text not null,
  name varchar(255) not null,
  parent_id text,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_categories primary key (id)
);
create table products (
  id text not null,
  sku varchar(255) not null,
  name varchar(255) not null,
  description text,
  product_status_id text not null,
  regular_price numeric default 0,
  discount_price numeric default 0,
  quantity integer default 0,
  taxable bool default false,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_products primary key (id)
);
create table tags (
  id text not null,
  name varchar(255) not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_tags primary key (id)
);
create table sales_orders (
  id text not null,
  order_date date not null,
  total numeric not null,
  coupon_id text,
  session_id varchar(255) not null,
  user_id text not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_sales_orders primary key (id)
);
create table coupons (
  id text not null,
  code varchar(255) not null,
  description text,
  active bool default true,
  value numeric,
  multiple bool default false,
  start_date timestamp with time zone,
  end_date timestamp with time zone,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_coupons primary key (id)
);
create table product_tags (
  product_id text not null,
  tag_id text not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_product_tags primary key (product_id, tag_id)
);
create table cc_transactions (
  id text not null,
  code varchar(255),
  order_id text not null,
  transdate timestamp with time zone,
  processor varchar(255) not null,
  processor_trans_id varchar(255) not null,
  amount numeric not null,
  cc_num varchar(255),
  cc_type varchar(255),
  response text,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_cc_transactions primary key (id)
);
create table sessions (
  id varchar(255) not null,
  data text,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_sessions primary key (id)
);
create table product_statuses (
  id text not null,
  name varchar(255) not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_product_statuses primary key (id)
);
create table product_categories (
  category_id text not null,
  product_id text not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_product_categories primary key (category_id, product_id)
);
create table order_products (
  id text not null,
  order_id text,
  sku varchar(255) not null,
  name varchar(255) not null,
  description text,
  price numeric not null,
  quantity integer not null,
  subtotal numeric not null,
  inserted_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  constraint pk_order_products primary key (id)
);
-- get_associations_create
alter table sales_orders
add constraint fk_coupon_order foreign key (coupon_id) references coupons (id);
alter table product_tags
add constraint fk_products_product_tags foreign key (product_id) references products (id);
alter table product_tags
add constraint fk_tags_product_tags foreign key (tag_id) references tags (id);
alter table user_roles
add constraint fk_roles_user_roles foreign key (role_id) references roles (id);
alter table user_roles
add constraint fk_users_user_roles foreign key (user_id) references users (id);
alter table product_categories
add constraint fk_category_products_categories foreign key (category_id) references categories (id);
alter table sales_orders
add constraint fk_user_sales_order foreign key (user_id) references users (id);
alter table sales_orders
add constraint fk_session_sales_order foreign key (session_id) references sessions (id);
alter table products
add constraint fk_product_statuses_product foreign key (product_status_id) references product_statuses (id);
alter table order_products
add constraint fk_sales_orders_order_products foreign key (order_id) references sales_orders (id);
alter table cc_transactions
add constraint fk_sales_order_cc_transaction foreign key (order_id) references sales_orders (id);
alter table product_categories
add constraint fk_product_product_category foreign key (product_id) references products (id);
alter table categories
add constraint fk_category_parent_category foreign key (parent_id) references categories (id);