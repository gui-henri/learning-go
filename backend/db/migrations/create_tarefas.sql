CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  descricao TEXT NOT NULL,
  prazo TEXT,
  concluida BOOLEAN DEFAULT FALSE,
  criada_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pacientes (
 id SERIAL PRIMARY KEY,
 version INT DEFAULT 0,
 last_updated TIMESTAMP,
 active BOOLEAN DEFAULT TRUE,
 gender VARCHAR,
 birth_date DATE,
 ethnicity VARCHAR,
 race VARCHAR,
 deceased BOOLEAN,
 deceased_datetime TIMESTAMP,
 marital_status VARCHAR,
 marital_status_display VARCHAR,
 managing_organization_id VARCHAR,
 name_family VARCHAR,
 name_given_1 VARCHAR,
 name_given_2 VARCHAR,
 builder_id VARCHAR,
 upid VARCHAR,
 created_at VARCHAR,



)