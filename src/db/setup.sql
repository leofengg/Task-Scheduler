CREATE EXTENSION IF NOT EXISTS "uuid-ossp"

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    command TEXT NOT NULL,
    scheduled_at TIMESTAMP NOT NULL,
    status TEXT CHECK (status IN ('PENDING', 'COMPLETED', 'IN_PROGRESS')),
    completed_at TIMESTAMP,
    started_at TIMESTAMP,
    locked_until TIMESTAMP,
)

CREATE INDEX idx_tasks_scheduled_at ON tasks (scheduled_at);