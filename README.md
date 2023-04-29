# Asteroids NeoWs API

API for the NASA Asteroids NeoWs service, which allows you to get the number of asteroids near Earth for a given date.

## How to use

The service has a REST architecture with the following methods:

### `GET`

Returns the number of asteroids for a given date. Different dates in the format `YYYY-MM-DD` can be passed as input:

```
/neo/count?dates=2020-01-20&dates=2020-03-20&dates=2020-05-24&dates=2020-08-22
```

### `POST`

Writes (or updates if the date already exists) data to the database.
An array of dates and the number of asteroids in JSON is received at the address `/neo/count`:

```
{
  "neo_counts": [
    {
      "date": "2020-01-20",
      "count": 12
    },
    {
      "date": "2020-02-26",
      "count": 9
    }
  ]
}
```
