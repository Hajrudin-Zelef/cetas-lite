---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-18ftn6l-vpn-ipsec-down-and-wont-reconnect-7b878970
title: "r-fortinet-comments-18ftn6l-vpn-ipsec-down-and-wont-reconnect-7b878970"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2023-12-11"]
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-18ftn6l-vpn-ipsec-down-and-wont-reconnect-7b878970.md
source_anchor: ""
source_lines: [1, 75]
sha256: ea3f9a8dc6aae82ef5c61c9a6350dbb83c35f518815e836a886b3dda42e2cfb9
---

# r-fortinet-comments-18ftn6l-vpn-ipsec-down-and-wont-reconnect-7b878970

VPN IPSec down and won't reconnect 
        
    Hi,
We have an IPSec connection between Fortigate 200E and 100F. The guy before me did the VPN connections but I figured that he had issues all the time so I need to sort this out once and for all.
I am running some diagnostics maybe with someone's help here will sort it out. I have first checked the public IP's of source and destination and can confirm that they have not changed.
Secondly, I ran the command "diagnose vpn ike gateway list name" and have the below output:
vd: root/0
name: xxxx
version: 1
interface: ppp1 75
addr: x.x.x.x:500 -> x.x.x.x:500
created: 7s ago
IKE SA: created 2/2
IPsec SA: created 0/0
id/spi: 207251 a92425779a0aef14/3d6a6f1d25d6e362
direction: responder
status: connecting, state 3, started 7s ago
id/spi: 207250 b436eb61b2c7a141/0000000000000000
direction: responder
status: connecting, state 3, started 7s ago
Is the above configuration correct? Is there supposed to be 2 IKE SA connections?
I then ran further diagnostics and got this:
2023-12-11 13:15:19.262102 ike 0:a0ee3b66adfc8900/0000000000000000:207145: responder: main mode get 1st message...
2023-12-11 13:15:19.262112 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID RFC 3947 4A131C81070358455C5728F20E95452F
2023-12-11 13:15:19.262119 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID draft-ietf-ipsec-nat-t-ike-03 7D9419A65310CA6F2C179D9215
529D56
2023-12-11 13:15:19.262126 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID draft-ietf-ipsec-nat-t-ike-02 CD60464335DF21F87CFDB2FC68
B6A448
2023-12-11 13:15:19.262132 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID draft-ietf-ipsec-nat-t-ike-02\n 90CB80913EBB696E086381B5
EC427B1F
2023-12-11 13:15:19.262138 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID draft-ietf-ipsec-nat-t-ike-01 16F6CA16E4A4066D83821A0F0A
EAA862
2023-12-11 13:15:19.262145 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID draft-ietf-ipsec-nat-t-ike-00 4485152D18B6BBCD0BE8A84695
79DDCC
2023-12-11 13:15:19.262151 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID DPD AFCAD71368A1F1C96B8696FC77570100
2023-12-11 13:15:19.262157 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID FRAGMENTATION 4048B7D56EBCE88525E7DE7F00D6C2D3
2023-12-11 13:15:19.262163 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID FRAGMENTATION 4048B7D56EBCE88525E7DE7F00D6C2D3C0000000
2023-12-11 13:15:19.262169 ike 0:a0ee3b66adfc8900/0000000000000000:207145: VID FORTIGATE 8299031757A36082C6A621DE00000000
2023-12-11 13:15:19.262187 ike 0:a0ee3b66adfc8900/0000000000000000:207145: negotiation result
2023-12-11 13:15:19.262194 ike 0:a0ee3b66adfc8900/0000000000000000:207145: proposal id = 1:
2023-12-11 13:15:19.262198 ike 0:a0ee3b66adfc8900/0000000000000000:207145: protocol id = ISAKMP:
2023-12-11 13:15:19.262202 ike 0:a0ee3b66adfc8900/0000000000000000:207145: trans_id = KEY_IKE.
2023-12-11 13:15:19.262212 ike 0:a0ee3b66adfc8900/0000000000000000:207145: encapsulation = IKE/none
2023-12-11 13:15:19.262215 ike 0:a0ee3b66adfc8900/0000000000000000:207145: type=OAKLEY_ENCRYPT_ALG, val=AES_CBC, key-len=128
2023-12-11 13:15:19.262219 ike 0:a0ee3b66adfc8900/0000000000000000:207145: type=OAKLEY_HASH_ALG, val=SHA2_256.
2023-12-11 13:15:19.262222 ike 0:a0ee3b66adfc8900/0000000000000000:207145: type=AUTH_METHOD, val=PRESHARED_KEY.
2023-12-11 13:15:19.262226 ike 0:a0ee3b66adfc8900/0000000000000000:207145: type=OAKLEY_GROUP, val=MODP2048.
2023-12-11 13:15:19.262229 ike 0:a0ee3b66adfc8900/0000000000000000:207145: ISAKMP SA lifetime=86400
2023-12-11 13:15:19.262236 ike 0:a0ee3b66adfc8900/0000000000000000:207145: SA proposal chosen, matched gateway
2023-12-11 13:15:19.262241 ike 0: found x.x.x.x 75 -> x.x.x.x:500
2023-12-11 13:15:19.262246 ike 0::207145: DPD negotiated
2023-12-11 13:15:19.262249 ike 0::207145: peer is FortiGate/FortiOS (v0 b0)
2023-12-11 13:15:19.262254 ike 0::207145: selected NAT-T version: RFC 3947
If negotiation is all correct, why is the tunnel remaining down? Thank you
Section des commentaires
u/retrogamer-999 issue solved! I did all the settings that you told me about, tunnel remained down and all I had to do was shut the tunnel down for around 5 to 10 minutes, bring it up and that's it :). Thanks a lot!!
congrats. keep in mind the wizard is great but does a lot of unecessary things. keep it clean and simple.
Thanks!
diagnose vpn ike gateway listis for phase 1. Phase 2 would bediagnose vpn tunnel list, although your output already says "IPsec SA: created 0/0", so there is no functioning phase 2.
There should be more in the debug to say exactly what the problem is.
Please check my pastebins in my reply to retrogamer-999
you need to run diagnose debug application ike -1. that will give you the most verbose output.
use PasteBin and post a sanitized output with 1.1.1.1 for the 100F and 2.2.2.2 for the 200F
also do a "show vpn ipsec phase1-interface" and "show vpn ipsec phase2-interface". remove any public IP's or replace them with 1.1.1.1 and 2.2.2.2 as stated above.
Hi,
My pastebins below thanks
https://pastebin.com/Ai8s2dQV
https://pastebin.com/6xCEAFU2
https://pastebin.com/NLXDj0G9
https://pastebin.com/K6WCtDcu
https://pastebin.com/diPeFqhR
https://pastebin.com/JZuhXmua
OK i was hoping for one link rather than 6.
As the VPN was made buy the wizard there really isnt much that can go wrong, but i would suggest that you simplfy your config. Set proposales to AES256-SHA256 for both phase1 & 2, enable auto negotiate on both ends.
Link 3 doesnt work...
