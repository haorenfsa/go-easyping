# go-easyping
`go-easyping` is aim to provide very easy to use `Ping` method  
inspired by `github.com/sparrc/go-ping` and `github.com/tatsushid/go-fastping`  
I've tried to use these 2 package but find it not easy enough for most users.  
So I write my own.

## Ping in 2 Lines of Code
simple use:
``` go
import "github.com/haorenfsa/go-easyping"
// ...
delay, err := easyping.Ping("127.0.0.1") // delay is a time.Duration
// ...
```

advanced:
``` go
import "github.com/haorenfsa/go-easyping"
// ...
opt := &easyping.Option{...}
result, err := easyping.AdvancedPing(opt)
// ...
```

# Installation
Install and update with `go get -u github.com/haorenfsa/go-easyping`

# Todo
- ~~add timeout~~
- ~~support host~~
- ~~handle errors~~
- support some options
- multiple pings with statistics