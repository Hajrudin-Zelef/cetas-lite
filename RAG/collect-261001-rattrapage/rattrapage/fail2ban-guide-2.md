---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-2
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [183, 423]
sha256: 47992d1ad774221313beb8f26b426bbf028097ab0ecac268212eebce208e6ada
---

# Guide fail2ban — Le bouclier anti-brute-force

- `<HOST>` est un **raccourci magique** : il capture l'adresse IP (IPv4 ou IPv6).
- `<F-USER>...</F-USER>` capture le nom d'utilisateur (utile pour les logs).
- `__prefix_line` vient de `common.conf` : il matche le préfixe date/hostname/processus.

> Règle d'or : **un filter ne décide jamais de bannir**. Il se contente de dire « cette ligne est un échec d'authentification, venant de cette IP ».

### Le jail : « sur quoi veiller et combien tolérer »

Un jail relie un filter à des logs et des seuils. C'est la **configuration opérationnelle**.

```ini
[sshd]
enabled  = true
filter   = sshd
logpath  = /var/log/auth.log
backend  = auto
port     = ssh
maxretry = 5
findtime = 10m
bantime  = 1h
action   = %(action_)s
```

Lecture : « surveille `/var/log/auth.log` avec le filter sshd ; si une IP échoue 5 fois en 10 minutes, bannis-la 1 heure sur le port ssh ».

### L'action : « que faire quand ça dépasse »

Une action est un fichier dans `/etc/fail2ban/action.d/<nom>.conf`. Elle définit les commandes à exécuter au **ban** et au **unban**.

Structure type (`action.d/iptables.conf`, simplifié) :

```ini
[Definition]
actionstart  = iptables -N f2b-sshd
               iptables -A f2b-sshd -j RETURN
               iptables -I INPUT -p tcp --dport ssh -j f2b-sshd

actionban    = iptables -I f2b-sshd 1 -s <ip> -j REJECT --reject-with icmp-port-unreachable

actionunban  = iptables -D f2b-sshd -s <ip> -j REJECT --reject-with icmp-port-unreachable

actionstop   = iptables -D INPUT -p tcp --dport ssh -j f2b-sshd
               iptables -F f2b-sshd
               iptables -X f2b-sshd
```

`<ip>` est remplacé par l'IP fautive. `actionstart` crée la chaîne, `actionstop` la nettoie.

### Schéma d'ensemble

```
/etc/fail2ban/
├── jail.conf / jail.local      ← les JAILS (quoi surveiller, seuils)
├── filter.d/*.conf              ← les FILTERS (regex de détection)
└── action.d/*.conf              ← les ACTIONS (ban/unban)

            ┌─────────────────────────────────┐
            │         fail2ban-server         │
            │                                 │
 log ──▶    │  jail[sshd]                     │
fichier     │   filter: sshd.conf             │
ou journal  │   seuils: 5 / 10m / 1h          │──▶ action ban/unban
            │                                 │     (iptables, nftables,
            └─────────────────────────────────┘      ufw, cloudflare...)
                    ▲
                    │ fail2ban-client (CLI)
                    │ get / set / ban / unban / status
```

### Le démon et le client

- `fail2ban-server` : le démon Python qui tourne en tâche de fond (surveille les logs, applique les bannissements).
- `fail2ban-client` : l'outil en ligne de commande qui dialogue avec le serveur via un socket (`/run/fail2ban/fail2ban.sock`).
- `fail2ban-regex` : l'outil de **test** des regex contre un fichier de log (indispensable, section 16).

---

## 3. Installation sur Debian/Ubuntu

### Installation standard

```bash
sudo apt update
sudo apt install -y fail2ban
```

C'est tout. Le paquet Debian/Ubuntu installe :

- le démon et les outils CLI,
- un service systemd `fail2ban.service` (activé au boot),
- des jails par défaut raisonnables (dont `sshd` activé),
- les dépendances Python nécessaires.

Vérification :

