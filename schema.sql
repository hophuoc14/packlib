create table users (
    id uuid primary key,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    username varchar(100) constraint users_username_key unique,
    password text not null,
    email varchar(50) constraint users_email_key unique,
    profile varchar(255)
);

create index idx_users_deleted_at on users(deleted_at);

create table departments (
    id uuid primary key,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    name varchar(255) constraint departments_name_key unique,
    description text
);

create index idx_departments_name on departments(name);
create index idx_departments_deleted_at on departments(deleted_at);

create table employees (
    id uuid primary key,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    full_name varchar(100),
    phone varchar(15) constraint employees_phone_key unique,
    age smallint,
    status varchar(50) not null, 
    user_id uuid 
        constraint employees_user_id_fkey 
            references users(id) 
                on update cascade on delete set null
);

create index idx_employees_deleted_at on employees(deleted_at);

create table employee_departments (
    employee_id uuid 
        constraint employee_departments_employee_id_fkey
            references employees(id)
                on update cascade on delete cascade,
    department_id uuid
        constraint employee_departments_department_id_fkey
            references departments(id)
                on update cascade on delete cascade,
    primary key (employee_id, department_id)
);

create table projects (
    id uuid primary key,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    name varchar(255) not null constraint projects_name_key unique,
    description text,
    department_id uuid 
        constraint projects_department_id_fkey 
            references departments(id) 
                on update cascade on delete set null
);

create index idx_projects_deleted_at on projects(deleted_at);

create table teams (
    id uuid primary key,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    name varchar(255) constraint teams_name_key unique,
    project_id uuid 
        constraint teams_project_id_fkey 
            references projects(id) 
                on update cascade on delete set null,
    employee_id uuid 
        constraint teams_employee_id_fkey 
            references employees(id) 
                on update cascade on delete set null
);

create index idx_teams_deleted_at on teams(deleted_at);
create index idx_teams_name on teams(name);

create table employee_teams (
    employee_id uuid 
        constraint employee_teams_employee_id_fkey
            references employees(id)
                on update cascade on delete cascade,
    team_id uuid
        constraint employee_teams_team_id_fkey
            references teams(id)
                on update cascade on delete cascade,
    primary key (employee_id, team_id)
);

