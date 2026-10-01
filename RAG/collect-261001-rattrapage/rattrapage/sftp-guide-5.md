---
id: collect-261001-rattrapage/rattrapage/sftp-guide-5
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [999, 1229]
sha256: dfd1be02334c0131024bcb47a2cbfbf04ef00edc3463bc78c5fb05ddf60b2c59
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

Permissions : clé `0600` appartenant à l'utilisateur du cron, script `0750`.

---

## 44. Clés dédiées par automatisme (principe)

- **Une paire de clés par flux automatisé**, jamais la clé d'un humain.
- Clé **sans phrase de passe** (sinon le cron bloque), stockée `0600`, propriétaire = compte du cron.
- Restreignez côté serveur si possible (`Match User` dédié + chroot + `AllowUsers user@ip`).
- Documentez : quel flux, quelle machine source, quel compte, date de création, date de rotation prévue.
- En cas de compromission de la machine source : révoquez **immédiatement** la clé (suppression de la ligne dans `authorized_keys`), générez-en une nouvelle.

---

## 45. Python + paramiko (alternative scriptée)

```bash
pip install paramiko
```

```python
import paramiko
cle = paramiko.Ed25519Key.from_private_key_file("/etc/svc_export/id_ed25519")
transport = paramiko.Transport(("sftp.partenaire.lan", 22))
transport.connect(username="sftp_expediteur", pkey=cle)
sftp = paramiko.SFTPClient.from_transport(transport)
sftp.chdir("depot")
sftp.put("/var/exports/rapport.csv", "depot/rapport.csv",
         callback=lambda e, t: print(f"{e}/{t}"))
sftp.close(); transport.close()
```

> paramiko est pratique pour la logique métier (renommage, contrôle, notification),
> mais `sftp -b` / `lftp` suffisent pour du transfert pur.

---

## 46. Transferts fiables : checksum et idempotence

- Après un `put`, vérifiez l'intégrité : comparez un checksum SHA256 calculé en local
  avec celui recalculé après `get` (ou demandez au partenaire de publier un `.sha256`).
- Rendez les dépôts **idempotents** : nommez les fichiers avec la date
  (`factures_2026-09-26.csv`) pour éviter les écrasements silencieux.
- Protocole « fichier témoin » : envoyez d'abord `fichier.csv`, puis `fichier.csv.ok`.
  Le consommateur ne traite que les fichiers ayant leur `.ok` (transfert terminé garanti).

```bash
# côté émetteur
sftp> put factures_2026-09-26.csv /depot/
sftp> put factures_2026-09-26.csv.ok /depot/
# côté consommateur : ignorer tout fichier sans .ok correspondant
```

---

## 47. rsync vs SFTP : que choisir ?

