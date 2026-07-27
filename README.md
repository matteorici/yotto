# UAV Telemetry Visualizer

A lightweight Go web application for visualizing recorded UAV telemetry.

The application loads local CSV flight logs and displays charts for speed, altitude, battery level, and GPS route.

## Features

- CSV import
- Speed chart
- Altitude chart
- Battery chart
- GPS map
- REST API

## Run

```bash
go run ./cmd
```

Open

```
http://localhost:8080
```

## Future Improvements

- Live telemetry
- Flight replay
- GeoJSON export
- GPX export
- Dark mode
