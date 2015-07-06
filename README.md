#ngrep

A network utility for checking the content of potentially unbounded streams.

## usage

```
# a service is running on port 5000, we expect it to say "hello"
# within four seconds of us connecting
ngrep -t=4 hello localhost 5000

# we want to check that a client program will make an http GET request
# on port 8001
ngrep -l GET localhost 8001
```

ngrep returns a 0 exit code upon successfuly finding its pattern, otherwise
it returns a nonzero exit code.


