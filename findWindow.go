package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var user32 = syscall.NewLazyDLL("user32.dll")
var procFindWindowW = user32.NewProc("FindWindowW")
var procShowWindow = user32.NewProc("ShowWindow")
var procSetForeground = user32.NewProc("SetForegroundWindow")

// FindWindow searches for the window by its title and returns a handle.
func FindWindow(title string) (uintptr, error) {
	// Convert title string to UTF-16
	titlePtr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return 0, err
	}

	// Call FindWindowW with the window title
	ret, _, _ := procFindWindowW.Call(uintptr(unsafe.Pointer(titlePtr)), 0)
	return ret, nil
}

const (
	SW_RESTORE = 9
)

func restoreAndFocusWindow(hwnd uintptr) error {
	// Attempt to restore the window
	_, _, err := procShowWindow.Call(hwnd, SW_RESTORE)
	if err != nil && err.Error() != "The operation completed successfully." {
		return fmt.Errorf("failed to restore window: %v", err)
	}

	// Set the window to the foreground
	_, _, err = procSetForeground.Call(hwnd)
	if err != nil && err.Error() != "The operation completed successfully." {
		return fmt.Errorf("failed to set window to foreground: %v", err)
	}

	return nil
}
