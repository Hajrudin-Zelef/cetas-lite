---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-6-4-15-administration-guide-173316-add-ldap-user-au-138184ea
title: "Add LDAP user authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-6-4-15-administration-guide-173316-add-ldap-user-au-138184ea.md
source_anchor: ""
source_lines: [1, 63]
sha256: bc0de470e7797f4c4054f5d2f3122880a9d6d5d0c4caa0f774becfb6383596df
---

# Add LDAP user authentication

This configuration adds LDAP user authentication to the FortiClient dialup VPN configuration (FortiClient as dialup client). You must have already generated and exported a CA certificate from your AD server.

###### To configure LDAP user authentication using the GUI:

1. Import the CA certificate into FortiGate:
        
  1. Go to *System > Certificates* .If the *Certificates* option is not visible, enable it in*Feature Visibility* . See Feature visibility for details.
  2. Click *Import > CA Certificate* .
  3. Set *Type* to*File* .
  4. Click *Upload* then find and select the certificate file.
  5. Click *OK* .The CA certificate now appears in the list of *External CA Certificates* . In this example, it is called*CA_Cert_1* .
  6. Optionally, rename the system generated *CA_Cert_1* to something more descriptive:```
config vpn certificate ca
    rename CA_Cert_1 to LDAPS-CA
end
```
2. Go to 
3. Configure the LDAP user:
        
  1. Go to *User & Authentication > LDAP Servers* and click*Create New* .
  2. Set *Name* to*ldaps-server* and specify*Server IP/Name* .
  3. Specify *Common Name Identifier* and*Distinguished Name* .
  4. Set *Bind Type* to*Regular* .
  5. Specify *Username* and*Password* .
  6. Enable *Secure Connection* and set*Protocol* to*LDAPS* .
  7. For *Certificate* , select*LDAP server CA**LDAPS-CA* from the list.
  8. Click *OK* .
4. Go to 
5. Add the LDAP user to the user group:
        
  1. Go to *User & Authentication > User Groups* and edit the*vpngroup* group.
  2. In *Remote Groups* , click*Add* to add the*ldaps-server* remote server.
  3. Click *OK* .
6. Go to 

###### To configure LDAP user authentication using the CLI:

1. Import the CA certificate using the GUI.
2. Configure the LDAP user:```
config user ldap
    edit "ldaps-server"
        set server "172.20.120.161"
        set cnid "cn"
        set dn "cn=Users,dc=qa,dc=fortinet,dc=com"
        set type regular
        set username "CN=Administrator,cn=users,DC=qa,DC=fortinet,DC=com"
        set password **********
        set group-member-check group-object
        set secure ldaps
        set ca-cert "LDAPS-CA"
        set port 636
    next
end
```
3. Add the LDAP user to the user group:```
config user group
    edit "vpngroup" 
        append member "ldaps-server"
    next 
end
```
