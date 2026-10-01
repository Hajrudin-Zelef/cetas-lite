---
id: collect-261001-rattrapage/rattrapage/debian-ubuntu-guide-6
title: "Debian & Ubuntu — Guide ultra-complet d'administration système"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "attribution", "benchmark"]
source: docs/RAG/collect-261001-rattrapage/debian_ubuntu_guide.md
source_anchor: ""
source_lines: [1334, 1602]
sha256: c38591b5cc5550762f690e3913823da3a4c5b71200fea61939cb71d382bdca89
---

# Debian & Ubuntu — Guide ultra-complet d'administration système

- Guide **Proxmox VE** approfondi (cluster, Ceph, HA, vzdump).
- Guide **Zabbix** : installation, templates, alerting SMS.
- Guide **Ansible** : rôles, inventaires dynamiques, AWX.
- Guide **Bareos** : déploiement complet.
- Durcissement **CIS Benchmark** : checklist exhaustive.
- Guide **GLPI** : déploiement + bonnes pratiques SAV.

---

*Fin du guide — 1600+ lignes. Administrer, c'est **automatiser,
superviser, sauvegarder, documenter**. Avec ça, ton volet systèmes
est aussi solide que ton volet énergies, Zelef.*
---

## 71. systemd en profondeur

### 71.1 Drop-ins (sans modifier le paquet)
```bash
systemctl edit nginx   # crée /etc/systemd/system/nginx.service.d/override.conf
```
```ini
[Service]
LimitNOFILE=65536
Restart=always
```

### 71.2 Sandboxing — durcir un service
```ini
[Service]
ProtectSystem=strict
ProtectHome=true
PrivateTmp=true
NoNewPrivileges=true
```
- `systemd-analyze security nginx` : score de durcissement.

### 71.3 Debug
```bash
systemd-analyze blame          # qui ralentit le boot ?
systemd-analyze critical-chain
systemctl status --no-pager -l
```

---

## 72. Réseau avancé — bonding, VLAN, bridge

```bash
# Bonding (Debian /etc/network/interfaces)
auto bond0
iface bond0 inet static
    address 192.168.1.10/24
    bond-mode 802.3ad
    bond-miimon 100
    bond-slaves eth0 eth1

# VLAN (paquet vlan)
auto eth0.10
iface eth0.10 inet static
    address 10.10.10.10/24
    vlan-raw-device eth0

# Bridge (KVM)
auto br0
iface br0 inet static
    address 192.168.1.10/24
    bridge_ports eth0
```
- Netplan (Ubuntu) : `bonds:`, `vlans:`, `bridges:` au même niveau
  que `ethernets:`.

---

## 73. Script d'audit mensuel du parc

```bash
#!/bin/bash
# /usr/local/bin/audit-mensuel.sh — à lancer via Ansible sur tout le parc
echo "=== AUDIT $(hostname) $(date +%F) ==="
echo "[OS] $(lsb_release -ds) — kernel $(uname -r)"
echo "[MAJ] $(apt list --upgradable 2>/dev/null | wc -l) paquets en attente"
echo "[SERVICES EN ÉCHEC]"; systemctl list-units --failed --no-legend
echo "[DISQUES]"; df -h / /var | awk 'NR>1{print $1" "$5" utilisé"}'
echo "[INODES]"; df -i / /var | awk 'NR>1{print $1" "$5}'
echo "[ÉCOUTE]"; ss -tlnp | awk 'NR>1{print $4}' | sort -u
echo "[DERNIERS LOGINS]"; last -n 5
echo "[CERTIFICATS]"; for c in /etc/ssl/certs/*.pem; do
  openssl x509 -enddate -noout -in "$c" 2>/dev/null; done | sort | head -5
```
- Sortie centralisée (mail, fichier partagé) → revue mensuelle
  en 30 min pour tout le parc.

---

## 74. Sécurité avancée

### 74.1 Jails fail2ban personnalisées
```ini
# /etc/fail2ban/jail.local
[nginx-badbot]
enabled = true
filter = nginx-badbot
logpath = /var/log/nginx/access.log
maxretry = 2
```

### 74.2 Portsentry / PSAD — détection de scan
- `portsentry` : bannit les IP qui scannent les ports.
- À n'activer qu'après les bases (pare-feu, fail2ban).

### 74.3 Mises à jour du kernel sans reboot
- **Canonical Livepatch** (Ubuntu Pro) : patchs kernel à chaud.
- Utile sur les serveurs qu'on ne peut pas rebooter souvent.

---

## 75. Cas pratiques supplémentaires

**Cas 11 — `apt update` : « Release file expired »**
Horloge fausse (NTP mort) → `timedatectl`, chrony, puis réessayer.

**Cas 12 — Docker remplit le disque**
`docker system df`, `docker system prune -a --volumes`
(attention : supprime les volumes non utilisés !), logs json
sans rotation → `log-opts` dans daemon.json.

**Cas 13 — NFS qui se fige**
`soft` vs `hard` dans les options de montage, `intr`, timeout ;
serveur NFS surchargé → iostat côté serveur.

**Cas 14 — Certificat expiré un dimanche**
Supervision des expirations (§21.4) + `certbot renew --dry-run`
en test mensuel. Renouvellement auto vérifié, pas supposé.

