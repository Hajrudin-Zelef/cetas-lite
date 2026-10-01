---
id: collect-261001-rattrapage/rattrapage/sftp-guide-2
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [216, 461]
sha256: 2306de574c288a68e4c794afb60f696813f882758251bc8c7abde2eb00585f3d
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

> ⚠️ `%u` = nom d'utilisateur. Ici le chroot est `/srv/sftp/sftp_acme` — adaptez la convention :
> soit nommez les dossiers comme les utilisateurs, soit utilisez un `Match User` par compte
> avec un `ChrootDirectory` explicite (voir § 11).

Après modification :

```bash
sudo sshd -t && sudo systemctl reload ssh
```

Ordre de lecture : les blocs `Match` doivent figurer **à la fin** de `sshd_config`,
après toutes les directives globales.

---

## 11. Chroot par utilisateur (quand %u ne suffit pas)

Si le dossier chroot ne porte pas le nom de l'utilisateur :

```
Match User sftp_acme
    ChrootDirectory /srv/sftp/client_acme
    ForceCommand internal-sftp
    PasswordAuthentication no

Match User svc_copieur_etage2
    ChrootDirectory /srv/sftp/scans_copieurs
    ForceCommand internal-sftp
    PasswordAuthentication no
```

Et par groupe pour un périmètre commun :

```
Match Group sftp_compta
    ChrootDirectory /srv/sftp/compta
    ForceCommand internal-sftp
```

Un utilisateur peut appartenir à plusieurs groupes : le **premier** bloc `Match` correspondant gagne.

---

## 12. Clés SSH uniquement : génération côté client

```bash
# Sur le poste client / le serveur d'envoi
ssh-keygen -t ed25519 -f ~/.ssh/id_sftp_acme -C "sftp_acme@client"
# Protégez la clé privée par une phrase de passe pour un usage interactif ;
# pour un batch automatisé, clé dédiée SANS phrase de passe, restreinte (voir § 44).
```

Déposer la clé publique côté serveur (l'utilisateur n'a pas de home inscriptible) :

```bash
sudo mkdir -p /etc/ssh/authorized_keys
sudo sh -c 'cat > /etc/ssh/authorized_keys/sftp_acme'   # coller la clé publique, Ctrl+D
sudo chown root:root /etc/ssh/authorized_keys/sftp_acme
sudo chmod 0644 /etc/ssh/authorized_keys/sftp_acme
```

Dans `sshd_config` (global ou dans le bloc Match) :

```
AuthorizedKeysFile /etc/ssh/authorized_keys/%u
```

Alternative : `AuthorizedKeysFile .ssh/authorized_keys` avec un home **hors chroot**.

---

## 13. Tester la première connexion

```bash
# Depuis le client possédant la clé privée
sftp -i ~/.ssh/id_sftp_acme sftp_acme@sftp.entreprise.lan
```

Session type :

```
sftp> pwd
Remote working directory: /
sftp> ls -la
drwxr-xr-x  4 root     root     4096 ... .
drwxr-xr-x  4 root     root     4096 ... ..
drwxr-x---  2 sftp_acme sftpusers 4096 ... depot
drwxr-x---  2 root     sftpusers 4096 ... restitutions
sftp> cd depot
sftp> put rapport.pdf
Uploading rapport.pdf to /depot/rapport.pdf
sftp> ls -l
-rw-r--r-- 1 sftp_acme sftpusers 182340 ... rapport.pdf
sftp> bye
```

Points à valider : pas de `!` (shell) disponible, `cd /` reste dans le chroot,
impossible de remonter au-dessus.

---

## 14. Commandes SFTP interactives essentielles

| Commande | Effet |
|---|---|
| `pwd` / `lpwd` | Dossier distant / local courant |
| `ls [-la]` / `lls` | Lister distant / local |
| `cd` / `lcd` | Changer de dossier distant / local |
| `get fichier` / `put fichier` | Télécharger / envoyer un fichier |
| `get -r dossier` / `put -r dossier` | Récursif |
| `mkdir` / `rmdir` / `rm` | Créer/supprimer |
| `rename a b` | Renommer/déplacer |
| `chmod` / `chown` / `chgrp` | Permissions (si autorisées) |
| `df -h` | Espace disque visible du chroot |
| `progress` | Barre de progression |
| `reget` / `reput` | Reprise d'un transfert interrompu |
| `bye` / `exit` / `quit` | Quitter |

Options utiles de la commande `sftp` :

| Option | Effet |
|---|---|
| `-i clé` | Clé privée à utiliser |
| `-P port` | Port non standard |
| `-b fichier` | Batch : exécute les commandes du fichier |
| `-B taille` | Taille des buffers (gros fichiers) |
| `-R nbre` | Nombre de requêtes en vol (débit) |
| `-o Opt=Val` | Passe une option ssh_config |

---

