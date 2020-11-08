CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL PRIMARY KEY,
    login varchar(50) NOT NULL,
    password varchar(255) NOT NULL,
    create_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE,
    update_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE
);

CREATE UNIQUE INDEX idx_login ON users (login);

INSERT INTO users(login, password) VALUES
('admin', '$2a$10$aZ9Uy/5sBvGtf8IBlTVCgur/VzJ4SREtXGt5jf0EENCyDv9YSoRw6'),
('user', '$2a$10$6seLs.6w0HMQqGLm3fNRKueLYGmHV8u5JHTow7Du18/FaHivB/mV.');