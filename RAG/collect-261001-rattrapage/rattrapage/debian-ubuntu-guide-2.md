---
id: collect-261001-rattrapage/rattrapage/debian-ubuntu-guide-2
title: "Debian & Ubuntu — Guide ultra-complet d'administration système"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/debian_ubuntu_guide.md
source_anchor: ""
source_lines: [255, 559]
sha256: f05e025cc0283124df3a1800d579ca2453de5d4d7677d840dce59c9e8a37a93d
---

# Debian & Ubuntu — Guide ultra-complet d'administration système

### 10.3 Aller plus loin
- **fail2ban** (§12), **2FA** (libpam-google-authenticator),
  bastion pour les accès externes.
- `~/.ssh/config` côté client : alias, clés par hôte.
---

## 11. Pare-feu — UFW et nftables

### 11.1 UFW (Ubuntu, simple)
```bash
apt install ufw
ufw default deny incoming
ufw default allow outgoing
ufw allow 22/tcp
ufw allow 80,443/tcp
ufw allow from 192.168.1.0/24 to any port 5432  # PostgreSQL local
ufw enable && ufw status verbose
```

### 11.2 nftables (Debian, le standard actuel)
```bash
apt install nftables
# /etc/nftables.conf — exemple minimal
table inet filter {
  chain input {
    type filter hook input priority 0; policy drop;
    ct state established,related accept
    iif lo accept
    tcp dport 22 accept
    tcp dport {80, 443} accept
  }
  chain forward { type filter hook forward priority 0; policy drop; }
  chain output { type filter hook output priority 0; policy accept; }
}
systemctl enable --now nftables
```
- iptables est l'ancienne génération (wrapper vers nftables).

---

## 12. fail2ban — anti brute-force

```bash
apt install fail2ban
# /etc/fail2ban/jail.local
[DEFAULT]
bantime = 1h
findtime = 10m
maxretry = 5

[sshd]
enabled = true
```
```bash
fail2ban-client status sshd
fail2ban-client unban <IP>   # débannir (le classique du lundi matin)
```
- Jails utiles : sshd, nginx-http-auth, postfix, recidive.

---

## 13. Réseau — Netplan (Ubuntu) vs interfaces (Debian)

### 13.1 Ubuntu — Netplan (`/etc/netplan/50-cloud-init.yaml`)
```yaml
network:
  version: 2
  ethernets:
    eth0:
      addresses: [192.168.1.10/24]
      routes:
        - to: default
          via: 192.168.1.1
      nameservers:
        addresses: [192.168.1.1, 8.8.8.8]
```
```bash
netplan try      # teste 120 s puis revient en arrière (sauve des vies)
netplan apply
```

### 13.2 Debian — `/etc/network/interfaces`
```
auto eth0
iface eth0 inet static
    address 192.168.1.10/24
    gateway 192.168.1.1
    dns-nameservers 192.168.1.1 8.8.8.8
```
```bash
ifup eth0 / ifdown eth0
```

### 13.3 Commandes communes
```bash
ip addr ; ip route ; ip -s link
ss -tlnp            # ports en écoute (remplace netstat)
ping -c3 8.8.8.8 ; traceroute / mtr
resolvectl status   # DNS (systemd-resolved)
nmcli               # si NetworkManager
```

---

## 14. Temps et DNS

### 14.1 NTP — chrony (recommandé)
```bash
apt install chrony
# /etc/chrony/chrony.conf : server ntp.local iburst
chronyc tracking ; chronyc sources
timedatectl set-ntp true
```
- Un serveur avec une horloge fausse = **Kerberos, TLS, logs et
  clustering cassés**.

### 14.2 DNS
- Client : `/etc/resolv.conf` (géré par systemd-resolved ou
  résolu via Netplan).
- Serveur local : **dnsmasq** (simple) ou BIND9 (complet) — §39.

---

## 15. Stockage — partitionnement et LVM

### 15.1 LVM — le standard serveur
```bash
pvcreate /dev/sdb
vgcreate vg_data /dev/sdb
lvcreate -L 50G -n lv_partage vg_data
mkfs.ext4 /dev/vg_data/lv_partage
# /etc/fstab :
/dev/vg_data/lv_partage /partage ext4 defaults 0 2
mount -a
```

### 15.2 Opérations courantes
```bash
lvextend -r -L +20G /dev/vg_data/lv_partage  # agrandir (à chaud !)
lvreduce -r -L -10G ...                       # réduire (démonter d'abord)
vgextend vg_data /dev/sdc                     # ajouter un disque
pvs ; vgs ; lvs                               # état
lvcreate -s -L 5G -n snap /dev/vg_data/lv_partage  # snapshot
```

### 15.3 RAID logiciel — mdadm
```bash
apt install mdadm
mdadm --create /dev/md0 --level=1 --raid-devices=2 /dev/sdb /dev/sdc
cat /proc/mdstat          # suivi de la resync
mdadm --detail /dev/md0
```
- RAID1 (miroir) pour le système, RAID5/6 pour le stockage.
- **Le RAID n'est pas une sauvegarde** (§24).

