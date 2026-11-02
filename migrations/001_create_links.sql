CREATE TABLE IF NOT EXISTS links (
    id BIGSERIAL PRIMARY KEY, 
    code VARCHAR(10) UNIQUE NOT NULL,
    original text not null,
    clicks bigint not null default 0,
    created_at timestamptz not null default now()
)