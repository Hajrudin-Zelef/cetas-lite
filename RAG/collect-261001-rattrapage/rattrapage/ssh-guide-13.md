---
id: collect-261001-rattrapage/rattrapage/ssh-guide-13
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [2553, 2709]
sha256: a8ae12d9296bab05661d07b782041e900d137bba0804a341bb0c3a1c3c7c4292
---

# Guide SSH approfondi

```
# /etc/ssh/sshd_config.d/99-hardening.conf — base durcie
Port 22
PermitRootLogin no
PasswordAuthentication no
ChallengeResponseAuthentication no
UsePAM yes
PubkeyAuthentication yes
PermitEmptyPasswords no
AllowGroups ssh-users
MaxAuthTries 3
MaxSessions 4
LoginGraceTime 30
MaxStartups 10:30:60
ClientAliveInterval 300
ClientAliveCountMax 2
UseDNS no
X11Forwarding no
AllowTcpForwarding no
PermitTunnel no
LogLevel VERBOSE
Banner /etc/ssh/banner.txt
DebianBanner no
PrintMotd no
PrintLastLog yes
KexAlgorithms curve25519-sha256,curve25519-sha256@libssh.org,diffie-hellman-group16-sha512
Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes128-ctr
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com

# Ouvertures ciblées APRÈS (en fin de fichier) :
# Match Group tunnel-users
#     AllowTcpForwarding local
```

Checklist de mise en service : `sshd -t` → `reload` → test depuis session de
secours → `sshd -T | grep -Ei 'permitrootlogin|passwordauthentication'` →
fail2ban actif → logs qui partent au central.

## 83. Glossaire

| Terme | Définition |
|---|---|
| Bastion / jump host | Serveur passerelle, unique point d'entrée SSH vers un réseau privé |
| `authorized_keys` | Fichier listant les clés publiques autorisées pour un compte |
| CA (SSH) | Autorité qui signe des certificats SSH utilisateurs/hôtes |
| Certificat SSH | Clé publique signée par une CA, avec principals et dates de validité |
| Chiffrement symétrique | Chiffrement des données avec une clé partagée (rapide) |
| Cipher | Algorithme de chiffrement symétrique (ex. chacha20-poly1305) |
| ECDSA / Ed25519 | Algorithmes de signature ; ed25519 = choix moderne |
| Forward secrecy | Propriété : la compromission future d'une clé ne déchiffre pas le passé |
| GSSAPI | Authentification via Kerberos |
| Handshake | Négociation initiale : algos, échange de clés, authentifications |
| KEX | Key exchange : protocole d'établissement de la clé de session |
| `known_hosts` | Empreintes des serveurs déjà rencontrés (protection MITM) |
| MAC | Code d'authentification de message : garantit l'intégrité |
| MITM | Homme-du-milieu : attaquant interposé entre client et serveur |
| Multiplexage | Réutilisation d'une connexion TCP pour plusieurs sessions |
| PAM | Modules d'authentification pluggables (Linux) |
| Passphrase | Mot de passe protégeant une clé privée |
| Principal | Identité portée par un certificat SSH (ex. `zelef`) |
| ProxyJump | Relais SSH chiffré de bout en bout via un bastion (`-J`) |
| SCP / SFTP | Protocoles de copie de fichiers sur SSH |
| SOCKS | Protocole de proxy ; `-D` ouvre un proxy SOCKS via SSH |
| TOFU | Trust On First Use : on fait confiance à la 1re empreinte vue |
| TOTP | Mot de passe à usage unique basé sur le temps (2FA) |
| Tunnel `-L/-R/-D` | Redirection de port locale / distante / dynamique via SSH |
| 2FA | Authentification à deux facteurs |

## 84. Quiz (10 questions + réponses)

**Q1.** Pourquoi l'authentification par clé est-elle structurellement plus sûre
qu'un mot de passe, même robuste ?
> **R.** La clé privée ne voyage jamais : le client signe un défi unique que le
> serveur vérifie avec la clé publique. Un mot de passe, lui, est transmis
> (dans le tunnel chiffré, certes) et peut être rejoué s'il est intercepté ou
> deviné. La clé résiste au phishing et au brute-force en ligne.

**Q2.** Que signifie `StrictHostKeyChecking=accept-new`, et pourquoi est-il
préférable à `no` dans les scripts ?
> **R.** Les hôtes inconnus sont ajoutés automatiquement, mais tout
> **changement** d'empreinte reste bloquant. `no` désactive toute vérification
> et expose aux attaques homme-du-milieu.

