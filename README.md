# rendering-engine

## Code Documentation

Path: rendering-engine/html-parser/pkg/lexer.go

```go
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0 // 0 is EOF
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}
```

#### Checking for End of Input String (EOF):
```go
if l.readPosition >= len(l.input):
```
This condition checks whether the variable `readPosition` has reached or exceeded the length of the input string `input`.
If `readPosition` is equal to or greater than the length of the input string, it implies that all characters have been read, and the lexer should mark the end of file by setting `l.ch` to 0 (which is used here to represent EOF).

#### Reading the Next Character:
If `readPosition` is less than the length of the input string, the next character from `input` is read and assigned to `l.ch`


#### Updating Position Variables:
```go
l.position = l.readPosition:  # Updates the current character position that the lexer is analyzing.
l.readPosition++:             # Increments readPosition, preparing it for the next call to readChar(), so that the next character will be read next time.
```

#### Example of Application
Suppose we have the following HTML input string:
```html
<!DOCTYPE html>
```
Let’s see how the `readChar()` function will read this string:

Initially, both `position` and `readPosition` are equal to 0.
At the first call to `readChar()`, the character at position 0 (which is `<`) is read, and `position` remains at 0, while `readPosition` is incremented to 1.
