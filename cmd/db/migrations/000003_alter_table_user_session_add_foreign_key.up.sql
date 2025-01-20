ALTER TABLE user_session
ADD CONSTRAINT fk_user_session_user_id
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