## 15. Reprise de transfert (reget / reput)

```bash
sftp> reput gros_fichier.iso
# reprend là où le transfert s'était arrêté (ajoute à la fin du fichier partiel)
sftp> reget archive.zip
```

- `reput`/`reget` comparent la taille du fichier partiel distant/local et reprennent l'envoi.
- Indispensable sur liens instables ou pour de gros fichiers (ISO, sauvegardes, scans groupés).
- Avec `lftp`, la reprise est automatique (`mirror --continue`, voir § 42).

---

## 16. Forcer le chiffrement et vérifier l'empreinte du serveur

Premier contact : vérifiez l'empreinte (fingerprint) hors bande :

```bash
# Côté serveur : empreintes à communiquer au partenaire
sudo ssh-keygen -lf /etc/ssh/ssh_host_ed25519_key.pub
sudo ssh-keygen -lf /etc/ssh/ssh_host_rsa_key.pub
# Format visuel "randomart" pour lecture au téléphone :
sudo ssh-keygen -lvf /etc/ssh/ssh_host_ed25519_key.pub
```

Côté client, au premier `sftp`, comparez l'empreinte affichée avec celle fournie
par l'exploitant (jamais « oui » aveugle en production).

---

## 17. Durcir sshd_config (global)

```
# --- Durcissement global ---
Port 22
Protocol 2
PermitRootLogin no
PasswordAuthentication no            # clés uniquement pour tout le monde
PubkeyAuthentication yes
ChallengeResponseAuthentication no
UsePAM yes
X11Forwarding no
AllowAgentForwarding no
AllowTcpForwarding no
PermitTunnel no
MaxAuthTries 3
LoginGraceTime 30
ClientAliveInterval 300
ClientAliveCountMax 2
MaxSessions 5
```

> Si certains comptes admin ont encore besoin du mot de passe, ne rouvrez pas globalement :
> utilisez un `Match User adm_xxx` ciblé, ou mieux, déployez-leur des clés.

---

## 18. Où stocker les clés autorisées ?

Trois stratégies :

| Stratégie | AuthorizedKeysFile | Avantages | Inconvénients |
|---|---|---|---|
| A. Centralisé (recommandé) | `/etc/ssh/authorized_keys/%u` | Hors chroot, géré par root, audit simple | Un fichier par utilisateur à créer |
| B. Home hors chroot | `/home/%u/.ssh/authorized_keys` avec home hors chroot | Classique | Le home doit exister hors chroot |
| C. Dans le chroot | `%h/.ssh/authorized_keys` | — | **Déconseillé** : le chroot appartient à root, l'utilisateur ne peut pas gérer ses clés, et sshd exige des permissions strictes |

Mise en place de la stratégie A :

```bash
sudo install -d -o root -g root -m 0755 /etc/ssh/authorized_keys
# un fichier par compte, 0644 root:root (voir § 12)
```

---

## 19. Rotation des clés

Procédure :

1. Le partenaire génère une nouvelle paire et transmet la **publique** par canal authentifié.
2. Ajoutez-la en **plus** de l'ancienne dans `/etc/ssh/authorized_keys/<user>`.
3. Le partenaire teste la nouvelle clé.
4. Retirez l'ancienne clé du fichier.
5. Tracez l'opération (qui, quand, pourquoi) dans votre journal d'exploitation.

```bash
# Ajouter sans écraser :
sudo sh -c 'cat >> /etc/ssh/authorized_keys/sftp_acme'   # coller, Ctrl+D
sudo chmod 0644 /etc/ssh/authorized_keys/sftp_acme
# Vérifier qu'il y a bien 2 lignes, tester, puis supprimer l'ancienne :
sudo sed -i '/ancienne-cle-commentaire/d' /etc/ssh/authorized_keys/sftp_acme
```

Fréquence conseillée : annuelle, ou à chaque départ d'un interlocuteur / compromission suspectée.

---

## 20. Comptes techniques : copieurs multifonctions (MFP)

Cas métier : les copieurs/IMFP du parc déposent leurs numérisations sur le SFTP
(lien avec la maintenance des copieurs : le scan-to-SFTP remplace le scan-to-SMB,
plus sûr et sans compte AD à gérer sur chaque machine).

```bash
sudo useradd -r -d /srv/sftp/scans_copieurs -s /usr/sbin/nologin \
  -G sftpusers -c "Dépôt scans copieurs" svc_copieur
sudo passwd -l svc_copieur
sudo mkdir -p /srv/sftp/scans_copieurs/entrees
sudo chown root:root /srv/sftp/scans_copieurs && sudo chmod 0755 /srv/sftp/scans_copieurs
sudo chown svc_copieur:sftpusers /srv/sftp/scans_copieurs/entrees
sudo chmod 0730 /srv/sftp/scans_copieurs/entrees   # écriture seule + lecture du nom
```

