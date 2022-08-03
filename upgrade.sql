-- 2022-08-01
CREATE TABLE public.t_user (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,
	username varchar NULL,
	"password" varchar(256) NULL,
	CONSTRAINT t_user_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX t_user_tenant_username_idx ON public.t_user USING btree (tenant_id,username);

-- 2022-08-03
CREATE TABLE public.t_wxaccount (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,
	openid varchar(256) NULL,
	unionid varchar(256) NULL,
	CONSTRAINT t_wxaccount_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX t_user_tenant_openid_idx ON public.t_wxaccount USING btree (tenant_id,openid);
CREATE UNIQUE INDEX t_user_unionid_idx ON public.t_wxaccount (tenant_id,unionid);

CREATE TABLE public.t_user_wxaccount (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,
	user_id uuid,
	wxaccount_id uuid,	
	CONSTRAINT t_user_wxaccount_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX t_user_userid_idx ON public.t_user_wxaccount USING btree (user_id);
CREATE UNIQUE INDEX t_user_wxaccountid_idx ON public.t_user_wxaccount USING btree (wxaccount_id);

