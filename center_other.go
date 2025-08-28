//go:build !darwin

package main

import (
    "context"
    wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Fallback: just center on current screen for non-macOS.
func CenterWindowOnMouseMonitor(ctx context.Context) {
    wailsRuntime.WindowCenter(ctx)
}

