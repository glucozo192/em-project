-- get_constraints_drop 
alter table user_roles drop constraint fk_roles_user_roles ;
alter table user_roles drop constraint fk_users_user_roles ;

-- get_schema_drop
drop table users;
drop table roles;
drop table user_roles;