**Cas 15 — Mot de passe root perdu (accès physique)**
Boot → grub → `init=/bin/bash` → `mount -o remount,rw /` →
`passwd` → reboot. (D'où l'importance du chiffrement LUKS et
de la sécurité physique.)

---

## 76. Le mot du chef de service systèmes

```
Un parc sain, c'est :
- des serveurs PATCHÉS (unattended-upgrades + fenêtre mensuelle),
- des sauvegardes TESTÉES (3-2-1, restauration 2×/an),
- des accès TRACÉS (1 compte/personne, sudo, logs distants),
- des configs VERSIONNÉES (Ansible, pas de bricolage SSH),
- des docs À JOUR (fiche par serveur),
- et une ASTREINTE qui dort tranquille (supervision + alertes).
```

---

*Fin du guide — 1600+ lignes. Systèmes + énergies + copieurs :
ton app couvre maintenant tout ton périmètre, Zelef.*
---

## 77. Fiche réflexe — commandes de survie par situation

```
SERVEUR LENT :
  htop → iotop → vmstat 1 → dmesg | tail → df -h

SERVICE DOWN :
  systemctl status X → journalctl -u X -n 50 → config -t → ss -tlnp

RÉSEAU COUPÉ :
  ip addr → ip route → ping GW → resolvectl status → ufw/nftables

DISQUE PLEIN :
  df -h → du -sh /* → journalctl --vacuum-size → lsof | grep deleted

APRÈS INTRUSION SUSPECTÉE :
  isoler → last → ausearch → aide --check → logs distants → reconstruire

AVANT TOUTE INTERVENTION RISQUÉE :
  snapshot (LVM/Borg/Proxmox) → backup → fenêtre → console dispo
```

---

## 78. Documentation officielle — les références

- Debian : debian.org/doc (Administrator's Handbook — gratuit)
- Ubuntu : ubuntu.com/server/docs
- systemd : freedesktop.org (man systemd.unit, systemd.service…)
- Ansible : docs.ansible.com
- Toujours `man <commande>` d'abord : la doc locale est la plus fiable.

---

## 79. Mini-projet — monter un serveur complet (fil rouge)

```
Jour 1 : Installation Ubuntu 24.04 (LVM, /var séparé), checklist §3
Jour 2 : SSH durci + UFW + fail2ban + unattended-upgrades
Jour 3 : LVM data + Samba + quotas (partage bureau)
Jour 4 : CUPS + imprimante réseau partagée
Jour 5 : Borg vers NAS + test de restauration
Jour 6 : Zabbix agent + alertes + Netdata
Jour 7 : Ansible (playbook de base) + fiche d'exploitation §47
```
- En une semaine, un serveur **prod-ready**, documenté, supervisé,
  sauvegardé. C'est le standard à viser pour chaque machine du parc.

---

*1600+ lignes. Guide terminé.*
---

## 80. Les services et ports à connaître par cœur

| Port | Service | Protocole | Usage |
|---|---|---|---|
| 22 | SSH | TCP | Administration |
| 53 | DNS | TCP/UDP | Résolution |
| 67/68 | DHCP | UDP | Attribution IP |
| 80/443 | HTTP/HTTPS | TCP | Web |
| 123 | NTP | UDP | Heure |
| 445 | SMB | TCP | Partages Windows |
| 2049 | NFS | TCP | Partages Linux |
| 3306 | MySQL/MariaDB | TCP | Base de données |
| 5432 | PostgreSQL | TCP | Base de données |
| 631 | IPP/CUPS | TCP | Impression |
| 514 | Syslog | TCP/UDP | Logs distants |
| 51820 | WireGuard | UDP | VPN |
| 5665/10050 | Zabbix | TCP | Supervision |
| 9090 | Cockpit | TCP | Admin web |
| 8006 | Proxmox | TCP | Admin web |

---

*Fin du guide — 1600+ lignes.*
---

## 81. Checklist de départ — nouveau serveur (à cocher)

```
□ OS installé (LVM, /var séparé), hostname FQDN, timezone
□ IP fixe, DNS, NTP synchronisé
□ apt full-upgrade + unattended-upgrades configuré
□ Utilisateur admin (clé SSH), root SSH interdit, sudo
□ UFW/nftables actif, fail2ban actif
□ Zabbix agent + Netdata, alertes testées
□ Borg configuré, 1re sauvegarde OK, test de restauration planifié
□ Logs distants configurés
□ Playbook Ansible à jour (serveur dans l'inventaire)
□ Fiche d'exploitation remplie (§47)
□ GLPI : fiche créée
□ Mots de passe dans le coffre
```

---

*Guide terminé — 1600+ lignes.*

---

## 82. Pour aller plus loin

- Guide **Proxmox VE** : cluster, Ceph, HA, vzdump.
- Guide **Zabbix** : templates, triggers, alerting SMS.
- Guide **Ansible** : rôles, AWX, inventaires dynamiques.
- Guide **Bareos** : déploiement complet multi-serveurs.
- Guide **GLPI** : ticketing + gestion de parc.

*1600+ lignes — fin.*
