# WeatherAPI - get weather data


## Overview

WeatherAPI is an HTTP Service that reports on Melbourne weather. This service will source its information from either of the below providers:

* weatherstack (primary):

      curl "h​ttp://api.weatherstack.com/current?access_key=YOUR_ACCESS_KEY&query=Melbourne​"

      Documentation: https://weatherstack.com/documentation

* OpenWeatherMap (failover):

      Curl "http://api.openweathermap.org/data/2.5/weather?q=melbourne,AU&appid=APP_ID"

      Documentation: https://openweathermap.org/current


## Development Environment

* Copy .env.sample to .env. 
* Enter WeatherStack_Access_Key and OpenWeatherMap_AppID.

### build

```docker build -t weatherapi .```

### run

```docker run -p 8080:8081 -d weatherapi```
