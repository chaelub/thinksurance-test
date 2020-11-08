CREATE TABLE IF NOT EXISTS roles_permissions (
    role_id bigint NOT NULL,
    permission_id bigint NOT NULL,
    create_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE,
    update_at timestamp without time zone NOT NULL DEFAULT NOW()::DATE
);

ALTER TABLE roles_permissions
    ADD CONSTRAINT fk_role_id_to_roles
        FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE roles_permissions
    ADD CONSTRAINT fk_permission_id_to_permissions
        FOREIGN KEY (permission_id) REFERENCES permissions(id) ON UPDATE CASCADE ON DELETE RESTRICT;

CREATE UNIQUE INDEX idx_role_permission ON roles_permissions (role_id, permission_id);

INSERT INTO roles_permissions(role_id, permission_id) VALUES
(1,1),
(2,2);