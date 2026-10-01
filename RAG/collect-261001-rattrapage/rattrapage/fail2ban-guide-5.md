---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-5
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [959, 1210]
sha256: 98b011bf528cc56430232282252cc52e0a8276df4fa3e97451c87178863d89f4
---

# Guide fail2ban — Le bouclier anti-brute-force

Lines: 2450 lines, 0 ignored, 27 matched, 2423 missed
```

Lecture : sur 2450 lignes, 27 matchent le filter. Si vous savez qu'il y a eu ~27 attaques, c'est bon.

### Tester une regex ad hoc (sans fichier)

```bash
sudo fail2ban-regex /var/log/nginx/error.log \
  'user "\S+" was not found in "[^"]+", client: <HOST>'
```

### Tester contre une ligne unique (debug rapide)

```bash
echo 'Sep 26 10:22:11 srv sshd[1234]: Failed password for root from 203.0.113.45 port 51234 ssh2' \
  | sudo fail2ban-regex - /etc/fail2ban/filter.d/sshd.conf --print-all-matched
```

L'option `--print-all-matched` affiche les lignes matchées : idéal pour comprendre pourquoi une ligne matche (ou pas).

### Les options indispensables

| Option | Usage |
|---|---|
| `--print-all-matched` | affiche les lignes qui matchent |
| `--print-all-missed` | affiche les lignes qui NE matchent pas |
| `--print-all-ignored` | affiche les lignes ignorées (ignoreregex) |
| `-v` | mode verbeux (détail par regex) |
| `--datepattern=...` | forcer un format de date |

### Cas d'école : ma regex ne matche pas

```bash
# 1. Vérifier que le format de DATE est reconnu
sudo fail2ban-regex /var/log/monapp.log /etc/fail2ban/filter.d/monapp.conf -v
# Si "Date template hits" = 0 → le problème vient de la date, pas de la regex !

# 2. Forcer le format de date
sudo fail2ban-regex --datepattern="%%Y-%%m-%%d %%H:%%M:%%S" /var/log/monapp.log ...
```

> 50 % des « ma regex ne marche pas » sont en réalité des **problèmes de date** : fail2ban doit d'abord extraire le timestamp avant d'appliquer la regex (section 46).

---

## 17. Créer un filter personnalisé pas à pas

Scénario : vous avez une application maison qui loggue dans `/var/log/myapp/auth.log` :

```
2026-09-26 14:32:10 [WARN] login failed for user 'jdupont' from 203.0.113.77 (bad password)
2026-09-26 14:32:15 [WARN] login failed for user 'jdupont' from 203.0.113.77 (bad password)
```

### Étape 1 : écrire la regex et la tester

La partie variable : le nom d'utilisateur et l'IP. La regex :

```
\[WARN\] login failed for user '[^']+' from <HOST> \(bad password\)
```

Test immédiat :

```bash
sudo fail2ban-regex /var/log/myapp/auth.log \
  "\[WARN\] login failed for user '[^']+' from <HOST> \(bad password\)" \
  --print-all-matched
```

### Étape 2 : créer le fichier filter

```ini
# /etc/fail2ban/filter.d/myapp.conf
[INCLUDES]
before = common.conf

[Definition]
failregex = \[WARN\] login failed for user '[^']+' from <HOST> \(bad password\)$
ignoreregex =
```

> `common.conf` fournit `__prefix_line` et les définitions de `<HOST>`. L'inclure est quasi obligatoire.

### Étape 3 : tester le fichier complet

```bash
sudo fail2ban-regex /var/log/myapp/auth.log /etc/fail2ban/filter.d/myapp.conf
# Lines: 120 lines, 0 ignored, 8 matched, 112 missed  → OK si 8 attaques connues
```

### Étape 4 : créer le jail

```ini
# /etc/fail2ban/jail.d/myapp.local
[myapp]
enabled  = true
port     = 8443
filter   = myapp
logpath  = /var/log/myapp/auth.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

### Étape 5 : recharger et vérifier

```bash
sudo fail2ban-client reload
sudo fail2ban-client status myapp
```

### Checklist du filter maison

