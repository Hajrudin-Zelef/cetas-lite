---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-16
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [3715, 3863]
sha256: 715195c772031d81ddd64fd15245facfc80cd28208e466cc7755034675e200e7
---

# Guide fail2ban — Le bouclier anti-brute-force

[Definition]
failregex = ^\[[^]]+\]: <\S+> IMAP Error: Login failed for \S+ from <HOST>\.
datepattern = ^\[%%d-%%b-%%Y %%H:%%M:%%S %%z\]
```

> Le `datepattern` explicite est nécessaire : le format `[26-Sep-2026 13:00:01 +0200]` n'est pas toujours auto-détecté. Testez avec `fail2ban-regex -v`.

### Jail (seuils « humains »)

```ini
[roundcube]
enabled  = true
port     = http,https
filter   = roundcube
logpath  = /var/log/roundcube/errors.log
maxretry = 10
findtime = 10m
bantime  = 30m
```

`maxretry = 10` : un utilisateur qui se trompe 5 fois + son téléphone qui réessaie 5 fois ne doit pas être banni. `bantime = 30m` : le temps d'appeler le support, pas une punition.

### Alternative : protéger dovecot plutôt que Roundcube

Si Roundcube s'authentifie en IMAP contre dovecot en local, les échecs apparaissent AUSSI dans les logs dovecot — mais avec l'IP du serveur webmail, pas celle du client ! Dans ce cas, le jail dovecot est inutile (il bannirait le webmail). Protégez au niveau Roundcube (vraie IP via `X-Forwarded-For` si proxy), pas au niveau IMAP.

---

## 68. Performance : fail2ban sur gros volume

### Où sont les limites ?

- **Lecture des logs** : négligeable (inotify + regex en Python, quelques % CPU).
- **Nombre de règles firewall** : le vrai goulot. 10 000 règles iptables linéaires = chaque paquet traverse 10 000 règles.
- **Base sqlite** : tient des centaines de milliers de bans sans broncher.

### Recommandations > 1000 bans simultanés

1. Utilisez `nftables-multiport` (sets en hachage, section 44) — pas iptables linéaire.
2. `bantime` modéré + `recidive` agressif (section 23) : moins de règles actives en permanence.
3. `dbpurgeage` adapté : ne gardez pas 6 mois d'historique si vous n'en faites rien.
4. Surveillez : `time sudo nft list table inet f2b-table | wc -l` (doit rester instantané).

### Le cas extrême : attaque DDoS applicative

fail2ban n'est pas un anti-DDoS (section 1). Face à 100 000 IP distinctes :

1. Montez le rate-limiting en amont (nginx `limit_req`, Cloudflare).
2. Basculez en mode « tout bloquer sauf » temporaire si le service doit survivre.
3. fail2ban continue de tourner, mais il ne sauvera pas le serveur seul.

---

## 69. Pense-bête de poche

À imprimer et coller sur le mur du bureau (ou dans le wiki d'équipe).

```
┌─ FAIL2BAN : PENSE-BÊTE ─────────────────────────────────┐
│                                                         │
│  STATUT                                                 │
│  sudo fail2ban-client status              # jails actifs│
│  sudo fail2ban-client status sshd         # détail jail │
│                                                         │
│  DÉBANNIR (le classique du lundi)                       │
│  sudo fail2ban-client set sshd unbanip 203.0.113.45     │
│  sudo fail2ban-client set recidive unbanip 203.0.113.45 │
│                                                         │
│  BANNIR À LA MAIN (incident)                            │
│  sudo fail2ban-client set sshd banip 203.0.113.45       │
│                                                         │
│  RECHARGER                                              │
│  sudo fail2ban-client --test   # VALIDER D'ABORD !      │
│  sudo fail2ban-client reload   # recharge en douceur    │
│                                                         │
│  TESTER UN FILTER                                       │
│  sudo fail2ban-regex /var/log/auth.log \                │
│      /etc/fail2ban/filter.d/sshd.conf                   │
│                                                         │
│  LOGS                                                   │
│  sudo tail -f /var/log/fail2ban.log                     │
│  sudo grep " Ban " /var/log/fail2ban.log | tail -20     │
│                                                         │
│  RÈGLES FIREWALL                                        │
│  sudo nft list table inet f2b-table        # nftables   │
│  sudo iptables -L f2b-sshd -n              # iptables   │
│                                                         │
│  URGENCE AUTO-BAN                                       │
│  1. Console hors-bande (PVE/IPMI/cloud)                 │
│  2. fail2ban-client set sshd unbanip MON_IP             │
│  3. Ajouter MON_IP à ignoreip dans jail.local           │
│                                                         │
│  FICHIERS                                               │
│  /etc/fail2ban/jail.local      ← VOTRE config           │
│  /etc/fail2ban/jail.conf       ← NE JAMAIS TOUCHER      │
│  /var/log/fail2ban.log         ← ce qui se passe       │
│                                                         │
│  TRIANGLE : maxretry / findtime / bantime               │
│  ignoreip : VOS réseaux + supervision, TOUJOURS         │
└─────────────────────────────────────────────────────────┘
```

---

## 70. Glossaire

| Terme | Définition |
|---|---|
| **Action** | Ce que fait fail2ban lors d'un ban/unban (ex: ajouter une règle nftables). Fichiers dans `action.d/`. |
| **Backend** | Méthode de surveillance des logs : `auto`, `systemd`, `pyinotify`, `polling`. |
| **Ban** | Blocage temporaire (ou permanent) d'une IP au firewall. |
| **banaction** | L'action de bannissement par défaut (ex: `nftables-multiport`). |
| **bantime** | Durée du bannissement. `-1` = permanent. |
| **Brute-force** | Attaque par essais répétés de mots de passe. |
| **CIDR** | Notation des plages d'IP : `192.168.1.0/24` = 256 adresses. |
| **Faux positif** | IP légitime bannie par erreur (utilisateur maladroit, sonde). |
| **failregex** | Expression régulière qui détecte un échec dans un log. |
| **Filter** | Fichier de `filter.d/` contenant les failregex d'un service. |
| **findtime** | Fenêtre de temps pendant laquelle on compte les échecs. |
| **ignoreip** | Liste blanche : ces IP ne sont jamais bannies. |
| **ignoreregex** | Regex d'exclusion : ces lignes ne sont jamais comptées. |
| **Jail** | Unité de surveillance : un filter + un log + des seuils + une action. |
| **journald** | Le journal système de systemd (`journalctl`). |
| **maxretry** | Nombre d'échecs qui déclenche le ban. |
| **recidive** | Jail qui surveille les récidivistes (déjà bannis N fois). |
| **REJECT / DROP** | Deux façons de bloquer : avec (REJECT) ou sans (DROP) réponse à l'attaquant. |
| **Unban** | Levée du bannissement (fin de bantime ou manuelle). |
| **WAF** | Pare-feu applicatif web — fail2ban n'en est pas un. |

---

## 71. Quiz : 10 questions

Répondez sans regarder les réponses (section 72). Objectif : 8/10 pour valider.

**Q1.** Que se passe-t-il si vous modifiez `/etc/fail2ban/jail.conf` puis que vous faites `apt upgrade` ?

**Q2.** Votre `ignoreip` contient `192.168.10.0/24`. Un attaquant depuis `192.168.10.77` échoue 50 fois en SSH. Est-il banni ? Pourquoi ?

**Q3.** Avec `maxretry = 5` et `findtime = 10m`, une IP échoue à 10h00, 10h02, 10h04, 10h06 puis 10h21. Est-elle bannie à 10h21 ? Expliquez le compteur.

**Q4.** `fail2ban-regex` affiche `Lines: 5000 lines, 0 ignored, 0 matched, 5000 missed` alors que le log contient des attaques visibles. Citez deux causes possibles (autres que « la regex est fausse »).

**Q5.** Votre site est derrière Cloudflare sans `set_real_ip_from`. Quel est le risque si vous activez un jail nginx ?

**Q6.** Quelle commande débannit `198.51.100.23` du jail `sshd` ? Et s'il est aussi dans `recidive` ?

**Q7.** Vous changez `maxretry` avec `fail2ban-client set sshd maxretry 3`, puis vous faites `fail2ban-client reload`. Que vaut `maxretry` après le reload ?

