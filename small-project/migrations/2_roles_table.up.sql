CREATE TABLE IF NOT EXISTS roles (
    id serial NOT NULL PRIMARY KEY,
    name varchar(50) NOT NULL,
    create_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE,
    update_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE
);

INSERT INTO roles(name) VALUES
('admin_role'),
('user_role');