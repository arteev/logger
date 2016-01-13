# logger
Package Golang for simple logger. 

Installation
------------

This package can be installed with the go get command:

    go get github.com/arteev/logger
    
    
Documentation
-------------

Example:
```go
...
logger.InitToConsole(logger.LevelError)
logger.Info.Println("The package Logger!")
logger.Error.Println("This is error message!!!")
...
buf := &bytes.Buffer{}
logger.InitToWriter(logger.LevelTrace,buf)
logger.Info.Println("Out to writer!")
buf.WriteTo(os.Stdout)
...	
```
    
License
-------

  MIT

Author
------

Arteev Aleksey