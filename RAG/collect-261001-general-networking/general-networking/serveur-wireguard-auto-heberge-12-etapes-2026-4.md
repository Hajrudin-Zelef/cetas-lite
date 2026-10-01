---
id: collect-261001-general-networking/general-networking/serveur-wireguard-auto-heberge-12-etapes-2026-4
title: "Les blocs [Peer] des clients seront ajoutés ici à l'étape 7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/serveur-wireguard-auto-heberge-12-etapes-2026.md
source_anchor: ""
source_lines: [207, 283]
sha256: 6379a46667ec244fbb87009a3cd5d4c2f00f342eadfa33fb828278a48ab149f1
---

# Les blocs [Peer] des clients seront ajoutés ici à l'étape 7

- **Oublier la redirection de port UDP.** C'est de loin l'erreur la plus fréquente : le tunnel semble configuré correctement, mais rien ne traverse la box tant que le port n'est pas ouvert et redirigé vers la bonne IP locale.
- **Utiliser une IP locale qui change (DHCP non fixé).** Si le serveur reçoit une nouvelle adresse locale après un redémarrage, la redirection de port pointe vers une machine qui n'existe plus. Réservez toujours un bail DHCP fixe.
- **Oublier d'activer l'IP forwarding.** Le tunnel se connecte, le handshake réussit, mais aucun paquet ne circule vers internet parce que`net.ipv4.ip_forward` est resté désactivé.
- **Réutiliser la même paire de clés sur plusieurs appareils.** Chaque pair doit avoir sa propre clé privée. Partager une clé entre deux appareils casse le suivi des connexions et complique la révocation en cas de perte.
- **Ignorer le CGNAT côté FAI.** Sur certaines offres mobiles ou fibre avec adressage partagé, la redirection de port est tout simplement impossible côté box. Aucune configuration WireGuard ne peut contourner cette limitation réseau.

## Exemple de sortie attendue à chaque étape clé

Voici ce que vous devriez voir si tout fonctionne correctement. Après le démarrage du service à l'étape 6, `wg show` doit afficher quelque chose comme :

```
interface: wg0
  public key: AbCdEf...GhIjKl=
  private key: (hidden)
  listening port: 51820
peer: XyZaBc...DeFgHi=
  endpoint: 82.65.xxx.xxx:41230
  allowed ips: 10.10.0.2/32
  latest handshake: 12 seconds ago
  transfer: 4.21 MiB received, 18.76 MiB sent
```
Un `latest handshake` qui affiche une valeur récente (moins de deux minutes) confirme que le client et le serveur communiquent bien. Si ce champ n'apparaît jamais ou reste vide, le problème se situe presque toujours au niveau réseau (redirection de port, pare-feu) plutôt que dans la configuration WireGuard elle-même.

## Dépannage : 8 problèmes courants et leurs solutions

- **Le tunnel se connecte mais aucun trafic internet ne passe.** Vérifiez que`net.ipv4.ip_forward=1` est actif avec`sysctl net.ipv4.ip_forward` , et que les règles NAT dans`PostUp` /`PostDown` pointent vers la bonne interface sortante.
- **Aucun handshake ne s'établit du tout.** Le port UDP n'est probablement pas atteignable de l'extérieur. Testez avec`nc -u -z -v votre-ip-publique 51820` depuis un réseau externe.
- **La connexion fonctionne en Wi-Fi mais pas en 4G/5G.** Ajoutez`PersistentKeepalive = 25` dans le bloc`[Interface]` du client pour maintenir le tunnel actif à travers le NAT mobile.
- **Erreur « Address already in use » au démarrage.** Un autre service utilise déjà le port 51820, ou une ancienne instance de`wg0` tourne encore. Vérifiez avec`ss -ulnp | grep 51820` .
- **Le DNS ne se résout pas une fois connecté.** Vérifiez la ligne`DNS =` dans le fichier client, et assurez-vous que`resolvconf` ou`systemd-resolved` est installé côté client.
- **Les clients ne se voient pas entre eux.** C'est normal par défaut : WireGuard n'active pas le routage entre pairs sans règles explicites. Ajoutez des règles`iptables FORWARD` spécifiques si nécessaire.
- **Le service ne démarre pas après un redémarrage du serveur.** Vérifiez que`wg-quick@wg0` est bien activé avec`systemctl is-enabled wg-quick@wg0` , sinon relancez la commande`enable` .
- **La redirection de port fonctionne mais la connexion reste lente.** Vérifiez le MTU de l'interface (souvent 1420 par défaut) : certaines connexions fibre nécessitent un ajustement manuel avec`MTU = 1380` dans le fichier de configuration pour éviter la fragmentation des paquets.