### 15.4 fstab — les options qui comptent
```
/dev/... /data ext4 defaults,noatime 0 2
UUID=xxx  /backup ext4 defaults,nofail 0 2   # nofail = boot même si absent
//srv/partage /mnt cifs credentials=/root/.smbcred,iocharset=utf8 0 0
```
- Après modif : `mount -a` puis `findmnt --verify`.

---

## 16. Systèmes de fichiers

- **ext4** : défaut, solide, partout.
- **XFS** : gros volumes, hautes performances (défaut RHEL).
- **btrfs/ZFS** : snapshots, checksums (ZFS via `zfs-dkms`, licence
  à noter).
```bash
df -h ; du -sh /var/* | sort -rh | head   # où part l'espace ?
tune2fs -l /dev/sda1 | grep -i reserved    # 5 % réservés root par défaut
```

---

## 17. Quotas disque

```bash
apt install quota
# fstab : ajouter usrquota,grpquota, remonter
quotacheck -cum /
quotaon -a
setquota -u prenom 10G 12G 0 0 /
repquota -a
```
- Indispensable sur les partages multi-utilisateurs.

---

## 18. Mémoire et swap

```bash
free -h ; vmstat 1 ; cat /proc/meminfo
swapon --show
```
- Serveur moderne : **zram** ou petit swap fichier (2–4 Go) plutôt
  qu'une partition.
```bash
fallocate -l 4G /swapfile && chmod 600 /swapfile
mkswap /swapfile && swapon /swapfile
# fstab : /swapfile none swap sw 0 0
```
- `vm.swappiness=10` (sysctl) : préférer la RAM au swap.

---

## 19. Performance — les outils

```bash
htop            # CPU/RAM par processus
iotop           # qui écrit sur le disque ?
nethogs         # qui utilise le réseau ?
vmstat 1 ; iostat -x 1   # (paquet sysstat)
sar -u ; sar -r          # historique (sysstat)
uptime ; cat /proc/loadavg
```
- **Load average** : ~1 par cœur = sain ; > 2× cœurs = investiguer.

---

## 20. sysctl — tuning kernel

```bash
# /etc/sysctl.d/99-perso.conf
net.ipv4.ip_forward=1              # routeur/VPN
net.core.somaxconn=4096            # serveurs web chargés
vm.swappiness=10
fs.file-max=100000
net.ipv4.tcp_syncookies=1          # anti SYN flood
```
```bash
sysctl --system && sysctl -a | grep ...
```
---

## 21. Supervision — voir avant la panne

### 21.1 Les niveaux
1. **Local** : htop, logs, `systemctl --failed` (visites).
2. **Centralisé** : Zabbix / Prometheus+Grafana / Netdata.
3. **Alertes** : SMS/mail sur seuils (disque > 85 %, charge, service down).

### 21.2 Netdata (rapide à déployer)
```bash
# 1 ligne, tableau de bord temps réel sur :19999
wget -O /tmp/netdata-kickstart.sh https://get.netdata.cloud/kickstart.sh && sh /tmp/netdata-kickstart.sh
```

### 21.3 Zabbix agent (le classique entreprise)
```bash
apt install zabbix-agent2
# /etc/zabbix/zabbix_agent2.conf : Server=IP_ZABBIX
systemctl enable --now zabbix-agent2
```
- Templates : Linux, systemd, disques, réseau.

### 21.4 Ce qu'il faut superviser (minimum)
- Ping/disponibilité, charge CPU, RAM, swap.
- Disques : % utilisé **et** inodes (`df -i`).
- Services critiques (systemd).
- Certificats TLS (expiration).
- Mises à jour en attente (`apt`).
- Température (salle serveurs).

---

## 22. Logs centralisés — rsyslog et Loki

```bash
# Client rsyslog → serveur central
# /etc/rsyslog.d/49-remote.conf
*.* @@192.168.1.5:514
systemctl restart rsyslog
```
- Serveur : rsyslog + **Loki** (Grafana) ou Graylog.
- **Pourquoi** : un pirate efface les logs locaux ; les logs
  distants racontent l'attaque. Et le dépannage multi-serveurs
  devient simple.

---

## 23. Sauvegardes — la discipline

### 23.1 La règle 3-2-1
- **3** copies, **2** supports différents, **1** hors site.

### 23.2 Borg (le meilleur rapport simplicité/efficacité)
```bash
apt install borgbackup
borg init --encryption=repokey /backup/repo
borg create /backup/repo::srv01-{now:%Y-%m-%d} /etc /home /var/www \
  --exclude /home/*/.cache
borg list /backup/repo
borg prune --keep-daily=7 --keep-weekly=4 --keep-monthly=6 /backup/repo
```
- Dédupliqué, chiffré, incrémental. **Le choix par défaut.**

### 23.3 Alternatives
- **restic** : comme Borg, vers S3/cloud (`restic -r s3:...`).
- **rsync** : simple miroir (`rsync -aAXv --delete src/ dst/`).
- **Bacula/Bareos** : gros parcs, pilotage central.