**Q3.** Dans `~/.ssh/config`, quelle règle de priorité s'applique entre deux
blocs `Host` qui définissent la même option ?
> **R.** La **première** valeur obtenue gagne. D'où : blocs spécifiques avant
> le bloc `Host *` générique, placé en dernier.

**Q4.** Cite trois dangers de l'agent forwarding (`ssh -A`).
> **R.** 1) Root sur le serveur distant peut utiliser ton agent (signer avec
> tes clés) tant que tu es connecté. 2) Le socket peut survivre à la
> déconnexion. 3) Il est inutile dans la plupart des cas : `ProxyJump`
> authentifie depuis ton poste sans exposer l'agent.

**Q5.** Quelle différence entre `ssh -L` et `ssh -D` ?
> **R.** `-L` redirige un port local fixe vers **une** destination fixée à
> l'ouverture ; `-D` ouvre un proxy **SOCKS** dont la destination est choisie
> connexion par connexion par l'application cliente.

**Q6.** À quoi sert `command=` dans `authorized_keys`, et avec quelles options
l'accompagner ?
> **R.** À forcer l'exécution d'une commande précise en ignorant celle du
> client (clés de service : backup, CI). À accompagner de `no-pty`,
> `no-agent-forwarding`, `no-X11-forwarding`, `no-port-forwarding` et
> idéalement `from=`.

**Q7.** Pourquoi une connexion SSH peut-elle mettre 20 secondes à s'établir
sur un réseau sain ? Deux causes et leurs remèdes.
> **R.** 1) `UseDNS yes` côté serveur : résolution inverse lente →
> `UseDNS no`. 2) `GSSAPIAuthentication yes` côté client : tentative Kerberos
> qui timeout → `GSSAPIAuthentication no`.

**Q8.** Que faire quand `ssh` répond `Too many authentication failures` ?
> **R.** Le client propose plus de clés que `MaxAuthTries` : forcer la bonne
> avec `-o IdentitiesOnly=yes -i <clé>`, puis figer ça dans le bloc `Host`
> (`IdentityFile` + `IdentitiesOnly yes`).

**Q9.** Quelles permissions doivent avoir `~/.ssh`, la clé privée, et
`authorized_keys` côté serveur (avec home) ?
> **R.** `~/.ssh` : `700` ; clé privée : `600` ; `~/.ssh/authorized_keys` :
> `600` ; home : `755` maximum (sinon `StrictModes` fait refuser la clé).

**Q10.** En une phrase : quel est l'apport principal des certificats SSH par
rapport aux `authorized_keys` ?
> **R.** La CA signe les clés au lieu de les déployer sur chaque serveur :
> onboarding/offboarding instantanés, expiration intégrée, révocation
> centralisée, et fin du TOFU pour les clés d'hôte.

## 85. Pour aller plus loin

- **Docs officielles** : `man ssh`, `man ssh_config`, `man sshd`,
  `man sshd_config`, `man ssh-keygen` — la référence absolue, à jour avec ta
  version (`man` > n'importe quel tuto).
- **Protocole** : RFC 4251 (architecture), 4252 (auth), 4253 (transport),
  4254 (connexion).
- **Mozilla Infosec** : « OpenSSH security guidelines » — recommandations de
  configuration durcies, mises à jour régulièrement.
- **SSH Certificates** : la doc de `ssh-keygen` (options `-s`, `-I`, `-n`,
  `-V`, `-z`) + les articles de Facebook/Meta et Netflix sur le SSH à grande
  échelle (bastions + certificats éphémères).
- **Outils** : `autossh` (tunnels persistants), `mosh` (shell sur liens
  instables — pas un remplacement de SSH, un complément), `sshuttle` (VPN
  paresseux sur SSH), `ansible` (déploiement d'`authorized_keys` et de
  `sshd_config` à l'échelle).
- **Pratique** : monte un labo (3 VM : poste, bastion, cible), applique les
  sections 42-48, casse volontairement la config et dépanne avec la méthode
  de la section 61. C'est en réparant qu'on apprend.
- **Veille** : suis les annonces OpenSSH (nouvelles versions = nouveaux
  algos, retraits d'anciens) et les CVE ; teste chaque montée de version
  avec `sshd -T` et tes cas limites (section 60).

---

*Fin du guide. 85 sections. Prochaine étape suggérée : versionner ton
`~/.ssh/config.d/` et ton `sshd_config.d/` dans le dépôt d'exploitation de
l'équipe, puis planifier l'inventaire trimestriel des `authorized_keys`
(section 59).*
