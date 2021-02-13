# go-zabbix-sender
Simple implementation of a Zabbix sender in Go using the `zabbix` package from https://github.com/adubkov/go-zabbix.  
The main application in my case is to send Zabbix traps from a Synology NAS. Tested on DS220+.

## Build

The current `Makefile` creates a Linux executable:
```
make
```

## Run
Example to update item `backup.status` with value `OK` for host `nas` on Zabbix server `192.168.0.10`:
```bash
./go-zabbix-sender -server 192.168.0.10 -host nas -key backup.status -value OK
```

## Help
```bash
$ ./go-zabbix-sender -h
Usage of ./go-zabbix-sender:
  -host string
    	Specify the host name the item belongs to (default "localhost")
  -key string
    	Item key
  -port int
    	Port of the Zabbix server trapper (default 10051)
  -server string
    	Hostname or IP of the Zabbix server (default "localhost")
  -value string
    	Item value
```