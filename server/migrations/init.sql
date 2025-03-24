CREATE TYPE priority AS ENUM('high', 'medium', 'low');
CREATE TYPE status AS ENUM('ready for work', 'in progress', 'done');


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL, -- Можно иднекс на email повесть, чтобы памяти много не забирало не делал
    hash_password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE families (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    creator_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE
);

ALTER TABLE users ADD COLUMN family_id INT DEFAULT NULL;
ALTER TABLE users ADD FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE SET NULL;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    assignee_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    priority priority NOT NULL,
    status status NOT NULL DEFAULT 'ready for work',
    deleted_at TIMESTAMP NULL,
    family_id INT NOT NULL,
    creator_id INT NOT NULL,
    FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (assignee_id) REFERENCES users(id) ON DELETE SET NULL
);


CREATE TABLE family_invitations (
    id SERIAL PRIMARY KEY,
    family_id INT NOT NULL,
    invited_user_id INT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, accepted, declined // да мб лучше было бы сделать + таблицу, но это ту мач
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE CASCADE,
    FOREIGN KEY (invited_user_id) REFERENCES users(id) ON DELETE CASCADE
);

