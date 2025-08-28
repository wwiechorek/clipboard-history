//go:build darwin

package main

/*
#cgo CFLAGS: -x objective-c -fmodules -fobjc-arc
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#import <dispatch/dispatch.h>

static NSScreen* screenContainingPoint(NSPoint point) {
    for (NSScreen *screen in [NSScreen screens]) {
        NSRect frame = [screen visibleFrame];
        if (NSPointInRect(point, frame)) {
            return screen;
        }
    }
    return [NSScreen mainScreen];
}

void centerMainWindowOnMouseScreen(void) {
    @autoreleasepool {
        dispatch_async(dispatch_get_main_queue(), ^{
            NSPoint mouse = [NSEvent mouseLocation];
            NSScreen *target = screenContainingPoint(mouse);
            if (target == nil) {
                target = [NSScreen mainScreen];
            }

            NSWindow *win = [NSApp mainWindow];
            if (win == nil) {
                win = [NSApp keyWindow];
            }
            if (win == nil) {
                return; // No window to move
            }

            NSRect vis = [target visibleFrame];
            NSRect wf = [win frame];

            CGFloat newX = vis.origin.x + (vis.size.width - wf.size.width) / 2.0;
            CGFloat newY = vis.origin.y + (vis.size.height - wf.size.height) / 2.0;

            [win setFrameOrigin:NSMakePoint(newX, newY)];
        });
    }
}
*/
import "C"

import "context"

// CenterWindowOnMouseMonitor recenters the main window on the monitor
// currently under the mouse pointer (macOS only).
func CenterWindowOnMouseMonitor(_ context.Context) {
    C.centerMainWindowOnMouseScreen()
}

