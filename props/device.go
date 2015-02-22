package props

import (
	"github.com/Shaked/gomobiledetect"
)

type Device struct {
	IsMobile bool   `json:"isMobile"`
	Browser  Type   `json:"browser"`
	OS       Type   `json:"os"`
	Type     Type   `json:"type"`
	Prop     string `json:"device"`
	Grade    string `json:"grade"`
}

func NewDevice() *Device {
	device := &Device{IsMobile: false}
	return device
}

func (d *Device) Configure(m *mobiledetect.MobileDetect) {
	isPhone := m.IsMobile()
	isTablet := m.IsTablet()

	if isPhone || isTablet {
		d.IsMobile = true
		d.Prop = "phone"

		if isPhone {
			d.Prop = "phone"
		} else if isTablet {
			d.Prop = "tablet"
		}

		detectOs(m, d)
		detectBrowser(m, d)
		detectType(m, d)
		d.Grade = m.MobileGrade()
	}
}

func detectOs(m *mobiledetect.MobileDetect, d *Device) {
	for key, p := range mapos {
		if m.IsKey(key) {
			d.OS = p
			break
		}
	}
}

func detectBrowser(m *mobiledetect.MobileDetect, d *Device) {
	for key, p := range mapbrows {
		if m.IsKey(key) {
			d.Browser = p
			break
		}
	}
}

func detectType(m *mobiledetect.MobileDetect, d *Device) {
	if d.Prop == "phone" {
		for key, gob := range mapmob {
			if m.IsKey(key) {
				d.Type = gob
				break
			}
		}

	} else if d.Prop == "tablet" {
		for key, gob := range maptab {
			if m.IsKey(key) {
				d.Type = gob
				break
			}
		}
	}
}

type Type struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var mapos = map[int]Type{
	mobiledetect.ANDROIDOS:       Type{"AndroidOS", 1},
	mobiledetect.IOS:             Type{"iOS", 7},
	mobiledetect.WINDOWSPHONEOS:  Type{"WindowsPhoneOS", 6},
	mobiledetect.WINDOWSMOBILEOS: Type{"WindowsMobileOS", 5},
	mobiledetect.BLACKBERRYOS:    Type{"BlackBerryOS", 2},
	mobiledetect.WEBOS:           Type{"webOS", 11},
	mobiledetect.SYMBIANOS:       Type{"SymbianOS", 4},
	mobiledetect.MEEGOOS:         Type{"MeeGoOS", 8},
	mobiledetect.MAEMOOS:         Type{"MaemoOS", 9},
	mobiledetect.JAVAOS:          Type{"JavaOS", 10},
	mobiledetect.BADAOS:          Type{"badaOS", 12},
	mobiledetect.BREWOS:          Type{"BREWOS", 13},
	mobiledetect.PALMOS:          Type{"PalmOS", 3},
}

var mapmob = map[int]Type{
	mobiledetect.IPHONE:       Type{"iPhone", 1},
	mobiledetect.BLACKBERRY:   Type{"BlackBerry", 2},
	mobiledetect.HTC:          Type{"HTC", 3},
	mobiledetect.NEXUS:        Type{"Nexus", 4},
	mobiledetect.DELL:         Type{"Dell", 5},
	mobiledetect.MOTOROLA:     Type{"Motorola", 6},
	mobiledetect.SAMSUNG:      Type{"Samsung", 7},
	mobiledetect.LG:           Type{"LG", 8},
	mobiledetect.SONY:         Type{"Sony", 9},
	mobiledetect.ASUS:         Type{"Asus", 10},
	mobiledetect.MICROMAX:     Type{"Micromax", 11},
	mobiledetect.PALM:         Type{"Palm", 12},
	mobiledetect.VERTU:        Type{"Vertu", 13},
	mobiledetect.PANTECH:      Type{"Pantech", 14},
	mobiledetect.FLY:          Type{"Fly", 15},
	mobiledetect.SIMVALLEY:    Type{"SimValley", 16},
	mobiledetect.GENERICPHONE: Type{"GenericPhone", 17},
	mobiledetect.IMOBILE:      Type{"IMobile", 18},
	mobiledetect.WOLFGANG:     Type{"Wolfgang", 19},
	mobiledetect.ALCATEL:      Type{"Alcatel", 20},
	mobiledetect.AMOI:         Type{"Amoi", 21},
	mobiledetect.INQ:          Type{"Inq", 22},
}

