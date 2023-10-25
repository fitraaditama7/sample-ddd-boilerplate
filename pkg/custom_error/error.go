package custom_error

import "github.com/gofiber/fiber/v2"

var BadRequestError = &HTTPError{StatusCode: fiber.StatusBadRequest, Message: "Bad Request"}
var ForbiddenError = &HTTPError{StatusCode: fiber.StatusForbidden, Message: "Forbidden"}
var UnauthorizedError = &HTTPError{StatusCode: fiber.StatusUnauthorized, Message: "Unauthorized"}
var NotFoundError = &HTTPError{StatusCode: fiber.StatusNotFound, Message: "Not Found"}
var InternalServerError = &HTTPError{StatusCode: fiber.StatusInternalServerError, Message: "Internal Server Error"}

var InvalidIDError = &HTTPError{StatusCode: fiber.StatusBadRequest, Message: "Invalid id"}
