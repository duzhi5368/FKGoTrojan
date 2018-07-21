/*
Author: FreeKnight
虚拟键值/windows值枚举
*/
//------------------------------------------------------------
package components
//------------------------------------------------------------
const (
	MB_OK                	= 0x00000000
	MB_OKCANCEL          	= 0x00000001
	MB_ABORTRETRYIGNORE  	= 0x00000002
	MB_YESNOCANCEL       	= 0x00000003
	MB_YESNO             	= 0x00000004
	MB_RETRYCANCEL       	= 0x00000005
	MB_CANCELTRYCONTINUE 	= 0x00000006
	MB_ICONHAND          	= 0x00000010
	MB_ICONQUESTION      	= 0x00000020
	MB_ICONEXCLAMATION   	= 0x00000030
	MB_ICONASTERISK      	= 0x00000040
	MB_USERICON          	= 0x00000080
	MB_ICONWARNING       	= MB_ICONEXCLAMATION
	MB_ICONERROR         	= MB_ICONHAND
	MB_ICONINFORMATION   	= MB_ICONASTERISK
	MB_ICONSTOP          	= MB_ICONHAND

	MB_DEFBUTTON1 		= 0x00000000
	MB_DEFBUTTON2 		= 0x00000100
	MB_DEFBUTTON3 		= 0x00000200
	MB_DEFBUTTON4 		= 0x00000300

	ERROR_ALREADY_EXISTS 	= 183

	MEM_COMMIT  		= 0x1000
	MEM_RESERVE 		= 0x2000

	PAGE_EXECUTE_READWRITE 	= 0x40

	// 虚拟键
	vk_BACK       = 0x08
	vk_TAB        = 0x09
	vk_CLEAR      = 0x0C
	vk_RETURN     = 0x0D
	vk_SHIFT      = 0x10
	vk_CONTROL    = 0x11
	vk_MENU       = 0x12
	vk_PAUSE      = 0x13
	vk_CAPITAL    = 0x14
	vk_ESCAPE     = 0x1B
	vk_SPACE      = 0x20
	vk_PRIOR      = 0x21
	vk_NEXT       = 0x22
	vk_END        = 0x23
	vk_HOME       = 0x24
	vk_LEFT       = 0x25
	vk_UP         = 0x26
	vk_RIGHT      = 0x27
	vk_DOWN       = 0x28
	vk_SELECT     = 0x29
	vk_PRINT      = 0x2A
	vk_EXECUTE    = 0x2B
	vk_SNAPSHOT   = 0x2C
	vk_INSERT     = 0x2D
	vk_DELETE     = 0x2E
	vk_LWIN       = 0x5B
	vk_RWIN       = 0x5C
	vk_APPS       = 0x5D
	vk_SLEEP      = 0x5F
	vk_NUMPAD0    = 0x60
	vk_NUMPAD1    = 0x61
	vk_NUMPAD2    = 0x62
	vk_NUMPAD3    = 0x63
	vk_NUMPAD4    = 0x64
	vk_NUMPAD5    = 0x65
	vk_NUMPAD6    = 0x66
	vk_NUMPAD7    = 0x67
	vk_NUMPAD8    = 0x68
	vk_NUMPAD9    = 0x69
	vk_MULTIPLY   = 0x6A
	vk_ADD        = 0x6B
	vk_SEPARATOR  = 0x6C
	vk_SUBTRACT   = 0x6D
	vk_DECIMAL    = 0x6E
	vk_DIVIDE     = 0x6F
	vk_F1         = 0x70
	vk_F2         = 0x71
	vk_F3         = 0x72
	vk_F4         = 0x73
	vk_F5         = 0x74
	vk_F6         = 0x75
	vk_F7         = 0x76
	vk_F8         = 0x77
	vk_F9         = 0x78
	vk_F10        = 0x79
	vk_F11        = 0x7A
	vk_F12        = 0x7B
	vk_NUMLOCK    = 0x90
	vk_SCROLL     = 0x91
	vk_LSHIFT     = 0xA0
	vk_RSHIFT     = 0xA1
	vk_LCONTROL   = 0xA2
	vk_RCONTROL   = 0xA3
	vk_LMENU      = 0xA4
	vk_RMENU      = 0xA5
	vk_OEM_1      = 0xBA
	vk_OEM_PLUS   = 0xBB
	vk_OEM_COMMA  = 0xBC
	vk_OEM_MINUS  = 0xBD
	vk_OEM_PERIOD = 0xBE
	vk_OEM_2      = 0xBF
	vk_OEM_3      = 0xC0
	vk_OEM_4      = 0xDB
	vk_OEM_5      = 0xDC
	vk_OEM_6      = 0xDD
	vk_OEM_7      = 0xDE
	vk_OEM_8      = 0xDF
)