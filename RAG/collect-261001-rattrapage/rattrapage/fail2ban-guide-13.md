---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-13
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["arr", "valuation"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [2928, 3179]
sha256: 05a2dea919287b443c90c43b9853d7c549b6637965b72ad25edc458fba84b60e
---

# Guide fail2ban — Le bouclier anti-brute-force

### Enrichir les logs avec le pays

```bash
sudo apt install -y geoip-bin geoip-database
geoiplookup 203.0.113.45
# GeoIP Country Edition: US, United States
```

Script d'enrichissement du rapport quotidien (section 48) : ajoutez une colonne pays au top 10.

### Bannir un pays : la réalité technique

fail2ban ne fait pas de GeoIP natif. Les approches :

1. **Listes d'IP par pays** (ex: ipdeny.com) chargées dans un ipset/nftables set, en amont de fail2ban. C'est du firewall statique, pas du fail2ban.
2. **Action personnalisée** qui consulte GeoIP au moment du ban et allonge le bantime si le pays est « inattendu ».

Exemple d'action « bantime selon pays » (idée, à adapter) :

```ini
# /etc/fail2ban/action.d/geo-ban.conf
[Definition]
actionban = /usr/local/bin/geo-ban.sh <ip> <name>
actionunban = /usr/local/bin/geo-unban.sh <ip>
```

> Recommandation pragmatique : **n'auto-bannissez pas par pays**. Utilisez la GeoIP pour le *reporting* (d'où viennent les attaques ?) et gardez les décisions de ban sur les comportements (échecs d'authentification).

---

## 51. Le fichier filter.d : anatomie complète

```ini
# /etc/fail2ban/filter.d/exemple.conf

[INCLUDES]
# Charge d'abord common.conf (définit <HOST>, __prefix_line...)
# puis les fichiers listés, dans l'ordre
before = common.conf

[Definition]
# La/les regex de détection (une par ligne, toutes en OU)
failregex = ^%(__prefix_line)sFailed password for <F-USER>\S+</F-USER> from <HOST> port \d+ ssh2$
            ^%(__prefix_line)sInvalid user \S+ from <HOST> port \d+$

# Regex d'exclusion : une ligne qui matche ignoreregex n'est JAMAIS comptée,
# même si elle matche failregex (pratique pour les faux positifs connus)
ignoreregex = ^%(__prefix_line)sFailed password for deploy from <HOST> port \d+ ssh2$

# Surcharges possibles :
# datepattern = ^%Y-%m-%d %H:%M:%S   (si le format de date n'est pas auto-détecté)
# journalmatch = _SYSTEMD_UNIT=sshd.service  (pour le backend systemd)
```

### Ordre d'évaluation pour chaque ligne de log

1. La date est-elle reconnue ? (non → ligne ignorée silencieusement)
2. `ignoreregex` matche-t-elle ? (oui → ignorée)
3. Une `failregex` matche-t-elle ? (oui → échec compté pour `<HOST>`)

### Héritage et surcharge d'un filter existant

Ne modifiez jamais un filter fourni : créez `filter.d/sshd.local` :

```ini
# /etc/fail2ban/filter.d/sshd.local — AJOUTE des regex au filter sshd
[Definition]
failregex = %(known/failregex)s
            ^%(__prefix_line)sMon pattern maison from <HOST>$
```

`%(known/failregex)s` reprend les regex du fichier de base. Votre `.local` les complète.

---

## 52. Le fichier action.d : anatomie complète

```ini
# /etc/fail2ban/action.d/exemple.conf

[Definition]
# Exécuté une fois au démarrage du jail
actionstart = iptables -N f2b-<name>
              iptables -A f2b-<name> -j RETURN
              iptables -I INPUT -p tcp --dport <port> -j f2b-<name>

# Exécuté à chaque ban
actionban = iptables -I f2b-<name> 1 -s <ip> -j <blocktype>

# Exécuté à chaque unban (fin de bantime ou unban manuel)
actionunban = iptables -D f2b-<name> -s <ip> -j <blocktype>

# Exécuté à l'arrêt du jail (doit ANNULER actionstart)
actionstop = iptables -D INPUT -p tcp --dport <port> -j f2b-<name>
             iptables -F f2b-<name>
             iptables -X f2b-<name>

# Vérification optionnelle avant ban (retourne 0 si déjà banni)
actioncheck = iptables -n -L f2b-<name> | grep -q 'f2b-<name>'

[Init]
# Valeurs par défaut des variables, surchargeables dans le jail
port = ssh
blocktype = REJECT --reject-with icmp-port-unreachable
```

### Variables disponibles dans les actions

| Variable | Contenu |
|---|---|
| `<ip>` | l'IP bannie/débannie |
| `<name>` | le nom du jail |
| `<port>` | le(s) port(s) du jail |
| `<failures>` | nombre d'échecs ayant déclenché le ban |
| `<time>` | timestamp du ban |
| `<matches>` | extrait des lignes de log fautives |

### Règle d'or des actions maison

`actionstop` doit **exactement annuler** `actionstart`, et `actionunban` doit **exactement annuler** `actionban`. Sinon : règles orphelines qui s'accumulent à chaque reload.

---

## 53. Jail.d : surcharges modulaires

Plutôt qu'un seul gros `jail.local`, découpez par service dans `/etc/fail2ban/jail.d/`. Les fichiers `.local` sont lus après `jail.local`.

```
# /etc/fail2ban/jail.d/00-defaults.local
[DEFAULT]
bantime  = 1h
findtime = 10m
maxretry = 5
ignoreip = 127.0.0.1/8 ::1 192.168.10.0/24
banaction = nftables-multiport

# /etc/fail2ban/jail.d/10-sshd.local
[sshd]
enabled = true

# /etc/fail2ban/jail.d/20-web.local
[nginx-http-auth]
enabled = true
...

# /etc/fail2ban/jail.d/30-mail.local
[postfix]
enabled = true
...

# /etc/fail2ban/jail.d/99-recidive.local
[recidive]
enabled = true
...
```

Avantages : chaque fichier a un responsable, les diffs git sont lisibles, Ansible déploie par rôle.

---

## 54. Ordre de fusion de la configuration

fail2ban lit les fichiers dans cet ordre strict. **La dernière valeur définie gagne.**

```
1. jail.conf
2. jail.d/*.conf          (ordre alphabétique)
3. jail.local
4. jail.d/*.local         (ordre alphabétique)
```

### Voir la configuration effective d'un jail

```bash
# fail2ban >= 1.0 : affiche la config fusionnée
sudo fail2ban-client --dp "jail:sshd:maxretry"
sudo fail2ban-client -d | grep -A15 "^\[sshd\]"
```

`fail2ban-client -d` (dump) affiche **toute** la configuration effective : c'est la vérité de référence quand on ne comprend plus quelle valeur s'applique.

### Piège classique

```ini
# jail.local
[DEFAULT]
maxretry = 5

# jail.d/99-strict.local
[sshd]
maxretry = 3    # ← gagne sur le 5 du DEFAULT pour le jail sshd
```

Si vous ne comprenez pas d'où vient une valeur, `fail2ban-client -d` + `grep`.

---

## 55. Durcissement de fail2ban lui-même

fail2ban tourne en **root** (il doit modifier le firewall). Quelques mesures :

### Réduire la surface

```ini
# /etc/fail2ban/fail2ban.conf — via fail2ban.d/
[Definition]
# Socket accessible uniquement à root
socket = /run/fail2ban/fail2ban.sock
# (permissions gérées par systemd, déjà restrictives par défaut)
```

Vérifiez :

```bash
ls -la /run/fail2ban/fail2ban.sock
# srwx------ 1 root root ... → seul root peut piloter le démon. Parfait.
```

### Protéger les fichiers de config

```bash
sudo chmod 640 /etc/fail2ban/action.d/webhook.conf   # contient un token
sudo chmod 640 /etc/fail2ban/action.d/cloudflare.local
sudo chown root:root /etc/fail2ban/jail.local
```

### Surveiller l'intégrité

- Mettez `/etc/fail2ban/` sous git (section 37) : toute modification non planifiée apparaît dans `git status`.
- Avec AIDE ou auditd, surveillez les écritures dans `/etc/fail2ban/`.

### fail2ban n'écoute pas sur le réseau

Le démon ne bind **aucun port** : il communique via socket Unix local. Il n'est pas exposable directement. Le risque vient des **actions** (curl vers des webhooks, tokens en clair) : soignez-les.

---

## 56. Bannir un pays entier : réaliste ?

### La réponse courte

Non, pas avec fail2ban seul. Et c'est rarement une bonne idée.

### Pourquoi c'est tentant (et pourquoi ça échoue)

- Tentant : « 80 % de mes attaques viennent de 3 pays où je n'ai aucun client ».
- Problème 1 : les attaquants louent des VPS **dans votre pays** (OVH, Hetzner...). Le géoblocage ne les arrête pas.
- Problème 2 : les plages d'IP par pays changent en permanence ; une liste statique se périme.
- Problème 3 : un client en déplacement à l'étranger (ou derrière un VPN) se fait bloquer → ticket au support.

### Si vous voulez quand même le faire (firewall, pas fail2ban)

