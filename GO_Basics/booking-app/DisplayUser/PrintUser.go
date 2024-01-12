package displayuser
import (
	"fmt"
)
func PrintUser(name *string) {
	fmt.Println("Welcome ", *name)
}