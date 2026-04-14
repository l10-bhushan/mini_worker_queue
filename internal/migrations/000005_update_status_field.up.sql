CREATE TYPE job_status AS ENUM ('pending', 'processing', 'completed', 'failed');
ALTER TABLE jobs ALTER COLUMN status TYPE job_status USING status::job_status;