---
id: collect-261001-rattrapage/rattrapage/debian-ubuntu-guide-1
title: "Debian & Ubuntu — Guide ultra-complet d'administration système"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/debian_ubuntu_guide.md
source_anchor: ""
source_lines: [1, 254]
sha256: e4e3f47c32a25d0aaf2897612d7aa59b1314f9cb89e3c17c8b21c440ac5ade8b
---

# Debian & Ubuntu — Guide ultra-complet d'administration système

> Tout pour administrer des serveurs Debian et Ubuntu en production :
> installation, APT, systemd, réseau, sécurité, SSH, stockage LVM/RAID,
> supervision, sauvegardes, KVM, Docker, Ansible, dépannage —
> **angle chef de service systèmes & énergies**.
>
> Les commandes et chemins varient légèrement selon les versions
> (Debian 12/13, Ubuntu 22.04/24.04 LTS) : **validez sur la version
> installée** (`lsb_release -a`, `cat /etc/os-release`).

---

## 1. Debian vs Ubuntu — que choisir ?

| Critère | Debian | Ubuntu |
|---|---|---|
| Philosophie | 100 % libre, communautaire | Commercial (Canonical), plus « clés en main » |
| Cycle | Stable ~2 ans (bookworm 12, trixie 13) | LTS tous les 2 ans (22.04, 24.04), support 5 ans (+10 ESM) |
| Paquets | Stables mais anciens | Plus récents sur LTS |
| Par défaut | Minimaliste | Plus d'outils préinstallés (cloud-init, snap) |
| Idéal pour | Serveurs stables, Proxmox, routeurs | Serveurs d'entreprise, cloud, postes |

- **En entreprise : Ubuntu LTS ou Debian Stable.** Jamais d'interim
  (6 mois) ni de Debian Testing en production.
- Les deux partagent **APT, systemd, 95 % des commandes** : ce guide
  couvre les deux, les différences sont signalées.

---

## 2. Installation — bien partir

### 2.1 Quel ISO ?
- **Debian** : netinstall (~700 Mo) — tout le reste vient du réseau.
- **Ubuntu Server** : ISO live-server (~2 Go).
- Toujours vérifier la **somme SHA256** de l'ISO.

### 2.2 Partitionnement recommandé (serveur)
```
 /boot/efi : 512 Mo (UEFI)
/boot      : 1 Go (ext4)
LVM :
  /        : 20–30 Go
  /var     : 20 Go (logs !)
  /home    : 10 Go
  swap     : 2–4 Go (ou zram)
  libre    : le reste (croissance, snapshots)
```
- **Séparer /var** : des logs qui explosent ne doivent jamais
  remplir la racine.

### 2.3 Automatisation
- **Debian** : preseed (`preseed.cfg`).
- **Ubuntu** : **cloud-init** / autoinstall (`user-data`).
- En parc : PXE + autoinstall = un serveur prêt en 15 min.

---

## 3. Premier boot — checklist

```
□ Hostname FQDN : hostnamectl set-hostname srv01.domaine.local
□ Fuseau horaire : timedatectl set-timezone Africa/Abidjan
□ NTP actif : timedatectl status (NTP synchronized: yes)
□ Réseau : IP fixe, passerelle, DNS (voir §13)
□ Mises à jour : apt update && apt full-upgrade
□ unattended-upgrades : installé et configuré (§6)
□ SSH : clés uniquement, root interdit (§10)
□ Pare-feu : UFW/nftables actif (§11)
□ Utilisateur admin + sudo (§9)
□ Supervision : agent installé (§22)
□ Sauvegarde : planifiée (§24)
```

---

## 4. APT en profondeur

### 4.1 Les commandes
```bash
apt update              # rafraîchir les listes
apt full-upgrade        # mise à jour complète (gère les dépendances)
apt install -y paquet
apt remove paquet       # garde la config
apt purge paquet        # supprime tout
apt autoremove --purge  # nettoyer les orphelins
apt search / apt show paquet
apt-mark hold paquet    # bloquer une version (kernel, appli critique)
apt list --upgradable
```

### 4.2 Les sources (`/etc/apt/sources.list` ou `sources.list.d/`)
```bash
# Debian 12 (bookworm) — exemple
deb http://deb.debian.org/debian bookworm main contrib non-free non-free-firmware
deb http://security.debian.org/debian-security bookworm-security main
deb http://deb.debian.org/debian bookworm-updates main
# Ubuntu 24.04 — exemple
deb http://archive.ubuntu.com/ubuntu noble main restricted universe multiverse
deb http://archive.ubuntu.com/ubuntu noble-security main restricted universe multiverse
```
- **Ne jamais mélanger** les dépôts de deux versions.
- Miroir local/proche : choisir le plus rapide (`netselect-apt`).

