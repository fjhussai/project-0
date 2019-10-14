CREATE TABLE user_accounts (
    uniqname varchar primary key,
    userfirst varchar,
    userlast varchar,
    password varchar,
    funds decimal);

insert into user_accounts values ('cardib', 'cardi', 'belcalis', 'bodakyellow', 54);


CREATE TABLE bank_accounts (
    acctnumber serial primary key,
    uniqname varchar REFERENCES user_accounts(uniqname),
    acctbalance decimal,
    acct_type varchar);

insert into bank_accounts values('cardib',54,'checking')


/*
CREATE TABLE employee_info (
    emp_number INTEGER
    emp_first VARCHAR
    emp_last VARCHAR
    emp_password VARCHAR
)

*/