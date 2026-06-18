package imaging

import (
	"github.com/ugparu/onvif/xsd"
	"github.com/ugparu/onvif/xsd/onvif"
)

type Capabilities struct {
	ImageStabilization xsd.Boolean `xml:"ImageStabilization,attr" xmlu:"ImageStabilization,attr"`
	Presets            xsd.Boolean `xml:"Presets,attr" xmlu:"Presets,attr"`
}

type ImagingPreset struct {
	Name  onvif.Name           `xmlu:"Name"`
	Token onvif.ReferenceToken `xml:"token,attr" xmlu:"token,attr"`
	Type  string               `xml:"type,attr" xmlu:"type,attr"`
}

type GetServiceCapabilities struct {
	XMLName string `xml:"timg:GetServiceCapabilities"`
}

type GetImagingSettings struct {
	XMLName          string               `xml:"timg:GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetImagingSettings struct {
	XMLName          string                  `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.ImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"timg:ForcePersistence"`
}

type GetOptions struct {
	XMLName          string               `xml:"timg:GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Move struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMove      `xml:"timg:Focus"`
}

type GetMoveOptions struct {
	XMLName          string               `xml:"timg:GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Stop struct {
	XMLName          string               `xml:"timg:Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string               `xml:"timg:GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string               `xml:"timg:GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"timg:GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"timg:SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"timg:PresetToken"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xmlu:"Capabilities"`
}

type GetImagingSettingsResponse struct {
	ImagingSettings onvif.ImagingSettings20 `xmlu:"ImagingSettings"`
}

type SetImagingSettingsResponse struct {
}

type GetOptionsResponse struct {
	ImagingOptions onvif.ImagingOptions20 `xmlu:"ImagingOptions"`
}

type MoveResponse struct {
}

type GetMoveOptionsResponse struct {
	MoveOptions onvif.MoveOptions20 `xmlu:"MoveOptions"`
}

type StopResponse struct {
}

type GetStatusResponse struct {
	Status onvif.ImagingStatus20 `xmlu:"Status"`
}

type GetPresetsResponse struct {
	Preset []ImagingPreset `xmlu:"Preset"`
}

type GetCurrentPresetResponse struct {
	Preset *ImagingPreset `xmlu:"Preset"`
}

type SetCurrentPresetResponse struct {
}
