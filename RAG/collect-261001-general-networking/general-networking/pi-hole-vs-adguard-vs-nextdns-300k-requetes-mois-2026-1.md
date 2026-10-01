---
id: collect-261001-general-networking/general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026-1
title: "Temps de requete affiche sur la ligne \"Query time\""
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Samsung"]
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026.md
source_anchor: ""
source_lines: [1, 32]
sha256: 5f1cee35de42b26c13a3a238785f857020decc1ed82c223beb7110208bef2b66
---

# Temps de requete affiche sur la ligne "Query time"

Depuis le début de l’année 2026, trois noms reviennent sans cesse dans les discussions sur le blocage publicitaire réseau : Pi-hole, AdGuard Home et NextDNS. Les trois filtrent les publicités, les traqueurs et les domaines malveillants directement au niveau DNS, avant même que votre navigateur ou votre application ne charge quoi que ce soit. Mais leurs approches divergent radicalement. Pi-hole mise sur un sinkhole DNS auto-hébergé et communautaire, AdGuard Home ajoute le chiffrement natif et une interface plus moderne, et NextDNS déporte tout dans le cloud pour éviter la moindre maintenance. Pour un lecteur en France ou ailleurs en Europe soucieux du RGPD et de la localisation de ses données, le choix ne se résume pas à « lequel bloque le plus de pubs ». Il touche à l’endroit où transitent vos requêtes DNS, à qui les héberge, et au temps que vous êtes prêt à consacrer à la maintenance. Ce comparatif détaille les trois outils avec des données de tarification vérifiées, des caractéristiques techniques à jour et des recommandations concrètes selon votre profil.

## Pourquoi le blocage DNS reste la protection réseau la plus efficace en 2026

Le principe du blocage DNS est simple. Chaque fois qu’un appareil veut charger une ressource, il interroge d’abord un serveur DNS pour transformer un nom de domaine en adresse IP. Pi-hole, AdGuard Home et NextDNS s’intercalent à cette étape précise. Si le domaine demandé figure sur une liste de blocage, le résolveur répond par un domaine inexistant ou une adresse nulle, et la ressource ne se charge jamais. Ni script publicitaire, ni traqueur, ni pixel de mesure.

L’intérêt par rapport à une extension de navigateur comme uBlock Origin tient à la couverture. Une extension ne protège que le navigateur dans lequel elle est installée. Un blocage DNS configuré sur la box, le routeur ou le DHCP protège tous les appareils du réseau : smart TV, console de jeu, montre connectée, application mobile qui n’accepte aucune extension. C’est aussi la seule méthode qui fonctionne uniformément sur un iPhone, une TV Samsung et un ordinateur portable sous Linux sans rien installer localement sur chaque appareil.

Cette approche a toutefois une limite connue et documentée par les trois projets eux-mêmes : elle ne bloque pas les publicités qui partagent le même domaine que le contenu. Les pré-rolls YouTube ou les coupures publicitaires Twitch, servis depuis les domaines googlevideo.com ou ttvnw.net, échappent largement au filtrage DNS parce que bloquer ces domaines casserait aussi la lecture vidéo. Pour ces cas précis, une extension de navigateur reste complémentaire. Le blocage DNS, en revanche, coupe la grande majorité des traqueurs publicitaires tiers, des domaines de télémétrie et d’une partie croissante des domaines de phishing et de malwares, ce qui en fait un outil de sécurité autant que de confort.

## Pi-hole en 2026 : le sinkhole DNS open source historique

Pi-hole reste le nom le plus connu du secteur, porté depuis 2014 par une communauté et non par une entreprise. Le projet est distribué sous licence EUPL 1.2, une licence open source approuvée par la Commission européenne, et reste gratuit sans aucune version payante. Techniquement, Pi-hole s’appuie sur deux briques : le cœur applicatif qui gère l’interface web et les listes, et FTL (Faster Than Light), le moteur de résolution DNS qui applique le filtrage en temps réel. L’ensemble tourne sur à peu près n’importe quoi : un Raspberry Pi, un conteneur Docker, une machine virtuelle ou un vieux PC recyclé.

