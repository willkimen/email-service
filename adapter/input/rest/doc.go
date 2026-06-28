/*
Package rest implements the HTTP/REST API endpoints and transport layer orchestration.

This package functions as a primary input adapter (Driving Adapter). It
acts as the entry point for synchronous HTTP external communication, receiving requests,
unmarshaling payloads, and driving the application core through its input ports.

# Architectural Role & Responsibilities

The rest package belongs strictly to the outermost infrastructure layer. Its structural
responsibilities are strictly isolated to:
  - Routing incoming HTTP requests to their respective handlers.
  - Parsing, deserializing (JSON), and sanitizing transport-specific Data Transfer Objects (DTOs).
  - Translating web transport details into strongly typed application messages.
  - Invoking business use cases via abstract input ports (driving interfaces).
  - Intercepting application-layer errors (e.g., apperrors.InvalidFieldError) and translating
    them into standardized HTTP status codes (such as 400 Bad Request, 422 Unprocessable Entity,
    or 500 Internal Server Error) and structured JSON responses.

By encapsulating HTTP mechanics here, the application core remains completely agnostic to
concepts like status codes, headers, cookies, and HTTP routers.
*/
package rest
