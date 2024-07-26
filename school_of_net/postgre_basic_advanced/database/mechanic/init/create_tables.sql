-- Tabela Cliente

CREATE TABLE public.cliente
(
    id_cliente serial NOT NULL,
    nome character varying(50) NOT NULL,
    sobrenome character varying(50) NOT NULL,
    contato text NOT NULL,
    observacao text,
    data_cadastro date,
    PRIMARY KEY (id_cliente)
);

ALTER TABLE IF EXISTS public.cliente
    OWNER to postgres;


-- Tabela marca dos carros
CREATE TABLE public.marca_carros
(
    id_marca serial NOT NULL,
    nome_marca character varying(50) NOT NULL,
    PRIMARY KEY (id_marca)
);

ALTER TABLE IF EXISTS public.marca_carros
    OWNER to postgres;


-- Tabela tipo conserto

CREATE TABLE public.tipo_conserto
(
    id_tipo_conserto serial NOT NULL,
    nome_conserto character varying(100) NOT NULL,
    PRIMARY KEY (id_tipo_conserto)
);

ALTER TABLE IF EXISTS public.tipo_conserto
    OWNER to postgres;


-- Tabela peça
CREATE TABLE public.peca
(
    id_peca serial NOT NULL,
    nomde_peca character varying(100) NOT NULL,
    modelo_numero character varying NOT NULL,
    preco_peca numeric(10, 2) NOT NULL,
    PRIMARY KEY (id_peca)
);

ALTER TABLE IF EXISTS public.peca
    OWNER to postgres;