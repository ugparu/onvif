package onvif

import (
	"github.com/ugparu/onvif/xsd"
)

// BUG(r): Enum types implemented as simple string

//TODO: enumerations
//TODO: type <typeName> struct {Any string} convert to type <typeName> AnyType
//TODO: process restrictions

//todo посмотреть все Extensions (Any string)
//todo что делать с xs:any = Any
//todo IntList и ему подобные. Проверить нужен ли слайс. Изменить на slice
//todo посмотреть можно ли заменить StreamType и ему подобные типы на вмтроенные типы
//todo оттестировать тип VideoSourceMode из-за Description-а

//todo в документации описать, что Capabilities повторяеся у каждого сервиса, поэтому у каждого свой Capabilities (MediaCapabilities)
//todo AuxiliaryData и другие simpleTypes, как реализовать рестрикшн
//todo Name и ему подобные необходимо изучить на наличие "List of..." ошибок

//todo Add in buit in

type ContentType string // minLength value="3"
type DNSName xsd.Token

type DeviceEntity struct {
	Token ReferenceToken `xml:"token,attr"`
}

type ReferenceToken xsd.String

type Name xsd.String

type IntRectangle struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type IntRectangleRange struct {
	XRange      IntRange
	YRange      IntRange
	WidthRange  IntRange
	HeightRange IntRange
}

type IntRange struct {
	Min int
	Max int
}

type FloatRange struct {
	Min float64 `xml:"Min"`
	Max float64 `xml:"Max"`
}

type OSDConfiguration struct {
	DeviceEntity                  `xml:"token,attr"`
	VideoSourceConfigurationToken OSDReference              `xml:"tt:VideoSourceConfigurationToken" xmlu:"VideoSourceConfigurationToken"`
	Type                          OSDType                   `xml:"tt:Type" xmlu:"Type"`
	Position                      OSDPosConfiguration       `xml:"tt:Position" xmlu:"Position"`
	TextString                    OSDTextConfiguration      `xml:"tt:TextString" xmlu:"TextString"`
	Image                         OSDImgConfiguration       `xml:"tt:Image" xmlu:"Image"`
	Extension                     OSDConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type OSDType xsd.String

type OSDPosConfiguration struct {
	Type      string                       `xml:"tt:Type" xmlu:"Type"`
	Pos       Vector                       `xml:"tt:Pos" xmlu:"Pos"`
	Extension OSDPosConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type Vector struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type OSDPosConfigurationExtension xsd.AnyType

type OSDReference ReferenceToken

type OSDTextConfiguration struct {
	IsPersistentText xsd.Boolean `xml:"IsPersistentText,attr"`

	Type            xsd.String                    `xml:"tt:Type" xmlu:"Type"`
	DateFormat      xsd.String                    `xml:"tt:DateFormat" xmlu:"DateFormat"`
	TimeFormat      xsd.String                    `xml:"tt:TimeFormat" xmlu:"TimeFormat"`
	FontSize        xsd.Int                       `xml:"tt:FontSize" xmlu:"FontSize"`
	FontColor       OSDColor                      `xml:"tt:FontColor" xmlu:"FontColor"`
	BackgroundColor OSDColor                      `xml:"tt:BackgroundColor" xmlu:"BackgroundColor"`
	PlainText       xsd.String                    `xml:"tt:PlainText" xmlu:"PlainText"`
	Extension       OSDTextConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type OSDColor struct {
	Transparent int `xml:"Transparent,attr"`

	Color Color `xml:"tt:Color" xmlu:"Color"`
}

type Color struct {
	X          float64    `xml:"X,attr"`
	Y          float64    `xml:"Y,attr"`
	Z          float64    `xml:"Z,attr"`
	Colorspace xsd.AnyURI `xml:"Colorspace,attr"`
}

type OSDTextConfigurationExtension xsd.AnyType

type OSDImgConfiguration struct {
	ImgPath   xsd.AnyURI                   `xml:"tt:ImgPath" xmlu:"ImgPath"`
	Extension OSDImgConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type OSDImgConfigurationExtension xsd.AnyType

type OSDConfigurationExtension xsd.AnyType

type VideoSource struct {
	DeviceEntity
	Framerate  float64
	Resolution VideoResolution
	Imaging    ImagingSettings
	Extension  VideoSourceExtension
}

type VideoResolution struct {
	Width  xsd.Int `xml:"tt:Width" xmlu:"Width"`
	Height xsd.Int `xml:"tt:Height" xmlu:"Height"`
}

type VideoResolutionClean struct {
	Width  xsd.Int
	Height xsd.Int
}

type ImagingSettings struct {
	BacklightCompensation BacklightCompensation
	Brightness            float64
	ColorSaturation       float64
	Contrast              float64
	Exposure              Exposure
	Focus                 FocusConfiguration
	IrCutFilter           IrCutFilterMode
	Sharpness             float64
	WideDynamicRange      WideDynamicRange
	WhiteBalance          WhiteBalance
	Extension             ImagingSettingsExtension
}

type BacklightCompensation struct {
	Mode  BacklightCompensationMode
	Level float64
}

type BacklightCompensationMode xsd.String

type Exposure struct {
	Mode            ExposureMode
	Priority        ExposurePriority
	Window          Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain         float64
	MaxGain         float64
	MinIris         float64
	MaxIris         float64
	ExposureTime    float64
	Gain            float64
	Iris            float64
}

type ExposureMode xsd.String

type ExposurePriority xsd.String

type Rectangle struct {
	Bottom float64 `xml:"bottom,attr"`
	Top    float64 `xml:"top,attr"`
	Right  float64 `xml:"right,attr"`
	Left   float64 `xml:"left,attr"`
}

type FocusConfiguration struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed  float64
	NearLimit     float64
	FarLimit      float64
}

type AutoFocusMode xsd.String

type IrCutFilterMode xsd.String

type WideDynamicRange struct {
	Mode  WideDynamicMode `xml:"tt:Mode"`
	Level float64         `xml:"tt:Level"`
}

type WideDynamicMode xsd.String

type WhiteBalance struct {
	Mode   WhiteBalanceMode
	CrGain float64
	CbGain float64
}

type WhiteBalanceMode xsd.String

type ImagingSettingsExtension xsd.AnyType

type VideoSourceExtension struct {
	Imaging   ImagingSettings20
	Extension VideoSourceExtension2
}

type ImagingSettings20 struct {
	BacklightCompensation *BacklightCompensation20    `xml:"tt:BacklightCompensation"`
	Brightness            float64                     `xml:"tt:Brightness,omitempty"`
	ColorSaturation       float64                     `xml:"tt:ColorSaturation,omitempty"`
	Contrast              float64                     `xml:"tt:Contrast,omitempty"`
	Exposure              *Exposure20                 `xml:"tt:Exposure"`
	Focus                 *FocusConfiguration20       `xml:"tt:Focus"`
	IrCutFilter           *IrCutFilterMode            `xml:"tt:IrCutFilter"`
	Sharpness             float64                     `xml:"tt:Sharpness,omitempty"`
	WideDynamicRange      *WideDynamicRange20         `xml:"tt:WideDynamicRange"`
	WhiteBalance          *WhiteBalance20             `xml:"tt:WhiteBalance"`
	Extension             *ImagingSettingsExtension20 `xml:"tt:Extension"`
}

type BacklightCompensation20 struct {
	Mode  BacklightCompensationMode `xml:"tt:Mode"`
	Level float64                   `xml:"tt:Level"`
}

type Exposure20 struct {
	Mode            ExposureMode     `xml:"tt:Mode,omitempty"`
	Priority        ExposurePriority `xml:"tt:Priority,omitempty"`
	Window          Rectangle        `xml:"tt:Window,omitempty"`
	MinExposureTime float64          `xml:"tt:MinExposureTime,omitempty"`
	MaxExposureTime float64          `xml:"tt:MaxExposureTime,omitempty"`
	MinGain         float64          `xml:"tt:MinGain,omitempty"`
	MaxGain         float64          `xml:"tt:MaxGain,omitempty"`
	MinIris         float64          `xml:"tt:MinIris,omitempty"`
	MaxIris         float64          `xml:"tt:MaxIris,omitempty"`
	ExposureTime    float64          `xml:"tt:ExposureTime,omitempty"`
	Gain            float64          `xml:"tt:Gain,omitempty"`
	Iris            float64          `xml:"tt:Iris,omitempty"`
}

type FocusConfiguration20 struct {
	AutoFocusMode AutoFocusMode                 `xml:"tt:AutoFocusMode"`
	DefaultSpeed  float64                       `xml:"tt:DefaultSpeed"`
	NearLimit     float64                       `xml:"tt:NearLimit"`
	FarLimit      float64                       `xml:"tt:FarLimit"`
	Extension     FocusConfiguration20Extension `xml:"tt:Extension"`
}

type FocusConfiguration20Extension xsd.AnyType

type WideDynamicRange20 struct {
	Mode  WideDynamicMode `xml:"tt:Mode"`
	Level float64         `xml:"tt:Level"`
}

type WhiteBalance20 struct {
	Mode      WhiteBalanceMode        `xml:"tt:Mode"`
	CrGain    float64                 `xml:"tt:CrGain"`
	CbGain    float64                 `xml:"tt:CbGain"`
	Extension WhiteBalance20Extension `xml:"tt:Extension"`
}

type WhiteBalance20Extension xsd.AnyType

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization          `xml:"tt:ImageStabilization"`
	Extension          ImagingSettingsExtension202 `xml:"tt:Extension"`
}

type ImageStabilization struct {
	Mode      ImageStabilizationMode      `xml:"tt:Mode"`
	Level     float64                     `xml:"tt:Level"`
	Extension ImageStabilizationExtension `xml:"tt:Extension"`
}

type ImageStabilizationMode xsd.String

type ImageStabilizationExtension xsd.AnyType

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment   `xml:"tt:IrCutFilterAutoAdjustment"`
	Extension                 ImagingSettingsExtension203 `xml:"tt:Extension"`
}

type IrCutFilterAutoAdjustment struct {
	BoundaryType   string                             `xml:"tt:BoundaryType"`
	BoundaryOffset float64                            `xml:"tt:BoundaryOffset"`
	ResponseTime   xsd.Duration                       `xml:"tt:ResponseTime"`
	Extension      IrCutFilterAutoAdjustmentExtension `xml:"tt:Extension"`
}

type IrCutFilterAutoAdjustmentExtension xsd.AnyType

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation            `xml:"tt:ToneCompensation"`
	Defogging        Defogging                   `xml:"tt:Defogging"`
	NoiseReduction   NoiseReduction              `xml:"tt:NoiseReduction"`
	Extension        ImagingSettingsExtension204 `xml:"tt:Extension"`
}