var maptab = map[int]Type{
	mobiledetect.IPAD:              Type{"iPad", 18},
	mobiledetect.SAMSUNGTABLET:     Type{"SamsungTablet", 20},
	mobiledetect.KINDLE:            Type{"Kindle", 21},
	mobiledetect.SURFACETABLET:     Type{"SurfaceTablet", 22},
	mobiledetect.HPTABLET:          Type{"HPTablet", 23},
	mobiledetect.ASUSTABLET:        Type{"AsusTablet", 24},
	mobiledetect.BLACKBERRYTABLET:  Type{"BlackBerryTablet", 25},
	mobiledetect.HTCTABLET:         Type{"HTCtablet", 26},
	mobiledetect.MOTOROLATABLET:    Type{"MotorolaTablet", 27},
	mobiledetect.NOOKTABLET:        Type{"NookTablet", 28},
	mobiledetect.ACERTABLET:        Type{"AcerTablet", 29},
	mobiledetect.TOSHIBATABLET:     Type{"ToshibaTablet", 30},
	mobiledetect.LGTABLET:          Type{"LGTablet", 31},
	mobiledetect.FUJITSUTABLET:     Type{"FujitsuTablet", 32},
	mobiledetect.PRESTIGIOTABLET:   Type{"PrestigioTablet", 33},
	mobiledetect.LENOVOTABLET:      Type{"LenovoTablet", 34},
	mobiledetect.YARVIKTABLET:      Type{"YarvikTablet", 35},
	mobiledetect.MEDIONTABLET:      Type{"MedionTablet", 36},
	mobiledetect.ARNOVATABLET:      Type{"ArnovaTablet", 37},
	mobiledetect.INTENSOTABLET:     Type{"IntensoTablet", 38},
	mobiledetect.IRUTABLET:         Type{"IRUTablet", 39},
	mobiledetect.MEGAFONTABLET:     Type{"MegafonTablet", 40},
	mobiledetect.EBODATABLET:       Type{"EbodaTablet", 41},
	mobiledetect.ALLVIEWTABLET:     Type{"AllViewTablet", 42},
	mobiledetect.ARCHOSTABLET:      Type{"ArchosTablet", 43},
	mobiledetect.AINOLTABLET:       Type{"AinolTablet", 44},
	mobiledetect.SONYTABLET:        Type{"SonyTablet", 45},
	mobiledetect.CUBETABLET:        Type{"CubeTablet", 46},
	mobiledetect.COBYTABLET:        Type{"CobyTablet", 47},
	mobiledetect.MIDTABLET:         Type{"MIDTablet", 48},
	mobiledetect.SMITTABLET:        Type{"SMiTTablet", 49},
	mobiledetect.ROCKCHIPTABLET:    Type{"RockChipTablet", 50},
	mobiledetect.FLYTABLET:         Type{"FlyTablet", 51},
	mobiledetect.BQTABLET:          Type{"bqTablet", 52},
	mobiledetect.HUAWEITABLET:      Type{"HuaweiTablet", 53},
	mobiledetect.NECTABLET:         Type{"NecTablet", 54},
	mobiledetect.PANTECHTABLET:     Type{"PantechTablet", 55},
	mobiledetect.BRONCHOTABLET:     Type{"BronchoTablet", 56},
	mobiledetect.VERSUSTABLET:      Type{"VersusTablet", 57},
	mobiledetect.ZYNCTABLET:        Type{"ZyncTablet", 58},
	mobiledetect.POSITIVOTABLET:    Type{"PositivoTablet", 59},
	mobiledetect.NABITABLET:        Type{"NabiTablet", 60},
	mobiledetect.KOBOTABLET:        Type{"KoboTablet", 61},
	mobiledetect.DANEWTABLET:       Type{"DanewTablet", 62},
	mobiledetect.TEXETTABLET:       Type{"TexetTablet", 63},
	mobiledetect.PLAYSTATIONTABLET: Type{"PlaystationTablet", 65},
	mobiledetect.TREKSTORTABLET:    Type{"TrekstorTablet", 66},
	mobiledetect.ADVANTABLET:       Type{"AdvanTablet", 67},
	mobiledetect.DANYTECHTABLET:    Type{"DanyTechTablet", 68},
	mobiledetect.GALAPADTABLET:     Type{"GalapadTablet", 69},
	mobiledetect.MICROMAXTABLET:    Type{"MicromaxTablet", 70},
	mobiledetect.KARBONNTABLET:     Type{"KarbonnTablet", 71},
	mobiledetect.ALLFINETABLET:     Type{"AllFineTablet", 72},
	mobiledetect.PROSCANTABLET:     Type{"PROSCANTablet", 73},
	mobiledetect.YONESTABLET:       Type{"YONESTablet", 74},
	mobiledetect.CHANGJIATABLET:    Type{"ChangJiaTablet", 75},
	mobiledetect.GUTABLET:          Type{"GUTablet", 76},
	mobiledetect.POINTOFVIEWTABLET: Type{"PointOfViewTablet", 77},
	mobiledetect.OVERMAXTABLET:     Type{"OvermaxTablet", 78},
	mobiledetect.DPSTABLET:         Type{"DPSTablet", 80},
	mobiledetect.MODECOMTABLET:     Type{"ModecomTablet", 86},
	mobiledetect.VONINOTABLET:      Type{"VoninoTablet", 87},
	mobiledetect.TOLINOTABLET:      Type{"TolinoTablet", 88},
	mobiledetect.HUDL:              Type{"Hudl", 89},
	mobiledetect.TELSTRATABLET:     Type{"TelstraTablet", 90},
	mobiledetect.GENERICTABLET:     Type{"GenericTablet", 91},
	mobiledetect.XOROTABLET:        Type{"XoroTablet", 92},
	mobiledetect.FX2TABLET:         Type{"Fx2Tablet", 93},
	mobiledetect.IJOYTABLET:        Type{"IJoytTablet", 94},
	mobiledetect.JXDTABLET:         Type{"JXDTablet", 95},
	mobiledetect.TECNOTABLET:       Type{"TecnoTablet", 96},
	mobiledetect.SKKTABLET:         Type{"SkkTablet", 97},
	mobiledetect.AMPETABLET:        Type{"AmpeTablet", 98},
	mobiledetect.AUDIOSONICTABLET:  Type{"AudioSonicTablet", 99},
	mobiledetect.IMOBILETABLET:     Type{"ImobileTablet", 100},
	mobiledetect.ROSSMOORTABLET:    Type{"RossmoorTablet", 101},
	mobiledetect.ESSENTIELBTABLET:  Type{"EssentielbTablet", 102},
	mobiledetect.VODAFONETABLET:    Type{"VodafoneTablet", 103},
	mobiledetect.STOREXTABLET:      Type{"StorexTablet", 104},
	mobiledetect.ECSTABLET:         Type{"EcsTablet", 105},
	mobiledetect.GOCLEVERTABLET:    Type{"GocleverTablet", 106},
	mobiledetect.CONCORDETABLET:    Type{"ConcordeTablet", 107},
	mobiledetect.MEDIATEKTABLET:    Type{"MediatekTablet", 108},
	mobiledetect.CRESTATABLET:      Type{"CrestaTablet", 109},
	mobiledetect.HCLTABLET:         Type{"HclTablet", 110},
	mobiledetect.PYLEAUDIOTABLET:   Type{"PyleaudioTablet", 111},
	mobiledetect.PHILIPSTABLET:     Type{"PhilipsTablet", 112},
	mobiledetect.MSITABLET:         Type{"MsiTablet", 113},
	mobiledetect.VISTURETABLET:     Type{"VistureTablet", 114},
	mobiledetect.VIEWSONICTABLET:   Type{"ViewSonicTablet", 115},
	mobiledetect.ODYSTABLET:        Type{"OdysTablet", 116},
	mobiledetect.CAPTIVATABLET:     Type{"CaptivaTablet", 117},
	mobiledetect.ICONBITTABLET:     Type{"IconBitTablet", 118},
	mobiledetect.TECLASTTABLET:     Type{"TeclastTablet", 119},
	mobiledetect.JAYTECHTABLET:     Type{"JaytechTablet", 120},
	mobiledetect.BLAUPUNKTTABLET:   Type{"BlauPunktTablet", 121},
	mobiledetect.DIGMATABLET:       Type{"DigmatTablet", 122},
	mobiledetect.EVOLIOTABLET:      Type{"EvolioTablet", 123},
	mobiledetect.POCKETBOOKTABLET:  Type{"PocketbookTablet", 124},
}

