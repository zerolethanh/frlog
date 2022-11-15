package frlog

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"sort"
)

type Options struct {
	PrintByPath bool
	RawPrint    bool
}

var DefaultOptions = Options{
	PrintByPath: true,
	RawPrint:    false,
}

// PrintAppStacks will print all routes
func PrintAppStacks(app *fiber.App, options ...*Options) {

	var opt Options
	if len(options) > 0 {
		opt = *options[0]
	} else {
		opt = DefaultOptions
	}

	stacks := lo.Flatten[*fiber.Route](app.Stack())

	color.Blue("-- App Route Stacks (%d) --", len(stacks))
	if opt.PrintByPath {
		printByPathStacks(stacks)
	}
	if opt.RawPrint {
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
	paths := sortedKeys(byPathStacks)

	for _, path := range paths {
		color.HiYellow(path)
		routes := byPathStacks[path]
		fmt.Print(" ➜")
		for _, route := range routes {
			params := getRouteParams(route)
			method := route.Method
			c := color.WhiteString

			switch method {
			case "OPTIONS":
				c = color.HiCyanString
			case "GET", "POST", "PUT", "PATCH":
				c = color.HiGreenString
			case "DELETE":
				c = color.HiRedString
			default:
				c = color.WhiteString
			}
			fmt.Print(c(" %s %s", method, params))
		}
		fmt.Println("")
	}
}

func getRouteParams(route *fiber.Route) string {
	var params string
	if len(route.Params) > 0 {
		params = fmt.Sprintf("Params: %v", route.Params)
	}
	return params
}

func sortedKeys(m map[string][]*fiber.Route) []string {
	keys := lo.Keys(m)
	sort.Strings(keys)
	return keys
}
