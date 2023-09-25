package main

import (
	"context"
	"fmt"
	"os"

	app "git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/app/technical"
)

func main() {
	a := app.New()

	if err := a.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.Run(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