### 4.3 Backports (Debian) — paquets récents sur Stable
```bash
echo "deb http://deb.debian.org/debian bookworm-backports main" \
  > /etc/apt/sources.list.d/backports.list
apt update && apt -t bookworm-backports install paquet
```

### 4.4 Pinning (`/etc/apt/preferences.d/`)
- Épingler une version d'un dépôt spécifique (ex. : kernel d'une
  version, le reste de stable). À manier avec précaution.

---

## 5. dpkg — quand APT ne suffit pas

```bash
dpkg -l | grep paquet        # lister
dpkg -L paquet               # fichiers installés
dpkg -S /chemin/fichier      # quel paquet fournit ce fichier ?
dpkg -i paquet.deb           # installer un .deb local
dpkg --configure -a          # réparer une install interrompue
debsums -c                   # vérifier l'intégrité des fichiers
```

---

## 6. Mises à jour automatiques — unattended-upgrades

```bash
apt install unattended-upgrades
dpkg-reconfigure -plow unattended-upgrades
```
- `/etc/apt/apt.conf.d/50unattended-upgrades` : choisir les origines
  (security **toujours**, updates selon politique).
- Reboot automatique si nécessaire : `Unattended-Upgrade::Automatic-Reboot "true";`
  + heure creuse (`Automatic-Reboot-Time "03:00"`).
- **Politique de chef de service** : security = auto ; le reste =
  fenêtre de maintenance mensuelle planifiée.
- Surveiller : `/var/log/unattended-upgrades/`.

---

## 7. systemd — le cœur du système

### 7.1 Services
```bash
systemctl status nginx
systemctl start|stop|restart|reload nginx
systemctl enable|disable nginx      # au boot
systemctl is-enabled nginx
systemctl list-units --failed       # services en échec !
systemctl cat nginx                 # voir l'unité
```

### 7.2 Créer un service (`/etc/systemd/system/monapp.service`)
```ini
[Unit]
Description=Mon application
After=network.target

[Service]
User=monapp
ExecStart=/opt/monapp/start.sh
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```
```bash
systemctl daemon-reload && systemctl enable --now monapp
```

### 7.3 Timers (remplacent cron pour les services)
```bash
systemctl list-timers --all
```
- Plus robustes que cron (journalisation, dépendances, `OnCalendar`).

### 7.4 Targets (niveaux)
```bash
systemctl get-default          # graphical.target ou multi-user.target
systemctl set-default multi-user.target   # serveur sans GUI
```

---

## 8. journald — les logs système

```bash
journalctl -u nginx                    # logs d'un service
journalctl -u nginx --since "1 hour ago"
journalctl -f                          # suivi temps réel
journalctl -p err -b                   # erreurs depuis le boot
journalctl --disk-usage                # taille des logs
```
- Config : `/etc/systemd/journald.conf` (`SystemMaxUse=500M`).
- **Logs persistants** : `mkdir /var/log/journal` (sinon en RAM).

---

## 9. Utilisateurs, groupes, sudo

```bash
adduser prenom            # crée user + home (Debian/Ubuntu)
usermod -aG sudo prenom   # groupe sudo (Ubuntu) / sudo (Debian)
usermod -aG docker prenom
passwd -l compte          # verrouiller (pas supprimer)
userdel -r ancien         # supprimer + home
id prenom ; groups prenom
```

### 9.1 sudo fin (`/etc/sudoers.d/`)
```
# Toujours via visudo !
deploiement ALL=(ALL) NOPASSWD: /usr/bin/systemctl restart monapp
%admins ALL=(ALL) ALL
```
- Principe du **moindre privilège** : jamais `ALL=(ALL) NOPASSWD: ALL`
  sauf cas justifié.
- Tracer : les commandes sudo sont dans les logs (`/var/log/auth.log`).

---

## 10. SSH — durcissement complet

### 10.1 Clés (fini les mots de passe)
```bash
ssh-keygen -t ed25519 -C "prenom@poste"
ssh-copy-id prenom@srv01
```

### 10.2 `/etc/ssh/sshd_config` — la config qui protège
```
PermitRootLogin no
PasswordAuthentication no
ChallengeResponseAuthentication no
PubkeyAuthentication yes
AllowUsers prenom1 prenom2
Port 22  (ou autre, mais le port ≠ une sécurité)
MaxAuthTries 3
LoginGraceTime 30
ClientAliveInterval 300
ClientAliveCountMax 2
```
```bash
sshd -t && systemctl reload sshd
```
- ⚠️ **Toujours** garder une session ouverte quand on touche à SSH.

