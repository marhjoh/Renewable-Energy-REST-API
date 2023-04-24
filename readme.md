# Assignment 2 #

> Group members: Kjetil Indrehus, Martin Johannessen, Sander Hauge.

This is an API which allows for: searching of reports on percentage of renewable energy in different countries' energy mix over time.
```
/energy/v1/renewables/current 
/energy/v1/renewables/history/
/energy/v1/notifications/ 
/energy/v1/status/
```

## Current endpoint ##

This endpoint retrieves the elements of the latest year currently available. The newest data in renewable-share-energy
is from 2021, and is therefore the current year of this project. Features of this endpoint:
* Search for country by name and country code.
* Add-on to get neighbouring countries
* Cache for reducing amount of calls to countries API.

It uses the file renewable-share-energy.csv and REST Countries API, which is retrieved from: http://129.241.150.113:8080/v3.1. 

**Using the endpoint**
```
REQUEST: GET
PATH: /energy/v1/renewables/current/{country?}{?neighbours=bool?}
```

Using no extra parameters will print all countries to the client.
If an optional parameter: /{country?}, is passed the corresponding country will be printed. This variable could be both
country codes, and also country name.
The query: {?neighbours=bool?}, may also be used, and will print information about the neighbouring countries of the
country passed. This query is dependent on the optional parameter country.

### Current Test ###

There is created a test class for the current endpoint.

To use the test, print into command line when in root project folder:
> go test .\internal\webserver\handlers\current_test.go


## History endpoint ##

This endpoint retrieves all elements from renewable-share-energy. When no query is passed it will return the mean of all
data based on each country. Functionality of history endpoint:
* Search for specific countries based on country code and name. 
* Allows for searching for specific years.
* Allows for searching between specific years.
* Sort by percentage and alphabetically, both descending and ascending. 
* Calculating the mean of a country.

It uses the file renewable-share-energy.csv and API for countries.

**Using the endpoint**
```
REQUEST: GET
PATH: /energy/v1/renewables/history/{country?}{?begin=year?}{?end=year?}{?mean=bool?}{?sortbyvalue=bool?}
```
These can also be combined, using "&" after "?". Begin and end query combined will find countries between the ones written.

**Example request:**
```
REQUEST: GET
PATH: /energy/v1/renewables/history/nor?begin=2011&end=2014&sortbyvalue=true
```

```json
[
{
"name": "Norway",
"isoCode": "NOR",
"year": 2012,
"percentage": 70.095116
},
{
"name": "Norway",
"isoCode": "NOR",
"year": 2014,
"percentage": 68.88728
},
{
"name": "Norway",
"isoCode": "NOR",
"year": 2013,
"percentage": 67.50864
},
{
"name": "Norway",
"isoCode": "NOR",
"year": 2011,
"percentage": 66.30012
}
]
```

### History test ###

There is created a test class for the history endpoint.

To use the test, print into command line when in root project folder:
> go test .\internal\webserver\handlers\history_test.go