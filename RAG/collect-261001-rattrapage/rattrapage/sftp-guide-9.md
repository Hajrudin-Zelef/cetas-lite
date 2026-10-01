---
id: collect-261001-rattrapage/rattrapage/sftp-guide-9
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [1773, 1925]
sha256: 58e64d5f35f969f2224a8ef8ce27e4f2cc58f002b111ec96fb97419f555d06ae
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

1. OpenSSH exige que chaque composant du chemin du chroot appartienne à root et ne soit
   inscriptible par personne d'autre (sécurité : sinon l'utilisateur pourrait modifier son
   environnement racine). Sinon : `fatal: bad ownership or modes for chroot directory`
   et connexion fermée aussitôt (voir § 62).
2. `internal-sftp` est intégré à sshd : aucune dépendance externe, il fonctionne dans un
   chroot vide. Le binaire `sftp-server` exigerait de recopier libc et devices dans chaque
   chroot. `internal-sftp` accepte aussi `-f`/`-l` pour la journalisation (voir § 6, § 30).
3. On refuse : la clé privée ne voyage jamais. On demande la clé **publique** (`.pub`),
   transmise par un canal authentifié, et on rappelle au partenaire de régénérer sa paire
   puisque sa privée a circulé en clair (voir § 12, § 44).
4. Dans `/etc/ssh/authorized_keys/%u` (centralisé, géré par root, hors chroot) : le chroot
   appartient à root, l'utilisateur ne pourrait pas gérer ses clés, et sshd impose des
   permissions strictes que le chroot complique (voir § 18).
5. `BatchMode=yes` fait échouer proprement au lieu d'afficher un prompt (mot de passe,
   host key) : sans elle, le cron se bloque en attendant une entrée qui ne viendra jamais (voir § 41, § 43).
6. Avec `Subsystem sftp internal-sftp -f AUTH -l VERBOSE` : les lignes `open ... WRITE/READ`
   et `close ... bytes read/written` dans `/var/log/sftp.log`, corrélées aux lignes
   `Accepted publickey for <user> from <ip>` de `auth.log` (voir § 30, § 32).
7. (a) Le dossier de dépôt appartient à root ou n'est pas inscriptible → `ls -la`/`stat` ;
   (b) quota atteint → `repquota` ; (c) disque plein → `df -h` (voir § 60).
8. Le `.ok` n'est déposé qu'après le transfert complet : le consommateur ne traite que les
   fichiers accompagnés de leur témoin, ce qui élimine les fichiers partiels (voir § 46).
9. **SFTP** : rsync exige un shell et le binaire rsync côté distant, incompatible avec
   `ForceCommand internal-sftp` ; SFTP fonctionne nativement avec des comptes sans shell
   en chroot (voir § 47).
10. `ForceCommand internal-sftp`, `AllowTcpForwarding no`, `AllowAgentForwarding no`,
    `X11Forwarding no` (et `PermitTunnel no`) — voir § 10 et § 39.

---

## 82. Pour aller plus loin

- **Documentation OpenSSH** : `man sshd_config` (directives `Match`, `ChrootDirectory`,
  `ForceCommand`), `man sftp`, `man sftp-server`.
- **Automatisation avancée** : remplacez les batchs shell par Ansible (module `ansible.builtin.copy`
  via SFTP, gestion des `authorized_keys` en masse).
- **Haute disponibilité** : deux serveurs SFTP derrière un VIP (keepalived), `/srv/sftp`
  répliqué (DRBD ou rsync), clés hôtes identiques sur les deux nœuds.
- **Chiffrement de bout en bout** : GPG systématique pour les flux sensibles (§ 26),
  avec gestion des clés via un trousseau d'équipe.
- **SIEM** : centralisez `/var/log/sftp.log` et `auth.log` (rsyslog TLS, § 36),
  créez des tableaux de bord (volumes par client, échecs par IP).
- **Conformité** : alignez rétention et journalisation sur vos obligations
  (bancaires, clients, RGPD : minimisation et durées définies).
- **Guides compagnons** : `debian_ubuntu_guide.md` (bases système), `proxmox_guide.md`
  (virtualisation et sauvegarde), `onduleurs_ups_guide.md` (continuité électrique du serveur).

---

## 83. Modèle de fiche d'exploitation par compte

