分布式全局ID生成器<br>
====
介绍：<br>
----
服务器端使用Go语言实现，支持多语言客户端（python，java，c++等等），现有功能支持64位和32位ID生成，能通过client端传值进行生成规则配置，计划在未来计划提供黑名单功能，支持号码过滤。<br>
准备：<br>
----
1、准备一套zeekeeper，一套redis，本地sqlite（详细安装请查阅官网）<br>
2、在res/config.ini 中进行相关配置<br>
3、服务端需要有go运行环境（详细安装请查阅官网）<br>
服务端：<br>
----
1、进入src/ids/ids.go <br>
2、运行 go run ids.go <br>
客户端：<br>
----
写那么多累死宝宝了，在protocol中有thrift文件，可根据自己需要的语言生成。（后续有时间提供各语言的demo）
