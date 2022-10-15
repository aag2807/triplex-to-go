# Triplex to Go
small enough to take on the go. 🚀

## About the project

A byte size library largely inspired by Triplex from lsolano <a href="https://github.com/lsolano/triplex">The Repository<a>. 
This is solely used for internal projects.

## Validation types ##
### Arguments ###

It validates the the method declared parameters, if an argument is invalid it will cause the function to panic.

```go
arguments := Arguments{}

func sum(a,b int) int{
    arguments.IsGreaterThan(a,b, "the first arg cannot be greater than the second")

    return a + b
}
```

### State ###

It validates a given peace of transient state, to ensure pre-conditions or invariants, if a condition is invalid it will cause the function to panic.

```go
state := State{}

func sum(a,b int) int{
    state.IsTrue(a != b, "The numbers cannot be equal")

    return a + b
}
```