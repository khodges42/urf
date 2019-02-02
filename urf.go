package main

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)


var executionLocked uint32

var urf = `
	Universal Reload Framework!
                     _.---.._
        _        _.-'         ''-.
      .'  '-,_.-'                 '''.
     (       _                     o  :
      '._ .-'  '-._         \  \-  ---]
                    '-.___.-')  )..-'
                           ==(_/=======[URF]
	
	... starting watch loop.
`


func main() {

	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Printf("Usage: urf dir_to_watch")
		os.Exit(1)
	}

	watchedDir := os.Args[1]
	makePath := fmt.Sprintf("%s/%s",watchedDir,"Makefile")
	makeTarget := "urf"

	if _, err := os.Stat(makePath); os.IsNotExist(err) {
		fmt.Printf("Expecting a Makefile in your watched directory.")
		os.Exit(1)
	}

	w := watcher.New()
	w.SetMaxEvents(1)

	w.FilterOps(watcher.Write, watcher.Create)

	fmt.Printf(urf)

	go func() {
		for {
			select {
			case <-w.Event:
				executeMake(watchedDir, makeTarget)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(watchedDir); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * 2000); err != nil {
		log.Fatalln(err)
	}
}

func executeMake(makePath string, makeTarget string)(){
	// Ensure there isn't a lock (to prevent infinite loops if files are created during make)
	if !atomic.CompareAndSwapUint32(&executionLocked, 0, 1){
		defer atomic.StoreUint32(&executionLocked, 0)
		return
	}

	atomic.StoreUint32(&executionLocked, 1) // Lock for next cycle
	fmt.Println("Found a change. Running make.")
	cmd := exec.Command("make", makeTarget, "-C", makePath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Make failed: %s\n", err)
	}
}