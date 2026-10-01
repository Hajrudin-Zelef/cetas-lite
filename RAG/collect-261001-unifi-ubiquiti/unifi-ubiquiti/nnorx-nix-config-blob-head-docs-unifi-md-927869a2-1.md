---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/nnorx-nix-config-blob-head-docs-unifi-md-927869a2-1
title: "nnorx-nix-config-blob-head-docs-unifi-md-927869a2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: ["2026-09-04", "2026-09-12"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/nnorx-nix-config-blob-head-docs-unifi-md-927869a2.md
source_anchor: ""
source_lines: [1, 94]
sha256: 59621ba0e58f316f726226026ddd6afc7ddd29d3fdeae45756c77b2a5a06890d
---

# nnorx-nix-config-blob-head-docs-unifi-md-927869a2

Runs on core5 as two pinned containers, the Network Application and its
MongoDB, on a private Docker network. Only the application publishes ports; the
database is reachable from nothing but the other container. See
modules/unifi.nix.
It manages the Flex switch and the U7 Pro. Its database holds adoption state, SSIDs, PSKs, VLAN assignments and port profiles, none of which are in this repo under any approach, which makes backups the part that matters.
Decided on evidence rather than taste. Both unifi and mongodb are unfree
(Ubiquiti's EULA and the SSPL), so Hydra does not build them and
cache.nixos.org does not carry them. The module path would mean core5
compiling MongoDB from source, CI attempting the same inside its 350-minute cap,
and the result being pushed to a public Cachix, which is redistribution of
both.
It also decouples controller upgrades from nix flake update. UniFi's database
migrations are one-way, so a lock bump that moved the controller would leave a
generation rollback facing a newer schema with an older binary, which does not
start.
MongoDB 5+ requires ARMv8.2-A. The Pi 5's Cortex-A76 has it and the Pi 4's Cortex-A72 does not, so this cannot move to core4 or lifeline without changing database.
Take the local admin option, not a Ubiquiti account. Signing in with a UI
account ties the controller to their cloud, which is the thing moving off Google
was meant to avoid. If it happens by accident the association lives in the
database, so the fix is to stop both containers, empty /var/lib/unifi/db and
/var/lib/unifi/config, and start again. Emptying db matters as much as
config: the MongoDB init hook only runs against an empty data directory, and
that is what recreates the application user.
Then two settings, neither expressible in Nix because they live in the controller's own database:
- Settings > System: turn off Remote Management and Analytics. Otherwise Google's telemetry has been swapped for Ubiquiti's.
- Settings > System > Backups: set a schedule. See the automated copy for what its interval costs.
The inform host used to be a third. The controller advertises an address for
devices to report to, and on a bridge network that is its container address in
172.16/12, which nothing on the LAN can reach. The symptom is not an error:
adoption appears to begin and then loops forever. modules/unifi.nix now seeds
system_ip into system.properties before every start, from lib/net.nix, so
it survives a volume wipe and follows the host if it renumbers.
Adoption state, SSIDs, PSKs and VLAN assignments live in MongoDB. The controller
writes its own backups to /var/lib/unifi/config/data/backup/, owned by the
core5 user so they can be copied without root:
scp -r core5:/var/lib/unifi/config/data/backup/ ./unifi-backup-$(date +%F)/
Treat those as sensitive. They contain Wi-Fi PSKs and device credentials, so they do not belong in this repo or any public location.
This is state that lives outside the flake and is not reproducible from Nix.
modules/unifi-backup.nix takes the newest file
the controller wrote, encrypts it to the nick age recipient, and pushes it to
a private repo. A daily timer on core5, and a no-op when the newest backup is
one it has already pushed.
It does not make a backup of its own. The controller's .unf is the format its
restore flow expects, and a mongodump would be a second, unsupported path
into a database whose migrations are one-way. The module's only job is moving
a file that already exists off the host that holds the original.
It depends on the controller's own schedule, and the schedule's interval
bounds how stale the off-box copy can be: on a monthly schedule, a change made
on the 2nd is not off the box until the 1st of the following month. Until the
first scheduled run, config/data/backup/autobackup/ is empty and the unit
fails loudly saying so, rather than exiting cleanly on nothing.
Two things on that settings page are easy to misread:
- The tooltip's path is wrong for this install. It names
/var/lib/unifi/backup/autobackup , the location on a Debian package install,
which does not exist in this container. The real directory is/config/data/backup/autobackup , on the volume, so scheduled files survive
the container being recreated.
- Enabled is not the same as having run. logs/backup.log records every
run. On 2026-09-12 it held three, all manual exports, beside a schedule that
had been switched on after its first possible run time and so had never
fired.
Backup Retention: Settings Only is the right choice. It keeps a backup around 30 KB, and settings are what a rebuild needs; statistics history is not.
age encryption needs only the public half of the key, so core5 holds nothing
that can read these back. That is what makes a private repo an acceptable
destination for a file carrying Wi-Fi PSKs: the destination is untrusted by
construction, and the private half is in Bitwarden and
~/.config/sops/age/keys.txt.
The deploy key has write access and core5 can read it, so a compromised core5
could push to the backup repo. It cannot erase what is already there: a ruleset
on nnorx/homelab-state blocks force-pushes and deletion of the default branch,
with no bypass actors, so the history this copy exists for survives the host it
exists to outlive. Verified on 2026-09-12 by a force-push from the owner's own
token, which was rejected.
The plaintext hash is what decides whether to commit, not the encrypted blob. The timer runs daily while the controller writes a new file only on its own schedule, and age uses a fresh ephemeral key per run, so the same file encrypts differently every time. Comparing ciphertext would re-commit an already-pushed backup every day.
Known gap: a silent stop. Nothing alerts on the unit failing.
systemctl status unifi-backup on core5 is the manual check until the Phase 8
monitoring work in router.md covers it.
Whether the repo can stand in for that check is not yet confirmed. If the controller writes a byte-different file on every run, as it probably does since the archive records when it was made, every scheduled run produces a commit, and a repo quiet for longer than one interval means the pipeline is broken. If identical settings produce identical files, commits happen only on real changes and a quiet repo proves nothing. Two consecutive scheduled files with different hashes and no config change between them settles it.
latest is not always the one you want. After a rebuild or a controller
reset, the new controller's first scheduled backup is of an empty config, and
the timer pushes it over the good one. Every earlier version is still in the
repo's history, and the ruleset above is what guarantees that:
git clone git@github.com:nnorx/homelab-state.git && cd homelab-state
git log --format='%h %ad %s' --date=short -- unifi/latest.unf.age
git show <commit>:unifi/latest.unf.age > pick.unf.age
age --decrypt --identity ~/.config/sops/age/keys.txt pick.unf.age > restore.unf
Pick the last commit from before the loss, not the newest. On a rebuilt core5,
systemctl stop unifi-backup.timer until the restore is done keeps an empty
controller's backup from landing on top in the meantime.
Then a fresh controller, and Settings > System > Backups > Restore. Expect to re-adopt: a restore brings back the saved device config, which is the thing that stranded the switch on 2026-09-04, so read the re-adoption section below before assuming it will come back clean.
A backup nobody has restored is not a backup. This path has not been drilled yet.
Change the digest in modules/unifi.nix. Get the new one with:
nix run nixpkgs#skopeo -- inspect --format '{{.Digest}}' \
  docker://lscr.io/linuxserver/unifi-network-application:<version>
Read Ubiquiti's release notes first. Downgrading needs a restore from backup, not a generation rollback, because the migrations are one-way.
