package console

import (
	"NetGo/src/utils"
	"fmt"
)

func PrettyBoot(portString string) {
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

func LogRequest(method string, path string, session *utils.Session) {
	fmt.Printf("📨 [%s] %s | Who: %v ClientId: %v\n\n", method, path, session.GetName, session.GetClientId)
}