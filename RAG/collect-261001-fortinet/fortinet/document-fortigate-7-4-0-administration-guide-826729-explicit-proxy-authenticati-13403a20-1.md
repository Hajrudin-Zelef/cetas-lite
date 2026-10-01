---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-826729-explicit-proxy-authenticati-13403a20-1
title: "Explicit proxy authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-826729-explicit-proxy-authenticati-13403a20.md
source_anchor: ""
source_lines: [1, 162]
sha256: 4dac6c3815d51f9d6d44357b654672a1338fbc82da29b4b2c0821697d0e61a1a
---

# Explicit proxy authentication

FortiGate supports multiple authentication methods. This topic explains using an external authentication server with Kerberos as the primary and NTLM as the fallback.

###### To configure Explicit Proxy with authentication:

###### To enable and configure explicit web proxy in the GUI:

1. 
                                                    Go to *Network > Explicit Proxy* .
2. 
                                                    Enable *Explicit Web Proxy* .
3. 
                                                    Select *port2* as the*Listen on Interfaces* and set the*HTTP Port* to*8080* .
4. 
                                                    Configure the remaining settings as needed.
5. 
                                                    Click *Apply* .

###### To enable and configure explicit web proxy in the CLI:

```
config web-proxy explicit
    set status enable
    set ftp-over-http enable
    set socks enable
    set http-incoming-port 8080
    set ipv6-status enable
    set unknown-http-version best-effort
end
config system interface
    edit "port2"
        set vdom "vdom1"
        set ip 10.1.100.1 255.255.255.0
        set allowaccess ping https ssh snmp http telnet
        set type physical
        set explicit-web-proxy enable
        set snmp-index 12
        end
    next
end
```
                                            
                                            Since we are using an external authentication server with Kerberos authentication as the primary and NTLM as the fallback, Kerberos authentication is configured first and then FSSO NTLM authentication is configured.

For successful authorization, the FortiGate checks if user belongs to one of the groups that is permitted in the security policy.

|  | When configuring an LDAP connection to an Active Directory server, an administrator must provide Active Directory user credentials.  | 

###### To configure an authentication server and create user groups in the GUI:

1. 
                                                    Configure Kerberos authentication: 
  1. 
                                                            Go to *User & Authentication > LDAP Servers* .
  2. 
                                                            Click *Create New* .
  3. 
                                                            Set the following: Name ldap-kerberos Server IP 172.18.62.220 Server Port 389 Common Name Identifier cn Distinguished Name dc=fortinetqa,dc=local
  4. 
                                                            Click *OK*
2. 
                                                            
3. 
                                                    Define Kerberos as an authentication service. This option is only available in the CLI. For information on generating a keytab, see Generating a keytab on a Windows server.
4. 
                                                    Configure FSSO NTLM authentication: FSSO NTLM authentication is supported in a Windows AD network. FSSO can also provide NTLM authentication service to the FortiGate unit. When a user makes a request that requires authentication, the FortiGate initiates NTLM negotiation with the client browser, but does not process the NTLM packets itself. Instead, it forwards all the NTLM packets to the FSSO service for processing. 
  1. 
                                                            Go to *Security Fabric > External Connectors* .
  2. 
                                                            Click *Create New* and select*FSSO Agent on Windows AD* from the*Endpoint/Identity* category.
  3. 
                                                            Set the *Name* to*FSSO* ,*Primary FSSO Agent* to*172.16.200.220* , and enter a password.
  4. 
                                                            Click *OK* .
5. 
                                                            
6. 
                                                    Create a user group for Kerberos authentication: 
  1. 
                                                            Go to *User & Authentication > User Groups* .
  2. 
                                                            Click *Create New* .
  3. 
                                                            Set the *Name* to*Ldap-Group* , and*Type* to*Firewall* .
  4. 
                                                            In the *Remote Groups* table, click*Add* , and set the*Remote Server* to the previously created*ldap-kerberos* server.
  5. 
                                                            Click *OK* .
7. 
                                                            
8. 
                                                    Create a user group for NTLM authentication: 
  1. 
                                                            Go to *User & Authentication > User Groups* .
  2. 
                                                            Click *Create New* .
  3. 
                                                            Set the *Name* to*NTLM-FSSO-Group* ,*Type* to*Fortinet Single Sign-On (FSSO)* , and add*FORTINETQA/FSSO* as a member.
  4. 
                                                            Click *OK* .
9. 
                                                            

###### To configure an authentication server and create user groups in the CLI:

1. 
                                                    Configure Kerberos authentication: ```
config user ldap
    edit "ldap-kerberos"
        set server "172.18.62.220"
        set cnid "cn"
        set dn "dc=fortinetqa,dc=local"
        set type regular
        set username "CN=root,CN=Users,DC=fortinetqa,DC=local"
        set password *********
    next
end
```
2. 
                                                    Define Kerberos as an authentication service: ```
config user krb-keytab
    edit "http_service"
        set pac-data disable
        set principal "HTTP/FGT.FORTINETQA.LOCAL@FORTINETQA.LOCAL"
        set ldap-server "ldap-kerberos"
        set keytab "BQIAAABFAAIAEEZPUlRJTkVUUUEuTE9DQUwABEhUVFAAFEZHVC5GT1JUSU5FVFFBLkxPQ0FMAAAAAQAAAAAEAAEACKLCMonpitnVAAAARQACABBGT1JUSU5FVFFBLkxPQ0FMAARIVFRQABRGR1QuRk9SVElORVRRQS5MT0NBTAAAAAEAAAAABAADAAiiwjKJ6YrZ1QAAAE0AAgAQRk9SVElORVRRQS5MT0NBTAAESFRUUAAURkdULkZPUlRJTkVUUUEuTE9DQUwAAAABAAAAAAQAFwAQUHo9uqR9cSkzyxdzKCEXdwAAAF0AAgAQRk9SVElORVRRQS5MT0NBTAAESFRUUAAURkdULkZPUlRJTkVUUUEuTE9DQUwAAAABAAAAAAQAEgAgzee854Aq1HhQiKJZvV4tL2Poy7hMIARQpK8MCB//BIAAAABNAAIAEEZPUlRJTkVUUUEuTE9DQUwABEhUVFAAFEZHVC5GT1JUSU5FVFFBLkxPQ0FMAAAAAQAAAAAEABEAEG49vHEiiBghr63Z/lnwYrU="
    next
end
```
For information on generating a keytab, see Generating a keytab on a Windows server.
3. 
                                                    Configure FSSO NTLM authentication: ```
config user fsso
    edit "1"
        set server "172.18.62.220"
        set password *********
    next
end
```
4. 
                                                    Create a user group for Kerberos authentication: ```
config user group
    edit "Ldap-Group"
        set member "ldap" "ldap-kerberos"
    next
end
```
5. 
                                                    Create a user group for NTLM authentication: ```
config user group
    edit "NTLM-FSSO-Group"
        set group-type fsso-service
        set member "FORTINETQA/FSSO"
    next
end
```

Explicit proxy authentication is managed by authentication schemes and rules. An authentication scheme must be created first, and then the authentication rule.

###### To create an authentication scheme and rules in the GUI:

