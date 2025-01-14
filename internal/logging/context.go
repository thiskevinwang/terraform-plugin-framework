package logging

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
)

// InitContext creates SDK logger contexts. The incoming context will
// already have the root SDK logger and root provider logger setup from
// terraform-plugin-go tf6server RPC handlers.
func InitContext(ctx context.Context) context.Context {
	ctx = tfsdklog.NewSubsystem(ctx, SubsystemFramework,
		// All calls are through the Framework* helper functions
		tfsdklog.WithAdditionalLocationOffset(1),
		tfsdklog.WithLevelFromEnv(EnvTfLogSdkFramework),
	)

	return ctx
}
