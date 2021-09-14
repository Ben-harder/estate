# estate

Webserver to manage household chore schedule written in Go.

### Brainstorming

A household is a madeup of houeshold members and chore schedules. 

The household member list has the names of those who live in the house stored in alphabetical order. The chore schedules are jobs with dates and responsibilities. 

A chore schedule doesn't care who's doing the chore, and only concerns itself with the next job, and updating that next job when one day has passed after the chore date.

Who's then responsible for connecting the alphabetical household members, to the chore list? I think there's a missing class here that would take in a chore's responsibilities, and the members who need to do it. It would then present that. 

The chore manager maintains the list of household chores and syncs them every x hours. This separates the schedules themselves from the household members. It uses the list of household members and list of schedules to do all this.

### Go Rules

Encapsulation as it is in Go makes me uncomfortable. To make myself feel more comfortable with it, I've enforced some rules on myself for how "classes" are created and represented in code.

Classes must have a constructor with name `New<Obj>`, an interface with name `<obj>Interface`, and a struct with name `<obj>`. Use uppercase first letter for methods or interfaces that should be exposed outside the package. Structs should never be exposed. 

### Example with a dog object
```
func NewDog(name string, breed string) dogInterface {
  return &dog{name: name, breed: breed)
}

type DogInterface interface {
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

### Why?

These rules I'm using force a 1:1 relationship between interfaces and their accompanying struct, which isn't how interfaces were intended to be used. I'm using them this way for a couple reasons. 

1. Having all of a class' methods in one central place helps visualize the class' purpose 
2. Returning an interface from a constructor hides object state completely, preventing any accidental break in encapsulation
3. Using interfaces as much as possible allows for more polymorphism*

\* This is a weaker reasoning since the interfaces I'm making are specific to the struct it wraps, rather than generalized interfaces. I'm aware that small interface are good practice. 
