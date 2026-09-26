---
id: collect-260926-mikrotik/mikrotik/questions-970165-strongswan-centos-7-to-mikrotik-router-l2tp-vpn-no-prposal-chos-6e8dc6a8
title: "ipsec.conf - strongSwan IPsec configuration file"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-970165-strongswan-centos-7-to-mikrotik-router-l2tp-vpn-no-prposal-chos-6e8dc6a8.md
source_anchor: ""
source_lines: [1, 73]
sha256: 3209eeceda49100fb8394079cf7bf14613806062a38cd77d6f9a3a4c560416ac
---

# ipsec.conf - strongSwan IPsec configuration file

I have reviewed existing Q&A on this issue and, maybe there's a hook I'm not seeing, but I don't know what else to try.
I get this output when attempting to launch VPN from CentOS "client":
[root@hostname etc]# strongswan up casanova_vpn
initiating Main Mode IKE_SA casanova_vpn[1] to <VPN_SERVER_PUBLIC_IP>
generating ID_PROT request 0 [ SA V V V V V ]
sending packet: from <CENTOS_7_PUBLIC_IP>[500] to <VPN_SERVER_PUBLIC_IP>[500] (176 bytes)
received packet: from <VPN_SERVER_PUBLIC_IP>[500] to <CENTOS_7_PUBLIC_IP>[500] (156 bytes)
parsed ID_PROT response 0 [ SA V V V V ]
received NAT-T (RFC 3947) vendor ID
received XAuth vendor ID
received DPD vendor ID
received FRAGMENTATION vendor ID
selected proposal: IKE:3DES_CBC/HMAC_SHA1_96/PRF_HMAC_SHA1/MODP_1024
generating ID_PROT request 0 [ KE No NAT-D NAT-D ]
sending packet: from <CENTOS_7_PUBLIC_IP>[500] to <VPN_SERVER_PUBLIC_IP>[500] (244 bytes)
received packet: from <VPN_SERVER_PUBLIC_IP>[500] to <CENTOS_7_PUBLIC_IP>[500] (236 bytes)
parsed ID_PROT response 0 [ KE No NAT-D NAT-D ]
generating ID_PROT request 0 [ ID HASH N(INITIAL_CONTACT) ]
sending packet: from <CENTOS_7_PUBLIC_IP>[500] to <VPN_SERVER_PUBLIC_IP>[500] (100 bytes)
received packet: from <VPN_SERVER_PUBLIC_IP>[500] to <CENTOS_7_PUBLIC_IP>[500] (68 bytes)
parsed ID_PROT response 0 [ ID HASH ]
IKE_SA casanova_vpn[1] established between <CENTOS_7_PUBLIC_IP>[<CENTOS_7_PUBLIC_IP>]...<VPN_SERVER_PUBLIC_IP>[<VPN_SERVER_PUBLIC_IP>]
scheduling reauthentication in 3394s
maximum IKE_SA lifetime 3574s
generating QUICK_MODE request 3035167021 [ HASH SA No KE ID ID ]
sending packet: from <CENTOS_7_PUBLIC_IP>[500] to <VPN_SERVER_PUBLIC_IP>[500] (300 bytes)
received packet: from <VPN_SERVER_PUBLIC_IP>[500] to <CENTOS_7_PUBLIC_IP>[500] (68 bytes)
parsed INFORMATIONAL_V1 request 3361583959 [ HASH N(NO_PROP) ]
received NO_PROPOSAL_CHOSEN error notify
establishing connection 'casanova_vpn' failed
[root@hostname etc]#
CentOS /etc/ipsec.conf: (I get it that 3des-sha1-modp1024 are weak. I'll bump up the levels when I get the tunnel working, and eventually migrate to certificates...trying to keep it minimal for debug...)
[root@hostname etc]# cat ipsec.conf
# ipsec.conf - strongSwan IPsec configuration file
# basic configuration
config setup
  # strictcrlpolicy=yes
  # uniqueids = no
# Add connections here.
# Sample VPN connections
conn %default
  ikelifetime=60m
  keylife=20m
  rekeymargin=3m
  keyingtries=1
  keyexchange=ikev1
  authby=secret
  ike=3des-sha1-modp1024!
  esp=3des-sha1-modp1024!
conn casanova_vpn
  keyexchange=ikev1
  left=%defaultroute
  auto=add
  authby=secret
  type=transport
  leftprotoport=17/1701
  rightprotoport=17/1701
  right=<VPN_SERVER_PUBLIC_IP>
[root@breezeview etc]#
Tik config:
ppp profile:
name="Vultr_vpn" local-address=172.16.101.1 remote-address=172.16.101.2 
     use-mpls=default use-compression=default use-encryption=yes 
     only-one=default change-tcp-mss=default use-upnp=default 
     address-list="" dns-server=8.8.8.8,8.8.4.4 on-up="" on-down=""
proposal:
 2    name="vultr" auth-algorithms=sha1 enc-algorithms=3des lifetime=1h 
      pfs-group=modp1024 
policy:
 T * group=default src-address=<CENTOS_7_PUBLIC_IP/32> dst-address=\<VPN_SERVER_PUBLIC_IP/32> protocol=all proposal=vultr template=yes
So...everything appears to line up, in theory. In practice...not so much.
Any insights much appreciated!
NO_PROPOSAL_CHOSENusually indicates a problem with the algorithm proposal, but that seems to match. However, it might also be a problem with the traffic selectors (e.g. I don't see any mention of ports/protocol in the Mikrotik policy configuration, or that transport mode should be used).
