package main

import (
	"io"
	"runtime"
	"sync"

	"github.com/xlab/android-go/app"
)

const assetSize = 1 << 20

// Read the same asset concurrently with one reader per go routines and check the returned values.
var testAssetsConcurrent = test{
	name: "assetsConcurrent",
	run: func(t test, a app.NativeActivity) {
		var wg sync.WaitGroup
		// Dispatch goroutines
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(i int) {
				// Open AssetReader
				var r, err = a.OpenAsset("testAsset")
				t.failIf(err != nil, "Can't open asset:", err)
				// Continue reading until there is a read error
				var b [32]byte
				for count := 0; count < assetSize; count += len(b) {
					// Read 32 bytes
					_, err = io.ReadFull(r, b[:])
					// Check for read errors
					t.failIf(err != nil, "Read error:", err)
					// Check content
					for j := range b {
						t.failIf(b[j] != byte(count+j), "Wrong content")
					}
				}
				// Read whole asset. Check for io.EOF
				_, err = r.Read(b[:])
				t.failIf(err != io.EOF, "Expected io.EOF, got:", err)
				// Tell parent goroutine that we are done
				wg.Done()
			}(i)
		}
		// Wait for child goroutines to finish
		wg.Wait()
	},
}

// Use the same asset reader alternatingly on two goroutines and check the returned values.
var testAssetsShared = test{
	name: "assetsShared",
	run: func(t test, a app.NativeActivity) {
		var wg sync.WaitGroup
		// Open shared AssetReader
		var r, err = a.OpenAsset("testAsset")
		t.failIf(err != nil, "Can't open asset:", err)
		// Channel for signalling goroutines that it is their turn to read.
		var c = make(chan int)
		// Dispatch go routines
		for i := 0; i < 2; i++ {
			wg.Add(1)
			go func(i int) {
				// Force goroutine to run on different thread than the other
				runtime.LockOSThread()
				// Loop over "my turn to read"-signals until channel is closed
				for count := range c {
					// Read 32 bytes
					var b [32]byte
					var _, err = io.ReadFull(r, b[:])
					// Check for io.EOF if the whole asset was read in the previous read.
					if count == assetSize {
						// The other goroutine read the last bytes, check for io.EOF
						t.failIf(err != io.EOF, "Expected io.EOF, got:", err)
						// Success. Tell the other goroutine to exit by closing the channel.
						close(c)
						break
					}
					// Check for read errors
					t.failIf(err != nil, "Read error:", err)
					// Check content
					for j := range b {
						t.failIf(b[j] != byte(count+j), "Wrong content")
					}
					// Send new offset as read signal
					c <- count + len(b)
				}
				// Tell parent goroutine that we are done.
				wg.Done()
			}(i)
		}
		// Send initial read signal to one of the goroutines
		c <- 0
		// Wait for child goroutines to finish
		wg.Wait()
	},
}
