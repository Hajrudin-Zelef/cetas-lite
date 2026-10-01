---
id: collect-261001-general-networking/general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026-4
title: "Temps de requete affiche sur la ligne \"Query time\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026.md
source_anchor: ""
source_lines: [137, 201]
sha256: 65440c99025771eea07c949a095fb0e418435f9213fb3e4c35f6e6b83510f393
---

# Temps de requete affiche sur la ligne "Query time"

- **Vous voulez le contrôle maximal sur vos données et un boîtier dédié :** Pi-hole, éventuellement complété par Unbound pour le chiffrement amont.
- **Vous voulez du chiffrement natif et une interface moderne sans rien ajouter :** AdGuard Home.
- **Vous n’avez ni l’envie ni le temps de gérer du matériel :** NextDNS, en particulier son offre gratuite pour un usage individuel léger.
- **Vous vous déplacez souvent avec plusieurs appareils :** NextDNS, pour la cohérence de la protection hors du réseau domestique.
- **Vous administrez un petit parc informatique professionnel :** AdGuard Home pour les règles par client, ou NextDNS Business si vous préférez déléguer l’infrastructure.
- **Vous possédez déjà un routeur sous OpenWrt :** AdGuard Home, seul à proposer un paquet natif installable directement dessus.
- **Vous voulez uniquement du contrôle parental simple, sans bricolage :** NextDNS ou AdGuard Home, tous deux dotés de filtres prêts à l’emploi.

## Avantages et inconvénients de chaque solution

### Pi-hole

**Avantages :** gratuit à vie, licence ouverte, communauté immense, listes de blocage abondantes, léger, fonctionne sur du matériel très modeste, aucune dépendance à une entreprise tierce.

**Inconvénients :** pas de chiffrement DNS natif, configuration du contrôle parental limitée et manuelle, pas d’application mobile, courbe d’apprentissage plus raide pour qui veut une protection complète avec chiffrement.

### AdGuard Home

**Avantages :** chiffrement DNS natif dans les deux sens, interface moderne, règles par client très fines, contrôle parental intégré, paquet natif OpenWrt, entièrement gratuit malgré une société éditrice commerciale.

**Inconvénients :** nécessite malgré tout un matériel à héberger et maintenir, pas d’application mobile dédiée, dépend des mises à jour d’une entreprise privée plutôt que d’un projet purement communautaire.

### NextDNS

**Avantages :** aucune installation ni maintenance, applications mobiles natives avec synchronisation, protection identique sur tous les réseaux, quota gratuit généreux qui ne coupe jamais l’accès à internet, tarifs Pro accessibles.

**Inconvénients :** dépendance à un service tiers et à sa disponibilité, entité légale incorporée aux États-Unis malgré des fondateurs français, quota de requêtes gratuites à surveiller pour les foyers très connectés, moins de transparence sur les listes de blocage propriétaires que sur des listes communautaires ouvertes.

## Guide de migration : changer d’outil sans coupure de service

Changer de solution de blocage DNS ne demande ni de tout réinstaller dans l’urgence ni de couper l’accès internet du foyer pendant l’opération. La méthode la plus sûre consiste à faire cohabiter l’ancien et le nouvel outil quelques jours avant de basculer complètement, quelle que soit la direction de la migration (Pi-hole vers AdGuard Home, un outil auto-hébergé vers NextDNS, ou l’inverse).

