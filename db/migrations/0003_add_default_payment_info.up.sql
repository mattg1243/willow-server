ALTER TABLE user_contact_info
ALTER COLUMN paymentInfo TYPE JSONB
USING paymentInfo::JSONB;

ALTER TABLE user_contact_info
ALTER COLUMN paymentInfo SET DEFAULT '{}'::jsonb;

-- handle existing null values
UPDATE user_contact_info
SET paymentInfo = '{}'::JSONB
WHERE paymentInfo IS NULL;

ALTER TABLE user_contact_info
ALTER COLUMN paymentInfo SET NOT NULL;