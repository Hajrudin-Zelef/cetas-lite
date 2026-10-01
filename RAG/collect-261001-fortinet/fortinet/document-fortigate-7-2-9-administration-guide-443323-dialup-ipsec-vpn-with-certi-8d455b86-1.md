---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-9-administration-guide-443323-dialup-ipsec-vpn-with-certi-8d455b86-1
title: "Dialup IPsec VPN with certificate authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-9-administration-guide-443323-dialup-ipsec-vpn-with-certi-8d455b86.md
source_anchor: ""
source_lines: [1, 166]
sha256: e9080483bca9aa2d156293be54d616d1ba7010d7d5c258636921e2d5a52920e6
---

# Dialup IPsec VPN with certificate authentication

# Dialup IPsec VPN with certificate authentication

In a dialup IPsec VPN setup, a company may choose to use X.509 certificates as their authentication solution for remote users. This method includes the option to verify the remote user using a user certificate, instead of a username and password. This method can be simpler for end users.

Administrators need to issue unique user certificates to each user for remote access management. The user certificate can be verified by the subject field, common name, or the principal name in the Subject Alternative Name (SAN) field.

## Subject field verification

This is the basic method that verifies the subject string defined in the PKI user setting matches a substring in the subject field of the user certificate. For example:

```
config user peer
    edit "tgerber"
        set ca "CA_Cert_2"
        set subject "CN=tgerber"
    next
end
```
                                            
                                            In this method, administrators can define the CN string to match the common name (CN) in the subject field of the certificate. For example:

```
config user peer
    edit "tgerber"
        set ca "CA_Cert_2"
        set cn "tgerber"
    next
end
```
                                            The matching certificate looks like the following:


A PKI user must be created on the FortiGate for each remote user that connects to the VPN with a unique user certificate.

## Principal name with LDAP integration

In this method, the PKI user setting references an LDAP server. When `ldap-mode` is set to `principal-name`, the UPN in the user certificate's SAN field is used to look up the user in the LDAP directory. If a match is found, then authentication succeeds. For example:

```
config user peer
    edit "ldap-peer"
        set ca "CA_Cert_2"
        set ldap-server "WIN2K16-KLHOME-LDAPS"
        set ldap-mode principal-name
    next
end
```
                                            The matching certificate looks like the following:


This method is more scalable because only one PKI user needs to be created on the FortiGate. Remote users connect with their unique user certificate that are matched against users in the LDAP server.

## Certificate management

Dialup IPsec VPN with certificate authentication requires careful certificate management planning. Assuming that a company's private certificate authority (CA) is used to generate and sign all the certificates, the following certificates are needed:

| Certificate type | Description | 
|---|---|
| Server certificate | The server certificate is used to identify the FortiGate IPsec dialup gateway. A CSR can be generated on the FortiGate and signed by the CA, or the CA can generate the private and public keys and export the certificate package to the FortiGate. | 
| User certificate | The user certificate is generated and signed by the CA with unique CNs in the subject field and/or unique Principal Names in the SAN field. They are used to identify the user that is connecting to the VPN. User certificates must be installed on client machines. | 
| CA certificate | The root CA certificate, and any subordinate CA that signed the actual user and server certificates, must be imported into the FortiGate and client machines. The CA certificate is used to verify the certificate chain of the server and user certificates. | 

## Example

In this example, a dialup IPsec VPN tunnel is configured with certificate authentication using the subject field verification method and the LDAP integration method.


The company CA, named root CA, signs all the server and user certificates. The user, tgerber@klhome.local, has a user certificate signed by root CA installed on their endpoint. The corresponding user account is also present under the company's Active Directory.

There are five major steps to configure this example:

The server certificate and CA certificate need to be imported into the FortiGate.

###### To import the server certificate:

1. Go to *System > Certificates* and select*Import > Local Certificate* .
2. For *Type* , select*PKCS #12 Certificate* .
3. Upload the key file exported from the CA and enter the password.
4. Click *OK* . The certificate now appears in the*Local Certificate* section.

###### To import the CA certificate:

1. Go to *System > Certificates* and select*Import > CA Certificate* .
2. For *Type* , select*File* .
3. Upload the CA certificate (usually a .CRT file). This certificate only contains the public key.
4. Click *OK* . The certificate now appears in the*Remote CA Certificate* section.

|  | If any subordinate CA is involved in signing the certificates, you need to import its certificate. | 

FortiGate PKI users do not appear in the GUI until at least one PKI user has been created in the CLI. The following instructions create the PKI users in the CLI.

###### To configure PKI users for subject field verification:

1. Create the PKI user and choose the CA certificate that was imported (if the certificate was signed by a subordinate CA, choose the subordinate CA's certificate):```
config user peer
    edit "tgerber"
        set ca "CA_Cert_2"
        set subject "CN=tgerber"
    next
end
```
For an example of CN field matching, see Common name verification.
2. Create additional users as needed.
3. Place the users into a peer group:```
config user peergrp
    edit "pki-users"
        set member "tgerber" <user> ... <user>
    next
end
```

###### To configure PKI users for LDAP integration:

1. Configure the LDAP server that users connect to for authentication:```
config user ldap
    edit "WIN2K16-KLHOME-LDAPS"
        set server "192.168.20.6"
        set cnid "sAMAccountName"
        set dn "dc=KLHOME,dc=local"
        set type regular
        set username "KLHOME\\Administrator"
        set password ************
        set secure ldaps
        set ca-cert "CA_Cert_1"
        set port 636
    next
end
```
2. Configure the PKI user to reference the LDAP server using the CA certificate that was imported:```
config user peer
    edit "ldap-peer"
        set ca "CA_Cert_2"
        set ldap-server "WIN2K16-KLHOME-LDAPS"
        set ldap-mode principal-name
    next
end
```
3. Place the user into a peer group:```
config user peergrp
    edit "pki-ldap"
        set member "ldap-peer"
    next
end
```

To configure the VPN, the address objects must be defined first so they can be used in the VPN and policy configurations. In this example, the VPN is configured in custom mode to define the authentication settings.

###### To configure the address objects:

1. Create the address range for the dialup clients:
  1. Go to *Policy & Objects > Addresses* and click*Create New > Address* .
  2. For *Name* , enter*remote-user-range* .
  3. For *Type* , select*IP Range* and enter*172.18.200.10-172.18.200.99* in the*IP Range* field.
  4. Click *OK* .
2. Go to 
3. Create the address subnet for the destination 192.168.20.0/24:
  1. Click *Create New > Address* .
  2. For *Name* , enter*192.168.20.0* .
  3. For *Type* , select*Subnet* and enter*192.168.20.0/24* in the*IP/Netmask* field.
  4. Click *OK* .
4. Click 

###### To configure the IPsec dialup tunnel:

