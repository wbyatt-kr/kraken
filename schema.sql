CREATE TABLE services (
  id BIGSERIAL PRIMARY KEY,
  backend text NOT NULL
);

CREATE TABLE routes (
  id BIGSERIAL PRIMARY KEY,
  path text NOT NULL,
  service_id BIGSERIAL NOT NULL
);