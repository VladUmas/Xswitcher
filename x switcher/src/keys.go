package main
// Key (name=code) definitions. To simplifize config.
// https://github.com/gvalkov/golang-evdev/blob/master/ecodes.go
// awk '{if (($1 ~ "^KEY_") && ($3 ~ "^[0-9]")) print "\t\"" $1 "\": " $3 ","}' ecodes.go.src | sed 's/KEY_//'
var (
	key_def = map[string]uint { // evdev way is too sophisticated for my goals...
	"RESERVED": 0,
	"ESC": 1,
	"1": 2,
	"2": 3,
	"3": 4,
	"4": 5,
	"5": 6,
	"6": 7,
	"7": 8,
	"8": 9,
	"9": 10,
	"0": 11,
	"MINUS": 12,
	"-" : 12,
	"EQUAL": 13,
	"=": 13,
	"BACKSPACE": 14,
	"BS": 14,
	"TAB": 15,
	"Q": 16,
	"W": 17,
	"E": 18,
	"R": 19,
	"T": 20,
	"Y": 21,
	"U": 22,
	"I": 23,
	"O": 24,
	"P": 25,
	"LEFTBRACE": 26,
	"[": 26,
	"RIGHTBRACE": 27,
	"]": 27,
	"ENTER": 28,
	"LEFTCTRL": 29,
	"L_CTRL": 29,
	"A": 30,
	"S": 31,
	"D": 32,
	"F": 33,
	"G": 34,
	"H": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"SEMICOLON": 39,
	";": 39,
	"APOSTROPHE": 40,
	"'": 40,
	"GRAVE": 41,
	"`": 41,
	"LEFTSHIFT": 42,
	"L_SHIFT": 42,
	"BACKSLASH": 43,
	"\\": 43,
	"Z": 44,
	"X": 45,
	"C": 46,
	"V": 47,
	"B": 48,
	"N": 49,
	"M": 50,
	"COMMA": 51,
	",": 51,
	"DOT": 52,
	".": 52,
	"SLASH": 53,
	"/": 53,
	"RIGHTSHIFT": 54,
	"R_SHIFT": 54,
	"KPASTERISK": 55,
	"LEFTALT": 56,
	"L_ALT": 56,
	"SPACE": 57,
	"_": 57,
	"CAPSLOCK": 58,
	"CAPS": 58,
	"F1": 59,
	"F2": 60,
	"F3": 61,
	"F4": 62,
	"F5": 63,
	"F6": 64,
	"F7": 65,
	"F8": 66,
	"F9": 67,
	"F10": 68,
	"NUMLOCK": 69,
	"N_LOCK": 69,
	"SCROLLLOCK": 70,
	"S_LOCK": 70,
	"KP7": 71,
	"KP8": 72,
	"KP9": 73,
	"KPMINUS": 74,
	"KP4": 75,
	"KP5": 76,
	"KP6": 77,
	"KPPLUS": 78,
	"KP1": 79,
	"KP2": 80,
	"KP3": 81,
	"KP0": 82,
	"KPDOT": 83,
	"ZENKAKUHANKAKU": 85,
	"102ND": 86,
	"F11": 87,
	"F12": 88,
	"RO": 89,
	"KATAKANA": 90,
	"HIRAGANA": 91,
	"HENKAN": 92,
	"KATAKANAHIRAGANA": 93,
	"MUHENKAN": 94,
	"KPJPCOMMA": 95,
	"KPENTER": 96,
	"RIGHTCTRL": 97,
	"R_CTRL": 97,
	"KPSLASH": 98,
	"SYSRQ": 99,
	"RIGHTALT": 100,
	"R_ALT": 100,
	"LINEFEED": 101,
	"HOME": 102,
	"UP": 103,
	"PAGEUP": 104,
	"LEFT": 105,
	"RIGHT": 106,
	"END": 107,
	"DOWN": 108,
	"PAGEDOWN": 109,
	"INSERT": 110,
	"DELETE": 111,
	"MACRO": 112,
	"MUTE": 113,
	"VOLUMEDOWN": 114,
	"VOLUMEUP": 115,
	"POWER": 116,
	"KPEQUAL": 117,
	"KPPLUSMINUS": 118,
	"PAUSE": 119,
	"SCALE": 120,
	"KPCOMMA": 121,
	"HANGEUL": 122,
	"HANJA": 123,
	"YEN": 124,
	"LEFTMETA": 125,
	"L_META": 125,
	"RIGHTMETA": 126,
	"R_META": 126,
	"COMPOSE": 127,
	"STOP": 128,
	"AGAIN": 129,
	"PROPS": 130,
	"UNDO": 131,
	"FRONT": 132,
	"COPY": 133,
	"OPEN": 134,
	"PASTE": 135,
	"FIND": 136,
	"CUT": 137,
	"HELP": 138,
	"MENU": 139,
	"CALC": 140,
	"SETUP": 141,
	"SLEEP": 142,
	"WAKEUP": 143,
	"FILE": 144,
	"SENDFILE": 145,
	"DELETEFILE": 146,
	"XFER": 147,
	"PROG1": 148,
	"PROG2": 149,
	"WWW": 150,
	"MSDOS": 151,
	"COFFEE": 152,
	"ROTATE_DISPLAY": 153,
	"CYCLEWINDOWS": 154,
	"MAIL": 155,
	"BOOKMARKS": 156,
	"COMPUTER": 157,
	"BACK": 158,
	"FORWARD": 159,
	"CLOSECD": 160,
	"EJECTCD": 161,
	"EJECTCLOSECD": 162,
	"NEXTSONG": 163,
	"PLAYPAUSE": 164,
	"PREVIOUSSONG": 165,
	"STOPCD": 166,
	"RECORD": 167,
	"REWIND": 168,
	"PHONE": 169,
	"ISO": 170,
	"CONFIG": 171,
	"HOMEPAGE": 172,
	"REFRESH": 173,
	"EXIT": 174,
	"MOVE": 175,
	"EDIT": 176,
	"SCROLLUP": 177,
	"SCROLLDOWN": 178,
	"KPLEFTPAREN": 179,
	"KPRIGHTPAREN": 180,
	"NEW": 181,
	"REDO": 182,
	"F13": 183,
	"F14": 184,
	"F15": 185,
	"F16": 186,
	"F17": 187,
	"F18": 188,
	"F19": 189,
	"F20": 190,
	"F21": 191,
	"F22": 192,
	"F23": 193,
	"F24": 194,
	"PLAYCD": 200,
	"PAUSECD": 201,
	"PROG3": 202,
	"PROG4": 203,
	"DASHBOARD": 204,
	"SUSPEND": 205,
	"CLOSE": 206,
	"PLAY": 207,
	"FASTFORWARD": 208,
	"BASSBOOST": 209,
	"PRINT": 210,
	"HP": 211,
	"CAMERA": 212,
	"SOUND": 213,
	"QUESTION": 214,
	"EMAIL": 215,
	"CHAT": 216,
	"SEARCH": 217,
	"CONNECT": 218,
	"FINANCE": 219,
	"SPORT": 220,
	"SHOP": 221,
	"ALTERASE": 222,
	"CANCEL": 223,
	"BRIGHTNESSDOWN": 224,
	"BRIGHTNESSUP": 225,
	"MEDIA": 226,
	"SWITCHVIDEOMODE": 227,
	"KBDILLUMTOGGLE": 228,
	"KBDILLUMDOWN": 229,
	"KBDILLUMUP": 230,
	"SEND": 231,
	"REPLY": 232,
	"FORWARDMAIL": 233,
	"SAVE": 234,
	"DOCUMENTS": 235,
	"BATTERY": 236,
	"BLUETOOTH": 237,
	"WLAN": 238,
	"UWB": 239,
	"UNKNOWN": 240,
	"VIDEO_NEXT": 241,
	"VIDEO_PREV": 242,
	"BRIGHTNESS_CYCLE": 243,
	"BRIGHTNESS_AUTO": 244,
	"DISPLAY_OFF": 245,
	"WWAN": 246,
	"RFKILL": 247,
	"MICMUTE": 248,
	"OK": 0x160,
	"SELECT": 0x161,
	"GOTO": 0x162,
	"CLEAR": 0x163,
	"POWER2": 0x164,
	"OPTION": 0x165,
	"INFO": 0x166,
	"TIME": 0x167,
	"VENDOR": 0x168,
	"ARCHIVE": 0x169,
	"PROGRAM": 0x16a,
	"CHANNEL": 0x16b,
	"FAVORITES": 0x16c,
	"EPG": 0x16d,
	"PVR": 0x16e,
	"MHP": 0x16f,
	"LANGUAGE": 0x170,
	"TITLE": 0x171,
	"SUBTITLE": 0x172,
	"ANGLE": 0x173,
	"ZOOM": 0x174,
	"MODE": 0x175,
	"KEYBOARD": 0x176,
	"SCREEN": 0x177,
	"PC": 0x178,
	"TV": 0x179,
	"TV2": 0x17a,
	"VCR": 0x17b,
	"VCR2": 0x17c,
	"SAT": 0x17d,
	"SAT2": 0x17e,
	"CD": 0x17f,
	"TAPE": 0x180,
	"RADIO": 0x181,
	"TUNER": 0x182,
	"PLAYER": 0x183,
	"TEXT": 0x184,
	"DVD": 0x185,
	"AUX": 0x186,
	"MP3": 0x187,
	"AUDIO": 0x188,
	"VIDEO": 0x189,
	"DIRECTORY": 0x18a,
	"LIST": 0x18b,
	"MEMO": 0x18c,
	"CALENDAR": 0x18d,
	"RED": 0x18e,
	"GREEN": 0x18f,
	"YELLOW": 0x190,
	"BLUE": 0x191,
	"CHANNELUP": 0x192,
	"CHANNELDOWN": 0x193,
	"FIRST": 0x194,
	"LAST": 0x195,
	"AB": 0x196,
	"NEXT": 0x197,
	"RESTART": 0x198,
	"SLOW": 0x199,
	"SHUFFLE": 0x19a,
	"BREAK": 0x19b,
	"PREVIOUS": 0x19c,
	"DIGITS": 0x19d,
	"TEEN": 0x19e,
	"TWEN": 0x19f,
	"VIDEOPHONE": 0x1a0,
	"GAMES": 0x1a1,
	"ZOOMIN": 0x1a2,
	"ZOOMOUT": 0x1a3,
	"ZOOMRESET": 0x1a4,
	"WORDPROCESSOR": 0x1a5,
	"EDITOR": 0x1a6,
	"SPREADSHEET": 0x1a7,
	"GRAPHICSEDITOR": 0x1a8,
	"PRESENTATION": 0x1a9,
	"DATABASE": 0x1aa,
	"NEWS": 0x1ab,
	"VOICEMAIL": 0x1ac,
	"ADDRESSBOOK": 0x1ad,
	"MESSENGER": 0x1ae,
	"DISPLAYTOGGLE": 0x1af,
	"SPELLCHECK": 0x1b0,
	"LOGOFF": 0x1b1,
	"DOLLAR": 0x1b2,
	"EURO": 0x1b3,
	"FRAMEBACK": 0x1b4,
	"FRAMEFORWARD": 0x1b5,
	"CONTEXT_MENU": 0x1b6,
	"MEDIA_REPEAT": 0x1b7,
	"10CHANNELSUP": 0x1b8,
	"10CHANNELSDOWN": 0x1b9,
	"IMAGES": 0x1ba,
	"DEL_EOL": 0x1c0,
	"DEL_EOS": 0x1c1,
	"INS_LINE": 0x1c2,
	"DEL_LINE": 0x1c3,
	"FN": 0x1d0,
	"FN_ESC": 0x1d1,
	"FN_F1": 0x1d2,
	"FN_F2": 0x1d3,
	"FN_F3": 0x1d4,
	"FN_F4": 0x1d5,
	"FN_F5": 0x1d6,
	"FN_F6": 0x1d7,
	"FN_F7": 0x1d8,
	"FN_F8": 0x1d9,
	"FN_F9": 0x1da,
	"FN_F10": 0x1db,
	"FN_F11": 0x1dc,
	"FN_F12": 0x1dd,
	"FN_1": 0x1de,
	"FN_2": 0x1df,
	"FN_D": 0x1e0,
	"FN_E": 0x1e1,
	"FN_F": 0x1e2,
	"FN_S": 0x1e3,
	"FN_B": 0x1e4,
	"BRL_DOT1": 0x1f1,
	"BRL_DOT2": 0x1f2,
	"BRL_DOT3": 0x1f3,
	"BRL_DOT4": 0x1f4,
	"BRL_DOT5": 0x1f5,
	"BRL_DOT6": 0x1f6,
	"BRL_DOT7": 0x1f7,
	"BRL_DOT8": 0x1f8,
	"BRL_DOT9": 0x1f9,
	"BRL_DOT10": 0x1fa,
	"NUMERIC_0": 0x200,
	"NUMERIC_1": 0x201,
	"NUMERIC_2": 0x202,
	"NUMERIC_3": 0x203,
	"NUMERIC_4": 0x204,
	"NUMERIC_5": 0x205,
	"NUMERIC_6": 0x206,
	"NUMERIC_7": 0x207,
	"NUMERIC_8": 0x208,
	"NUMERIC_9": 0x209,
	"NUMERIC_STAR": 0x20a,
	"NUMERIC_POUND": 0x20b,
	"NUMERIC_A": 0x20c,
	"NUMERIC_B": 0x20d,
	"NUMERIC_C": 0x20e,
	"NUMERIC_D": 0x20f,
	"CAMERA_FOCUS": 0x210,
	"WPS_BUTTON": 0x211,
	"TOUCHPAD_TOGGLE": 0x212,
	"TOUCHPAD_ON": 0x213,
	"TOUCHPAD_OFF": 0x214,
	"CAMERA_ZOOMIN": 0x215,
	"CAMERA_ZOOMOUT": 0x216,
	"CAMERA_UP": 0x217,
	"CAMERA_DOWN": 0x218,
	"CAMERA_LEFT": 0x219,
	"CAMERA_RIGHT": 0x21a,
	"ATTENDANT_ON": 0x21b,
	"ATTENDANT_OFF": 0x21c,
	"ATTENDANT_TOGGLE": 0x21d,
	"LIGHTS_TOGGLE": 0x21e,
	"ALS_TOGGLE": 0x230,
	"BUTTONCONFIG": 0x240,
	"TASKMANAGER": 0x241,
	"JOURNAL": 0x242,
	"CONTROLPANEL": 0x243,
	"APPSELECT": 0x244,
	"SCREENSAVER": 0x245,
	"VOICECOMMAND": 0x246,
	"BRIGHTNESS_MIN": 0x250,
	"BRIGHTNESS_MAX": 0x251,
	"KBDINPUTASSIST_PREV": 0x260,
	"KBDINPUTASSIST_NEXT": 0x261,
	"KBDINPUTASSIST_PREVGROUP": 0x262,
	"KBDINPUTASSIST_NEXTGROUP": 0x263,
	"KBDINPUTASSIST_ACCEPT": 0x264,
	"KBDINPUTASSIST_CANCEL": 0x265,
	"RIGHT_UP": 0x266,
	"RIGHT_DOWN": 0x267,
	"LEFT_UP": 0x268,
	"LEFT_DOWN": 0x269,
	"ROOT_MENU": 0x26a,
	"MEDIA_TOP_MENU": 0x26b,
	"NUMERIC_11": 0x26c,
	"NUMERIC_12": 0x26d,
	"AUDIO_DESC": 0x26e,
	"3D_MODE": 0x26f,
	"NEXT_FAVORITE": 0x270,
	"STOP_RECORD": 0x271,
	"PAUSE_RECORD": 0x272,
	"VOD": 0x273,
	"UNMUTE": 0x274,
	"FASTREVERSE": 0x275,
	"SLOWREVERSE": 0x276,
	"DATA": 0x275,
	"MAX": 0x2ff,
	} // 768 elements 0..767
)
