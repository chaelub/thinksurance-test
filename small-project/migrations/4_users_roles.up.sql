CREATE TABLE IF NOT EXISTS users_roles (
    user_id bigint NOT NULL UNIQUE,
    role_id bigint NOT NULL,
    create_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE,
    update_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE
);

ALTER TABLE users_roles
    ADD CONSTRAINT fk_user_id_to_users
        FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE users_roles
    ADD CONSTRAINT fk_country_id_to_roles
        FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE RESTRICT;

INSERT INTO users_roles(user_id, role_id) VALUES
(1,1),
(2,2);