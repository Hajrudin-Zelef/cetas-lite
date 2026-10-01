---
id: collect-261001-general-networking/general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026-1
title: "repérez le nom de votre interface, par exemple eth0 ou ens18"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "distribution", "incident", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: 005a0cc4cefb2cd07d1a84b0be73cfbcb695fbb89dc426bb457fa3c132882adf
---

# repérez le nom de votre interface, par exemple eth0 ou ens18

Suricata vient de passer un cap important : depuis le 9 juillet 2026, la branche 8.0.x est la seule activement maintenue par l’Open Information Security Foundation (OISF), la branche 7.x étant officiellement en fin de vie (EOL) depuis le 7 juillet 2026 avec l’ultime correctif 7.0.17. Le rythme de publication est resté soutenu depuis : après la 8.0.6 de juillet, l’OISF a livré la 8.0.7 en septembre 2026, confirmant un calendrier de correctifs régulier sur la nouvelle branche. Pour les administrateurs réseau, RSSI et développeurs qui gèrent l’infrastructure d’une PME ou d’un service public en France, ce basculement n’est pas un détail technique : sans mise à jour, les serveurs restent exposés sans les correctifs de sécurité les plus récents. Ce tutoriel vous montre, étape par étape, comment déployer Suricata 8.0.7, la dernière version stable disponible en septembre 2026, comme système de détection et de prévention d’intrusion (IDS/IPS) sur un serveur Debian ou Ubuntu, configurer des règles de détection pertinentes, et l’intégrer dans une chaîne de supervision réseau complète.

Suricata est un moteur open source d’inspection profonde de paquets (DPI), capable d’analyser le trafic HTTP, TLS, DNS et bien d’autres protocoles en temps réel pour repérer des signatures d’attaque, des comportements suspects ou des flux non conformes à une politique de sécurité. Contrairement à un pare-feu classique qui filtre sur des règles statiques de ports et d’adresses, Suricata inspecte le contenu des paquets et peut donc détecter une exploitation de vulnérabilité, un malware en communication avec son serveur de commande, ou une exfiltration de données. Il tourne aussi bien en mode IDS (détection passive, alertes) qu’en mode IPS (blocage actif en ligne), et sert de moteur de détection dans des distributions de sécurité populaires comme Security Onion, dont la version 3.2.0 est sortie fin juillet 2026.

## Pourquoi déployer Suricata 8.0.6 en 2026

La fin de vie de Suricata 7 change la donne pour toute organisation qui exploite encore cette branche. Selon la politique EOL publiée par l’OISF, les versions 7.0.x sont en fin de vie depuis le 7 juillet 2026 avec la publication de l’ultime correctif 7.0.17, après une série de mises à jour régulières qui a vu se succéder la 7.0.13 en novembre 2025, la 7.0.14 en janvier 2026, la 7.0.15 en mars puis la 7.0.16 en mai 2026 ; les versions 6.0.x sont quant à elles en fin de vie depuis août 2024. Concrètement, cela signifie qu’aucun correctif de sécurité ne sera plus publié pour ces branches, ce qui expose les moteurs de détection eux-mêmes à des vulnérabilités non patchées, un comble pour un outil censé protéger le réseau.

Suricata 8.0.6, publié le 9 juillet 2026 en même temps que le correctif 7.0.17 de la branche sortante, a rapidement été suivi par la 8.0.7 en septembre 2026, confirmant le rythme de publication engagé par l’OISF depuis le lancement de la branche 8.0.x avec la 8.0.3 en janvier 2026, puis la 8.0.4 en mars et la 8.0.5 en mai. La 8.0.7 est aujourd’hui le choix recommandé pour un déploiement neuf. L’écosystème a suivi rapidement : le paquet Debian `suricata 1:8.0.6-1` a été accepté dans *unstable* et *stable-backports* le 10 juillet 2026, ce qui permet une installation propre sur Debian 12 et Debian 13 sans avoir besoin de compiler depuis les sources, en attendant la mise à jour du paquet vers la 8.0.7. Security Onion, la distribution de supervision réseau la plus utilisée dans les SOC (Security Operations Center) en Europe, a également suivi le mouvement : la version 3.1.0 sortie le 28 mai 2026 embarque Suricata 8.0.5, et la 3.2.0 de fin juillet 2026 introduit en plus des fonctionnalités d’IA agentique et un support étendu des règles Sigma pour la corrélation d’événements.

