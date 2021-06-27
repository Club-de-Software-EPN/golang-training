# Arrays & Slices
### Arrays
In Golang:
- An array is a numbered sequence of elements of a single type.
- Their size cannot be changed (fixed size).

```
var variable [size]type
```
Or,
```
variable := [size]type{element1, element2, ..., elemente_size}
```

### Slices
In Golang:
- Are similar to arrays but their size can be changed on runtime.
- Are a flexible view into the elements of an array.
- Slices don't store any data, they only describes a section of an array.

```
a[lowBound : highBound]
```
Or,
```
mySlice := make([]int, 3) // [0, 0, 0]
```
The make function allocates a zeroed array and returns a slice that refers to that array. To specify a capacity, pass a third argument :

```
var length int = 0
var capacity int = 5
mySlice := make([]int, length, capacity) // []
```


### Slice literals
- They are like an array literal without the length.
```
variable := []type{element1, element2, ...}
```

### Functions
- **append()**

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
```
result = append(mySlice, newValue)
```
- **len() y cap()**

The length and capacity of a slice can be obtained using len() and cap().
```
mySlice := make([]int, 0, 5) // []
fmt.Println(len(mySlice)) // 0
fmt.Println(cap(mySlice)) // 5
```

### Basic for-each loop
- We use *range* expression to iterate over an array or a slice.
- Two values are returned for each iteration: the index and the element at that index.

```
var sliceLiteral = []string{"a", "e", "i", "o", "u"}
for index, element := range sliceLiteral {
    fmt.Printf("index: %d, element: %s\n", index, element)
}
```
