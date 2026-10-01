---
id: collect-261001-general-networking/general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026-5
title: "Vérifier la version installée"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026.md
source_anchor: ""
source_lines: [420, 467]
sha256: 60f4ad0b738e9a9bf9df83217f09903bc75360b3337bd6067a4b7c08f86eecc2
---

# Vérifier la version installée

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Le service refuse de démarrer | Erreur de syntaxe dans un fichier jail.d/ | Lancer `fail2ban-client -t` pour valider la config avant redémarrage | 
| Aucune IP jamais bannie malgré des échecs visibles | logpath incorrect ou journalmatch mal ciblé | Tester avec `fail2ban-regex` contre le journal réel | 
| Je suis bloqué hors de mon propre serveur | ignoreip non configuré avant activation du jail | Passer par la console de secours de l’hébergeur, puis unbanip | 
| Le jail sshd ne démarre pas avec backend systemd | python3-systemd manquant | `sudo apt install python3-systemd` puis redémarrer | 
| Les modifications de jail.conf n’ont aucun effet | Fichier surchargé par jail.local ou defaults-debian.conf | Déplacer les changements vers jail.local | 
| Les IPv6 ne sont jamais bannies | Action de bannissement encore configurée en iptables v4 uniquement | Basculer vers banaction = nftables qui gère nativement IPv4/IPv6 | 
| Trop d’emails de notification | action = %(action_mwl)s appliqué à tous les jails | Limiter la notification email aux jails critiques (sshd, recidive) | 
| Une IP légitime reste bannie après correction de ignoreip | Le bannissement existant n’est pas retiré automatiquement | Débannir manuellement avec fail2ban-client set JAIL unbanip IP | 

## Foire aux questions

**Fail2ban fonctionne-t-il sur Windows Server ?**

Non, Fail2ban est un outil natif Linux qui s’appuie sur les journaux système et les pare-feu netfilter/nftables. Sur Windows, des équivalents comme IPBan ou des règles de pare-feu natives couvrent un besoin similaire, mais avec une architecture différente.

**Fail2ban ralentit-il les performances du serveur ?**

L’impact est négligeable sur un serveur correctement dimensionné : le démon consomme peu de mémoire et de CPU, sauf en cas de backend systemd mal filtré qui lirait l’intégralité du journal sans `journalmatch` ciblé, ce qui peut générer une charge inutile.

**Faut-il utiliser Fail2ban si j’ai déjà désactivé l’authentification par mot de passe SSH ?**

Oui, cela reste utile. Fail2ban continue de ralentir les tentatives d’énumération de noms d’utilisateurs et protège les autres services exposés (web, mail) qui ne bénéficient pas forcément d’une authentification par clé.

**Combien de temps une IP reste-t-elle bannie par défaut ?**

Le réglage par défaut de `bantime` est d’une heure (3600 secondes), mais c’est entièrement personnalisable. Avec `bantime.increment` activé, la durée double à chaque récidive jusqu’à un plafond que vous définissez.

**Peut-on utiliser Fail2ban derrière un reverse proxy ou un CDN ?**

Oui, mais il faut s’assurer que les journaux contiennent la véritable IP du visiteur via l’en-tête `X-Forwarded-For` plutôt que l’IP du proxy lui-même, sans quoi Fail2ban bannirait le proxy et non l’attaquant réel.

**Fail2ban remplace-t-il un pare-feu applicatif web (WAF) ?**

Non. Fail2ban agit essentiellement sur les couches réseau (IP et ports) en réaction à des motifs de journaux. Un WAF inspecte le trafic HTTP en profondeur en temps réel. Les deux outils sont complémentaires plutôt que substituables.

**Comment savoir si mon serveur est réellement ciblé par des attaques ?**

Consultez `/var/log/fail2ban.log` et comptez les occurrences de bannissement sur 24 heures. Un VPS exposé sur Internet reçoit généralement plusieurs dizaines à centaines de tentatives de connexion SSH automatisées par jour, même sans ciblage spécifique.

**Existe-t-il une interface graphique pour Fail2ban ?**

Pas d’interface officielle. Certains outils tiers proposent des tableaux de bord (via Grafana avec export des logs, ou des panneaux d’hébergement comme CyberPanel), mais l’administration native reste en ligne de commande via `fail2ban-client`.

**Fail2ban est-il suffisant seul pour sécuriser un serveur exposé sur Internet ?**

Non, il s’agit d’une couche parmi d’autres. Un déploiement solide combine authentification par clé SSH ou passkeys, mises à jour de sécurité régulières, pare-feu correctement configuré, et surveillance des journaux, avec Fail2ban comme filet automatique contre les tentatives de force brute répétées plutôt que comme unique ligne de défense.