type ToneCompensation struct {
	Mode      string                    `xml:"tt:Mode"`
	Level     float64                   `xml:"tt:Level"`
	Extension ToneCompensationExtension `xml:"tt:Extension"`
}

type ToneCompensationExtension xsd.AnyType

type Defogging struct {
	Mode      string
	Level     float64
	Extension DefoggingExtension
}

type DefoggingExtension xsd.AnyType

type NoiseReduction struct {
	Level float64 `xml:"tt:Level" xmlu:"Level"`
}

type ImagingSettingsExtension204 xsd.AnyType

type VideoSourceExtension2 xsd.AnyType

type AudioSource struct {
	DeviceEntity
	Channels int
}

type AudioOutput struct {
	DeviceEntity
}

type Profile struct {
	Token                       ReferenceToken `xml:"token,attr"`
	Fixed                       bool           `xml:"fixed,attr"`
	Name                        Name
	VideoSourceConfiguration    VideoSourceConfiguration
	AudioSourceConfiguration    AudioSourceConfiguration
	VideoEncoderConfiguration   VideoEncoderConfiguration
	AudioEncoderConfiguration   AudioEncoderConfiguration
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration
	PTZConfiguration            PTZConfiguration
	MetadataConfiguration       MetadataConfiguration
	Extension                   ProfileExtension
}

type VideoSourceConfiguration struct {
	ConfigurationEntity
	ViewMode    string                            `xml:"ViewMode,attr"`
	SourceToken ReferenceToken                    `xml:"tt:SourceToken" xmlu:"SourceToken"`
	Bounds      IntRectangle                      `xml:"tt:Bounds" xmlu:"Bounds"`
	Extension   VideoSourceConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type ConfigurationEntity struct {
	Token    ReferenceToken `xml:"token,attr"`
	Name     Name           `xml:"tt:Name" xmlu:"Name" `
	UseCount int            `xml:"tt:UseCount" xmlu:"UseCount"`
}

type ConfigurationEntityClean struct {
	Token    ReferenceToken `xml:"token,attr"`
	Name     Name
	UseCount int
}

type VideoSourceConfigurationExtension struct {
	Rotate    Rotate                             `xml:"tt:Rotate" xmlu:"Rotate"`
	Extension VideoSourceConfigurationExtension2 `xml:"tt:Extension" xmlu:"Extension"`
}

type Rotate struct {
	Mode      RotateMode      `xml:"tt:Mode" xmlu:"Mode"`
	Degree    xsd.Int         `xml:"tt:Degree"  xmlu:"Degree"`
	Extension RotateExtension `xml:"tt:Extension"  xmlu:"Extension"`
}

type RotateMode xsd.String

type RotateExtension xsd.AnyType

type VideoSourceConfigurationExtension2 struct {
	LensDescription  LensDescription  `xml:"tt:LensDescription"`
	SceneOrientation SceneOrientation `xml:"tt:SceneOrientation"`
}

type LensDescription struct {
	FocalLength float64        `xml:"FocalLength,attr"`
	Offset      LensOffset     `xml:"tt:Offset"`
	Projection  LensProjection `xml:"tt:Projection"`
	XFactor     float64        `xml:"tt:XFactor"`
}

type LensOffset struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type LensProjection struct {
	Angle         float64 `xml:"tt:Angle"`
	Radius        float64 `xml:"tt:Radius"`
	Transmittance float64 `xml:"tt:Transmittance"`
}

type SceneOrientation struct {
	Mode        SceneOrientationMode `xml:"tt:Mode"`
	Orientation xsd.String           `xml:"tt:Orientation"`
}

type SceneOrientationMode xsd.String

type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken ReferenceToken `xml:"tt:SourceToken" xmlu:"SourceToken"`
}

type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       VideoEncoding          `xml:"tt:Encoding" xmlu:"Encoding"`
	Resolution     VideoResolution        `xml:"tt:Resolution" xmlu:"Resolution"`
	Quality        float64                `xml:"tt:Quality" xmlu:"Quality"`
	RateControl    VideoRateControl       `xml:"tt:RateControl" xmlu:"RateControl"`
	MPEG4          Mpeg4Configuration     `xml:"tt:MPEG4" xmlu:"MPEG4"`
	H264           H264Configuration      `xml:"tt:H264" xmlu:"H264"`
	Multicast      MulticastConfiguration `xml:"tt:Multicast" xmlu:"Multicast"`
	SessionTimeout xsd.Duration           `xml:"tt:SessionTimeout" xmlu:"SessionTimeout"`
}

type VideoEncoderConfigurationClean struct {
	ConfigurationEntityClean
	Encoding       VideoEncoding
	Resolution     VideoResolutionClean
	Quality        float64
	RateControl    VideoRateControlClean
	MPEG4          Mpeg4ConfigurationClean
	H264           H264ConfigurationClean
	Multicast      MulticastConfigurationClean
	SessionTimeout xsd.Duration
}

type VideoEncoding xsd.String

type VideoRateControl struct {
	FrameRateLimit   xsd.Int `xml:"tt:FrameRateLimit" xmlu:"FrameRateLimit"`
	EncodingInterval xsd.Int `xml:"tt:EncodingInterval" xmlu:"EncodingInterval"`
	BitrateLimit     xsd.Int `xml:"tt:BitrateLimit" xmlu:"BitrateLimit"`
}

type VideoRateControlClean struct {
	FrameRateLimit   xsd.Int
	EncodingInterval xsd.Int
	BitrateLimit     xsd.Int
}

type Mpeg4Configuration struct {
	GovLength    xsd.Int      `xml:"tt:GovLength" xmlu:"GovLength"`
	Mpeg4Profile Mpeg4Profile `xml:"tt:Mpeg4Profile" xmlu:"Mpeg4Profile"`
}

type Mpeg4ConfigurationClean struct {
	GovLength    xsd.Int
	Mpeg4Profile Mpeg4Profile
}

type Mpeg4Profile xsd.String

type H264Configuration struct {
	GovLength   xsd.Int     `xml:"tt:GovLength" xmlu:"GovLength"`
	H264Profile H264Profile `xml:"tt:H264Profile" xmlu:"H264Profile"`
}

type H264ConfigurationClean struct {
	GovLength   xsd.Int
	H264Profile H264Profile
}

type H264Profile xsd.String

type MulticastConfiguration struct {
	Address   IPAddress   `xml:"tt:Address" xmlu:"Address"`
	Port      int         `xml:"tt:Port" xmlu:"Port"`
	TTL       int         `xml:"tt:TTL" xmlu:"TTL"`
	AutoStart xsd.Boolean `xml:"tt:AutoStart" xmlu:"AutoStart"`
}

type MulticastConfigurationClean struct {
	Address   IPAddressClean
	Port      int
	TTL       int
	AutoStart xsd.Boolean
}

type IPAddress struct {
	Type        IPType      `xml:"tt:Type" xmlu:"Type"`
	IPv4Address IPv4Address `xml:"tt:IPv4Address" xmlu:"IPv4Address"`
	IPv6Address IPv6Address `xml:"tt:IPv6Address" xmlu:"IPv6Address"`
}

type IPAddressClean struct {
	Type        IPType
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type IPType xsd.String

// IPv4 address
type IPv4Address xsd.Token

// IPv6 address
type IPv6Address xsd.Token

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       AudioEncoding          `xml:"tt:Encoding" xmlu:"Encoding"`
	Bitrate        int                    `xml:"tt:Bitrate" xmlu:"Bitrate"`
	SampleRate     int                    `xml:"tt:SampleRate" xmlu:"SampleRate"`
	Multicast      MulticastConfiguration `xml:"tt:Multicast" xmlu:"Multicast"`
	SessionTimeout xsd.Duration           `xml:"tt:SessionTimeout" xmlu:"SessionTimeout"`
}

type AudioEncoding xsd.String

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"tt:AnalyticsEngineConfiguration"`
	RuleEngineConfiguration      RuleEngineConfiguration      `xml:"tt:RuleEngineConfiguration"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule Config                                `xml:"tt:AnalyticsModule"`
	Extension       AnalyticsEngineConfigurationExtension `xml:"tt:Extension"`
}

type Config struct {
	Name       string    `xml:"Name,attr"`
	Type       xsd.QName `xml:"Type,attr"`
	Parameters ItemList  `xml:"tt:Parameters"`
}

type ItemList struct {
	SimpleItem  SimpleItem        `xml:"tt:SimpleItem"`
	ElementItem ElementItem       `xml:"tt:ElementItem"`
	Extension   ItemListExtension `xml:"tt:Extension"`
}

type SimpleItem struct {
	Name  string            `xml:"Name,attr"`
	Value xsd.AnySimpleType `xml:"Value,attr"`
}

type ElementItem struct {
	Name string `xml:"Name,attr"`
}

type ItemListExtension xsd.AnyType

type AnalyticsEngineConfigurationExtension xsd.AnyType

type RuleEngineConfiguration struct {
	Rule      Config                           `xml:"tt:Rule"`
	Extension RuleEngineConfigurationExtension `xml:"tt:Extension"`
}

type RuleEngineConfigurationExtension xsd.AnyType

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp                               int                       `xml:"MoveRamp,attr"`
	PresetRamp                             int                       `xml:"PresetRamp,attr"`
	PresetTourRamp                         int                       `xml:"PresetTourRamp,attr"`
	NodeToken                              ReferenceToken            `xml:"NodeToken"`
	DefaultAbsolutePantTiltPositionSpace   xsd.AnyURI                `xml:"DefaultAbsolutePantTiltPositionSpace"`
	DefaultAbsoluteZoomPositionSpace       xsd.AnyURI                `xml:"DefaultAbsoluteZoomPositionSpace"`
	DefaultRelativePanTiltTranslationSpace xsd.AnyURI                `xml:"DefaultRelativePanTiltTranslationSpace"`
	DefaultRelativeZoomTranslationSpace    xsd.AnyURI                `xml:"DefaultRelativeZoomTranslationSpace"`
	DefaultContinuousPanTiltVelocitySpace  xsd.AnyURI                `xml:"DefaultContinuousPanTiltVelocitySpace"`
	DefaultContinuousZoomVelocitySpace     xsd.AnyURI                `xml:"DefaultContinuousZoomVelocitySpace"`
	DefaultPTZSpeed                        PTZSpeed                  `xml:"DefaultPTZSpeed"`
	DefaultPTZTimeout                      xsd.Duration              `xml:"DefaultPTZTimeout"`
	PanTiltLimits                          PanTiltLimits             `xml:"PanTiltLimits"`
	ZoomLimits                             ZoomLimits                `xml:"ZoomLimits"`
	Extension                              PTZConfigurationExtension `xml:"Extension"`
}

type PTZSpeed struct {
	PanTilt Vector2D `xml:"tt:PanTilt"`
	Zoom    Vector1D `xml:"tt:Zoom"`
}

type Vector2D struct {
	X     float64    `xml:"x,attr"`
	Y     float64    `xml:"y,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type Vector1D struct {
	X     float64    `xml:"x,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type PanTiltLimits struct {
	Range Space2DDescription `xml:"Range"`
}

type Space2DDescription struct {
	URI    xsd.AnyURI `xml:"URI"`
	XRange FloatRange `xml:"XRange"`
	YRange FloatRange `xml:"YRange"`
}

type ZoomLimits struct {
	Range Space1DDescription `xml:"Range"`
}

type Space1DDescription struct {
	URI    xsd.AnyURI `xml:"URI"`
	XRange FloatRange `xml:"XRange"`
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection         `xml:"tt:PTControlDirection"`
	Extension          PTZConfigurationExtension2 `xml:"tt:Extension"`
}

type PTControlDirection struct {
	EFlip     EFlip                       `xml:"tt:EFlip"`
	Reverse   Reverse                     `xml:"tt:Reverse"`
	Extension PTControlDirectionExtension `xml:"tt:Extension"`
}

type EFlip struct {
	Mode EFlipMode `xml:"tt:Mode"`
}

type EFlipMode xsd.String

type Reverse struct {
	Mode ReverseMode `xml:"tt:Mode"`
}

type ReverseMode xsd.String

type PTControlDirectionExtension xsd.AnyType

type PTZConfigurationExtension2 xsd.AnyType

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType              string                         `xml:"CompressionType,attr"`
	PTZStatus                    PTZFilter                      `xml:"tt:PTZStatus"`
	Events                       EventSubscription              `xml:"tt:Events"`
	Analytics                    xsd.Boolean                    `xml:"tt:Analytics"`
	Multicast                    MulticastConfiguration         `xml:"tt:Multicast"`
	SessionTimeout               xsd.Duration                   `xml:"tt:SessionTimeout"`
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration   `xml:"tt:AnalyticsEngineConfiguration"`
	Extension                    MetadataConfigurationExtension `xml:"tt:Extension"`
}

type PTZFilter struct {
	Status   bool `xml:"tt:Status"`
	Position bool `xml:"tt:Position"`
}

type EventSubscription struct {
	Filter             FilterType `xml:"tt:Filter"`
	SubscriptionPolicy `xml:"tt:SubscriptionPolicy"`
}

type FilterType xsd.AnyType

type SubscriptionPolicy xsd.AnyType

type MetadataConfigurationExtension xsd.AnyType

type ProfileExtension struct {
	AudioOutputConfiguration  AudioOutputConfiguration
	AudioDecoderConfiguration AudioDecoderConfiguration
	Extension                 ProfileExtension2
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken ReferenceToken `xml:"tt:OutputToken" xmlu:"OutputToken"`
	SendPrimacy xsd.AnyURI     `xml:"tt:SendPrimacy" xmlu:"SendPrimacy"`
	OutputLevel int            `xml:"tt:OutputLevel" xmlu:"OutputLevel"`
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type ProfileExtension2 xsd.AnyType

type VideoSourceConfigurationOptions struct {
	MaximumNumberOfProfiles    int `xml:"MaximumNumberOfProfiles,attr"`
	BoundsRange                IntRectangleRange
	VideoSourceTokensAvailable ReferenceToken
	Extension                  VideoSourceConfigurationOptionsExtension
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate    RotateOptions
	Extension VideoSourceConfigurationOptionsExtension2
}

type RotateOptions struct {
	Mode       RotateMode
	DegreeList IntList
	Extension  RotateOptionsExtension
}

type IntList struct {
	Items []int
}

type RotateOptionsExtension xsd.AnyType

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode SceneOrientationMode
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange
	JPEG         JpegOptions
	MPEG4        Mpeg4Options
	H264         H264Options
	Extension    VideoEncoderOptionsExtension
}

type JpegOptions struct {
	ResolutionsAvailable  VideoResolution
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
}

type Mpeg4Options struct {
	ResolutionsAvailable   VideoResolution
	GovLengthRange         IntRange
	FrameRateRange         IntRange
	EncodingIntervalRange  IntRange
	Mpeg4ProfilesSupported Mpeg4Profile
}

type H264Options struct {
	ResolutionsAvailable  VideoResolution
	GovLengthRange        IntRange
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
	H264ProfilesSupported H264Profile
}

type VideoEncoderOptionsExtension struct {
	JPEG      JpegOptions2
	MPEG4     Mpeg4Options2
	H264      H264Options2
	Extension VideoEncoderOptionsExtension2
}

type JpegOptions2 struct {
	JpegOptions
	BitrateRange IntRange
}

type Mpeg4Options2 struct {
	Mpeg4Options
	BitrateRange IntRange
}

type H264Options2 struct {
	H264Options
	BitrateRange IntRange
}

type VideoEncoderOptionsExtension2 xsd.AnyType

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable ReferenceToken
	Extension            AudioSourceOptionsExtension
}

type AudioSourceOptionsExtension xsd.AnyType

type AudioEncoderConfigurationOptions struct {
	Options AudioEncoderConfigurationOption
}

type AudioEncoderConfigurationOption struct {
	Encoding       AudioEncoding
	BitrateList    IntList
	SampleRateList IntList
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions
	Extension              MetadataConfigurationOptionsExtension
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported   bool
	ZoomStatusSupported      bool
	PanTiltPositionSupported bool
	ZoomPositionSupported    bool
	Extension                PTZStatusFilterOptionsExtension
}

type PTZStatusFilterOptionsExtension xsd.AnyType

type MetadataConfigurationOptionsExtension struct {
	CompressionType string
	Extension       MetadataConfigurationOptionsExtension2
}

type MetadataConfigurationOptionsExtension2 xsd.AnyType

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable ReferenceToken
	SendPrimacyOptions    xsd.AnyURI
	OutputLevelRange      IntRange
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions  AACDecOptions
	G711DecOptions G711DecOptions
	G726DecOptions G726DecOptions
	Extension      AudioDecoderConfigurationOptionsExtension
}

type AACDecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G711DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G726DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type AudioDecoderConfigurationOptionsExtension xsd.AnyType

type StreamSetup struct {
	Stream    StreamType `xml:"tt:Stream" xmlu:"Stream"`
	Transport Transport  `xml:"tt:Transport" xmlu:"Transport"`
}

type StreamType xsd.String

type Transport struct {
	Protocol TransportProtocol `xml:"tt:Protocol" xmlu:"Protocol"`
	Tunnel   *Transport        `xml:"tt:Tunnel" xmlu:"Tunnel"`
}

// enum
type TransportProtocol xsd.String

type MediaUri struct {
	Uri                 xsd.AnyURI
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
	Timeout             xsd.Duration
}

type VideoSourceMode struct {
	Token         ReferenceToken `xml:"token,attr"`
	Enabled       bool           `xml:"Enabled,attr"`
	MaxFramerate  float64
	MaxResolution VideoResolution
	Encodings     EncodingTypes
	Reboot        bool
	Description   Description
	Extension     VideoSourceModeExtension
}

type EncodingTypes struct {
	EncodingTypes []string
}

type Description struct {
	Description string
}

type VideoSourceModeExtension xsd.AnyType

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs MaximumNumberOfOSDs
	Type                OSDType
	PositionOption      string
	TextOption          OSDTextOptions
	ImageOption         OSDImgOptions
	Extension           OSDConfigurationOptionsExtension
}

type MaximumNumberOfOSDs struct {
	Total       int `xml:"Total,attr"`
	Image       int `xml:"Image,attr"`
	PlainText   int `xml:"PlainText,attr"`
	Date        int `xml:"Date,attr"`
	Time        int `xml:"Time,attr"`
	DateAndTime int `xml:"DateAndTime,attr"`
}

type OSDTextOptions struct {
	Type            string
	FontSizeRange   IntRange
	DateFormat      string
	TimeFormat      string
	FontColor       OSDColorOptions
	BackgroundColor OSDColorOptions
	Extension       OSDTextOptionsExtension
}

type OSDColorOptions struct {
	Color       ColorOptions
	Transparent IntRange
	Extension   OSDColorOptionsExtension
}

type ColorOptions struct {
	ColorList       Color
	ColorspaceRange ColorspaceRange
}

type ColorspaceRange struct {
	X          FloatRange
	Y          FloatRange
	Z          FloatRange
	Colorspace xsd.AnyURI
}

type OSDColorOptionsExtension xsd.AnyType

type OSDTextOptionsExtension xsd.AnyType

type OSDImgOptions struct {
	FormatsSupported StringAttrList `xml:"FormatsSupported,attr"`
	MaxSize          int            `xml:"MaxSize,attr"`
	MaxWidth         int            `xml:"MaxWidth,attr"`
	MaxHeight        int            `xml:"MaxHeight,attr"`

	ImagePath xsd.AnyURI
	Extension OSDImgOptionsExtension
}

type StringAttrList struct {
	AttrList []string
}

type OSDImgOptionsExtension xsd.AnyType

type OSDConfigurationOptionsExtension xsd.AnyType

//PTZ

