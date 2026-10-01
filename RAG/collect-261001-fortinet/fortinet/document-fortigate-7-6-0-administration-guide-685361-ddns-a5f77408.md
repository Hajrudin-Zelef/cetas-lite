---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-administration-guide-685361-ddns-a5f77408
title: "diagnose debug application ddnscd -1"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-administration-guide-685361-ddns-a5f77408.md
source_anchor: ""
source_lines: [1, 130]
sha256: 2203629438e789d5fcbd5297361c986086cde99d5ef4b469af261e150c6b6157
---

# diagnose debug application ddnscd -1

DDNS
DDNS
If your external IP address changes regularly and you want a static domain name, you can configure the external interface to use a dynamic DNS (DDNS) service. This ensures that external users and customers can always connect to your company firewall. You can configure FortiGuard as the DDNS server using the GUI or CLI.
Multiple DDNS interfaces can be configured in the GUI. The number of DDNS entries that can be configured is restricted by table size, with limits of 16, 32, and 64 entries for entry-level, mid-range, and high-end FortiGates respectively.
A license or subscription is not required to use the DDNS service, but configuring DDNS in the GUI is not supported if:
- 
                                                    The FortiGate model is a 1000-series or higher.
- 
                                                    The FortiGate is a VM.
- 
                                                    The DNS server is not using FortiGuard as the DNS.
|  | DDNS is not supported in transparent mode. | 
Sample topology
In this example, FortiGuard DDNS is enabled and the DDNS server is set to float-zone.com. Other DDNS server options include fortiddns.com and fortidyndns.com.
To configure multiple DDNS entries in the GUI:
- 
                                                    Go to Network > DNS.
- 
                                                    In the Dynamic DNS table, click Create new. The New DDNS Entry pane opens.
- 
                                                    Configure the DDNS entry settings: 
  - 
                                                            Select the Interface with the dynamic connection.
  - 
                                                            Select the Server that you have an account with.
  - 
                                                            Enter the Unique Location.
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Click Create new and repeat step 3 to add more entries.
- 
                                                    Click Apply.
To configure the FortiGuard DDNS service as an IPv4 DDNS server in the CLI:
config system ddns
    edit 1
        set ddns-server FortiGuardDDNS
        set server-type ipv4
        set ddns-domain "branch.float-zone.com"
        set addr-type ipv4
        set use-public-ip enable
        set monitor-interface "wan1"
    next
end
                                            To configure the FortiGuard DDNS service as an IPv6 DDNS server in the CLI:
config system ddns
    edit 1
        set ddns-server FortiGuardDDNS
        set server-type ipv6
        set ddns-domain "fgtatest001.float-zone.com"
        set addr-type ipv6
        set monitor-interface "wan1"
    next
end
                                            DDNS servers other than FortiGuard
If you do not have a FortiGuard subscription, or want to use a different DDNS server, you can configure a DDNS server for each interface. Only the first configure port appears in the GUI.
The available commands vary depending on the selected DDNS server.
To configure DDNS servers other than FortiGuard in the CLI:
config system ddns
    edit <DDNS_ID>
        set monitor-interface <external_interface>
        set ddns-server <ddns_server_selection>
        set server-type {ipv4 | ipv6}
        set ddns-server-addr <address>
        set addr-type ipv6 {ipv4 | ipv6}
    next
end
                                            To configure an IPv6 DDNS client with generic DDNS on port 3 in the CLI:
config system ddns
    edit 1
        set ddns-server genericDDNS
        set server-type ipv6
        set ddns-server-addr "2004:16:16:16::2" "16.16.16.2" "ddns.genericddns.com"
        set ddns-domain "test.com"
        set addr-type ipv6
        set monitor-interface "port3"
    next
end
                                            Refresh DDNS IP addresses
When using a public IP that is not assigned to the FortiGate, the FortiGate cannot trigger an update when the IP address changes. The FortiGate can be configured to refresh DDNS IP addresses by periodically checking the DDNS server at an update interval.
To configure FortiGate to refresh DDNS IP addresses in the CLI:
config system ddns
    edit 1
        set use-public-ip enable
        set update-interval <seconds>
    next
end
                                            When update-interval is set to 0:
- 
                                                    For FortiGuard DDNS, the interval is 300 seconds.
- 
                                                    For third part DDNS servers, the interval is assigned by the DDNS server.
Disable cleartext
When clear-text is disabled, FortiGate uses the SSL connection to send and receive DDNS updates.
To disable cleartext and set the SSL certificate in the CLI:
config system ddns
    edit 2
        set clear-text disable
        set ssl-certificate <cert_name>
    next
end
                                            DDNS update override
A DHCP server has an override command option that allows DHCP server communications to go through DDNS to perform updates for the DHCP client. This enforces a DDNS update of the A field every time even if the DHCP client does not request it. This allows support for the allow, ignore, and deny client‑updates options.
To enable DDNS update override in the CLI:
config system dhcp server
    edit 1
        set ddns-update enable
        set ddns-update-override enable
        set ddns-server-ip <ddns_server_ip>
        set ddns-zone <ddns_zone>
    next
end
                                            Troubleshooting
To debug DDNS:
# diagnose debug application ddnscd -1
# diagnose debug enable
To check if a DDNS server is available:
# diagnose test application ddnscd 3
Not available:
FortiDDNS status:
ddns_ip=0.0.0.0, ddns_ip6=::, ddns_port=443 svr_num=0 domain_num=0
Available:
FortiDDNS status:
ddns_ip=208.91.113.230, ddns_ip6=::, ddns_port=443 svr_num=1 domain_num=3
svr[0]= 208.91.113.230
domain[0]= fortiddns.com
domain[1]= fortidyndns.com
domain[2]= float-zone.com
