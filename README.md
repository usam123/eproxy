获取二进制文件或者自己编译
----
build/proxy (linux amd64)

添加可执行权限
----
进入proxy所在目录， chmod u+x proxy

运行
----
nohup ./proxy &
```
简单实用步骤：
$  wget https://github.com/usam123/eproxy/raw/master/build/proxy
$  sudo chmod 0777 proxy
$  sudo nohup ./proxy > proxy.log 2>&1 &

配置页面地址.
http://yourIP_or_Domain/config.
如果域名接入CF，配置页面只能用http://yourIp/config 访问.
```


