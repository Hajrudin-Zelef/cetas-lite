---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/nnorx-nix-config-blob-head-docs-unifi-md-927869a2-2
title: "nnorx-nix-config-blob-head-docs-unifi-md-927869a2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2026-09-02", "2026-09-04"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/nnorx-nix-config-blob-head-docs-unifi-md-927869a2.md
source_anchor: ""
source_lines: [95, 131]
sha256: e1f37f5cb530b0009867de1e14c9c64a30cc0cc85b55ae77f3d093993b522b3a
---

# nnorx-nix-config-blob-head-docs-unifi-md-927869a2

Automatic device updates are off, deliberately. The Flex switch supplies PoE to all three Pis, so updating its firmware drops power to the entire DNS layer. That reboot is deferred, not avoided.
Do it on its own evening, with a client pointed at a public resolver first, and expect every Pi to hard-reset.
The switch became unadoptable on 2026-09-02 and stayed that way. Three independent blockers, each of whose obvious fix needs one of the others already cleared, which is why this needs a deliberate order rather than an attempt.
- Its stored inform URL pointed at a retired address that no longer exists
- Device SSH is disabled on it, so set-inform cannot repair that in place,
and enabling device SSH is itself a setting pushed over inform
- Its management rides the untagged native VLAN
A reset device needs a pool to come back on, and there is not one by
default. servers deliberately carries no DHCP pool, for the reasons in
network.md. So step 0 is putting one back, temporarily: add to the
servers segment in lib/net.nix, deploy gate, and remove it again afterwards.
      pool = {
        first = "192.168.20.100";
        last = "192.168.20.150";
      };
Without it a factory-defaulted UniFi device gets no lease and falls back to its
built-in 192.168.1.20, which gate does not address. Leaving the pool in place
permanently trades a rare bootstrap problem for a constant one on every client
VLAN, which is the wrong way round.
Then:
- 
Forget the device in the controller first, then factory reset the switch. Order matters, and this is the step that was missed: two resets in a row appeared to fail because the controller still held the device record and auto-re-adopted within seconds, pushing the same broken saved config back every time. Forgetting deletes the saved config so the reset has something to stick to. Hold reset until the LED changes, roughly ten seconds. Watch the LED, not the clock: white means it worked, blue means it did not. A short press only reboots, and the two are easy to confuse because both make the device drop and return.
- 
Nothing to do on gate. servers is the untagged native VLAN and a
factory-default switch tags nothing, so the two already agree.
- 
Adopt the switch in the controller. Discovery is L2 and everything shares a VLAN at this point, so it should appear on its own.
- 
Change no native network on any port. Add only tagged trusted, iot, work and guest to gate's uplink port and the AP's port, which is what SSIDs need. Everything else stays on the factory default. Setting a port's native network to the servers network object splits its traffic onto VLAN 20, away from
the switch's own management and every port still on the default.
- 
Set the statics: switch .2 , AP.3 . Then remove the transitional pool.
A factory-default UniFi device beacons for a controller on UDP 10001 every few seconds. A switch sending only STP and LLDP, with no DHCP and no beacon, is not waiting for adoption however default it looks:
ssh gate 'tcpdump -i lan0 -nn -e ether host <switch-mac>'
That answers it in twenty seconds and is worth reaching for before trying another reset.
A config change made while the controller cannot reach a device is queued, not lost. It applies whenever contact resumes, which may be much later and during something unrelated. On 2026-09-04 a port profile set while core5 was down landed minutes afterwards, mid-way through an unrelated gate deploy, and flipped that port from tagged to untagged between two packet captures. It looked exactly like the deploy had broken the network. If a change appears not to have applied, assume it is pending rather than lost, and do not stack another change on top of it.
Do not reassure yourself that port profiles live in the controller and come back on adoption. They do, and on 2026-09-04 that was the problem rather than the consolation: the saved profiles were what kept stranding the switch, and discarding them via Forget is what finally broke the loop.
