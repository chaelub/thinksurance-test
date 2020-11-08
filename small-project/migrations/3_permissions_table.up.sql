CREATE TABLE IF NOT EXISTS permissions (
    id serial NOT NULL PRIMARY KEY,
    source varchar(50) NOT NULL,
    mode varchar(1) NOT NULL,
    create_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE,
    update_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE
);

INSERT INTO permissions(source, mode) VALUES
('*','*'),
('main_page','r');