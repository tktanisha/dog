package dog

import "fmt"

func Greet(str string) string {
	return fmt.Sprintf("Hello, " + str)

}

func Bark(Str string) string {
	return fmt.Sprintf(" %s is doing bhow bhow", Str)
}

func Meow(Str string) string {
	return fmt.Sprintf(" %s is doing meow meow", Str)
}
