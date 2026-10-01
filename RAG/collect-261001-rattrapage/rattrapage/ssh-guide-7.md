---
id: collect-261001-rattrapage/rattrapage/ssh-guide-7
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["incident", "valuation"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [1306, 1515]
sha256: 3b7665d69f7b1931902f856801e64b7dab29a8db8b0b8b59f72054360f2e1d22
---

# Guide SSH approfondi

- `PermitRootLogin no` : root ne se connecte plus directement. Passe par un
  utilisateur normal + `sudo` (traçabilité : on sait **qui** est devenu root).
  `prohibit-password` n'autorise root que par clé : compromis acceptable en
  transition, mais `no` reste la cible.
- `PasswordAuthentication no` : coupe les attaques par dictionnaire (99 % du
  bruit Internet). **Vérifie que tes clés fonctionnent avant** (sinon tu te
  enfermes dehors — d'où la session de secours, section 41).
- `UsePAM yes` : garde PAM pour la gestion de session/motd même sans auth par
  mot de passe. Avec `ChallengeResponseAuthentication no`, PAM ne proposera pas
  de 2FA par défaut (section 49 pour l'activer proprement).
- `PermitEmptyPasswords no` (défaut) : à vérifier explicitement.

Vérification :

```bash
serveur# sshd -T | grep -Ei 'permitrootlogin|passwordauthentication|pubkeyauthentication'
permitrootlogin no
passwordauthentication no
pubkeyauthentication yes
```

## 43. Durcissement : `AllowUsers` / `AllowGroups`

Liste blanche : seuls ces utilisateurs/groupes peuvent se connecter. Tout le
reste est rejeté **avant** même l'authentification.

```
# /etc/ssh/sshd_config.d/20-access.conf
AllowGroups ssh-users admins
# ou, plus fin :
# AllowUsers zelef deploy@10.0.0.0/8
DenyUsers backup-test
```

- `AllowUsers` accepte le format `user@hôte-ou-motif` : `deploy@10.0.5.*`
  restreint ce compte au réseau de déploiement.
- Crée un groupe `ssh-users` : l'ajout/retrait d'accès devient `usermod -aG`,
  traçable et réversible, au lieu d'éditer la config SSH.
- Ordre d'évaluation : `DenyUsers`, `AllowUsers`, `DenyGroups`, `AllowGroups`.
  Si `AllowUsers`/`AllowGroups` est présent, seuls les listés passent.

## 44. Durcissement : `MaxAuthTries`, `MaxSessions`, `LoginGraceTime`

```
# /etc/ssh/sshd_config.d/30-limits.conf
MaxAuthTries 3            # essais d'authentification par connexion (défaut 6)
MaxSessions 4             # sessions multiplexées max par connexion réseau
LoginGraceTime 30         # 30 s pour s'authentifier, sinon déconnexion
MaxStartups 10:30:60      # anti-flood : voir ci-dessous
ClientAliveInterval 300   # détecte les sessions mortes côté serveur
ClientAliveCountMax 2
```

- `MaxStartups 10:30:60` : à partir de 10 connexions non authentifiées
  simultanées, 30 % de rejet, jusqu'à 100 % à 60. Freine les floods sans
  bloquer un usage légitime (Ansible ouvre beaucoup de connexions : ajuste si
  besoin, ex. `20:30:100`).
- `ClientAliveInterval`/`ClientAliveCountMax` : le serveur coupe les sessions
  mortes (complément de `ServerAliveInterval` côté client, section 13).
- Ces limites réduisent l'impact des scans mais **ne remplacent pas fail2ban**
  (section 48).

## 45. Algorithmes modernes : Ciphers, MACs, KexAlgorithms

```
# /etc/ssh/sshd_config.d/40-crypto.conf
KexAlgorithms curve25519-sha256,curve25519-sha256@libssh.org,diffie-hellman-group16-sha512
Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes128-ctr
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com
HostKeyAlgorithms ssh-ed25519,rsa-sha2-512,rsa-sha2-256
CASignatureAlgorithms ssh-ed25519,rsa-sha2-512,rsa-sha2-256
```

- `-etm` (encrypt-then-mac) : les MAC modernes, à préférer aux `-etm`-less.
- Ce qui est **exclu** ici (et c'est voulu) : `3des`, `blowfish`, `arcfour`,
  `hmac-md5`, `hmac-sha1`, `diffie-hellman-group1-sha1`, `ssh-dss`.
- **Compatibilité** : des clients très anciens (vieux équipements réseau) ne
  supporteront pas cette liste. Plutôt que d'affaiblir le serveur global,
  utilise un `Match` ciblé (section 14 côté serveur) ou mets ces équipements
  derrière un bastion avec une config dédiée.
- Après changement : `sshd -t && systemctl reload ssh`, puis teste depuis tes
  clients habituels (`ssh -v` montre les algorithmes négociés, ligne
  `debug1: kex: ...`).

## 46. Désactiver ce qui ne sert pas

```
# /etc/ssh/sshd_config.d/50-features.conf
X11Forwarding no
AllowTcpForwarding no        # coupe -L/-R/-D ... à assouplir par Match si besoin
PermitTunnel no              # pas de VPN layer3 sur SSH
# Si les tunnels sont nécessaires pour certains utilisateurs seulement :
# Match Group tunnel-users
#     AllowTcpForwarding local
```

Stratégie : **tout fermé par défaut**, ouverture au cas par cas via `Match`
(en fin de fichier, rappel section 41) :

```
Match Group dev-tunnel
    AllowTcpForwarding local
    PermitOpen db-interne:5432
```

`PermitOpen` restreint les destinations des tunnels locaux : même avec un
tunnel autorisé, l'utilisateur ne peut viser que la base de données prévue.
Défense en profondeur concrète.

## 47. Bannières et informations divulguées

```
# /etc/ssh/sshd_config.d/60-banner.conf
Banner /etc/ssh/banner.txt
DebianBanner no              # ne pas annoncer "Debian-*" dans la version
VersionAddendum none         # OpenSSH >= 9.8 : masque la version patch
PrintMotd no                 # le motd via PAM suffit
PrintLastLog yes             # utile : "Last login: ..." alerte en cas d'accès anormal
```

Contenu type de `/etc/ssh/banner.txt` :

```
********************************************************************
* Accès réservé aux personnes autorisées. Toute activité est       *
* journalisée. Déconnectez-vous immédiatement si vous n'êtes pas   *
* autorisé à accéder à ce système.                                 *
********************************************************************
```

Pourquoi : la bannière légale aide en cas d'incident (prouve l'avertissement),
et masquer la version exacte complique le ciblage automatisé des CVE. Ce
n'est pas une protection forte, c'est de l'hygiène.

## 48. fail2ban pour SSH

fail2ban bannit temporairement les IP qui accumulent des échecs (complément
indispensable même avec `PasswordAuthentication no` : ça coupe le bruit et les
tentatives sur les clés).

```bash
serveur# apt install -y fail2ban
```

`/etc/fail2ban/jail.d/sshd.local` :

```ini
[sshd]
enabled = true
port = 22
maxretry = 5
findtime = 10m
bantime = 1h
# Recidive : les récidivistes sont bannis plus longtemps
[recidive]
enabled = true
```

```bash
serveur# systemctl enable --now fail2ban
serveur# fail2ban-client status sshd
serveur# fail2ban-client set sshd unbanip 203.0.113.7   # débannir (toi, un jour)
```

Bonnes pratiques :

- Mets ton IP de bureau / ton bastion en `ignoreip` (dans `jail.local`) pour
  ne jamais te bannir toi-même.
- `bantime` progressif : 1h puis `recidive` 1 semaine pour les robots
  persistants.
- Sur un parc, centralise : fail2ban + journald distant, ou un WAF/règles
  firewall partagées. Un attaquant banni sur web1 doit l'être partout.
- Alternative/complément : `sshd` + nftables `recent`, ou port-knocking pour
  les serveurs très exposés.

## 49. 2FA par TOTP

Deuxième facteur via PAM (`libpam-google-authenticator`), compatible avec les
applis TOTP (FreeOTP, Aegis — préfère-les à Google Authenticator).

```bash
serveur# apt install -y libpam-google-authenticator
serveur$ google-authenticator   # en tant que l'utilisateur concerné
# -> flasher le QR code, noter les codes de secours (sur papier, au coffre)
```

`/etc/pam.d/sshd` : ajouter (ou décommenter selon la distro) :

```
auth required pam_google_authenticator.so nullok
```

`nullok` = les utilisateurs sans TOTP configuré passent quand même (phase de
transition). **Retire `nullok`** quand tout le monde est enrôlé.

`/etc/ssh/sshd_config.d/70-2fa.conf` :

```
ChallengeResponseAuthentication yes
AuthenticationMethods publickey,keyboard-interactive
```

Ici : clé **ET** TOTP obligatoires (la virgule = « et »). Sans la ligne
`AuthenticationMethods`, le TOTP serait proposé en *alternative* à la clé,
pas en *complément* — nuance critique (section 50).

## 50. Clé + 2FA combinés (`AuthenticationMethods`)

Syntaxe : la virgule = ET, l'espace = OU.

