CREATE TABLE sucursal
(
    id        VARCHAR(255) NOT NULL,
    nombre    VARCHAR(255) NOT NULL,
    direccion VARCHAR(255) NOT NULL,
    localidad VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;