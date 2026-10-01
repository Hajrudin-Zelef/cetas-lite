---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647-1
title: "site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["China"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647.md
source_anchor: ""
source_lines: [1, 110]
sha256: a1b1f2f0211eb550b3430b233dd3c71e928660310a67c3631a9feacc94cbf166
---

# site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647

After much troubleshooting, I was able to successfully establish OpenVPN “Site-to-Site” VPN tunnel between my primary OPNsense firewall and an edge Ubiquiti EdgeRouter (ERLite-3). This guide will show you how to setup the OPNsense side using the new “instances” configuration and how to setup your EdgeRouter OpenVPN profile manually.
The Problem
There are three main challenges at play here: 1) preshared keys are being depreciated and 2) finding interoperability between major versions of OpenVPN ciphers and 3) the Ubiquiti documentation, guides and knowledge around setting up site-to-site VPNs all focus on preshared keys and there are limitation in the EdgeOS gui which does not pass validation thus we must pivot to a traditional OpenVPN profile configuration on the EdgeRouter.
In summary, the EdgeRouter interface does not validate/support the options we need to connect to a newer version of OpenVPN.
Configuration
Assumptions
This article assumes you have your own:
- Root CA
- Certificate & key for Site A acting as the OpenVPN server
- Certificate & key for Site B acting as the OpenVPN client
- Both Site A & Site B certificates are signed by your Root CA
This article will not help you create a Root CA, issuing and signing certificates. I will say there is a nifty tool called XCA which I highly recommend for helping you do this and keep track of all your certificates. Here is a XCA How-to Tutorial. You’ll want to use the TLS_server profile for Site A and TLS_client Profile for Site B within XCA.
Additionally, I assume you’ve generated a TLS-CRPYT key. You can learn to do that here: EasyRSA3 OpenVPN Howto. This example uses a TLS-CRYPT key not the TLS-CRYPT-v2 key. TLS-CRYPT-v2 is not supported in OpenVPN 2.4 (EdgeRouter). This (TLS-CRYPT-v2) was introduced in OpenVPN 2.5.
Contextual Notes
This tutorial is focused around OpenVPN operating via UDP and not TCP. You can change everything to be TCP if you wish, just make sure the ports and protocols are consistent throughout your configuration. I also opt to operate on higher port ranges to along with firewall inbound source address filtering. Defining your source (Site B) IP address or through an Alias you can define a DDNS FQDN to limit the scope of who can talk inbound to that port. This is good for the security-minded folk. If you don’t care that Russia or China knocks on your OpenVPN front door make Step 9, Source: any!
Topology
OPNsense “Site A”
- VPN -> OpenVPN -> Instances [new] -> Static Keys (tab) -> Click “+”
- Add your generated TLS-CRYPT key:
  - Description: OpenVPN TLS-CRYPT Key
  - Mode: crypt (Encrypt and authenticate all control channel packets)
  - Static Key: <copy/paste key here with —–BEGIN OpenVPN Static key V1—– and —–END OpenVPN Static key V1—–>
- Press “Save.
- VPN -> OpenVPN -> Instances [new] -> Click “+”
- Toggle “Advance Mode”
  - Role: Server
  - Description: OpenVPN Site-to-Site
  - Protocol: UDP (IPv4)
  - Port: <23222>
  - Bind Address: <leave blank>
  - Type: tun
  - (Optional) Verbosity: 4 (Normal)
  - (Optional) Keep alive interval: 60
  - (Optional) Keep alive timeout: 300
  - Server (IPv4): <10.255.0.0/24>
  - Topology: subnet
  - Trust: <Site B Certificate>
  - Certificate Authority: <Root CA>
  - Verify Client Certificate: required
  - Certificate Depth: Two (Client+Intermediate+Server)
  - TLS static key: <select OpenVPN key from step 1>
  - Auth: SHA256 (256-bit)
  - Data Ciphers: AES-256-GCM, AES-128-GCM
  - Data Ciphers Fallback: Nothing selected
  - Authentication: Nothing selected
  - Local Network: <10.1.0.0/24>
  - Remote Network: <10.2.0.0/24>
  - .. Leave the rest of the settings default …
