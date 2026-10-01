---
id: collect-261001-rattrapage/rattrapage/sftp-guide-1
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [1, 215]
sha256: 8342465c7b6498f9de2b294c5a1f4cbf744c0f27a22dd93dc6ac34ff71cac492
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

**Public :** chef de service systèmes & énergies, administrateurs, techniciens d'exploitation.
**Objectif :** mettre en place, sécuriser, superviser et dépanner un serveur SFTP d'entreprise
(OpenSSH, Debian/Ubuntu) : dépôts clients, échanges comptables/bancaires, réception de scans de copieurs.
**Pré-requis :** notions d'administration Linux (Debian/Ubuntu), SSH, permissions Unix.

> Ce guide suit les bonnes pratiques « production » : clés uniquement, chroot systématique,
> pas de shell pour les comptes SFTP, journalisation détaillée, quotas, supervision et sauvegarde.

---

## 1. SFTP : de quoi parle-t-on ?

- **SFTP (SSH File Transfer Protocol)** n'a rien à voir avec FTP : c'est un sous-protocole de SSH (port 22/TCP).
- Tout transite dans un tunnel chiffré : authentification, commandes, données.
- Remplace avantageusement FTP/FTPS en entreprise : un seul port à ouvrir, un seul service à durcir.
- Le serveur SFTP d'OpenSSH se présente sous deux formes :
  - le binaire externe `sftp-server` ;
  - le serveur intégré `internal-sftp` (recommandé : fonctionne dans un chroot sans binaires externes).

| Protocole | Port(s) | Chiffrement | Recommandation |
|---|---|---|---|
| FTP | 21 + ports passifs | Aucun | À bannir |
| FTPS (FTP+TLS) | 21 + ports passifs | TLS | Toléré si partenaire l'impose (cf. § 62) |
| SFTP | 22 | SSH | **Recommandé** |
| SCP | 22 | SSH | Obsolète (utiliser SFTP) |
| rsync-over-SSH | 22 | SSH | Excellent pour synchro/miroir (cf. § 47) |

---

## 2. Glossaire rapide (version détaillée au § 74)

| Terme | Sens |
|---|---|
| chroot | Enfermer un utilisateur dans un sous-dossier qui devient sa racine `/` |
| subsystem | Sous-système SSH invoqué à la demande (ici `sftp`) |
| ForceCommand | Directive sshd qui impose la commande exécutée, quel que soit le client |
| Match | Bloc conditionnel dans `sshd_config` (par utilisateur, groupe, adresse…) |
| internal-sftp | Serveur SFTP intégré à sshd, sans dépendance externe |
| Batch mode | Mode non interactif de `sftp` (`-b fichier`) pour les scripts |
| Quota | Limite d'espace disque par utilisateur/groupe |
| Clé SSH | Paire cryptographique (publique/privée) pour s'authentifier sans mot de passe |

---

## 3. Architecture cible

```
/srv/sftp/                  <- racine des chroots (root:root, 0755)
├── client_acme/            <- chroot de l'utilisateur sftp_acme
│   ├── depot/              <- inscriptible par sftp_acme (dépôt entrant)
│   ├── restitutions/       <- inscriptible par l'exploitant, lecture seule pour le client
│   └── .ssh/               <- NON : les clés restent hors chroot !
├── compta/                <- chroot du service comptabilité
│   ├── envoi_banque/
│   └── recus_banque/
└── scans_copieurs/         <- chroot du compte technique des copieurs
    └── entrees/            <- les MFP y déposent les PDF numérisés
```

Principes :

