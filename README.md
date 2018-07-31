# WMI exporter

[![Build status](https://ci.appveyor.com/api/projects/status/ljwan71as6pf2joe?svg=true)](https://ci.appveyor.com/project/martinlindhe/wmi-exporter)

Prometheus exporter for Windows machines, using the WMI (Windows Management Instrumentation).


## Collectors

Name     | Description | Enabled by default
---------|-------------|--------------------
ad | [Win32_PerfRawData_DirectoryServices_DirectoryServices](https://msdn.microsoft.com/en-us/library/ms803980.aspx) Active Directory |
cpu | [Win32_PerfRawData_PerfOS_Processor](https://msdn.microsoft.com/en-us/library/aa394317(v=vs.90).aspx) metrics (cpu usage) | &#10003;
cs | [Win32_ComputerSystem](https://msdn.microsoft.com/en-us/library/aa394102) metrics (system properties, num cpus/total memory) | &#10003;
dns | [Win32_PerfRawData_DNS_DNS](https://technet.microsoft.com/en-us/library/cc977686.aspx) metrics (DNS Server) |
hyperv | Performance counters for Hyper-V hosts |
iis | [Win32_PerfRawData_W3SVC_WebService](https://msdn.microsoft.com/en-us/library/aa394345) IIS metrics |
logical_disk | [Win32_PerfRawData_PerfDisk_LogicalDisk](https://msdn.microsoft.com/en-us/windows/hardware/aa394307(v=vs.71)) metrics (disk I/O) | &#10003;
net | [Win32_PerfRawData_Tcpip_NetworkInterface](https://technet.microsoft.com/en-us/security/aa394340(v=vs.80)) metrics (network interface I/O) | &#10003;
msmq | [Win32_PerfRawData_MSMQ_MSMQQueue](http://wutils.com/wmi/root/cimv2/win32_perfrawdata_msmq_msmqqueue/) metrics (MSMQ/journal count) |
[mssql] | a meta collector that loads the following mssql_* collectors  |
mssql_availreplica | [Win32_PerfRawData_MSSQLSERVER_SQLServerAvailabilityReplica](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-availability-replica) metrics (SQL Server Availaility Group Replica) |
mssql_bufman | [Win32_PerfRawData_MSSQLSERVER_SQLServerBufferManager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object) metrics (SQL Server Buffer Manager) |
mssql_databases | [Win32_PerfRawData_MSSQLSERVER_SQLServerDatabases](  https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object) metrics (SQL Servers Databases) |
mssql_dbreplica | [Win32_PerfRawData_MSSQLSERVER_SQLServerDatabaseReplica](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-database-replica) (SQL Server Database Replica) |
mssql_genstats | [Win32_PerfRawData_MSSQLSERVER_SQLServerGeneralStatistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-general-statistics-object) (SQL Server General Statistics) |
mssql_locks | [Win32_PerfRawData_MSSQLSERVER_SQLServerLocks](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-locks-object) (SQL Server Locks) |
mssql_memmgr | [Win32_PerfRawData_MSSQLSERVER_SQLServerMemoryManager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-memory-manager-object) (SQL Server Memory Manager) |
mssql_sqlstats | [Win32_PerfRawData_MSSQLSERVER_SQLServerSQLStatistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-sql-statistics-object) (SQL Server SQL Statistics) |
os | [Win32_OperatingSystem](https://msdn.microsoft.com/en-us/library/aa394239) metrics (memory, processes, users) | &#10003;
process | [Win32_PerfRawData_PerfProc_Process](https://msdn.microsoft.com/en-us/library/aa394323(v=vs.85).aspx) metrics (per-process stats) |
service | [Win32_Service](https://msdn.microsoft.com/en-us/library/aa394418(v=vs.85).aspx) metrics (service states) | &#10003;
system | Win32_PerfRawData_PerfOS_System metrics (system calls) | &#10003;
tcp | [Win32_PerfRawData_Tcpip_TCPv4](https://msdn.microsoft.com/en-us/library/aa394341(v=vs.85).aspx) metrics (tcp connections) |
textfile | Read prometheus metrics from a text file | &#10003;
vmware | Performance counters installed by the Vmware Guest agent |

The HELP texts shows the WMI data source, please see MSDN documentation for details.

## Installation
The latest release can be downloaded from the [releases page](https://github.com/martinlindhe/wmi_exporter/releases).

Each release provides a .msi installer. The installer will setup the WMI Exporter as a Windows service, as well as create an exception in the Windows Firewall.

If the installer is run without any parameters, the exporter will run with default settings for enabled collectors, ports, etc. The following parameters are available:

Name | Description
-----|------------
`ENABLED_COLLECTORS` | As the `--collectors.enabled` flag, provide a comma-separated list of enabled collectors
`LISTEN_ADDR` | The IP address to bind to. Defaults to 0.0.0.0
`LISTEN_PORT` | The port to bind to. Defaults to 9182.
`METRICS_PATH` | The path at which to serve metrics. Defaults to `/metrics`
`TEXTFILE_DIR` | As the `--collector.textfile.directory` flag, provide a directory to read text files with metrics from

Parameters are sent to the installer via `msiexec`. Example invocation:

```powershell
msiexec /i <path-to-msi-file> ENABLED_COLLECTORS=os,iis LISTEN_PORT=5000
```

## Roadmap

See [open issues](https://github.com/martinlindhe/wmi_exporter/issues)


## Usage

    go get -u github.com/golang/dep
    go get -u github.com/prometheus/promu
    go get -u github.com/martinlindhe/wmi_exporter
    cd $env:GOPATH/src/github.com/martinlindhe/wmi_exporter
    promu build -v .
    .\wmi_exporter.exe

The prometheus metrics will be exposed on [localhost:9182](http://localhost:9182)

## Examples

### Enable only service collector and specify a custom query

    .\wmi_exporter.exe --collectors.enabled "service" --collector.service.services-where "Name='wmi_exporter'"

### Enable only process collector and specify a custom query

    .\wmi_exporter.exe --collectors.enabled "process" --collector.process.processes-where "Name LIKE 'firefox%'"

When there are multiple processes with the same name, WMI represents those after the first instance as `process-name#index`. So to get them all, rather than just the first one, the query needs to be a wildcard search.

## License

Under [MIT](LICENSE)
