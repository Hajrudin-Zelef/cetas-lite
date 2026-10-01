---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-951346-68fbba0a-1
title: "document-fortigate-7-4-8-administration-guide-951346-68fbba0a"
domain: fortinet
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-951346-68fbba0a.md
source_anchor: ""
source_lines: [1, 80]
sha256: 579972a204ff5aa3906162c9f5f42a5cb31b0d15ac59cfefb3983e5d886a6997
---

# document-fortigate-7-4-8-administration-guide-951346-68fbba0a

SAML-based authentication for FortiClient remote access dialup IPsec VPN clients
SAML-based authentication for FortiClient remote access dialup IPsec VPN clients is now supported. This feature requires FortiClient 7.2.4 and FortiClient supports only using IKEv2. Two factor authentication using FortiToken push is also supported.
The FortiGate authd daemon has been enhanced to support SAML authentication and accepts local-in traffic from the FortiClient by the TCP port number configured in the auth-ike-saml-port setting (0 - 65535, default = 1001). Currently, this setting can only be configured in the CLI as follows:
config system global
    set auth-ike-saml-port <integer>
end
                                            This allows the FortiGate to act as a SAML service provider (SP) for IKEv2 FortiClient remote access IPsec VPN clients by forwarding the FortiClient's SAML request to the configured SAML identity provider (IdP) for user authentication.
The ike-saml-server setting enables a configured SAML server to listen on a FortiGate interface for SAML authentication requests from FortiClient remote access IPsec VPN clients. It must be configured on the interface that is directly receiving the SAML authentication requests from FortiClient. This setting can be configured in the CLI:
config system interface
    edit <name>
        set ike-saml-server <saml_server>
    next
end
                                            |  | The ike-saml-server setting must be configured on the interface that is the first point of contact for FortiClient traffic. For example, if FortiClient user SAML authentication traffic is always routed to the FortiGate on the WAN1 interface, then ike-saml-server must be configured for WAN1. If it is configured for WAN2, then the authentication traffic will not reach it on WAN1, even is the FortiGate allows traffic to flow from WAN1 to WAN2. | 
FortiClient will validate the certificate presented to it by FortiGate during its initial SAML connection. This certificate can be configured on the FortiGate from the GUI under User & Authentication > Authentication Settings > Certificate under User Authentication Options. To import the certificate on the FortiGate, see Import a certificate.
This certificate can also be configured in the CLI:
config user setting
     set auth-cert <certificate>
end
                                            To prevent an invalid server certificate prompt on FortiClient, the certificate's common name (CN) should match the IPsec VPN remote gateway's FQDN. If the certificate is signed by a custom Certificate Authority or one that is not well-known, the Certificate Authority's (CA) certificate should be imported in FortiClient endpoint's Trusted Root Certificate Authority store. For details on installing a CA certificate on the endpoint, see Installing certificates on the client.
SAML authentication flow with IPsec
The SAML Authentication flow when using IPsec where FortiGate is the Service Provider (SP), FortiAuthenticator, Entra ID, Okta, or another SAML IdP is the Identity Provider (IdP) and FortiClient is the web-browser as follows:
- 
                                                    When the FortiClient user clicks on Connect on FortiClient to connect to IPsec VPN Gateway (i.e. FortiGate), FortiClient first initiates a connection to FortiGate on the auth-ike-saml-port configured on FortiGate.
- 
                                                    The FortiGate sends a SAML Authentication Requests inside a redirect to FortiClient. The redirect consists of URLs to reach the IdP.
- 
                                                    FortiClient uses these redirects to send SAML Authentication Request to the IdP after which the login page on the IdP opens up.
- 
                                                    The user authenticates to the IdP using their SAML credentials configured on the IdP.
- 
                                                    The IdP sends a SAML Authentication Response that contains the user and group information in form of SAML Assertions to FortiClient.
- 
                                                    FortiClient sends a SAML Authentication Response to FortiGate.
- 
                                                    The FortiGate consumes the SAML Authentication Response and SAML Assertions after verifying the IdP using its IdP's certificate and provides FortiClient with a temporary token ID.
- 
                                                    FortiClient initiates IPsec tunnel and presents the token ID for authentication. Upon successful verification of token ID, IPsec tunnel establishes.
SAML configuration example with different IdPs
We will now see how to configure IPsec with SAML authentication using different IdPs on FortiGate and FortiClient using the following example:
The configuration steps on the FortiGate, different IdPs and FortiClient are as follows:
|  | Only Configuring SAML IdP and SAML SP is unique to individual IdPs. All other steps listed above are the same on FortiGate and FortiClient when using different IdPs. | 
Configuring IKE-SAML authentication port number on FortiGate
Configure a suitable TCP port number for SAML authentication (auth-ike-saml-port) used by FortiGate. This example uses port 9443 and the setting is configurable using CLI.
config system global
    set auth-ike-saml-port 9443
end 
                                            Configuring IPsec VPN certificate
In this step, using either the GUI or the CLI, configure the IPsec VPN certificate that is presented to FortiClient upon its initial connection.
To configure the IPsec VPN certificate in the GUI:
- 
                                                    Go to User & Authentication > Authentication Settings.
- 
                                                    Select the certificate from the Certificate dropdown menu. To import the certificate on FortiGate, see Import a certificate.
To configure the IPsec VPN certificate in the CLI:
If the certificate VPN_Certificate has already been imported on the FortiGate, then use the following CLI commands:
config user setting
    set auth-cert "VPN_Certificate"
end
                                            Configuring SAML IdP and SAML SP
The SAML configuration on SP (FortiGate) will vary based on selected IdPs from the list below. Select the preferred combination of SP and IdP as per your requirement from the following list.
- 
                                                    Configure FortiAuthenticator as SAML IdP and FortiGate as SAML SP
- 
                                                    Configure Microsoft Entra ID as SAML IdP and FortiGate as SAML SP
|  | SAML IdPs other than FortiAuthenticator or Microsoft Entra ID can be used. Please refer to the documentation of the respective SAML IdP for details. | 
Configuring IPsec IKEv2 on FortiGate
Configuring Remote access VPN on FortiGate enables FortiClient to connect to the IPsec VPN gateway configured on FortiGate. FortiClient 7.2.4 GA and above supports only IKEv2 for SAML authentication. The example discussed uses full-tunnel IPsec VPN. For split-tunnel configuration and other advanced configurations as per your requirement, see Remote access.
To configure IPsec VPN on FortiGate with FortiClient as the dialup client:
- 
                                                    Go to VPN > IPsec Tunnels.
- 
                                                    Click Create New > IPsec Tunnel. The VPN Creation Wizard is displayed.
- 
                                                    Enter the Name as FCT_SAML. This example does not use the VPN wizard for the IPsec tunnel configuration but rather configures a Custom IPsec tunnel.
- 
                                                    Configure the Template type as Custom.
- 
                                                    Click Next.
- 
