---
id: collect-261001-meraki/meraki/questions-8252-openswan-ipsec-connection-to-cisco-meraki-mx-appliance-6aea6123
title: "questions-8252-openswan-ipsec-connection-to-cisco-meraki-mx-appliance-6aea6123"
domain: meraki
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/questions-8252-openswan-ipsec-connection-to-cisco-meraki-mx-appliance-6aea6123.md
source_anchor: ""
source_lines: [1, 21]
sha256: 151158ce5ad1957dfc1c917171c1fcde834554cf438cd01583453098c7c25165
---

# questions-8252-openswan-ipsec-connection-to-cisco-meraki-mx-appliance-6aea6123

I have service I need to make VPN connections to.
I have this working perfectly fine in windows, setup was quick, easy and straightforward.
I am now trying to achieve the same connection in Linux using OpenSWAN but my connection is failing with:
117 "MyConnection" #115: STATE_QUICK_I1: initiate
010 "MyConnection" #115: STATE_QUICK_I1: retransmission; will wait 20s for response
010 "MyConnection" #115: STATE_QUICK_I1: retransmission; will wait 40s for response
031 "MyConnection" #115: max number of retransmissions (2) reached STATE_QUICK_I1.  No 
  acceptable response to our first Quick Mode message: perhaps peer likes no proposal
000 "MyConnection" #115: starting keying attempt 2 of an unlimited number, but releasing 
  whack
I'm not exactly sure what this means, though have Googled a lot though I'm guessing I have some configuration parameters wrong.
In windows I have the following configuration defined (everything else is default)
Username and password
L2TP/IPSec selected
PSK defined in advanced options of the L2TP/IPSec connection
Require Encryption turned on
Unencrypted password (PAP) checked
And it all just works - unfortunately I'm not sure how to translate these settings fully to relevant ipsec.conf parameters.
The VPN I'm connecting to is a Cisco meraki MX appliance if that helps...
I guess if anyone has a sample config for an openSWAN connection to Cisco meraki MX appliance that would be a helpful starting point, but more specifically if someone can translate the windows VPN settings to ipsec.conf options that would be the most useful thing
No acceptable response to our first Quick Mode message: perhaps peer likes no proposalsuggests that perhaps you need to configured aggressive mode, or perhaps you haven't configured the correct transforms. Please check your configuration and provide more detailed logging output.
