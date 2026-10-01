---
id: collect-261001-rattrapage/rattrapage/debian-ubuntu-guide-4
title: "Debian & Ubuntu — Guide ultra-complet d'administration système"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/debian_ubuntu_guide.md
source_anchor: ""
source_lines: [835, 1098]
sha256: fa28eff22edbff29737c3a9e8ab82b5280f411d537b227c2cf7a7bcf2aaf1f29
---

# Debian & Ubuntu — Guide ultra-complet d'administration système

```bash
apt install postfix   # type : "Site Internet" ou "Satellite"
# /etc/postfix/main.cf : relayhost = [smtp.fai]:587 + SASL
echo "test" | mail -s "Test" admin@example.com
mailq ; postsuper -d ALL   # file d'attente
```
- Un serveur mail **complet** (anti-spam, etc.) est un métier :
  pour les alertes, un relais SMTP suffit. Pour l'hébergement mail,
  envisagez un prestataire.

---

## 40. Cron et systemd timers

```bash
crontab -e
# m h dom mon dow commande
0 2 * * * /usr/local/bin/backup.sh >> /var/log/backup.log 2>&1
```
- Scripts cron : **shebang, chemins absolus, log, verrou**
  (`flock -n` contre les chevauchements).
- Alternative moderne : systemd timer (§7.3).

---

## 41. Mises à niveau de version

### 41.1 Ubuntu LTS → LTS
```bash
do-release-upgrade   # jamais avec -d en prod
```
- Snapshot/backup **avant**, fenêtre de maintenance, jamais à distance
  sans console (IPMI/iDRAC).

### 41.2 Debian Stable → Stable
```bash
sed -i 's/bookworm/trixie/g' /etc/apt/sources.list*
apt update && apt full-upgrade
```
- Lire les **release notes** (changements majeurs : ex. merged-/usr).
- Même discipline : backup, fenêtre, console.

---

## 42. Snap (Ubuntu) — cohabitation avec APT

```bash
snap list ; snap refresh
snap set system refresh.timer=sun,02:00  # fenêtre de maj
```
- Snaps : paquets universels, mises à jour auto, plus lourds.
- En serveur : préférer APT quand le paquet existe ; snap pour le
  reste (certbot historiquement, etc.).
---

## 43. Rescue — quand ça ne boote plus

```
1. Boot sur ISO en mode rescue (ou init=/bin/bash au grub)
2. Monter : mount /dev/vg0/lv_root /mnt ; mount --bind /dev /mnt/dev ...
   chroot /mnt
3. Réparer :
   - grub-install /dev/sda && update-grub
   - dpkg --configure -a (maj interrompue)
   - passwd (root perdu — avec accès console physique)
   - fsck (filesystem corrompu, démonté !)
4. Logs : journalctl -b -1 (boot précédent)
```
- **Accès console** (IPMI/iDRAC/KVM) : obligatoire sur tout serveur
  distant. Sans console, un grub cassé = déplacement.

---

## 44. Dépannage — 10 cas terrain

**Cas 1 — Disque plein (`/`)**
`du -sh /*`, `journalctl --vacuum-size=200M`, logs applicatifs
(`logrotate` absent ?), `/tmp`, vieux kernels
(`apt autoremove --purge`).

**Cas 2 — « No space left » mais `df` OK**
Inodes épuisés (`df -i`) — des millions de petits fichiers
(cache, sessions PHP, mails).

**Cas 3 — Service qui ne démarre pas**
`systemctl status`, `journalctl -u`, `sshd -t` / `nginx -t`
(config !), permissions, port déjà pris (`ss -tlnp`).

**Cas 4 — SSH coupé après changement réseau**
`netplan try` aurait sauvé. Console/IPMI, vérifier Netplan,
routes, pare-feu (UFW qui bloque le nouveau port…).

**Cas 5 — Charge à 50, tout est lent**
`htop` (tri CPU), `iotop` (I/O), `dmesg` (disque en train de
mourir ?), swap (`free`).

**Cas 6 — DNS qui ne résout plus**
`resolvectl status`, `/etc/resolv.conf` (lien symbolique ?),
`systemd-resolved` redémarré, pare-feu port 53.

**Cas 7 — Mise à jour qui casse le boot**
`dpkg --configure -a`, kernel précédent au grub, snapshot
(Borg/LVM) → rollback.

**Cas 8 — Samba inaccessible depuis Windows**
`testparm`, `smbstatus`, pare-feu (445/tcp), `nmbd` (vieux
Windows), identifiants (`smbpasswd`).

**Cas 9 — Imprimante CUPS bloquée**
`lpstat -o`, `cancel -a`, ping du copieur, file SNMP,
`error_log` en debug — et vérifier le copieur lui-même
(vos guides !).

**Cas 10 — Compromission suspectée**
Isoler (pare-feu sortant), `last`, `ausearch`, AIDE
(`aide --check`), logs distants (§22), **ne pas nettoyer** :
reconstruire depuis une source saine + changer tous les secrets.

---

## 45. Scripts bash utiles

