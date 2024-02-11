CREATE TABLE cliente_saldo (
    id INT AUTO_INCREMENT PRIMARY KEY,
    saldo INT NOT NULL,
    limite INT NOT NULL
);

CREATE TABLE transacoes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    valor INT NOT NULL,
    cliente_id INT NOT NULL,
    tipo VARCHAR(1) NOT NULL,
    descricao VARCHAR(10) NOT NULL,
    realizada_em TIMESTAMP NOT NULL
);

ALTER TABLE transacoes ADD INDEX ix_cliente_id (cliente_id);

INSERT INTO cliente_saldo (saldo, limite)
VALUES
    (0, 1000 * 100),
    (0, 800 * 100),
    (0, 10000 * 100),
    (0, 100000 * 100),
    (0, 5000 * 100);