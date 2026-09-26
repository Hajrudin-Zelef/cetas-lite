---
id: collect-260926-mikrotik/mikrotik/remplacer-livebox-guide-de-connexion-fibre-directement-sur-un-routeur-voire-meme-78ebcdaf
title: "La fibre Orange à 8Gbps, sur un routeur MikroTik 10Gbps CCR2004, via un ONT SFP+"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["attention"]
source: docs/RAG/lot-mikrotik/forum/misc/remplacer-livebox-guide-de-connexion-fibre-directement-sur-un-routeur-voire-meme-78ebcdaf.md
source_anchor: ""
source_lines: [1, 40]
sha256: a712f882c3e66cca824af447c7a6fa148bd0853e6bbc7987dc73496d7057acba
---

# La fibre Orange à 8Gbps, sur un routeur MikroTik 10Gbps CCR2004, via un ONT SFP+

*Source : https://lafibre.info/remplacer-livebox/guide-de-connexion-fibre-directement-sur-un-routeur-voire-meme-en-2gbps/4644/*

**Contexte :** suivi du tutoriel ci-dessous. IPv4 publique OK, NAT OK, ping 8.8.8.8 et google.fr OK, mais accès à certains sites seulement (lafibre.info OK, Paypal/Google KO). Question : comment ajouter l'IPv6 — le tuto ne permet-il de récupérer qu'une IPv4 sur le réseau Orange ? Faut-il obligatoirement configurer l'IPv6 pour accéder à certains sites ?

**Tutoriel — configuration IPv4 pour Mikrotik CCR2004-1G-12S+2XS** (inspiré du post de Seb59, janvier 2017) :

**Organisation des ports :** sans logique particulière (inspiration Seb59) — les ports utilisés ne sont pas nécessairement adjacents, pour diluer la dissipation thermique et repousser le seuil d'intervention du refroidissement actif.

D'origine, le CCR2004 démarre en 192.168.88.1/24 sur l'interface eth1 1 Gbps. Le plus simple : s'y connecter avec WinBox et écraser certains réglages en une étape via Quick Set (en haut à gauche) :
- Adresse IP de la Gateway : 192.168.1.1/24 (Local Network / IP Address)
- Router Identity : Mikrotik CCR2004 (utile pour différencier le routeur des switchs 10 Gbps)
- Password : à changer impérativement (des petits malins tenteront de se connecter aussitôt)

Puis déconnecter/reconnecter WinBox. Le plus simple est de passer par SSH, moins sensible aux caractères accentués.

**Bridge pour les VLAN :** on utilise un bridge pour lier tous les VLAN nécessaires. *Mise à jour du 1er janvier 2023 :* la surcharge de l'adresse MAC de ce bridge permet l'envoi automatique d'un DUID DHCPv6 conforme aux attentes d'Orange.

On écrase les IP des interfaces d'accès. On crée une interface esclave sur le VLAN 832.

**Client DHCP :** définir les options nécessaires à la configuration du client DHCP pour l'accès Orange. *Mise à jour du 1er janvier 2023 :* convertir l'identifiant fti et le mot de passe via la page ad hoc de @Kgersen (jsfiddle.net/kgersen/3mnsc6wy) et un convertisseur ASCII→hex (rapidtables.com). Attention : la valeur de l'option 90 comprend bien 140 caractères en hexa.

On lancera le client DHCP un peu plus tard — d'abord le serveur DHCP côté LAN (interface ether12-LAN).

**Queues :** RouterOS permet de définir de nombreux types de queues (limitation/priorisation du trafic — cf. wiki.mikrotik.com/wiki/Manual:Queue). Rester simple : FIFO (premier arrivé, premier servi).

**Point critique :** si on lance le client DHCP requérant l'IP publique à Orange sans marquer CE trafic DHCP sortant au niveau de priorité COS=6 (selon la localisation géographique), cela ne fonctionnera pas.

Ajouter l'interface esclave vlan832-internet au bridge br-wan, sur lequel s'envoie la requête DHCP (à COS=6 forgé). On peut enfin lancer l'authentification DHCP. L'IP WAN apparaît alors sur le bridge br-wan dans WinBox (/IP/ADDRESS).

**DNS pour le LAN, puis firewall :** on a une adresse IP externe mais le NAT IPv4 n'est pas encore lancé. Créer des règles de firewall : d'abord ajouter à la liste « support » toutes les tranches d'IP utilisées — **très important**, car si la tranche IP LAN n'y est pas, le firewall coupe l'accès au routeur lui-même. Puis les règles firewall elles-mêmes (avec un traitement spécial pour les tentatives d'ouverture de shell SSH depuis l'interface WAN).

Une fois le firewall activé, lancer le NAT. Le routeur est connecté et route.

**Port forwarding :** exemple — serveur web HTTP port 80 sur l'IP LAN 192.168.1.5. Bien préciser l'interface d'entrée (`in-interface=br-wan`), sinon toutes les connexions sortantes vers un port identique seraient renvoyées vers le réseau interne.

**Fin :** désactiver Telnet et FTP. Le plus sûr : SSH, uniquement depuis l'intérieur du réseau.

**Annexe :** script de mise à jour toutes les 4 heures d'une blocklist d'IP à droper depuis le WAN (cf. post lafibre.info « script d'inclusion de blocklistes de moches »).
