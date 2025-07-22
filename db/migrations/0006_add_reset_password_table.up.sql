CREATE TABLE IF NOT EXISTS reset_password (
  id SERIAL PRIMARY KEY,
  user_id UUID UNIQUE NOT NULL ,
  reset_token VARCHAR(255) UNIQUE NOT NULL,
  requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- add cron job to delete expired rows
-- CREATE EXTENSION IF NOT EXISTS pg_cron;

-- SELECT cron.schedule('delete_expired_resets', '0 * * * *', $$
--   DELETE FROM reset_password
--   WHERE expires_at <= NOW();
-- $$);