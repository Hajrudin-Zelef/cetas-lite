---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-11bn813-wireguard-config-broken-after-upgrade-94116175
title: "Wireguard Config Broken After Upgrade"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-11bn813-wireguard-config-broken-after-upgrade-94116175.md
source_anchor: ""
source_lines: [1, 52]
sha256: e5a318a8e292e668476e62fbcdd160e8a04b057c40f15073533cda7aa8ccfcc4
---

# Wireguard Config Broken After Upgrade

*Source : https://www.reddit.com/r/mikrotik/comments/11bn813/wireguard_config_broken_after_upgrade/*
*Auteur : u/nmbgeek | Score : 3 | r/mikrotik*

I went to update a RB3011 this morning running 7.5 to 7.7 so I created a backup and completed the upgrade.  I don't *think* I changed anything else.  Prior to the upgrade my Wireguard connections were working great.  It is 4 offices that are all interconnected with Wireguard and by that I mean each office has a connection back to each office.  This office and one other have an RB3011 and the other 2 are Ubiquiti.  I also have a config for my PC to allow remoting in to each of the offices.  The other 3 offices are still connected and working as expected however the upgraded router stopped passing any Wireguard traffic.  I can still see packets in the Firewall view on the accept input rule for the Wireguard port so traffic is getting to the router however the Wireguard interface is showing no activity.  I tried some other allow rules from other threads even though these are not present in the other office with the RB3011 (RouterOS v7.1).

For troubleshooting I have downgraded back to 7.5, restored backup I created, restored another backup from 12/2022 when everything was working, and still no luck.  I created an export of the current config and redacted public IPs and keys.  Any help would be appreciated.

[https://pastebin.com/6twbhkpG](https://pastebin.com/6twbhkpG)

---

## Commentaires

**nmbgeek** (score 3):

Update on this.  I did find that if I specify the peer endpoint ip and port on the Mikrotik the connection does work.  This has never been required before and will not allow practical use of dynamic IPs or being behind a firewall on the client side.

  **10698** (score 4):

  Try a script/scheduler like:
  
  `:interface wireguard peer set endpoint-address=[:resolve domain-name="remotehost01.dynamic-host.tld"] number=[:interface wireguard peer find where interface="wireguard1"]`
  
  I hope this helps.

    **nmbgeek** (score 1):

    Adding the peer's wireguard interface IP in the allowed addresses seems to have resolved my problem.  I feel like I had it at one point in the past and then removed it after finding it not necessary.

**onosendai1979** (score 1):

Just adding my experience, perhaps someone will find it useful, since I found this post by looking for a solution myself.

I updated my router from 7.6 to 7.8 over the weekend and wireguard VPN stopped working for me too. I'm using wireguard in a road-warrior setup, my phone is connected to my home network, when I'm away.

The problem for me was that after the upgrade the wireguard interface ended up in the WAN interface list instead of in the LAN interfaces list. Manually changing that solved everything for me.

  **nmbgeek** (score 2):

  Mine never switched interface lists and are still in my LAN.  Hopefully this can help someone else though.

  **Ttyrim** (score 1):

  For future reference. I have had this same issue. After the update, the wireguard interface stopped working. I have checked everything, it did show up as LAN interface, but didn't work. Setting it to WAN, then setting it back to LAN solved the problem.

    **Willing_Atmosphere_1** (score 1):

    Man you're a life-saver!  
    I've tried everything until I stumbled upon your comment.  
    Disabling and re-enabling the peer did the trick.
