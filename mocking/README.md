You have been asked to write a program which counts down from 3, printing each number on a new line (with a 1-second pause) and when it reaches zero it will print "Go!" and exit.

```
Copy
3
2
1
Go!
```

We'll tackle this by writing a function called `Countdown` which we will then put inside a `main` program so it looks something like this:

```
package main

func main() {
	Countdown()
}
```