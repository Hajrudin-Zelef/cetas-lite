---
id: collect-261001-meraki/meraki/r-meraki-comments-6u4hiu-how-is-one-supposed-to-manage-firewall-outbound-0bd5eb06
title: "r-meraki-comments-6u4hiu-how-is-one-supposed-to-manage-firewall-outbound-0bd5eb06"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-6u4hiu-how-is-one-supposed-to-manage-firewall-outbound-0bd5eb06.md
source_anchor: ""
source_lines: [1, 24]
sha256: d2fcb7be0fb21694fc9b4d4eea2c9c837e304f631f651e19abbd297bd07b0196
---

# r-meraki-comments-6u4hiu-how-is-one-supposed-to-manage-firewall-outbound-0bd5eb06

How is one supposed to manage Firewall > Outbound Rules in MX devices? This is wacky. 
        
    This is not just a rant. This is a real question that starts with a rant.
The more I experience this terrible UI the more angry I get. I just can't believe that Meraki has not improved any of the forms since I became a customer a few years ago. Data entry is tedious and error prone. The worst of it, for me, is when trying to manage Outbound Rules on the Firewall page.
- 
      You can't reference networks or hosts by name.
- 
      They don't provide an enable/disable option.
- 
      There's no way to edit rules en masse. (They should at least provide a text view of the rules so you can do bulk edits. CSV or JSON would be great.)
- 
      The description fields are SERIOUSLY too short.
They have hardware that supports networks many orders of magnitude larger (and, therefore, more complex) than mine so certainly there is a way to manage this without wanting to pull all your hair out.
Can anyone provide some guidance?
The only thing I can think of is to maintain a spreadsheet that I edit first before carefully copying and pasting fields into the form.
Section des commentaires
Absolutely agree -- by far one of the worst aspects of using the MX series. It's so bad that it's forced us to keep using our sonicwalls for our DMZ traffic.
I'll second this. I don't see why their products cost as much as they do but only seem to be half thought out. Time and time again I find that they just flat out don't support something I need to be able to do.
I feel like if you come from a heavy firewall/networking background into Meraki, you see all the holes and half-assed features and think 'WTF, why can't I do XYZ. This VPN support is fucking terrible, these firewall rules sucks, etc.'
If you come from a background maybe managing Linksys/Netgear-type devices and see Meraki, you'd think 'wow, look at the nice graphs, the auto-VPN just works magically, this is awesome'.
This is exactly what killed us looking at the MX as an edge device at our remote locations. It needs a global ruleset and better troubleshooting features to be anything more than a small business solution.
There are templates, which allow for the unified configuration of the firewall config across many branches.
It's still the mostly limited IP based firewall, though.
When you get to templates, it actually allows for quite a bit of granular control, and hostname-based rules are a thing that's coming (have worked with them)