1. **Un chroot par périmètre** (client, service, usage), jamais un chroot unique partagé.
2. Le répertoire chroot appartient à `root:root` en `0755` (exigence OpenSSH : chaque composant du chemin doit appartenir à root et ne pas être inscriptible par l'utilisateur).
3. Les dossiers de dépôt à l'intérieur appartiennent à l'utilisateur SFTP.
4. Les clés publiques sont gérées **hors chroot** (`/etc/ssh/authorized_keys/%u` ou `~/.ssh` si pas de chroot sur le home — voir § 18).

---

## 4. Installation d'OpenSSH Server (Debian/Ubuntu)

```bash
sudo apt update
sudo apt install -y openssh-server
sudo systemctl enable --now ssh
ssh -V            # vérifie la version, ex : OpenSSH_9.6p1
ss -tlnp | grep :22
```

Fichiers importants :

| Fichier | Rôle |
|---|---|
| `/etc/ssh/sshd_config` | Configuration du serveur |
| `/etc/ssh/sshd_config.d/` | Fragments de configuration (priorité selon l'ordre) |
| `/var/log/auth.log` | Journal d'authentification (Debian/Ubuntu) |
| `/var/log/syslog` | Journal système (si SyslogFacility AUTH par défaut) |

> Sauvegardez toujours `sshd_config` avant de le modifier :
> `sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.bak-$(date +%F)`

---

## 5. Vérifier la configuration avant de recharger

```bash
sudo sshd -t          # teste la syntaxe ; silencieux = OK
sudo sshd -T | grep -i -E 'chroot|forcecommand|subsystem|passwordauth'
sudo systemctl reload ssh
```

Règle d'or : **ne jamais fermer votre session SSH active** avant d'avoir testé une nouvelle
connexion dans un second terminal. Une erreur de `sshd_config` peut vous enfermer dehors.

---

## 6. Le subsystem SFTP : internal-sftp vs sftp-server

Dans `/etc/ssh/sshd_config` :

```
# Recommandé : serveur intégré, compatible chroot sans binaires
Subsystem sftp internal-sftp

# Ancienne forme (binaire externe) — à éviter avec chroot :
# Subsystem sftp /usr/lib/openssh/sftp-server
```

Pourquoi `internal-sftp` ?

- Aucune dépendance (libc, /dev/*) à recopier dans chaque chroot.
- Accepte les options `-f` (niveau de log) et `-l` (facility) directement.
- Se combine avec `ChrootDirectory` + `ForceCommand internal-sftp` sans bidouillage.

Exemple avec journalisation renforcée (voir § 50) :

```
Subsystem sftp internal-sftp -f AUTH -l VERBOSE
```

---

## 7. Créer le groupe sftpusers

```bash
sudo groupadd sftpusers
```

Tous les comptes « transfert uniquement » seront membres de ce groupe. Les blocs `Match Group sftpusers`
dans `sshd_config` appliqueront d'un coup : chroot, pas de shell, clés uniquement.

Convention de nommage proposée :

| Préfixe | Usage | Exemple |
|---|---|---|
| `sftp_` | Comptes clients/partenaires externes | `sftp_acme`, `sftp_banque` |
| `svc_` | Comptes techniques internes | `svc_copieur_etage2`, `svc_compta` |
| `adm_` | Comptes d'exploitation (hors chroot) | `adm_sftp` |

---

## 8. Créer un utilisateur SFTP sans shell

```bash
# Compte client ACME : pas de shell, pas de mot de passe, groupe sftpusers
sudo useradd -m -d /srv/sftp/client_acme -s /usr/sbin/nologin \
  -G sftpusers -c "Dépôt SFTP client ACME" sftp_acme
sudo passwd -l sftp_acme      # verrouille le mot de passe
```

- `-s /usr/sbin/nologin` : aucun shell interactif possible.
- `-G sftpusers` : rattache au groupe (groupe principal laissé à son groupe privé).
- `passwd -l` : verrouille l'authentification par mot de passe (les clés SSH restent utilisables).

Vérification :

```bash
id sftp_acme
getent passwd sftp_acme
```

---

## 9. Arborescence chroot correcte (permissions root:root)

```bash
sudo mkdir -p /srv/sftp/client_acme/{depot,restitutions}
sudo chown root:root /srv/sftp /srv/sftp/client_acme
sudo chmod 0755 /srv/sftp /srv/sftp/client_acme
sudo chown sftp_acme:sftpusers /srv/sftp/client_acme/depot
sudo chmod 0750 /srv/sftp/client_acme/depot
sudo chown root:sftpusers /srv/sftp/client_acme/restitutions
sudo chmod 0750 /srv/sftp/client_acme/restitutions
```

Règles OpenSSH (strictes, sinon la connexion échoue) :

- Chaque composant du chemin du `ChrootDirectory` doit appartenir à **root**.
- Aucun composant ne doit être **inscriptible** par l'utilisateur ou son groupe.
- À l'intérieur du chroot, créez des sous-dossiers dont l'utilisateur est propriétaire.

Test de conformité (script, voir § 71) :

```bash
namei -l /srv/sftp/client_acme
# tout, jusqu'à client_acme inclus : root root, sans bit w pour group/other
```

---

## 10. Bloc Match dans sshd_config (le cœur du dispositif)

```
# --- SFTP entreprise : comptes de transfert uniquement ---
Match Group sftpusers
    ChrootDirectory /srv/sftp/%u
    ForceCommand internal-sftp
    PasswordAuthentication no
    PermitTunnel no
    AllowAgentForwarding no
    AllowTcpForwarding no
    X11Forwarding no
```

