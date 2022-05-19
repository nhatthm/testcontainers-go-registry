CREATE TABLE customer (
      id          INT NOT NULL PRIMARY KEY IDENTITY(1,1),
      country     VARCHAR(2) NOT NULL,
      first_name  VARCHAR(50) NOT NULL,
      last_name   VARCHAR(50) NOT NULL,
      email       VARCHAR(200) NOT NULL,
      created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at  DATETIME
);

CREATE INDEX customer_country ON customer(country);
CREATE UNIQUE INDEX customer_email ON customer(email);

INSERT INTO customer VALUES ('US', 'John', 'Doe', 'john.doe@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
