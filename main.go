package frlog

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type Options struct {
	PrintByPath bool
	RawPrint    bool
}

var DefaultOptions = Options{
	PrintByPath: true,
	RawPrint:    false,
}

func PrintAppStacks(app *fiber.App, options *Options) {

	if options == nil {
		options = &DefaultOptions
	}

	stacks := lo.Flatten[*fiber.Route](app.Stack())

	if options.PrintByPath {
		printByPathStacks(stacks)
	}
	if options.RawPrint {
		printByJson(stacks)
	}

}

func printByJson(stacks []*fiber.Route) {
	var strs []string
	for _, route := range stacks {
		str := fmt.Sprintf("%s %s", route.Method, route.Path)
		strs = append(strs, str)
	}
	var data, _ = json.MarshalIndent(lo.Uniq(strs), "", " ")
	fmt.Println(string(data))
}

func printByPathStacks(stacks []*fiber.Route) {
	byPathStacks := lo.GroupBy(stacks, func(stack *fiber.Route) string {
		return stack.Path
	})
	for path, stacks := range byPathStacks {
		fmt.Println("--", path)
		//var methods string
		for _, stack := range stacks {
			var params string
			//methods += stack.Method + " "
			if len(stack.Params) > 0 {
				params = fmt.Sprintf("Params: %v", stack.Params)
			}
			fmt.Print(fmt.Sprintf("     %s %s", stack.Method, params))
		}
		fmt.Println("")
		//fmt.Println("  ", methods)
	}
}
