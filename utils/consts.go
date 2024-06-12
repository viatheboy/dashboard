package utils

// MONGODB_URI is the MongoDB URI for the service
const MONGODB_URI = "mongodb+srv://viatheyboy:uvL3rwoDOrt42G8g@dashboards.ylq2uxl.mongodb.net/?retryWrites=true&w=majority&appName=dashboards"

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
const ENDPOINT_DASHBOARD = "dashboard"

// ENDPOINT_DASHBOARD is the endpoint for the dashboard
const PATH_DASHBOARD = DEFAULT_PATH + ENDPOINT_DASHBOARD + SLASH

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
