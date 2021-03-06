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

### system interfaces

```ts
interface IReader {
    Read(): void 
}

interface IWriter {
    Write(input: Bytes): void 
}

type Reader = IReader;
type Writer = IWriter;
type ReaderWrite = Reader & Writer;




type Result<T> = { result: T }

function Result<T>(result: T): Result<T> {
    return {result};
}

interface IError {
    Message?: string
    Code?: number
}


type Try<T> = Result<T> | IError
type Possible<T> = T | IError

type Some<T> = { value: T }

function Some<T>(value: T): Some<T> {
    return {value};
}


type None = {}
const None: None = {};

type Maybe<T> = Some<T> | None
 

```

# Compiler Strategies

## Highly Concurrent and Distributed

As time has continued, so has progress. Unfortunately for code, this often means more code. Although the cloud services industry has turned to microservice architectures to allow of polyglot development, developer velocity, and highly decoupled systems, large companies such as Google and Microsoft, have admitted that the properties of shared repositories provide strong motiviations for not adopting a strictly microservice strategy. The design philosphies behind microservice exist without and independently from separate repositories. With these factors under consideration, the compiler should be designed to maximize these traits, without requiring them. Therefore, a goal of this project is to allow for quick compilations to aid in developer velocity, despite any size complexity of a project. The compiler will make use of not only concurrency primitives, but distributed primitives as well, to allow for horizontal scaling. The compiler should work well on projects of any size.

## The Compilation Pipeline

There are several components to compilers. We can think of compilers in many ways. Throughout this project, we will refer to parts of the system in the following ways:

### Syntax, Semantics, and Pragmatics

Syntax refers to the language constructs of the system; usually sugars.

Semantics refers to the structure of the language; usually the semantic tree. 

Pragmatics refers to the implementation of the language: usually the compiler.


Lexical Analysis
Lexical Analysis is the process by which source code is scanned for Tokens, and Lexemes are generated. The input to the process is the byte stream of a source code file. In our case, this is a stream of UTF-8 encoded bytes. The output of this process is a stream of Tokens or Lexemes. In our case, Tokens refer to a type, and a Lexeme is an instance of this token type.

Syntactical Parsing
Syntacitcal Parsing is the process by which the tokens or lexemes are analyzed for language constructs. The input to this process is the stream of Lexemes, which themselves contain Tokens. The output of this process is a graph, or a Syntax Tree.

### Directories

Files and Directories form a graph, or a tree. 


![File System Graph](docs/img/fsgraph.svg)
<details> 

digraph G { DIRECTORY [shape=box]; FILE; source [shape=box]; component1 [shape=box]; component2 [shape=box]; component3 [shape=box]; c3subdir1 [shape=box]; c3subdir2 [shape=box]; source -> component1; source -> component2; source -> component3; component1 -> c1file1; component1 -> c1file2; component2 -> c2file1; component2 -> c2file2; component2 -> c2file3; component3 -> c3file1; component3 -> c3subdir1; component3 -> c3subdir2; c3subdir1 -> c3sd1file1; c3subdir1 -> c3sd1file2; c3subdir1 -> c3sd1file3; c3subdir2 -> c3sd2file1; }
</details>

We can use this tree to define boundaries of concurrency. Each subgraph is dispatched to a compiler processor, which performs lexical scanning. Once scanning is complete, the subgraphs are merged in a mutation which combines the results of all subgraph processors. We can guarentee this merge using a lattice structure. 

### Concurrency Models

### Distributed Models


## Resources

### Probabilistic 
   - [ Erik Meijer ](https://www.youtube.com/watch?v=NKeHrApPWlo)

### Parsing
   - [Tolerant Parsing ](https://github.com/Microsoft/tolerant-php-parser/blob/master/docs/HowItWorks.md)

### Tooling
   - [Language Server (Protocol)](https://code.visualstudio.com/docs/extensions/example-language-server)

### Scheduling
  Books
  
  - Intelligent Scheduling                          *Monte Zweben, Mark S. Fox*
  - Scheduling: Theory, Algorithms, and Systems     *Michael Pinedo*
         
### Collection
  Books
  - The Garbage Collection Handbook                 *Richard Jones, Antony Hosking, Eliot Moss*
  - [GC Handbook](http://gchandbook.org/)

### Actors

  - [Hewitt, Meijer and Szyperski: The Actor Model](https://youtu.be/7erJ1DV_Tlo)

### Stats
  - [ Beating Machine Learning ](https://www.youtube.com/watch?v=NKeHrApPWlo)
