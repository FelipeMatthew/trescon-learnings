-- Checar preço positivo
ALTER TABLE IF EXISTS public.peca
    ADD CONSTRAINT preco_positivo CHECK (preco_peca>=0)
    NOT VALID;


--  Deixar nome da marca unico 
ALTER TABLE IF EXISTS public.marca_carros
    ADD CONSTRAINT nome_marca_unico UNIQUE (nome_marca);


-- Modelo de peça unico
ALTER TABLE IF EXISTS public.peca
    ADD CONSTRAINT modelo_peca UNIQUE (modelo_numero);