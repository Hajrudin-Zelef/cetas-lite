---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647-2
title: "site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-opnsense-pfsense/site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647.md
source_anchor: ""
source_lines: [111, 237]
sha256: daff370fc8b51a4c35f30a3462bcf05988dcc1699fce2c1f7a18c9d5576050ba
---

# site-to-site-openvpn-between-opnsense-and-ubiquiti-edgerouter-edgeos-529a9647

library versions: OpenSSL 1.0.2u  20 Dec 2019, LZO 2.08
Originally developed by James Yonan
Copyright (C) 2002-2018 OpenVPN Inc <sales@openvpn.net>
Compile time defines: enable_async_push=no enable_comp_stub=no enable_crypto=yes enable_crypto_ofb_cfb=yes enable_debug=yes enable_def_auth=yes enable_dependency_tracking=no enable_dlopen=unknown enable_dlopen_self=unknown enable_dlopen_self_static=unknown enable_fast_install=needless enable_fragment=yes enable_iproute2=yes enable_libtool_lock=yes enable_lz4=yes enable_lzo=yes enable_maintainer_mode=no enable_management=yes enable_multihome=yes enable_pam_dlopen=no enable_pedantic=no enable_pf=yes enable_pkcs11=yes enable_plugin_auth_pam=yes enable_plugin_down_root=yes enable_plugins=yes enable_port_share=yes enable_selinux=no enable_server=yes enable_shared=yes enable_shared_with_static_runtimes=no enable_silent_rules=no enable_small=no enable_static=yes enable_strict=no enable_strict_options=no enable_systemd=yes enable_werror=no enable_win32_dll=yes enable_x509_alt_username=yes with_aix_soname=aix with_crypto_library=openssl with_gnu_ld=yes with_mem_check=no with_sysroot=no
EdgeRouter OpenVPN Cipher Information
root@edgerouter:/usr/sbin#.\openvpn --show-ciphers
The following ciphers and cipher modes are available for use
with OpenVPN.  Each cipher shown below may be use as a
parameter to the --cipher option.  The default key size is
shown as well as whether or not it can be changed with the
--keysize directive.  Using a CBC or GCM mode is recommended.
In static key mode only CBC mode is allowed.
AES-128-CBC  (128 bit key, 128 bit block)
AES-128-CFB  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-CFB1  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-CFB8  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-GCM  (128 bit key, 128 bit block, TLS client/server mode only)
AES-128-OFB  (128 bit key, 128 bit block, TLS client/server mode only)
AES-192-CBC  (192 bit key, 128 bit block)
AES-192-CFB  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-CFB1  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-CFB8  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-GCM  (192 bit key, 128 bit block, TLS client/server mode only)
AES-192-OFB  (192 bit key, 128 bit block, TLS client/server mode only)
AES-256-CBC  (256 bit key, 128 bit block)
AES-256-CFB  (256 bit key, 128 bit block, TLS client/server mode only)
AES-256-CFB1  (256 bit key, 128 bit block, TLS client/server mode only)
AES-256-CFB8  (256 bit key, 128 bit block, TLS client/server mode only)
AES-256-GCM  (256 bit key, 128 bit block, TLS client/server mode only)
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
SEED-CBC  (128 bit key, 128 bit block)
SEED-CFB  (128 bit key, 128 bit block, TLS client/server mode only)
SEED-OFB  (128 bit key, 128 bit block, TLS client/server mode only)
The following ciphers have a block size of less than 128 bits,
and are therefore deprecated.  Do not use unless you have to.
BF-CBC  (128 bit key by default, 64 bit block)
BF-CFB  (128 bit key by default, 64 bit block, TLS client/server mode only)
BF-OFB  (128 bit key by default, 64 bit block, TLS client/server mode only)
...
EdgeRouter OpenVPN Configuration
Back to EdgeRoute “Site B” Configuration
client
tls-client
dev-type tun
dev vtun1
rport 23222
remote ddns.noip.org
remote-cert-tls server
lport 24222
tun-mtu 1500
auth SHA256
cipher AES-256-GCM
keysize 256
explicit-exit-notify
persist-tun
persist-key
fast-io
auth-nocache
float
resolv-retry 3
persist-tun
verb 4
writepid /var/run/openvpn-vtun1.pid
status /var/run/openvpn/status/vtun1.status 30
log /var/log/openvpn/ovpn.log
ping 10
ping-restart 60
<tls-crypt>
-----BEGIN OpenVPN Static key V1-----
..
-----END OpenVPN Static key V1-----
</tls-crypt>
<ca>
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----
</ca>
<cert>
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----
</cert>
<key>
-----BEGIN RSA PRIVATE KEY-----
...
-----END RSA PRIVATE KEY-----
</key>
EdgeRouter OpenVPN Connection Log File
sudo cat /var/log/openvpn/ovpn.logThu Jan 25 11:47:28 2024 us=120846 OpenVPN 2.4.7 mips-unknown-linux-gnu [SSL (OpenSSL)] [LZO] [LZ4] [EPOLL] [PKCS11] [MH/PKTINFO] [AEAD] built on Apr 22 2022
Thu Jan 25 11:47:28 2024 us=121133 library versions: OpenSSL 1.0.2u  20 Dec 2019, LZO 2.08
Thu Jan 25 11:47:28 2024 us=135172 Outgoing Control Channel Encryption: Cipher 'AES-256-CTR' initialized with 256 bit key
Thu Jan 25 11:47:28 2024 us=135560 Outgoing Control Channel Encryption: Using 256 bit message hash 'SHA256' for HMAC authentication
Thu Jan 25 11:47:28 2024 us=135839 Incoming Control Channel Encryption: Cipher 'AES-256-CTR' initialized with 256 bit key
Thu Jan 25 11:47:28 2024 us=136156 Incoming Control Channel Encryption: Using 256 bit message hash 'SHA256' for HMAC authentication
Thu Jan 25 11:47:28 2024 us=137063 Control Channel MTU parms [ L:1621 D:1156 EF:94 EB:0 ET:0 EL:3 ]
Thu Jan 25 11:47:28 2024 us=464874 Data Channel MTU parms [ L:1621 D:1450 EF:121 EB:406 ET:0 EL:3 ]
Thu Jan 25 11:47:28 2024 us=465237 Local Options String (VER=V4): 'V4,dev-type tun,link-mtu 1549,tun-mtu 1500,proto UDPv4,cipher AES-256-GCM,auth [null-digest],keysize 256,key-method 2,tls-client'
Thu Jan 25 11:47:28 2024 us=465414 Expected Remote Options String (VER=V4): 'V4,dev-type tun,link-mtu 1549,tun-mtu 1500,proto UDPv4,cipher AES-256-GCM,auth [null-digest],keysize 256,key-method 2,tls-server'
Thu Jan 25 11:47:28 2024 us=465625 TCP/UDP: Preserving recently used remote address: [AF_INET]<ip obscured>:23222
Thu Jan 25 11:47:28 2024 us=465841 Socket Buffers: R=[294912->294912] S=[294912->294912]
Thu Jan 25 11:47:28 2024 us=466054 UDP link local (bound): [AF_INET][undef]:24222
Thu Jan 25 11:47:28 2024 us=466237 UDP link remote: [AF_INET]<ip obscured>:23222
Thu Jan 25 11:47:28 2024 us=513089 TLS: Initial packet from [AF_INET]<ip obscured>:23222, sid=36f84160 d53e60ae
Thu Jan 25 11:47:28 2024 us=606744 VERIFY OK: depth=1, CN=<ca obscured>
Thu Jan 25 11:47:28 2024 us=626181 VERIFY KU OK
Thu Jan 25 11:47:28 2024 us=626486 Validating certificate extended key usage
Thu Jan 25 11:47:28 2024 us=626735 ++ Certificate has EKU (str) TLS Web Server Authentication, expects TLS Web Server Authentication
Thu Jan 25 11:47:28 2024 us=626953 VERIFY EKU OK
Thu Jan 25 11:47:28 2024 us=627157 VERIFY OK: depth=0, CN=<name obscured> 
Thu Jan 25 11:47:31 2024 us=333020 WARNING: 'link-mtu' is used inconsistently, local='link-mtu 1549', remote='link-mtu 1553'
Thu Jan 25 11:47:31 2024 us=333346 WARNING: 'cipher' is present in local config but missing in remote config, local='cipher AES-256-GCM'
Thu Jan 25 11:47:31 2024 us=333567 WARNING: 'auth' is used inconsistently, local='auth [null-digest]', remote='auth SHA256'
