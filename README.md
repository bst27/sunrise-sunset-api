# About
An HTTP API to get sunrise and sunset times.

# Example
This example shows request and response. All request fields are mandatory.
```
http://localhost:8080/sunrise-sunset?utcOffset=2&latitude=51.963691&longitude=7.610362&date=2020-08-17
```
```
{
    "Date": "2020-08-17",
    "Latitude": 51.963691,
    "Longitude": 7.610362,
    "Sunrise": "06:17:31",
    "Sunset": "20:48:25",
    "UtcOffset": 2
}
```