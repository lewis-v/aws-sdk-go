// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lambdaiface provides an interface for the AWS Lambda.
package lambdaiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaAPI is the interface type for lambda.Lambda.
type LambdaAPI interface {
	AddPermissionRequest(*lambda.AddPermissionInput) (*service.Request, *lambda.AddPermissionOutput)

	AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)

	CreateEventSourceMappingRequest(*lambda.CreateEventSourceMappingInput) (*service.Request, *lambda.EventSourceMappingConfiguration)

	CreateEventSourceMapping(*lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	CreateFunctionRequest(*lambda.CreateFunctionInput) (*service.Request, *lambda.FunctionConfiguration)

	CreateFunction(*lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)

	DeleteEventSourceMappingRequest(*lambda.DeleteEventSourceMappingInput) (*service.Request, *lambda.EventSourceMappingConfiguration)

	DeleteEventSourceMapping(*lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	DeleteFunctionRequest(*lambda.DeleteFunctionInput) (*service.Request, *lambda.DeleteFunctionOutput)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)

	GetEventSourceMappingRequest(*lambda.GetEventSourceMappingInput) (*service.Request, *lambda.EventSourceMappingConfiguration)

	GetEventSourceMapping(*lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	GetFunctionRequest(*lambda.GetFunctionInput) (*service.Request, *lambda.GetFunctionOutput)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)

	GetFunctionConfigurationRequest(*lambda.GetFunctionConfigurationInput) (*service.Request, *lambda.FunctionConfiguration)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)

	GetPolicyRequest(*lambda.GetPolicyInput) (*service.Request, *lambda.GetPolicyOutput)

	GetPolicy(*lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)

	InvokeRequest(*lambda.InvokeInput) (*service.Request, *lambda.InvokeOutput)

	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, error)

	InvokeAsyncRequest(*lambda.InvokeAsyncInput) (*service.Request, *lambda.InvokeAsyncOutput)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error)

	ListEventSourceMappingsRequest(*lambda.ListEventSourceMappingsInput) (*service.Request, *lambda.ListEventSourceMappingsOutput)

	ListEventSourceMappings(*lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)

	ListEventSourceMappingsPages(*lambda.ListEventSourceMappingsInput, func(*lambda.ListEventSourceMappingsOutput, bool) bool) error

	ListFunctionsRequest(*lambda.ListFunctionsInput) (*service.Request, *lambda.ListFunctionsOutput)

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)

	ListFunctionsPages(*lambda.ListFunctionsInput, func(*lambda.ListFunctionsOutput, bool) bool) error

	RemovePermissionRequest(*lambda.RemovePermissionInput) (*service.Request, *lambda.RemovePermissionOutput)

	RemovePermission(*lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)

	UpdateEventSourceMappingRequest(*lambda.UpdateEventSourceMappingInput) (*service.Request, *lambda.EventSourceMappingConfiguration)

	UpdateEventSourceMapping(*lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	UpdateFunctionCodeRequest(*lambda.UpdateFunctionCodeInput) (*service.Request, *lambda.FunctionConfiguration)

	UpdateFunctionCode(*lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)

	UpdateFunctionConfigurationRequest(*lambda.UpdateFunctionConfigurationInput) (*service.Request, *lambda.FunctionConfiguration)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
}
