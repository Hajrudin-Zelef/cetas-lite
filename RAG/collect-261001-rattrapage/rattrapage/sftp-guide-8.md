---
id: collect-261001-rattrapage/rattrapage/sftp-guide-8
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [1619, 1772]
sha256: d4c393a54b10657b504a66609764d86ea8f7b70e4a1b53bd177c7b78bfc08f26
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

**Points de vigilance :** prévenir ACME que `restitutions/` est en lecture seule ;
planifier la purge à 90 jours (§ 69) ; ajouter le rapport quotidien (§ 71.2).

---

## 76. Cas pratique 2 — échange bancaire sécurisé (commenté)

**Contexte :** la comptabilité envoie les virements et récupère les relevés ; la banque impose un SFTP (tant mieux).

```bash
# 1. Deux comptes distincts : un interne, un partenaire
sudo useradd -m -d /srv/sftp/compta -s /usr/sbin/nologin -G sftp_compta \
  -c "Comptabilité interne" svc_compta
sudo useradd -m -d /srv/sftp/compta -s /usr/sbin/nologin -G sftp_compta \
  -c "Partenaire bancaire" sftp_banque
sudo passwd -l svc_compta sftp_banque

# 2. Chroot commun au groupe (voir § 11)
# Match Group sftp_compta
#     ChrootDirectory /srv/sftp/compta
#     ForceCommand internal-sftp

# 3. Permissions croisées : chacun écrit chez soi, lit chez l'autre
sudo mkdir -p /srv/sftp/compta/{envoi_banque,recus_banque,archives}
sudo chown root:root /srv/sftp/compta && sudo chmod 0755 /srv/sftp/compta
sudo chown svc_compta:sftp_compta /srv/sftp/compta/envoi_banque && sudo chmod 0750 /srv/sftp/compta/envoi_banque
sudo chown sftp_banque:sftp_compta /srv/sftp/compta/recus_banque && sudo chmod 0750 /srv/sftp/compta/recus_banque
sudo chown root:root /srv/sftp/compta/archives && sudo chmod 0750 /srv/sftp/compta/archives
```

**Exploitation :**

- Les fichiers de virement sont chiffrés GPG avant dépôt (§ 26).
- Un script déplace chaque nuit les fichiers traités vers `archives/` puis `chattr +i`.
- Rétention selon les obligations légales/comptables (à valider avec la DAF).
- Restriction IP : `AllowUsers sftp_banque@<ip_banque>` si l'IP est fixe (§ 37).

---

## 77. Cas pratique 3 — réception des scans des copieurs (commenté)

**Contexte :** 6 copieurs déposent leurs numérisations ; le service veut un tri automatique par machine et par jour (voir § 20-21).

```bash
# 1. Compte technique unique (ou un par site si besoin de cloisonner)
sudo useradd -r -d /srv/sftp/scans_copieurs -s /usr/sbin/nologin -G sftpusers \
  -c "Dépôt scans copieurs" svc_copieur
sudo passwd -l svc_copieur

# 2. Dépôt en écriture seule : le copieur ne doit ni relire ni lister les autres scans
sudo mkdir -p /srv/sftp/scans_copieurs/entrees/{etage1,etage2,atelier}
sudo chown root:root /srv/sftp/scans_copieurs && sudo chmod 0755 /srv/sftp/scans_copieurs
sudo chown svc_copieur:sftpusers /srv/sftp/scans_copieurs/entrees
sudo chmod 0730 /srv/sftp/scans_copieurs/entrees
sudo chown svc_copieur:sftpusers /srv/sftp/scans_copieurs/entrees/*
sudo chmod 0730 /srv/sftp/scans_copieurs/entrees/*

# 3. Clé dédiée installée sur chaque copieur (ou une paire par copieur si le modèle le permet)
# 4. Tri automatique toutes les 5 minutes (script § 21)
# 5. Purge à 30 jours après tri (§ 69)
```

**Lien maintenance :** le scan-to-SFTP supprime la dépendance au SMB (plus de compte
de service AD à maintenir sur chaque MFP, plus de SMBv1 qui traîne). En cas de changement
de copieur, il suffit de déployer la clé sur la nouvelle machine.

