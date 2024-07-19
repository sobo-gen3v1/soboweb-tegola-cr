## For Reference

```sql
DROP TYPE IF EXISTS public.geometric_type;
CREATE TYPE public.geometric_type AS ENUM ('Point','Polygon','Linestring','MultiPoint','MultiLineString','MultiPolygon','GeometryCollection');

CREATE TABLE IF NOT EXISTS public.m_areas (
    ogc_fid bigserial NOT NULL,
    wkb_geometry public.geometry NOT NULL,
    geometric_type public.geometric_type NOT NULL,
    area_code character varying,
    properties text,
    other_tags public.hstore,
    created_at timestamptz default CURRENT_TIMESTAMP,
    updated_at timestamptz default CURRENT_TIMESTAMP,
    deleted_at timestamptz
);

ALTER TABLE public.m_areas OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.d_areas (
    ogc_fid bigserial NOT NULL,
    wkb_geometry public.geometry NOT NULL,
    geometric_type public.geometric_type NOT NULL,
    gis_data_id varchar,
    area_code varchar,
    properties text,
    other_tags public.hstore,
    created_at timestamptz default CURRENT_TIMESTAMP,
    updated_at timestamptz default CURRENT_TIMESTAMP,
    deleted_at timestamptz
);

ALTER TABLE public.d_areas OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.c_areas (
    geom_id bigserial NOT NULL,
    ogc_fid bigint NOT NULL,
    properties text,
    other_tags public.hstore,
    created_at timestamptz default CURRENT_TIMESTAMP,
    updated_at timestamptz default CURRENT_TIMESTAMP,
    deleted_at timestamptz
);

ALTER TABLE public.c_areas OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.v_areas (
    "version" VARCHAR NOT NULL PRIMARY KEY,
     layer_name VARCHAR NOT NULL,
     start_time TIMESTAMPTZ NOT NULL,
     end_time TIMESTAMPTZ NOT NULL,
     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NULL,
     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NULL,
     deleted_at TIMESTAMPTZ NULL
);

ALTER TABLE public.v_areas OWNER TO postgres;

```