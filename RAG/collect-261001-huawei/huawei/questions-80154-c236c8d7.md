---
id: collect-261001-huawei/huawei/questions-80154-c236c8d7
title: "questions-80154-c236c8d7"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2022-10-07"]
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/questions-80154-c236c8d7.md
source_anchor: ""
source_lines: [1, 27]
sha256: 5c520404e1c6565c5dce270524e01c87e853c350c3b8670477e5603b371743cb
---

# questions-80154-c236c8d7

i have a CE6851-48S6Q-HI with the following information:
VRP (R) software, Version 8.180 (CE6851HI V200R005C10SPC800)
HUAWEI CE6851-48S6Q-HI uptime is 1 day, 9 hours, 41 minutes
CE6851-48S6Q-HI(Master) 1 : uptime is  1 day, 9 hours, 40 minutes
        StartupTime 2022/10/07   07:33:05+04:30
Memory    Size    : 2048 M bytes
Flash     Size    : 1024 M bytes
CE6851-48S6Q-HI version information
1. PCB    Version : CEM48S6QP04    VER B
2. MAB    Version : 1
3. Board  Type    : CE6851-48S6Q-HI
4. CPLD1  Version : 103
5. CPLD2  Version : 103
6. BIOS   Version : 433
yesterday my switch rebooted without any reason and when i checked "display device board reset 1" command it shows this:
Board 1 reset information:
-- 1. DATE:2022-10-07  TIME:07:33:05+04:30  BARCODE:2102X  RESET Num:1
--    Reason:Board register, BarCode is 2102X.
--    BootMode:NORMAL
--    BootCode:0x060100ff
-- 2. DATE:2022-10-07  TIME:07:29:18+04:30  BARCODE:21023X  RESET Num:1
--    Reason:LAN Switch parity/ecc error, and reset board.(CPU Reset)
--    BootMode:NORMAL
--    BootCode:0x80000160
i searched a bit and find https://support.huawei.com/enterprise/en/doc/EDOC1100164803 and i think it was a software issue or bug that cause this,
Do you have any idea? can i ignore this?
Thank you.
