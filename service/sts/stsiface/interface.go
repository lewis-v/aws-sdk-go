// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package stsiface provides an interface for the AWS Security Token Service.
package stsiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/sts"
)

// STSAPI is the interface type for sts.STS.
type STSAPI interface {
	AssumeRoleRequest(*sts.AssumeRoleInput) (*service.Request, *sts.AssumeRoleOutput)

	AssumeRole(*sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error)

	AssumeRoleWithSAMLRequest(*sts.AssumeRoleWithSAMLInput) (*service.Request, *sts.AssumeRoleWithSAMLOutput)

	AssumeRoleWithSAML(*sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error)

	AssumeRoleWithWebIdentityRequest(*sts.AssumeRoleWithWebIdentityInput) (*service.Request, *sts.AssumeRoleWithWebIdentityOutput)

	AssumeRoleWithWebIdentity(*sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error)

	DecodeAuthorizationMessageRequest(*sts.DecodeAuthorizationMessageInput) (*service.Request, *sts.DecodeAuthorizationMessageOutput)

	DecodeAuthorizationMessage(*sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error)

	GetFederationTokenRequest(*sts.GetFederationTokenInput) (*service.Request, *sts.GetFederationTokenOutput)

	GetFederationToken(*sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error)

	GetSessionTokenRequest(*sts.GetSessionTokenInput) (*service.Request, *sts.GetSessionTokenOutput)

	GetSessionToken(*sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error)
}
