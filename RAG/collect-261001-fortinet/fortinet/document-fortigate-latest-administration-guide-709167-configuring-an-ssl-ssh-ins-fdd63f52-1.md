---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52-1
title: "document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52.md
source_anchor: ""
source_lines: [1, 56]
sha256: 47944a22f1fd1a0426a55d736b1eed50f8eb1784e727d09d40e980320f90a286
---

# document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52

Configuring an SSL/SSH inspection profile
Configuring an SSL/SSH inspection profile
The custom-deep-inspection profile can be edited or new SSL/SSH inspection profiles can be configured to be used in firewall policies.
To configure an SSL/SSH inspection profile in the GUI:
- 
                                                    Go to Security Profiles > SSL/SSH Inspection and click Create New.
- 
                                                    Configure the following settings: Name Enter a unique name for the profile. Comments Enter a comment (optional). SSL Inspection Options Enable SSL Inspection of Enable SSL inspection of: 
  - 
                                                                            Multiple Clients Connecting to Multiple Servers: Use this option for generic policies where the destination is unknown. This is normally used when inspecting outbound internet traffic. Other SSL Inspection Options become available to configure if this option is selected.
  - 
                                                                            Protecting SSL Server: Use this option when setting up a profile customized for a specific SSL server with a specific certificate. Define the certificate using the Server certificate field. See Protecting an SSL server for more information.
 Inspection method Define the inspection method: 
  - 
                                                                            SSL Certificate Inspection: Only inspects the certificate, by way of the headers up to the SSL/TLS layer, and not the contents of the traffic. Allows Encrypted Client Hello (ECH) to be allowed or blocked.
  - 
                                                                            Full SSL Inspection: Inspects the SSL/TLS encrypted traffic payload. See Deep inspection.
 CA certificate Use the dropdown menu to select one of the installed certificates for the inspection of the packets. Click Download to save the certificate. Blocked certificates Block or allow potentially malicious certificates. Select View Blocked Certificates for a detailed list of blocked certificates, including the listing reason and date. Untrusted SSL certificates Configure the action to take when a server certificate is not issued by a trusted CA. 
  - 
                                                                            Allow: Allow the untrusted server certificate. This is the default value.
  - 
                                                                            Block: Block the session.
  - 
                                                                            Ignore: This option is for Full SSL inspection only. It re-signs the server certificate as trusted. When configured in the GUI for certificate inspection it has no effect and the setting is not saved.
 Click View Trusted CAs List to see a list of the factory bundled and user imported CAs that are trusted by the FortiGate. Server certificate SNI check Check the SNI in the hello message with the CN or SAN field in the returned server certificate: 
  - 
                                                                            Enable: If it is mismatched, use the CN in the server certificate for URL filtering.
  - 
                                                                            Strict: If it is mismatched, close the connection.
  - 
                                                                            Disable: Server certificate SNI check is disabled.
 Enforce SSL cipher compliance Enable/disable SSL cipher compliance. This option is for Full SSL inspection only. Enforce SSL negotiation compliance Enable/disable SSL negotiation compliance. This option is for Full SSL inspection only. RPC over HTTPS Enable/disable inspection of Remote Procedure Calls (RPC) over HTTPS traffic. This option is for Full SSL inspection only. MAPI over HTTPS Enable/disable inspection of Messaging Application Programming Interface (MAPI) over HTTPS traffic. Protocol Port Mapping Inspect all ports Inspect all ports with the IPS engine by enabling Inspect all ports. Choose whether to disable DNS over TLS. HTTPS, SMTPS, POP3S, IMAPS, FTPS If Inspect all ports is disabled, specify the port through which traffic will be inspected in the field next to the listed protocols. See Inspecting all ports for details. DNS over TLS Enable/disable inspection of DNS over TLS. Encrypted Client Hello Allow or block TLS connections that use Encrypted Client Hello (ECH). HTTP/3 When Inspection method is set to SSL Certificate Inspection or HTTPS is disabled, HTTP/3 is set to Bypass and cannot be changed. When Inspection method is set to Full SSL Inspection, choose whether to Inspect, Bypass, or Block HTTP/3. DNS over QUIC Available when Inspection method is set to Full SSL Inspection. Choose whether to Inspect, Bypass, or Block DNS over QUIC. Exempt from SSL Inspection These options are for Full SSL inspection only. Use the menus in this section to specify any reputable websites, FortiGuard Web Categories, or addresses that will be exempt from SSL inspection: 
  - 
                                                                            Reputable Websites: Enable this option to exempt any websites identified by FortiGuard as reputable.
  - 
                                                                            Web Categories: The categories of Finance and Banking, Health and Wellness, and Personal Privacy have been added by default. These categories are the most likely to have applications that will require a specific certificate.
  - 
                                                                            Addresses: These can be any of the address objects that have an interface of any.
  - 
                                                                            Log SSL exemptions: Enable this option to log all SSL exemptions.
 See Exempt web sites from deep inspection for more information. SSH Inspection Options SSH deep scan Enable/disable SSH protocol packet deep scanning capabilities. SSH port will become available if SSH deep scan is enabled. SSH port Define what ports will search for SSH protocol packets: 
  - 
                                                                            Any: Select this option to search all traffic regardless of service or TCP/IP port for packets that conform to the SSH protocol.
  - 
                                                                            Specify: Select this option and enter the port number to restrict the search for SSH protocol packets to the TCP/IP port number specified. This is not as comprehensive but it is easier on the performance of the firewall.
 Common Options Invalid SSL certificates Allow or block the passing of traffic in invalid certificates. Additional common options that provide more granularity with actions for different types of invalid SSL certificates will become available if Invalid SSL certificates is set to Custom: 
  - 
                                                                            Expired certificates: Action to take when the server certificate is expired. The default action is block.
  - 
                                                                            Revoked certificates: Action to take when the server certificate is revoked. The default action is block.
  - 
                                                                            Validation timed-out certificates: Action to take when the server certificate validation times out. For certificate inspection, the default action is allow. For deep inspection, the default action is Keep Untrusted & Allow.
  - 
                                                                            Validation failed certificates: Action to take when the server certificate validation fails. The default action is block.
 For deep inspection, the above options have the following actions: 
  - 
