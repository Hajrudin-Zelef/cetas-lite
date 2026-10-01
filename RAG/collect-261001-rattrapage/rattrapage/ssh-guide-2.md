---
id: collect-261001-rattrapage/rattrapage/ssh-guide-2
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [172, 397]
sha256: 889a26937a8a29db9fce18cef312e7d6a2a3d0d52710a6b97029675506065245
---

# Guide SSH approfondi

Retiens surtout : avec l'authentification par clé, **rien de secret ne transite**
(même chiffré) hormis la signature d'un défi unique. C'est structurellement
plus sûr qu'un mot de passe, même avec un bon mot de passe.

Algorithmes à éviter aujourd'hui : `3des`, `blowfish`, `arcfour`, `md5`-based
MACs, `diffie-hellman-group1-sha1`, clés RSA < 2048 bits, DSA. Voir section 45
pour la configuration moderne.

## 4. Handshake : établissement d'une session (notions)

Séquence simplifiée d'une connexion `ssh zelef@serveur` :

1. **TCP** : le client ouvre une connexion vers le port 22 du serveur.
2. **Négociation** : client et serveur échangent leurs listes d'algorithmes
   (KEX, chiffrement, MAC, compression) et choisissent le meilleur commun.
3. **Échange de clés** : Diffie-Hellman / Curve25519 produit une clé de session
   éphémère ; tout le reste de la session est chiffré et intègre.
4. **Authentification du serveur** : le serveur signe avec sa clé d'hôte ; le
   client compare l'empreinte à `known_hosts` (sections 7-8).
