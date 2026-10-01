---
id: collect-261001-rattrapage/rattrapage/sftp-guide-3
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [462, 696]
sha256: 2fe4633fb246fcc42e8083ff644bc09ebb95596cf57671c9c660963eeca82c1c
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

Sur le copieur (ex. panneau Kyocera/équivalent) : protocole SFTP, port 22,
authentification par clé (coller la privée générée pour ce compte) ou login/mot de passe
**dédié** si le modèle ne gère pas les clés — dans ce cas, `Match User` avec
`PasswordAuthentication yes` ciblé (voir § 58).

Organisation conseillée : un sous-dossier par machine dans `entrees/` pour trier ensuite :

```
/srv/sftp/scans_copieurs/entrees/
├── etage1_accueil/
├── etage2_compta/
└── atelier/
```

---

## 21. Réception et tri automatisé des scans

Script de tri (toutes les 5 min via cron, ou déclenché par inotify) :

```bash
#!/bin/bash
# /usr/local/sbin/tri_scans.sh — classe les PDF par machine et par date
SRC=/srv/sftp/scans_copieurs/entrees
DST=/srv/partage/scans_tries
for machine in "$SRC"/*/; do
  m=$(basename "$machine")
  for f in "$machine"*.pdf "$machine"*.PDF; do
    [ -e "$f" ] || continue
    jour=$(date +%F)
    dest="$DST/$m/$jour"
    mkdir -p "$dest"
    mv "$f" "$dest/"
    logger -t tri_scans "classé $f -> $dest/"
  done
done
```

```cron
*/5 * * * * root /usr/local/sbin/tri_scans.sh
```

> Le script tourne en root (ou avec ACL sur `entrees/`) : le compte `svc_copieur`
> n'a accès qu'en écriture à son dépôt, il ne peut ni lire ni supprimer les scans des autres machines.

---

## 22. Échanges comptables et bancaires

Exigences typiques : confidentialité, traçabilité, horodatage, pas de suppression.

Arborescence :

```
/srv/sftp/compta/
├── envoi_banque/      # la compta dépose les fichiers de virement (écriture compta)
├── recus_banque/      # la banque dépose les relevés (écriture banque, lecture compta)
└── archives/          # écriture exploitant seul (gel des échanges)
```

Durcissement :

- Deux comptes distincts : `svc_compta` (interne) et `sftp_banque` (partenaire).
- Permissions croisées : chacun écrit dans son dossier, lit (sans écrire) dans celui de l'autre.
- Interdire la suppression : pas de droit `w` sur le dossier parent pour les comptes SFTP,
  ou utiliser l'attribut immuable sur les fichiers archivés (`chattr +i`).

```bash
sudo chown svc_compta:sftp_compta /srv/sftp/compta/envoi_banque
sudo chmod 0750 /srv/sftp/compta/envoi_banque
sudo chown sftp_banque:sftp_compta /srv/sftp/compta/recus_banque
sudo chmod 0750 /srv/sftp/compta/recus_banque
```

---

## 23. Dépôt de fichiers clients (espace d'échange)

Modèle « boîte aux lettres » par client :

- `depot/` : le client **écrit**, ne **relit pas** (0750, propriétaire client) — évite qu'un client
  voie les fichiers d'un autre et limite l'exfiltration.
- `restitutions/` : l'exploitant **écrit**, le client **lit** (0750, propriétaire root, groupe du client).

Checklist de création d'un espace client :

- [ ] Utilisateur `sftp_<client>` créé, shell nologin, mot de passe verrouillé
- [ ] Membre de `sftpusers`
- [ ] Chroot `/srv/sftp/<perimetre>` root:root 0755
- [ ] Sous-dossiers avec bons propriétaires/permissions
- [ ] Clé publique installée dans `/etc/ssh/authorized_keys/<user>`
- [ ] Test de connexion + `put`/`get` depuis un poste de test
- [ ] Quota appliqué (voir § 24)
- [ ] Fiche d'exploitation renseignée (contacts, volumétrie, rétention)

---

## 24. Quotas disque (project quota XFS ou quota ext4)

Sans quota, un client peut remplir le disque et bloquer tout le serveur.

**Option ext4 (classique)** :

```bash
sudo apt install -y quota quotatool
# activer dans /etc/fstab : ajouter usrquota,grpquota aux options de /srv
# UUID=... /srv ext4 defaults,usrquota,grpquota 0 2
sudo mount -o remount /srv
sudo quotacheck -cum /srv && sudo quotaon /srv
sudo setquota -u sftp_acme 10485760 11534336 0 0 /srv   # 10 Go soft, 11 Go hard
sudo repquota /srv
```

**Option XFS (project quotas, conseillée si /srv est en XFS)** :

