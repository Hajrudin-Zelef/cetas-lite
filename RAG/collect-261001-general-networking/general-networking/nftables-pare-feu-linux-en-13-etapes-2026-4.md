---
id: collect-261001-general-networking/general-networking/nftables-pare-feu-linux-en-13-etapes-2026-4
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "valuation"]
source: docs/RAG/collect-261001-general-networking/nftables-pare-feu-linux-en-13-etapes-2026.md
source_anchor: ""
source_lines: [283, 333]
sha256: a2d481d9278938a98beb73b10cf1d0ebd57d3e289a53431f2e951cd87eed88b5
---

# Debian / Ubuntu

- **Le service nftables refuse de démarrer après modification du fichier de configuration.** Lancez`nft -c -f /etc/nftables.conf` pour identifier la ligne fautive avant de relancer le service, la plupart du temps une virgule ou une accolade mal fermée.
- **La connexion SSH se coupe immédiatement après application des règles.** Gardez toujours une session console active en parallèle, ou programmez un rollback automatique avec`at now + 5 minutes` qui restaure l’ancienne configuration si vous ne confirmez pas manuellement.
- **Les règles semblent ignorées malgré une configuration correcte.** Vérifiez qu’aucun ancien service iptables, ufw ou firewalld ne tourne encore en parallèle avec`systemctl status ufw firewalld` .
- **Le NAT ne fonctionne pas malgré une règle masquerade correcte.** Confirmez que`net.ipv4.ip_forward` vaut bien 1 avec`sysctl net.ipv4.ip_forward` , sinon aucun paquet ne transitera entre interfaces.
- **Les logs n’apparaissent jamais dans journalctl.** Vérifiez que le service systemd-journald tourne correctement et que la règle log est bien positionnée avant le verdict final drop, pas après.
- **Le compteur d’une règle reste à zéro alors que le trafic correspondant existe.** Une règle placée trop bas dans la chaîne peut être court-circuitée par une règle antérieure qui matche déjà le paquet, notamment la règle established,related.
- **Erreur “Could not process rule: No such file or directory” lors d’un ajout de règle.** La table ou la chaîne référencée n’existe pas encore, vérifiez l’orthographe exacte avec`nft list tables` .
- **Les performances se dégradent avec plusieurs milliers de règles individuelles.** Remplacez les règles répétitives par des sets nommés, qui utilisent des structures de données optimisées côté noyau au lieu d’une évaluation séquentielle.
- **Après une mise à jour de la distribution, les anciennes règles iptables ne sont plus prises en compte.** Certaines migrations de version basculent automatiquement le backend vers nft, exécutez une conversion avec iptables-restore-translate pour récupérer vos anciennes règles.

## Astuces avancées pour les administrateurs expérimentés

Une fois les bases maîtrisées, plusieurs fonctionnalités permettent d’aller plus loin sur des serveurs à fort trafic ou des infrastructures complexes.

Les flowtables accélèrent le transit de paquets déjà établis en les faisant passer par un chemin de traitement raccourci dans le noyau, particulièrement utile sur un routeur qui encaisse un débit important entre plusieurs interfaces. Elles évitent de repasser par l’intégralité de la chaîne de règles pour chaque paquet d’une connexion déjà validée.

```
sudo nft add flowtable inet filtre acceleration { hook ingress priority 0 \; devices = { eth0, eth1 } \; }
sudo nft add rule inet filtre transit ct state established flow add @acceleration
```
L’intégration avec CrowdSec ou Fail2ban permet d’automatiser complètement le bannissement dynamique : ces outils analysent les logs en continu et injectent directement des adresses IP malveillantes dans un set nftables avec timeout, sans intervention manuelle. C’est la combinaison recommandée pour un serveur exposé qui reçoit un volume élevé de tentatives d’intrusion automatisées.

Pour les infrastructures multi-serveurs, centraliser les logs nftables vers un outil comme Wazuh ou Graylog permet de corréler des tentatives d’attaque réparties sur plusieurs machines et de détecter des campagnes de scan coordonnées qui passeraient inaperçues serveur par serveur.

## nftables et Docker : éviter les conflits de règles

Docker manipule directement les règles de filtrage réseau pour exposer les ports de ses conteneurs, et cette manipulation entre régulièrement en collision avec une configuration nftables écrite à la main. Le démon Docker crée ses propres chaînes et injecte des règles d’acceptation qui peuvent contourner silencieusement une politique drop pourtant censée tout bloquer par défaut. Un administrateur qui publie un conteneur avec `-p 5432:5432` peut découvrir, parfois des semaines plus tard, que ce port PostgreSQL est accessible depuis Internet malgré une chaîne d’entrée nftables strictement restrictive.

Ce comportement s’explique par la façon dont Docker gère le NAT et le transit de paquets pour ses réseaux de conteneurs : les paquets destinés à un port publié passent par la chaîne forward et par des règles NAT propres à Docker, souvent insérées avec une priorité qui les place avant les règles de filtrage écrites manuellement. La documentation Docker recommande d’ailleurs explicitement de ne pas gérer ces chaînes générées automatiquement.

```
# Inspecter les chaînes injectées par Docker
sudo nft list chain ip docker docker
sudo nft list chain inet filtre transit
# Restreindre l'exposition d'un conteneur à une IP précise plutôt qu'à toutes les interfaces
docker run -p 203.0.113.10:5432:5432 postgres
```
La solution la plus fiable reste de limiter les ports publiés à une adresse IP spécifique plutôt qu’à `0.0.0.0`, et de vérifier systématiquement avec `nmap` depuis l’extérieur après chaque déploiement de conteneur, exactement comme à l’étape 12 de ce tutoriel. Sur un serveur qui héberge à la fois des services système et des conteneurs Docker, considérez ces deux mondes comme deux couches de filtrage distinctes qu’il faut auditer séparément.

## nftables sur les hébergeurs cloud européens : OVHcloud, Scaleway, Hetzner

La majorité des serveurs Linux déployés en France et en Europe tournent aujourd’hui chez des hébergeurs comme OVHcloud, Scaleway ou Hetzner, qui proposent tous un pare-feu réseau en amont, au niveau de l’infrastructure plutôt que du système d’exploitation. Ce filtrage périphérique ne dispense pas d’une configuration nftables locale : les deux couches se complètent et répondent à des besoins différents.

Le pare-feu réseau de l’hébergeur bloque le trafic avant même qu’il n’atteigne la carte réseau virtuelle du serveur, ce qui protège contre des volumes d’attaque qu’un pare-feu local ne pourrait pas encaisser seul, comme certaines tentatives de saturation réseau. nftables, de son côté, offre une granularité que les consoles d’hébergeur ne permettent généralement pas : limitation de débit par connexion, journalisation détaillée, sets dynamiques alimentés par des outils comme CrowdSec, ou logique conditionnelle complexe basée sur l’état des connexions.

Une pratique de défense en profondeur consiste à dupliquer les règles essentielles aux deux niveaux : bloquer un port sensible à la fois dans la console de l’hébergeur et dans `/etc/nftables.conf`. Si une des deux couches est mal configurée ou temporairement désactivée pendant une opération de maintenance, l’autre continue de protéger le serveur. Cette redondance coûte quelques minutes de configuration supplémentaires mais évite qu’une simple erreur de manipulation sur une console web n’expose brutalement un service critique.

## nftables vs ufw vs firewalld : quel outil choisir en 2026

nftables reste le moteur sous-jacent dans les trois cas, mais le choix de l’interface change radicalement l’expérience d’administration selon le contexte d’usage.