La branche v6, qui a succédé à la v5 avec une réécriture importante de l’interface et du moteur de configuration, constitue la version de référence en 2026. Pi-hole fonctionne en **DNS forwarder** : il ne résout pas lui-même les domaines mais transmet chaque requête non bloquée à un résolveur en amont (Cloudflare, Quad9, ou votre FAI). C’est là sa principale faiblesse technique : Pi-hole ne chiffre pas nativement ces requêtes sortantes. Pour obtenir du DNS-over-HTTPS ou DNS-over-TLS, il faut ajouter manuellement un composant tiers comme Unbound ou cloudflared, une étape supplémentaire que beaucoup d’utilisateurs débutants ignorent ou repoussent.

Ce que Pi-hole fait remarquablement bien, c’est la gestion communautaire des listes de blocage. Des dizaines de listes maintenues publiquement, référencées dans la documentation officielle (StevenBlack, OISD, HaGeZi), s’ajoutent en quelques clics et couvrent ensemble plusieurs centaines de milliers de domaines publicitaires et de traqueurs. La gestion des clients se fait par groupes, ce qui permet d’appliquer des règles différentes à un sous-réseau d’invités ou aux appareils des enfants. Notre tutoriel d’installation de Pi-hole en 12 étapes détaille la procédure complète pour une mise en service sur Raspberry Pi en moins de 40 minutes.

## AdGuard Home en 2026 : le successeur chiffré et plus permissif

AdGuard Home est développé par la société AdGuard, fondée en 2009 et connue pour son extension de navigateur et son application de blocage publicitaire mobile. Il faut bien distinguer trois produits différents portant un nom proche : AdGuard (l’extension et l’appli, payantes), AdGuard DNS (le service cloud public, gratuit avec option payante), et AdGuard Home, le serveur DNS auto-hébergé dont il est question ici. AdGuard Home est publié sous licence GPLv3 sur GitHub, entièrement gratuit, sans version payante propre à ce produit précis.

La différence technique majeure avec Pi-hole tient au chiffrement. AdGuard Home intègre nativement le support de DNS-over-HTTPS, DNS-over-TLS, DNS-over-QUIC et DNSCrypt, à la fois pour recevoir les requêtes des appareils du réseau et pour interroger les résolveurs en amont, par défaut Quad9. Aucun composant tiers à installer : le chiffrement fonctionne dès l’activation dans les paramètres. La branche stable v0.107.x équipe la majorité des installations en production en 2026, pendant que la série v0.108 progresse en version bêta avec des correctifs de sécurité et un renforcement de la validation des upstreams DoH.

L’interface d’administration se distingue aussi par sa granularité. Chaque client peut recevoir des règles individuelles par adresse IP, plage CIDR ou adresse MAC : upstream DNS dédié, listes de blocage spécifiques, plages horaires d’accès, blocage de services précis comme TikTok ou YouTube en un clic. Le contrôle parental et le SafeSearch forcé sont intégrés nativement, là où Pi-hole demande de bricoler des règles regex. Côté plateformes, AdGuard Home couvre Linux, Windows, macOS, FreeBSD, Docker et surtout OpenWrt, où il existe en paquet natif directement installable sur un routeur, un avantage que ni Pi-hole ni NextDNS ne proposent.

## NextDNS en 2026 : la solution cloud sans matériel à gérer

NextDNS change complètement de philosophie. Pas de Raspberry Pi, pas de conteneur à maintenir, pas de mise à jour à surveiller : on crée un compte, on configure ses listes de blocage depuis un tableau de bord, et on pointe ses appareils vers les serveurs NextDNS. Le filtrage, les statistiques et les règles de sécurité tournent entièrement côté cloud, avec des points de présence répartis pour limiter la latence ajoutée par ce saut réseau supplémentaire.

### Qui se cache derrière NextDNS ?

