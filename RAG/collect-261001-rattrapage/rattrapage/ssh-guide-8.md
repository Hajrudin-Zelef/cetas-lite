---
id: collect-261001-rattrapage/rattrapage/ssh-guide-8
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [1516, 1717]
sha256: 991f2fef4d0b63c0bfd190d974cc15af7cb72d000f67fdaa4521394c845623b9
---

# Guide SSH approfondi

```
# Clé ET (mot de passe OU TOTP) — souple
AuthenticationMethods publickey,keyboard-interactive password,keyboard-interactive

# Clé ET TOTP — strict (recommandé pour l'admin)
AuthenticationMethods publickey,keyboard-interactive

# Par population, avec Match (en fin de fichier !)
Match Group admins
    AuthenticationMethods publickey,keyboard-interactive
Match Group deploy
    AuthenticationMethods publickey
```

Points d'attention :

- Teste **toujours** depuis une session de secours (section 41) : une
  `AuthenticationMethods` mal écrite = porte claquée.
- `keyboard-interactive` couvre le TOTP via PAM **et** les questions PAM
  classiques : vérifie qu'aucun module PAM n'accepte un simple mot de passe
  en fallback (relis `/etc/pam.d/sshd` en entier).
- Ansible et scripts non interactifs **ne passent pas** le TOTP : prévois un
  groupe/une clé de service en `publickey` seul, restreinte (`command=`,
  `from=`, sections 23-24), ou des certificats à courte durée de vie
  (section 51).

## 51. Certificats SSH : tour d'horizon

Les certificats SSH remplacent la gestion manuelle d'`authorized_keys` : une
**autorité de certification (CA)** interne signe les clés publiques des
utilisateurs (et des hôtes). Le serveur fait confiance à la CA, pas à chaque
clé individuelle.

Avantages décisifs en équipe / parc :

- **Zéro `authorized_keys` à déployer** : ajouter un arrivant = signer sa clé,
  pas toucher 200 serveurs.
- **Expiration intégrée** : `-V -1w:+52w` (valide 1 semaine avant → 52 après,
  typiquement courte : quelques heures/jours pour les accès sensibles).
- **Principals** : le certificat dit *qui* tu es (`zelef`, `deploy`) ; le
  serveur mappe les principals aux comptes autorisés.
- **Révocation** centralisée via `RevokedKeys` (section 53).
- Fini le TOFU : les clés d'hôte sont elles aussi signées par la CA.

Inconvénients : il faut opérer une CA (sécuriser sa clé privée **comme un
secret critique**, idéalement HSM ou machine dédiée offline), et outiller la
signature (petit service interne, HashiCorp Vault SSH, ou script signé).

Quand l'adopter : à partir d'une dizaine de serveurs et/ou d'arrivées-départs
fréquents. En dessous, `authorized_keys` déployé par Ansible reste très bien.

## 52. Certificats : créer une CA et signer des clés

**Sur la machine CA (dédiée, chiffrée, sauvegardée)** :

```bash
ca$ ssh-keygen -t ed25519 -f /etc/ssh/ca_user -C "CA utilisateurs $(date +%Y)"
ca$ ssh-keygen -t ed25519 -f /etc/ssh/ca_host -C "CA hotes $(date +%Y)"
# Protéger : chmod 600, passphrase forte, sauvegarde chiffrée hors site
```

**Signer la clé d'un utilisateur** (l'utilisateur envoie sa `.pub`, jamais sa
privée) :

```bash
ca$ ssh-keygen -s /etc/ssh/ca_user -I "zelef-20260926" -n zelef,deploy \
    -V -5m:+8h -z 42 /tmp/id_ed25519_zelef.pub
# -I : identifiant (audit), -n : principals, -V : validité, -z : n° série
# Produit : /tmp/id_ed25519_zelef-cert.pub -> à renvoyer à l'utilisateur
```

Côté client, le certificat se place à côté de la clé et est utilisé
automatiquement :

```bash
client$ ls ~/.ssh/
id_ed25519  id_ed25519-cert.pub  id_ed25519.pub
client$ ssh -i ~/.ssh/id_ed25519 web1   # le -cert.pub est trouvé seul
client$ ssh-keygen -L -f ~/.ssh/id_ed25519-cert.pub   # inspecter le certificat
```

**Signer une clé d'hôte** (fini le TOFU) :

```bash
ca$ ssh-keygen -s /etc/ssh/ca_host -I "web1-20260926" -h -n web1.example.com \
    -V -1w:+52w /tmp/ssh_host_ed25519_key.pub
# -h : certificat d'HÔTE, -n : noms d'hôte couverts
```

## 53. Certificats côté serveur : confiance et révocation

`/etc/ssh/sshd_config.d/80-certs.conf` :