var mapbrows = map[int]Type{
	mobiledetect.CHROME:         Type{"Chrome", 1},
	mobiledetect.DOLFIN:         Type{"Dolfin", 2},
	mobiledetect.OPERA:          Type{"Opera", 3},
	mobiledetect.SKYFIRE:        Type{"Skyfire", 4},
	mobiledetect.IE:             Type{"IE", 5},
	mobiledetect.FIREFOX:        Type{"Firefox", 6},
	mobiledetect.BOLT:           Type{"Bolt", 7},
	mobiledetect.TEASHARK:       Type{"TeaShark", 8},
	mobiledetect.BLAZER:         Type{"Blazer", 9},
	mobiledetect.SAFARI:         Type{"Safari", 10},
	mobiledetect.TIZEN:          Type{"Tizen", 11},
	mobiledetect.UCBROWSER:      Type{"UCBrowser", 12},
	mobiledetect.DIIGOBROWSER:   Type{"DiigoBrowser", 13},
	mobiledetect.PUFFIN:         Type{"Puffin", 14},
	mobiledetect.MERCURY:        Type{"Mercury", 15},
	mobiledetect.GENERICBROWSER: Type{"GenericBrowser", 16},
	mobiledetect.NETFRONT:       Type{"NetFront", 17},
	mobiledetect.OBIGOBROWSER:   Type{"Obigo", 18},
	mobiledetect.BAIDUBROWSER:   Type{"Baidu", 19},
	mobiledetect.BAIDUBOXAPP:    Type{"Baidubox", 20},
}
