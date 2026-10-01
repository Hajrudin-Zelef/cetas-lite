---
id: collect-261001-meraki/meraki/questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17-1
title: "questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17.md
source_anchor: ""
source_lines: [1, 46]
sha256: c401f3bd8c386412bb93d7eb08144c33cebbaa4ab7cf4622133b1aa9886fac53
---

# questions-87141-ikev2-issues-between-meraki-and-opnsense-3c4d3b17

I have a meraki VPN mesh which consists of 3 meraki firewalls and 1 OPNSense firewall. There are 3 IKEv2 IPsec connections setup on the OPNSense firewall, one for each meraki. They're all configured identically minus the remote endpoint address and the remote subnets. The tunnels between the OPNSense and meraki firewalls come up and work for a while but will drop offline after some time and need a restart in order to come back up. I figure I have a settings mismatch but I can't find it. I'm not sure if the meraki lifetime is equivalent to the re-auth or rekey time but I've tried both and the behavior doesn't seem to change. What's the best way to get a meraki device connected to OPNSense?
The time it takes to drop is based on the phase 1 re-auth/lifetime of the tunnel. If I set the lifetime(meraki) and re-auth(OPNSense) timers higher then it takes longer for the tunnel to die.
Note about the logs: 12:32 is when I restart the IPsec service on OPNSense to force a tunnel rebuild, 12:30 is when the tunnel died.
On the meraki side I have these logs(IPs have been replaced)
May 9 12:32:58      Non-Meraki VPN  Non-Meraki VPN negotiation  msg: <remote-peer-2|1403> CHILD_SA net-2{2050} established with SPIs cf753c86(inbound) c46fb58a(outbound) and TS 172.16.0.0/24 === 172.17.0.0/24
May 9 12:32:58      Non-Meraki VPN  Non-Meraki VPN negotiation  msg: <remote-peer-2|1403> IKE_SA remote-peer-2[1403] established between 192.0.2.114[192.0.2.114]...192.0.2.13[192.0.2.13]
May 9 12:32:28      Non-Meraki VPN  Non-Meraki VPN negotiation  msg: <remote-peer-2|1400> deleting IKE_SA remote-peer-2[1400] between 192.0.2.114[192.0.2.114]...192.0.2.13[192.0.2.13]
May 9 12:30:12      Non-Meraki VPN  Non-Meraki VPN negotiation  msg: <remote-peer-2|1400> closing CHILD_SA net-2{2046} with SPIs cdb2bcf7(inbound) (30205337 bytes) c8e6aa8f(outbound) (28885177 bytes) and TS 172.16.0.0/24 === 172.17.0.0/24
May 9 12:30:12      Non-Meraki VPN  Non-Meraki VPN negotiation  msg: May 9 19:30:12 10[IKE] <remote-peer-2|1400> outbound CHILD_SA net-2{2048} established with SPIs c075fb18(inbound) ca176942(outbound) and TS 172.16.0.0/24 === 172.17.0.0/24
May 9 12:30:12      Non-Meraki VPN  Non-Meraki VPN negotiation  msg: May 9 19:30:12 10[IKE] <remote-peer-2|1400> inbound CHILD_SA net-2{2048} established with SPIs c075fb18(inbound) ca176942(outbound) and TS 172.16.0.0/24 === 172.17.0.0/24
On the OPNSense I have these for the same time frame
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 2
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 2
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 2
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 2
2025-05-09T12:32:58-07:00   Informational   charon   09[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> CHILD_SA 4c518c5c-c2f4-4049-a569-60a578ab4fe9{2} established with SPIs c46fb58a_i cf753c86_o and TS 172.17.0.0/24 === 172.16.0.0/24
2025-05-09T12:32:58-07:00   Informational   charon   09[CFG] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> selected proposal: ESP:AES_CBC_256/HMAC_SHA2_256_128/NO_EXT_SEQ
2025-05-09T12:32:58-07:00   Informational   charon   09[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> maximum IKE_SA lifetime 3875s
2025-05-09T12:32:58-07:00   Informational   charon   09[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> scheduling reauthentication in 3515s
2025-05-09T12:32:58-07:00   Informational   charon   09[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> IKE_SA 964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9[2] established between 192.0.2.13[192.0.2.13]...192.0.2.114[192.0.2.114]
2025-05-09T12:32:58-07:00   Informational   charon   09[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> authentication of '192.0.2.114' with pre-shared key successful
2025-05-09T12:32:58-07:00   Informational   charon   09[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> parsed IKE_AUTH response 1 [ IDr AUTH SA TSi TSr ]
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 1
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 1
2025-05-09T12:32:58-07:00   Informational   charon   08[NET] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> received packet: from 192.0.2.239[500] to 192.0.2.13[500] (224 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   09[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> received packet: from 192.0.2.114[500] to 192.0.2.13[500] (272 bytes)
2025-05-09T12:32:58-07:00   Notice  charon   [UPDOWN] received up-client event for reqid 1
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> CHILD_SA 9864a88a-52fa-40e2-bcc0-eea0fb2757c4{1} established with SPIs c8a9de60_i c17ae8d4_o and TS 172.17.0.0/24 === 172.18.0.0/24
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> selected proposal: ESP:AES_CBC_256/HMAC_SHA2_256_128/NO_EXT_SEQ
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> maximum IKE_SA lifetime 3713s
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> scheduling reauthentication in 3353s
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> IKE_SA 8600b45c-2c67-49ed-b27f-593e09665e7a[1] established between 192.0.2.13[192.0.2.13]...192.0.2.226[192.0.2.226]
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> authentication of '192.0.2.226' with pre-shared key successful
2025-05-09T12:32:58-07:00   Informational   charon   13[ENC] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> parsed IKE_AUTH response 1 [ IDr AUTH SA TSi TSr ]
2025-05-09T12:32:58-07:00   Informational   charon   13[NET] <8600b45c-2c67-49ed-b27f-593e09665e7a|1> received packet: from 192.0.2.226[500] to 192.0.2.13[500] (256 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   13[NET] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> sending packet: from 192.0.2.13[500] to 192.0.2.239[500] (256 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   13[ENC] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> generating IKE_AUTH request 1 [ IDi AUTH N(ESP_TFC_PAD_N) SA TSi TSr N(MULT_AUTH) N(EAP_ONLY) N(MSG_ID_SYN_SUP) ]
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> establishing CHILD_SA 8be16212-f51a-41bf-9f73-ae47d3f691f0{3}
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> authentication of '192.0.2.13' (myself) with pre-shared key
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> no IDi configured, fall back on IP address
2025-05-09T12:32:58-07:00   Informational   charon   13[CFG] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> selected proposal: IKE:AES_CBC_256/HMAC_SHA2_256_128/PRF_HMAC_SHA2_256/MODP_2048
2025-05-09T12:32:58-07:00   Informational   charon   13[ENC] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> parsed IKE_SA_INIT response 0 [ SA KE No N(NATD_S_IP) N(NATD_D_IP) N(FRAG_SUP) N(HASH_ALG) N(CHDLESS_SUP) N(MULT_AUTH) ]
2025-05-09T12:32:58-07:00   Informational   charon   13[NET] <302ac6f2-a390-4413-be90-6f58d4db9d1c|3> received packet: from 192.0.2.239[500] to 192.0.2.13[500] (472 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   13[NET] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> sending packet: from 192.0.2.13[500] to 192.0.2.114[500] (304 bytes)
2025-05-09T12:32:58-07:00   Informational   charon   13[ENC] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> generating IKE_AUTH request 1 [ IDi AUTH N(ESP_TFC_PAD_N) SA TSi TSr N(MULT_AUTH) N(EAP_ONLY) N(MSG_ID_SYN_SUP) ]
2025-05-09T12:32:58-07:00   Informational   charon   13[IKE] <964c3b6b-7e05-4eaa-8d7b-e8d7de9f06b9|2> establishing CHILD_SA 4c518c5c-c2f4-4049-a569-60a578ab4fe9{2}
