package beta

import (
	"fmt"

	alpha "github.com/Artemis-Cooperative/alpha"
	packages "github.com/Artemis-Cooperative/interfaces/packages"
)

var alphaInstance packages.Alpha = alpha.Alpha{}

type Beta struct{}

func (b Beta) Run() {
	fmt.Println("beta")
	alphaInstance.Run()
}
