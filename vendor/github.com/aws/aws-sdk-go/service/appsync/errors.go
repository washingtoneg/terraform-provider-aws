// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have access to perform this operation on this resource.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeApiKeyLimitExceededException for service response error code
	// "ApiKeyLimitExceededException".
	//
	// The API key exceeded a limit. Try your request again.
	ErrCodeApiKeyLimitExceededException = "ApiKeyLimitExceededException"

	// ErrCodeApiKeyValidityOutOfBoundsException for service response error code
	// "ApiKeyValidityOutOfBoundsException".
	//
	// The API key expiration must be set to a value between 1 and 365 days from
	// creation (for CreateApiKey) or from update (for UpdateApiKey).
	ErrCodeApiKeyValidityOutOfBoundsException = "ApiKeyValidityOutOfBoundsException"

	// ErrCodeApiLimitExceededException for service response error code
	// "ApiLimitExceededException".
	//
	// The GraphQL API exceeded a limit. Try your request again.
	ErrCodeApiLimitExceededException = "ApiLimitExceededException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The request is not well formed. For example, a value is invalid or a required
	// field is missing. Check the field values, and then try again.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Another modification is in progress at this time and it must complete before
	// you can make your change.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeGraphQLSchemaException for service response error code
	// "GraphQLSchemaException".
	//
	// The GraphQL schema is not valid.
	ErrCodeGraphQLSchemaException = "GraphQLSchemaException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// An internal AWS AppSync error occurred. Try your request again.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request exceeded a limit. Try your request again.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The resource specified in the request was not found. Check the resource,
	// and then try again.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// You are not authorized to perform this operation.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":              newErrorAccessDeniedException,
	"ApiKeyLimitExceededException":       newErrorApiKeyLimitExceededException,
	"ApiKeyValidityOutOfBoundsException": newErrorApiKeyValidityOutOfBoundsException,
	"ApiLimitExceededException":          newErrorApiLimitExceededException,
	"BadRequestException":                newErrorBadRequestException,
	"ConcurrentModificationException":    newErrorConcurrentModificationException,
	"GraphQLSchemaException":             newErrorGraphQLSchemaException,
	"InternalFailureException":           newErrorInternalFailureException,
	"LimitExceededException":             newErrorLimitExceededException,
	"NotFoundException":                  newErrorNotFoundException,
	"UnauthorizedException":              newErrorUnauthorizedException,
}
