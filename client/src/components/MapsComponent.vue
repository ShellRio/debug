<template>
  <div class="map-wrap h-full w-full">
    <div class="map" ref="mapContainer"></div>
  </div>
</template>

<script>
import maplibregl from "maplibre-gl";
import { getAll } from "@/services/users.js";

export default {
  mounted() {
    this.initializeMap();
  },
  methods: {
    initializeMap() {
      maplibregl.accessToken = "YOUR_MAPLIBRE_GL_ACCESS_TOKEN";

      const map = new maplibregl.Map({
        container: this.$refs.mapContainer,
        style:
          "https://api.maptiler.com/maps/5bf070bd-db6a-47ab-99b1-715b7773481f/style.json?key=s2pMtJD3WQhyX32yxJLp",
        center: [7.88818, 51.18898],
        zoom: 5,
        interactive: true,
      });

      // Appel à votre service pour obtenir les données des victimes
      getAll()
        .then((data) => {
          // Loop à travers les données et ajouter les marqueurs avec les événements de clic
          data.forEach((item) => {
            let markerColor = "red";
            if (item.status === "on-Line") {
              markerColor = "green"; // Changer la couleur du marqueur si le statut est "on-Line"
            }

            const marker = new maplibregl.Marker({
              color: markerColor, // Utiliser la couleur définie en fonction du statut
            })
              .setLngLat([item.longitude, item.latitude])
              .addTo(map);

            // Créer la bulle d'informations
            const popup = new maplibregl.Popup({
              closeButton: true,
              closeOnClick: false,
            }).setHTML(`<h3>${item.id}</h3>
                        <p>Status: ${item.status}</p>
                        <p>IP: ${item.ip}</p>
                        <p>Pays: ${item.pays}</p>
                        <p>Ville: ${item.ville}</p>`);

            // Ajouter l'événement de clic pour afficher la bulle d'informations
            marker.getElement().addEventListener("click", () => {
              popup.addTo(map);
            });
          });
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>

<style scoped>
@import "~maplibre-gl/dist/maplibre-gl.css";

.map-wrap {
  position: relative;
  width: 100%;
  height: calc(
    100vh - 77px
  ); /* calculate height of the screen minus the heading */
}

.map {
  position: absolute;
  width: 100%;
  height: 100%;
}
</style>
