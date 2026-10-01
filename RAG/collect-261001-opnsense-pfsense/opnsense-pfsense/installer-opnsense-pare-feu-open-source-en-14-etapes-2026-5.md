---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026-5
title: "Sous Linux, décompression de l'image téléchargée"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel", "open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [265, 344]
sha256: 1c263a5aee6041e09ea53a93cde4404a5435a89d87d304f10b8be5e63117413a
---

# Sous Linux, décompression de l'image téléchargée

1. **Impossible d’accéder à l’interface web après configuration** : vérifiez que le poste client est bien sur le même sous-réseau que le LAN, videz le cache DNS local, et confirmez qu’aucune règle de pare-feu ne bloque le port 443 sur l’interface LAN elle-même (Firewall → Rules → LAN).
2. **Aucune connexion Internet après l’installation** : contrôlez le type de WAN configuré (DHCP, statique, PPPoE) correspond bien à celui exigé par votre fournisseur d’accès, et vérifiez sous Interfaces → Overview que l’interface WAN a bien obtenu une adresse IP publique.
3. **Le tunnel WireGuard ne s’établit pas** : confirmez que le port UDP 51820 est ouvert sur le pare-feu WAN, que les clés publiques sont correctement échangées entre pairs, et que l’horloge système du serveur et du client ne dérive pas de plus de quelques minutes.
4. **Suricata consomme un CPU anormalement élevé** : réduisez le nombre de catégories de règles actives, passez en mode IDS plutôt qu’IPS si le matériel est limité, et vérifiez que le nombre de threads alloués correspond au nombre de cœurs physiques disponibles.
5. **Bascule CARP intempestive entre les deux nœuds** : ce symptôme provient généralement d’un lien de synchronisation pfsync saturé ou instable ; dédiez une interface réseau distincte, idéalement à 1 Gbps minimum, uniquement pour ce trafic de synchronisation.
6. **Échec de mise à jour du firmware** : videz l’espace disque disponible (une mise à jour nécessite généralement 2 à 4 Go d’espace libre), vérifiez la connectivité vers les dépôts OPNsense, puis relancez via`opnsense-update -c` pour un diagnostic en ligne de commande.
7. **Certificat Let’s Encrypt (ACME) qui ne se renouvelle pas** : vérifiez que le port 80 reste accessible depuis Internet pour la validation HTTP-01, ou basculez vers une validation DNS-01 si le port 80 est bloqué par ailleurs.
8. **Débit réseau anormalement bas malgré un matériel dimensionné** : désactivez le mode IPS de Suricata pour tester, vérifiez les statistiques d’interruptions (interrupts) au niveau du système, et confirmez que les cartes réseau utilisent bien un pilote FreeBSD natif (igb, em, ix) plutôt qu’un pilote générique.

## Conseils avancés pour aller plus loin

Une fois l’installation de base stabilisée, plusieurs ajustements permettent de tirer davantage parti d’OPNsense en environnement exigeant. Le plugin os-netflow, couplé à un outil de visualisation comme Insight, donne une vision granulaire de la consommation de bande passante par machine et par protocole, utile pour détecter une exfiltration de données ou un poste compromis participant à un botnet.

Pour la segmentation réseau, la création de VLAN directement dans OPNsense (Interfaces → Other Types → VLAN) permet d’isoler un réseau invité, un réseau IoT et un réseau de production sur une seule carte réseau physique reliée à un commutateur administrable compatible 802.1Q. Chaque VLAN reçoit ensuite son propre jeu de règles de pare-feu, empêchant par exemple une caméra IP compromise de communiquer avec un poste de travail du réseau de production.

Enfin, pour les environnements virtualisés (Proxmox, KVM, VMware ESXi), privilégiez les pilotes VirtIO pour les interfaces réseau plutôt que l’émulation d’une carte Intel générique : le gain de performance est significatif, à condition de vérifier au préalable que le noyau FreeBSD embarqué dans votre version d’OPNsense supporte correctement le pilote VirtIO de votre hyperviseur, ce qui est le cas pour toutes les versions de la série 26.x.

## Projet complet : pare-feu domestique sécurisé avec VPN et IDS

Pour synthétiser l’ensemble du tutoriel, voici l’architecture complète d’un projet fonctionnel que vous pouvez reproduire de bout en bout sur un mini-PC à deux cartes réseau.

