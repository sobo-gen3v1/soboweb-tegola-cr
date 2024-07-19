import { CSSProperties, createRef, useEffect, useRef } from 'react'
import { Map } from 'maplibre-gl'

// import { Map, View } from 'ol'
// import { VectorTile as LayerVectorTile } from 'ol/layer'
// import { VectorTile as SrcVectorTile } from 'ol/source'
// import { MVT } from 'ol/format'
// import { createXYZ } from 'ol/tilegrid'

function App() {
  const ref = createRef<HTMLDivElement>()
  const mapRef = useRef<Map>()

  // using maplibre
  useEffect(() => {
    if (ref.current && !mapRef.current) {
      mapRef.current = new Map({
        container: ref.current,
        style: 'https://api.maptiler.com/maps/streets/style.json?key=' + import.meta.env.VITE_APP_OPEN_STREET_TOKEN,
        center: [139.65031060, 35.67619190],
        zoom: 8,
      })

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
            "http://127.0.0.1:8181/maps/sobo/v1/{z}/{x}/{y}.vector.pbf",
          ],
        })

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
        mapRef.current?.addLayer({
          id: "areas_point",
          source: "sobo",
          "source-layer": "areas_point",
          type: "circle",
          paint: {
            "circle-color": "#FF0000",
          }
        })
        mapRef.current?.addLayer({
          id: "areas_linestring",
          source: "sobo",
          "source-layer": "areas_linestring",
          type: "line",
          layout: {
            "line-cap": 'round',
            "line-join": 'round'
          },
          paint: {
            "line-color": "#FF0000",
          }
        })
      })

    }
  }, [])

  const mapStyle: CSSProperties = {
    width: "90%",
    height: "70%",
    position: "absolute",
    margin: "0 auto",
    color: "white"
  }

  return (
    <>
      <div ref={ref} style={mapStyle}></div>
    </>
  )
}

export default App
