---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-2-2-administration-guide-33548-configuring-client-bd07f012
title: "Configuring client certificate authentication on the LDAP server"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-2-2-administration-guide-33548-configuring-client-bd07f012.md
source_anchor: ""
source_lines: [1, 110]
sha256: 794dca46dfa555794a74ec3e9aa3ba5f4752c2b113f83b7a67fc7a4ccc197fe5
---

# Configuring client certificate authentication on the LDAP server

Administrators can configure a FortiGate client certificate in the LDAP server configuration when the FortiGate connects to an LDAPS server that requires client certificate authentication.

```
config user ldap
    edit <ldap_server>
        set client-cert-auth {enable | disable}
        set client-cert <source>
    next
end
```
                                            ## Example

In this example, the FortiGate is configured as an explicit web proxy. It connects to the Windows AD server through LDAPS, where the Windows server requires a client certificate to connect. The client certificate is configured in the CLI.

The endpoint PC connecting to the web server will first need to authenticate to the explicit web proxy before accessing the server.


While this example demonstrates an LDAP client certificate for an explicit proxy configuration, LDAP client certificates can be used in firewall authentication, transparent proxy, ZTNA, and where ever LDAP configurations are used on the FortiGate.

###### To configure a client certificate on the LDAP server:

1. 
                                                    Enable the explicit web proxy on port2: ```
config system interface
    edit "port2"
        set explicit-web-proxy enable
    next
end
```
2. 
                                                    Upload the client certificate to the FortiGate: ```
config vpn certificate local
    edit "Zach"
        set password **********
        set private-key <private key>
        set certificate <certificate>
    next
end
```
3. 
                                                    Configure the LDAP server settings: ```
config user ldap
    edit "ldaps"
        set server "172.16.200.57"
        set server-identity-check disable
        set cnid "CN"
        set dn "CN=Users,DC=ftnt,DC=com"
        set secure ldaps
        set port 636
        
```
**set client-cert-auth enable
        set client-cert "Zach"** next
end
4. 
                                                    Configure the authentication scheme: ```
config authentication scheme
    edit "1"
        set method basic
        set user-database "ldaps"
    next
end
```
5. 
                                                    Configure the authentication rule: ```
config authentication rule
    edit "1"
        set srcintf "port2"
        set srcaddr "all"
        set dstaddr "all"
        set active-auth-method "1"
    next
end
```
6. 
                                                    Configure the user group: ```
config user group
    edit "test"
        set member "ldaps"
    next
end
```
7. 
                                                    Configure the proxy policy with the user group: ```
config firewall proxy-policy
    edit 1
        set proxy explicit-web
        set dstintf "port3"
        set srcaddr "all"
        set dstaddr "all"
        set service "webproxy"
        set action accept
        set schedule "always"
        set srcaddr6 "all"
        set dstaddr6 "all"
        
```
**set groups "test"** set utm-status enable
        set ssl-ssh-profile "deep-inspection-clone"
        set av-profile "av"
    next
end

### Testing and verification

When traffic from the endpoint PC matches a policy and triggers authentication, the FortiGate starts the LDAPS TLS connection handshake with the Windows AD. The LDAPS server requests a client certificate to identify the FortiGate as a client. The FortiGate provides a configured client certificate, issued to zach.com, to the LDAPS server.

The following communication between the FortiGate and the LDAPS server shows the client certificate is sent by the FortiGate:
