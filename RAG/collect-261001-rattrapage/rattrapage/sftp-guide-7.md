---
id: collect-261001-rattrapage/rattrapage/sftp-guide-7
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [1437, 1618]
sha256: 343eee7ebcff383c95922aeb452e65958a3d7b5b86a9dcf6ea96cfb74783ff43
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

Sans purge, tout dépôt finit par saturer. Définir par périmètre :

| Périmètre | Rétention | Action |
|---|---|---|
| Dépôt clients (`depot/`) | 90 jours après traitement | Archivage puis suppression |
| Échanges bancaires | Durée légale (à valider) | Archivage chiffré |
| Scans copieurs (`entrees/`) | 30 jours après tri | Suppression auto |
| Restitutions | 180 jours | Archivage |

```bash
# Purge des scans triés de plus de 30 jours (avec log)
find /srv/partage/scans_tries -type f -mtime +30 -delete -print | \
  logger -t purge_scans
```

> Toujours logger les purges, et exclure les périmètres à obligation légale.

---

## 70. Cas 11 — horodatage incohérent entre client et serveur

- Tous les serveurs en NTP/chrony : `timedatectl status` → `System clock synchronized: yes`.
- Les logs (`auth.log`, `sftp.log`) utilisent l'heure serveur : la documenter (UTC ou locale ?).
- Dans les rapports, convertir explicitement : « heures Europe/Paris ».
- Un fichier daté du futur bloque les `mirror --only-newer` : purger ou retoucher avec `touch`.

---

## 71. Scripts d'exploitation prêts à l'emploi

### 71.1 Contrôle des permissions des chroots

```bash
#!/bin/bash
# /usr/local/sbin/check_chroot.sh — à lancer après chaque création/modif
BASE=/srv/sftp
OK=1
for d in "$BASE"/*/; do
  own=$(stat -c '%U:%G %a' "$d")
  if [ "$own" != "root:root 755" ]; then
    echo "ANOMALIE: $d -> $own (attendu root:root 755)"
    OK=0
  fi
done
[ $OK -eq 1 ] && echo "Tous les chroots sont conformes."
exit $((1-OK))
```

### 71.2 Rapport quotidien des transferts

```bash
#!/bin/bash
# /usr/local/sbin/rapport_sftp.sh — envoi par mail chaque matin
LOG=/var/log/sftp.log
HIER=$(date -d yesterday '+%b %e')
{
  echo "=== Transferts SFTP du $HIER ==="
  grep "$HIER" "$LOG" | grep -c 'flags WRITE' | xargs echo "Uploads :"
  grep "$HIER" "$LOG" | grep -c 'flags READ'  | xargs echo "Downloads :"
  echo "=== Top fichiers ==="
  grep "$HIER" "$LOG" | grep 'close' | awk '{print $(NF-2), $NF}' | sort -rn | head -20
} | mail -s "[SFTP] Rapport quotidien" admin@entreprise.lan
```

```cron
0 7 * * * root /usr/local/sbin/rapport_sftp.sh
```

---

## 72. Pense-bête de poche (une page)

```
CONNEXION
  sftp -i cle user@hote
  sftp -b batch.txt -o BatchMode=yes user@hote

TRANSFERT
  put / get / mput / mget / reput / reget
  lftp : mirror --reverse --continue --delete src/ /dst/

DIAGNOSTIC
  sftp -v ...                    # côté client
  tail -f /var/log/auth.log      # côté serveur
  sshd -T | grep -i chroot       # config effective
  namei -l /srv/sftp/xxx         # permissions du chroot

PERMISSIONS CHROOT
  chroot            -> root:root 0755 (toute la hiérarchie !)
  dossier de dépôt  -> user:sftpusers 0750

BLOC MATCH TYPE
  Match Group sftpusers
      ChrootDirectory /srv/sftp/%u
      ForceCommand internal-sftp
      PasswordAuthentication no
      AllowTcpForwarding no / X11Forwarding no

URGENCES
  Révoquer une clé : vider /etc/ssh/authorized_keys/<user>
  Bannir une IP   : fail2ban-client set sshd banip <ip>
  Disque plein    : df -h /srv ; du -sh /srv/sftp/* ; purger selon rétention
```

---

## 73. Glossaire détaillé

