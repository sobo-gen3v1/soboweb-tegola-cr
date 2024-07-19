CREATE INDEX geom_index_display_areas ON d_areas USING GIST (wkb_geometry);
