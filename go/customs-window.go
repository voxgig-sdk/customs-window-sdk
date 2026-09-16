package voxgigcustomswindowsdk

import (
	"github.com/voxgig-sdk/customs-window-sdk/go/core"
	"github.com/voxgig-sdk/customs-window-sdk/go/entity"
	"github.com/voxgig-sdk/customs-window-sdk/go/feature"
	_ "github.com/voxgig-sdk/customs-window-sdk/go/utility"
)

// Type aliases preserve external API.
type CustomsWindowSDK = core.CustomsWindowSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type CustomsWindowEntity = core.CustomsWindowEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type CustomsWindowError = core.CustomsWindowError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewDebugFeatureFunc = func() core.Feature {
		return feature.NewDebugFeature()
	}
	core.NewIdempotencyFeatureFunc = func() core.Feature {
		return feature.NewIdempotencyFeature()
	}
	core.NewMetricsFeatureFunc = func() core.Feature {
		return feature.NewMetricsFeature()
	}
	core.NewPagingFeatureFunc = func() core.Feature {
		return feature.NewPagingFeature()
	}
	core.NewRatelimitFeatureFunc = func() core.Feature {
		return feature.NewRatelimitFeature()
	}
	core.NewRetryFeatureFunc = func() core.Feature {
		return feature.NewRetryFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewTimeoutFeatureFunc = func() core.Feature {
		return feature.NewTimeoutFeature()
	}
	core.NewBulkUploadEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewBulkUploadEntity(client, entopts)
	}
	core.NewFileEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewFileEntity(client, entopts)
	}
	core.NewPaginatedBulkUploadListListEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewPaginatedBulkUploadListListEntity(client, entopts)
	}
	core.NewPaginatedPartyListListEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewPaginatedPartyListListEntity(client, entopts)
	}
	core.NewPaginatedSubmissionListListEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewPaginatedSubmissionListListEntity(client, entopts)
	}
	core.NewPartyEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewPartyEntity(client, entopts)
	}
	core.NewSubmissionEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewSubmissionEntity(client, entopts)
	}
	core.NewSubmissionDetailEntityFunc = func(client *core.CustomsWindowSDK, entopts map[string]any) core.CustomsWindowEntity {
		return entity.NewSubmissionDetailEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewCustomsWindowSDK = core.NewCustomsWindowSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewCustomsWindowSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *CustomsWindowSDK  { return NewCustomsWindowSDK(nil) }
func Test() *CustomsWindowSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewDebugFeature = feature.NewDebugFeature
var NewIdempotencyFeature = feature.NewIdempotencyFeature
var NewMetricsFeature = feature.NewMetricsFeature
var NewPagingFeature = feature.NewPagingFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
