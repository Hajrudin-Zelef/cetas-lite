---
id: collect-260926-mikrotik/mikrotik/blog-fr-comment-depanner-mikrotik-routeros-2025-95e700f1-2
title: "Exemple : Tester si le pare-feu bloque le trafic"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/blog-fr-comment-depanner-mikrotik-routeros-2025-95e700f1.md
source_anchor: ""
source_lines: [130, 267]
sha256: ecf91846241608b274db0dbaf59756824e266de5434fb3f7dff06c099ee38788
---

# Exemple : Tester si le pare-feu bloque le trafic

Utilisez l'outil Torch (`/tool torch`) dans Winbox ou WebFig pour surveiller le trafic en temps réel par adresse IP, protocole ou port. Pour une analyse à plus long terme, configurez le flux de trafic (`/ip traffic-flow`) et exportez les données vers un collecteur pour des rapports détaillés sur l'utilisation de la bande passante.

**Introduction à la Section :** Les problèmes de connectivité sont parmi les problèmes les plus courants rencontrés par les administrateurs réseau. Cette section se concentre sur le diagnostic et la résolution méthodique de divers types de problèmes de connexion dans les environnements MikroTik.

**Explication :** Les problèmes de connectivité peuvent provenir de problèmes de couche physique, d'erreurs de configuration, de problèmes de routage ou de perturbations de service. Une approche systématique permet d'identifier la cause spécifique.

**Détails Techniques :** Nous explorerons les problèmes de connectivité courants à différentes couches réseau, des connexions physiques aux services de niveau application, avec des commandes et configurations MikroTik spécifiques pour résoudre chaque type de problème.

**Avantages et Applications :**

- Restaurer rapidement l'accès Internet pour les utilisateurs
- Résoudre les problèmes de connectivité réseau interne
- Corriger les problèmes de connexion VPN
- Résoudre les échecs de résolution DNS
- Dépanner les problèmes de routage et de passerelle

**Instructions Étape par Étape pour le Dépannage de la Connectivité :**

1. 
**Vérifications de la Couche Physique :**
  - Vérifier l'état de l'interface et la détection de lien :
plaintext```
/interface ethernet print
```
  - Rechercher les erreurs ou les rejets d'interface :
plaintext```
/interface ethernet print stats
```
  - Pour les liaisons sans fil, examiner la force du signal et le CCQ :
plaintext```
/interface wireless registration-table print
```
  - Réinitialiser les interfaces problématiques :
plaintext```
/interface ethernet reset-mac-address ether1
```
2. Vérifier l'état de l'interface et la détection de lien :
3. 
**Vérification de la Configuration IP :**
  - Vérifier les adresses IP sur les interfaces :
plaintext```
/ip address print
```
  - Vérifier le fonctionnement du client DHCP :
plaintext```
/ip dhcp-client print
```
  - Tester l'accessibilité de la passerelle :
plaintext```
/ping [adresse-ip-passerelle] count=5
```
  - Examiner la table ARP :
plaintext```
/ip arp print
```
4. Vérifier les adresses IP sur les interfaces :
5. 
**Dépannage du Routage :**
  - Vérifier les entrées de la table de routage :
plaintext```
/ip route print
```
  - Rechercher les conflits de routage ou les routes manquantes :
plaintext```
/ip route print detail
```
  - Tester des chemins de route spécifiques :
plaintext```
/ping 8.8.8.8 routing-table=main count=5
```
  - Examiner le processus de sélection des routes :
plaintext```
/ip route get 8.8.8.8
```
6. Vérifier les entrées de la table de routage :
7. 
**Problèmes de Résolution DNS :**
  - Vérifier la configuration du serveur DNS :
plaintext```
/ip dns print
```
  - Tester la résolution DNS :
plaintext```
/tool dns-lookup name=google.com server=8.8.8.8
```
  - Vérifier le cache DNS :
plaintext```
/ip dns cache print
```
  - Vider le cache DNS si nécessaire :
plaintext```
/ip dns cache flush
```
8. Vérifier la configuration du serveur DNS :
9. 
**Vérification du Pare-feu et du NAT :**
  - Vérifier les règles de pare-feu qui pourraient bloquer le trafic :
plaintext```
/ip firewall filter print
```
  - Vérifier la configuration NAT :
plaintext```
/ip firewall nat print
```
  - Désactiver temporairement le pare-feu pour les tests (à utiliser avec prudence) :
plaintext```
/ip firewall filter disable [find]
```
  - Surveiller le suivi des connexions :
plaintext```
/ip firewall connection print where dst-address=ip-problématique
```
10. Vérifier les règles de pare-feu qui pourraient bloquer le trafic :

**Résumé de la Section :** Des problèmes de connectivité peuvent survenir à plusieurs couches de la pile réseau. En suivant une approche systématique de la couche physique vers le haut, vous pouvez identifier et résoudre efficacement la cause première des problèmes de connexion dans votre environnement MikroTik.

**Mini-FAQ :**

Cela indique souvent un problème de résolution DNS. Si vous pouvez pinguer des adresses IP mais pas des noms de domaine, vérifiez la configuration de votre serveur DNS, recherchez les règles de pare-feu bloquant le trafic DNS (port UDP/TCP 53) et assurez-vous que vos serveurs DNS sont joignables depuis le routeur.

Tout d'abord, déterminez ce que les appareils qui fonctionnent et ceux qui ne fonctionnent pas ont en commun. Vérifiez s'ils se trouvent sur des interfaces, des VLAN ou des sous-réseaux IP différents. Vérifiez que le DHCP fonctionne correctement pour le sous-réseau affecté et examinez les règles de pare-feu qui pourraient filtrer le trafic en fonction des adresses source ou des adresses MAC.

**Introduction à la Section :** Les problèmes de performance peuvent être plus difficiles à diagnostiquer que les pannes complètes car le réseau fonctionne toujours mais pas à des niveaux optimaux. Cette section se concentre sur l'identification et la résolution des goulots d'étranglement de performance dans les routeurs MikroTik.

**Explication :** Les problèmes de performance se manifestent généralement par des vitesses lentes, une latence élevée, une perte de paquets ou une connectivité intermittente. Ces problèmes peuvent découler de contraintes de ressources, d'inefficacités de configuration ou de facteurs externes.

**Détails Techniques :** Nous explorerons des méthodes pour identifier les goulots d'étranglement des ressources, optimiser les configurations et implémenter des fonctionnalités améliorant les performances dans RouterOS.

**Avantages et Applications :**

- Améliorer le débit réseau global
- Réduire la latence pour les applications sensibles au temps
- Éliminer la perte de paquets et l'instabilité de la connexion
- Optimiser l'utilisation des ressources
- Améliorer la qualité de service pour le trafic critique

**Instructions Étape par Étape pour le Dépannage des Performances :**