| Terme | Définition |
|---|---|
| SFTP | SSH File Transfer Protocol : transfert de fichiers dans un tunnel SSH (port 22) |
| FTPS | FTP + TLS : FTP historique sécurisé par TLS (ports 21 + plage passive) |
| SCP | Ancien utilitaire de copie via SSH, supplanté par SFTP |
| Subsystem | Service invoqué par sshd à la demande du client (`sftp`, ici `internal-sftp`) |
| internal-sftp | Implémentation du serveur SFTP intégrée à sshd, sans binaire externe |
| sftp-server | Binaire externe historique (`/usr/lib/openssh/sftp-server`) |
| ChrootDirectory | Directive sshd : dossier qui devient `/` pour l'utilisateur |
| ForceCommand | Force la commande exécutée pour une session (ici `internal-sftp`) |
| Match | Bloc conditionnel de `sshd_config` (User, Group, Address, Host) |
| AuthorizedKeysFile | Fichier(s) contenant les clés publiques autorisées (`%u` = utilisateur) |
| BatchMode | Option client : échoue au lieu d'interagir (indispensable en script/cron) |
| nologin | Shell factice (`/usr/sbin/nologin`) : interdit toute connexion interactive |
| Quota | Limite d'espace (bloc) ou de nombre de fichiers (inode) par utilisateur |
| SyslogFacility | Canal syslog utilisé par sshd pour journaliser (AUTH par défaut) |
| VERBOSE | Niveau de log donnant les open/close de fichiers (qui a transféré quoi) |
| fail2ban | Bannit temporairement les IP après N échecs d'authentification |
| chattr +i | Attribut immuable : fichier insupprimable même par root (sans retirer l'attribut) |
| MFP | Multifunction Printer : copieur multifonction (impression/scan) |
| Rétention | Durée de conservation des fichiers avant archivage/suppression |
| 3-2-1 | Règle de sauvegarde : 3 copies, 2 supports différents, 1 hors site |

---

## 74. Les 12 erreurs classiques (à afficher en salle d'exploitation)

| # | Erreur | Conséquence | Réflexe |
|---|---|---|---|
| 1 | `chown -R user` sur le chroot | Connexion fermée aussitôt | Le chroot reste `root:root 0755` |
| 2 | Oublier `sshd -t` + `reload` | Config non appliquée, panne au reboot | Toujours tester avant de recharger |
| 3 | Fermer sa session SSH avant de tester | Enfermé dehors | Tester dans un 2ᵉ terminal |
| 4 | Clé **privée** déposée côté serveur | Authentification impossible | Seule la **publique** va dans `authorized_keys` |
| 5 | Bloc `Match` pas en fin de fichier | Directives ignorées / erreurs | `Match` toujours après le global |
| 6 | Pas de quota | Un client remplit le disque | Quota dès la création du compte |
| 7 | Mot de passe au lieu de clé « pour dépanner » | Dette de sécurité | Clé dédiée, `Match` ciblé, traçé |
| 8 | `Subsystem sftp` pointant vers l'ancien binaire + chroot | Échec du subsystem | `internal-sftp` partout |
| 9 | Droits `0777` « pour que ça marche » | Fuite de données entre clients | `0750` + bons propriétaires |
| 10 | Pas de `.ok` / contrôle d'intégrité | Fichiers incomplets traités | Protocole fichier témoin + checksum |
| 11 | Logs jamais regardés | Incident découvert trop tard | Rapport quotidien + alertes |
| 12 | Pas de sauvegarde des clés hôtes | Alertes MITM chez tous les clients après restauration | Sauvegarder `/etc/ssh/` |

---

## 75. Cas pratique 1 — onboarding d'un nouveau client (commenté)

**Contexte :** le client « ACME » doit déposer chaque soir ses bons de commande (CSV, ~50 Mo/jour).

```bash
# 1. Création du compte (voir checklist § 51)
U=sftp_acme
sudo useradd -m -d /srv/sftp/client_acme -s /usr/sbin/nologin -G sftpusers \
  -c "Dépôt SFTP client ACME" $U
sudo passwd -l $U

# 2. Arborescence : le client écrit dans depot/, lit dans restitutions/
sudo mkdir -p /srv/sftp/client_acme/{depot,restitutions}
sudo chown root:root /srv/sftp/client_acme && sudo chmod 0755 /srv/sftp/client_acme
sudo chown $U:sftpusers /srv/sftp/client_acme/depot && sudo chmod 0750 /srv/sftp/client_acme/depot
sudo chown root:sftpusers /srv/sftp/client_acme/restitutions && sudo chmod 0750 /srv/sftp/client_acme/restitutions

# 3. Clé publique fournie par ACME (canal authentifié : mail signé, portail, téléphone)
sudo sh -c "cat > /etc/ssh/authorized_keys/$U"   # coller la clé, Ctrl+D
sudo chmod 0644 /etc/ssh/authorized_keys/$U

# 4. Quota : 50 Mo/jour x 90 jours de rétention ≈ 5 Go -> quota 8 Go
sudo setquota -u $U 8388608 8912896 0 0 /srv

# 5. Test croisé : depuis un poste de test, put + get + tentative d'évasion du chroot
# 6. Fiche d'exploitation : contacts ACME, volumétrie, rétention 90 j, rotation clé annuelle
# 7. Communiquer au client : hôte, port, utilisateur, empreinte du serveur (§ 16),
#    arborescence vue de son côté, protocole .ok, format de nommage des fichiers
```

