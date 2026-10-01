---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2026-06"]
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1.md
source_anchor: ""
source_lines: [41, 90]
sha256: 49edb47adc5e999114fb46f236f73bd37c027bfa19f223bb7e1a98272fd63c0a
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1

This is the default option. With this option, the MX Appliance will enroll in a public trusted certificate using the DDNS hostname of the Meraki network. This publicly trusted certificate renews automatically. The DDNS hostname is not easy to remember, hence, it is highly recommended to use an AnyConnect profile to create a DDNS alias to simplify user interaction. For more information see, how to create a profile. Dashboard administrators do not have to worry about interacting with public CAs to get a signed certificate.
DDNS hostname is configurable on MX Appliances in Passthrough/VPN Concentrator mode when AnyConnect is enabled.
Automatic certificate generation is not supported for networks hosted on dashboard.meraki.cn or dashboard.meraki.ca or dashboard.meraki.in
If the MX is in High Availability (HA) mode with a virtual IP and behind a NAT device, we recommend using the custom certificates feature to enable you manage your certificates and DNS records. The automatic DDNS hostname certificates may not suffice.
Option 2: Custom hostname certificates
- Requires MX firmware version 16.16+.
- Custom hostname certificates do not renew automatically. Administrators will need to renew certificates manually in addition to managing their DNS record (to enable their hostname resolve to the MX IP on the Internet).
- Custom hostname certificates are supported in High Availability mode. Adminstrators are required to download Certificate Signing Requests (CSRs) and upload certificates for both Primary and Spare MX Appliances with the custom certificates Primary | Spare tab only visible when the MX Appliance is in High Availability mode. Each MX device in HA needs to have its' own CSR generated with a different Common Name (CN). If you wish to have a shared hostname between the two MX in the HA pair, you will need to put the shared hostname in the "Subject Alt Name (Hostnames)" field.
- Wildcard certificates are not supported.
- Changing the Server Certificate Generation Method back to the default option of Auto-Generated from Custom will cause the MX to delete the current uploaded custom certificate.
Warning: A factory reset or device replacement (RMA) will generate a new private key for that MX. As such, a new CSR will need to be generated, signed and the resulting certificates will need to be re-added under Security & SD-WAN > Configure > Client VPN > AnyConnect.
Note: Utilizing a custom hostname certificate will replace the auto generated DDNS hostname on the MX.
Administrators can generate a CSR, that can be signed by a public CA. The signed certificate should be uploaded to the MX Appliance via the dashboard. This option allows administrators to use a preferred hostname. For example, vpn.abc.com
Step 1. Generate CSR
Step 2. Get the CSR signed by a public CA of your choice
Step 3. Upload the signed certificate and CA chain from your CA*
 
*Note: A chain certificate must establish a full chain of trust back to a root CA, including any intermediate certificates that sit between the device and root certs. Such certificates are self-signed by the CA providing them, as the following example demonstrates:
Image courtesy of Mozilla Software Foundation and Wikipedia
Note that both the Subject Common Name and Issuer Common name are equal.
An incomplete or invalid chain of trust will result in the error "Failed verifying Device Cert with Cert Chain" being seen on dashboard when you go to upload the certificates.
Questions on how to obtain such a certificate should be brought up to whatever entity is providing the ones in question. For more details regarding certificate chains, see the following DigiCert Article.
Option 3: Self-signed certificates:
Only available for testing purposes.
 
Note on Certificates
Starting in June 2026, Public Certificate Authorities will no longer issue TLS certificates that contain both the serverAuth EKU and the clientAuth EKU in the same certificate. Browsers such as Chrome will prevent the use of the Client Authentication purpose in the Extended Key Usage (EKU) extension on server certificates. This restriction does not impact Cisco Secure Client (formerly AnyConnect), as it does not leverage the Chrome browser for certificate-based authentication. However, users configured for SAML authentication with external browser workflows may be affected if the external browser is Chrome and certificate-based authentication is required. For more information on these changes for Chrome, please refer to this link https://googlechrome.github.io/chromerootprogram/
Authentication Methods
AnyConnect supports authentication with either SAML, RADIUS, Active Directory, Meraki Cloud and Certificate authentication. For more details on authentication configuration, refer to AnyConnect Authentication Methods.
Client Routing
- Send all traffic through VPN: This is the same as full tunneling. All traffic from the client is sent over the VPN tunnel.
- Send all traffic except traffic going to these destinations: This is the same as full tunnel with exclusions, when configured, the client will send all traffic over the VPN except traffic destined for the configured subnet. This option is not supported on Android devices.
- Only send traffic going to these destinations: This is the same as spilt tunneling, when configured, the client will only send traffic destined for the configured subnet over the VPN. Every other traffic sent over the local network.
When selecting "Only send traffic going to these destinations" or "Send all traffic execpt traffic going to these destinations" at least one IP must be entered otherwise the following error will be presented:
"The AnyConnect traffic destination (Client Routing) must contain at least one valid subnet in CIDR format."
Dynamic Client Routing
Dynamic Client Routing is only supported on MX16.5+ firmware
Dynamic Client Routing is only supported on Windows and Mac platforms. It is not supported Linux or any mobile platforms.
Dynamic split tunneling/client routing allows for the specification of traffic that should be included or excluded in the VPN tunnel based on domain name rather than IP/Classless Inter-Domain Routing (CIDR) notation. This is critical for services that do not have dedicated or fixed IP addresses. Dynamic split tunneling can be used with or without the regular split tunneling feature.
Please note that every hostname configured is treated as a wildcard. For example, cisco.com is treated as *.cisco.com. Wildcards, for example, *.cisco.com cannot be configured on the dashboard. For the end user, routes are populated when a user tries to access the specified hostname. 
      
When selecting "Only send traffic going to these destination" or "Send all traffic except traffic going to these hostnames" at least one FQDN should be entered otherwise the following error message will be presented:
"The AnyConnect traffic destination (Dynamic Client Routing) must contain at least one valid destination."
Local LAN access
Local LAN access may be desired when full-tunneling is configured (Send all traffic through VPN), but users still require the ability to communicate with their local network. For example, a client that is allowed local LAN access while connected to the MX in full tunnel mode is able to print to a local printer at home, while other traffic flows through the tunnel.
To enable local LAN access, two things need to be done. Local LAN access will not work if both conditions are not satisfied.
1. Configure the MX: Select "Send all traffic except traffic going to these destinations" option on the dashboard and configure a 0.0.0.0/32 route. This will cause the AnyConnect client to automatically exclude traffic destined for the user's local network from going over the tunnel.
 
2. Configure the Client: Enable Allow local LAN Access on the AnyConnect Client. This can be enabled manually or via the AnyConnect profile.
                
