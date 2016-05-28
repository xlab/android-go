package egl

import "errors"

func Error() error {
	switch GetError() {
	case Success:
		return nil
	case NotInitialized:
		return ErrNotInitialized
	case BadAccess:
		return ErrBadAccess
	case BadAlloc:
		return ErrBadAlloc
	case BadAttribute:
		return ErrBadAttribute
	case BadConfig:
		return ErrBadConfig
	case BadContext:
		return ErrBadContext
	case BadCurrentSurface:
		return ErrBadCurrentSurface
	case BadDisplay:
		return ErrBadDisplay
	case BadMatch:
		return ErrBadMatch
	case BadNativePixmap:
		return ErrBadNativePixmap
	case BadNativeWindow:
		return ErrBadNativeWindow
	case BadParameter:
		return ErrBadParameter
	case BadSurface:
		return ErrBadSurface
	case ContextLost:
		return ErrContextLost
	default:
		return ErrUnknown
	}
}

var (
	ErrUnknown           = errors.New("unknown error")
	ErrNotInitialized    = errors.New("not initialized")
	ErrBadAccess         = errors.New("bad access")
	ErrBadAlloc          = errors.New("bad alloc")
	ErrBadAttribute      = errors.New("bad attribute")
	ErrBadConfig         = errors.New("bad config")
	ErrBadContext        = errors.New("bad context")
	ErrBadCurrentSurface = errors.New("bad current surface")
	ErrBadDisplay        = errors.New("bad display")
	ErrBadMatch          = errors.New("bad match")
	ErrBadNativePixmap   = errors.New("bad native pixmap")
	ErrBadNativeWindow   = errors.New("bad native window")
	ErrBadParameter      = errors.New("bad parameter")
	ErrBadSurface        = errors.New("bad surface")
	ErrContextLost       = errors.New("context lost")
)
