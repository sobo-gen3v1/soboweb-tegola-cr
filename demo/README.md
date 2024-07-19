## App.tsx

```tsx
mapRef.current.on("load", () => {
    mapRef.current?.addSource("sobo", {
        type: "vector",
        tiles: [
        // /maps/:map_name/:layer_name/:version/:z/:x/:y

        // Return vector tiles for a map. The URI supports the following variables:

        // - `:map_name` is the name of the map as defined in the `config.toml` file.
        // - `:layer_name` is the name of the provider name as defined in the `config.toml` file.
        // - `:version` is the version of api.
        // - `:z` is the zoom level of the map.
        // - `:x` is the row of the tile at the zoom level.
        // - `:y` is the column of the tile at the zoom level.
        "http://127.0.0.1:8181/maps/sobo/v1/{z}/{x}/{y}.vector.pbf",  // <-- modify the url here
        ],
    })
    // configure the rendering style here
    mapRef.current?.addLayer({
        id: "areas_polygon",
        // config.toml
        // [[maps]]
        // name = "bonn"
        source: "sobo",
        // config.toml
        // [[providers.layers]]
        // name = "main_roads"
        "source-layer": "areas_polygon",
        type: "fill",
        paint: {
            "fill-color": "#ff0000",
            "fill-opacity": 0.8
        }
    })
})
```