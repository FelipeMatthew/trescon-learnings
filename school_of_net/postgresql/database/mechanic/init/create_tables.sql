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
    nome_peca character varying(100) NOT NULL,
    modelo_numero character varying NOT NULL,
    preco_peca numeric(10, 2) NOT NULL,
    PRIMARY KEY (id_peca)
);

ALTER TABLE IF EXISTS public.peca
    OWNER to postgres;

ALTER TABLE IF EXISTS public.peca
    ADD CONSTRAINT preco_positivo CHECK (preco_peca>=0)
    NOT VALID;


-- Tabela carro serie
CREATE TABLE public.carro_series
(
    id_serie serial NOT NULL,
    nome_serie character varying(50) NOT NULL,
    id_marca integer NOT NULL,
    PRIMARY KEY (id_serie),
    CONSTRAINT referencia_id_marca FOREIGN KEY (id_marca)
        REFERENCES public.marca_carros (id_marca) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public.carro_series
    OWNER to postgres;


-- Tabela veículo

CREATE TABLE public.veiculo
(
    id_veiculo serial NOT NULL,
    id_cliente integer NOT NULL,
    id_serie integer NOT NULL,
    ano smallint NOT NULL,
    numero_serie character varying(100) NOT NULL,
    placa character varying(100) NOT NULL,
    nota text,
    PRIMARY KEY (id_veiculo),
    CONSTRAINT referencia_serie_carro FOREIGN KEY (id_serie)
        REFERENCES public.carro_series (id_serie) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT referencia_id_cliente FOREIGN KEY (id_cliente)
        REFERENCES public.cliente (id_cliente) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public.veiculo
    OWNER to postgres;


-- Tabela conserto
CREATE TABLE public.consertos
(
    id_conserto serial NOT NULL,
    descricao text NOT NULL,
    chegada timestamp with time zone NOT NULL,
    inicio_hora_conserto timestamp without time zone NOT NULL,
    fim_hora_conserto timestamp without time zone NOT NULL,
    despensa timestamp without time zone NOT NULL,
    custo_servico numeric(10, 2) NOT NULL,
    id_veiculo integer NOT NULL,
    id_tipo_conserto integer NOT NULL,
    PRIMARY KEY (id_conserto),
    CONSTRAINT referencia_id_veiculo FOREIGN KEY (id_veiculo)
        REFERENCES public.veiculo (id_veiculo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT referencia_id_tipo_conserto FOREIGN KEY (id_tipo_conserto)
        REFERENCES public.tipo_conserto (id_tipo_conserto) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public.consertos
    OWNER to postgres;