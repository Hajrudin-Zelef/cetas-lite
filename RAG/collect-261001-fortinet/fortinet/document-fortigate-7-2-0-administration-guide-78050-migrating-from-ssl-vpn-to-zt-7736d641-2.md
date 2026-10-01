---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-0-administration-guide-78050-migrating-from-ssl-vpn-to-zt-7736d641-2
title: "Migrating from SSL VPN to ZTNA"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-0-administration-guide-78050-migrating-from-ssl-vpn-to-zt-7736d641.md
source_anchor: ""
source_lines: [156, 288]
sha256: 5bf71d393fceac77bc86f64c255b0e981291b61f38d5260091a7d679d9af180c
---

# Migrating from SSL VPN to ZTNA

6. 
                                                    Select the *Default certificate* . Clients will be presented with this certificate when they connect to the access proxy VIP.
7. 
                                                    Add server mapping: 
  1. 
                                                            In the *Service/server mapping* table, click*Create New* .
  2. 
                                                            Set *Service* to*HTTPS* .
  3. 
                                                            Set *Virtual Host* to*Any Host* .
  4. 
                                                            Configure the path as needed. For example, to map to *winserver.fgdocs.com/fortigate* , enter*/fortigate* .
  5. 
                                                            Add a server: 
    1. 
                                                                    In the *Servers* table, click*Create New* .
    2. 
                                                                    Set *IP* to*192.168.20.6* .
    3. 
                                                                    Set *Port* to*443* .
    4. 
                                                                    Click *OK* .
  6. 
                                                                    
  7. 
                                                            Click *OK* .
8. 
                                                            
9. 
                                                    Click *OK* .

The authentication scheme defines the authentication method that is applied. In this example, basic HTTP authentication is used so that users are prompted for a username and password the first time that they connect to a website through the HTTPS access proxy. The LDAP server defined for the SSL VPN configurations can be reused here.

###### To configure an authentication scheme:

1. 
                                                    Go to *Policy & Objects > Authentication Rules* and click*Create New > Authentication Scheme* .
2. 
                                                    Set the name to *ZTNA-Auth-scheme* .
3. 
                                                    Set *Method* to*Basic* .
4. 
                                                    Set *User database* to*Other* and select*WIN2K16-KLHOME-LDAPS* as the LDAP server.
5. 
                                                    Click *OK* .

The authentication rule defines the proxy sources and destination that require authentication, and what authentication scheme is applied. In this example, active authentication through the basic HTTP prompt is used and applied to all sources.

###### To configure an authentication rule:

1. 
                                                    Go to *Policy & Objects > Authentication Rules* and click*Create New > Authentication Rule* .
2. 
                                                    Set the name to *ZTNA-Auth-rule* .
3. 
                                                    Set *Source Address* to*all* .
4. 
                                                    Set *Protocol* to*HTTP* .
5. 
                                                    Enable *Authentication Scheme* and select*ZTNA-Auth-scheme* .
6. 
                                                    Click *OK* .

A user or user group must be applied to the ZTNA rule used to control user access. The authenticated user from the authentication scheme and rule must match the user or user group in the ZTNA rule. The user group, *KLHOME-ALLOWED-VPN*, defined  in the SSL VPN configurations is reused in this example. The ZTNA tag, *Malicious-File-Detected*, is used to define a rule to deny access when the connecting device has the malicious file detected.

###### To configure ZTNA rules to allow and deny traffic based on ZTNA tags:

1. 
                                                    Go to *Policy & Objects > ZTNA* and select the*ZTNA Rules* tab.
2. 
                                                    Create a rule to deny traffic: 
  1. 
                                                            Click *Create New* .
  2. 
                                                            Set *Name* to*ZTNA-Deny-malicious* .
  3. 
                                                            Set *Incoming Interface* to*port1* .
  4. 
                                                            Set *Source* to*all* , then click the*+* and from the*User* tab, select the*KLHOME-ALLOWED-VPN* group.
  5. 
                                                            Add the ZTNA tag *Malicious-File-Detected* .This tag is dynamically retrieved from EMS when the Zero Trust tagging rule is first created.
  6. 
                                                            Select the ZTNA server *WIN2K16-P1* .
  7. 
                                                            Set *Action* to*DENY* .
  8. 
                                                            Enable *Log Violation Traffic* .
  9. 
                                                            Click *OK* .
3. 
                                                            
4. 
                                                    Create a rule to allow traffic: 
  1. 
                                                            Click *Create New* .
  2. 
                                                            Set *Name* to*proxy-WIN2K16-P1* .
  3. 
                                                            Set *Incoming Interface* to*port1* .
  4. 
                                                            Set *Source* to*all* , then click the*+* and from the*User* tab, select the*KLHOME-ALLOWED-VPN* group. The*Source* can also be set to specific IP addresses to only allow those addresses to connect to this HTTPS access proxy.
  5. 
                                                            Add the ZTNA tag *Low* .
  6. 
                                                            Select the ZTNA server *WIN2K16-P1* .
  7. 
                                                            Set *Action* to*ACCEPT* .
  8. 
                                                            Configure the remaining options as needed.
  9. 
                                                            Click *OK* .
5. 
                                                            
6. 
                                                    In the *ZTNA Rules* list, make sure that the deny rule (*ZTNA-Deny-malicious* ) is above the allow rule (*proxy-WIN2K16-P1* ).

## Testing the connection

Once ZTNA is configured, connect to the FortiGate access proxy using an endpoint that is registered to EMS. The user should be prompted for their device certificate, username, and password the first time they connect. Once they have authenticated and they pass the security posture checks, they will be allowed to access the website.

See ZTNA HTTPS access proxy example and ZTNA HTTPS access proxy with basic authentication example for sample verifications and results.

## Disabling the SSL VPN

Once testing is complete and the ZTNA servers and policies are configured, the users can be migrated to using ZTNA. Use the following checklist to verify if the remote users are ready to migrate:

1. The users have installed a supported FortiClient version and have installed the ZTNA module.
2. The endpoints can register to FortiClient EMS.
3. If using a TCP forwarding access proxy, ensure that ZTNA rules are either pushed from FortiClient EMS, or the users know how to configure them manually.

Next, SSL VPN access can be disabled in a phased approach by disabling SSL VPN firewall policies that allow access to resources that are accessible using ZTNA.

Once all applications and resources have been migrated, the SSL VPN can be disabled entirely by going to *VPN > SSL-VPN Settings*, and deselecting the *Enable SSL-VPN* toggle.
