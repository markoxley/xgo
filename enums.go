package xgo

type RotateEnum byte

const (
	RotateLeft RotateEnum = iota
	RotateRight
)

type DirectionEnum byte

const (
	DirectionForward DirectionEnum = iota
	DirectionBackward
	DirectionLeft
	DirectionRight
)

type SpeedFrequencyEnum byte

const (
	ServoSpeed SpeedFrequencyEnum = iota
	SteppedFrequency
)

type SpeedEnum byte

const (
	SpeedFast   SpeedEnum = 0xf0
	SpeedNormal SpeedEnum = 0x80
	Speedslow   SpeedEnum = 0x10
)

type TranslationDirectionEnum byte

const (
	TranslationForward TranslationDirectionEnum = iota
	TranslationBackward
	TranslationLeftShift
	TranslationRightShift
)

type RotateDirectionEnum byte

const (
	RotateTurnLeft RotateDirectionEnum = iota
	RotateTurnRight
)

type OrientationEnum byte

const (
	OrientationX OrientationEnum = iota
	OrientationY
	OrientationZ
)

type BodyDirectionXYZEnum byte

const (
	BodyDirectionX BodyDirectionXYZEnum = iota
	BodyDirectionY
	BodyDirectionZ
)

type TranslationXYZEnum byte

const (
	TranslationX TranslationXYZEnum = iota
	TranslationY
	TranslationZ
)

type SwitchEnum byte

const (
	TurnOn SwitchEnum = iota
	TurnOff
)

type ServoSwitchEnum byte

const (
	Load ServoSwitchEnum = iota
	Unload
)

type BodyPartsEnum byte

const (
	BodyPartLeftFront BodyPartsEnum = iota
	BodyPartLeftHind
	BodyPartRightFront
	BodyPartRightHind
)

type JointEnum byte

const (
	JointUpper JointEnum = iota
	JointMiddle
	JointBelow
)

type TurnJointEnum byte

const (
	TurnJointUpper TurnJointEnum = iota
	TurnJointMiddle
	TurnJointBelow
)

type ActionEnum byte

const (
	ActionDefaultPosture ActionEnum = iota
	ActionGoProne
	ActionStand
	ActionCrawlForward
	ActionWhirl
	ActionSurPlace
	ActionSquat
	ActionTwirlRoll
	ActionTwirlPitch
	ActionTwirlYaw
	ActionTriaxialRotation
	ActionPee
	ActionSitDown
	ActionWave
	ActionStretchOneself
	ActionPlayPendulum
	ActionRequestFeeding
	ActionLookingForFood
	ActionHandshake
)
