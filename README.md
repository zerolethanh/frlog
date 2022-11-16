# frlog
fiber routes log


## Install
```sh
go get github.com/zerolethanh/frlog
```

## Usage

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zerolethanh/frlog"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
	app.Put("/login/:uid", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		return c.SendString("User")
	})
	app.Patch("/fly/:from-:to", func(c *fiber.Ctx) error {
		return c.SendString("Fly")
	})

	frlog.PrintAppStacks(app)

	log.Fatalln(app.Listen(":3000"))
}

```

> Output:
<img width="293" alt="Ảnh màn hình 2022-11-16 lúc 14 22 08" src="https://user-images.githubusercontent.com/2741804/202113314-7712045f-4383-47e7-8ef7-40317ae92cef.png">



## LICENSE
MIT
