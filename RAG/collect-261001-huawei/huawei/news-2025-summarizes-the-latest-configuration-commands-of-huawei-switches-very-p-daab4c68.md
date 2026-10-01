---
id: collect-261001-huawei/huawei/news-2025-summarizes-the-latest-configuration-commands-of-huawei-switches-very-p-daab4c68
title: "Activer le service Telnet"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "sol"]
source: docs/RAG/collect-261001-huawei/news-2025-summarizes-the-latest-configuration-commands-of-huawei-switches-very-p-daab4c68.md
source_anchor: ""
source_lines: [1, 81]
sha256: d1d0acc9ec1d6dff2cbf9b7f02044f62ef8f1460d04f7588210e9141605c5b10
---

# Activer le service Telnet

Cet article résume les dernières commandes de configuration des commutateurs Huawei pour 2025. Cet article couvre deux parties: 1.Exemples de configuration des commutateurs HuaweiAllons voir.
1. Commandes de configuration de la catégorie 6 du commutateur Huawei
1. Entrée des vues de configuration:
affichage système → Entrez affichage système (mode de configuration).
sysname SW1 → Modifier le nom du commutateur.
fuseau horaire BJ ajouter 08:00:00 → Mettez le fuseau horaire (heure de Pékin).
afficher la configuration actuelle → Afficher la configuration en cours.
2. VLAN et configuration des ports (fonctions de base)
1. IP de gestion et connexion à distance
Interface Vlanif 10 # Entrez le VLAN de gestion (pas VLAN 1 est recommandé)
Adresse IP 172.16.1.1 255. Je vous en prie.255.255.0 # Configurer l' adresse IP de gestion
# Activer le service Telnet
Interface utilisateur vty 0 4
mode d'authentification aaa # Forcer l'authentification par nom d'utilisateur et mot de passe aaaprotocol telnet entrant # Activer le protocole telnet
mot de passe d'administrateur local-utilisateur Admin@123 # Le mot de passe doit contenir des lettres majuscules, minuscules et numériques
Telnet de type service administrateur utilisateur local
2Sécurité portuaire et liaison MAC
Interface GigabitEthernet 0/0/5port-sécurité activé # Activé
Port-security max-mac-num 2 # Limiter le nombre d'adresses MAC (pour empêcher les appareils non autorisés)
# Lier automatiquement l'adresse MAC du premier accès
3. Configuration des fonctions avancées (agrégation et service DHCP)
1. Aggrégation des liens (Eth-Trunk)
interface Eth-Trunk 1 # Créer un groupe d'agrégation
le port principal GigabitEthernet de 0/0/3 à 0/0/5 # Ajouter les ports membres
type de liaison de port trunkload-balance src-dst-mac
Équilibrage de la charge basé sur le MAC source/destination
2. Service DHCP
DHCP activé # DHCP activé globalement
Puits d'adresses invité # Créer un pool d'adresses
le réseau 192.168.30.0 masque 24
liste de passerelle 192.168.30.1
interface Vlanif 30
dhcp sélectionne global # Apply pool d'adresses
IV. Durcissement de la sécurité et ACL
1. Liste de contrôle d' accès (ACL)
ACL 3000 # ACL avancée (3000-3999)
La règle 5 est de refuser la source IP 192.168.10Je suis désolée.0.0.255 à destination 192.168.20Je suis désolée.0.0.255 # Interdire l' accès au réseau
Interface Vlanif 10
Filtreur de trafic entrant acl 3000 # Appliquer la politique de l'ACL
2. Droits de la console et de l' utilisateur
Console d'interface utilisateur 0
mot de passe de mode d'authentification
définir le mot de passe d'authentification Huawei@123 # mot de passe de la console
niveau de privilège de l'opérateur utilisateur local 3 # Niveaux de privilège de l'utilisateur (0-15, 3 pour l'administrateur)
5. Commandes courantes pour le dépannage
Résumé de l'interface d'affichage
afficher l'adresse mac
affichage de l'eth-trunk 1
Le ping 192.168.1.1
Réinitialiser l'interface des compteurs GigabitEthernet 0/0/1
6. Routage IP et fonctions de couche 3 1. Le routage inter-VLAN utilise VLANIF pour atteindre la communication sous-réseau:
Adresse IP 192.168.10.1 24
Interface Vlanif 20
Adresse IP 192.168.20.1 24
# Après que le PC ait défini la passerelle correspondante, la communication VLAN croisée peut être réalisée
2. Route statique pointant vers le routeur ou le pare-feu principal:
IP-route statique 0.0.0Je suis désolée.0.0.0 192. Je vous en prie.168.100.254 # Route par défaut
2. Exemple de configuration: les VLAN communiquent par l'interface VLANIF. Voici un exemple simple:
Comme le montre la figure précédente, dans un réseau d'entreprise, l'utilisateur 1 et l'utilisateur 2 ont le même service mais appartiennent à des VLAN différents et sont situés sur des segments de réseau différents.Il est maintenant nécessaire de permettre la communication entre l'utilisateur 1 et l'utilisateur 2Ce scénario est le plus courant dans les réseaux de bureaux.
Stratégie de configuration: créer un VLAN, ajouter une interface au VLAN, créer une interface VLANIF et attribuer une adresse IP à l'interface.Jetons un coup d'œil à la configuration.
1. Créez un VLAN
[HUAWEI] sysname Switch //Nommer le commutateur Switch
[Commutateur] Vlan lot 10 20 //Créer le VLAN 10 vlan2
2. Configurer les interfaces pour rejoindre les VLAN
[Commutateur] interface gigabitethernet 0/0/1 // Entrez le port 0/0/1
[Switch-GigabitEthernet0/0/1] accès au port de type liaison //Mettre le port en mode d'accès
[Switch-GigabitEthernet0/0/1] port par défaut vlan 10 //Assigner le port à VLAN 10
[Switch-GigabitEthernet0/0/1] déconnecter //Sortir
[Commutateur] interface Gigabitethernet 0/0/2 // Entrez le port 0/0/2
[Switch-GigabitEthernet0/0/2] accès au port de type liaison //Mettre le port en mode d'accès
[Switch-GigabitEthernet0/0/2] port par défaut vlan 20 //Assigner le port à VLAN 10
[Switch-GigabitEthernet0/0/2] déconnecter //Sortir
L' étape clé est ici.
3. Configurer l' adresse IP de l' interface VLANIF
[Switch] interface vlanif 10 //Entrer la vue de l'interface logique du VLAN 10
[Commutateur-Vlanif10] adresse IP 10.10.10.2 24 //Configurer une adresse IP pour elle
[Commutateur-Vlanif10] arrête //Sorte
[Switch] interface vlanif 20 //Entrez la vue de l'interface logique du VLAN 20
[Switch-Vlanif20] adresse IP 10.10.20.2 24 //Configurer une adresse IP pour elle
[Switch-Vlanif20] arrête //Sorte
Une fois la configuration terminée, leurs interfaces ont des adresses IP et peuvent communiquer entre elles.