```bash
fail2ban-client --version
# Fail2Ban v1.0.2

sudo systemctl status fail2ban
# ● fail2ban.service - Fail2Ban Service
#   Active: active (running) ...

sudo fail2ban-client status
# Status
# |- Number of jail:	1
# `- Jail list:	sshd
```

### Versions par distribution (repères 2026)

| Distribution | Version fail2ban (dépôt) | Backend nftables natif |
|---|---|---|
| Debian 11 (bullseye) | 0.11.2 | partiel (banaction nftables-multiport OK) |
| Debian 12 (bookworm) | 1.0.2 | oui |
| Debian 13 (trixie) | 1.1.x | oui |
| Ubuntu 20.04 | 0.11.1 | partiel |
| Ubuntu 22.04 | 0.11.2 | partiel |
| Ubuntu 24.04 | 1.0.2 | oui |

> Sur Debian 11 / Ubuntu ≤ 22.04, `iptables` est en réalité un wrapper vers nftables (`iptables-nft`). fail2ban fonctionne quand même, mais préférez explicitement `banaction = nftables-multiport` sur les systèmes récents (section 19).

### Installer une version plus récente (optionnel)

Si vous voulez les derniers filters (WordPress, nouveaux services), deux options :

```bash
# Option A : backports Debian (préférée, propre)
sudo apt -t bookworm-backports install fail2ban

# Option B : depuis les sources GitHub (à éviter en prod sauf besoin précis)
# git clone https://github.com/fail2ban/fail2ban.git
```

> En production, **restez sur le paquet de la distribution** sauf raison impérieuse : la version packagée est testée avec le backend firewall du système.

### Dépendances utiles

```bash
# whois : pour l'action d'alerte mail enrichie (optionnel)
sudo apt install -y whois

# ipset : pour la variante haute performance (optionnel, section 49)
sudo apt install -y ipset
```

### Vérifier que le service démarre au boot

```bash
sudo systemctl enable --now fail2ban
sudo systemctl is-enabled fail2ban
# enabled
```

---

## 4. Arborescence et fichiers

```
/etc/fail2ban/
├── fail2ban.conf          # config du démon (socket, loglevel, dbfile) — rarement touché
├── fail2ban.d/            # surcharges modulaires de fail2ban.conf
├── jail.conf              # JAILS par défaut — NE JAMAIS MODIFIER
├── jail.d/                # vos surcharges de jails (fichiers .local ou .conf)
│   └── defaults-debian.conf
├── jail.local             # VOTRE fichier principal (à créer)
├── filter.d/              # ~100 filters fournis
│   ├── sshd.conf
│   ├── nginx-http-auth.conf
│   └── ...
└── action.d/              # ~50 actions fournies
    ├── iptables.conf
    ├── nftables.conf
    ├── ufw.conf
    └── ...
```

### La règle d'or : ne jamais toucher `jail.conf`

`jail.conf` est **écrasé à chaque mise à jour du paquet**. Toute votre configuration va dans :

1. **`/etc/fail2ban/jail.local`** — le fichier principal (recommandé), ou
2. **`/etc/fail2ban/jail.d/*.local`** — découpé par service (pratique avec Ansible).

fail2ban fusionne `jail.conf` + `jail.d/*.conf` + `jail.local` + `jail.d/*.local`, dans cet ordre ; **la dernière valeur définie gagne**.

> Astuce : mettez `jail.local` sous version (git) ou sauvegardez-le (section 37). C'est LE fichier à ne pas perdre.

### Où sont les logs de fail2ban lui-même

- `/var/log/fail2ban.log` — le journal du démon (bans, unbans, erreurs).
- `journalctl -u fail2ban` — la même chose via systemd (selon config).

### Où est la base de données

`/var/lib/fail2ban/fail2ban.sqlite3` — elle mémorise les bannissements pour les restaurer après un redémarrage (section 29).

---

## 5. Premier jail.local minimal

Voici le plus petit `jail.local` utile en production. Il protège SSH et ne touche à rien d'autre.

```ini
# /etc/fail2ban/jail.local — configuration minimale de production

[DEFAULT]
# Une IP bannie le reste 1 heure
bantime = 1h

# Fenêtre de comptage : 10 minutes
findtime = 10m

# Nombre d'échecs tolérés dans la fenêtre
maxretry = 5

# NE JAMAIS bannir ces adresses (à adapter !)
ignoreip = 127.0.0.1/8 ::1 192.168.1.0/24 10.0.0.0/8

# Action par défaut : bannissement nftables multi-ports
banaction = nftables-multiport

[sshd]
enabled = true
```

Mise en service :

```bash
# 1. Vérifier la syntaxe (fail2ban >= 0.10)
sudo fail2ban-client --test

# 2. Recharger
sudo fail2ban-client reload

# 3. Contrôler
sudo fail2ban-client status sshd
```

Sortie attendue :

