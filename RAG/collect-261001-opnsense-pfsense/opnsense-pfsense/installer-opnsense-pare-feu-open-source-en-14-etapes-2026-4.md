---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026-4
title: "Sous Linux, décompression de l'image téléchargée"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [199, 264]
sha256: 6e85e793d83ff028b99c66b4e41e0ec6b6b8ccd744a010f25974a7c0d1858223
---

# Sous Linux, décompression de l'image téléchargée

```
# Exemple de configuration d'un groupe de passerelles (System > Gateways > Group)
Group Name: WAN_FAILOVER
Gateway: WAN_FIBRE   Tier: 1   (priorité principale)
Gateway: WAN_4G       Tier: 2   (bascule de secours)
Trigger Level: Member Down
# Règle de pare-feu utilisant ce groupe comme passerelle
Firewall > Rules > LAN
Action: Pass, Gateway: WAN_FAILOVER
```
Ce mécanisme de bascule (failover) reste transparent pour les utilisateurs du réseau local : seules les connexions actives au moment de la coupure sont interrompues, les nouvelles connexions empruntant automatiquement le lien de secours en quelques secondes. Pour une répartition de charge plutôt qu’un simple basculement, il suffit d’attribuer le même niveau de priorité (tier) aux deux passerelles : OPNsense distribue alors les nouvelles connexions entre les deux liens selon un mécanisme de round-robin pondéré.

## OPNsense vs pfSense : quelles différences concrètes en 2026

La question revient systématiquement chez les administrateurs qui migrent d’une solution vers l’autre. Les deux projets partagent une origine commune (OPNsense est un fork de pfSense datant de 2015) mais ont divergé sur la gouvernance, la licence et le rythme de publication.

| Critère | OPNsense | pfSense | 
|---|---|---|
| Éditeur | Deciso B.V. (Pays-Bas) | Netgate (États-Unis) | 
| Dernière version gratuite | 26.7.3 (27 août 2026) | pfSense CE 2.8.1 (septembre 2025) | 
| Édition commerciale | Business Edition 26.4.2, même base de code open source | pfSense Plus, fonctions réservées aux appliances Netgate | 
| Licence | Open source BSD sur l’ensemble du code | CE open source, Plus partiellement propriétaire | 
| Base système | FreeBSD 15.1 | FreeBSD (version antérieure) | 
| Rythme de publication | Plusieurs versions mineures par an | Une à deux versions majeures par an | 

Pour un déploiement neuf sans contrainte de compatibilité avec du matériel Netgate existant, OPNsense présente l’avantage d’un code entièrement ouvert et d’un cycle de correctifs plus rapide. pfSense conserve un écosystème d’appliances certifiées plus large et reste pertinent pour les organisations déjà investies dans du matériel Netgate.

## Édition Business, appliances et coûts de support

Deciso commercialise des appliances matérielles préinstallées et des contrats de support autour de l’édition Business (26.4.2 au 14 août 2026). Les tarifs précis dépendent du modèle d’appliance et du niveau de support choisi, et sont publiés directement sur le site de Deciso plutôt que dans un catalogue tarifaire fixe. Pour une utilisation personnelle ou associative, la Community Edition installée sur du matériel générique reste totalement gratuite et couvre l’immense majorité des besoins, y compris Suricata, WireGuard et la haute disponibilité.

## Cinq pièges fréquents à éviter

La majorité des échecs d’installation ou des pannes réseau après mise en production d’OPNsense proviennent d’une poignée d’erreurs récurrentes, largement documentées sur les forums de la communauté.

1. **Inversion des interfaces WAN et LAN** : brancher le câble du fournisseur d’accès sur l’interface configurée en LAN expose immédiatement le réseau interne à Internet sans filtrage. Vérifiez toujours physiquement chaque câble avant l’assignation finale.
2. **Verrouillage total après une règle de pare-feu trop stricte** : supprimer par erreur la règle autorisant l’accès HTTPS depuis le LAN vers l’interface de gestion coupe l’accès à l’interface web. La console série ou un accès physique au clavier reste alors le seul recours pour rétablir une règle d’urgence.
3. **Conflit de serveurs DHCP** : laisser le DHCP de la box internet actif en parallèle de celui d’OPNsense provoque des attributions d’adresses IP incohérentes et des pertes de connectivité intermittentes. Désactivez systématiquement le DHCP de la box en amont.
4. **Décharge matérielle (offloading) incompatible avec certaines cartes réseau** : certains contrôleurs Realtek ou des cartes réseau virtuelles mal configurées provoquent des instabilités si les fonctions de décharge de checksum (checksum offloading) ou TSO restent actives. Désactivez-les sous Interfaces → Settings en cas de symptômes de latence ou de paquets perdus.
5. **Interception TLS pour l’IDS sans PKI maîtrisée** : tenter de déchiffrer le trafic HTTPS pour l’inspection Suricata sans déployer correctement un certificat d’autorité de confiance sur tous les postes clients génère des erreurs de certificat en cascade côté utilisateurs. Cette fonctionnalité doit rester réservée aux environnements d’entreprise avec un service informatique capable de gérer un magasin de certificats centralisé.

## Résultats attendus après une installation réussie

Une fois toutes les étapes complétées, le tableau de bord (Dashboard) de l’interface web affiche l’état de santé du système : charge CPU, utilisation mémoire, état des interfaces réseau, et un flux des dernières entrées de journal. Voici un exemple représentatif de ce que vous devez observer sur une installation fonctionnelle.

```
État général attendu après installation :
Interfaces:
  WAN (igb0)   [UP]   IP: 203.0.113.45/24    Trafic entrant/sortant actif
  LAN (igb1)   [UP]   IP: 192.168.1.1/24     4 baux DHCP actifs
Services:
  unbound (DNS resolver)     [running]
  dhcpd (LAN)                [running]
  wireguard                  [running]  1 instance, 2 peers connectés
  suricata (IDS mode)        [running]  Interface: WAN
Firewall:
  Règles actives: 12
  Connexions suivies (states): 340/500000
  Dernière mise à jour firmware: 26.7.3 (27 août 2026)
```
Si votre tableau de bord ressemble à cet exemple, avec des interfaces “UP”, des services actifs sans erreur rouge et un nombre de connexions suivies cohérent avec votre trafic réel, l’installation est opérationnelle et prête pour un usage en production.

## Dépannage : 8 problèmes courants et leurs solutions

Voici les incidents les plus fréquemment rapportés par la communauté OPNsense, avec la cause probable et la correction associée.