type PTZNode struct {
	DeviceEntity
	FixedHomePosition      xsd.Boolean `xml:"FixedHomePosition,attr"`
	GeoMove                xsd.Boolean `xml:"GeoMove,attr"`
	Name                   Name
	SupportedPTZSpaces     PTZSpaces
	MaximumNumberOfPresets int
	HomeSupported          xsd.Boolean
	AuxiliaryCommands      AuxiliaryData
	Extension              PTZNodeExtension
}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    Space2DDescription
	AbsoluteZoomPositionSpace       Space1DDescription
	RelativePanTiltTranslationSpace Space2DDescription
	RelativeZoomTranslationSpace    Space1DDescription
	ContinuousPanTiltVelocitySpace  Space2DDescription
	ContinuousZoomVelocitySpace     Space1DDescription
	PanTiltSpeedSpace               Space1DDescription
	ZoomSpeedSpace                  Space1DDescription
	Extension                       PTZSpacesExtension
}

type PTZSpacesExtension xsd.AnyType

// TODO: restriction
type AuxiliaryData xsd.String

type PTZNodeExtension struct {
	SupportedPresetTour PTZPresetTourSupported
	Extension           PTZNodeExtension2
}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours int
	PTZPresetTourOperation     PTZPresetTourOperation
	Extension                  PTZPresetTourSupportedExtension
}

type PTZPresetTourOperation xsd.String
type PTZPresetTourSupportedExtension xsd.AnyType

type PTZNodeExtension2 xsd.AnyType

type PTZConfigurationOptions struct {
	PTZRamps           IntAttrList `xml:"PTZRamps,attr"`
	Spaces             PTZSpaces
	PTZTimeout         DurationRange
	PTControlDirection PTControlDirectionOptions
	Extension          PTZConfigurationOptions2
}

type IntAttrList struct {
	IntAttrList []int
}

type DurationRange struct {
	Min xsd.Duration
	Max xsd.Duration
}

type PTControlDirectionOptions struct {
	EFlip     EFlipOptions
	Reverse   ReverseOptions
	Extension PTControlDirectionOptionsExtension
}

type EFlipOptions struct {
	Mode      EFlipMode
	Extension EFlipOptionsExtension
}

type EFlipOptionsExtension xsd.AnyType

type ReverseOptions struct {
	Mode      ReverseMode
	Extension ReverseOptionsExtension
}

type ReverseOptionsExtension xsd.AnyType

type PTControlDirectionOptionsExtension xsd.AnyType

type PTZConfigurationOptions2 xsd.AnyType

type PTZPreset struct {
	Token       ReferenceToken `xml:"token,attr"`
	Name        Name
	PTZPosition PTZVector
}

type PTZVector struct {
	PanTilt Vector2D `xml:"tt:PanTilt"`
	Zoom    Vector1D `xml:"tt:Zoom"`
}

type PTZStatus struct {
	Position   PTZVector
	MoveStatus PTZMoveStatus
	Error      string
	UtcTime    xsd.DateTime
}

type PTZMoveStatus struct {
	PanTilt MoveStatus
	Zoom    MoveStatus
}

type MoveStatus struct {
	Status string
}

type GeoLocation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

type PresetTour struct {
	Token             ReferenceToken                 `xml:"token,attr"`
	Name              Name                           `xml:"tt:Name"`
	Status            PTZPresetTourStatus            `xml:"tt:Status"`
	AutoStart         xsd.Boolean                    `xml:"tt:AutoStart"`
	StartingCondition PTZPresetTourStartingCondition `xml:"tt:StartingCondition"`
	TourSpot          PTZPresetTourSpot              `xml:"tt:TourSpot"`
	Extension         PTZPresetTourExtension         `xml:"tt:Extension"`
}

type PTZPresetTourStatus struct {
	State           PTZPresetTourState           `xml:"tt:State"`
	CurrentTourSpot PTZPresetTourSpot            `xml:"tt:CurrentTourSpot"`
	Extension       PTZPresetTourStatusExtension `xml:"tt:Extension"`
}

type PTZPresetTourState xsd.String

type PTZPresetTourSpot struct {
	PresetDetail PTZPresetTourPresetDetail  `xml:"tt:PresetDetail"`
	Speed        PTZSpeed                   `xml:"tt:Speed"`
	StayTime     xsd.Duration               `xml:"tt:StayTime"`
	Extension    PTZPresetTourSpotExtension `xml:"tt:Extension"`
}

type PTZPresetTourPresetDetail struct {
	PresetToken   ReferenceToken             `xml:"tt:PresetToken"`
	Home          xsd.Boolean                `xml:"tt:Home"`
	PTZPosition   PTZVector                  `xml:"tt:PTZPosition"`
	TypeExtension PTZPresetTourTypeExtension `xml:"tt:TypeExtension"`
}

type PTZPresetTourTypeExtension xsd.AnyType

type PTZPresetTourSpotExtension xsd.AnyType

type PTZPresetTourStatusExtension xsd.AnyType

type PTZPresetTourStartingCondition struct {
	RandomPresetOrder xsd.Boolean                             `xml:"RandomPresetOrder,attr"`
	RecurringTime     xsd.Int                                 `xml:"tt:RecurringTime"`
	RecurringDuration xsd.Duration                            `xml:"tt:RecurringDuration"`
	Direction         PTZPresetTourDirection                  `xml:"tt:Direction"`
	Extension         PTZPresetTourStartingConditionExtension `xml:"tt:Extension"`
}

type PTZPresetTourDirection xsd.String

type PTZPresetTourStartingConditionExtension xsd.AnyType

type PTZPresetTourExtension xsd.AnyType

type PTZPresetTourOptions struct {
	AutoStart         xsd.Boolean
	StartingCondition PTZPresetTourStartingConditionOptions
	TourSpot          PTZPresetTourSpotOptions
}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime     IntRange
	RecurringDuration DurationRange
	Direction         PTZPresetTourDirection
	Extension         PTZPresetTourStartingConditionOptionsExtension
}

type PTZPresetTourStartingConditionOptionsExtension xsd.AnyType

type PTZPresetTourSpotOptions struct {
	PresetDetail PTZPresetTourPresetDetailOptions
	StayTime     DurationRange
}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          ReferenceToken
	Home                 xsd.Boolean
	PanTiltPositionSpace Space2DDescription
	ZoomPositionSpace    Space1DDescription
	Extension            PTZPresetTourPresetDetailOptionsExtension
}

type PTZPresetTourPresetDetailOptionsExtension xsd.AnyType

//Device

type OnvifVersion struct {
	Major int
	Minor int
}

type SetDateTimeType xsd.String

type TimeZone struct {
	TZ xsd.Token `xml:"tt:TZ"`
}

type SystemDateTime struct {
	DateTimeType    SetDateTimeType
	DaylightSavings xsd.Boolean
	TimeZone        TimeZone
	UTCDateTime     xsd.DateTime
	LocalDateTime   xsd.DateTime
	Extension       SystemDateTimeExtension
}

type SystemDateTimeExtension xsd.AnyType

type FactoryDefaultType xsd.String

type AttachmentData struct {
	ContentType ContentType `xml:"contentType,attr"`
	Include     Include     `xml:"inc:Include"`
}

type Include struct {
	Href xsd.AnyURI `xml:"href,attr"`
}

type BackupFile struct {
	Name string         `xml:"tt:Name"`
	Data AttachmentData `xml:"tt:Data"`
}

type SystemLogType xsd.String

type SystemLog struct {
	Binary AttachmentData
	String string
}

type SupportInformation struct {
	Binary AttachmentData
	String string
}

type Scope struct {
	ScopeDef  ScopeDefinition
	ScopeItem xsd.AnyURI
}

type ScopeDefinition xsd.String

type DiscoveryMode xsd.String