Pour une entreprise soumise à la directive NIS2 ou au Cyber Resilience Act, disposer d’un IDS/IPS à jour et documenté fait partie des exigences de base en matière de détection d’incident. Un moteur de détection open source, gratuit et audité par une fondation indépendante, reste l’option la plus accessible pour les PME qui n’ont pas le budget d’une solution commerciale de type NDR (Network Detection and Response).

Security Onion 3.2.0 illustre également une tendance de fond de l’écosystème en 2026 : l’ajout de capacités d’IA agentique en complément de la détection par signatures, avec un modèle Gemma proposé comme option hébergée pour assister l’analyse d’alertes. Cela ne remplace pas Suricata, mais montre que le moteur de détection par signatures reste la brique de base sur laquelle viennent désormais se greffer des couches de corrélation plus intelligentes. Pour une PME ou une collectivité qui déploie son premier IDS/IPS, l’essentiel reste néanmoins de maîtriser Suricata lui-même avant d’envisager ces couches additionnelles.

## Prérequis : matériel, logiciels et versions

Avant de commencer l’installation, vérifiez que votre environnement correspond à ces prérequis. Ce tutoriel a été rédigé et testé sur Debian 13 (Trixie), mais fonctionne également sur Debian 12 (Bookworm) avec les backports, ainsi que sur Ubuntu 24.04 LTS et 25.10.

| Composant | Version minimale recommandée | Remarque | 
|---|---|---|
| Système d’exploitation | Debian 12/13 ou Ubuntu 24.04+ | Autres distributions compatibles via compilation source | 
| Suricata | 8.0.6 (juillet 2026) | Suricata 7.x est EOL depuis le 7 juillet 2026 | 
| RAM | 4 Go minimum, 8 Go recommandé | Dépend du débit réseau à inspecter | 
| CPU | 4 cœurs minimum | Suricata exploite le multi-threading natif | 
| Espace disque | 50 Go minimum | Pour les logs, PCAP et règles de détection | 
| Accès réseau | Interface en mode promiscuous ou port mirroring (SPAN) | Nécessaire pour capturer tout le trafic à inspecter | 
| Droits | Accès root ou sudo | Requis pour l’installation de paquets et la configuration réseau | 

Vous aurez également besoin d’un accès à Internet pour télécharger les règles de détection (les “rulesets”) depuis des sources comme Emerging Threats Open ou Abuse.ch, et d’une machine de test isolée si vous voulez d’abord valider votre configuration avant un déploiement en production. Si vous n’avez pas encore de labo de test, notre tutoriel pour installer Kali Linux 2026.2 dans VirtualBox peut servir de base pour générer du trafic d’attaque contrôlé et valider vos règles Suricata.

## Étape 1 : Préparer le système et identifier l’interface réseau

Commencez par mettre à jour votre système et identifier l’interface réseau que Suricata va surveiller. Cette interface doit voir passer tout le trafic à inspecter, ce qui implique généralement un port mirroring (SPAN) sur votre switch si vous inspectez du trafic autre que celui de la machine elle-même.

```
sudo apt update && sudo apt upgrade -y
ip a
# repérez le nom de votre interface, par exemple eth0 ou ens18
sudo ethtool -k eth0 | grep -i offload
```
Désactivez les fonctions d’offload matériel (segmentation offload, generic receive offload) sur l’interface de capture, car elles peuvent fragmenter la vue que Suricata a du trafic réel et provoquer des faux négatifs.

`sudo ethtool -K eth0 tso off gso off gro off lro off`
## Étape 2 : Installer Suricata 8.0.6 depuis le dépôt officiel

Sur Debian, le paquet `suricata 1:8.0.6-1` est disponible dans *unstable* et via *stable-backports* depuis le 10 juillet 2026. Sur Debian 12 stable, ajoutez le dépôt backports pour obtenir la version à jour plutôt que l’ancienne version packagée dans la branche stable.