## Astuces avancées pour aller plus loin

Une fois le serveur de base opérationnel, plusieurs ajustements améliorent la robustesse et l'usage au quotidien. Premièrement, envisagez un split-tunneling sélectif : plutôt que de router tout le trafic via `AllowedIPs = 0.0.0.0/0`, limitez le VPN à certains sous-réseaux spécifiques (votre réseau domestique, par exemple) pour économiser la bande passante du serveur et garder une navigation directe pour le reste. Deuxièmement, si vous gérez plusieurs serveurs WireGuard sur des sites différents, une configuration site-à-site avec des routes croisées dans les blocs `AllowedIPs` permet de relier deux réseaux locaux entre eux, un usage fréquent pour connecter un bureau et un domicile.

Pensez aussi à automatiser la génération des fichiers clients avec un script bash simple qui crée la paire de clés, ajoute le bloc `[Peer]` côté serveur et génère un QR code avec `qrencode -t ansiutf8 < client.conf` pour un scan direct depuis l'application mobile. Enfin, pour la supervision, un export Prometheus des métriques `wg show` (via l'exporter `prometheus-wireguard-exporter`) permet de suivre les volumes de transfert et la fraîcheur des handshakes dans un tableau de bord Grafana, utile si vous exploitez le VPN pour plusieurs utilisateurs.

## Projet complet : script d'installation automatisé

Pour synthétiser l'ensemble des étapes précédentes, voici un script bash qui automatise l'installation initiale du serveur (à adapter avec votre interface réseau et votre plage d'adresses).

```
#!/bin/bash
set -e
INTERFACE="eth0"
WG_PORT="51820"
WG_SUBNET="10.10.0.1/24"
apt update && apt install -y wireguard wireguard-tools qrencode
mkdir -p /etc/wireguard
cd /etc/wireguard
umask 077
wg genkey | tee server_private.key | wg pubkey > server_public.key
cat > /etc/wireguard/wg0.conf <
```
Ce script couvre l'installation, la génération des clés serveur, l'activation du routage IP et le démarrage du service. Il reste à ajouter manuellement les blocs `[Peer]` pour chaque client, comme décrit à l'étape 7, puis à configurer la redirection de port sur votre box.

## Surveiller et maintenir le serveur dans la durée

Un serveur WireGuard qui fonctionne le jour de l'installation n'est pas nécessairement un serveur qui fonctionnera correctement six mois plus tard. Trois habitudes simples évitent la majorité des mauvaises surprises. D'abord, planifiez une vérification mensuelle des mises à jour de sécurité du système avec `sudo apt update && sudo apt list --upgradable`, en particulier sur le noyau et le paquet `wireguard-tools` lui-même. Ensuite, tenez un inventaire à jour des pairs autorisés : sur un serveur partagé entre plusieurs personnes, il est facile d'oublier de retirer un ancien collègue ou un téléphone remplacé depuis longtemps.

Enfin, surveillez les tentatives de connexion suspectes sur le port SSH resté ouvert. Un serveur exposé sur internet reçoit en permanence des tentatives de scan automatisées, même si WireGuard lui-même reste invisible aux scanners de ports classiques (le protocole ne répond tout simplement pas aux paquets non authentifiés, contrairement à SSH ou à un service web). C'est l'un des avantages de sécurité les moins connus de WireGuard : un scan de ports standard sur le port UDP 51820 ne révèle strictement rien à un attaquant qui ne possède pas déjà une clé valide.

```
# Vérifier les mises à jour disponibles
sudo apt update && sudo apt list --upgradable | grep -i wireguard
# Lister les pairs actuellement configurés et leur dernière activité
sudo wg show wg0 | grep -E "peer|handshake"
# Renouveler la paire de clés du serveur (opération annuelle recommandée)
wg genkey | tee server_private_new.key | wg pubkey > server_public_new.key
```
Si vous renouvelez la clé privée du serveur, n'oubliez pas de mettre à jour la clé publique correspondante dans le fichier de configuration de chaque client, sans quoi tous les tunnels existants cesseront de fonctionner simultanément. Cette opération demande un peu de coordination si plusieurs utilisateurs dépendent du même serveur.

## Sécurité et conformité RGPD : ce que change l'auto-hébergement