```
# Faire confiance à la CA pour les utilisateurs...
TrustedUserCAKeys /etc/ssh/ca_user.pub
# ...et pour les hôtes (les clients vérifient via KnownHostsCommand ou config)
TrustedCAKeys /etc/ssh/ca_host.pub
HostCertificate /etc/ssh/ssh_host_ed25519_key-cert.pub
# Révocation centralisée
RevokedKeys /etc/ssh/revoked_keys
```

Mapper les principals aux comptes locaux si besoin (`AuthorizedPrincipalsFile`) :

```
AuthorizedPrincipalsFile /etc/ssh/auth_principals/%u
# /etc/ssh/auth_principals/zelef contient : zelef
# /etc/ssh/auth_principals/deploy contient : deploy ci-deploys
```

**Révocation** (départ d'un salarié, clé compromise) :

```bash
serveur# ssh-keygen -k -f /etc/ssh/revoked_keys /tmp/cert-a-revoquer-cert.pub
serveur# systemctl reload ssh
```

Exploitation au quotidien : un petit script/CI signe les clés (`-V :+8h`),
les utilisateurs récupèrent leur certificat le matin (ou via Vault), aucun
`authorized_keys` à maintenir. Le départ d'un collaborateur = ne plus signer
+ ajouter le dernier certificat à `revoked_keys`. C'est le niveau « équipe
structurée » de la gestion SSH.

---

# PARTIE SERVEUR — EXPLOITATION

## 54. Supervision : les logs d'authentification

Sur Debian/Ubuntu : `/var/log/auth.log`. Sur RHEL : `/var/log/secure`.
Via systemd partout : `journalctl`.

```bash
serveur# tail -f /var/log/auth.log
serveur# journalctl -u ssh -f
# Connexions réussies :
serveur# grep "Accepted" /var/log/auth.log | tail
# Échecs :
serveur# grep "Failed" /var/log/auth.log | tail
# Déconnectés / timeouts :
serveur# grep -E "Disconnected|timeout" /var/log/auth.log | tail
```

Avec `LogLevel VERBOSE` (recommandé en production), chaque connexion réussie
journalise **l'empreinte de la clé utilisée** :

```
Accepted publickey for zelef from 192.0.2.25 port 51234 ssh2:
ED25519 SHA256:AbCdEfGhIjKlMnOpQrStUvWxYz0123456789ABc
```

C'est précieux : en cas d'incident, tu sais **quelle clé** a été utilisée, pas
seulement quel compte. Conserve ces logs (forward vers un SIEM/syslog
central : une machine compromise ne doit pas pouvoir effacer ses traces
locales).

Requêtes utiles :

```bash
# Top des IP en échec sur 24h
serveur# grep "Failed password" /var/log/auth.log | grep -oE '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+' | sort | uniq -c | sort -rn | head
# Connexions root tentées (doivent être à zéro si PermitRootLogin no)
serveur# grep "Failed.*root" /var/log/auth.log | wc -l
# Tunnels ouverts (avec LogLevel VERBOSE)
serveur# grep -i "forward" /var/log/auth.log | tail
```

## 55. Audit des connexions : `last`, `who`, `ss`

```bash
serveur# last -n 20            # historique des connexions (wtmp)
serveur# lastb -n 20           # tentatives ÉCHOUÉES (btmp) — à surveiller
serveur# who / w               # qui est connecté MAINTENANT
serveur# ss -tnp | grep :22    # connexions TCP actuelles sur SSH
serveur# ps aux | grep sshd    # processus sshd par session
```

Lecture d'une ligne `sshd` : `sshd: zelef@pts/0` = session interactive de
zelef ; `sshd: zelef [net]` = phase réseau (pas encore authentifié) ; trop de
`[net]` persistants = scan ou flood en cours (croiser avec `MaxStartups`,
section 44, et fail2ban, section 48).

Rituel hebdo (5 minutes) sur les serveurs exposés :

1. `lastb | head` : des IP inconnues en masse ?
2. `grep Accepted /var/log/auth.log` : des comptes/heures inhabituels ?
3. `fail2ban-client status sshd` : le bannissement tourne-t-il ?
4. Clés `authorized_keys` : cf. inventaire section 59 (mensuel).

## 56. Alertes sur connexions suspectes

Exemples de signaux à alerter (SIEM, ou simple script + mail) :

- Connexion réussie **root** alors que `PermitRootLogin no` (config dérivées ?).
- Connexion depuis un pays/ASN inhabituel pour l'équipe.
- Connexion à 3h du matin un dimanche (selon tes horaires).
- `Accepted` pour un compte de service normalement non interactif.
- Rafale de `Failed password` depuis une IP interne (mouvement latéral ?).

Exemple minimaliste (cron toutes les 5 min, à adapter) :

