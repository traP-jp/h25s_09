CREATE TABLE messages (
    id          CHAR(36)    PRIMARY KEY,
    author      VARCHAR(32) NOT NULL,
    message     TEXT        NOT NULL,
    replies_to  CHAR(36)    DEFAULT NULL,
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_author (author),
    INDEX idx_replies_to (replies_to),
    INDEX idx_created_at (created_at)
);

CREATE TABLE message_images (
    id CHAR(36) PRIMARY KEY,
    message_id  CHAR(36)    NOT NULL,
    data        MEDIUMBLOB  NOT NULL,
    mime        VARCHAR(64) NOT NULL,
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (message_id) REFERENCES messages(id),
    INDEX idx_message_id (message_id)
);

CREATE TABLE message_reactions (
    message_id  CHAR(36)    NOT NULL,
    username    VARCHAR(32) NOT NULL,
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (message_id) REFERENCES messages(id),
    INDEX idx_message_id (message_id)
);

CREATE TABLE achievements (
    id          INT         NOT NULL,
    username    VARCHAR(32) NOT NULL,
    achieved_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_username (username),
    INDEX idx_username_achieved_at (username, achieved_at)
);
