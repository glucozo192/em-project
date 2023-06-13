-- get_constraints_drop 
alter table sales_orders drop constraint fk_coupon_order ;
alter table product_tags drop constraint fk_products_product_tags ;
alter table product_tags drop constraint fk_tags_product_tags ;
alter table user_roles drop constraint fk_roles_user_roles ;
alter table user_roles drop constraint fk_users_user_roles ;
alter table product_categories drop constraint fk_category_products_categories ;
alter table sales_orders drop constraint fk_user_sales_order ;
alter table sales_orders drop constraint fk_session_sales_order ;
alter table products drop constraint fk_product_statuses_product ;
alter table order_products drop constraint fk_sales_orders_order_products ;
alter table cc_transactions drop constraint fk_sales_order_cc_transaction ;
alter table product_categories drop constraint fk_product_product_category ;
alter table categories drop constraint fk_category_parent_category ;

-- get_schema_drop
drop table users;
drop table roles;
drop table user_roles;
drop table categories;
drop table products;
drop table tags;
drop table sales_orders;
drop table coupons;
drop table product_tags;
drop table cc_transactions;
drop table sessions;
drop table product_statuses;
drop table product_categories;
drop table order_products;