```bash
# /etc/fstab : ...,pquota
sudo xfs_quota -x -c 'project -s -p /srv/sftp/client_acme 101' /srv
sudo xfs_quota -x -c 'limit -p bhard=10g 101' /srv
```

Surveillance : alerte si usage > 80 % (voir § 52).

---

## 25. Limiter la bande passante par utilisateur (optionnel)

`sshd_config` ne gère pas nativement la limitation de débit SFTP. Solutions :

- **trickle** côté client pour les envois planifiés.
- **QoS réseau** (tc) sur le serveur si un flux écrase les autres.
- En pratique : les quotas + fenêtres de transfert planifiées (cron) suffisent le plus souvent.

---

## 26. Chiffrement au repos (si données sensibles)

Le chiffrement SSH ne protège que le **transport**. Pour les données au repos :

| Solution | Cas d'usage |
|---|---|
| LUKS sur le volume `/srv` | Vol/perte du serveur, mise au rebut des disques |
| eCryptfs / fscrypt | Chiffrement par dossier (plus complexe à opérer) |
| Chiffrement applicatif (GPG) par le partenaire | Secret de bout en bout, le serveur ne voit que du chiffré |

Exemple GPG côté partenaire avant dépôt :

```bash
gpg --encrypt --recipient compta@entreprise.lan virements.csv
# déposer virements.csv.gpg sur le SFTP
```

---

## 27. Sauvegarde des dépôts SFTP

Ce qui doit être sauvegardé :

- [ ] Le contenu des chroots (`/srv/sftp`) — données métier
- [ ] `/etc/ssh/` (clés hôtes + `authorized_keys/`) — sans les clés hôtes, tous les clients verront une alerte MITM après restauration
- [ ] `/etc/ssh/sshd_config*` et fragments
- [ ] Scripts d'exploitation (`/usr/local/sbin/tri_scans.sh`, etc.)
- [ ] Journaux (`/var/log/auth.log`, syslog) selon la durée légale de conservation

Exemple avec Borg (voir aussi le guide Debian/Ubuntu) :

```bash
export BORG_REPO=/mnt/sauvegardes/borg-sftp
borg create --stats $BORG_REPO::'{hostname}-{now:%F}' \
  /srv/sftp /etc/ssh /usr/local/sbin/tri_scans.sh
borg prune --keep-daily=7 --keep-weekly=4 --keep-monthly=6 $BORG_REPO
```

Règle 3-2-1 : 3 copies, 2 supports, 1 hors site. Testez la restauration **au moins** une fois par trimestre.

---

## 28. Restauration après sinistre

Procédure type :

1. Réinstaller Debian/Ubuntu + `openssh-server`.
2. Restaurer `/etc/ssh/` (clés hôtes identiques → pas d'alerte côté clients).
3. Restaurer `sshd_config` et vérifier avec `sshd -t`.
4. Recréer utilisateurs/groupes avec **les mêmes UID/GID** (sinon les propriétaires des fichiers restaurés seront faux).
5. Restaurer `/srv/sftp` avec les permissions.
6. Rejouer le script de contrôle des permissions (voir § 71).
7. Test de connexion avec un compte de test avant réouverture aux partenaires.

> Astuce : sauvegardez `/etc/passwd`, `/etc/group`, `/etc/shadow` (ou au minimum la liste
> des UID/GID SFTP) pour recréer les comptes à l'identique.

---

## 29. Journalisation de base : auth.log

Par défaut, OpenSSH journalise via syslog (facility AUTH) :

```bash
# Connexions réussies / échouées
sudo grep -i "sshd" /var/log/auth.log | tail -20
# Échecs d'authentification
sudo grep "Failed password" /var/log/auth.log
sudo grep "Accepted publickey" /var/log/auth.log
```

Exemple de ligne utile :

```
sshd[1234]: Accepted publickey for sftp_acme from 203.0.113.10 port 51234 ssh2: ED25519 SHA256:...
```

On y voit : qui (`sftp_acme`), d'où (`203.0.113.10`), avec quelle clé (empreinte SHA256).

---

## 30. Journalisation détaillée des transferts (VERBOSE)

Pour savoir **qui a transféré quoi**, augmentez le niveau de log du subsystem :

```
Subsystem sftp internal-sftp -f AUTH -l VERBOSE
```

Puis `sudo systemctl reload ssh`. Vous obtiendrez dans `/var/log/auth.log` :

```
sftp-server[1234]: open "/depot/rapport.pdf" flags WRITE,CREATE,TRUNCATE mode 0644
sftp-server[1234]: close "/depot/rapport.pdf" bytes read 0 written 182340
```

- `open ... WRITE` = début d'envoi (upload) ; `READ` = début de téléchargement.
- `close ... bytes read/written` = fin de transfert avec les volumes.
- Corrélé avec la ligne `Accepted publickey for <user>`, on reconstitue : utilisateur, IP, fichier, sens, volume, heure.

