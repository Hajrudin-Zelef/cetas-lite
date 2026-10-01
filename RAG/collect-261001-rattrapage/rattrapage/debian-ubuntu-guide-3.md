---
id: collect-261001-rattrapage/rattrapage/debian-ubuntu-guide-3
title: "Debian & Ubuntu — Guide ultra-complet d'administration système"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["scout"]
source: docs/RAG/collect-261001-rattrapage/debian_ubuntu_guide.md
source_anchor: ""
source_lines: [560, 834]
sha256: 87dbf850e98f497bf31a281a0d58327884970c620307c91ec486f5cf3ae35411
---

# Debian & Ubuntu — Guide ultra-complet d'administration système

### 23.4 Ce qu'on sauvegarde (serveur type)
- `/etc` (configs), `/home`, `/var/www`, bases de données
  (**dump** avant : `pg_dump`, `mysqldump` — jamais les fichiers
  à chaud).
- **Tester la restauration 2×/an.** Une sauvegarde non testée
  n'existe pas.

---

## 24. KVM — virtualisation

```bash
apt install qemu-kvm libvirt-daemon-system virtinst cockpit
virsh list --all
virt-install --name vm01 --ram 4096 --vcpus 2 \
  --disk size=20 --os-variant ubuntu24.04 \
  --network bridge=br0 --graphics none \
  --location /iso/ubuntu-24.04.iso --extra-args "console=ttyS0"
virsh start|shutdown|snapshot-create vm01
```
- **Cockpit** (`:9090`) : gestion web simple (VM, stockage, réseau).
- Bridge réseau pour des VM joignables sur le LAN.
- Snapshots avant toute mise à jour risquée.

---

## 25. Docker — conteneurs

```bash
# Installation officielle (dépôt Docker, pas le paquet snap)
apt install docker.io docker-compose-plugin   # simple (dépôt Ubuntu/Debian)
# ou dépôt officiel docker.com pour les dernières versions
docker run -d --name web --restart unless-stopped -p 8080:80 nginx
docker compose up -d   # avec compose.yaml
docker ps ; docker logs -f web ; docker exec -it web sh
```
- Données : **volumes** (`-v data:/var/lib/...`), jamais dans le
  conteneur.
- Mises à jour : `docker compose pull && docker compose up -d`.
- Sécurité : ne pas exposer le socket Docker, user non-root dans
  les images, scanner (`docker scout`).

---

## 26. LXC/LXD — conteneurs système

- Plus légers que les VM, plus « système » que Docker.
- `lxd init`, `lxc launch ubuntu:24.04 c1`, `lxc exec c1 -- bash`.
- Idéal : un service par conteneur, snapshots instantanés.

---

## 27. Ansible — automatisation du parc

```yaml
# playbook.yml
- hosts: serveurs
  become: true
  tasks:
    - name: Mises à jour de sécurité
      apt: upgrade=dist update_cache=yes
    - name: fail2ban installé
      apt: name=fail2ban state=present
    - name: SSH durci
      lineinfile:
        path: /etc/ssh/sshd_config
        regexp: '^PermitRootLogin'
        line: 'PermitRootLogin no'
      notify: reload sshd
  handlers:
    - name: reload sshd
      service: name=ssh state=reloaded
```
```bash
ansible-playbook -i inventaire playbook.yml
```
- **Idempotence** : on peut relancer sans risque.
- En parc de 10+ serveurs, Ansible n'est plus une option.

---

## 28. Durcissement — la checklist sécurité

```
□ Mises à jour auto (security) — §6
□ SSH clés uniquement, root interdit — §10
□ Pare-feu actif, ports minimaux — §11
□ fail2ban — §12
□ Utilisateurs : pas de compte partagé, sudo tracé — §9
□ AppArmor en enforcing (aa-status)
□ Pas de service inutile : systemctl list-unit-files --state=enabled
□ Mots de passe : shadow, pas de compte sans mot de passe inutile
□ /tmp en noexec (fstab : tmpfs /tmp tmpfs defaults,noexec,nosuid 0 0)
□ Bannières : pas de version affichée (ServerTokens Prod)
□ Sauvegardes chiffrées hors site — §23
□ Logs centralisés — §22
```

---

## 29. AppArmor (Debian & Ubuntu)

```bash
aa-status                    # profils chargés
aa-enforce /etc/apparmor.d/usr.sbin.nginx
aa-complain ...              # mode permissif (debug)
```
- Ubuntu : profils actifs par défaut (plus strict).
- Si une appli est bloquée : `dmesg | grep -i apparmor` avant de
  tout désactiver.

---

## 30. auditd et AIDE — traçabilité et intégrité

