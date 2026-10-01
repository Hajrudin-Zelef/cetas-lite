---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-173316-add-ldap-user-authenticatio-11959b23
title: "document-fortigate-7-4-0-administration-guide-173316-add-ldap-user-authenticatio-11959b23"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-173316-add-ldap-user-authenticatio-11959b23.md
source_anchor: ""
source_lines: [1, 51]
sha256: 36257274819d88b008742230b3252a9e7ee06a8e850376e8f26f65299808300a
---

# document-fortigate-7-4-0-administration-guide-173316-add-ldap-user-authenticatio-11959b23

Add LDAP user authentication
Add LDAP user authentication
This configuration adds LDAP user authentication to the FortiClient dialup VPN configuration (FortiClient as dialup client). You must have already generated and exported a CA certificate from your AD server.
To configure LDAP user authentication using the GUI:
- Import the CA certificate into FortiGate:
        
  - Go to System > Certificates.If the Certificates option is not visible, enable it in Feature Visibility. See Feature visibility for details.
  - Click Import > CA Certificate.
  - Set Type to File.
  - Click Upload then find and select the certificate file.
  - Click OK.The CA certificate now appears in the list of External CA Certificates. In this example, it is called CA_Cert_1.
  - Optionally, rename the system generated CA_Cert_1 to something more descriptive:config vpn certificate ca
    rename CA_Cert_1 to LDAPS-CA
end
- Go to System > Certificates.
- Configure the LDAP user:
        
  - Go to User & Authentication > LDAP Servers and click Create New.
  - Set Name to ldaps-server and specify Server IP/Name.
  - Specify Common Name Identifier and Distinguished Name.
  - Set Bind Type to Regular.
  - Specify Username and Password.
  - Enable Secure Connection and set Protocol to LDAPS.
  - For Certificate, select LDAP server CA LDAPS-CA from the list.
  - Click OK.
- Add the LDAP user to the user group:
        
  - Go to User & Authentication > User Groups and edit the vpngroup group.
  - In Remote Groups, click Add to add the ldaps-server remote server.
  - Click OK.
To configure LDAP user authentication using the CLI:
- Import the CA certificate using the GUI.
- Configure the LDAP user:config user ldap
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
- Add the LDAP user to the user group:config user group
    edit "vpngroup" 
        append member "ldaps-server"
    next 
end
