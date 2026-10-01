---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-0-administration-guide-78050-migrating-from-ssl-vpn-to-zt-7736d641-1
title: "Migrating from SSL VPN to ZTNA"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-0-administration-guide-78050-migrating-from-ssl-vpn-to-zt-7736d641.md
source_anchor: ""
source_lines: [1, 155]
sha256: 88fe039d9f7f9268600df10d38e743ee5687a3f1e5d7b9f1403edd3e32224f52
---

# Migrating from SSL VPN to ZTNA

# Migrating from SSL VPN to ZTNA

ZTNA can be used to replace VPN-based teleworking solutions to enhance the user experience and to increase security. A typical teleworking configuration may utilize SSL VPN tunnel or web portal mode with LDAP user authentication. Common objects defined for this setup can be reused when migrating to ZTNA, such as the remote LDAP server, user group, and address objects.

### SSL VPN tunnel mode access with LDAP user authentication

Remote users that are in the *ALLOWED-VPN* active directory group have access to a specific web server when they connect through the SSL VPN tunnel. The FortiGate enables split tunneling to the web server so that only traffic to that destination is routed through the tunnel. The web server hosts internal websites that are only accessible by employees.


### SSL VPN web mode access with LDAP user authentication

Remote users that are in the *ALLOWED-VPN* active directory group have access to a specific web server when they connect through the SSL VPN web portal. The web server hosts internal websites that are only accessible by employees. The predefined bookmark to the internal website is the only site that allows remote access.


## Common configurations

This section includes configurations for common objects used in the SSL VPN configuration that can be reused in the ZTNA deployment:

###### To configure an LDAP server:

```
config user ldap
    edit "WIN2K16-KLHOME-LDAPS"
        set server "192.168.20.6"
        set server-identity-check disable
        set cnid "sAMAccountName"
        set dn "dc=KLHOME,dc=local"
        set type regular
        set username "KLHOME\\Administrator"
        set password **********
        set secure ldaps
        set ca-cert "CA_Cert_1"
        set port 636
    next
end
```
                                            
                                            ###### To configure the user group:

```
config user group
    edit "KLHOME-ALLOWED-VPN"
        set member "WIN2K16-KLHOME-LDAPS"
        config match
            edit 1
                set server-name "WIN2K16-KLHOME-LDAPS"
                set group-name "CN=ALLOWED-VPN,DC=KLHOME,DC=local"
            next
        end
    next
end
```
                                            
                                            Firewall addresses can be reused in the server settings for TCP forwarding configurations.

###### To configure the firewall address:

```
config firewall address
    edit "winserver"
        set subnet 192.168.20.6 255.255.255.255
    next
end
```
                                            ## Migrating to ZTNA

The preceding simple SSL VPN tunnel and web mode teleworking solutions can be migrated to ZTNA configurations, providing device authentication using client certificates and additional security posture checks.


Instead of connecting to the SSL VPN tunnel or web portal, the remote user connects to the HTTPS access proxy that forwards traffic to the web server after authentication and security posture checks are completed. This provides granular control over who can access the web resource using role-based access control. It also gives the user transparent access to the website using only their browser.

Migrating to ZTNA includes the following steps:

The first step to configure ZTNA is to connect to and authorize a FortiClient EMS using the EMS connector. There are different ways to connect to an on-premise FortiClient EMS server and a FortiClient EMS Cloud. Refer to the first step of Configure a FortiClient EMS connector for instructions.

ZTNA tags and tagging rules define security posture checks that connecting devices must pass before they are allowed to access protected resources and applications. In the following example, a Zero Trust tagging rule is configured to detect if a virus file exists on an endpoint.

###### To configure a Zero Trust tagging rule on the FortiClient EMS:

1. 
                                                    Log in to the FortiClient EMS.
2. 
                                                    Go to *Zero Trust Tags > Zero Trust Tagging Rules* , and click*Add* .
3. 
                                                    In the *Name* field, enter*Malicious-File-Detected* .
4. 
                                                    In the *Tag Endpoint As* dropdown list, select*Malicious-File-Detected* .EMS uses this tag to dynamically group together endpoints that satisfy the rule, as well as any other rules that are configured to use this tag.
5. 
                                                    Click *Add Rule* then configure the rule:
  1. 
                                                            For *OS* , select*Windows* .
  2. 
                                                            From the *Rule Type* dropdown list, select*File* and click  the*+* button.
  3. 
                                                            Enter a file name, such as *C:\virus.txt* .
  4. 
                                                            Click *Save* .
6. 
                                                            
7. 
                                                    Click *Save* .

A ZTNA solution requires users to be registered and connected to the FortiClient EMS server. When an EMS server is behind the FortiGate, a VIP needs to be defined to allow remote users access to register to the FortiClient EMS. The only port required to be forwarded is TCP/8013. This VIP also needs to be applied in a firewall policy to allow this traffic.

###### To configure a VIP to allow traffic to the EMS server:

1. Go to *Policy & Objects > Virtual IPs* and click*Create New > Virtual IP* .
2. Set *Name* to*VIP-EMS* .
3. Configure the VIP settings:
  1. Set *Interface* to*port1* .
  2. Set *External IP address/range* to*192.168.2.5* .
  3. Set *Map to* to*192.168.20.10* .
  4. Enable *Port Forwarding* .
  5. Set *External service port* to*8013* .
  6. Set *Map to IPv4 port* to*8013* .
4. Set 
5. Click *OK* .

###### To configure the firewall policy:

1. Go to *Policy & Objects >Firewall Policy* and click*Create New* .
2. Set *Name* to*ZTNA-VIP* .
3. Configure the policy settings:
  1. Set *Incoming Interface* to*port1* .
  2. Set *Outgoing Interface* to*port3* .
  3. Set *Source* to*all* .
  4. Set *Destination* to*VIP-EMS* .
  5. For *Service* , select an option that is for TCP/8013.
  6. Disable *NAT* .
  7. Configure the remaining options as needed.
4. Set 
5. Click *OK* .

The ZTNA server defines the external IP and port used for the FortiGate access proxy. It also defines the protected resources that can be accessed through the HTTPS access proxy or TCP forwarding access proxy. The following configuration defines a HTTPS access proxy for accessing the web server on 192.168.20.6.

###### To configure a ZTNA server for HTTPS access proxy:

1. 
                                                    Go to *Policy & Objects > ZTNA* and select the*ZTNA Servers* tab.
2. 
                                                    Click *Create New* .
3. 
                                                    Set *Name* to*WIN2K16-P1* .
4. 
                                                    Configure the network settings: 
  1. 
                                                            Set *External interface* to*port1* .
  2. 
                                                            Set *External IP* to*192.168.2.86* .
  3. 
                                                            Set *External port* to*8443* .
5. 
                                                            