type NetworkHost struct {
	Type        NetworkHostType      `xml:"tt:Type" xmlu:"Type"`
	IPv4Address IPv4Address          `xml:"tt:IPv4Address" xmlu:"IPv4Address"`
	IPv6Address IPv6Address          `xml:"tt:IPv6Address" xmlu:"IPv6Address"`
	DNSname     DNSName              `xml:"tt:DNSname" xmlu:"DNSname"`
	Extension   NetworkHostExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type NetworkHostType xsd.String

type NetworkHostExtension xsd.String

type RemoteUser struct {
	Username           string      `xml:"tt:Username" xmlu:"Username"`
	Password           string      `xml:"tt:Password" xmlu:"Password"`
	UseDerivedPassword xsd.Boolean `xml:"tt:UseDerivedPassword" xmlu:"UseDerivedPassword"`
}

type User struct {
	Username  string        `xml:"tt:Username" xmlu:"Username"`
	Password  string        `xml:"tt:Password" xmlu:"Password"`
	UserLevel UserLevel     `xml:"tt:UserLevel" xmlu:"UserLevel"`
	Extension UserExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type UserLevel xsd.String

type UserExtension xsd.String

type CapabilityCategory xsd.String

// Capabilities of device
type Capabilities struct {
	Analytics AnalyticsCapabilities
	Device    DeviceCapabilities
	Events    EventCapabilities
	Imaging   ImagingCapabilities
	Media     MediaCapabilities
	PTZ       PTZCapabilities
	Extension CapabilitiesExtension
}

// AnalyticsCapabilities Check
type AnalyticsCapabilities struct {
	XAddr                  xsd.AnyURI
	RuleSupport            xsd.Boolean
	AnalyticsModuleSupport xsd.Boolean
}

// DeviceCapabilities Check
type DeviceCapabilities struct {
	XAddr     xsd.AnyURI
	Network   NetworkCapabilities
	System    SystemCapabilities
	IO        IOCapabilities
	Security  SecurityCapabilities
	Extension DeviceCapabilitiesExtension
}

// NetworkCapabilities Check
type NetworkCapabilities struct {
	IPFilter          xsd.Boolean
	ZeroConfiguration xsd.Boolean
	IPVersion6        xsd.Boolean
	DynDNS            xsd.Boolean
	Extension         NetworkCapabilitiesExtension
}

// NetworkCapabilitiesExtension Check
type NetworkCapabilitiesExtension struct {
	Dot11Configuration xsd.Boolean
	Extension          NetworkCapabilitiesExtension2
}

// NetworkCapabilitiesExtension2 Extension2
type NetworkCapabilitiesExtension2 xsd.AnyType

// SystemCapabilities check
type SystemCapabilities struct {
	DiscoveryResolve  xsd.Boolean
	DiscoveryBye      xsd.Boolean
	RemoteDiscovery   xsd.Boolean
	SystemBackup      xsd.Boolean
	SystemLogging     xsd.Boolean
	FirmwareUpgrade   xsd.Boolean
	SupportedVersions OnvifVersion
	Extension         SystemCapabilitiesExtension
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    xsd.Boolean
	HttpSystemBackup       xsd.Boolean
	HttpSystemLogging      xsd.Boolean
	HttpSupportInformation xsd.Boolean
	Extension              SystemCapabilitiesExtension2
}

type SystemCapabilitiesExtension2 xsd.AnyType

type IOCapabilities struct {
	InputConnectors int
	RelayOutputs    int
	Extension       IOCapabilitiesExtension
}

type IOCapabilitiesExtension struct {
	Auxiliary         xsd.Boolean
	AuxiliaryCommands AuxiliaryData
	Extension         IOCapabilitiesExtension2
}

type IOCapabilitiesExtension2 xsd.AnyType

type SecurityCapabilities struct {
	TLS1_1               xsd.Boolean
	TLS1_2               xsd.Boolean
	OnboardKeyGeneration xsd.Boolean
	AccessPolicyConfig   xsd.Boolean
	X_509Token           xsd.Boolean
	SAMLToken            xsd.Boolean
	KerberosToken        xsd.Boolean
	RELToken             xsd.Boolean
	Extension            SecurityCapabilitiesExtension
}

type SecurityCapabilitiesExtension struct {
	TLS1_0    xsd.Boolean
	Extension SecurityCapabilitiesExtension2
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              xsd.Boolean
	SupportedEAPMethod int
	RemoteUserHandling xsd.Boolean
}

type DeviceCapabilitiesExtension xsd.AnyType

type EventCapabilities struct {
	XAddr                                         xsd.AnyURI
	WSSubscriptionPolicySupport                   xsd.Boolean
	WSPullPointSupport                            xsd.Boolean
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean
}

type ImagingCapabilities struct {
	XAddr xsd.AnyURI
}

type MediaCapabilities struct {
	XAddr                 xsd.AnyURI
	StreamingCapabilities RealTimeStreamingCapabilities
	Extension             MediaCapabilitiesExtension
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast xsd.Boolean
	RTP_TCP      xsd.Boolean
	RTP_RTSP_TCP xsd.Boolean
	Extension    RealTimeStreamingCapabilitiesExtension
}

type RealTimeStreamingCapabilitiesExtension xsd.AnyType

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int
}

type PTZCapabilities struct {
	XAddr xsd.AnyURI
}

type CapabilitiesExtension struct {
	DeviceIO        DeviceIOCapabilities
	Display         DisplayCapabilities
	Recording       RecordingCapabilities
	Search          SearchCapabilities
	Replay          ReplayCapabilities
	Receiver        ReceiverCapabilities
	AnalyticsDevice AnalyticsDeviceCapabilities
	Extensions      CapabilitiesExtension2
}

type DeviceIOCapabilities struct {
	XAddr        xsd.AnyURI
	VideoSources int
	VideoOutputs int
	AudioSources int
	AudioOutputs int
	RelayOutputs int
}

type DisplayCapabilities struct {
	XAddr       xsd.AnyURI
	FixedLayout xsd.Boolean
}

type RecordingCapabilities struct {
	XAddr              xsd.AnyURI
	ReceiverSource     xsd.Boolean
	MediaProfileSource xsd.Boolean
	DynamicRecordings  xsd.Boolean
	DynamicTracks      xsd.Boolean
	MaxStringLength    int
}

type SearchCapabilities struct {
	XAddr          xsd.AnyURI
	MetadataSearch xsd.Boolean
}

type ReplayCapabilities struct {
	XAddr xsd.AnyURI
}

type ReceiverCapabilities struct {
	XAddr                xsd.AnyURI
	RTP_Multicast        xsd.Boolean
	RTP_TCP              xsd.Boolean
	RTP_RTSP_TCP         xsd.Boolean
	SupportedReceivers   int
	MaximumRTSPURILength int
}

type AnalyticsDeviceCapabilities struct {
	XAddr       xsd.AnyURI
	RuleSupport xsd.Boolean
	Extension   AnalyticsDeviceExtension
}

type AnalyticsDeviceExtension xsd.AnyType

type CapabilitiesExtension2 xsd.AnyType

type HostnameInformation struct {
	FromDHCP  xsd.Boolean
	Name      xsd.Token
	Extension HostnameInformationExtension
}

type HostnameInformationExtension xsd.AnyType

type DNSInformation struct {
	FromDHCP     xsd.Boolean
	SearchDomain xsd.Token
	DNSFromDHCP  IPAddress
	DNSManual    IPAddress
	Extension    DNSInformationExtension
}

type DNSInformationExtension xsd.AnyType

type NTPInformation struct {
	FromDHCP    xsd.Boolean
	NTPFromDHCP NetworkHost
	NTPManual   NetworkHost
	Extension   NTPInformationExtension
}

type NTPInformationExtension xsd.AnyType

type DynamicDNSInformation struct {
	Type      DynamicDNSType
	Name      DNSName
	TTL       xsd.Duration
	Extension DynamicDNSInformationExtension
}

// TODO: enumeration
type DynamicDNSType xsd.String

type DynamicDNSInformationExtension xsd.AnyType

type NetworkInterface struct {
	DeviceEntity
	Enabled   xsd.Boolean
	Info      NetworkInterfaceInfo
	Link      NetworkInterfaceLink
	IPv4      IPv4NetworkInterface
	IPv6      IPv6NetworkInterface
	Extension NetworkInterfaceExtension
}

type NetworkInterfaceInfo struct {
	Name      xsd.String
	HwAddress HwAddress
	MTU       xsd.Int
}

type HwAddress xsd.Token

type NetworkInterfaceLink struct {
	AdminSettings NetworkInterfaceConnectionSetting
	OperSettings  NetworkInterfaceConnectionSetting
	InterfaceType IANA_IfTypes `xml:"IANA-IfTypes"`
}

type IANA_IfTypes xsd.Int

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation xsd.Boolean `xml:"tt:AutoNegotiation" xmlu:"AutoNegotiation"`
	Speed           xsd.Int     `xml:"tt:Speed" xmlu:"Speed"`
	Duplex          Duplex      `xml:"tt:Duplex" xmlu:"Duplex"`
}