1. Installation d’OPNsense 26.7.3 sur un mini-PC équipé de 8 Go de RAM, 120 Go de SSD et deux ports Ethernet Intel
2. Interface WAN en DHCP côté box opérateur placée en mode bridge, interface LAN en `192.168.1.1/24` avec DHCP activé
3. Résolveur Unbound configuré en mode récursif avec enregistrement automatique des baux DHCP
4. VLAN dédié pour les objets connectés (`192.168.20.0/24` ), isolé du réseau principal par des règles de pare-feu strictes
5. Plugin os-wireguard installé, tunnel VPN nomade actif sur le port UDP 51820 pour l’accès distant sécurisé
6. Plugin os-ids avec Suricata en mode IDS sur l’interface WAN, jeu de règles ET Open activé
7. Plugin os-acme-client configuré pour renouveler automatiquement un certificat Let’s Encrypt sur l’interface de gestion
8. Sauvegarde automatique hebdomadaire du fichier `config.xml` vers un stockage externe

```
# Résumé des règles de pare-feu essentielles du projet complet
# Sur l'interface WAN
Allow UDP 51820 (WireGuard) from any to WAN address
# Sur l'interface WireGuard
Allow TCP/UDP from 10.10.10.0/24 to 192.168.1.0/24
# Sur le VLAN IoT
Block from 192.168.20.0/24 to 192.168.1.0/24
Allow from 192.168.20.0/24 to WAN (Internet uniquement)
# Sur le LAN principal
Allow LAN net to any (règle par défaut)
```
Ce projet couvre la quasi-totalité des besoins d’un particulier exigeant ou d’une très petite structure : accès distant sécurisé, isolation des objets connectés, détection d’intrusion et chiffrement TLS automatisé, le tout sur du matériel accessible et sans licence à payer.

### Foire aux questions

**OPNsense est-il gratuit ?**

Oui, la Community Edition est entièrement gratuite et open source. Seule l’édition Business, destinée aux entreprises souhaitant un support contractuel avec Deciso, implique un coût.

**Quelle est la différence entre OPNsense et pfSense ?**

OPNsense est un fork de pfSense créé en 2015. Les deux projets partagent une origine commune mais divergent sur la gouvernance, la licence (OPNsense reste intégralement open source) et le rythme de publication, plus soutenu chez OPNsense.

**Combien de RAM faut-il pour faire tourner OPNsense correctement ?**

4 Go suffisent pour un usage basique (pare-feu, DHCP, DNS). Comptez 8 Go minimum si vous activez Suricata en IDS/IPS, et 16 Go ou plus pour un usage combinant VPN, IDS et haute disponibilité.

**Peut-on installer OPNsense sur une machine virtuelle ?**

Oui, OPNsense fonctionne sans problème sous Proxmox, KVM ou VMware ESXi, à condition d’utiliser des pilotes réseau VirtIO pour de meilleures performances plutôt que l’émulation d’une carte réseau générique.

**OPNsense supporte-t-il IPv6 ?**

Oui, et le support IPv6 a été spécifiquement amélioré dans la série 26.7 sortie en juillet 2026, notamment pour la délégation de préfixe (DHCPv6-PD) reçue depuis certains fournisseurs d’accès.

**Faut-il des cartes réseau spécifiques ?**

OPNsense fonctionne avec la majorité des cartes réseau x86-64 reconnues par FreeBSD, mais les modèles Intel utilisant les pilotes igb, em ou ix offrent la meilleure stabilité et les meilleures performances, en particulier sous forte charge.

**Comment migrer une configuration pfSense existante vers OPNsense ?**

Il n’existe pas d’outil de migration automatique officiel entre les deux projets en raison de leur divergence de code depuis 2015. La reconstruction manuelle des règles de pare-feu et des services à partir d’une exportation de configuration reste la méthode recommandée par la communauté.

**À quelle fréquence faut-il mettre à jour OPNsense ?**

Le projet publie des versions mineures plusieurs fois par an (trois rien que pour la branche 26.7 entre juillet et août 2026). Une vérification hebdomadaire des mises à jour, et une application rapide après toute annonce de faille de sécurité sur le forum officiel, est recommandée.

### Related Coverage

