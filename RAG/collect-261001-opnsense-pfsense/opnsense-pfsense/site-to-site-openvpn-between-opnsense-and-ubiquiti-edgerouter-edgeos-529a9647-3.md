---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647-3
title: "site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-opnsense-pfsense/site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647.md
source_anchor: ""
source_lines: [238, 324]
sha256: ed5c19a0ded0221b9226f50b745ccb138021cf5124572a853b544fb0a61a6ffe
---

# site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647

Thu Jan 25 11:47:31 2024 us=333878 WARNING: 'keysize' is used inconsistently, local='keysize 256', remote='keysize 128'
Thu Jan 25 11:47:31 2024 us=334874 Control Channel: TLSv1.2, cipher TLSv1/SSLv3 ECDHE-RSA-AES256-GCM-SHA384, 4096 bit RSA
Thu Jan 25 11:47:31 2024 us=335202 [<name obscured>] Peer Connection Initiated with [AF_INET]<ip obscured>:23222
Thu Jan 25 11:47:32 2024 us=409756 SENT CONTROL [<name obscured>]: 'PUSH_REQUEST' (status=1)
Thu Jan 25 11:47:32 2024 us=463398 PUSH: Received control message: 'PUSH_REPLY,route 10.1.0.0 255.255.0.0,route-gateway 10.255.0.1,topology subnet,ping 60,ping-restart 300,ifconfig 10.255.0.2 255.255.255.0,peer-id 0,cipher AES-256-GCM'
Thu Jan 25 11:47:32 2024 us=464096 OPTIONS IMPORT: timers and/or timeouts modified
Thu Jan 25 11:47:32 2024 us=464310 OPTIONS IMPORT: --ifconfig/up options modified
Thu Jan 25 11:47:32 2024 us=464450 OPTIONS IMPORT: route options modified
Thu Jan 25 11:47:32 2024 us=464577 OPTIONS IMPORT: route-related options modified
Thu Jan 25 11:47:32 2024 us=464700 OPTIONS IMPORT: peer-id set
Thu Jan 25 11:47:32 2024 us=464825 OPTIONS IMPORT: adjusting link_mtu to 1624
Thu Jan 25 11:47:32 2024 us=464947 OPTIONS IMPORT: data channel crypto options modified
Thu Jan 25 11:47:32 2024 us=465135 Data Channel MTU parms [ L:1552 D:1450 EF:52 EB:406 ET:0 EL:3 ]
Thu Jan 25 11:47:32 2024 us=466086 Outgoing Data Channel: Cipher 'AES-256-GCM' initialized with 256 bit key
Thu Jan 25 11:47:32 2024 us=466345 Incoming Data Channel: Cipher 'AES-256-GCM' initialized with 256 bit key
Thu Jan 25 11:47:32 2024 us=467418 ROUTE_GATEWAY <ip obscured>/<netmask obscured> IFACE=eth0 HWADDR=04:18:d6:f1:23:25
Thu Jan 25 11:47:32 2024 us=473094 TUN/TAP device vtun1 opened
Thu Jan 25 11:47:32 2024 us=473396 TUN/TAP TX queue length set to 100
Thu Jan 25 11:47:32 2024 us=473806 do_ifconfig, tt->did_ifconfig_ipv6_setup=0
Thu Jan 25 11:47:32 2024 us=474133 /sbin/ip link set dev vtun1 up mtu 1500
Thu Jan 25 11:47:32 2024 us=491786 /sbin/ip addr add dev vtun1 10.255.0.2/24 broadcast 10.255.0.255
Thu Jan 25 11:47:32 2024 us=507385 /sbin/ip route add 10.1.0.0/16 via 10.225.0.1
Thu Jan 25 11:47:32 2024 us=518722 Initialization Sequence Completed
OPNsense OpenVPN Version Information
root@opnsense:~ $ openvpn --version
OpenVPN 2.6.8 amd64-portbld-freebsd13.2 [SSL (OpenSSL)] [LZO] [LZ4] [PKCS11] [MH/RECVDA] [AEAD]
library versions: OpenSSL 1.1.1w  11 Sep 2023, LZO 2.10
Originally developed by James Yonan
Copyright (C) 2002-2023 OpenVPN Inc <sales@openvpn.net>
Compile time defines: enable_async_push=no enable_comp_stub=no enable_crypto_ofb_cfb=yes enable_dco=no enable_debug=yes enable_dlopen=unknown enable_dlopen_self=unknown enable_dlopen_self_static=unknown enable_fast_install=needless enable_fragment=yes enable_iproute2=no enable_libtool_lock=yes enable_lz4=yes enable_lzo=yes enable_management=yes enable_pam_dlopen=no enable_pedantic=no enable_pkcs11=yes enable_plugin_auth_pam=yes enable_plugin_down_root=yes enable_plugins=yes enable_port_share=yes enable_selinux=no enable_shared=yes enable_shared_with_static_runtimes=no enable_silent_rules=no enable_small=no enable_static=yes enable_strict=yes enable_strict_options=no enable_systemd=no enable_unit_tests=no enable_werror=no enable_win32_dll=yes enable_wolfssl_options_h=yes enable_x509_alt_username=no with_aix_soname=aix with_crypto_library=openssl with_gnu_ld=yes with_mem_check=no with_openssl_engine=auto with_sysroot=no
OPNsense OpenVPN Cipher Information
root@opnsense:~ $ openvpn --show-ciphers
The following ciphers and cipher modes are available for use
with OpenVPN.  Each cipher shown below may be used as a
parameter to the --data-ciphers (or --cipher) option. In static
key mode only CBC mode is allowed.
See also openssl list -cipher-algorithms
AES-128-CBC  (128 bit key, 128 bit block)
AES-128-CFB  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-CFB1  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-CFB8  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-OFB  (128 bit key, 128 bit block, TLS client/server mode only)
AES-192-CBC  (192 bit key, 128 bit block)
AES-192-CFB  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-CFB1  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-CFB8  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-OFB  (192 bit key, 128 bit block, TLS client/server mode only)
AES-256-CBC  (256 bit key, 128 bit block)
AES-256-CFB  (256 bit key, 128 bit block, TLS client/server mode only)
AES-256-CFB1  (256 bit key, 128 bit block, TLS client/server mode only)
AES-256-CFB8  (256 bit key, 128 bit block, TLS client/server mode only)
AES-256-OFB  (256 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-128-CBC  (128 bit key, 128 bit block)
CAMELLIA-128-CFB  (128 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-128-CFB1  (128 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-128-CFB8  (128 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-128-OFB  (128 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-192-CBC  (192 bit key, 128 bit block)
CAMELLIA-192-CFB  (192 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-192-CFB1  (192 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-192-CFB8  (192 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-192-OFB  (192 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-256-CBC  (256 bit key, 128 bit block)
CAMELLIA-256-CFB  (256 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-256-CFB1  (256 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-256-CFB8  (256 bit key, 128 bit block, TLS client/server mode only)
CAMELLIA-256-OFB  (256 bit key, 128 bit block, TLS client/server mode only)
CHACHA20-POLY1305  (256 bit key, stream cipher, TLS client/server mode only)
SEED-CBC  (128 bit key, 128 bit block)
SEED-CFB  (128 bit key, 128 bit block, TLS client/server mode only)
SEED-OFB  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-GCM  (128 bit key, 128 bit block, TLS client/server mode only)
AES-192-GCM  (192 bit key, 128 bit block, TLS client/server mode only)
AES-256-GCM  (256 bit key, 128 bit block, TLS client/server mode only)
The following ciphers have a block size of less than 128 bits,
and are therefore deprecated.  Do not use unless you have to.
BF-CBC  (128 bit key, 64 bit block)
BF-CFB  (128 bit key, 64 bit block, TLS client/server mode only)
BF-OFB  (128 bit key, 64 bit block, TLS client/server mode only)
...
References
- https://community.openvpn.net/openvpn/wiki/CipherNegotiation
- https://hohnstaedt.de/xca
- https://hohnstaedt.de/xca/index.php/documentation/stepbystep
- https://docs.netgate.com/pfsense/en/latest/recipes/openvpn-s2s-tls.html
- https://docs.netgate.com/pfsense/en/latest/vpn/openvpn/configure-overrides.html
- https://docs.netgate.com/pfsense/en/latest/troubleshooting/openvpn-iroute.html
