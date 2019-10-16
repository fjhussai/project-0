CREATE TABLE user_accounts (
    uniqname varchar primary key,
    userfirst varchar,
    userlast varchar,
    password varchar);

CREATE TABLE bank_accounts (
    acctnumber varchar primary key,
    uniqname varchar,
    acctbalance decimal,
    accttype varchar);

insert into bank_accounts values('5464','cardib',54,'checking');

insert into user_accounts values ('cardib', 'cardi', 'belcalis', 'bodakyellow');
insert into user_accounts values ('schrutefarms','dwight', 'schrute','password');
insert into user_accounts values ('BEASLEY', 'pamela', 'beasley','password');
insert into user_accounts values ('bigtuna', 'James','Halpert', 'password');
insert into user_accounts values ('sprinkles','angela','schrute','password');
insert into user_accounts values ('smartypants', 'oscar', 'martinez','password')

/*
CREATE TABLE employee_info (
    emp_number INTEGER,
    emp_first VARCHAR,
    emp_last VARCHAR,
    emp_password VARCHAR
);
*/