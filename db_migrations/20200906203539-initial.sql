
-- +migrate Up
CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    CONSTRAINT user_email_unique UNIQUE (email)
);

CREATE TABLE "item" (
    id UUID PRIMARY KEY,
    user_id UUID,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT fk_item_user FOREIGN KEY (user_id) REFERENCES "user" (id)
);

INSERT INTO "user" (id,name,email) 
VALUES (
    '4468918f-cb92-426d-a0dc-b452b94a151c', 
    'Aron Aronson', 
    'aronaronson@example.com'
);

INSERT INTO "item" (id,name,user_id) 
VALUES (
    '80b3448e-63ee-11eb-ae93-0242ac130002', 
    'The Unstoppable Force', 
    '4468918f-cb92-426d-a0dc-b452b94a151c'
);

-- +migrate Down
DROP TABLE "item";
DROP TABLE "user";
