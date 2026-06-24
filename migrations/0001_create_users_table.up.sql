CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NULL DEFAULT NULL,
    email_verified_at TIMESTAMP NULL DEFAULT NULL,
    password VARCHAR(255) NOT NULL,
    remember_token VARCHAR(100) NULL DEFAULT NULL,
    avatar VARCHAR(500) NULL DEFAULT NULL,
    status ENUM(
        'pending',
        'active',
        'inactive',
        'blocked'
    ) NOT NULL DEFAULT 'pending',
    last_login_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by BIGINT UNSIGNED NULL DEFAULT NULL,
    updated_by BIGINT UNSIGNED NULL DEFAULT NULL,
    deleted_by BIGINT UNSIGNED NULL DEFAULT NULL,

    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_phone_number (phone_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
