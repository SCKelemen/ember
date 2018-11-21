# ember 🔥

```
// declare name space
package main

// import classes
import "collections"

// type declarations / unions
type int ( int8 | int16 | int32 | int64 )
type uint ( uint8 | uint16 | uint32 | uint64 )

// type aliases
type byte uint8
type rune int32

// interfaces
interface ISortAlgorithm {
    Sort(collection: ICollection) ICollection 
}

// abstract classes 
abstract class SortAlgorithm<T> : ISortAlgorithm {
    string _name

    SortAlgorithm(name: String) {
        _name = name
    }

    string Name => _name

    abstract Sort(collection: ICollection<T>) ICollection<T>
}

// classes
class QuickSort<T> : SortAlgorithm {
    QuickSort() : base("QuickSort") {
        // empty constructor
    } 

    Sort(collection: ICollection<T>) ICollection<T> {
        // implement quick sort
    }
}

// actors
actor Main {

}
```



### declarations

```

func GetPriceForUser() float32 {

    let price = 1.23        // assign-once 
    var discount = 0.15     // variable
    const cost = 1.00       // compile-time constant

    return floor( price * ( 1 - discount ), cost )
}


```

# Compiler Strategies

## Highly Concurrent and Distributed

As time has continued, so has progress. Unfortunately for code, this often means more code. Although the cloud services industry has turned to microservice architectures to allow of polyglot development, developer velocity, and highly decoupled systems, large companies such as Google and Microsoft, have admitted that the properties of shared repositories provide strong motiviations for not adopting a strictly microservice strategy. The design philosphies behind microservice exist without and independently from separate repositories. With these factors under consideration, the compiler should be designed to maximize these traits, without requiring them. Therefore, a goal of this project is to allow for quick compilations to aid in developer velocity, despite any size complexity of a project. The compiler will make use of not only concurrency primitives, but distributed primitives as well, to allow for horizontal scaling. The compiler should work well on projects of any size.

### Concurrency Models

### Distributed Models
