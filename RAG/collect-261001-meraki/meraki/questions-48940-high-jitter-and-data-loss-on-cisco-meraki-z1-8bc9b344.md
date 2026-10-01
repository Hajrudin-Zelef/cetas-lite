---
id: collect-261001-meraki/meraki/questions-48940-high-jitter-and-data-loss-on-cisco-meraki-z1-8bc9b344
title: "questions-48940-high-jitter-and-data-loss-on-cisco-meraki-z1-8bc9b344"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-48940-high-jitter-and-data-loss-on-cisco-meraki-z1-8bc9b344.md
source_anchor: ""
source_lines: [1, 14]
sha256: 642e4a9781a1e27f50855dbd723ff96f655801304f3774f7518bcbb246246109
---

# questions-48940-high-jitter-and-data-loss-on-cisco-meraki-z1-8bc9b344

Our office is running a Cisco Meraki Z1 Security Appliance and recently (March 5th) the internet has been dropping out intermittently. The appliance remains online but the internet either doesn't work or is unbearably slow. It begins when people start showing up to the office, but then speeds will pick up an hour or so later, but then they may drop again another hour later.
Here is the uplink details for the network currently:
The Z1 datasheet says it supports up to 5 clients. this seems quite low. Our office averages 80-100 clients. In this case, the fact we would have this device is absurd. Unfortunately the guy who set up the Meraki Z1 is no longer with the company, so I can't ask him what his intentions were by buying a device that is so limited.
Here is the traffic analysis from today:
As you can see, around 9:00 when people arrive at the office, the data loss sky rockets. This makes complete sense to me seeing how we have about 100 clients running on a device which only supports 5 client devices. My question is, why is this only happening now? This device was installed nearly 3 years ago and this has never happened before until now.
No configuration for the device has changed, and we have done the following troubleshooting:
- Checked power cable to ensure sufficient power supply
- Checked data cables to ensure they are healthy and working
- Reset the Z1 device to factory settings
- Connected to ISP modem and ensure a healthy connection is present
- Rolled back firmware to see if problem persists, it did.
I am looking at buying either the MX64 (50 clients) or MX84 (200 clients) but first I am asking the question, why is this problem only happening now after using this device for 3 years?
Update:
I should also mention that the appliances received a firmware updated on March 3rd, however Cisco support team has said that this is not the cause of the problem. We rolled back the firmware to test, and the problem seemed to persist.
