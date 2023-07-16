
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


SET TIMEZONE="Europe/Moscow";


CREATE TABLE clients (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    ip_address inet
);


CREATE TABLE info (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    key BYTEA NOT NULL,
    value BYTEA NOT NULL,
    client_id UUID,
    read_only BOOL DEFAULT FALSE,
    CONSTRAINT f_k_client FOREIGN KEY (client_id) REFERENCES clients(id)
);

