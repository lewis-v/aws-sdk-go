// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudtrailiface provides an interface for the AWS CloudTrail.
package cloudtrailiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

// CloudTrailAPI is the interface type for cloudtrail.CloudTrail.
type CloudTrailAPI interface {
	CreateTrailRequest(*cloudtrail.CreateTrailInput) (*service.Request, *cloudtrail.CreateTrailOutput)

	CreateTrail(*cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)

	DeleteTrailRequest(*cloudtrail.DeleteTrailInput) (*service.Request, *cloudtrail.DeleteTrailOutput)

	DeleteTrail(*cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)

	DescribeTrailsRequest(*cloudtrail.DescribeTrailsInput) (*service.Request, *cloudtrail.DescribeTrailsOutput)

	DescribeTrails(*cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)

	GetTrailStatusRequest(*cloudtrail.GetTrailStatusInput) (*service.Request, *cloudtrail.GetTrailStatusOutput)

	GetTrailStatus(*cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)

	LookupEventsRequest(*cloudtrail.LookupEventsInput) (*service.Request, *cloudtrail.LookupEventsOutput)

	LookupEvents(*cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)

	StartLoggingRequest(*cloudtrail.StartLoggingInput) (*service.Request, *cloudtrail.StartLoggingOutput)

	StartLogging(*cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)

	StopLoggingRequest(*cloudtrail.StopLoggingInput) (*service.Request, *cloudtrail.StopLoggingOutput)

	StopLogging(*cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)

	UpdateTrailRequest(*cloudtrail.UpdateTrailInput) (*service.Request, *cloudtrail.UpdateTrailOutput)

	UpdateTrail(*cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
}
