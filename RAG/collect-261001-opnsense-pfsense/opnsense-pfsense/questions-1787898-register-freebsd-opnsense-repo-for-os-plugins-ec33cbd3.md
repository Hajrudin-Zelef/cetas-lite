---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1787898-register-freebsd-opnsense-repo-for-os-plugins-ec33cbd3
title: "Register FreeBSD OPNsense Repo for `os-` plugins"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1787898-register-freebsd-opnsense-repo-for-os-plugins-ec33cbd3.md
source_anchor: ""
source_lines: [1, 31]
sha256: bdb7853933043513cd8218e6584de5377589bc7b60e4ab9f5ad9a5a494126a72
---

# Register FreeBSD OPNsense Repo for `os-` plugins

*Score : 0 | Source : https://superuser.com/questions/1787898/register-freebsd-opnsense-repo-for-os-plugins*

I have a community version of OPNsense runnung. I want to install the package os-wireguard-go on my machine. I found out that plugins will only be installable with a pro license. I also found this repo:
https://pkg.opnsense.org/FreeBSD:13:amd64/23.1/latest/All/
which seem to have all plungins available. I would like to add it to the
/usr/local/etc/pkg/repos/OPNsense.conf
But I get a certificate error when doing so.
I would like to do something like:
fetch -o /usr/local/share/opnsense/pkg/fingerprints/pkg.opnsense.org.pub 
https://pkg.opnsense.org/pub/pkg.opnsense.org.pub
to obtain that fingerprint. I'm stuck here.
Could I also just fetch the https://pkg.opnsense.org/FreeBSD:13:amd64/23.1/latest/All/os-wireguard-go-1.13_5.pkg file and install it or is this something stupid to do?
The docs somehow say that the tooling uses the FreeBSD repo, which is disabled, but it is super confusing.

---

### Reponse — score 1

I created a new .conf called OPNsenseCore.conf and added the following lines and it worked:
OPNsenseCore: {
  fingerprints: "/usr/local/etc/pkg/fingerprints/OPNsense",
  url: "pkg+https://pkg.opnsense.org/${ABI}/23.1/latest",
  signature_type: "fingerprints",
  mirror_type: "srv",
  priority: 11,
  enabled: yes
}
the pgk update and you can see all plugins in the OPNsense WebUI.
Downside is that the original version is 23.4 and the repo only offers to 23.1. Maybe someone can comment on this. I somehow saw that opnsense-patch -c plugin should somehow make it possible to update to the latest version. Also maybe someone could comment on that.
