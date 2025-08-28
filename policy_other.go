//go:build !darwin

package main

// applyAccessoryPolicy is a no-op on non-macOS platforms.
func applyAccessoryPolicy() {}

