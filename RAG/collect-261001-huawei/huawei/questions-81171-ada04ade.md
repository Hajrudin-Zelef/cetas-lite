---
id: collect-261001-huawei/huawei/questions-81171-ada04ade
title: "questions-81171-ada04ade"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["consumer", "licenses"]
source: docs/RAG/collect-261001-huawei/questions-81171-ada04ade.md
source_anchor: ""
source_lines: [1, 11]
sha256: e93c293eaf0c6687e11ab11ee83b908312455dec6da574d35ac7100a84e1e0da
---

# questions-81171-ada04ade

How can I find out what date the switch is` deployed first time or any closer date or anything that I can estimate? Consider the switch is huwaei/dell/cisco/aruba.
Thanks
Usually you'd have that date in your inventory.
For HPE/Aruba gear you can use the warranty check (requires registration as of late): https://support.hpe.com/connect/s/?card=wc.
Cisco: https://connectthedots.cisco.com/connectdots/serviceWarrantyFinderRequest
Dell: https://www.dell.com/support/contents/en-us/Category/Warranty
Huawei: https://support.huawei.com/enterprise/en/warranty
Feature licenses might tell you when they were activated but often only show how long they're valid.
Other than tools like that you're likely condemned to digging up invoices...
For enterprise grade hardware, the best bet is to find the date of manufacture sticker on the chassis. Beyond that, you'd have to have your own history database, or find it in the vendor's database. I don't know of any hardware that keeps any sort of "warrantee" activation date in NVRAM. (consumer devices, sure, and they're dirt simple to reset.)
If the firmware has never been updated (and how would you know that?), maybe the flash filesystem has a date?
