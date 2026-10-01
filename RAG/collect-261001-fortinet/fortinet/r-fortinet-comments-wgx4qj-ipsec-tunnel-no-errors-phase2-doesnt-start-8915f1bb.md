---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-wgx4qj-ipsec-tunnel-no-errors-phase2-doesnt-start-8915f1bb
title: "r-fortinet-comments-wgx4qj-ipsec-tunnel-no-errors-phase2-doesnt-start-8915f1bb"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-wgx4qj-ipsec-tunnel-no-errors-phase2-doesnt-start-8915f1bb.md
source_anchor: ""
source_lines: [1, 48]
sha256: 8dffe0a3f53d48916f7ebacdfc124042ebfa8e10091e5c2fac5aadc450a3699d
---

# r-fortinet-comments-wgx4qj-ipsec-tunnel-no-errors-phase2-doesnt-start-8915f1bb

IPSEC Tunnel - No errors, Phase2 doesn't start 
        
    Hi All,
I am trying to get an IPSEC Tunnel up and running and phase1 says it negotiate success according to the logs, then Phase2 never attempts. The phase1 gets torn down and starts all over again.
I am on fortios 7.0.5 fg60poe.
Anyone have any ideas as to what might be happening?
EDIT
Sorry for the lack of details.
Source is a Fortigate 60E with a Frontier DSL connection using PPPoE on WAN1 with a static IP (note, I am not using the unnumbered IP to set the static, that would not work for some reason)
Destination is a Cisco ASA on a Static IP. We have 10 locations deployed with Fortigates, all came up fine on the VPN tunnel but this location. The only real difference here is the PPPoE piece. Configs should be identical across all of them. I checked the ASA, it has the same IPSEC policies applied as other locations.
In the Phase1 setup, I have NAT Traversal set to Enable (This is how all of our sites are set up) Authentication and Phase1 Proposal matches my other sites and settings exactly
Phase2: Subnets and Named Address objects failed for local and remote addresses Enabled Replay Detection - Unchecked Enabled PFS - Checked Local Port - Checked Remote Port - Checked Protocol - Checked Auto-Negotiated - Checked Autokey Keep Alive - Checked
EDIT 2
It looks like this is resolved. I deleted out the entries on the ASA and recreated them and the tunnel came right up. I am not sure why that would have made a difference, but something was stuck some where. Thanks for all of the suggestions!
Section des commentaires
Enable tunnel debugging in CLI, you should obviously replace 1.1.1.1 with the other end of the IPsec tunnel endpoint.
Disable debugging when you're done:
Happy reading, there will be lots of output to go through.
I did this and was trying to scan through it. I see no failures or errors (nothing obvious). I see a message about received nat-d payload mat not detected.
That is all I am seeing. Any ideas on that?
Post what you see here so we can take a look
Make sure you hey the outputs of these debug commands while RESPONDING to VPN requests. Meaning : ask your peer to initiate the tunnel rather than sending traffic yourself.
try sending some traffic over the tunnel and see if phase2 is getting up
Pretty sure I was doing this, but I am not certain. I had to.leave the site, but will attempt this shortly.
Attempted to send traffic from both sides, tunnel will not come up.
I had a similar issue. Try giving the gate a reboot and see if the phase 2 come up after that.
I spent countless hours troubleshooting and more with TAC. Eventually we gaveup figuring out the root problem. Something with the firmware possibly. This was on 6.4.7 in a HA pair.
Rebooted, no dice.
to get around this reboot i deleted all the parts, addresses, address groups, routes, etc that were used by the wizard to build the tunnel.
ASA still has fortinet's number on tunnel stability lol i have asa's not rebooted for many years :)
I had two sites that I had to do this at, I had to delete everything out and rebuild the VPN tunnels, routes and polices, but that is not working at this site.
Can you share config?
Do you have a policy configured?
Do you have route configured?
I will get this info posted when I can get back to a computer. I have policies and routes in place
I updated my original post with additional info. I have the route in place and policies configured.
1.) what is the device on the other Side If no Fortigate may read recommendations from other Vendor. Maybe you have to fallback to ikev1
2.) Negotiation success do not meen that initiated an SPI. May the Fortigate and the other device have talkt to another and the Fortigate has get a matching ISAKMP but not put together because of Routing or Firewall policy problems, DNS Match, Password or Certificates, DPD or AutoNegotiation and so on. Only the Proposal (AES128/SHA512/DH21) matched…
3.) is there a NAT-T or not? Did you get traffic on IP50 oder TCP 4500/500
4.) Tunnel Mode or Interface Mode
And not often but I’m some constellations seen that a Problem with Phase2 kicks Phase1
Thanks for this. I will be digging into it more this weekend to get it figured out. I have a few hours of driving to get home, but I do have remote access to it
Okay 👍🏻 Get home safely an we will stay ready to help you 😉
Well, share more on the other device type.
When this happens in my case. Its almost always phase selectors. I normally go with 0.0.0.0 0.0.0.0 on both networks make a static route.
Okay is it possible that you could post the CLI Configuration from both sides because Fortigate to ASA is a often used constellation.
There a often many different interpretations of how is IKEv2 working from Vendor to Vendor.
I used " set auto-negotiate enable" in the phase 2 config.