```bash
apt install auditd audispd-plugins
auditctl -w /etc/shadow -p wa -k identifiants
ausearch -k identifiants
```
```bash
apt install aide
aideinit   # base de référence
aide --check   # détecte les fichiers modifiés (intrusion ?)
```
- AIDE : la base **hors serveur** (sinon l'attaquant la modifie).
---

## 31. Serveur web — nginx + TLS

```bash
apt install nginx certbot python3-certbot-nginx
# /etc/nginx/sites-available/monsite
server {
    listen 80;
    server_name monsite.local;
    root /var/www/monsite;
    location / { try_files $uri $uri/ =404; }
}
ln -s ../sites-available/monsite /etc/nginx/sites-enabled/
nginx -t && systemctl reload nginx
certbot --nginx -d monsite.example.com   # Let's Encrypt (auto-renew via timer)
```
- Renouvellement : timer systemd `certbot.timer` (vérifier).
- Reverse proxy vers une appli :
```
location / { proxy_pass http://127.0.0.1:3000; proxy_set_header Host $host; }
```

---

## 32. PHP / Python / Node — les runtimes

- **PHP** : `apt install php-fpm php-mysql`, pool par site
  (`/etc/php/*/fpm/pool.d/`), nginx → `fastcgi_pass`.
- **Python** : venv par projet (`python3 -m venv`), **jamais** de
  pip en root sur le système (PEP 668 : `--break-system-packages`
  interdit en prod).
- **Node** : via NodeSource ou nvm (pas le paquet système ancien).

---

## 33. PostgreSQL

```bash
apt install postgresql
sudo -u postgres psql
# CREATE USER monapp WITH PASSWORD '...'; CREATE DATABASE monapp OWNER monapp;
# /etc/postgresql/*/main/postgresql.conf : listen_addresses
# /etc/postgresql/*/main/pg_hba.conf : host monapp monapp 192.168.1.0/24 scram-sha-256
pg_dump monapp | gzip > /backup/monapp-$(date +%F).sql.gz   # dump quotidien (cron)
```
- Sauvegarde : dump + archivage WAL pour la PITR (point-in-time recovery).

---

## 34. MariaDB / MySQL

```bash
apt install mariadb-server
mysql_secure_installation
mysqldump --all-databases | gzip > /backup/mysql-$(date +%F).sql.gz
```
- `mysqltuner` : script de réglage (à relancer après quelques jours
  de charge réelle).

---

## 35. Partages fichiers — Samba et NFS

### 35.1 Samba (Windows)
```bash
apt install samba
# /etc/samba/smb.conf
[partage]
  path = /srv/partage
  valid users = @bureau
  read only = no
  create mask = 0660
smbpasswd -a prenom
systemctl restart smbd
```
- **Corbeille réseau** : module `recycle` (les suppressions
  accidentelles…).
- Intégration AD : `realm` / winbind (parc Windows).

### 35.2 NFS (Linux)
```bash
apt install nfs-kernel-server
# /etc/exports : /srv/data 192.168.1.0/24(rw,sync,no_subtree_check)
exportfs -ra
# client : mount srv:/srv/data /mnt
```
- NFSv4 + Kerberos pour la sécurité ; sinon réseau de confiance
  uniquement.

---

## 36. Serveur d'impression CUPS — ⚠️ votre métier !

```bash
apt install cups
# https://srv:631 (interface web)
lpadmin -p Kyocera-TASKalfa -E -v socket://192.168.1.50 -m everywhere
cupsenable Kyocera-TASKalfa ; cupsaccept Kyocera-TASKalfa
lpstat -p -d
```
- Pilote **générique IPP Everywhere** (`-m everywhere`) : fonctionne
  avec la plupart des copieurs récents sans pilote propriétaire.
- File d'attente bloquée : `cancel -a`, vérifier le copieur (papier ?
  erreur ? — voir vos guides copieurs !).
- **Samba + CUPS** : partage des imprimantes vers Windows
  (`printing = cups`, `print$` pour les pilotes).
- Logs : `/var/log/cups/error_log` (niveau debug : `LogLevel debug`).

---

## 37. DNS/DHCP — dnsmasq et BIND

```bash
apt install dnsmasq
# /etc/dnsmasq.conf
interface=eth0
dhcp-range=192.168.1.100,192.168.1.200,24h
address=/srv01.local/192.168.1.10
systemctl restart dnsmasq
```
- dnsmasq = DNS + DHCP en un paquet, parfait pour un site PME.
- BIND9 pour les zones complexes / vues.

---

## 38. VPN — WireGuard

```bash
apt install wireguard
wg genkey | tee /etc/wireguard/privatekey | wg pubkey > /etc/wireguard/publickey
# /etc/wireguard/wg0.conf
[Interface]
PrivateKey = ...
Address = 10.8.0.1/24
ListenPort = 51820
PostUp = ufw route allow in on wg0 out on eth0
[Peer]
PublicKey = ...
AllowedIPs = 10.8.0.2/32
wg-quick up wg0 && systemctl enable wg-quick@wg0
```
- Moderne, rapide, simple : **le choix par défaut** (OpenVPN si
  contrainte client).

---

## 39. Postfix — mail de base (notifications)

