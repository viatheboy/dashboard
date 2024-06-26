package utils

// DEFAULT_PORT is the default port for the service
const DEFAULT_PORT = "8080"

// DEFAULT_PATH is the root path.
const DEFAULT_PATH = "/dashboards/" + VERSION + SLASH

// VERSION is the current version of the service.
const VERSION = "v1"

// LOCALHOST is the localhost address
const LOCALHOST = "localhost"

// ROOT is the root path
const SLASH = "/"

// ENDPOINT_REGISTRATIONS holds the name for the registration endpoint
const ENDPOINT_REGISTRATIONS = "registrations"

// PATH_REGISTRATIONS is the path to the registration endpoint
const PATH_REGISTRATIONS = DEFAULT_PATH + ENDPOINT_REGISTRATIONS + SLASH

// PATH_AUTHENTICATION is the path to the authentication endpoint
const ENDPOINT_AUTHENTICATION = "auth"

// ENDPOINT_AUTHENTICATION is the endpoint for authentication
const PATH_AUTHENTICATION = DEFAULT_PATH + ENDPOINT_AUTHENTICATION + SLASH

// PATH_DASHBOARD is the path to the dashboard
const ENDPOINT_WEATHER = "weather"

// ENDPOINT_DASHBOARD is the endpoint for the dashboard
const PATH_WEATHER = DEFAULT_PATH + ENDPOINT_WEATHER + SLASH

// COLLECTION_USERS is the MongoDB collection for users
const COLLECTION_USERS = "users"

// ENDPOINT_SET_COOKIE is the endpoint for setting a cookie
const ENDPOINT_SET_COOKIE = "set-cookie"

// PATH_SET_COOKIE is the path to the set-cookie endpoint
const PATH_SET_COOKIE = DEFAULT_PATH + ENDPOINT_SET_COOKIE + SLASH

// ENDPOINT_GET_COOKIE is the endpoint for getting a cookie
const ENDPOINT_GET_COOKIE = "get-cookie"

// PATH_GET_COOKIE is the path to the get-cookie endpoint
const PATH_GET_COOKIE = DEFAULT_PATH + ENDPOINT_GET_COOKIE + SLASH

// ENDPOINT_DELETE_COOKIE is the endpoint for deleting a cookie
const ENDPOINT_DELETE_COOKIE = "delete-cookie"

// PATH_DELETE_COOKIE is the path to the delete-cookie endpoint
const PATH_DELETE_COOKIE = DEFAULT_PATH + ENDPOINT_DELETE_COOKIE + SLASH

// GEOCODING_API is the URL for the geo-coding API
const GEOCODING_API = "https://geocoding-api.open-meteo.com/v1/search/?name="

// GEOLOCATIONS_TEST_FILE is the local path to the test file for geolocations
const GEOLOCATIONS_TEST_FILE = "../test_data/geolocations.json"

// WEATHER_TEST_FILE is the local path to the test file for weather
const WEATHER_TEST_FILE = "../test_data/weather.json"

// WEATHER_API_BASE is the base URL for the weather API
const WEATHER_API_BASE = "https://api.open-meteo.com/v1/forecast?"

// WEATHER_API_LAT is the latitude parameter for the weather API
const WEATHER_API_LAT = "latitude="

// WEATHER_API_LON is the longitude parameter for the weather API
const WEATHER_API_LON = "&longitude="

// WEATHER_API_HOURLY is the hourly parameters for the weather API
const WEATHER_API_PARAMETERS = "&hourly=temperature_2m,precipitation,cloud_cover,wind_speed_10m&wind_speed_unit=ms"

// WEEKLY_HOURS is the total number of hours in a week
const WEEKLY_HOURS = 168

// CONDITION_RAINY is the condition for rainy weather returned by the internal weather determiner
const CONDITION_RAINY = "rainy"

// CONDITION_CLOUDY is the condition for cloudy weather returned by the internal weather determiner
const CONDITION_CLOUDY = "cloudy"

// CONDITION_MOSTLY_CLOUDY is the condition for mostly cloudy weather returned by the internal weather determiner
const CONDITION_MOSTLY_CLOUDY = "mostly cloudy"

// CONDITION_PARTLY_CLOUDY is the condition for partly cloudy weather returned by the internal weather determiner
const CONDITION_PARTLY_CLOUDY = "partly cloudy"

// CONDITION_MOSTLY_SUNNY is the condition for mostly sunny weather returned by the internal weather determiner
const CONDITION_MOSTLY_SUNNY = "mostly sunny"

// CONDITION_MOSTLY_CLEAR is the condition for mostly clear weather returned by the internal weather determiner
const CONDITION_MOSTLY_CLEAR = "mostly clear"

// CONDITION_CLEAR_DAY is the condition for clear day weather returned by the internal weather determiner
const CONDITION_CLEAR_DAY = "clear day"

// TIME_FORMAT is the format for the time string returned in the weather handler
const TIME_FORMAT = "2006-01-02T15:04"
