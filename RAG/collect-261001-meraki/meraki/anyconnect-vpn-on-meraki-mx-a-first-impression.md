---
id: collect-261001-meraki/meraki/anyconnect-vpn-on-meraki-mx-a-first-impression
title: "anyconnect-vpn-on-meraki-mx-a-first-impression"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/anyconnect-vpn-on-meraki-mx-a-first-impression.md
source_anchor: ""
source_lines: [1, 104]
sha256: 4866d6609952c2f9a8dca29f0faaa4ed1230cc5fb7b092e64c24af3a5ba0a2e4
---

# anyconnect-vpn-on-meraki-mx-a-first-impression

The support for AnyConnect VPNs is probably one of the most wanted features for Meraki customers. It was first announced at Cisco Live 2015 (at least that is where I first heard of it) and after no more than six years the first public beta (v16.4) is available. Lets look at it.

My first try was with a Meraki Z3 which should be supported, but that device did not want to enroll a public certificate. Either it kept a self-signed-certificate or did not enable the AnyConnect server at all. Well, early Beta …

The next try was my MX68 (which I got from Meraki for my recognition as a Meraki All-Star, thanks again for that!). With this device the AnyConnect VPN was working.

### Configuration

The configuration is Meraki-easy as expected. For a basic setup we need:

1. Enable AnyConnect Client VPN
2. Change or accept the AnyConnect-port (default 443) and login-banner (default “You have successfully connected to client vpn.”)
3. Upload a client profile (optional, but I would always do so)
4. Configure the Authentication (RADIUS, Meraki Cloud or AD)
5. Configure the AnyConnect VPN subnet, Nameservers and DNS Suffix
6. Configure Split Tunneling

Thats all that has to be done and it is working.

### What is different to an AnyConnect implementation on the ASA

#### Certificate Enrollment

The certificate is automatically deployed for the DDNS hostname of the MX. It comes from the QuoVadis Root CA which should be trusted on all relevant systems and is valid for three months. The documentation says that it should auto-renew before it expires.

I expected that they implement an automatic Let’s Encrypt enrolment, but at least at the moment that is not possible. It’s also not possible to import your own certificate.

#### Crypto

This is a disappointment. On all my ASA implementations, I only enable TLS 1.2 with next-generation encryption and disable everything that has no Forward Secrecy (FS).

The MX also only uses TLS/DTLS 1.2 which is great. But there are a lot of non-FS algorithms enabled. SSLLabs only rates the VPN-Server with a “B” which is not state of the art any more. Having a default config (that can not be tuned) that gives a “B” is a little bit awkward nowadays.

#### Authentication

The default Authentication is AAA only. But you can also use double authentication (certificate and AAA) which I didn’t test yet.

There is no dedicated MFA-Config, but with RADIUS we can access any MFA server of our choice. After increasing the RADIUS timeout (default 3 seconds) MFA with the DUO authentication proxy directly worked like a charm.

The Authentication Protocol is “PAP_ASCII”, so there is no password-management for AnyConnect-users on the MX.

#### Authorization

On the ASA you can configure different IP subnets for different user groups, this is not possible with the MX and all users share the same VPN-subnet. It is also not possible to use a DHCP-server for address assignment.

In contrast to the legacy client VPN where all remote access users had to share the same “permit any” authorisation, with AnyConnect the RADIUS server can apply a group-policy to the session with the help of the RADIUS attribute “Filter-Id””.

Be carefull with the group-policy-names. If you configure the Filter-Id as “RA-USER”, and the RADIUS-server automatically appends an “.in” to the attribute, the group-policy has to be named “RA-USER.in” in the Meraki dashboard.

Same as with the AnyConnect pool, also the Split-Tunnel-config is global and can not be configured per user-group.

#### AnyConnect Profiles

As of now, only VPN-profiles can be pushed to the client. My first test did not work because the filename was like an FQDN (vpn.example.com.xml). After replacing the dots with dashes and only keeping the dot of the extension, it worked. The Meraki-Cloud added a second “.xml” so the profile name resulted in vpn-example-com.xml.xml but that does not harm anything.

There is no Profile-Editor embedded, the profiles have to be created in the standalone Profile-Editor or in a text editor.

Meraki-All-Star PhilipDAth created an online-version to generate a basic profile: https://www.ifm.net.nz/cookbooks/online-anyconnect-profile-editor.html

#### Redundancy

If the ASA is has multiple ISPs-interfaces, the ASA can be configured to accept connections on all interfaces. The MX only accepts AnyConnect-connections on the primary WAN-interface. But on the failure of the primary interface, the DDNS entry is updated to the IP of the secondary interface and that interface accepts the connections. Switching over took a couple of minutes which is not as good as configuring backup-servers in the profile, but at least we have basic redundancy.

#### AnyConnect versions

While the ASA supports a wide range of AnyConnect versions, the MX needs at least AnyConnect 4.8. But you should run a recent version anyhow.

The AnyConnect client can not be deployed from the MX as it is possible from the ASA. You need to implement any type of pre-installation.

#### Licensing

While in Beta, no extra license is needed, you even can download the AnyConnect client through the dashboard. But it is documented that the AnyConnect PLUS license is needed when this feature goes GA. I expect that we will have to connect the dashboard account to Cisco Smart licensing for that.

### Conclusion

The AnyConnect implementation on the Meraki MX is by far not as powerful as on the ASA. But probably no one expected that.

There are a couple of restrictions, but at least for me, I can probably arrange with it. I only hope that it does not take another couple of years for this release to become GA as most of my customers will not run Beta-code.

#### References:

**AnyConnect on the MX Appliance**

https://documentation.meraki.com/MX/AnyConnect_on_the_MX_Appliance

**AnyConnect Troubleshooting Guide**

https://documentation.meraki.com/MX/AnyConnect_on_the_MX_Appliance/AnyConnect_Troubleshooting_Guide

**AnyConnect on ASA vs. MX**

https://documentation.meraki.com/MX/AnyConnect_on_the_MX_Appliance/AnyConnect_on_ASA_vs._MX

**AnyConnect Client Download and Deployment**

https://documentation.meraki.com/MX/AnyConnect_on_the_MX_Appliance/Client_deployment

## Update

With recent Firmware-versions, Meraki now has the SSL Labs rating I was hoping for:

Is it possible to set a timeout – 8 hours – so the user has to authenticate again

There is no explicit option for that. Would be worth a test if the corresponding radius-attribute is used when sent from the server. But I doubt that.
