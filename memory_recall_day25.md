# Day 25 - Error Handling Middleware
- Middleware wraps every request to catch errors and panics.
- uuid.New() assigns a request ID per request.
- Logs include request ID + path.
- Returns uniform JSON errors via middleware.
- Centralized error handling = cleaner code + consistent responses. 