#!/bin/bash

sleep 3
psql -h 127.0.0.1 -p 5432 -U postgres -d sobo < ./c_areas.sql
psql -h 127.0.0.1 -p 5432 -U postgres -d sobo < ./d_areas.sql
psql -h 127.0.0.1 -p 5432 -U postgres -d sobo < ./v_areas.sql
