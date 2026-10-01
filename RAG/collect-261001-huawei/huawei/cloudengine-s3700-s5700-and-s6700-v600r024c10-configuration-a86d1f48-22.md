---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-22
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [2283, 2441]
sha256: 1f975be4ac7e9a10b8ddef0ebddbdd24a14d835a983aa747310c271f84c0c6eb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Excess-Offset Attack
                  An attacker sends a fragment with a large offset value to a target device, which
                  then allocates memory space to store all fragments, consuming a large amount of
                  resources.
                  The offset field is 13 bits long and measured in units of 8-byte blocks, so its
                  theoretical maximum value is 8191. In most cases, however, the value should be
                  smaller than 8190, since an offset of 8190 will lead to the payload (8190 x 8 =

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  35
Security Configuration
Security Configuration                                               3 Local Attack Defense Configuration


                  65520) exceeding the maximum (65515). Therefore, the maximum offset is 8189,
                  and the last fragment has a maximum of 3-byte IP payload (8189 x 8 - 65515).

                  The device considers packets with an offset value larger than 8189 malicious and
                  discards them.


Repeated Fragment Attack
                  A repeated fragment attack is an attack in which fragments are sent to a target
                  host multiple times. Such an attack is carried out in one of two ways:

                  ●      The same fragments are sent multiple times, causing high CPU usage or a
                         memory error on the target host.
                  ●      Different fragments with the same offset are sent. In this case, the target host
                         cannot determine which fragment will be reserved and which fragment will
                         be discarded, or whether all fragments need to be discarded. As a result, the
                         target host may experience high CPU usage or a memory error.

                  With defense against packet fragment attacks enabled, the device applies the
                  committed access rate (CAR) limit to packet fragments, reserves the first
                  fragment, and discards all the remaining repeated fragments to protect the CPU.


Syndrop Attack
                  A Syndrop attack exploits the mechanism of IP fragmentation to put the second
                  fragment into the first. The offset of the second fragment is smaller than that of
                  the first fragment, and the offset plus the data field of the second fragment does
                  not exceed the tail of the first fragment. A Syndrop attack uses TCP packets that
                  carry a SYN flag and also IP payload.

                  As shown in Figure 3-5:

                  ●      The IP payload in the first fragment is 28 bytes, and the IP header is 20 bytes.
                  ●      The IP payload in the second fragment is 4 bytes, the IP header is 20 bytes,
                         and the offset is 24 (should be 28).

                  Figure 3-5 Syndrop attack




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              36
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration


                  Syndrop attacks cause the system to restart or even crash. With defense against
                  fragmentation attacks enabled, the device discards all fragments in a Syndrop
                  attack.

NewTear Attack
                  A NewTear attack uses error fragments. As shown in Figure 3-6, the protocol is
                  UDP.
                  ●      The IP payload in the first fragment is 28 bytes including the UDP header. The
                         UDP checksum is 0.
                  ●      The IP payload in the second fragment is 4 bytes, and the offset is 24 (should
                         be 28).

                  Figure 3-6 NewTear attack




                  NewTear attacks cause the system to restart or even crash. With defense against
                  fragmentation attacks enabled, the device discards all fragments in a NewTear
                  attack.

Bonk Attack
                  A Bonk attack uses error fragments. As shown in Figure 3-7, the protocol is UDP.
                  ●      The IP payload in the first fragment is 36 bytes including the UDP header. The
                         UDP checksum is 0.
                  ●      The IP payload in the second fragment is 4 bytes, and the offset is 32 (should
                         be 36).




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            37
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration


                  Figure 3-7 Bonk attack




                  Bonk attacks cause the system to restart or even crash. With defense against
                  fragmentation attacks enabled, the device discards all fragments in a Bonk attack.

Nesta Attack
                  A Nesta attack uses error fragments. As shown in Figure 3-8:
                  ●      The IP payload in the first fragment is 18 bytes, the protocol used is UDP, and
                         the checksum is 0.
                  ●      The offset in the second fragment is 48 and the IP payload is 116 bytes.
                  ●      The offset in the third fragment is 0, the More Frag field is 1 (meaning there
                         are more fragments), the IP option (all EOLs) is 40 bytes, and the IP payload
                         is 224 bytes.

                  Figure 3-8 Nesta attack




                  Nesta attacks cause the system to restart or even crash. With defense against
                  fragmentation attacks enabled, the device discards all fragments in a Nesta attack.

3.9.2 Configuring Defense Against Fragmentation Attacks
Context
                  With defense against fragmentation attacks enabled, the device rate-limits the
                  received fragmented packets, discarding any that exceeds the limit.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            38
Security Configuration
Security Configuration                                                3 Local Attack Defense Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable defense against fragmentation attacks.
                  anti-attack fragment enable

                         NOTE

                  You can also run the anti-attack enable command in the system view to enable attack defense
                  against all attack packets, including fragmented packets.

         Step 3 Configure the rate limit for fragmented packets.
                  anti-attack fragment car cir cir-num

                  ----End

Verifying the Configuration
                  Run the display anti-attack statistics fragment command to check statistics
                  relating to defense against fragmentation attacks on the device.


3.10 Configuring Defense Against TCP SYN Flood
Attacks

