CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS patient (
    internal_id BIGSERIAL PRIMARY KEY,
    id VARCHAR UNIQUE NOT NULL,
    last_updated TIMESTAMPTZ,
    active BOOLEAN DEFAULT TRUE,
    gender VARCHAR(20), -- PROMOTED FIELD
    birth_date DATE, -- PROMOTED FIELD
    deceased BOOLEAN,
    full_name VARCHAR(255), -- PROMOTED FIELD
    cpf VARCHAR(11) UNIQUE, -- PROMOTED FIELD
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    resource_json JSONB NOT NULL
);

CREATE INDEX idx_patient_id ON patient (id);

CREATE INDEX idx_patient_cpf ON patient (cpf);

CREATE INDEX idx_patient_birth_date ON patient (birth_date);

CREATE INDEX idx_patient_full_name ON patient USING GIN (
    to_tsvector('portuguese', full_name)
);

-- Índice GIN para buscas flexíveis dentro do JSONB (para campos não promovidos).
CREATE INDEX idx_patient_recurso_gin ON patient USING GIN (resource_json);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    descricao TEXT NOT NULL,
    prazo TEXT,
    concluida BOOLEAN DEFAULT FALSE,
    patient_id VARCHAR,
    criada_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_patient Foreign Key (patient_id) REFERENCES patient (id) ON DELETE CASCADE
);

CREATE TABLE avaliation (
    id SERIAL PRIMARY KEY,
    resource_json JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    active BOOLEAN DEFAULT TRUE
);