## Preparing

1. Start up the docker containers (redis and postgresql and migration)
    - `cd docker`
    - `docker compose up -d`
2. Import test data in the test-data directory.
    - `d_areas.sql` is the spectee data.
    - `c_areas.sql` is the polygon data from `m_areas` table.
    - `v_areas.sql` is the version data, `version=v1/v2/v3` was enable.

## Build & Run Project

1. Run `build.ps1`/`build.sh` script.
2. Run `run.ps1`/`run.sh` script.

## Access From URL

1. Without layer name: `http://127.0.0.1:8181/maps/sobo/v1/{z}/{x}/{y}.vector.pbf`
2. With layer name: `http://127.0.0.1:8181/maps/sobo/areas_polygon/v1/{z}/{x}/{y}.vector.pbf`
