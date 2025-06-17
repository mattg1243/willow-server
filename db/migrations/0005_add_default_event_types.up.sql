INSERT INTO event_types (id, user_id, source, title, charge, created_at, updated_at)
VALUES
('f6fa9eec-6510-43c3-988c-a21ece02d29b', null, 'default', 'Phone Call', true, NOW(), NOW()),
('3fb0f5c7-e2b2-4240-8b86-71bcc48938e2', null, 'default', 'Refund', true, NOW(), NOW()),
('d8b4bdb9-35d2-45a6-a90e-720a23cb78de', null, 'default', 'Payment', false, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;