- [ ] La regex matche les vraies lignes d'échec (`--print-all-matched`)
- [ ] La regex ne matche PAS les lignes de succès (`--print-all-missed` relu à la main)
- [ ] `<HOST>` capture bien l'IP de l'attaquant (pas celle d'un proxy)
- [ ] Le format de date est détecté (section 16)
- [ ] Le fichier est dans `filter.d/`, pas écrasé par les mises à jour (c'est votre fichier, il est safe)
- [ ] Le jail pointe vers le bon `logpath`

---

## 18. Regex : guide de survie

Pas besoin d'être un expert regex. Voici les 20 % qui couvrent 80 % des filters.

### Les briques de base

| Motif | Signification | Exemple |
|---|---|---|
| `<HOST>` | une adresse IP (v4 ou v6) | `from <HOST>` |
| `\d+` | un ou plusieurs chiffres | `port \d+` |
| `\S+` | caractères non-espaces | `user "\S+"` |
| `[^']+` | tout sauf une apostrophe | `'[^']+'` |
| `.*` | n'importe quoi (gourmand) | à éviter si possible |
| `(?:...)` | groupe non-capturant | `(?:password\|publickey)` |
| `^` / `$` | début / fin de ligne | `^...Failed...$` |
| `\.` | un point littéral | `203\.0\.113\.` |
| `\s*` | espaces éventuels | `ssh2\s*$` |

### Les raccourcis fail2ban

Définis dans `filter.d/common.conf` :

| Raccourci | Équivalent | Usage |
|---|---|---|
| `<HOST>` | `(?:::f{4,6}:)?(?P<host>[\w\-.^_]*\w)` | capture l'IP |
| `<F-USER>...</F-USER>` | capture nommée `user` | le login attaqué |
| `<F-MLFID>...</F-MLFID>` | capture multi-lignes | logs sur plusieurs lignes |
| `__prefix_line` | `^...date...host...process:` | préfixe syslog standard |

### Les 5 règles d'un bon failregex

1. **Ancrez avec `<HOST>`** : toute failregex doit contenir `<HOST>`, sinon fail2ban ne sait pas qui bannir (et refuse de charger le filter).
2. **Soyez spécifique** : `Failed` seul matchera trop large. `Failed password for` est précis.
3. **Échappez les caractères spéciaux** : `[`, `]`, `(`, `)`, `.` ont un sens en regex. `[WARN]` doit s'écrire `\[WARN\]`.
4. **Préférez `[^']+` à `.*`** : moins gourmand, moins de faux positifs.
5. **Testez les lignes de succès** : assurez-vous que « login successful » ne matche pas votre regex d'échec.

### Exemple décortiqué

```
^\s*\S+ sshd\[\d+\]: Failed (?:password|publickey) for (?:invalid user )?\S+ from <HOST>( port \d+)?(?: ssh2)?\s*$
```

| Morceau | Rôle |
|---|---|
| `^\s*` | début de ligne, espaces éventuels |
| `\S+ sshd\[\d+\]:` | hostname + `sshd[pid]:` |
| `Failed (?:password\|publickey)` | le type d'échec (sans capturer) |
| `for (?:invalid user )?\S+` | `for root` ou `for invalid user admin` |
| `from <HOST>` | **l'IP de l'attaquant** |
| `( port \d+)?` | port source, optionnel |
| `(?: ssh2)?\s*$` | suffixe `ssh2` optionnel, fin de ligne |

### Tester mentalement : la méthode en 3 questions

1. Si je remplace `<HOST>` par `203.0.113.45`, ma regex matche-t-elle la vraie ligne de log ? (copiez-collez, ne devinez pas)
2. Matche-t-elle aussi une ligne de **succès** ? (si oui, resserrez)
3. Le format de date en début de ligne est-il dans la liste des `Date template hits` ?

---

## 19. Les actions : iptables, nftables, ufw

L'action transforme une décision (« bannir 203.0.113.45 ») en règles firewall. Le choix dépend de votre système.

### Tableau de décision

| Système | Pare-feu actif | banaction recommandée |
|---|---|---|
| Debian 12/13, Ubuntu 24.04 | nftables | `nftables-multiport` |
| Debian 11, Ubuntu 22.04 | iptables-nft (wrapper) | `iptables-multiport` (fonctionne) ou `nftables-multiport` |
| Ubuntu avec `ufw enable` | ufw | `ufw` |
| Haute performance (gros volume) | ipset/iptables | `iptables-ipset-proto6` |

### Configurer l'action par défaut

```ini
# /etc/fail2ban/jail.local
[DEFAULT]
# Choix moderne (Debian 12+, Ubuntu 24.04+)
banaction = nftables-multiport
banaction_allports = nftables-allports
```

```ini
# Variante classique iptables
[DEFAULT]
banaction = iptables-multiport
banaction_allports = iptables-allports
```

```ini
# Variante UFW (si ufw est votre firewall)
[DEFAULT]
banaction = ufw
banaction_allports = ufw
```

### Ce que fait nftables-multiport

À l'activation du jail (`actionstart`) :

```nft
# (simplifié)
table inet f2b-table {
  chain f2b-sshd {
    # les IP bannies sont ajoutées ici, en tête de chaîne
  }
  chain input {
    type filter hook input priority 0;
    tcp dport == ssh jump f2b-sshd
  }
}
```

Au ban (`actionban`) : `nft add element inet f2b-table f2b-sshd { 203.0.113.45 }`
Au unban : `nft delete element ...`

Vérifier à la main :

```bash
sudo nft list table inet f2b-table
```

### Ce que fait iptables-multiport

Crée une chaîne `f2b-sshd`, y insère les IP en `REJECT`, et branche la chaîne sur `INPUT` pour les ports concernés.

