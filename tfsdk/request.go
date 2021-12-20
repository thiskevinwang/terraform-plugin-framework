package tfsdk

// ConfigureProviderRequest represents a request containing the values the user
// specified for the provider configuration block, along with other runtime
// information from Terraform or the Plugin SDK. An instance of this request
// struct is supplied as an argument to the provider's Configure function.
type ConfigureProviderRequest struct {
	// TerraformVersion is the version of Terraform executing the request.
	// This is supplied for logging, analytics, and User-Agent purposes
	// only. Providers should not try to gate provider behavior on
	// Terraform versions.
	TerraformVersion string

	// Config is the configuration the user supplied for the provider. This
	// information should usually be persisted to the underlying type
	// that's implementing the Provider interface, for use in later
	// resource CRUD operations.
	Config ReadOnlyData
}

// CreateResourceRequest represents a request for the provider to create a
// resource. An instance of this request struct is supplied as an argument to
// the resource's Create function.
type CreateResourceRequest struct {
	// Config is the configuration the user supplied for the resource.
	//
	// This configuration may contain unknown values if a user uses
	// interpolation or other functionality that would prevent Terraform
	// from knowing the value at request time.
	Config ReadOnlyData

	// Plan is the planned state for the resource.
	Plan *Data

	// ProviderMeta is metadata from the provider_meta block of the module.
	ProviderMeta ReadOnlyData
}

// ReadResourceRequest represents a request for the provider to read a
// resource, i.e., update values in state according to the real state of the
// resource. An instance of this request struct is supplied as an argument to
// the resource's Read function.
type ReadResourceRequest struct {
	// State is the current state of the resource prior to the Read
	// operation.
	State *Data

	// ProviderMeta is metadata from the provider_meta block of the module.
	ProviderMeta ReadOnlyData
}

// UpdateResourceRequest represents a request for the provider to update a
// resource. An instance of this request struct is supplied as an argument to
// the resource's Update function.
type UpdateResourceRequest struct {
	// Config is the configuration the user supplied for the resource.
	//
	// This configuration may contain unknown values if a user uses
	// interpolation or other functionality that would prevent Terraform
	// from knowing the value at request time.
	Config ReadOnlyData

	// Plan is the planned state for the resource.
	Plan *Data

	// State is the current state of the resource prior to the Update
	// operation.
	State *Data

	// ProviderMeta is metadata from the provider_meta block of the module.
	ProviderMeta ReadOnlyData
}

// DeleteResourceRequest represents a request for the provider to delete a
// resource. An instance of this request struct is supplied as an argument to
// the resource's Delete function.
type DeleteResourceRequest struct {
	// State is the current state of the resource prior to the Delete
	// operation.
	State *Data

	// ProviderMeta is metadata from the provider_meta block of the module.
	ProviderMeta ReadOnlyData
}

// ModifyResourcePlanRequest represents a request for the provider to modify the
// planned new state that Terraform has generated for the resource.
type ModifyResourcePlanRequest struct {
	// Config is the configuration the user supplied for the resource.
	//
	// This configuration may contain unknown values if a user uses
	// interpolation or other functionality that would prevent Terraform
	// from knowing the value at request time.
	Config ReadOnlyData

	// State is the current state of the resource.
	State *Data

	// Plan is the planned new state for the resource.
	Plan *Data

	// ProviderMeta is metadata from the provider_meta block of the module.
	ProviderMeta ReadOnlyData
}

// ReadDataSourceRequest represents a request for the provider to read a data
// source, i.e., update values in state according to the real state of the
// data source. An instance of this request struct is supplied as an argument
// to the data source's Read function.
type ReadDataSourceRequest struct {
	// Config is the configuration the user supplied for the data source.
	//
	// This configuration may contain unknown values if a user uses
	// interpolation or other functionality that would prevent Terraform
	// from knowing the value at request time.
	Config ReadOnlyData

	// ProviderMeta is metadata from the provider_meta block of the module.
	ProviderMeta ReadOnlyData
}
