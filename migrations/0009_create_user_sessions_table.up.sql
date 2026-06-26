CREATE TABLE user_sessions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    user_id BIGINT UNSIGNED NOT NULL,

    refresh_token_hash VARCHAR(255) NOT NULL,

    device_name VARCHAR(255),

    user_agent VARCHAR(500),

    ip_address VARCHAR(45),

    expired_at DATETIME NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,

    created_by BIGINT UNSIGNED NULL,
    updated_by BIGINT UNSIGNED NULL,
    deleted_by BIGINT UNSIGNED NULL,

    CONSTRAINT fk_user_sessions_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    UNIQUE KEY uk_refresh_token_hash (refresh_token_hash),

    INDEX idx_user_id (user_id),
    INDEX idx_expired_at (expired_at),
    INDEX idx_deleted_at (deleted_at)
);