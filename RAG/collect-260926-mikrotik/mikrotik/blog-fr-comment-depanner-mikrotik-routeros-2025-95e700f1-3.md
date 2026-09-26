---
id: collect-260926-mikrotik/mikrotik/blog-fr-comment-depanner-mikrotik-routeros-2025-95e700f1-3
title: "Exemple : Tester si le pare-feu bloque le trafic"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/blog-fr-comment-depanner-mikrotik-routeros-2025-95e700f1.md
source_anchor: ""
source_lines: [268, 374]
sha256: 9e5089051eb9de65b17cb6f7f258dad03c6fac82a3bc8de4abb6d8a3d8b43160
---

# Exemple : Tester si le pare-feu bloque le trafic

1. 
**Identifier les Goulots d'Étranglement des Ressources :**
  - Vérifier l'utilisation du CPU :
plaintext```
/system resource cpu print
```
  - Surveiller l'utilisation du CPU par processus :
plaintext```
/tool profile
```
  - Examiner l'utilisation de la mémoire :
plaintext```
/system resource print
```
  - Vérifier l'utilisation et la santé du disque :
plaintext```
/disk print
```
2. Vérifier l'utilisation du CPU :
3. 
**Analyser les Modèles de Trafic :**
  - Utiliser Torch pour identifier le trafic gourmand en bande passante :
plaintext```
/tool torch interface=ether1 ip-protocol=any
```
  - Examiner la table des connexions pour un grand nombre de connexions :
plaintext```
/ip firewall connection print count-only
```
  - Vérifier la congestion de l'interface :
plaintext```
/interface monitor-traffic ether1 once
```
  - Identifier les principaux consommateurs de bande passante avec le flux de trafic :
plaintext```
/ip traffic-flow print
```
4. Utiliser Torch pour identifier le trafic gourmand en bande passante :
5. 
**Optimiser la Configuration du Pare-feu :**
  - Déplacer les règles fréquemment mises en correspondance en haut :
plaintext```
/ip firewall filter print stats
```
  - Utiliser les aides au suivi des connexions de manière appropriée :
plaintext```
/ip firewall connection tracking print
```
  - Implémenter le fasttrack pour le trafic de confiance :
plaintext```
/ip firewall filter add chain=forward action=fasttrack-connection connection-state=established,related comment="FastTrack"
```
  - Limiter les débits de connexion pour la protection potentielle contre les attaques DoS :
plaintext```
/ip firewall filter add chain=input protocol=tcp dst-port=22 connection-limit=3,32 action=drop comment="Protection anti-brute force SSH"
```
6. Déplacer les règles fréquemment mises en correspondance en haut :
7. 
**Mettre en Œuvre la Qualité de Service (QoS) :**
  - Identifier les types de trafic nécessitant une priorisation
  - Créer des types de files d'attente pour différentes classes de trafic :
plaintext```
/queue type add name=streaming-video kind=pcq pcq-classifier=dst-address pcq-rate=10M
```
  - Implémenter des files d'attente simples pour la gestion de la bande passante :
plaintext```
/queue simple add name=limit-youtube target=192.168.1.0/24 dst-address=youtube-ip-ranges queue=streaming-video max-limit=20M/20M
```
  - Ou utiliser des arbres de files d'attente plus avancés pour des scénarios complexes
8. 
**Optimiser les Performances Sans Fil (si applicable) :**
  - Sélectionner les fréquences optimales avec l'analyse spectrale :
plaintext```
/interface wireless spectral-scan wlan1
```
  - Ajuster la largeur de canal de manière appropriée
  - Implémenter une liste d'accès sans fil pour empêcher les connexions non autorisées
  - Activer la compression sans fil si bénéfique :
plaintext```
/interface wireless set wlan1 compression=yes
```
9. Sélectionner les fréquences optimales avec l'analyse spectrale :

**Résumé de la Section :** L'optimisation des performances nécessite une approche méthodique pour identifier les goulots d'étranglement et mettre en œuvre des solutions appropriées. En surveillant l'utilisation des ressources, en analysant les modèles de trafic et en optimisant les configurations, vous pouvez améliorer considérablement les performances et la fiabilité de votre routeur MikroTik.

**Mini-FAQ :**

Surveillez simultanément l'utilisation du CPU et le débit de l'interface. Si l'utilisation du CPU monte à près de 100 % alors que la bande passante reste inférieure à la capacité de l'interface, vous avez probablement un goulot d'étranglement lié au CPU. Si les interfaces affichent une utilisation élevée constante proche de leur capacité maximale tandis que l'utilisation du CPU reste raisonnable, vous êtes confronté à des contraintes de bande passante.

Oui, de manière significative. Les routeurs MikroTik possèdent de nombreuses fonctionnalités puissantes, mais l'activation de services inutiles consomme des ressources. N'activez que les fonctionnalités dont vous avez réellement besoin, surtout sur du matériel d'entrée de gamme. Des services comme le proxy, le proxy web, le SNMP, les tests de bande passante et la journalisation étendue peuvent tous avoir un impact sur les performances lorsqu'ils ne sont pas nécessaires.

**Introduction à la Section :** De nombreux problèmes MikroTik proviennent d'erreurs de configuration ou de problèmes liés aux logiciels. Cette section se concentre sur l'identification et la résolution de ces types de problèmes.

**Explication :** Les problèmes de configuration peuvent aller de simples erreurs de syntaxe à des problèmes d'interaction complexes entre différentes fonctionnalités de RouterOS. Les problèmes logiciels peuvent inclure des bugs dans des versions spécifiques de RouterOS ou des conflits de paquets.

**Détails Techniques :** Nous explorerons des méthodes pour identifier les incohérences de configuration, résoudre les problèmes spécifiques à une version et mettre en œuvre les meilleures pratiques pour la gestion de la configuration.

**Avantages et Applications :**

- Éliminer les erreurs de configuration causant des problèmes réseau
- Résoudre les bugs RouterOS spécifiques à une version
- Mettre en œuvre des pratiques de configuration plus robustes
- Récupérer après des mises à jour échouées ou des configurations corrompues
- Maintenir des sauvegardes de configuration pour une récupération rapide

**Instructions Étape par Étape pour le Dépannage de la Configuration :**

