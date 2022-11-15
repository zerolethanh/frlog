# frlog
fiber routes log


## Install
```sh
go get github.com/zerolethanh/frlog
```

## Usage

```go

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", func(c *fiber.Ctx) error {
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
<img width="313" alt="Ảnh màn hình 2022-11-15 lúc 18 53 41" src="https://user-images.githubusercontent.com/2741804/201913350-b5f806e9-3bdf-4b80-b121-6c98d768648f.png">



## LICENSE
MIT