- Press “Save”
- VPN -> OpenVPN -> Client Specific Overrides -> Click “+”
  - Enabled: Checked
  - Servers: <select the one you just created in step 2>
  - Common name: <name of your Site B certificate CN>
  - Local Network: <10.1.0.0/24>
  - Remote Network: <10.2.0.0/24>
- Press “Save”
- Firewall -> Rules -> WAN-> Click “+”
  - Action: Pass
  - Interface: WAN
  - Direction: in
  - TCP/IP Version: IPv4
  - Protocol: UDP
  - Source: <source ip address or fqdn/ddns> -or- <any>
  - Destination: WAN address
  - Destination port range: from (other) <23222> to (other) <23222>
  - Description: OpenVPN Tunnel
- Press “Save”
- Press “Apply changes”
An Important Note About “Client Specific Overrides”
In short, they (Client Specific Overrides) are required for site-to-site deployments. They serve a very specific purpose and it’s tied to routing tables within OpenVPN. They are called “iroutes” which is different than OPNsense system’s routing table. You may see the routes in the OPNsense system but that information doesn’t correlate the SSL certificate that was authenticated and which routes that unique client (in this case) have on the other end of the wire.
I had a scenario where I had the VPN tunnel up and I was able to ping the 10.255.0.2 (which means the tunnel is up, passing traffic but I couldn’t ping or connect to anything on the 10.2.0.0/24 network. It was tied to a misconfiguration on the Client Specific Overrides.
It’s worthy noting, iroutes within OpenVPN only apply when your tunnel network is larger than /30 network meaning you have the potential of more than one site-to-site client in your tunnel network. Learn more about iroutes, here.
Ubiquiti EdgeRouter “Site B”
To setup the EdgeRouter this will be done via terminal. You can’t do it through the web interface, it’s simply not supported.
- Connect via SSH to your EdgeRouter
- bash
- sudo su –
- mkdir /var/log/openvpn
- chown root:vyattacfg /var/log/openvpn
- chmod 777 /var/log/openvpn
- vi /config/site-to-site.ovpn
- Modify the configuration and put it into vim (right mouse click). Click here for the configuration file
  - remote <ddns.noip.org> (Remote Address) – Update to your fqdn/ip
  - rport <23222> – (Remote Port) Update to Site A’s port number
  - lport <24222> – (Local Port) Update to a higher port number locally at random
  - writepid </var/run/openvpn-vtun1.pid> – (PID file) Change filename to reflect the correct vtun device on the EdgeRouter
  - status <status /var/run/openvpn/status/vtun1.status 30> – (Status file) Change filename to reflect the correct vtun device on the EdgeRouter
  - Insert your certificates into the <tls-crypt> <ca> <cert> and <key> sections.
- Save by pressing colon “:” typing “wq” and pressing enter
- chown root:vyattacfg /config/site-to-site.ovpn
- chmod 744 /config/site-to-site.ovpn
- exit
- configure
- set interfaces openvpn vtun1 config-file /config/site-to-site.ovpn
- commit
- save
This configuration on the EdgeRouter side, since it’s role is a client will receive routing information that will get pushed from the OPNsense instance. Make sure you’re local and remote routes (specifically subnets) are correct on the OPNsense OpenVPN Instance. If you push bad routes, you’ll never route traffic.
Troubleshooting
If you run into an issue on either side, please check the OpenVPN logs, they are very telling when you have ‘verb 4’ / verbose 4 set. Generally speaking, you don’t need to get more aggressive with logging unless you’ve hit a software bug that needs more detailed dump logging. If you encounter any issues on the EdgeRouter side you’ll need to disable/renable the vtun interface. You can do this by issuing the following command:
configure
set interfaces openvpn vtun1 disable
commit
delete interfaces openvpn vtun1 disable
commit
save
Lastly, check your firewall rules. You’re routes maybe correct but if you have a deny policy or don’t have an open firewall acl, again you won’t route traffic.
Technical Dumps
EdgeRouter OpenVPN Version Information
root@edgerouter:/var/log# openvpn --version
OpenVPN 2.4.7 mips-unknown-linux-gnu [SSL (OpenSSL)] [LZO] [LZ4] [EPOLL] [PKCS11] [MH/PKTINFO] [AEAD] built on Apr 22 2022
