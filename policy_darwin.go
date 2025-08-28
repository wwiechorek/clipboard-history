//go:build darwin

package main

/*
#cgo CFLAGS: -x objective-c -fmodules -fobjc-arc
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#import <dispatch/dispatch.h>

void setAccessoryPolicy(void) {
    @autoreleasepool {
        dispatch_async(dispatch_get_main_queue(), ^{
            NSApplication *app = [NSApplication sharedApplication];
            [app setActivationPolicy:NSApplicationActivationPolicyAccessory];
        });
    }
}
*/
import "C"

// applyAccessoryPolicy sets NSApplicationActivationPolicyAccessory on macOS.
func applyAccessoryPolicy() {
    C.setAccessoryPolicy()
}

