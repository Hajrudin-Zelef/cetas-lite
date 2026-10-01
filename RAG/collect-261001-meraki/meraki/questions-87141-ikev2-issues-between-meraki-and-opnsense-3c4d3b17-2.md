---
id: collect-261001-meraki/meraki/questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17-2
title: "questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17.md
source_anchor: ""
source_lines: [47, 96]
sha256: e0a36123dd139ba9d4118beb646d1d2166ebffa2c937b95186027eca7569eedb
---

# questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17

2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> authentication of '192.0.2.13' (myself) with pre-shared key
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> no IDi configured, fall back on IP address
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> selected proposal: IKE:AES_CBC_256/HMAC_SHA2_256_128/PRF_HMAC_SHA2_256/MODP_2048
2025-05-09T12:32:58-07:00   Informational   charon   13[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> parsed IKE_SA_INIT response 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(CHDLESS_SUP) N(MULT_AUTH) ]
2025-05-09T12:32:58-07:00   Informational   charon   13[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received packet: from 192.0.2.114[500] to 192.0.2.13[500] (472 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   10[NET] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> sending packet: from 192.0.2.13[500] to 192.0.2.226[500] (288 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   10[ENC] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> generating IKE_AUTH request 1 [ IDi AUTH N(ESP_TFC_PAD_N) SA TSi TSr N(MULT_AUTH) N(EAP_ONLY) N(MSG_ID_SYN_SUP) ]
2025-05-09T12:32:58-07:00   Informational   charon   10[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> establishing CHILD_SA 9864a88a-52fa-40e2-bcc0-eea0fb2757c4{1}
2025-05-09T12:32:58-07:00   Informational   charon   10[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> authentication of '192.0.2.13' (myself) with pre-shared key
2025-05-09T12:32:58-07:00   Informational   charon   10[CFG] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> no IDi configured, fall back on IP address
2025-05-09T12:32:58-07:00   Informational   charon   10[CFG] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> selected proposal: IKE:AES_CBC_256/HMAC_SHA2_256_128/PRF_HMAC_SHA2_256/MODP_2048
2025-05-09T12:32:58-07:00   Informational   charon   10[ENC] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> parsed IKE_SA_INIT response 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(CHDLESS_SUP) N(MULT_AUTH) ]
2025-05-09T12:32:58-07:00   Informational   charon   10[NET] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> received packet: from 192.0.2.226[500] to 192.0.2.13[500] (472 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   11[NET] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> sending packet: from 192.0.2.13[500] to 192.0.2.239[500] (464 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   11[ENC] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> generating IKE_SA_INIT request 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(REDIR_SUP) ]
2025-05-09T12:32:58-07:00   Informational   charon   11[IKE] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> initiating IKE_SA 302ac6f2-a390-4413-be90-6f58d4db9d1c[3] to 192.0.2.239
2025-05-09T12:32:58-07:00   Informational   charon   11[CFG] initiating '8be16212-f51a-41bf-9f73-ae47d3f691f0'
2025-05-09T12:32:58-07:00   Informational   charon   11[CFG] added vici connection: 302ac6f2-a390-4413-be90-6f58d4db9d1c
2025-05-09T12:32:58-07:00   Informational   charon   13[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> sending packet: from 192.0.2.13[500] to 192.0.2.114[500] (464 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   13[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> generating IKE_SA_INIT request 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(REDIR_SUP) ]
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> initiating IKE_SA 964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9[2] to 192.0.2.114
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] initiating '4c518c5c-c2f4-4049-a569-60a578ab4fe9'
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] added vici connection: 964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9
2025-05-09T12:32:58-07:00   Informational   charon   12[NET] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> sending packet: from 192.0.2.13[500] to 192.0.2.226[500] (464 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   12[ENC] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> generating IKE_SA_INIT request 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(REDIR_SUP) ]
2025-05-09T12:32:58-07:00   Informational   charon   12[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> initiating IKE_SA 8600b45c-2c67-49ed-b27f-593e09665e7a[1] to 192.0.2.226
2025-05-09T12:32:58-07:00   Informational   charon   12[CFG] initiating '9864a88a-52fa-40e2-bcc0-eea0fb2757c4'
2025-05-09T12:32:58-07:00   Informational   charon   12[CFG] added vici connection: 8600b45c-2c67-49ed-b27f-593e09665e7a
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] loaded IKE shared key with id 'ike-d83b554f-e474-44f2-bb12-e65a6ca98dae' for: '192.0.2.13'
2025-05-09T12:32:57-07:00   Informational   charon   00[JOB] spawning 16 worker threads
2025-05-09T12:32:57-07:00   Informational   charon   00[LIB] loaded plugins: charon aes des blowfish rc2 sha2 sha1 md4 md5 random nonce x509 revocation constraints pubkey pkcs1 pkcs7 pkcs12 pgp dnskey sshkey pem openssl pkcs8 fips-prf curve25519 xcbc cmac hmac kdf gcm drbg curl attr kernel-pfkey kernel-pfroute resolve socket-default stroke vici updown eap-identity eap-md5 eap-mschapv2 eap-radius eap-tls eap-ttls eap-peap xauth-generic xauth-eap xauth-pam whitelist addrblock counters
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loaded 0 RADIUS server configurations
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loading secrets from '/usr/local/etc/ipsec.secrets'
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loading crls from '/usr/local/etc/ipsec.d/crls'
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loading attribute certificates from '/usr/local/etc/ipsec.d/acerts'
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loading ocsp signer certificates from '/usr/local/etc/ipsec.d/ocspcerts'
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loading aa certificates from '/usr/local/etc/ipsec.d/aacerts'
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] loading ca certificates from '/usr/local/etc/ipsec.d/cacerts'
2025-05-09T12:32:57-07:00   Informational   charon   00[CFG] using '/sbin/resolvconf' to install DNS servers
2025-05-09T12:32:57-07:00   Informational   charon   00[LIB] providers loaded by OpenSSL: default legacy
2025-05-09T12:32:57-07:00   Informational   charon   00[DMN] Starting IKE charon daemon (strongSwan 5.9.14, FreeBSD 14.1-RELEASE-p2, amd64)
2025-05-09T12:32:57-07:00   Informational   charon   00[IKE] <7> destroying IKE_SA in state CONNECTING without notification
2025-05-09T12:32:57-07:00   Informational   charon   00[IKE] <8> destroying IKE_SA in state CONNECTING without notification
2025-05-09T12:32:57-07:00   Informational   charon   00[NET] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> sending packet: from 192.0.2.13[500] to 192.0.2.239[500] (80 bytes)
2025-05-09T12:32:57-07:00   Informational   charon   00[ENC] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> generating INFORMATIONAL request 2 [ D ]
2025-05-09T12:32:57-07:00   Informational   charon   00[IKE] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> sending DELETE for IKE_SA 302ac6f2-a390-4413-be90-6f58d4db9d1c[3]
2025-05-09T12:32:57-07:00   Informational   charon   00[IKE] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> deleting IKE_SA 302ac6f2-a390-4413-be90-6f58d4db9d1c[3] between 192.0.2.13[192.0.2.13]...192.0.2.239[192.0.2.239]
2025-05-09T12:32:57-07:00   Informational   charon   00[DMN] SIGTERM received, shutting down
2025-05-09T12:32:34-07:00   Informational   charon   06[NET] <8> sending packet: from 192.0.2.13[500] to 192.0.2.226[500] (472 bytes)
2025-05-09T12:32:34-07:00   Informational   charon   06[ENC] <8> generating IKE_SA_INIT response 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(CHDLESS_SUP) N(MULT_AUTH) ]
