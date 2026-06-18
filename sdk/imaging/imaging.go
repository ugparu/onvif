package imaging

//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetServiceCapabilities
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetImagingSettings
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging SetImagingSettings
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetOptions
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging Move
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetMoveOptions
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging Stop
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetStatus
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetPresets
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging GetCurrentPreset
//go:generate go run github.com/ugparu/onvif/sdk/codegen imaging imaging SetCurrentPreset
