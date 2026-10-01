---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-1
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [1, 182]
sha256: ca6a03ae51c4bf21537b37c73894b3c3a33d3e1dba85858c0cfdaa1af1ad5abe
---

# Guide fail2ban — Le bouclier anti-brute-force

> **Audience :** sysadmins et chefs de service systèmes (production réelle).
> **Prérequis :** Debian 11/12/13 ou Ubuntu 20.04/22.04/24.04, accès root/sudo, bases iptables/nftables.
> **Objectif :** installer, configurer, exploiter et superviser fail2ban sans se bannir soi-même.

---

# Sommaire

1. Principe : scan de logs → bannissement
2. Architecture : jails, filters, actions
3. Installation sur Debian/Ubuntu
4. Arborescence et fichiers
5. Premier jail.local minimal
6. bantime, findtime, maxretry expliqués
7. Le jail sshd en détail
8. Le jail sshd (modes et variantes)
9. nginx-http-auth : protéger le basic auth
10. nginx-badbot et nginx-botsearch
11. postfix : relais et SASL
12. dovecot : POP3/IMAP sous pression
13. proftpd et vsftpd
14. recidive : les récidivistes
15. Activer plusieurs jails sans conflit
16. fail2ban-regex : tester son filtre
17. Créer un filter personnalisé pas à pas
18. Regex : guide de survie
19. Les actions : iptables, nftables, ufw
20. Action nftables en détail
21. Action ufw
22. Actions multiples et combinées
23. Tuning par service : valeurs recommandées
24. ignoreip : whitelist (son réseau, la supervision !)
25. Ne pas se bannir soi-même : 7 techniques
26. fail2ban-client : usage quotidien
27. Débannir : le classique du lundi
28. Bannissement temporaire vs permanent
29. Persistent bans (sqlite db)
30. Logs et niveaux de debug
31. Troubleshooting des logs
32. Intégration Cloudflare (bannir à l'edge)
33. AbuseIPDB : remonter les attaquants
34. Monitoring Zabbix
35. Monitoring Prometheus
36. Alertes mail et webhooks
37. Sauvegarde et restauration de jail.local
38. Mise à jour de fail2ban
39. 10+ erreurs classiques et solutions
40. Cas pratique : protéger un serveur web exposé
41. Cas pratique : durcir un Proxmox VE
42. Cas pratique : sécuriser un bastion SSH
43. Cas pratique : protéger un webmail Roundcube
44. Performance : fail2ban sur gros volume
45. Backend : auto, systemd, pyinotify, polling
46. Le fuseau horaire : la cause cachée
47. IPv6 avec fail2ban
48. Multi-instance et mode socket
49. Bannir via ipset pour la performance
50. Exclure les IPs légitimes : failregex d'exclusion
51. Rotation de logs et fail2ban
52. Fail2ban derrière un proxy inverse
53. Tester en sécurité : mode dry-run mental
54. Scripts d'administration utiles
55. Intégration Ansible
56. Intégration avec GeoIP
57. Le fichier filter.d : anatomie
58. Le fichier action.d : anatomie
59. Jail.d : surcharges modulaires
60. Ordre de fusion de la configuration
61. Durcissement de fail2ban lui-même
62. Bannir un pays entier : réaliste ?
63. Corrélation avec auditd
64. Fail2ban et Docker : attention
65. Fail2ban sur Proxmox : règles avancées
66. Recette de jail.local « gold »
67. Checklist de mise en production
68. Checklist de revue hebdomadaire
69. Checklist de revue mensuelle
70. Pense-bête de poche
71. Glossaire
72. Quiz : 10 questions
73. Réponses du quiz
74. Pour aller plus loin
75. Références officielles

---

## 1. Principe : scan de logs → bannissement

fail2ban n'est pas un pare-feu : c'est un **gardien qui lit les journaux**.

Le cycle de vie d'une attaque contre SSH ressemble à ceci :

```
┌──────────┐   tente de se connecter   ┌──────────┐
│ Attaquant │ ───────────────────────▶ │   sshd   │
│203.0.113.│ ◀──── refusé ─────────── │ port 22  │
└──────────┘                           └────┬─────┘
                                            │ écrit
                                            ▼
                                   /var/log/auth.log
                                   "Failed password ..."
                                            │
                                            ▼
                                   ┌───────────────┐
                                   │   fail2ban    │
                                   │  jail[sshd]   │
                                   │ maxretry=5    │
                                   └───────┬───────┘
                                           │ 5 échecs < findtime
                                           ▼
                              ┌────────────────────────┐
                              │ action ban :           │
                              │ iptables -A ... DROP   │
                              └────────────────────────┘
                                           │
                                           ▼
                                   Attaquant bloqué
                                   (bantime = 1h)
```

Trois étapes, toujours les mêmes :

1. **Observer** : fail2ban surveille en continu des fichiers de logs (ou le journal systemd).
2. **Compter** : chaque ligne qui correspond au filtre incrémente un compteur par IP.
3. **Agir** : si le compteur dépasse `maxretry` pendant `findtime`, une action bannit l'IP pendant `bantime`.

Le bannissement est **temporaire par défaut** : après `bantime`, l'IP est débannie automatiquement. C'est volontaire : on veut bloquer les robots, pas devenir soi-même une source de déni de service contre des utilisateurs légitimes dont le mot de passe a expiré.

### Pourquoi c'est efficace en production

- Un scan SSH classique frappe 50 à 500 fois par heure. Avec `maxretry=5` et `findtime=10m`, le bot est bloqué après ~30 secondes.
- Le CPU du serveur est soulagé : `iptables`/`nftables` jette les paquets avant qu'ils atteignent sshd.
- Les logs restent lisibles : au lieu de 10 000 lignes « Failed password », on n'en voit que 5 par attaquant.

### Ce que fail2ban ne fait pas

| Ce n'est pas... | Parce que... |
|---|---|
| Un pare-feu | il ne contrôle que les IP déjà vues dans les logs |
| Un WAF | il ne comprend pas le HTTP, il lit des lignes de texte |
| Une authentification forte | il ne remplace ni les clés SSH ni le 2FA |
| Un anti-DDoS | 100 000 IP distinctes satureront le compteur avant tout |

fail2ban est une **couche** de défense en profondeur, à combiner avec : clés SSH, 2FA, pare-feu restrictif, mises à jour.

### Exemple concret : à quoi ressemble une attaque

Extrait réel (anonymisé) de `/var/log/auth.log` :

```
Sep 26 08:14:02 srv-web sshd[2041]: Failed password for root from 203.0.113.45 port 51234 ssh2
Sep 26 08:14:04 srv-web sshd[2043]: Failed password for root from 203.0.113.45 port 51235 ssh2
Sep 26 08:14:06 srv-web sshd[2045]: Failed password for invalid user admin from 203.0.113.45 port 51236 ssh2
Sep 26 08:14:08 srv-web sshd[2047]: Failed password for invalid user oracle from 203.0.113.45 port 51237 ssh2
Sep 26 08:14:10 srv-web sshd[2049]: Failed password for root from 203.0.113.45 port 51238 ssh2
```

Cinq lignes en huit secondes : c'est exactement le schéma qu'un jail `sshd` détecte.

---

## 2. Architecture : jails, filters, actions

fail2ban est organisé en trois concepts. Les retenir, c'est comprendre 80 % de l'outil.

### Le filter : « quoi chercher »

Un filter est un fichier dans `/etc/fail2ban/filter.d/<nom>.conf` qui contient des **expressions régulières** (`failregex`). Il décrit *ce qu'est une attaque* dans un log donné.

Exemple simplifié de `filter.d/sshd.conf` :

```ini
[INCLUDES]
before = common.conf

[Definition]
failregex = ^%(__prefix_line)sFailed (?:password|publickey) for <F-USER>.*</F-USER> from <HOST>( port \d+)?(?: ssh2)?\s*$
ignoreregex =
```

