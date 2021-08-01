# estate

Webserver to manage household chore schedule.

### Go Rules

This server is written in Go with certain rules in place.

#### Objects

Objects must have a new function with name `new<Obj>`, an interface with name `<obj>Interface`, and a struct with name `<obj>`. Use uppercase first letter for methods or interfaces that should be exposed outside the package. Structs should never be exposed. 

Example with a dog object:
```
func newDog(name string, breed string) dogInterface {
  return &dog{name: name, breed: breed)
}

type dogInterface interface {
  bark()
}

type dog struct {
  name string
  breed string
}

func (d *dog) bark() {
  fmt.Println("ruff!")
}
```
