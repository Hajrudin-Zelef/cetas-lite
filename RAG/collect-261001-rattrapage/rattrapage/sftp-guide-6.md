---
id: collect-261001-rattrapage/rattrapage/sftp-guide-6
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [1230, 1436]
sha256: 04ec8c3e4db8df7a0f9ab10d708e08013d8b32cef16ea56a4fabb499da10fe9e
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

```
listen=YES
anonymous_enable=NO
local_enable=YES
write_enable=YES
chroot_local_user=YES
allow_writeable_chroot=YES
ssl_enable=YES
force_local_logins_ssl=YES
force_local_data_ssl=YES
ssl_tlsv1_2=YES
rsa_cert_file=/etc/ssl/certs/vsftpd.pem
rsa_private_key_file=/etc/ssl/private/vsftpd.key
pasv_min_port=50000
pasv_max_port=50100
```

Points de vigilance :

- Certificat TLS valide (Let's Encrypt ou PKI interne), TLS 1.2 minimum.
- Ouvrir le port 21 **et** la plage passive dans le pare-feu (contraignant : c'est
  l'un des gros désavantages de FTPS vs SFTP).
- `chroot_local_user=YES` + comptes sans shell comme pour SFTP.
- Préférez toujours renégocier vers SFTP quand c'est possible.

---

## 57. Couper le FTP classique (port 21 non chiffré)

- Ne **jamais** laisser `ssl_enable=NO` avec des comptes réels.
- Si un vsftpd historique tourne encore en clair : plan de migration
  (informer les partenaires, délai, bascule, extinction).
- Surveillez : `sudo ss -tlnp | grep :21` ne doit rien afficher une fois la migration terminée.

---

## 58. Mot de passe : quand on ne peut pas l'éviter (copieurs anciens)

Certains MFP anciens ne gèrent pas les clés SSH. Dernier recours, ciblé :

```
Match User svc_copieur_ancien
    ChrootDirectory /srv/sftp/scans_copieurs
    ForceCommand internal-sftp
    PasswordAuthentication yes
```

- Mot de passe **long et unique** (généré, pas choisi), stocké au coffre.
- Restreindre l'IP source si le copieur a une IP fixe (`AllowUsers`).
- Planifier le remplacement du matériel : le mot de passe est une dette de sécurité.
- Surveiller les connexions de ce compte dans les logs (alerte sur toute anomalie).

---

## 59. Dépannage — méthode générale

1. Reproduire avec `sftp -v` (verbeux) : la cause est souvent dans les 20 dernières lignes.
2. Regarder côté serveur en temps réel : `sudo journalctl -u ssh -f` et `/var/log/auth.log`.
3. Vérifier dans l'ordre : réseau/port → authentification → autorisations (Match) → chroot/permissions → quotas/disque.
4. Ne changer qu'**une** chose à la fois, `sshd -t`, `reload`, retester.

```bash
# Client : mode verbeux
sftp -v -i ~/.ssh/id_sftp_acme sftp_acme@sftp.entreprise.lan
# Serveur : suivre les logs
sudo tail -f /var/log/auth.log
```

---

## 60. Cas 1 — « Permission denied » dans le chroot

**Symptômes :** connexion OK, mais `put` échoue : `Permission denied`.

**Causes fréquentes :**

| Cause | Diagnostic | Remède |
|---|---|---|
| Le dossier appartient à root | `ls -la` dans le chroot | `chown user:groupe dossier` |
| Droits 0755 sur le dossier de dépôt | `stat dossier` | `chmod 0750` + bon propriétaire |
| Quota atteint | `repquota /srv` | augmenter ou purger |
| Disque plein | `df -h /srv` | libérer de l'espace |
| Fichier existant non inscriptible | `ls -l fichier` | vérifier le propriétaire |

> Rappel : le **chroot lui-même** doit rester `root:root 0755` ; c'est le **sous-dossier**
> de dépôt qui doit appartenir à l'utilisateur.

---

## 61. Cas 2 — la connexion se ferme aussitôt

**Symptômes :** `Connection closed` / `client_loop: send disconnect` juste après l'authentification.

**Pistes :**

1. Permissions du chroot non conformes (composant inscriptible par l'utilisateur) → sshd refuse silencieusement. Vérifier avec `namei -l /srv/sftp/...`.
2. `ChrootDirectory` pointe vers un chemin inexistant → `sshd -T | grep chroot`.
3. `ForceCommand internal-sftp` + `Subsystem sftp /usr/lib/openssh/sftp-server` incohérent → uniformiser sur `internal-sftp`.
4. Shell inexistant dans `/etc/shells` ? Avec `internal-sftp` ce n'est pas requis, mais un `Match` mal ordonné peut appliquer un autre bloc.
5. Regarder `auth.log` : `fatal: bad ownership or modes for chroot directory` = diagnostic direct.

---

## 62. Cas 3 — « Bad ownership or modes for chroot directory »

C'est le message explicite du cas 2. Procédure de réparation :

```bash
C=/srv/sftp/client_acme
sudo chown root:root "$C"
sudo chmod 0755 "$C"
# remonter la hiérarchie : /srv, /srv/sftp doivent aussi être root:root sans w groupe/autre
sudo chown root:root /srv /srv/sftp
sudo chmod 0755 /srv /srv/sftp
namei -l "$C"   # contrôle visuel final
sudo systemctl reload ssh
```

> Piège classique : avoir fait `chown -R sftp_acme /srv/sftp/client_acme`
> « pour que ça marche » — c'est exactement ce qui casse le chroot.

---

## 63. Cas 4 — clé refusée (« Permission denied (publickey) »)

Checklist :

- [ ] La clé **publique** (pas la privée !) est dans `/etc/ssh/authorized_keys/<user>`
- [ ] Fichier `0644 root:root`, dossier `0755` — sshd est strict aussi ici
- [ ] Pas de saut de ligne coupé / pas de retour chariot Windows dans la clé
- [ ] `AuthorizedKeysFile` pointe bien vers `/etc/ssh/authorized_keys/%u`
- [ ] Le compte n'est pas verrouillé par `DenyUsers` / `AllowUsers`
- [ ] Côté client : la bonne clé est proposée (`sftp -v` montre les clés essayées)
- [ ] L'empreinte dans `auth.log` correspond à la clé attendue

```bash
# Comparer l'empreinte proposée par le client et celle installée
sudo ssh-keygen -lf /etc/ssh/authorized_keys/sftp_acme
```

---

## 64. Cas 5 — « subsystem request failed on channel 0 »

- Le client demande le subsystem `sftp` mais le serveur ne le propose pas :
  vérifier la ligne `Subsystem sftp internal-sftp` (décommentée, syntaxe valide).
- Après modification : `sshd -t` puis `reload` (un `reload` oublié est la cause n°1).
- Conflit avec un `Match` qui écrase le subsystem ? `sshd -T -C user=sftp_acme,addr=... | grep subsystem`.

---

## 65. Cas 6 — transferts lents

| Piste | Action |
|---|---|
| Fenêtre TCP / latence | `sftp -R 64 -B 262144` (plus de requêtes en vol, gros buffers) |
| Chiffrement CPU-bound | Choisir une cipher rapide : `-o Ciphers=aes128-gcm@openssh.com` |
| Parallélisme | `lftp` avec `--parallel=4`, ou découper en plusieurs fichiers |
| Disque saturé côté serveur | `iostat -x 1`, vérifier le stockage |
| QoS / shaping opérateur | Tester hors heures ouvrées, mesurer avec `iperf3` |

```bash
# Test comparatif rapide
time sftp -b batch.txt -i cle user@hote   # batch avec un get d'un gros fichier
```

---

## 66. Cas 7 — caractères accentués / espaces dans les noms de fichiers

- SFTP transmet les noms en UTF-8 : forcer `LC_ALL=C.UTF-8` côté scripts si la locale est en POSIX.
- Dans les batchs, quoter les noms avec espaces : `put "mon fichier.pdf"`.
- Côté Windows (WinSCP), vérifier le jeu de caractères : UTF-8 par défaut, à conserver.
- Pour les flux automatisés : **normaliser** les noms à la source (sans accents, sans espaces,
  `facture_2026-09-26.pdf`) — voir § 46.

---

## 67. Cas 8 — le partenaire « ne voit pas » le fichier déposé

- Le fichier est-il complet ? Protocole `.ok` (voir § 46) : sans lui, le consommateur attend.
- Droits : fichier déposé en `0600` par défaut (umask du client) → le consommateur (autre UID)
  ne peut pas le lire. Harmoniser avec `umask` ou forcer les droits après dépôt.
- Mauvais dossier : `pwd` dans le chroot ≠ chemin absolu du serveur. Documenter l'arborescence vue par le client.
- Horloge désynchronisée : un fichier « du futur » peut perturber les tris par date → NTP partout (chrony).

---

## 68. Cas 9 — « Disk quota exceeded » / « No space left on device »

```bash
df -h /srv                          # disque plein ?
df -i /srv                          # inodes épuisés ? (beaucoup de petits fichiers)
sudo repquota -u /srv | grep sftp_  # quota utilisateur atteint ?
```

Actions :

1. Identifier les gros consommateurs : `du -sh /srv/sftp/*`.
2. Purger selon la politique de rétention (voir § 69) — jamais à la main sans traçabilité.
3. Augmenter le quota si légitime (`setquota`), avec validation.
4. Alerter : un dépôt qui grossit brutalement peut signaler un dysfonctionnement chez l'émetteur.

---

## 69. Cas 10 — saturation par accumulation (politique de rétention)

