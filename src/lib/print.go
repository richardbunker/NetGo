package lib

import (
	"fmt"
)

func StartUpMessage(portString string) {
	fmt.Print(`
    _   __     __  ______    
   / | / /__  / /_/ ____/___ 
  /  |/ / _ \/ __/ / __/ __ \
 / /|  /  __/ /_/ /_/ / /_/ /
/_/ |_/\___/\__/\____/\____/ 

`)
	fmt.Println("🚀 Launching NetGo...")
	fmt.Println()
	fmt.Printf("📡 Server is listening on port %s\n\n", portString)
}
