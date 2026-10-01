---
id: collect-261001-meraki/meraki/questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17-3
title: "questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17.md
source_anchor: ""
source_lines: [97, 162]
sha256: 8d0f4c3ea4b363a86e2a0ba171fe7e284b944c368efd94814a2d910224d6b47a
---

# questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17

2025-05-09T12:32:34-07:00   Informational   charon   06[CFG] <8> selected proposal: IKE:AES_CBC_256/HMAC_SHA2_256_128/PRF_HMAC_SHA2_256/MODP_2048
2025-05-09T12:32:34-07:00   Informational   charon   06[IKE] <8> 192.0.2.226 is initiating an IKE_SA
2025-05-09T12:32:34-07:00   Informational   charon   06[ENC] <8> parsed IKE_SA_INIT request 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(REDIR_SUP) ]
2025-05-09T12:32:34-07:00   Informational   charon   06[NET] <8> received packet: from 192.0.2.226[500] to 192.0.2.13[500] (464 bytes)
2025-05-09T12:32:33-07:00   Informational   charon   06[NET] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> sending packet: from 192.0.2.13[500] to 192.0.2.226[500] (80 bytes)
2025-05-09T12:32:33-07:00   Informational   charon   06[ENC] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> generating INFORMATIONAL response 2 [ ]
2025-05-09T12:32:33-07:00   Informational   charon   06[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> IKE_SA deleted
2025-05-09T12:32:33-07:00   Informational   charon   06[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> deleting IKE_SA 8600b45c-2c67-49ed-b27f-593e09665e7a[1] between 192.0.2.13[192.0.2.13]...192.0.2.226[192.0.2.226]
2025-05-09T12:32:33-07:00   Informational   charon   06[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> received DELETE for IKE_SA 8600b45c-2c67-49ed-b27f-593e09665e7a[1]
2025-05-09T12:32:33-07:00   Informational   charon   06[ENC] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> parsed INFORMATIONAL request 2 [ D ]
2025-05-09T12:32:33-07:00   Informational   charon   06[NET] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> received packet: from 192.0.2.226[500] to 192.0.2.13[500] (80 bytes)
2025-05-09T12:32:28-07:00   Informational   charon   06[NET] <7> sending packet: from 192.0.2.13[500] to 192.0.2.114[500] (472 bytes)
2025-05-09T12:32:28-07:00   Informational   charon   06[ENC] <7> generating IKE_SA_INIT response 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(CHDLESS_SUP) N(MULT_AUTH) ]
2025-05-09T12:32:28-07:00   Informational   charon   06[CFG] <7> selected proposal: IKE:AES_CBC_256/HMAC_SHA2_256_128/PRF_HMAC_SHA2_256/MODP_2048
2025-05-09T12:32:28-07:00   Informational   charon   06[IKE] <7> 192.0.2.114 is initiating an IKE_SA
2025-05-09T12:32:28-07:00   Informational   charon   06[ENC] <7> parsed IKE_SA_INIT request 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(REDIR_SUP) ]
2025-05-09T12:32:28-07:00   Informational   charon   06[NET] <7> received packet: from 192.0.2.114[500] to 192.0.2.13[500] (464 bytes)
2025-05-09T12:32:28-07:00   Informational   charon   06[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> sending packet: from 192.0.2.13[500] to 192.0.2.114[500] (80 bytes)
2025-05-09T12:32:28-07:00   Informational   charon   06[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> generating INFORMATIONAL response 2 [ ]
2025-05-09T12:32:28-07:00   Informational   charon   06[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> IKE_SA deleted
2025-05-09T12:32:28-07:00   Informational   charon   06[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> deleting IKE_SA 964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9[2] between 192.0.2.13[192.0.2.13]...192.0.2.114[192.0.2.114]
2025-05-09T12:32:28-07:00   Informational   charon   06[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received DELETE for IKE_SA 964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9[2]
2025-05-09T12:32:28-07:00   Informational   charon   06[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> parsed INFORMATIONAL request 2 [ D ]
2025-05-09T12:32:28-07:00   Informational   charon   06[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received packet: from 192.0.2.114[500] to 192.0.2.13[500] (80 bytes)
2025-05-09T12:30:12-07:00   Informational   charon   16[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> sending packet: from 192.0.2.13[500] to 192.0.2.114[500] (80 bytes)
2025-05-09T12:30:12-07:00   Informational   charon   16[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> generating INFORMATIONAL response 1 [ D ]
2025-05-09T12:30:12-07:00   Informational   charon   16[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> outbound CHILD_SA 4c518c5c-c2f4-4049-a569-60a578ab4fe9{6} established with SPIs ca176942_i c075fb18_o and TS 172.17.0.0/24 === 172.16.0.0/24
2025-05-09T12:30:12-07:00   Informational   charon   16[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> CHILD_SA closed
2025-05-09T12:30:12-07:00   Informational   charon   16[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> sending DELETE for ESP CHILD_SA with SPI c8e6aa8f
2025-05-09T12:30:12-07:00   Informational   charon   16[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> closing CHILD_SA 4c518c5c-c2f4-4049-a569-60a578ab4fe9{2} with SPIs c8e6aa8f_i (29201377 bytes) cdb2bcf7_o (33601900 bytes) and TS 172.17.0.0/24 === 172.16.0.0/24
2025-05-09T12:30:12-07:00   Informational   charon   16[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received DELETE for ESP CHILD_SA with SPI cdb2bcf7
2025-05-09T12:30:12-07:00   Informational   charon   16[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> parsed INFORMATIONAL request 1 [ D ]
2025-05-09T12:30:12-07:00   Informational   charon   16[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received packet: from 192.0.2.114[500] to 192.0.2.13[500] (80 bytes)
2025-05-09T12:30:12-07:00   Informational   charon   08[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> sending packet: from 192.0.2.13[500] to 192.0.2.114[500] (528 bytes)
2025-05-09T12:30:12-07:00   Informational   charon   08[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> generating CREATE_CHILD_SA response 0 [ N(ESP_TFC_PAD_N) SA No KE TSi TSr ]
2025-05-09T12:30:12-07:00   Informational   charon   08[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> inbound CHILD_SA 4c518c5c-c2f4-4049-a569-60a578ab4fe9{6} established with SPIs ca176942_i c075fb18_o and TS 172.17.0.0/24 === 172.16.0.0/24
2025-05-09T12:30:12-07:00   Informational   charon   08[CFG] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> selected proposal: ESP:AES_CBC_256/HMAC_SHA2_256_128/MODP_2048/NO_EXT_SEQ
2025-05-09T12:30:12-07:00   Informational   charon   08[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> parsed CREATE_CHILD_SA request 0 [ N(REKEY_SA) SA No KE TSi TSr ]
2025-05-09T12:30:12-07:00   Informational   charon   08[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received packet: from 192.0.2.114[500] to 192.0.2.13[500] (528 bytes)
On the OPNSense side my tunnel settings are
Phase 1:
Proposals: aes256-sha256-modp2048 [DH14]
Unique: no
Version: IKEv2
Re-auth time: 28800
DPD delay: 10
Phase 2:
Mode: Tunnel
Policies: yes
Start action: start
close action: none
DPD action: trap
ESP proposals: aes256-sha256-modp2048 [DH14]
Rekey time: 3600
On the meraki side
Phase 1:
Encryption: AES256
Authentication: SHA256
Pseudo-random function: Default
Diffie-Hellman group: 14
Lifetime: 28800
Phase 2:
Encryption: AES256
Authentication: SHA256
PFS group: 14
Lifetime: 3600
