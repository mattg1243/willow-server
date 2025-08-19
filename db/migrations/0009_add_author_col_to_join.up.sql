ALTER TABLE events_payment_types
ADD COLUMN author UUID,
ADD CONSTRAINT fk_author FOREIGN KEY (author) REFERENCES users(id);