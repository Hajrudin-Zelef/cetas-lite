---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-17
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [3864, 3960]
sha256: f7b9cef72a28527bd9541b452935fdc217dd9a4ee1a41d1ccac61ea7721d5ead
---

# Guide fail2ban — Le bouclier anti-brute-force

**Q8.** Pourquoi un `bantime = 1w` sur le jail sshd est-il généralement une mauvaise idée, et que recommande-t-on à la place ?

**Q9.** Sur un serveur mail avec des utilisateurs humains (smartphones), quelles valeurs conseillez-vous pour le jail dovecot, et pourquoi ?

**Q10.** Vous déployez fail2ban sur un hôte Docker et les conteneurs restent joignables malgré les bans. Expliquez et proposez une correction.

---

## 72. Réponses du quiz

**R1.** `jail.conf` est **écrasé par le paquet** à la mise à jour : toute votre configuration est perdue. Il faut mettre la config dans `jail.local` (jamais touché par les mises à jour).

**R2.** Non, il n'est jamais banni : `ignoreip` est une **whitelist absolue**. C'est pour cela qu'il ne faut y mettre que des réseaux de confiance (LAN admin, supervision) — jamais un réseau exposé à des attaquants.

**R3.** Non. À 10h21, l'échec de 10h00 a plus de 10 minutes : il sort de la fenêtre glissante. Compteur à 10h21 : 10h02, 10h04, 10h06, 10h21 = 4 échecs < 5. Pas de ban. (Un 5e échec avant 10h12 aurait déclenché le ban.)

**R4.** (1) Le **format de date** n'est pas reconnu : fail2ban ignore les lignes dont il ne comprend pas le timestamp — vérifiez `Date template hits`. (2) Le **backend** lit le mauvais fichier / le fichier est vide (les attaques sont ailleurs, ex: dans le journal systemd alors que le jail lit un fichier).

**R5.** Les logs nginx voient les **IP de Cloudflare**, pas celles des visiteurs : fail2ban risque de **bannir les IP de Cloudflare lui-même**, coupant le site pour tout le monde. Il faut `set_real_ip_from` + `real_ip_header CF-Connecting-IP`.

**R6.** `sudo fail2ban-client set sshd unbanip 198.51.100.23`. Si aussi dans recidive : ajouter `sudo fail2ban-client set recidive unbanip 198.51.100.23` (chaque jail a sa propre liste de bans).

**R7.** La valeur de `jail.local` (le `set` à chaud est **perdu au reload**). Pour un changement durable, modifiez `jail.local` puis `reload`.

**R8.** Les botnets utilisent des IP jetables : l'IP bannie ne reviendra jamais, la règle firewall reste pour rien et la table grossit (ralentissement). Recommandé : `bantime` modéré (1h) + jail `recidive` agressif (1 semaine, tous ports) pour les vrais persistants, éventuellement `bantime.increment = true`.

**R9.** `maxretry = 10` (seuil haut : smartphones qui réessaient en boucle avec un ancien mot de passe), `bantime = 30m` (court : un utilisateur légitime bloqué 1h appelle le support). Les services à humains exigent des seuils tolérants.

**R10.** Le trafic vers les conteneurs passe par la chaîne `FORWARD`/`DOCKER-USER`, pas par `INPUT` où fail2ban insère ses règles par défaut. Correction : utiliser une action qui bannit dans `DOCKER-USER` (ex: `banaction = iptables-multiport[chain=DOCKER-USER]`) et tester depuis l'extérieur.

---

## 73. Pour aller plus loin

### Documentation officielle

- Site : https://www.fail2ban.org
- Wiki : https://github.com/fail2ban/fail2ban/wiki
- Manuel : `man jail.conf`, `man fail2ban-regex`, `man fail2ban-client`

### Sujets à creuser après ce guide

1. **Écrire des filters pour vos applis métier** (section 17) : chaque application maison exposée mérite son filter.
2. **Corrélation multi-serveurs** : un attaquant qui frappe 3 serveurs à 2 échecs chacun passe sous les radars. Pistes : centraliser les logs (rsyslog → Graylog/ELK) et bannir depuis la corrélation, ou partager les bans via une action qui publie sur un bus (Redis, webhook).
3. **Threat intelligence** : croiser les IP bannies avec des flux (AbuseIPDB, Blocklist.de, FireHOL) pour le reporting.
4. **IPv6 /64** : adapter les actions pour bannir le préfixe /64 d'un attaquant v6.
5. **Tests automatisés** : intégrer `fail2ban-regex` dans votre CI Ansible (un filter non testé ne part pas en prod).
6. **Alternatives complémentaires** : `crowdsec` (approche collaborative moderne, même philosophie avec une communauté), `sshguard` (plus léger).

### Livres et références réseau

- La documentation `nftables` (https://wiki.nftables.org) pour comprendre ce que font vraiment les actions.
- Les RFC 5737 / 3849 : les plages d'IP de documentation utilisées dans ce guide (`203.0.113.0/24`, `198.51.100.0/24`, `192.0.2.0/24`).

### Ce que ce guide ne couvre pas (volontairement)

- L'écriture d'actions complexes multi-étapes (report AbuseIPDB + webhook + firewall) : faisable, mais chaque brique est déjà documentée sections 22, 32, 33, 52.
- Le tuning fin des regex PCRE avancées (lookahead/lookbehind) : rarement nécessaire pour des logs.
- L'intégration avec un SIEM : chaque SIEM a son connecteur, le principe reste « fail2ban écrit, le SIEM lit `/var/log/fail2ban.log` ».

---

## 74. Références officielles

| Ressource | URL |
|---|---|
| Site officiel fail2ban | https://www.fail2ban.org |
| Dépôt GitHub | https://github.com/fail2ban/fail2ban |
| Wiki (filters, actions, FAQ) | https://github.com/fail2ban/fail2ban/wiki |
| Man page jail.conf | `man jail.conf` sur votre système |
| Paquet Debian | https://packages.debian.org/stable/fail2ban |
| AbuseIPDB API | https://docs.abuseipdb.com |
| Cloudflare API (firewall) | https://developers.cloudflare.com/api/ |
| Plages IP Cloudflare | https://www.cloudflare.com/ips/ |
| Documentation nftables | https://wiki.nftables.org |

---

## 75. Note finale : la philosophie du gardien

fail2ban n'est pas un mur : c'est un **videur**. Il ne rend pas votre serveur invulnérable ; il rend les attaques automatisées **non rentables** :

- Le bot qui scanne 10 000 serveurs abandonne le vôtre après 5 essais.
- L'attaquant persistant se retrouve banni sur tous les ports pendant une semaine.
- Vous, vous dormez, et le lundi matin le rapport quotidien vous dit qui a frappé.

Les trois habitudes qui font la différence en production :

1. **`ignoreip` d'abord** : on ne protège bien que ce qu'on ne casse pas.
2. **Tester avant de bannir** : `fail2ban-regex` n'est pas une option.
3. **Relire les logs** : un gardien qu'on ne supervise pas finit par bannir les clients.

Bon courage, et que vos `Currently banned` restent à zéro les nuits calmes. 🛡️

---

*Guide rédigé pour Zelef — chef de service systèmes. Vérifié : toutes les commandes et tous les blocs de configuration suivent la syntaxe fail2ban ≥ 0.11 / 1.x. Les adresses IP utilisées sont des plages de documentation (RFC 5737). Adaptez `ignoreip`, les `logpath` et les ports à votre infrastructure avant toute mise en production.*