1. Exportez vos listes de blocage personnalisées et vos règles locales depuis l’outil actuel avant toute manipulation. Pi-hole et AdGuard Home proposent tous deux un export de configuration au format texte ou YAML depuis leur interface.
2. Installez le nouvel outil sur une adresse IP différente de l’ancien, sans encore toucher au DHCP ni à la box internet. Les deux solutions doivent pouvoir tourner en parallèle sur le réseau.
3. Recréez ou importez vos listes de blocage et vos règles personnalisées sur le nouvel outil, en ajoutant les mêmes listes communautaires que celles utilisées précédemment.
4. Configurez manuellement un seul appareil de test (un ordinateur portable suffit) pour pointer vers le nouveau résolveur, sans modifier encore les réglages globaux du réseau.
5. Vérifiez la résolution et le blocage effectif avec `dig` ou`nslookup` sur quelques domaines publicitaires connus, et confirmez que le chiffrement DNS fonctionne si c’est l’un des objectifs de la migration.
6. Une fois le test concluant, modifiez le DNS distribué par le DHCP de votre box ou routeur pour pointer l’ensemble du réseau vers le nouvel outil.
7. Laissez l’ancien outil actif et joignable pendant 48 heures au minimum, en solution de repli, avant de l’arrêter ou de le désinstaller définitivement.
8. Si vous migrez vers NextDNS, installez l’application mobile sur les appareils itinérants et activez la synchronisation pour retrouver la même politique de filtrage hors du réseau domestique.

Cette approche progressive évite le scénario le plus fréquent lors d’une migration ratée : un DHCP mal reconfiguré qui laisse une partie des appareils sans résolution DNS fonctionnelle, ou une liste de blocage mal recopiée qui casse l’accès à un service légitime.

## Au-delà du DNS : quand passer à un pare-feu réseau complet

Le blocage DNS a une limite structurelle qu’aucun des trois outils ne peut dépasser seul : il n’agit qu’au niveau de la résolution de noms, pas sur l’ensemble du trafic réseau. Une application qui contacte directement une adresse IP sans passer par une résolution DNS classique, ou un appareil configuré pour utiliser un DNS chiffré tiers en dur (certains téléphones récents imposent leur propre DoH par défaut), peut contourner Pi-hole, AdGuard Home ou NextDNS.

Pour un filtrage réellement exhaustif, l’étape suivante consiste à passer par un pare-feu réseau complet capable d’inspecter et de bloquer au niveau des paquets plutôt qu’au niveau des noms de domaine. Nous avons consacré un comparatif détaillé à ce sujet dans notre article pfSense vs OPNsense, qui reste la référence pour qui veut aller plus loin que le simple filtrage DNS avec un routeur ou un boîtier dédié. Ces solutions se combinent d’ailleurs très bien avec AdGuard Home ou Pi-hole, souvent installés comme service DNS interne derrière le pare-feu plutôt qu’en remplacement de celui-ci.

## Le verdict 2026 : quel bloqueur DNS choisir

Sur les critères purement techniques, AdGuard Home ressort comme le choix le plus complet pour qui accepte d’auto-héberger : chiffrement natif dans les deux sens, interface moderne, règles par client détaillées, contrôle parental intégré, le tout gratuitement et sans composant tiers à ajouter. C’est aujourd’hui la recommandation par défaut pour un nouveau déploiement chez un utilisateur technique.

Pi-hole garde toute sa pertinence pour qui privilégie un projet purement communautaire, une empreinte mémoire minimale et un historique de dix ans de stabilité, quitte à ajouter Unbound pour obtenir un chiffrement équivalent à celui d’AdGuard Home. NextDNS, de son côté, s’impose dès que la mobilité ou l’absence totale de maintenance priment sur la souveraineté complète des données, avec un tarif Pro à 1,99 € par mois qui reste accessible pour lever la limite de 300 000 requêtes du plan gratuit.

Le choix se résume finalement à une question d’arbitrage plutôt qu’à un vainqueur unique. Si vous avez déjà un serveur ou un Raspberry Pi qui tourne chez vous, AdGuard Home l’emporte. Si vous voulez zéro matériel et une protection identique sur votre téléphone en déplacement, NextDNS l’emporte. Si vous voulez la solution la plus documentée par la communauté depuis le plus longtemps, Pi-hole l’emporte.

## Foire aux questions

### Pi-hole, AdGuard Home ou NextDNS : lequel bloque le plus de publicités ?

À listes de blocage identiques, l’efficacité de blocage est quasiment équivalente, car les trois outils appliquent le même principe de filtrage par domaine. NextDNS ajoute des sources de détection propriétaires en plus des listes communautaires, ce qui peut légèrement améliorer la couverture contre le phishing et les domaines nouvellement créés.

### Le blocage DNS ralentit-il la connexion internet ?