**Dépannage typique :** « le copieur n'envoie plus » → vérifier (1) l'espace disque,
(2) que l'IP du copieur n'a pas changé (si restriction), (3) `auth.log` pour l'erreur exacte,
(4) l'horloge du copieur (certificats/clés sensibles au temps).

---

## 78. Cas pratique 4 — migration FTPS vers SFTP d'un partenaire (commenté)

**Contexte :** un partenaire historique utilise encore FTPS (vsftpd). Objectif : bascule sans coupure.

1. **Préparer** le compte SFTP en parallèle (`sftp_partenaire`), sans toucher au FTPS.
2. **Communiquer** : hôte, port 22, utilisateur, empreinte serveur, arborescence, protocole `.ok`.
3. **Doubler** les flux pendant 2 semaines : le partenaire envoie sur les deux canaux ;
   comparer les volumes (rapport § 71.2 vs logs vsftpd).
4. **Bascule** : le partenaire n'envoie plus qu'en SFTP ; lecture seule sur le FTPS.
5. **Extinction** : après 1 mois sans fichier sur le FTPS, couper vsftpd, fermer le port 21
   et la plage passive au pare-feu, archiver les logs.
6. **Bilan** : mettre à jour la fiche d'exploitation et prévenir la supervision
   (les sondes FTPS doivent pointer vers le SFTP).

---

## 79. Cas pratique 5 — export automatisé nocturne (commenté)

**Contexte :** l'ERP exporte chaque nuit à 1h un fichier vers le SFTP d'un partenaire.

```bash
#!/bin/bash
# /usr/local/bin/export_erp.sh — appelé par cron à 1h00
set -euo pipefail
CLE=/etc/erp_sftp/id_ed25519        # 0600, propriétaire erp
HOTE=sftp.partenaire.lan
USER=sftp_erp
SRC=/var/erp/exports
DEST=depot
DATE=$(date +%F)

# 1. L'ERP a déjà généré ${SRC}/export_${DATE}.csv (vérifier sa présence)
FICHIER="${SRC}/export_${DATE}.csv"
[ -f "$FICHIER" ] || { echo "export manquant"; exit 1; }

# 2. Checksum pour le contrôle d'intégrité côté partenaire
sha256sum "$FICHIER" > "${FICHIER}.sha256"

# 3. Envoi : fichier puis .sha256 puis témoin .ok
sftp -b /dev/stdin -i "$CLE" -o BatchMode=yes -o ConnectTimeout=20 "$USER@$HOTE" <<EOF
cd $DEST
put $FICHIER
put ${FICHIER}.sha256
put /dev/null export_${DATE}.csv.ok
bye
EOF

# 4. Traçabilité locale
logger -t export_erp "export $DATE envoyé, code $?"
```

```cron
0 1 * * * erp /usr/local/bin/export_erp.sh >>/var/log/export_erp.log 2>&1
```

**Garde-fous :** `BatchMode=yes` (pas de blocage), `ConnectTimeout`, nom de fichier daté
(idempotence), checksum, témoin `.ok`, log centralisé, alerte si le fichier source est absent,
clé dédiée (§ 44), rotation annuelle.

---

## 80. Quiz — 10 questions

1. Pourquoi le répertoire désigné par `ChrootDirectory` doit-il appartenir à `root:root`
   en `0755`, et que se passe-t-il sinon ?
2. Quelle différence entre `Subsystem sftp internal-sftp` et
   `Subsystem sftp /usr/lib/openssh/sftp-server` dans un contexte chroot ?
3. Un partenaire vous envoie sa clé **privée** par mail « pour configurer le serveur ».
   Que faites-vous ?
4. Où stocker les clés publiques des comptes chrootés, et pourquoi pas dans le chroot ?
5. À quoi sert l'option `-o BatchMode=yes` dans un script cron, et que se passe-t-il sans elle ?
6. Comment savoir, après coup, quel utilisateur a déposé quel fichier et quel volume ?
7. Un `put` échoue avec `Permission denied` alors que la connexion réussit : citez 3 causes possibles
   et leur diagnostic.
8. Pourquoi faut-il un protocole de fichier témoin (`.ok`) pour les flux automatisés ?
9. rsync-over-SSH ou SFTP pour un partenaire externe sans shell : lequel et pourquoi ?
10. Citez 4 directives du bloc `Match Group sftpusers` qui empêchent tout usage autre que le transfert.

---

## 81. Quiz — réponses

