监控`otter channel`和`node`运行状态，自动发现所有的`channel`和`otter node`
当`channel`挂起时触发报警，同时监控`otter node`，当node异常时同样触发报警。

刚学习golang，写的不是很好，不喜勿喷
### 测试环境
- CentOS  7.6
- Zabbix server 5.2.5

### 下载
[linux系统](https://github.com/typ431127/otter-zabbix/releases/download/1.0/Otter-zabbix-linux-amd64.zip) 
### 参数说明
```
参数说明:
  -id string
        channel ID (default "0")
  -node
        启用node监控
  -server string
        zk地址 (default "192.168.1.1:2181")
  -timeout duration
        设置连接超时时间 (default 10s)
  -zabbixdiscovery
        返回zabbix自动发现channel json数据
  -zabbixnodediscovery
        返回zabbix自动发现node json数据
默认channel返回值说明:
  START 启动
  STOP  停止
  PAUSE 挂起
  NONE  不存在
otter node返回值说明:
  1 正常
  0 停止
```
### 文件说明
`zbx_otter_templates.yaml` zabbix监控模板，测试版本:5.2.5

### zabbix配置
`zabbix_agentd.conf 配置示例 替换为自己的zk地址`

![image](https://user-images.githubusercontent.com/20376675/120777092-82aaf280-c557-11eb-8578-6bab46ca027d.png)

```
UserParameter=otter_discovery,/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -zabbixdiscovery
UserParameter=otter_nodediscovery,/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -zabbixnodediscovery
UserParameter=otter.status[*],/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -id $1
UserParameter=otter.nodestatus[*],/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -node -id $1

```
### 效果展示
![image](https://user-images.githubusercontent.com/20376675/120776753-23e57900-c557-11eb-9c9d-1e1ea56e8d6e.png)

