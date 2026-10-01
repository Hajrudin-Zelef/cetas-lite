---
id: collect-261001-cisco/cisco/cisco-vpn-client-configuration-setup-for-ios-router-1
title: "cisco-vpn-client-configuration-setup-for-ios-router"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/cisco-vpn-client-configuration-setup-for-ios-router.md
source_anchor: ""
source_lines: [1, 70]
sha256: 5ccda0eda97f14a1f6566d6dd633ebb633f3481c1d289489f935bd25fd573d71
---

# cisco-vpn-client-configuration-setup-for-ios-router

Cisco VPN Client Configuration - Setup for IOS Router

Remote VPN access is an extremely popular service amongst Cisco routers and ASA Firewalls. The flexibility of having remote access to our corporate network and its resources literally from anywhere in the world, has proven extremely useful and in many cases irreplaceable. All that is required is fast Internet connection and your user credentials to log in – all the rest are taken care by your Cisco router or firewall appliance.

To initiate the connection, we use the Cisco VPN client, available for Windows operating systems (XP, Vista, Windows 7 - 32 & 64bit), Linux, Mac OS X10.4 & 10.5 and Solaris UltraSPARC (32 & 64bit), making it widely available for most users around the globe. Cisco VPN Clients are available for download from our Cisco Downloads section.

The Cisco VPN also introduces the concept of ‘Split Tunneling'. Split tunneling is a feature that allows a remote VPN client access the company's LAN, but at the same time surf the Internet. In this setup, only traffic destined to the company's LAN is sent through the VPN tunnel (encrypted) while all other traffic (Internet) is routed normally as it would if the user was not connected to the company VPN.

Some companies have a strict policy that does not allow the remote VPN client access the Internet while connected to the company network (split tunneling disabled) while others allow restricted access to the Internet via the VPN tunnel (rare)! In this case, all traffic is tunnelled through the VPN and there's usually a web proxy that will provide the remote client restricted Internet access.

From all the above, split tunneling is the most common configuration of Cisco VPN configuration today, however for educational purposes, we will be covering all methods.

Setting up a Cisco router to accept remote Cisco VPN clients is not an extremely difficult task. Following each step shown in this article will guarantee it will work flawlessly.

Below is a typical diagram of a company network providing VPN access to remote users in order to access the company's network resources.

The VPN established is an IPSec secure tunnel and all traffic is encrypted using the configured encryption algorithm:

The Cisco IPSec VPN has two levels of protection as far as credentials concern. The remote client must have valid group authentication credential, followed by valid user credential.

The group credentials are entered once and stored in the VPN connection entry, however the user credentials are not stored and requested every time a connection is established:

We should note that configuring your router to support Point-to-Point Tunnel Protocol VPN (PPTP) is an alternative method and covered on our Cisco PPTP Router Configuration article, however PPTP VPN is an older, less secure and less flexible solution. We highly recommend using Cisco IPSec VPN only.

In order to configure Cisco IPSec VPN client support, the router must be running at least the 'Advanced Security' IOS otherwise most of the commands that follow will not be available at the CLI prompt!

To begin, we need to enable the router's 'aaa model' which stands for 'Authentication, Authorisation and Accounting'. AAA provides a method for identifying users who are logged in to a router and have access to servers or other resources.

AAA also identifies the level of access that has been granted to each user and monitors user activity to produce accounting information.

We enable the 'aaa new-model' service followed by X-Auth for user authentication and then group authentication (network vpn_group_ml_1):

R1# configure terminal R1(config)# aaa new-model R1(config)# aaa authentication login default local R1(config)# aaa authentication login vpn_xauth_ml_1 local R1(config)# aaa authentication login sslvpn local R1(config)# aaa authorization network vpn_group_ml_1 local R1(config)# aaa session-id common

When trying to establish an IPSec tunnel, there are two main phase negotiations where the remote client negotiates the security policies and encryption method with the Cisco VPN router.

Now we create the user accounts that will be provided to our remote users. Each time they try to connect to our VPN, they will be required to enter this information:

We next create an Internet Security Association and Key Management Protocol (ISAKMP) policy for Phase 1 negotiations. In this example, we've create two ISAKMP policies, and configure the encryption (encr), authentication method, hash algorithm and set the Diffie-Hellman group:

We now create a group and configure the DNS server and other parameters as required. These parameters are passed down to the client as soon as it successfully authenticates to the group:

R1(config)# crypto isakmp client configuration group CCLIENT-VPN R1(config-isakmp-group)# key firewall.cx R1(config-isakmp-group)# dns 10.0.0.10 R1(config-isakmp-group)# pool VPN-Pool R1(config-isakmp-group)# acl 120 R1(config-isakmp-group)# max-users 5 R1(config-isakmp-group)# exit R1(config)# ip local pool VPN-Pool 192.168.0.20 192.168.0.25

The above configuration is for the 'CCLIENT-VPN' group with a pre-share key (authentication method configured previously) of 'firewall.cx'. Users authenticating to this group will have their DNS set to 10.0.0.10. A maximum of 5 users are allowed to connect simultaneously to this group and will have access to the resources governed by access-list 120.

Lastly, users authenticating to this group will obtain their IP address from the pool named 'VPN-Pool' that provides the range of IP address: 192.168.0.20 up to 192.168.0.25.

Creation of the Phase 2 Policy is next. This is for actual data encryption & IPSec phase 2 authentication:

The transformation named 'encrypto-method-1' is then applied to an IPSec profile named 'VPN-Profile-1':

R1(config)# crypto ipsec profile VPN-Profile-1 R1(ipsec-profile)# set transform-set encrypt-method-1

Note the encryption and authentication method of our IPSec crypto tunnel as shown by a connected VPN client to the router with the above configuration:

Now its time to start binding all the above together by creating a virtual-template interface that will act as a 'virtual interface' for our incoming VPN clients. Remote VPN clients will obtain an IP address that is part of our internal network (see diagram above - 192.168.0.x/24) so we therefore do not require this virtual interface to have an ip address and configure it as an 'ip unnumbered' interface on our router's LAN interface.

Setting an interface as an ip unnumbered enables IP processing through it without assigning an explicit IP address, however you must bind it to a physical interface that does have an IP address configured, usually your LAN interface:

R1(config)# interface Virtual-Template2 type tunnel R1(config-if)# ip unnumbered FastEthernet0/0 R1(config-if)# tunnel mode ipsec ipv4 R1(config-if)# tunnel protection ipsec profile VPN-Profile-1

Above, our virtual template also inherits our configured encryption method via the 'ipsec profile VPN-Profile-1' command which sets the transform method to 'encrypt-method-1' (check previous configuration block) which in turn equals to 'esp-3des esp-sha-hmac'.

Notice how Cisco's CLI configuration follows a logical structure. You configure specific parameters which are then used in other sections of the configuration. If this logic is understood by the engineer, then decoding any given Cisco configuration becomes an easy task.

So far we've enabled the authentication mechanisms (aaa), created an ISAKMP policy, created the VPN group and set its parameters, configured the encryption method (transform-set) and binded it to the virtual template the remote VPN user will connect to.

Second-last step is to create one last ISAKMP profile to connect the VPN group with the virtual template:

