// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ResendSubscriptionConfirmationRequest wrapper for the ResendSubscriptionConfirmation operation
type ResendSubscriptionConfirmationRequest struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription to resend the confirmation for.
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ResendSubscriptionConfirmationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ResendSubscriptionConfirmationRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ResendSubscriptionConfirmationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ResendSubscriptionConfirmationResponse wrapper for the ResendSubscriptionConfirmation operation
type ResendSubscriptionConfirmationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Subscription instance
	Subscription `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ResendSubscriptionConfirmationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ResendSubscriptionConfirmationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
