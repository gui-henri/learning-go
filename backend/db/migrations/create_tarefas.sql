CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  descricao TEXT NOT NULL,
  prazo TEXT,
  concluida BOOLEAN DEFAULT FALSE,
  criada_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS patient (
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
 resource_json JSON,
 data_source VARCHAR
);

CREATE TABLE IF NOT EXISTS HUMAN_NAME (
    id VARCHAR(255) PRIMARY KEY,
    use VARCHAR(255),
    display VARCHAR(255),
    family VARCHAR(255),
    given_1 VARCHAR(255),
    given_2 VARCHAR(255),
    prefix_1 VARCHAR(255),
    suffix_1 VARCHAR(255),
    period_start TIMESTAMP,
    period_end TIMESTAMP
);

CREATE TABLE IF NOT EXISTS identifier (
  id SERIAL PRIMARY KEY,
  system VARCHAR,
  value VARCHAR,
  use VARCHAR,
  type_display VARCHAR,
  type_code VARCHAR,
  assigner_display VARCHAR
);

CREATE TABLE IF NOT EXISTS ORGANIZATION (
    id VARCHAR PRIMARY KEY,
    version INT,
    last_updated TIMESTAMP,
    name VARCHAR,
    address_line_1 VARCHAR,
    address_line_2 VARCHAR,
    address_city VARCHAR,
    address_state VARCHAR,
    address_postal_code VARCHAR,
    builder_id VARCHAR,
    created_at TIMESTAMP,
    resource_json JSON,
    data_source VARCHAR
);
