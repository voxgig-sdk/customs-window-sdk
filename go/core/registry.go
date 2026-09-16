package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewDebugFeatureFunc func() Feature

var NewIdempotencyFeatureFunc func() Feature

var NewMetricsFeatureFunc func() Feature

var NewPagingFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewBulkUploadEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewFileEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewPaginatedBulkUploadListListEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewPaginatedPartyListListEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewPaginatedSubmissionListListEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewPartyEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewSubmissionEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

var NewSubmissionDetailEntityFunc func(client *CustomsWindowSDK, entopts map[string]any) CustomsWindowEntity

