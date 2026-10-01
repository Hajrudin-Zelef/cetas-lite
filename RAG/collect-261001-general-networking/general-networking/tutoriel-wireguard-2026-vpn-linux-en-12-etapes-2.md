---
id: collect-261001-general-networking/general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes-2
title: "Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes.md
source_anchor: ""
source_lines: [47, 206]
sha256: 54b9d3cabb4d336f50f1e3299181bb6fafd6fdaf463457dc26d952afff2f8af9
---

# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13

```
# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13
sudo apt update
sudo apt install -y wireguard wireguard-tools resolvconf
# Rocky Linux 9 / Alma Linux 9 / RHEL 9
sudo dnf install -y epel-release
sudo dnf install -y wireguard-tools
# Fedora 41+
sudo dnf install -y wireguard-tools
# Vérification
wg --version
# Attendu : wireguard-tools v1.0.x - https://git.zx2c4.com/wireguard-tools/
modinfo wireguard | head -5
# Confirme que le module est compilé dans le kernel
```
Si la commande `modinfo wireguard` renvoie une erreur, c’est que votre kernel est antérieur à 5.6. Sur Ubuntu 20.04 (kernel 5.4), installez le HWE pour basculer sur un kernel 5.15+ : `sudo apt install -y linux-generic-hwe-20.04` puis redémarrez. Sur les vieilles distributions CentOS 7, optez pour Rocky Linux 9 ou migrez vers Debian 12+.

## Étape 2 : Génération des paires de clés Curve25519 pour le serveur

WireGuard utilise des clés **Curve25519** de 256 bits encodées en Base64. Chaque pair (serveur ou client) possède une **clé privée** et une **clé publique** dérivée. Optionnellement, vous pouvez ajouter une **clé pré-partagée** (PSK) entre deux pairs pour renforcer la résistance contre une éventuelle attaque post-quantique sur Curve25519.

```
cd /etc/wireguard
sudo umask 077
# Clé privée serveur
sudo wg genkey | sudo tee server_private.key
# Exemple : kP3xQ9vR2L5N8mY1aT7fE4cB6dW0sH3jK9lZ2qX8vU4=
# Clé publique serveur dérivée
sudo cat server_private.key | wg pubkey | sudo tee server_public.key
# Exemple : aB1cD2eF3gH4iJ5kL6mN7oP8qR9sT0uV1wX2yZ3aB4c=
# Clé pré-partagée optionnelle (renforce la PFS)
sudo wg genpsk | sudo tee server_preshared.key
ls -l /etc/wireguard/
# -rw------- 1 root root 45 avr 30 11:22 server_private.key
# -rw------- 1 root root 45 avr 30 11:22 server_public.key
# -rw------- 1 root root 45 avr 30 11:22 server_preshared.key
```
Le `umask 077` est essentiel : les fichiers ne doivent être lisibles que par `root`. Une clé privée WireGuard exposée équivaut à offrir l’accès complet au tunnel à n’importe qui. Conservez ces fichiers hors de tout dépôt Git, hors de toute sauvegarde non chiffrée et hors d’un cloud public sans envelope encryption.

## Étape 3 : Configuration du serveur WireGuard avec wg-quick

Créez le fichier `/etc/wireguard/wg0.conf` qui définit l’interface tunnel `wg0`. Nous utilisons le sous-réseau privé `10.66.0.0/24` pour les pairs et le port UDP standard **51820**. Ajoutez les règles iptables qui activent le NAT vers Internet (utile pour un VPN d’accès distant).

```
sudo nano /etc/wireguard/wg0.conf
[Interface]
# Adresse interne du serveur dans le tunnel
Address = 10.66.0.1/24
# Port UDP d'écoute (par défaut 51820)
ListenPort = 51820
# Clé privée serveur (contenu de server_private.key)
PrivateKey = kP3xQ9vR2L5N8mY1aT7fE4cB6dW0sH3jK9lZ2qX8vU4=
# Sauvegarde de la configuration des pairs (utile pour wg-quick save)
SaveConfig = false
# MTU recommandé pour éviter la fragmentation UDP
MTU = 1420
# NAT et forwarding au démarrage
PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
# Premier client (à compléter à l'étape 5)
# [Peer]
# PublicKey = ...
# AllowedIPs = 10.66.0.2/32
# Sécurisez les permissions
sudo chmod 600 /etc/wireguard/wg0.conf
```
Adaptez `eth0` à votre interface de sortie réelle (`ens3`, `enp1s0`, etc.). Vous pouvez l’identifier avec `ip route show default`. Pour un usage IPv6, ajoutez une seconde paire de règles avec `ip6tables` et déclarez `Address = 10.66.0.1/24, fd42:42:42::1/64`.

## Étape 4 : Activation du routage IP et règles de pare-feu

Sans IP forwarding, votre serveur ne pourra pas relayer les paquets entre l’interface `wg0` et l’Internet. Activez-le de façon persistante via `sysctl`, puis ouvrez le port UDP 51820 sur le pare-feu hôte (UFW, firewalld ou nftables).

