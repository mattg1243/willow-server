CREATE TABLE IF NOT EXISTS payment_types (
  id SERIAL PRIMARY KEY,
  user_id UUID,
  "name" VARCHAR (255) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- default types
INSERT INTO payment_types (user_id, "name")
VALUES
(null, 'Check'), (null, 'Venmo'), (null, 
'PayPal'), (null, 'Credit card');

-- create join table
CREATE TABLE IF NOT EXISTS events_payment_types (
  event_id UUID NOT NULL,
  payment_type_id SERIAL NOT NULL,
  FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
  FOREIGN KEY (payment_type_id) REFERENCES payment_types(id) ON DELETE CASCADE
);