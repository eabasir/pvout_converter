# PVOUT Converter

The project uses the data of [Global Sunlight Atlas](https://globalsolaratlas.info/download/world) to ingest the global photovoltaic power potential data into a Postgres database. The Goespatial information has the the `Point` type as in the PostGIS.

## Usage
The project needs the `.env` file including the needed data for the `internal/configs`.

```
go run cmd/pvout_converter/main.go
```