```bash
#!/bin/bash
# /usr/local/bin/sante.sh — état rapide du serveur
set -euo pipefail
echo "=== $(hostname) $(date) ==="
uptime
echo "--- Disques ---" ; df -h / /var | tail -n +2
echo "--- Mémoire ---" ; free -h | head -2
echo "--- Services en échec ---" ; systemctl list-units --failed --no-legend
echo "--- Mises à jour ---" ; apt list --upgradable 2>/dev/null | wc -l
echo "--- Écoute réseau ---" ; ss -tlnp | head -20
```

```bash
#!/bin/bash
# backup rapide /etc + confs (Borg en §23 pour le sérieux)
/usr/bin/tar -czf /backup/etc-$(date +%F).tar.gz /etc /root
find /backup -name "etc-*.tar.gz" -mtime +30 -delete
```

---

## 46. Gestion d'équipe — accès et traçabilité

- **1 compte par personne**, jamais de compte partagé.
- Groupes : `sudo` (admins), `deploiement` (restart appli),
  `lecture-logs` (support).
- sudoers fin (§9.1) + logs centralisés (§22).
- Départ d'un collaborateur : `passwd -l`, supprimer la clé
  `~/.ssh/authorized_keys`, revoir les accès (checklist RH).
- Bastion SSH pour les accès externes, 2FA.

---

## 47. Documentation d'exploitation (par serveur)

```
SERVEUR : srv01 — rôle : fichiers + impression
OS : Ubuntu 24.04 LTS — Installé le : ___ par ___
IP : 192.168.1.10/24 — GW : .1 — DNS : .1
Services : samba, cups, zabbix-agent2
Particularités : LVM 2 To, RAID1 système
Sauvegarde : Borg → NAS 22h, testé le ___
Contrat : ___ — Astreinte : ___
Mots de passe : dans le coffre (jamais ici !)
```

---

## 48. KPI pour le chef de service

- **Disponibilité** par serveur/service (objectif 99,5 %+).
- **MTTR** incidents, **délai de patch** (CVE critiques < 72 h).
- **Sauvegardes** : % réussies, dernier test de restauration.
- **Parc** : OS à jour, EOL à venir (tableau des versions).
- **Rapport mensuel** : 1 page — incidents, patchs, sauvegardes,
  risques (serveurs en fin de vie, disques > 80 %).

---

## 49. Fin de vie des versions — anticiper

| Version | Fin de support standard |
|---|---|
| Debian 11 (bullseye) | 2024 (LTS 2026) |
| Debian 12 (bookworm) | ~2026 (LTS ~2028) |
| Ubuntu 20.04 LTS | 2025 (ESM 2030) |
| Ubuntu 22.04 LTS | 2027 (ESM 2032) |
| Ubuntu 24.04 LTS | 2029 (ESM 2034) |

- Planifier les migrations **18 mois avant** la fin, pas 2 mois après.
- ESM Ubuntu (Extended Security Maintenance) : un sursis, pas une
  stratégie.

---

## 50. Glossaire

| Terme | Signification |
|---|---|
| LTS | Long Term Support (5 ans) |
| APT | Gestionnaire de paquets Debian/Ubuntu |
| systemd | Init + services + logs |
| LVM | Gestionnaire de volumes logiques |
| Netplan | Config réseau (Ubuntu) |
| UFW | Pare-feu simplifié (Ubuntu) |
| nftables | Pare-feu noyau actuel |
| chroot | Changer de racine (rescue) |
| 3-2-1 | Règle des sauvegardes |
| ESM | Support étendu Ubuntu |
| Idempotence | Relançable sans effet de bord (Ansible) |
---

## 51. cloud-init — le standard du déploiement

```yaml
# user-data (Ubuntu autoinstall / cloud)
#cloud-config
hostname: srv01
manage_etc_hosts: true
users:
  - name: admin
    sudo: ALL=(ALL) NOPASSWD:ALL
    ssh_authorized_keys:
      - ssh-ed25519 AAAA... admin@poste
packages: [qemu-guest-agent, zabbix-agent2, fail2ban]
runcmd:
  - [systemctl, enable, --now, zabbix-agent2]
timezone: Africa/Abidjan
```
- Un serveur **identique, reproductible, sans intervention** :
  la base de tout parc sérieux.

---

## 52. PXE — déploiement réseau

```
Client (PXE) → DHCP (option boot) → TFTP (bootloader)
  → HTTP (ISO / autoinstall) → installation auto (cloud-init)
```
- Outils : dnsmasq (DHCP+TFTP), serveur HTTP simple.
- En parc : un nouveau serveur physique/VM = branché, allumé,
  15 min plus tard il est dans Zabbix et Ansible.

---

## 53. Proxmox VE — virtualisation (basé sur Debian)

- Hyperviseur complet : KVM + LXC, web UI (`:8006`), clustering,
  Ceph intégré, sauvegardes planifiées.
```bash
# Sur Debian 12 : dépôt Proxmox, puis apt install proxmox-ve
```
- Sauvegardes : `vzdump` planifié → NFS/Borg.
- Cluster : 3 nœuds minimum (quorum), HA des VM.
- **Excellent choix PME** : un seul outil pour tout virtualiser.

---

## 54. Haute disponibilité

