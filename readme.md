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
result, err := easyping.Ping(opt)
// ...
```

# Installation
Install and update with `go get -u github.com/haorenfsa/go-easyping`

# About Protocal

ICMP Type Code

|TYPE|CODE|Description|(Q)uery/(E)rror|
|---|---|---|---|
|0|0|Echo Reply——回显应答（Ping应答）|`Q`|
|3|0|Network Unreachable——网络不可达|E|
|3|1|Host Unreachable——主机不可达|E|
|3|2|Protocol Unreachable——协议不可达|E|
|3|3|Port Unreachable——端口不可达|E|
|3|4|Fragmentation needed but no frag. bit set——需要进行分片但设置不分片比特|E|
|3|5|Source routing failed——源站选路失败|E|
|3|6|Destination network unknown——目的网络未知|E|
|3|7|Destination host unknown——目的主机未知|E|
|3|8|Source host isolated (obsolete)——源主机被隔离（作废不用）|E|
|3|9|Destination network administratively prohibited——目的网络被强制禁止|E|
|3|10|Destination host administratively prohibited——目的主机被强制禁止|E|
|3|11|Network unreachable for TOS——由于服务类型TOS，网络不可达|E|
|3|12|Host unreachable for TOS——由于服务类型TOS，主机不可达|E|
|3|13|Communication administratively prohibited by filtering——由于过滤，通信被强制禁止|E|
|3|14|Host precedence violation——主机越权|E|
|3|15|Precedence cutoff in effect——优先中止生效|E|
|4|0|Source quench——源端被关闭（基本流控制）|-|
|5|0|Redirect for network——对网络重定向|-|
|5|1|Redirect for host——对主机重定向|-|
|5|2|Redirect for TOS and network——对服务类型和网络重定向|-|
|5|3|Redirect for TOS and host——对服务类型和主机重定向|-|
|8|0|Echo request——回显请求（Ping请求）|`Q`|
|9|0|Router advertisement——路由器通告|E|
|10|0|Route solicitation——路由器请求|E|
|11|0|TTL equals 0 during transit——传输期间生存时间为0|E|
|11|1|TTL equals 0 during reassembly——在数据报组装期间生存时间为0|E|
|12|0|IP header bad (catchall error)——坏的IP首部（包括各种差错）|E|
|12|1|Required options missing——缺少必需的选项|E|
|13|0|Timestamp request (obsolete)——时间戳请求（作废不用）|`Q`|
|14|-|Timestamp reply (obsolete)——时间戳应答（作废不用）|`Q`|
|15|0|Information request (obsolete)——信息请求（作废不用）|`Q`|
|16|0|Information reply (obsolete)——信息应答（作废不用）|`Q`|
|17|0|Address mask request——地址掩码请求|`Q`|
|18|0|Address mask reply——地址掩码应答|`Q`|

above table reference: http://www.cnitblog.com/yang55xiaoguang/articles/59581.html

# Todo
- add timeout
- ~~support host~~
- handle errors
- support some options
