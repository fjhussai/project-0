CREATE TABLE user_accounts (
    uniqname varchar primary key,
    userfirst varchar,
    userlast varchar,
    password varchar,
    funds decimal);

CREATE TABLE bank_accounts (
    acctnumber varchar primary key,
    uniqname varchar,
    acctbalance decimal,
    accttype varchar);

insert into bank_accounts values('5464','cardib',54,'checking');

insert into user_accounts values ('cardib', 'cardi', 'belcalis', 'bodakyellow', 54);
insert into user_accounts values ('schrutefarms','dwight', 'schrute','password', 100);
insert into user_accounts values ('BEASLEY', 'pamela', 'beasley','password', 84);
insert into user_accounts values ('bigtuna', 'James','Halpert', 'password', 84);
insert into user_accounts values ('sprinkles','angela','schrute','password', 5);
insert into user_accounts values ('smartypants', 'oscar', 'martinez','password',100)

/*
CREATE TABLE employee_info (
    emp_number INTEGER,
    emp_first VARCHAR,
    emp_last VARCHAR,
    emp_password VARCHAR
);
*/