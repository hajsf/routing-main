/**
 * @license
 * Copyright 2019 Google LLC. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
 function initMap() {
    const directionsService = new google.maps.DirectionsService();
    const directionsRenderer = new google.maps.DirectionsRenderer();
    const map = new google.maps.Map(document.getElementById("map"), {
      zoom: 6,
      center: { lat: 26.372326, lng: 50.041064 },
    });
  
    directionsRenderer.setMap(map);
    document.getElementById("submit").addEventListener("click", () => {
      calculateAndDisplayRoute(directionsService, directionsRenderer);
    });
  }
  
  function calculateAndDisplayRoute(directionsService, directionsRenderer) {
    const waypts = [];
    const checkboxArray = document.getElementById("waypoints");
  
    for (let i = 0; i < checkboxArray.length; i++) {
      if (checkboxArray.options[i].selected) {
        waypts.push({
          location: checkboxArray[i].value,
          stopover: true,
        });
      }
    }
  
    directionsService
      .route({
        origin: document.getElementById("start").value,
        destination: document.getElementById("end").value,
        waypoints: waypts,
        optimizeWaypoints: true,
        travelMode: google.maps.TravelMode.DRIVING,
        language: "ar"
       //TransitOptions: {
         //   arrivalTime: new Date("2022-08-31T23:50:21.817Z"),
        //},
      })
      .then((response) => {
        directionsRenderer.setDirections(response);
  
        const route = response.routes[0];
        const summaryPanel = document.getElementById("directions-panel");
  
        summaryPanel.innerHTML = "";
        arrivalTime = new Date("2022-08-31T23:50:21.817Z")
        // For each route, display summary information.
        for (let i = 0; i < route.legs.length; i++) {
          const routeSegment = i + 1;
            console.log(route.legs[i])
          summaryPanel.innerHTML +=
            "<b>مرحلة الطريق رقم: " + routeSegment + "</b><br>";
        //  summaryPanel.innerHTML += route.legs[i].start_address + " to ";
        //  summaryPanel.innerHTML += route.legs[i].end_address + "<br>";
          summaryPanel.innerHTML += "مدة الرحله: ";
          summaryPanel.innerHTML += route.legs[i].duration.text + "<br>";
          summaryPanel.innerHTML += "مسافة الرحله: ";
          summaryPanel.innerHTML += route.legs[i].distance.text + "<br><br>";
        }
      })
      .catch((e) => window.alert("Directions request failed due to " + e));
  }
  
  window.initMap = initMap;
  