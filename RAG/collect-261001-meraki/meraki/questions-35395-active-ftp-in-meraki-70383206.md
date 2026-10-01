---
id: collect-261001-meraki/meraki/questions-35395-active-ftp-in-meraki-70383206
title: "questions-35395-active-ftp-in-meraki-70383206"
domain: meraki
role: reference
task: reference
actors: []
dates: ["2016-12", "2017-08-15"]
keywords: []
source: docs/RAG/collect-261001-meraki/questions-35395-active-ftp-in-meraki-70383206.md
source_anchor: ""
source_lines: [1, 16]
sha256: 8835d37b85ee63299a42c84ee49f128d87550a5531fa47cc60d80c965fe62a13
---

# questions-35395-active-ftp-in-meraki-70383206

Has anyone configured Active FTP in Meraki? If so, will you please assist me with how to configure it on an MX100?
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you could provide and accept your own answer.Ron Maupin– Ron Maupin ♦2017-08-15 03:42:11 +00:00Commented Aug 15, 2017 at 3:42
2 Answers 2
Cisco maintains documentation for this type of thing. It takes about 2 seconds to locate the specific document on how to do this. Active and Passive FTP Overview and Configuration:
  Active FTP Overview
An active FTP session involves the following steps:
- The client sends the PORT command to an FTP server. The source port is a random, high-numbered port. The destination port is 21.
- The server responds with an ACK.
- The server initiates a connection to the client with source port 20 and the destination port specified in the client’s PORT command.
- The client sends an ACK to the server. The FTP session has now been established. Client firewalls are often configured to block incoming connections. This causes step 3 of the above process to fail, as shown below:
MX Configuration for Active FTP
Configuration for active FTP on an MX appliance is a simple process. Firewall rules must be constructed to allow inbound connections on port 21 and 20. Additional information about constructing firewall rules can be found here, and the following example below details a 1:1 NAT rule that allows inbound connections to an internal FTP server.
By default, MX appliances allow all outbound connections, so no additional firewall configuration is necessary. The following diagram outlines the flow of active FTP traffic, and where the MX comes into play:
Just to add to this, client outbound active FTP (Ron's first picture) is not supported as of December 2016 without installing beta firmware on the Meraki MX appliances. Check the first note on the Meraki documentation:
This is because Meraki MX does not ship with a FTP inspection engine like other vendors for client FTP sessions. There is a beta firmware available to correct this, have to open a Meraki TAC case to have it pushed down for upgrade.
