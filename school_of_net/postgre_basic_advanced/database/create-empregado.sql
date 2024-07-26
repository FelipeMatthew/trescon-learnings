CREATE TABLE public.empregados
(
    id_empregado integer NOT NULL,
    CONSTRAINT id_empregados_pk PRIMARY KEY (id_empregado)
);

ALTER TABLE IF EXISTS public.empregados
    OWNER to postgres;

-- Inserindo colunas
ALTER TABLE IF EXISTS public.empregados
    ADD COLUMN nome character varying(50) NOT NULL;

ALTER TABLE IF EXISTS public.empregados
    ADD COLUMN sobrenome character varying(50) NOT NULL;