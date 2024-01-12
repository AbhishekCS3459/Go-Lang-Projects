package users
import (
	"fmt"
)
func ChangeUser(name *string, newName string) {
	fmt.Println("Changed Name from", *name, "to", newName)
	*name = newName 
}
