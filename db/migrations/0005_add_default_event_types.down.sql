DELETE FROM event_types
WHERE id = ANY(ARRAY[
  'f6fa9eec-6510-43c3-988c-a21ece02d29b'::uuid, 
  '3fb0f5c7-e2b2-4240-8b86-71bcc48938e2'::uuid,
  'd8b4bdb9-35d2-45a6-a90e-720a23cb78de'::uuid
  ]);