// TODO: enum
type Duplex xsd.String

type NetworkInterfaceExtension struct {
	InterfaceType IANA_IfTypes
	Dot3          Dot3Configuration
	Dot11         Dot11Configuration
	Extension     NetworkInterfaceExtension2
}

type NetworkInterfaceExtension2 xsd.AnyType

type Dot11Configuration struct {
	SSID     Dot11SSIDType                  `xml:"tt:SSID" xmlu:"SSID"`
	Mode     Dot11StationMode               `xml:"tt:Mode" xmlu:"Mode"`
	Alias    Name                           `xml:"tt:Alias" xmlu:"Alias"`
	Priority NetworkInterfaceConfigPriority `xml:"tt:Priority" xmlu:"Priority"`
	Security Dot11SecurityConfiguration     `xml:"tt:Security" xmlu:"Security"`
}

type Dot11SecurityConfiguration struct {
	Mode      Dot11SecurityMode                   `xml:"tt:Mode" xmlu:"Mode"`
	Algorithm Dot11Cipher                         `xml:"tt:Algorithm" xmlu:"Algorithm"`
	PSK       Dot11PSKSet                         `xml:"tt:PSK" xmlu:"PSK"`
	Dot1X     ReferenceToken                      `xml:"tt:Dot1X" xmlu:"Dot1X"`
	Extension Dot11SecurityConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type Dot11SecurityConfigurationExtension xsd.AnyType

type Dot11PSKSet struct {
	Key        Dot11PSK             `xml:"tt:Key" xmlu:"Key"`
	Passphrase Dot11PSKPassphrase   `xml:"tt:Passphrase" xmlu:"Passphrase"`
	Extension  Dot11PSKSetExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type Dot11PSKSetExtension xsd.AnyType

type Dot11PSKPassphrase xsd.String

type Dot11PSK xsd.HexBinary

// TODO: enumeration
type Dot11Cipher xsd.String

// TODO: enumeration
type Dot11SecurityMode xsd.String

// TODO: restrictions
type NetworkInterfaceConfigPriority xsd.Integer

// TODO: enumeration
type Dot11StationMode xsd.String

// TODO: restrictions
type Dot11SSIDType xsd.HexBinary

type Dot3Configuration xsd.String

type IPv6NetworkInterface struct {
	Enabled xsd.Boolean
	Config  IPv6Configuration
}

type IPv6Configuration struct {
	AcceptRouterAdvert xsd.Boolean
	DHCP               IPv6DHCPConfiguration
	Manual             PrefixedIPv6Address
	LinkLocal          PrefixedIPv6Address
	FromDHCP           PrefixedIPv6Address
	FromRA             PrefixedIPv6Address
	Extension          IPv6ConfigurationExtension
}

type IPv6ConfigurationExtension xsd.AnyType

type PrefixedIPv6Address struct {
	Address      IPv6Address `xml:"tt:Address" xmlu:"Address"`
	PrefixLength xsd.Int     `xml:"tt:PrefixLength" xmlu:"PrefixLength"`
}

// TODO: enumeration
type IPv6DHCPConfiguration xsd.String

type IPv4NetworkInterface struct {
	Enabled xsd.Boolean
	Config  IPv4Configuration
}

type IPv4Configuration struct {
	Manual    PrefixedIPv4Address
	LinkLocal PrefixedIPv4Address
	FromDHCP  PrefixedIPv4Address
	DHCP      xsd.Boolean
}

// optional, unbounded
type PrefixedIPv4Address struct {
	Address      IPv4Address `xml:"tt:Address" xmlu:"Address"`
	PrefixLength xsd.Int     `xml:"tt:PrefixLength" xmlu:"PrefixLength"`
}

type NetworkInterfaceSetConfiguration struct {
	Enabled   xsd.Boolean                               `xml:"tt:Enabled" xmlu:"Enabled"`
	Link      NetworkInterfaceConnectionSetting         `xml:"tt:Link" xmlu:"Link"`
	MTU       xsd.Int                                   `xml:"tt:MTU" xmlu:"MTU"`
	IPv4      IPv4NetworkInterfaceSetConfiguration      `xml:"tt:IPv4" xmlu:"IPv4"`
	IPv6      IPv6NetworkInterfaceSetConfiguration      `xml:"tt:IPv6" xmlu:"IPv6"`
	Extension NetworkInterfaceSetConfigurationExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3      Dot3Configuration                          `xml:"tt:Dot3" xmlu:"Dot3"`
	Dot11     Dot11Configuration                         `xml:"tt:Dot11" xmlu:"Dot11"`
	Extension NetworkInterfaceSetConfigurationExtension2 `xml:"tt:Extension" xmlu:"Extension"`
}

type NetworkInterfaceSetConfigurationExtension2 xsd.AnyType

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled            xsd.Boolean           `xml:"tt:Enabled" xmlu:"Enabled"`
	AcceptRouterAdvert xsd.Boolean           `xml:"tt:AcceptRouterAdvert" xmlu:"AcceptRouterAdvert"`
	Manual             PrefixedIPv6Address   `xml:"tt:Manual" xmlu:"Manual"`
	DHCP               IPv6DHCPConfiguration `xml:"tt:DHCP" xmlu:"DHCP"`
}

type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled xsd.Boolean         `xml:"tt:Enabled" xmlu:"Enabled"`
	Manual  PrefixedIPv4Address `xml:"tt:Manual" xmlu:"Manual"`
	DHCP    xsd.Boolean         `xml:"tt:DHCP" xmlu:"DHCP"`
}

type NetworkProtocol struct {
	Name      NetworkProtocolType      `xml:"tt:Name" xmlu:"Name"`
	Enabled   xsd.Boolean              `xml:"tt:Enabled" xmlu:"Enabled"`
	Port      xsd.Int                  `xml:"tt:Port" xmlu:"Port"`
	Extension NetworkProtocolExtension `xml:"tt:Extension" xmlu:"Extension"`
}

type NetworkProtocolExtension xsd.AnyType

// TODO: enumeration
type NetworkProtocolType xsd.String

