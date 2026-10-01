---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-b2849ef6
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-b2849ef6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-b2849ef6.md
source_anchor: ""
source_lines: [1, 48]
sha256: 10592366918a11cced17623fffc69c1eba039643179f871831cdfe524e33d525
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-b2849ef6

Prerequisites for SNMP
Supported SNMP Versions
This software release supports the following SNMP versions:
- 
                                 
                                 SNMPv1—The Simple Network Management Protocol, a Full Internet Standard, defined in RFC 1157.
- 
                                 
                                 SNMPv2C replaces the Party-based Administrative and Security Framework of SNMPv2Classic with the community-string-based Administrative Framework of SNMPv2C while retaining the bulk retrieval and improved error handling of SNMPv2Classic. It has these features: 
  - 
                                       
                                       SNMPv2—Version 2 of the Simple Network Management Protocol, a Draft Internet Standard, defined in RFCs 1902 through 1907.
  - 
                                       
                                       SNMPv2C—The community-string-based Administrative Framework for SNMPv2, an Experimental Internet Protocol defined in RFC 1901.
- 
                                       
                                       
- 
                                 
                                 SNMPv3—Version 3 of the SNMP is an interoperable standards-based protocol defined in RFCs 2273 to 2275. SNMPv3 provides secure access to devices by authenticating and encrypting packets over the network and includes these security features: 
  - 
                                       
                                       Message integrity—Ensures that a packet was not tampered with in transit.
  - 
                                       
                                       Authentication—Determines that the message is from a valid source.
  - 
                                       
                                       Encryption—Mixes the contents of a package to prevent it from being read by an unauthorized source.
- 
                                       
                                       
| Note | To select encryption, enter the priv keyword. | 
Both SNMPv1 and SNMPv2C use a community-based form of security. The community of managers able to access the agent’s MIB is defined by an IP address access control list and password.
SNMPv2C includes a bulk retrieval function and more detailed error message reporting to management stations. The bulk retrieval function retrieves tables and large quantities of information, minimizing the number of round-trips required. The SNMPv2C improved error-handling includes expanded error codes that distinguish different kinds of error conditions; these conditions are reported through a single error code in SNMPv1. Error return codes in SNMPv2C report the error type.
SNMPv3 provides for both security models and security levels. A security model is an authentication strategy set up for a user and the group within which the user resides. A security level is the permitted level of security within a security model. A combination of the security level and the security model determine which security method is used when handling an SNMP packet. Available security models are SNMPv1, SNMPv2C, and SNMPv3.
The following table identifies characteristics and compares different combinations of security models and levels:
| Table 1. SNMP Security Models and Levels |  |  |  |  | 
|---|---|---|---|---|
| Model | Level | Authentication | Encryption | Result | 
|---|---|---|---|---|
| SNMPv1 | noAuthNoPriv | Community string | No | Uses a community string match for authentication. | 
| SNMPv2C | noAuthNoPriv | Community string | No | Uses a community string match for authentication. | 
| SNMPv3 | noAuthNoPriv | Username | No | Uses a username match for authentication. | 
| SNMPv3 | authNoPriv | Message Digest 5 (MD5) or Secure Hash Algorithm (SHA) | No | Provides authentication based on the HMAC-MD5 or HMAC-SHA algorithms. | 
| SNMPv3 | authPriv | MD5 or SHA | Data Encryption Standard (DES) or Advanced Encryption Standard (AES) | Provides authentication based on the HMAC-MD5 or HMAC-SHA algorithms. Allows specifying the User-based Security Model (USM) with these encryption algorithms:  | 
You must configure the SNMP agent to use the SNMP version supported by the management station. Because an agent can communicate with multiple managers, you can configure the software to support communications using SNMPv1, SNMPv2C, or SNMPv3.
