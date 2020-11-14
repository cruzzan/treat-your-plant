CREATE TABLE plants(
    id uuid NOT NULL CONSTRAINT events_pkey PRIMARY KEY,
    name text,
    plant_type_id int NOT NULL CONSTRAINT plant_type_id_fk REFERENCES plant_types,
    created timestamp with time zone not null,
    updated timestamp with time zone not null
);

CREATE TABLE users(
    id uuid NOT NULL PRIMARY KEY,
    username text NOT NULL,
    created timestamp with time zone not null,
    updated timestamp with time zone not null
);

CREATE TABLE user_plants(
    user_id uuid NOT NULL CONSTRAINT user_plants_user_id_fk REFERENCES user,
    plant_id uuid NOT NULL CONSTRAINT user_plants_plant_id_fk REFERENCES plants,
    created timestamp with time zone not null,
    UNIQUE (user_id, plant_id)
);

CREATE TABLE plant_types(
    id serial NOT NULL CONSTRAINT plant_types_pkey PRIMARY KEY,
    latin_names json,
    common_names json,
    created timestamp with time zone not null,
    updated timestamp with time zone not null
);

CREATE TABLE tips(
    id bigserial NOT NULL CONSTRAINT tip_pkey PRIMARY KEY,
    type text NOT NULL,
    description text,
    tip_manifest json,
    schedule text,
    created timestamp with time zone not null,
    updated timestamp with time zone not null
);

CREATE TABLE plant_type_tips(
    plant_type_id bigint NOT NULL CONSTRAINT plant_type_tips_plant_type_id_fk REFERENCES user,
    tip_id uuid NOT NULL CONSTRAINT plant_type_tips_tip_id_fk REFERENCES plants,
    created timestamp with time zone not null,
    UNIQUE (plant_type_id, tip_id)
);

CREATE TABLE plant_tips(
    plant_id uuid NOT NULL CONSTRAINT plant_tips_plant_id_fk REFERENCES user,
    tip_id uuid NOT NULL CONSTRAINT plant_tips_tip_id_fk REFERENCES plants,
    created timestamp with time zone not null,
    UNIQUE (plant_id, tip_id)
);
