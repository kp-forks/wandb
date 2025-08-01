// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the lifecycle configuration information set on the bucket. For
// information about lifecycle configuration, see [Object Lifecycle Management].
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, object size, or any
// combination of these. Accordingly, this section describes the latest API, which
// is compatible with the new functionality. The previous version of the API
// supported filtering based only on an object key name prefix, which is supported
// for general purpose buckets for backward compatibility. For the related API
// description, see [GetBucketLifecycle].
//
// Lifecyle configurations for directory buckets only support expiring objects and
// cancelling multipart uploads. Expiring of versioned objects, transitions and tag
// filters are not supported.
//
// Permissions
//   - General purpose bucket permissions - By default, all Amazon S3 resources
//     are private, including buckets, objects, and related subresources (for example,
//     lifecycle configuration and website configuration). Only the resource owner
//     (that is, the Amazon Web Services account that created it) can access the
//     resource. The resource owner can optionally grant access permissions to others
//     by writing an access policy. For this operation, a user must have the
//     s3:GetLifecycleConfiguration permission.
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
//   - Directory bucket permissions - You must have the
//     s3express:GetLifecycleConfiguration permission in an IAM identity-based policy
//     to use this operation. Cross-account access to this API operation isn't
//     supported. The resource owner can optionally grant access permissions to others
//     by creating a role or user for them as long as they are within the same account
//     as the owner and resource.
//
// For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM]in
//
//	the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
//
//	operation to the Regional endpoint. These endpoints support path-style requests
//	in the format https://s3express-control.region-code.amazonaws.com/bucket-name
//	. Virtual-hosted-style requests aren't supported. For more information about
//	endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
//	information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// GetBucketLifecycleConfiguration has the following special error:
//
//   - Error code: NoSuchLifecycleConfiguration
//
//   - Description: The lifecycle configuration does not exist.
//
//   - HTTP Status Code: 404 Not Found
//
//   - SOAP Fault Code Prefix: Client
//
// The following operations are related to GetBucketLifecycleConfiguration :
//
// [GetBucketLifecycle]
//
// [PutBucketLifecycle]
//
// [DeleteBucketLifecycle]
//
// [GetBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketLifecycleConfiguration", params, optFns, c.addOperationGetBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketLifecycleConfigurationInput struct {

	// The name of the bucket for which to get the lifecycle information.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	//
	// This parameter applies to general purpose buckets only. It is not supported for
	// directory bucket lifecycle configurations.
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *GetBucketLifecycleConfigurationInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type GetBucketLifecycleConfigurationOutput struct {

	// Container for a lifecycle rule.
	Rules []types.LifecycleRule

	// Indicates which default minimum object size behavior is applied to the
	// lifecycle configuration.
	//
	// This parameter applies to general purpose buckets only. It isn't supported for
	// directory bucket lifecycle configurations.
	//
	//   - all_storage_classes_128K - Objects smaller than 128 KB will not transition
	//   to any storage class by default.
	//
	//   - varies_by_storage_class - Objects smaller than 128 KB will transition to
	//   Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default,
	//   all other storage classes will prevent transitions smaller than 128 KB.
	//
	// To customize the minimum object size for any transition you can add a filter
	// that specifies a custom ObjectSizeGreaterThan or ObjectSizeLessThan in the body
	// of your transition rule. Custom filters always take precedence over the default
	// transition behavior.
	TransitionDefaultMinimumObjectSize types.TransitionDefaultMinimumObjectSize

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBucketLifecycleConfiguration"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIsExpressUserAgent(stack); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketLifecycleConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addInterceptAttempt(stack, options); err != nil {
		return err
	}
	if err = addInterceptExecution(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptTransmit(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeDeserialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterDeserialization(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func (v *GetBucketLifecycleConfigurationInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBucketLifecycleConfiguration",
	}
}

// getGetBucketLifecycleConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getGetBucketLifecycleConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketLifecycleConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketLifecycleConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketLifecycleConfigurationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