```
# Activation persistante du forwarding
echo 'net.ipv4.ip_forward=1' | sudo tee /etc/sysctl.d/99-wireguard.conf
echo 'net.ipv6.conf.all.forwarding=1' | sudo tee -a /etc/sysctl.d/99-wireguard.conf
sudo sysctl --system
# Vérification
sysctl net.ipv4.ip_forward
# net.ipv4.ip_forward = 1
# UFW (Ubuntu/Debian)
sudo ufw allow 51820/udp
sudo ufw allow OpenSSH
sudo ufw enable
# firewalld (RHEL/Rocky/Fedora)
sudo firewall-cmd --permanent --add-port=51820/udp
sudo firewall-cmd --permanent --add-masquerade
sudo firewall-cmd --reload
# nftables (configuration moderne)
sudo nft add rule inet filter input udp dport 51820 accept
```
Si votre VPS est derrière un pare-feu opérateur (Hetzner Cloud Firewall, AWS Security Group, OVH IP Firewall), pensez à **autoriser explicitement** l’UDP 51820 entrant. Sur AWS EC2, créez une règle **Custom UDP, Port 51820, Source 0.0.0.0/0** ou restreignez-la aux IP publiques de vos clients.

## Étape 5 : Génération des clés et de la configuration d’un client

Chaque pair distant a besoin de sa propre paire de clés. Générez-la côté serveur (par commodité), notez les valeurs, puis ajoutez le bloc `[Peer]` correspondant dans `wg0.conf` du serveur. Le fichier final côté client sera ensuite transmis de façon sécurisée.

```
cd /etc/wireguard
sudo wg genkey | sudo tee client1_private.key | wg pubkey | sudo tee client1_public.key
CLIENT_PRIV=$(sudo cat client1_private.key)
CLIENT_PUB=$(sudo cat client1_public.key)
SERVER_PUB=$(sudo cat server_public.key)
PSK=$(sudo cat server_preshared.key)
# Ajout du peer côté serveur (à appender à wg0.conf)
sudo tee -a /etc/wireguard/wg0.conf >/dev/null <<EOF
[Peer]
# client1 (laptop bureau)
PublicKey = $CLIENT_PUB
PresharedKey = $PSK
AllowedIPs = 10.66.0.2/32
EOF
# Génération de la configuration côté client
sudo tee /etc/wireguard/client1.conf >/dev/null <<EOF
[Interface]
PrivateKey = $CLIENT_PRIV
Address = 10.66.0.2/32
DNS = 1.1.1.1, 9.9.9.9
MTU = 1420
[Peer]
PublicKey = $SERVER_PUB
PresharedKey = $PSK
Endpoint = vpn.exemple.fr:51820
AllowedIPs = 0.0.0.0/0, ::/0
PersistentKeepalive = 25
EOF
```
Le paramètre `AllowedIPs = 0.0.0.0/0, ::/0` route **tout le trafic** du client à travers le tunnel (mode « full tunnel »). Pour un VPN d’entreprise n’exposant que le LAN interne, remplacez par `AllowedIPs = 10.0.0.0/8, 192.168.0.0/16` (mode « split tunnel ») afin d’éviter de pousser inutilement le trafic streaming dans le tunnel.

## Étape 6 : Démarrage du serveur et activation systemd

Activez l’unité `[email protected]` pour démarrer WireGuard automatiquement au boot. `wg-quick` est un wrapper bash léger qui lit votre fichier `.conf`, crée l’interface et applique les hooks `PostUp`/`PostDown`.

```
# Démarrage immédiat
sudo wg-quick up wg0
# Activation au démarrage
sudo systemctl enable [email protected]
# Vérification de l'état
sudo systemctl status [email protected]
sudo wg show
# Sortie attendue :
# interface: wg0
#   public key: aB1cD2eF3gH4iJ5kL6mN7oP8qR9sT0uV1wX2yZ3aB4c=
#   private key: (hidden)
#   listening port: 51820
#
# peer: zX9yY8wW7vV6uU5tT4sS3rR2qQ1pP0oO9nN8mM7lL6k=
#   preshared key: (hidden)
#   allowed ips: 10.66.0.2/32
#   latest handshake: 12 seconds ago
#   transfer: 1.04 MiB received, 5.83 MiB sent
#   persistent keepalive: every 25 seconds
```
Pour rejouer la configuration sans redémarrer (par exemple après ajout d’un nouveau peer), utilisez la combinaison atomique : `sudo wg syncconf wg0 <(wg-quick strip wg0)`. Cette commande applique uniquement le diff sans interrompre les sessions actives — comportement crucial en production.

## Étape 7 : Configuration des clients Linux, macOS et Windows

Transférez le fichier `client1.conf` à votre poste client par un canal chiffré (SCP, SFTP, magic-wormhole). Évitez systématiquement Slack, e-mail non chiffré et clés USB non chiffrées.

