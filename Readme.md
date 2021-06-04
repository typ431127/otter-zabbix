监控`otter channel`和`node`运行状态，当channel挂起时触发报警，同时监控otter node，当node异常时同样触发报警

刚学习golang，写的不是很好，不喜勿喷
### 下载
[linux系统](https://github.com/typ431127/otter-zabbix/releases/download/1.0/Otter-zabbix-linux-amd64.zip) 

### 文件说明
`zbx_otter_templates.yaml` zabbix监控模板，测试版本:5.2.5

### zabbix配置
`zabbix_agentd.conf 配置示例`
```
UserParameter=otter_discovery,/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -zabbixdiscovery
UserParameter=otter_nodediscovery,/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -zabbixnodediscovery
UserParameter=otter.status[*],/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -id $1
UserParameter=otter.nodestatus[*],/etc/zabbix/scripts/otter_zabbix -server 192.168.1.1:2181 -node -id $1

```
