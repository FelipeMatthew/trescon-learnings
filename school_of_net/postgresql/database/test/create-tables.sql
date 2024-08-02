CREATE TABLE public.empregados
(
    id_empregado serial NOT NULL,
    CONSTRAINT id_empregados_pk PRIMARY KEY (id_empregado)
);

ALTER TABLE IF EXISTS public.empregados
    OWNER to postgres;

-- Inserindo colunas
ALTER TABLE IF EXISTS public.empregados
    ADD COLUMN nome character varying(50) NOT NULL;

ALTER TABLE IF EXISTS public.empregados
    ADD COLUMN sobrenome character varying(50) NOT NULL;



-- Departamento

CREATE TABLE public.departamento
(
    id_departamento serial NOT NULL,
    departamento character varying(50) NOT NULL,
    PRIMARY KEY (id_departamento)
);

ALTER TABLE IF EXISTS public.departamento
    OWNER to postgres;