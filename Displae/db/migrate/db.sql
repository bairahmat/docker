create table IF NOT EXISTS shipper (id serial,name string unique,logo string not null default '',primary key(id));

create table IF NOT EXISTS bank (id serial,name string unique,bank_code string unique ,swift_code string unique,country string not null references country(id),index (country),primary key(id));

create table IF NOT EXISTS country (id serial,code string unique,name string unique,primary key(id));

create table IF NOT EXISTS payment_method (id serial,name string unique, primary key(id));

create table IF NOT EXISTS user_account (id serial,username string unique,password string,phone_number string unique,country int not null references country(id),email string unique,index(country),primary key (id));

create table IF NOT EXISTS user_login_activity (id serial,user_id int not null references user_account(id),login_time timestamptz,os string,ip_address string,index(user_id));

create table IF NOT EXISTS user_profile (id serial,user_id int not null references user_acccount(id),real_name string not null default '',bank_account_number string not null default '',index(user_id),primary key (id));

create table IF NOT EXISTS user_profile_shipper (user_profile_id int not null references user_profile(id),shipper_id int not null references shipper(id),index(user_profile_id),index(shipper_id));

create table IF NOT EXISTS shipment_address (id serial,user_profile_id int not null references user_profile(id),address string not null default '',address_notes string not null default '',index(user_profile),primary key(id));

create table IF NOT EXISTS cart(id serial,user_id int not null references user_account(id),index(user_id),primary key(id));

create table IF NOT EXISTS cart_item (id serial,cart_id int not null references cart(id),seller_id int not null references user_account(id),product_id int not null refrences product(id),qty int not null default 1,index(cart_id),index(seller_id),index(product_id),primary key(id));

create table IF NOT EXISTS cart_shipper (cart_id int not null references cart(id),shipper_id int not null references shipper(id),index(cart_id),index(shipper_id),);

create table IF NOT EXISTS transaction (id serial,user_id int not null references user_account(id),transaction_date timestamptz,last_transaction_state int,last_transaction_state_date timestamptz,primary key(id));

create table IF NOT EXISTS transaction_detail (id serial,transaction_id int not null references transaction(id),seller_id int not null references user_account(id),product_id int not null references product ,shipper_id int not null references shipper(id),qty int not null default 1,index(transaction_id),index(seller_id),index(product_id),index(shipper_id),primary key (id));

create table IF NOT EXISTS transaction_state (id serial, transaction_id int not null references transaction(id),transaction_state int,state_date timestamptz,index(transaction_id),primary key(serial));

create table IF NOT EXISTS payment (id serial,transaction_id int not null references transaction(id),method int not null references payment_method(id),total_amount decimal not null default 0,unique_number decimal,index(transaction_id),primary_key(id));

create table IF NOT EXISTS tagihan (id serial,transaction_id int not null references transaction(id),nomor_tagihan string unique,total_tagihan decimal,index(transaction_id),primary key (id));

create table IF NOT EXISTS product (id serial,user_id string not null references user_account(id),name string,price decimal,description string,url string,sellable bool not null default false,index(user_id),primary key (id));