5. **Authentification du client** : `publickey` (signature d'un défi),
   `password` ou `keyboard-interactive` (2FA), selon la config du serveur.
6. **Ouverture des canaux** : shell, commande distante, SFTP, tunnels.

En pratique, les étapes 1 à 4 prennent une fraction de seconde en LAN ; c'est
le multiplexage (section 28) qui les rend quasi gratuites pour les connexions
suivantes vers le même hôte.

## 5. Méthodes d'authentification

| Méthode | Principe | Usage |
|---|---|---|
| `publickey` | Signature par clé privée du client | **Standard en production** |
| `password` | Mot de passe transmis dans le tunnel chiffré | À désactiver sur Internet (section 42) |
| `keyboard-interactive` | Échanges question/réponse (PAM) | Supporte le **2FA TOTP** (section 49) |
| `hostbased` | Confiance entre machines | Rare, fragile, à éviter |
| `gssapi-with-mic` | Kerberos | Environnements AD/Kerberos |

Ordre essayé par le client : configurable via `PreferredAuthentications`.
En production, vise : `publickey` obligatoire, éventuellement `+ keyboard-interactive`
pour le second facteur (section 50).

## 6. Installation d'OpenSSH

**Debian / Ubuntu :**

```bash
client$ sudo apt update && sudo apt install -y openssh-client
serveur# apt update && apt install -y openssh-server
serveur# systemctl enable --now ssh
serveur# systemctl status ssh
```

**RHEL / Rocky / AlmaLinux :**

```bash
serveur# dnf install -y openssh-server
serveur# systemctl enable --now sshd
```

**Windows 10/11** (client et serveur intégrés depuis 2018) :

```powershell
# En administrateur
Add-WindowsCapability -Online -Name OpenSSH.Client~~~~0.0.1.0
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
Start-Service sshd
Set-Service -Name sshd -StartupType Automatic
```

**macOS** : client préinstallé. Serveur : Réglages → Général → Partage →
« Connexion à distance » (ou `sudo systemsetup -setremotelogin on`).

Vérifier la version (utile pour le dépannage et les CVE) :

```bash
client$ ssh -V
OpenSSH_9.6p1 Ubuntu-3ubuntu13.2, OpenSSL 3.0.13 30 Jan 2024
```

## 7. Premier contact : `ssh`, clés d'hôte, `known_hosts`

```bash
client$ ssh zelef@192.0.2.10
The authenticity of host '192.0.2.10 (192.0.2.10)' can't be established.
ED25519 key fingerprint is SHA256:AbCdEfGhIjKlMnOpQrStUvWxYz0123456789ABc.
This key is not known by any other names.
Are you sure you want to continue connecting (yes/no/[fingerprint])?
```

Ce message apparaît **une seule fois par hôte** : c'est le modèle TOFU
(Trust On First Use). En tapant `yes`, l'empreinte est ajoutée à
`~/.ssh/known_hosts`. Aux connexions suivantes, si l'empreinte **change**,
SSH refuse de se connecter : c'est la protection contre l'homme-du-milieu.

Commandes utiles :

```bash
client$ ssh-keygen -l -f ~/.ssh/known_hosts   # lister les empreintes connues
client$ ssh-keygen -F 192.0.2.10              # chercher un hôte précis
client$ ssh-keygen -R 192.0.2.10              # retirer un hôte (après réinstall)
client$ ssh -o StrictHostKeyChecking=accept-new zelef@192.0.2.10
```

`StrictHostKeyChecking=accept-new` : ajoute automatiquement les hôtes inconnus
mais **refuse** toujours un changement d'empreinte. Bon compromis pour les
scripts (jamais `=no`, qui désactive toute vérification).

## 8. Vérifier l'empreinte d'un hôte (TOFU)

Le TOFU aveugle (`yes` sans réfléchir) est vulnérable à une attaque
homme-du-milieu **lors de la première connexion**. En production, vérifie
l'empreinte par un canal indépendant :

**Côté serveur** (à faire une fois, à conserver dans ta doc d'infra) :

```bash
serveur# for f in /etc/ssh/ssh_host_*_key.pub; do
  echo "== $f"; ssh-keygen -l -f "$f"
done
# 256 SHA256:AbCdEfGhIjKlMnOpQrStUvWxYz0123456789ABc root@web1 (ED25519)
# 3072 SHA256:ZyXwVuTsRqPoNmLkJiHgFeDcBa9876543210ZYx root@web1 (RSA)
```

**Alternatives propres au TOFU aveugle :**

- **SSHFP dans le DNS** (avec DNSSEC) : `ssh -o VerifyHostKeyDNS=yes`.
  Publie avec `ssh-keygen -r web1.example.com`.
- **Certificats SSH** (section 51) : plus de TOFU du tout, la CA signe les
  clés d'hôte.
- **Inventaire interne** : page wiki avec les empreintes de chaque serveur,
  vérifiée hors bande.

En pratique d'équipe : documente les empreintes au provisionnement
(Ansible peut les afficher), et apprends à tes collègues à comparer
**au moins les 8 premiers et 8 derniers caractères** avant de taper `yes`.

## 9. `ssh_config` client : structure

Le fichier `~/.ssh/config` (permissions `600`) évite de retaper les options.
Priorité : **la première valeur obtenue gagne** (contrairement à
`sshd_config`). Ordre de lecture : ligne de commande > `~/.ssh/config` >
`/etc/ssh/ssh_config`.

```bash
client$ ls -l ~/.ssh/config
-rw------- 1 zelef zelef 2341 Sep 26 2026 /home/zelef/.ssh/config
client$ chmod 600 ~/.ssh/config   # si besoin
```

Structure type :

```
# Bloc global : s'applique à tout (mettre les défauts personnels)
Host *
    ServerAliveInterval 60
    ServerAliveCountMax 3
    HashKnownHosts yes
    StrictHostKeyChecking accept-new

# Blocs spécifiques : le premier bloc correspondant gagne pour chaque option
Host web1
    HostName 192.0.2.10
    User zelef
    Port 2222
    IdentityFile ~/.ssh/id_ed25519_web
```

Tester la résolution effective d'un hôte (indispensable au débogage) :

```bash
client$ ssh -G web1 | grep -E '^(hostname|user|port|identityfile|proxyjump) '
hostname 192.0.2.10
user zelef
port 2222
identityfile ~/.ssh/id_ed25519_web
```

`ssh -G` affiche la configuration finale après fusion : c'est **l'outil n°1**
pour comprendre « pourquoi SSH ne prend pas ma clé / mon port ».

## 10. `Host` et alias

`Host` définit un motif ; tout ce qui suit s'applique aux hôtes correspondants.
Les motifs supportent `*` et `?`.

```
# Alias simple et mémorisable
Host web1
    HostName 192.0.2.10
    User zelef

# Motif : tous les serveurs d'un site
Host *.site-paris.example.com
    User admin
    ProxyJump bastion-paris

# Motif : tout un /24 via un bastion (pratique, à manier avec soin)
Host 10.20.30.*
    User zelef
    ProxyJump bastion

# Plusieurs noms pour le même bloc
Host db db-principal
    HostName 192.0.2.50
    User postgres-admin
```

Règles d'or :

- Des alias **courts et stables** (`web1`, `db`, `bastion`) plutôt que des IP
  dans les scripts et la doc d'équipe.
- Le bloc `Host *` **toujours en dernier** : comme la première valeur gagne,
  un `Host *` placé en premier écraserait les blocs spécifiques.
- `ssh web1` utilise l'alias ; `ssh 192.0.2.10` utilise l'IP brute (et donc
  potentiellement d'autres options — attention aux surprises avec les motifs
  en `*`).

## 11. `HostName`, `User`, `Port`

```
Host web1
    HostName 192.0.2.10        # le vrai nom/IP ; l'alias reste "web1"
    User zelef                 # fini les zelef@ à chaque fois
    Port 2222                  # si sshd n'écoute pas sur 22
```

