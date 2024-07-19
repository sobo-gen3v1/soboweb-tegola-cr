#!/bin/bash

sleep 3
psql -h postgis -p 5432 -U postgres -d sobo < ./data/map-schema.sql
psql -h postgis -p 5432 -U postgres -d sobo -f ./data/sql/add_index.sql
psql -h postgis -p 5432 -U postgres -d sobo < ./data/map-data.sql

rm ./data/map-schema.sql
rm ./data/map-data.sql