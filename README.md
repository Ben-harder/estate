# estate

Webserver to manage household chore schedule written in Go.

### Go Rules

Encapsulation as it is in Go makes me uncomfortable. To make myself feel more comfortable with it, I've enforced some rules on myself for how "classes" are created and represented in code.

Classes must have a constructor with name `New<Obj>`, an interface with name `<obj>Interface`, and a struct with name `<obj>`. Use uppercase first letter for methods or interfaces that should be exposed outside the package. Structs should never be exposed. 

### Example with a dog object
```
func NewDog(name string, breed string) DogInterface {
  return &dog{name: name, breed: breed)
}

type DogInterface interface {
  Bark()
}

type dog struct {
  name string
  breed string
}

func (d *dog) Bark() {
  fmt.Println("ruff!")
}
```

### Why?

These rules I'm using force a 1:1 relationship between interfaces and their accompanying struct, which isn't how interfaces were intended to be used. I'm using them this way for a couple reasons. 

1. Having all of a class' methods in one central place helps visualize the class' purpose 
2. Returning an interface from a constructor hides object state completely, preventing any accidental break in encapsulation
3. Using interfaces as much as possible allows for more polymorphism*

\* This is a weaker reasoning since the interfaces I'm making are specific to the struct it wraps, rather than generalized interfaces. I'm aware that small interface are good practice. 
