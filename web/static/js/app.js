fetch("/api/telemetry")

.then(r=>r.json())

.then(data=>{

console.log(data);

// здесь строятся графики Chart.js

// и маршрут Leaflet

});
