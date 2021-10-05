package main

import (
	evdev "github.com/gvalkov/golang-evdev"
)

var ModKeyToMask = map[uint8]uint8{
	evdev.KEY_RIGHTMETA:  1 << 7,
	evdev.KEY_RIGHTALT:   1 << 6,
	evdev.KEY_RIGHTSHIFT: 1 << 5,
	evdev.KEY_RIGHTCTRL:  1 << 4,
	evdev.KEY_LEFTMETA:   1 << 3,
	evdev.KEY_LEFTALT:    1 << 2,
	evdev.KEY_LEFTSHIFT:  1 << 1,
	evdev.KEY_LEFTCTRL:   1 << 0,
}

var ScanCodeToKeyCode = map[uint8]uint8{
	evdev.KEY_A:          0x04,
	evdev.KEY_B:          0x05,
	evdev.KEY_C:          0x06,
	evdev.KEY_D:          0x07,
	evdev.KEY_E:          0x08,
	evdev.KEY_F:          0x09,
	evdev.KEY_G:          0x0a,
	evdev.KEY_H:          0x0b,
	evdev.KEY_I:          0x0c,
	evdev.KEY_J:          0x0d,
	evdev.KEY_K:          0x0e,
	evdev.KEY_L:          0x0f,
	evdev.KEY_M:          0x10,
	evdev.KEY_N:          0x11,
	evdev.KEY_O:          0x12,
	evdev.KEY_P:          0x13,
	evdev.KEY_Q:          0x14,
	evdev.KEY_R:          0x15,
	evdev.KEY_S:          0x16,
	evdev.KEY_T:          0x17,
	evdev.KEY_U:          0x18,
	evdev.KEY_V:          0x19,
	evdev.KEY_W:          0x1a,
	evdev.KEY_X:          0x1b,
	evdev.KEY_Y:          0x1c,
	evdev.KEY_Z:          0x1d,
	evdev.KEY_1:          0x1e,
	evdev.KEY_2:          0x1f,
	evdev.KEY_3:          0x20,
	evdev.KEY_4:          0x21,
	evdev.KEY_5:          0x22,
	evdev.KEY_6:          0x23,
	evdev.KEY_7:          0x24,
	evdev.KEY_8:          0x25,
	evdev.KEY_9:          0x26,
	evdev.KEY_0:          0x27,
	evdev.KEY_ENTER:      0x28,
	evdev.KEY_ESC:        0x29,
	evdev.KEY_BACKSPACE:  0x2a,
	evdev.KEY_TAB:        0x2b,
	evdev.KEY_SPACE:      0x2c,
	evdev.KEY_MINUS:      0x2d,
	evdev.KEY_EQUAL:      0x2e,
	evdev.KEY_LEFTBRACE:  0x2f,
	evdev.KEY_RIGHTBRACE: 0x30,
	evdev.KEY_BACKSLASH:  0x31,
	// evdev.KEY_BACKSLASH:  0x32, (Keyboard Non-US # and ~)
	evdev.KEY_SEMICOLON:  0x33,
	evdev.KEY_APOSTROPHE: 0x34,
	evdev.KEY_GRAVE:      0x35,
	evdev.KEY_COMMA:      0x36,
	evdev.KEY_DOT:        0x37,
	evdev.KEY_SLASH:      0x38,
	evdev.KEY_CAPSLOCK:   0x39,
	evdev.KEY_F1:         0x3a,
	evdev.KEY_F2:         0x3b,
	evdev.KEY_F3:         0x3c,
	evdev.KEY_F4:         0x3d,
	evdev.KEY_F5:         0x3e,
	evdev.KEY_F6:         0x3f,
	evdev.KEY_F7:         0x40,
	evdev.KEY_F8:         0x41,
	evdev.KEY_F9:         0x42,
	evdev.KEY_F10:        0x43,
	evdev.KEY_F11:        0x44,
	evdev.KEY_F12:        0x45,
	evdev.KEY_SYSRQ:      0x46,
	evdev.KEY_SCROLLLOCK: 0x47,
	evdev.KEY_PAUSE:      0x48,
	evdev.KEY_INSERT:     0x49,
	evdev.KEY_HOME:       0x4a,
	evdev.KEY_PAGEUP:     0x4b,
	evdev.KEY_DELETE:     0x4c,
	evdev.KEY_END:        0x4d,
	evdev.KEY_PAGEDOWN:   0x4e,
	evdev.KEY_RIGHT:      0x4f,
	evdev.KEY_LEFT:       0x50,
	evdev.KEY_DOWN:       0x51,
	evdev.KEY_UP:         0x52,
	evdev.KEY_NUMLOCK:    0x53,
	evdev.KEY_KPSLASH:    0x54,
	evdev.KEY_KPASTERISK: 0x55,
	evdev.KEY_KPMINUS:    0x56,
	evdev.KEY_KPPLUS:     0x57,
	evdev.KEY_KPENTER:    0x58,
	evdev.KEY_KP1:        0x59,
	evdev.KEY_KP2:        0x5a,
	evdev.KEY_KP3:        0x5b,
	evdev.KEY_KP4:        0x5c,
	evdev.KEY_KP5:        0x5d,
	evdev.KEY_KP6:        0x5e,
	evdev.KEY_KP7:        0x5f,
	evdev.KEY_KP8:        0x60,
	evdev.KEY_KP9:        0x61,
	evdev.KEY_KP0:        0x62,
	evdev.KEY_KPDOT:      0x63,
	evdev.KEY_102ND:      0x64,
	evdev.KEY_COMPOSE:    0x65,
	evdev.KEY_POWER:      0x66,
	evdev.KEY_KPEQUAL:    0x67,
	evdev.KEY_F13:        0x68,
	evdev.KEY_F14:        0x69,
	evdev.KEY_F15:        0x6a,
	evdev.KEY_F16:        0x6b,
	evdev.KEY_F17:        0x6c,
	evdev.KEY_F18:        0x6d,
	evdev.KEY_F19:        0x6e,
	evdev.KEY_F20:        0x6f,
	evdev.KEY_F21:        0x70,
	evdev.KEY_F22:        0x71,
	evdev.KEY_F23:        0x72,
	evdev.KEY_F24:        0x73,
	evdev.KEY_OPEN:       0x74,
	evdev.KEY_HELP:       0x75,
	evdev.KEY_PROPS:      0x76,
	evdev.KEY_FRONT:      0x77,
	evdev.KEY_STOP:       0x78,
	evdev.KEY_AGAIN:      0x79,
	evdev.KEY_UNDO:       0x7a,
	evdev.KEY_CUT:        0x7b,
	evdev.KEY_COPY:       0x7c,
	evdev.KEY_PASTE:      0x7d,
	evdev.KEY_FIND:       0x7e,
	evdev.KEY_MUTE:       0x7f,
	evdev.KEY_VOLUMEUP:   0x80,
	evdev.KEY_VOLUMEDOWN: 0x81,
	// Keyboard Locking Caps Lock: 0x82,
	// Keyboard Locking Num Lock: 0x83,
	// Keyboard Locking Scroll Lock: 0x84,
	evdev.KEY_KPCOMMA: 0x85,
	// evdev.KEY_KPEQUAL:          0x86,
	evdev.KEY_RO:               0x87,
	evdev.KEY_KATAKANAHIRAGANA: 0x88,
	evdev.KEY_YEN:              0x89,
	evdev.KEY_HENKAN:           0x8a,
	evdev.KEY_MUHENKAN:         0x8b,
	evdev.KEY_KPJPCOMMA:        0x8c,
	// Keyboard International7: 0x8d,
	// Keyboard International8: 0x8e,
	// Keyboard International9: 0x8f,
	evdev.KEY_HANGEUL:        0x90,
	evdev.KEY_HANJA:          0x91,
	evdev.KEY_KATAKANA:       0x92,
	evdev.KEY_HIRAGANA:       0x93,
	evdev.KEY_ZENKAKUHANKAKU: 0x94,
	// Keyboard LANG6: 0x95,
	// Keyboard LANG6: 0x96,
	// Keyboard LANG6: 0x97,
	// Keyboard LANG6: 0x98,
	// Keyboard Alternate Erase: 0x99,
	// Keyboard SysReq/Attention: 0x9a,
	// Keyboard Cancel: 0x9b,
	// Keyboard Clear: 0x9c,
	// Keyboard Prior: 0x9d,
	// Keyboard Return: 0x9e,
	// Keyboard Separator: 0x9f,
	// Keyboard Out: 0xa0,
	// Keyboard Oper: 0xa1,
	// Keyboard Clear/Again: 0xa2,
	// Keyboard CrSel/Props: 0xa3,
	// Keyboard ExSel: 0xa4,
	// Reserved: 0xa5 to 0xaf
	// Keypad 00: 0xb0,
	// Keypad 000: 0xb1,
	// Thousands Separator: 0xb2,
	// Decimal Separator: 0xb3,
	// Currency Unit: 0xb4,
	// Currency Sub-unit: 0xb5,
	evdev.KEY_KPLEFTPAREN:  0xb6,
	evdev.KEY_KPRIGHTPAREN: 0xb7,
	// Keypad {: 0xb8
	// Keypad }: 0xb9
	// Keypad Tab: 0xba
	// Keypad Backspace: 0xbb
	// Keypad A: 0xbc
	// Keypad B: 0xbd
	// Keypad C: 0xbe
	// Keypad D: 0xbf
	// Keypad E: 0xc0
	// Keypad F: 0xc1
	// Keypad XOR: 0xc2
	// Keypad ^: 0xc3
	// Keypad %: 0xc4
	// Keypad <: 0xc5
	// Keypad >: 0xc6
	// Keypad &: 0xc7
	// Keypad &&: 0xc8
	// Keypad |: 0xc9
	// Keypad ||: 0xca
	// Keypad :: 0xcb
	// Keypad #: 0xcc
	// Keypad Space: 0xcd
	// Keypad @: 0xce
	// Keypad !: 0xcf
	// Keypad Memory Store: 0xd0
	// Keypad Memory Recall: 0xd1
	// Keypad Memory Clear: 0xd2
	// Keypad Memory Add: 0xd3
	// Keypad Memory Subtract: 0xd4
	// Keypad Memory Multiply: 0xd5
	// Keypad Memory Divide: 0xd6
	// Keypad +/-: 0xd7
	// Keypad Clear: 0xd8
	// Keypad Clear Entry: 0xd9
	// Keypad Binary: 0xda
	// Keypad Octal: 0xdb
	// Keypad Decimal: 0xdc
	// Keypad Hexadecimal: 0xdd
	evdev.KEY_LEFTCTRL:     0xe0,
	evdev.KEY_LEFTSHIFT:    0xe1,
	evdev.KEY_LEFTALT:      0xe2,
	evdev.KEY_LEFTMETA:     0xe3,
	evdev.KEY_RIGHTCTRL:    0xe4,
	evdev.KEY_RIGHTSHIFT:   0xe5,
	evdev.KEY_RIGHTALT:     0xe6,
	evdev.KEY_RIGHTMETA:    0xe7,
	evdev.KEY_PLAYPAUSE:    0xe8,
	evdev.KEY_STOPCD:       0xe9,
	evdev.KEY_PREVIOUSSONG: 0xea,
	evdev.KEY_NEXTSONG:     0xeb,
	evdev.KEY_EJECTCD:      0xec,
	// evdev.KEY_VOLUMEUP:     0xed,
	// evdev.KEY_VOLUMEDOWN:   0xee,
	// evdev.KEY_MUTE:         0xef,
	evdev.KEY_WWW:     0xf0,
	evdev.KEY_BACK:    0xf1,
	evdev.KEY_FORWARD: 0xf2,
	// evdev.KEY_STOP:         0xf3,
	// evdev.KEY_FIND:         0xf4,
	evdev.KEY_SCROLLUP:   0xf5,
	evdev.KEY_SCROLLDOWN: 0xf6,
	evdev.KEY_EDIT:       0xf7,
	evdev.KEY_SLEEP:      0xf8,
	evdev.KEY_COFFEE:     0xf9,
	evdev.KEY_REFRESH:    0xfa,
	evdev.KEY_CALC:       0xfb,
}
