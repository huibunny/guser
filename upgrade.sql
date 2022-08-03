-- 2022-08-01
CREATE TABLE public.t_user (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,
	username varchar NULL,
	"password" varchar NULL
);
CREATE UNIQUE INDEX t_user_tenant_id_idx ON public.t_user (tenant_id,username);
