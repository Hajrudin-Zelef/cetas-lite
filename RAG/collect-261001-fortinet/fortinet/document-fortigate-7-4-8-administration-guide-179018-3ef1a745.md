---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-179018-3ef1a745
title: "Using FortiManager as a local FortiGuard server"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-179018-3ef1a745.md
source_anchor: ""
source_lines: [1, 143]
sha256: 9995174024cc8cd06160df8c5bf22c761adaa8fd83ecce8d29651991125a909d
---

# Using FortiManager as a local FortiGuard server

# Using FortiManager as a local FortiGuard server

FortiManager can provide a local FortiGuard server with port 443 access.

Anycast FortiGuard settings force the rating process to use port 443, even with an override server. Using a unique address in the same subnet as the FortiManager access IP address, the FortiManager can provide local FortiGuard updates and rating access with a dedicated IP address and port 443.

|  | On FortiManager, use the *Bind to IP* addresses for the update and rating services over TCP/443. The *Bind to IP* address does not need to be configured for update services if the default port was not changed to TCP/443. See Configuring network interfaces  in the FortiManager Administration Guide for more information. | 

###### To use a FortiManager as a local FortiGuard server in the GUI:

1. 
                                                    Go to *System > FortiGuard*
2. 
                                                    In the *Override FortiGuard Servers* table, click*Create New* . The*Create New Override FortiGuard Server* pane opens.
3. 
                                                    Select the server address type: *IPv4* ,*IPv6* , or*FQDN* .
4. 
                                                    Enter the FortiManager address in the *Address* field.
5. 
                                                    Select the type of server: *AntiVirus & IPS Updates* ,*Filtering* , or*Both* .
6. 
                                                    Click *OK* .
7. 
                                                    Click *Create New* again to add a second override FortiManager for filtering.
8. 
                                                    Click *OK* , then click*Apply* .

###### To use a FortiManager as a local FortiGuard server in the CLI:

```
config system central-management
    set type fortimanager
    set fmg "172.18.37.148"
    config server-list
        edit 1
            set server-type update
            set server-address 172.18.37.150
        next
        edit 2
            set server-type rating
            set server-address 172.18.37.149
        next
    end
    set fmg-update-port 443
    set include-default-servers enable
end
```
                                            When `fmg-update-port` is set to `443`, the update process will use port 443 to connect to the override update server, which is the local FortiGuard server in the FortiManager. If this is not set, the update process will use port 8890, and the server address setting must be the FortiManager access IP address. Override FortiGuard services come from the server list that is the local FortiGuard server in the FortiManager, and use the traditional, non-OCSP TLS handshake. If override servers in the FortiManager are not available, the default FortiGuard servers are connected, and the anycast OCSP TLS handshake is used. 

The FortiManager IP address used in `set server-address` (for example, `set server-address 172.18.37.149`) corresponds to the *Bind to IP* setting configured on the FortiManager interface.

## HA considerations

When FortiGate and FortiManager units are both in high availability (HA) clusters, and FortiGate is using FortiManager as an update server, you must configure the following secondary IP addresses:

- 
                                                    Configure secondary IP addresses for the service on both FortiManager units in the HA cluster
- 
                                                    Configure FortiGate to use the secondary FortiManager IP addresses.

In this example, the primary FortiManager unit is configured to use the secondary IP address of `10.4.1.204` for the rating service, and the secondary FortiManager unit is configured to use the secondary IP address of `10.4.1.205` for the rating service. FortiGate is configured to use the secondary IP addresses on FortiManager.

###### To configure FortiManager units in an HA cluster as web filter rating servers:

1. 
                                                    On the primary FortiManager in the HA cluster, configure a secondary IP address for the rating service: The `set rating-service-ip` is set to the secondary IP address for the rating query.```
config system interface
    edit "port1"
        set ip 10.4.1.104 255.255.0.0
        set allowaccess ping https ssh snmp http webservice
        set serviceaccess fgtupdates fclupdates webfilter-antispam
        
```
**set rating-service-ip 10.4.1.204 255.255.0.0** set type physical
    next
end
2. 
                                                    On the secondary FortiManager in the HA cluster, configure a secondary IP address for the rating service: The `set rating-service-ip` is set to the secondary IP address for the rating query.```
config system interface
    edit "port1"
        set ip 10.4.1.105 255.255.0.0
        set allowaccess ping https ssh snmp http webservice
        set serviceaccess fgtupdates fclupdates webfilter-antispam
        
```
**set rating-service-ip 10.4.1.205 255.255.0.0** set type physical
    next
end
3. 
                                                    On both the primary and secondary FortiManager units in the HA cluster, enable web filter query: ```
config fmupdate service
    
```
**set query-webfilter enable** end
4. 
                                                    On both the primary and secondary FortiManager units in the HA cluster, enable all web filter logs: ```
config fmupdate web-spam fgd-setting
    
```
**set wf-log all** end
5. 
                                                    As needed, run the `diagnose debug application fgdsvr 255` command.For additional details, use shell and run `bash$ tail -f /var/log/fgdsvr.log` to display the incoming URL queries.

###### To configure FortiGate:

1. 
                                                    On FortiGate, configure central management to use the secondary FortiManager IP addresses for the rating service: The `set rating-service-ip` is set to the secondary IP address on the primary and secondary FortiManager units in the HA cluster.```
config system central-management
    set type fortimanager
    config server-list
        edit 1
            set server-type update
            set server-address 10.4.1.104              
        next
        edit 3
            set server-type update
            set server-address 10.4.1.105
        next
        edit 2
            set server-type rating
            
```
**set server-address 10.4.1.204** next
        edit 4
            set server-type rating**set server-address 10.4.1.205** next
    end
    set include-default-servers disable
end
2. 
                                                    Disable Anycast: ```
config system fortiguard
    
```
**set fortiguard-anycast disable** end
3. 
                                                    As needed, run the following diagnose commands: 
  - 
                                                            Run `diagnose debug app urlfilter -1` to show the rating query response.
  - 
                                                            Run `diagnose debug rating` to show the rating service information.
4.