| Critère | rsync-over-SSH | SFTP |
|---|---|---|
| Delta de transfert | Oui (n'envoie que les différences) | Non (fichier entier) |
| Reprise | Oui (`--partial`) | Oui (`reget`/`reput`, manuel) |
| Miroir | Natif (`--delete`) | Via `lftp mirror` |
| Chroot simple | Non (besoin d'un shell + rsync distant) | Oui (`internal-sftp`) |
| Comptes sans shell | Compliqué | **Natif** |
| Partenaires non techniques | Non | Oui (FileZilla/WinSCP) |

**Règle pratique :**

- Interne / technique / gros volumes redondants → **rsync over SSH** (avec un compte shell restreint, `rrsync`).
- Partenaires, clients, copieurs, comptes sans shell → **SFTP**.

---

## 48. rsync-over-SSH en pratique (pour mémoire)

```bash
# Envoi avec reprise, compression, suppression distante
rsync -avz --partial --progress -e "ssh -i ~/.ssh/id_rsync -p 22" \
  /var/exports/ svc_rsync@srv.entreprise.lan:/srv/depots/
```

> rsync exige un shell et le binaire rsync côté distant : **incompatible** avec
> `ForceCommand internal-sftp`. Pour des partenaires en chroot, restez sur SFTP.

---

## 49. Clients graphiques : FileZilla

Configuration d'un site (Gestionnaire de sites) :

| Champ | Valeur |
|---|---|
| Protocole | SFTP - SSH File Transfer Protocol |
| Hôte | `sftp.entreprise.lan` |
| Port | 22 |
| Type d'authentification | Fichier de clé |
| Utilisateur | `sftp_acme` |
| Fichier de clé | `id_sftp_acme.ppk` (converti via PuTTYgen si besoin) ou PEM direct |

Réglages recommandés : Édition → Paramètres → Transferts → limiter les transferts
simultanés à 2–4, activer la reprise.

---

## 50. Clients graphiques : WinSCP

- Nouveau site → Protocole **SFTP**, port 22, authentification par clé
  (`.ppk` via PuTTYgen, ou clé OpenSSH récente acceptée directement).
- Avancé → SSH → Authentification : cocher « Tentative d'authentification par clé »,
  désactiver « Tentative d'authentification par mot de passe » si clés uniquement.
- Scripting WinSCP pour automatisation Windows :

```bat
winscp.com /command ^
  "open sftp://sftp_acme@sftp.entreprise.lan/ -privatekey=C:\cles\id_sftp_acme.ppk -hostkey=""ssh-ed25519 256 AA:BB:...""" ^
  "put C:\exports\*.csv /depot/" ^
  "exit"
```

> `-hostkey` fige l'empreinte du serveur : le script échoue si elle change (anti-MITM).

---

## 51. Fiche réflexe : créer un compte SFTP complet (checklist)

```bash
U=sftp_nouveauclient ; PERIM=client_nouveau
sudo useradd -m -d /srv/sftp/$PERIM -s /usr/sbin/nologin -G sftpusers -c "Dépôt SFTP $PERIM" $U
sudo passwd -l $U
sudo mkdir -p /srv/sftp/$PERIM/{depot,restitutions}
sudo chown root:root /srv/sftp/$PERIM && sudo chmod 0755 /srv/sftp/$PERIM
sudo chown $U:sftpusers /srv/sftp/$PERIM/depot && sudo chmod 0750 /srv/sftp/$PERIM/depot
sudo chown root:$U /srv/sftp/$PERIM/restitutions 2>/dev/null || sudo chown root:sftpusers /srv/sftp/$PERIM/restitutions
sudo chmod 0750 /srv/sftp/$PERIM/restitutions
# clé publique du partenaire :
sudo sh -c "cat > /etc/ssh/authorized_keys/$U"   # coller, Ctrl+D
sudo chmod 0644 /etc/ssh/authorized_keys/$U
# bloc Match si chroot != /srv/sftp/$U (voir § 11), puis :
sudo sshd -t && sudo systemctl reload ssh
# quota, test de connexion, fiche d'exploitation
```

---

## 52. Supprimer / désactiver proprement un compte

```bash
U=sftp_ancienclient
# 1. Révoquer l'accès immédiatement (sans supprimer les données)
sudo sh -c "echo '# révoqué le $(date +%F)' > /etc/ssh/authorized_keys/$U"
# 2. Vérifier qu'il n'y a plus de session
sudo ss -tnp | grep :22
# 3. Archiver les données (selon contrat / obligations légales)
sudo tar -czf /srv/archives/${U}_$(date +%F).tar.gz /srv/sftp/<perimetre>
# 4. Après le délai de rétention : supprimer
sudo userdel -r $U   # -r supprime aussi le home si c'est le chroot : ATTENTION
sudo rm -rf /srv/sftp/<perimetre>
```

> Ne supprimez jamais un chroot sans archive préalable et validation du responsable.

---

## 53. Sécurité : durcissement du serveur hôte

- Mises à jour automatiques de sécurité (`unattended-upgrades`) — OpenSSH est critique.
- Partition `/srv` dédiée (un dépôt plein ne bloque pas le système).
- Pas d'autres services exposés sur la machine SFTP (cloisonnement).
- `PermitRootLogin no`, comptes admin en clés uniquement.
- Audit périodique : `lynis`, ou au minimum relecture annuelle de `sshd_config`.
- Sauvegarde chiffrée hors site (voir § 27).

---

## 54. Sécurité : journaliser aussi les échecs

Les échecs sont dans `/var/log/auth.log` :

```bash
# Top 10 des IP en échec aujourd'hui
sudo grep "$(date '+%b %e')" /var/log/auth.log | grep -i "failed" | \
  grep -oE '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+' | sort | uniq -c | sort -rn | head
# Utilisateurs inexistants tentés
sudo grep "Invalid user" /var/log/auth.log | awk '{print $(NF-2)}' | sort | uniq -c | sort -rn | head
```

Toute tentative sur un compte SFTP existant avec mot de passe doit être **impossible**
(`PasswordAuthentication no`) : si vous en voyez réussir, c'est une urgence.

---

## 55. fail2ban : aller plus loin (filtres SFTP)

Le jail `sshd` couvre déjà les échecs d'authentification. Ajoutez une surveillance
des scans agressifs :

```ini
# /etc/fail2ban/jail.d/sshd.local — version renforcée
[sshd]
enabled = true
maxretry = 5
findtime = 600
bantime = 86400
```

Et un jail récidiviste :

```ini
[recidive]
enabled = true
filter = recidive
logpath = /var/log/fail2ban.log
banaction = %(banaction_allports)s
bantime = 604800
findtime = 86400
maxretry = 3
```

---

## 56. TLS/FTPS : quand le partenaire impose le FTP(S)

Parfois un partenaire (banque, administration) n'accepte que FTPS. Dans ce cas,
**vsftpd** en mode minimal et sécurisé, sur une machine dédiée si possible.

```bash
sudo apt install -y vsftpd
```

`/etc/vsftpd.conf` (extrait sécurisé) :

