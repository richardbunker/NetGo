package console

import (
	"fmt"
)

func PrettyBoot() {
	fmt.Print(`
N   N EEEEE TTTTT GGG   OOO 
NN  N E       T  G     O   O
N N N EEEE    T  G  GG O   O
N  NN E       T  G   G O   O
N   N EEEEE   T   GGG   OOO

`)
	fmt.Println("🚀 Server preparing to launch...")
	fmt.Println("📡 Server is listening on port 3000")
}