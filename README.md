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
