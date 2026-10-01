---
id: collect-261001-general-networking/general-networking/troubleshooting-tip-fortiguard-ddns-ip-update-fails-community-2
title: "troubleshooting-tip-fortiguard-ddns-ip-update-fails-community"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/troubleshooting-tip-fortiguard-ddns-ip-update-fails-community.md
source_anchor: ""
source_lines: [210, 253]
sha256: 7a7791021ac05e836bfdeecc662177c8ceea735044aad62d60a5e3457107f891
---

# troubleshooting-tip-fortiguard-ddns-ip-update-fails-community

```
config system global
    set ssl-min-proto-version TLS1.0
end
```
**Note:** An MTU mismatch on the interface can also cause a TLS handshake failure like this, as an alternative to lowering ssl-min-proto-version. Verify the path MTU first with execute ping-options df-bit yes, followed by execute ping-options data-size <MTU size> and **execute ping 8.8.8.8**, then adjust the interface MTU if needed. See Technical Tip: How to adjust the Maximum Transmission Unit (MTU) value on a FortiGate interface for details on adjusting the MTU.


One more reason why the DDNS update might fail is if the WAN interface IP is assigned via DHCP and the 'override internal DNS' setting is enabled. The DNS server from the ISP may not be able to resolve the DDNS domain (globalddns.fortinet.net) and retrieve the IP for the FortiGuard DDNS servers.

In cases like this, the override has to be disabled:

In the GUI:


In the CLI:

```
config system interface
    edit "<wan interface>"
        set dns-server-override disable
    next
end
```

Restarting the FortiGate DDNS Client (ddnscd):

If FortiGuard **DDNS updates fail or become unresponsive**, restart the internal DDNS client daemon **(ddnscd)** to clear the session and force a fresh update.**Example use case:**

This can be useful when DDNS appears stuck, not updating despite correct configuration and connectivity.

Run the following command to terminate the running DDNS client daemon:

In the CLI:

`fnsysctl killall ddnscd`

This uses the FortiGate internal **fnsysctl** utility to kill all instances of the **ddnscd** process. Upon termination, the system will automatically restart the process as needed. 

If none of these steps allow the correct update of the IP, contact the Fortinet TAC team by creating a ticket for the issue and providing the above logs.**Updating the device with a new ISP link:**

If the ISP link is getting changed with a new public-ip and the DDNS resolving to that entry also needs to be changed, reach out to the TAC team to delete the old DDNS entry from the database.

After that, under DDNS settings via the CLI, delete the copy config for the previous one and delete that entry. Once done, paste the copied configuration and only change the attribute for set monitor-interface to the new WAN port.