```
Compte            : sftp_acme
Périmètre/chroot  : /srv/sftp/client_acme
Usage             : dépôt nocturne des bons de commande (CSV)
Responsable int.  : _______________      Contact partenaire : _______________
Volumétrie        : ~50 Mo/jour          Rétention : 90 jours
Quota             : 8 Go (soft) / 8,5 Go (hard)
Clé installée le  : _______________      Rotation prévue le : _______________
Empreinte clé     : SHA256:_______________
Restriction IP    : _______________
Créé le           : _______________ par _______________
Observations      : _______________
```

---

## 84. Checklist de mise en production du serveur

- [ ] OpenSSH installé, `sshd_config` durci (§ 17) et testé (`sshd -t`)
- [ ] `Subsystem sftp internal-sftp -f AUTH -l VERBOSE` actif
- [ ] Bloc `Match Group sftpusers` en fin de fichier (§ 10)
- [ ] Groupe `sftpusers` créé, comptes sans shell, mots de passe verrouillés
- [ ] Chroots `root:root 0755` vérifiés (`check_chroot.sh`, § 71.1)
- [ ] Clés publiques installées (`/etc/ssh/authorized_keys/%u`), test de connexion OK
- [ ] Quotas appliqués et alertes configurées (§ 24, § 35)
- [ ] fail2ban actif sur `sshd` (§ 38, § 55)
- [ ] Pare-feu : 22/TCP restreint aux sources légitimes (§ 37)
- [ ] Logs SFTP séparés + logrotate + centralisation (§ 31, § 36)
- [ ] Rapport quotidien des transferts (§ 71.2) et sonde de service (§ 34)
- [ ] Sauvegarde `/srv/sftp` + `/etc/ssh` testée (restauration à blanc, § 28)
- [ ] Politique de rétention définie par périmètre (§ 69)
- [ ] Fiches d'exploitation remplies (§ 83)
- [ ] Procédure de révocation de clé connue de l'astreinte (§ 52)

---

## 85. Conclusion

Un serveur SFTP d'entreprise tient sur cinq piliers : **chroot systématique**,
**clés uniquement**, **pas de shell**, **journalisation qui dit qui a transféré quoi**,
et **exploitation** (quotas, supervision, sauvegarde, rétention). Avec ce socle,
les dépôts clients, les échanges bancaires et les scans des copieurs reposent sur
une infrastructure simple, auditable et robuste — un seul port, un seul service,
des procédures écrites. Relisez ce guide à chaque onboarding de partenaire, et
faites-en la référence de votre équipe.

---

## 86. Durcissement cryptographique (ciphers, KEX, MAC)

Restreignez les algorithmes aux valeurs modernes dans `sshd_config` :

```
# N'autoriser que les échanges de clés robustes
KexAlgorithms curve25519-sha256,curve25519-sha256@libssh.org,ecdh-sha2-nistp521,ecdh-sha2-nistp384,ecdh-sha2-nistp256,diffie-hellman-group18-sha512
# Chiffrements AEAD / modernes uniquement
Ciphers aes128-gcm@openssh.com,aes256-gcm@openssh.com,chacha20-poly1305@openssh.com,aes256-ctr,aes128-ctr
# MAC avec EtM (encrypt-then-mac)
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com
# Clés hôtes : privilégier Ed25519
HostKey /etc/ssh/ssh_host_ed25519_key
```

Vérifiez ce que négocie réellement un client :

```bash
# Côté client, en verbeux : repérer "kex: algorithm:" et "cipher:"
sftp -v sftp_acme@sftp.entreprise.lan 2>&1 | grep -E "kex:|cipher|MAC"
# Scanner la configuration offerte par le serveur :
nmap --script ssh2-enum-algos -p 22 sftp.entreprise.lan
```

> Attention : un durcissement trop agressif peut casser de vieux clients/partenaires
> (copieurs anciens, librairies Java datées). Testez chaque partenaire après changement
> et gardez une fenêtre de compatibilité documentée.

---

## 87. SFTP derrière un bastion (jump host)

Si le serveur SFTP n'est pas exposé directement sur Internet :

```bash
# Le client passe par le bastion (une seule commande)
sftp -J admin@bastion.entreprise.lan sftp_acme@sftp-prive.entreprise.lan

# Ou en ~/.ssh/config côté client :
Host sftp-entreprise
    HostName sftp-prive.entreprise.lan
    User sftp_acme
    IdentityFile ~/.ssh/id_sftp_acme
    ProxyJump admin@bastion.entreprise.lan
```

Avantages : le serveur SFTP n'a que le bastion comme source autorisée sur le port 22
(`AllowUsers *@bastion` ou règle pare-feu), journalisation centralisée des accès,
pas d'exposition directe. Inconvénient : le bastion devient critique (HA, supervision).

---

## 88. Montée en charge et tuning

