---
id: collect-261001-rattrapage/rattrapage/debian-ubuntu-guide-5
title: "Debian & Ubuntu — Guide ultra-complet d'administration système"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agents"]
source: docs/RAG/collect-261001-rattrapage/debian_ubuntu_guide.md
source_anchor: ""
source_lines: [1099, 1333]
sha256: 018dc23b66fc961e3c3faa2a117758925be69373b8fd961b4e2901ffb1c43ac2
---

# Debian & Ubuntu — Guide ultra-complet d'administration système

### 54.1 keepalived — IP flottante
```
2 serveurs (ex. reverse proxy) partagent une VIP.
Le maître tombe → le backup prend la VIP en ~3 s.
```
```bash
apt install keepalived
# /etc/keepalived/keepalived.conf : vrrp_instance, virtual_ipaddress
```

### 54.2 Bases de données
- PostgreSQL : réplication streaming (primaire → standby) +
  Patroni pour le failover auto.
- MariaDB : Galera (multi-maître) ou réplication classique.

### 54.3 Fichiers partagés
- DRBD (miroir bloc réseau) ou NFS redondé, ou Ceph/GlusterFS.

---

## 55. Reverse proxy avancé

- **nginx** : suffisant dans 90 % des cas (§31).
- **Traefik** : découverte auto des conteneurs Docker, TLS auto
  (Let's Encrypt), tableau de bord.
- **HAProxy** : quand il faut du TCP pur / de la haute perf.

---

## 56. Autorité de certification interne

```bash
# CA maison pour les services internes (intranet, imprimantes...)
openssl req -x509 -newkey rsa:4096 -keyout ca.key -out ca.crt -days 3650
# Signer les CSR des serveurs, déployer ca.crt sur les postes
```
- Évite les erreurs TLS internes et les certificats autosignés
  sauvages. Durée : 5–10 ans, **procédure de renouvellement écrite**.

---

## 57. Sécuriser un serveur web

```
□ TLS 1.2+ uniquement, HSTS
□ Headers : X-Frame-Options, X-Content-Type-Options, CSP
□ ServerTokens Prod / server_tokens off
□ WAF : ModSecurity (nginx) pour les applis exposées
□ Rate limiting (nginx limit_req) anti-scraping/brute-force
□ Séparer les pools PHP par site (un site compromis ≠ tous)
□ Mises à jour auto des paquets security
```

---

## 58. Chiffrement disque — LUKS

```bash
cryptsetup luksFormat /dev/sdb
cryptsetup open /dev/sdb data_crypt
mkfs.ext4 /dev/mapper/data_crypt
# /etc/crypttab : data_crypt UUID=xxx none luks
# /etc/fstab : /dev/mapper/data_crypt /data ext4 defaults 0 2
```
- Indispensable : portables, serveurs en site non sécurisé,
  données sensibles.
- ⚠️ Passphrase au boot = pas de reboot à distance sans console
  (ou Tang/Clevis pour le déverrouillage réseau).

---

## 59. 2FA sur SSH

```bash
apt install libpam-google-authenticator
google-authenticator   # par utilisateur
# /etc/pam.d/sshd : auth required pam_google_authenticator.so
# sshd_config : ChallengeResponseAuthentication yes
#              AuthenticationMethods publickey,keyboard-interactive
```
- Clé + code TOTP : même avec la clé volée, pas d'accès.

---

## 60. Bastion SSH — l'unique porte d'entrée

```
Internet → [BASTION] → réseau serveurs (pas d'SSH direct)
```
- Le bastion : minimal, à jour, 2FA, fail2ban agressif, logs
  centralisés.
- `ProxyJump` côté client :
```
Host srv-interne
  ProxyJump admin@bastion
```
- Sessions enregistrées (tlog) pour les environnements sensibles.

---

## 61. Conteneurs en production — bonnes pratiques

```
□ 1 processus par conteneur
□ Images officielles/slimes, tag fixé (jamais :latest en prod)
□ Données dans des volumes nommés (sauvegardés !)
□ restart: unless-stopped
□ Réseaux dédiés par appli
□ Secrets : fichiers ou vault, jamais en clair dans compose
□ Mises à jour : watchtower (avec fenêtre) ou pipeline CI/CD
□ Logs : driver json-file + rotation, ou centralisés (Loki)
```

---

## 62. Bareos — sauvegarde d'entreprise

- Serveur central (Director), agents (File Daemon) sur chaque
  serveur, stockage (fichiers/disque/bande).
- Politiques : Full hebdo + Diff quotidien + Incr, rétention
  configurée, catalogue en base.
- À partir de ~20 serveurs, Bareos/Bacula bat les scripts.

---

## 63. Gestion des secrets

- **Jamais** de mot de passe en clair dans les scripts ou Git.
- Solutions : Ansible Vault (chiffré), HashiCorp Vault (central),
  gestionnaire d'équipe (Bitwarden/Vaultwarden auto-hébergé).
```bash
ansible-vault create secrets.yml
ansible-vault edit secrets.yml
```

---

## 64. Inventaire et CMDB — savoir ce qu'on a

| Champ | Exemple |
|---|---|
| Nom / rôle | srv01 / fichiers+impression |
| OS | Ubuntu 24.04 LTS |
| IP / MAC | 192.168.1.10 |
| Garantie | jusqu'au ___ |
| Contrat | Silver, échéance ___ |
| Sauvegarde | Borg 22h, testée le ___ |

- Outils : NetBox (infra), GLPI (parc + tickets) — **GLPI est
  aussi votre ticketing SAV**, parfait pour une équipe
  systèmes/énergies.

---

## 65. GLPI — tickets et parc (recommandé)

- Pourquoi : chaque panne devient un ticket traçable, chaque
  équipement a sa fiche, les SLA se mesurent.
- Installation : LAMP + GLPI (dépôt ou archive officielle),
  sauvegarde de la base **quotidienne**.
- Usage chef de service : tableau de bord (tickets ouverts,
  temps de résolution), base de connaissances (vos guides !),
  gestion des contrats fournisseurs.
---

## 66. Tableau récap — Debian vs Ubuntu en pratique

| Sujet | Debian | Ubuntu |
|---|---|---|
| Config réseau | `/etc/network/interfaces` | Netplan (`/etc/netplan/`) |
| Pare-feu préinstallé | nftables (à configurer) | UFW disponible |
| NTP | systemd-timesyncd (chrony en option) | chrony (serveur) |
| Mises à jour auto | unattended-upgrades (à installer) | préinstallé sur serveur |
| Montée de version | édition sources.list | `do-release-upgrade` |
| Paquets universels | flatpak (manuel) | snap (intégré) |
| Déploiement auto | preseed | cloud-init / autoinstall |
| Support commercial | communautaire | Canonical (Ubuntu Pro/ESM) |
| Idéal | stabilité max, Proxmox | parc entreprise, cloud |

---

## 67. Pense-bête de poche

```
apt update && apt full-upgrade | systemctl --failed | journalctl -p err -b
df -h && df -i | sshd -t && nginx -t | ufw status | fail2ban-client status
timedatectl | chronyc tracking | ss -tlnp | ip -brief addr
vgdisplay -s / lvs | cat /proc/mdstat | borg list /backup/repo
NE JAMAIS : reboot sans console à distance, :latest en prod,
            mdp en clair, rm -rf sans vérifier, upgrade sans backup.
```

---

## 68. Questions pièges (validez vos connaissances)

1. Pourquoi séparer `/var` ? → Les logs ne doivent pas remplir `/`.
2. `apt upgrade` vs `full-upgrade` ? → full-upgrade gère les
   changements de dépendances (nouveaux paquets, suppressions).
3. Le RAID remplace-t-il la sauvegarde ? → Non (fausse manip,
   ransomware, incendie).
4. Pourquoi `netplan try` ? → Retour auto après 120 s si on perd
   le réseau.
5. SSH : pourquoi garder une session ouverte ? → Pour réparer si
   la nouvelle config bloque.
6. Swap sur un serveur avec 64 Go de RAM ? → Petit swap/zram quand
   même (le noyau en a besoin).
7. `rm` sur un fichier ouvert par un processus ? → L'espace n'est
   libéré qu'à la fermeture (`lsof | grep deleted`).
8. Pourquoi les logs centralisés ? → L'attaquant efface les logs
   locaux.
9. Charge 8 sur 4 cœurs : grave ? → Oui, investiguer (2× les cœurs).
10. Quand tester une sauvegarde ? → 2×/an minimum, et après chaque
    changement majeur.

---

## 69. Erreurs classiques des débutants

1. **Installer avec DHCP** puis perdre l'IP → toujours IP fixe.
2. **Oublier unattended-upgrades** → serveur non patché 2 ans.
3. **Root en SSH avec mot de passe** → compromis en quelques jours.
4. **Pas de sauvegarde testée** → découverte le jour de la panne.
5. **`/var/log` qui explose** → logrotate à vérifier.
6. **Firewall qui bloque tout après `ufw enable` en SSH** → règle
   22 **avant** d'activer.
7. **Kernel hold oublié** → reboot sur un kernel qui ne boot pas.
8. **Permissions 777** « pour que ça marche » → revoir les groupes.
9. **Docker en root + socket exposé** → faille critique.
10. **Ne rien documenter** → le savoir part avec le technicien.

---

## 70. Prochaines étapes possibles