type NetworkGateway struct {
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type NetworkZeroConfiguration struct {
	InterfaceToken ReferenceToken
	Enabled        xsd.Boolean
	Addresses      IPv4Address
	Extension      NetworkZeroConfigurationExtension
}

type NetworkZeroConfigurationExtension struct {
	Additional *NetworkZeroConfiguration
	Extension  NetworkZeroConfigurationExtension2
}

type NetworkZeroConfigurationExtension2 xsd.AnyType

type IPAddressFilter struct {
	Type        IPAddressFilterType      `xml:"tt:Type" xmlu:"Type"`
	IPv4Address PrefixedIPv4Address      `xml:"tt:IPv4Address,omitempty" xmlu:"IPv4Address,omitempty"`
	IPv6Address PrefixedIPv6Address      `xml:"tt:IPv6Address,omitempty" xmlu:"IPv6Address,omitempty"`
	Extension   IPAddressFilterExtension `xml:"tt:Extension,omitempty" xmlu:"Extension,omitempty"`
}

type IPAddressFilterExtension xsd.AnyType

// enum { 'Allow', 'Deny' }
// TODO: enumeration
type IPAddressFilterType xsd.String

// TODO: attribite <xs:attribute ref="xmime:contentType" use="optional"/>
type BinaryData struct {
	X    ContentType      `xml:"xmime:contentType,attr"`
	Data xsd.Base64Binary `xml:"tt:Data"`
}

type Certificate struct {
	CertificateID xsd.Token  `xml:"tt:CertificateID"`
	Certificate   BinaryData `xml:"tt:Certificate"`
}

type CertificateStatus struct {
	CertificateID xsd.Token   `xml:"tt:CertificateID"`
	Status        xsd.Boolean `xml:"tt:Status"`
}

type RelayOutput struct {
	DeviceEntity
	Properties RelayOutputSettings
}

type RelayOutputSettings struct {
	Mode      RelayMode      `xml:"tt:Mode"`
	DelayTime xsd.Duration   `xml:"tt:DelayTime"`
	IdleState RelayIdleState `xml:"tt:IdleState"`
}

// TODO:enumeration
type RelayIdleState xsd.String

// TODO: enumeration
type RelayMode xsd.String

// TODO: enumeration
type RelayLogicalState xsd.String

type CertificateWithPrivateKey struct {
	CertificateID xsd.Token  `xml:"tt:CertificateID"`
	Certificate   BinaryData `xml:"tt:Certificate"`
	PrivateKey    BinaryData `xml:"tt:PrivateKey"`
}

type CertificateInformation struct {
	CertificateID      xsd.Token
	IssuerDN           xsd.String
	SubjectDN          xsd.String
	KeyUsage           CertificateUsage
	ExtendedKeyUsage   CertificateUsage
	KeyLength          xsd.Int
	Version            xsd.String
	SerialNum          xsd.String
	SignatureAlgorithm xsd.String
	Validity           DateTimeRange
	Extension          CertificateInformationExtension
}

type CertificateInformationExtension xsd.AnyType

type DateTimeRange struct {
	From  xsd.DateTime
	Until xsd.DateTime
}

type CertificateUsage struct {
	Critical         xsd.Boolean `xml:"Critical,attr"`
	CertificateUsage xsd.String
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken              `xml:"tt:Dot1XConfigurationToken"`
	Identity                xsd.String                  `xml:"tt:Identity"`
	AnonymousID             xsd.String                  `xml:"tt:AnonymousID,omitempty"`
	EAPMethod               xsd.Int                     `xml:"tt:EAPMethod"`
	CACertificateID         xsd.Token                   `xml:"tt:CACertificateID,omitempty"`
	EAPMethodConfiguration  EAPMethodConfiguration      `xml:"tt:EAPMethodConfiguration,omitempty"`
	Extension               Dot1XConfigurationExtension `xml:"tt:Extension,omitempty"`
}

type Dot1XConfigurationExtension xsd.AnyType

type EAPMethodConfiguration struct {
	TLSConfiguration TLSConfiguration   `xml:"tt:TLSConfiguration,omitempty"`
	Password         xsd.String         `xml:"tt:Password,omitempty"`
	Extension        EapMethodExtension `xml:"tt:Extension,omitempty"`
}

type EapMethodExtension xsd.AnyType

type TLSConfiguration struct {
	CertificateID xsd.Token `xml:"tt:CertificateID,omitempty"`
}

type Dot11Capabilities struct {
	TKIP                  xsd.Boolean
	ScanAvailableNetworks xsd.Boolean
	MultipleConfiguration xsd.Boolean
	AdHocStationMode      xsd.Boolean
	WEP                   xsd.Boolean
}

type Dot11Status struct {
	SSID              Dot11SSIDType
	BSSID             xsd.String
	PairCipher        Dot11Cipher
	GroupCipher       Dot11Cipher
	SignalStrength    Dot11SignalStrength
	ActiveConfigAlias ReferenceToken
}

// TODO: enumeration
type Dot11SignalStrength xsd.String

type Dot11AvailableNetworks struct {
	SSID                  Dot11SSIDType
	BSSID                 xsd.String
	AuthAndMangementSuite Dot11AuthAndMangementSuite
	PairCipher            Dot11Cipher
	GroupCipher           Dot11Cipher
	SignalStrength        Dot11SignalStrength
	Extension             Dot11AvailableNetworksExtension
}

type Dot11AvailableNetworksExtension xsd.AnyType

// TODO: enumeration
type Dot11AuthAndMangementSuite xsd.String

type SystemLogUriList struct {
	SystemLog SystemLogUri
}

type SystemLogUri struct {
	Type SystemLogType
	Uri  xsd.AnyURI
}

type LocationEntity struct {
	Entity    xsd.String     `xml:"Entity,attr"`
	Token     ReferenceToken `xml:"Token,attr"`
	Fixed     xsd.Boolean    `xml:"Fixed,attr"`
	GeoSource xsd.AnyURI     `xml:"GeoSource,attr"`
	AutoGeo   xsd.Boolean    `xml:"AutoGeo,attr"`

	GeoLocation      GeoLocation      `xml:"tt:GeoLocation"`
	GeoOrientation   GeoOrientation   `xml:"tt:GeoOrientation"`
	LocalLocation    LocalLocation    `xml:"tt:LocalLocation"`
	LocalOrientation LocalOrientation `xml:"tt:LocalOrientation"`
}

type LocalOrientation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

type LocalLocation struct {
	X xsd.Float `xml:"x,attr"`
	Y xsd.Float `xml:"y,attr"`
	Z xsd.Float `xml:"z,attr"`
}

type GeoOrientation struct {
	Roll  xsd.Float `xml:"roll,attr"`
	Pitch xsd.Float `xml:"pitch,attr"`
	Yaw   xsd.Float `xml:"yaw,attr"`
}

type FocusMove struct {
	Absolute   AbsoluteFocus   `xml:"tt:Absolute"`
	Relative   RelativeFocus   `xml:"tt:Relative"`
	Continuous ContinuousFocus `xml:"tt:Continuous"`
}

type ContinuousFocus struct {
	Speed xsd.Float `xml:"tt:Speed" xmlu:"Speed"`
}

type RelativeFocus struct {
	Distance xsd.Float `xml:"tt:Distance" xmlu:"Distance"`
	Speed    xsd.Float `xml:"tt:Speed" xmlu:"Speed"`
}

type AbsoluteFocus struct {
	Position xsd.Float `xml:"tt:Position" xmlu:"Position"`
	Speed    xsd.Float `xml:"tt:Speed" xmlu:"Speed"`
}

type DateTime struct {
	Time Time `xml:"tt:Time" xmlu:"Time"`
	Date Date `xml:"tt:Date" xmlu:"Date"`
}

type Time struct {
	Hour   xsd.Int `xml:"tt:Hour" xmlu:"Hour"`
	Minute xsd.Int `xml:"tt:Minute" xmlu:"Minute"`
	Second xsd.Int `xml:"tt:Second" xmlu:"Second"`
}

type Date struct {
	Year  xsd.Int `xml:"tt:Year" xmlu:"Year"`
	Month xsd.Int `xml:"tt:Month" xmlu:"Month"`
	Day   xsd.Int `xml:"tt:Day" xmlu:"Day"`
}
