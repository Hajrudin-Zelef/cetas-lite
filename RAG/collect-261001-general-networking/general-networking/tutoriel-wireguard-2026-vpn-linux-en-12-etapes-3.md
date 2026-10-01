---
id: collect-261001-general-networking/general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes-3
title: "Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Apple"]
dates: []
keywords: ["aws", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes.md
source_anchor: ""
source_lines: [207, 349]
sha256: afa34986c776f0f4f33413c536307e2914396fcd05a02b81da29364ac8cf9340
---

# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13

```
# --- Linux desktop ---
sudo apt install -y wireguard wireguard-tools
sudo cp client1.conf /etc/wireguard/wg0.conf
sudo chmod 600 /etc/wireguard/wg0.conf
sudo wg-quick up wg0
sudo systemctl enable [email protected]
# Test du tunnel
ping -c 4 10.66.0.1
curl -s ifconfig.io
# Doit afficher l'IP publique du serveur
# --- macOS (via Homebrew + App Store) ---
brew install wireguard-tools
# Ouvrir WireGuard.app, Importer le tunnel depuis client1.conf
# --- Windows 10/11 ---
# Télécharger https://download.wireguard.com/windows-client/
# Lancer WireGuard.msi, Importer Tunnel(s) from File... > client1.conf
# Cliquer Activate
```
Sur macOS, l’application officielle apporte une intégration parfaite avec le panneau Réseau et un statut visible dans la barre de menus. Sur Windows 11 25H2, l’application s’appuie désormais sur la branche **WireGuard for Windows 1.1**, lancée en deux versions successives en mai 2026 (suivi Versio.io), puis affinée par la version **1.1.1** publiée le 20 septembre 2026 — confirmée à la fois par le dépôt GitHub officiel et par l’annonce de Jason A. Donenfeld (ZX2C4) — qui supporte enfin nativement les **profils par défaut DNS-over-HTTPS**, ce qui simplifie le déploiement en environnement Active Directory.

## Étape 8 : Génération de QR codes pour iOS et Android

Sur mobile, la méthode la plus rapide consiste à scanner un **QR code** contenant la configuration. Installez `qrencode` côté serveur, puis générez le code à partir du fichier `client1.conf`. Le QR code apparaît directement dans le terminal.

```
# Installation
sudo apt install -y qrencode
# QR code dans le terminal
sudo qrencode -t ansiutf8 < /etc/wireguard/client1.conf
# Ou en fichier PNG
sudo qrencode -o /tmp/client1.png < /etc/wireguard/client1.conf
# --- iOS ---
# 1. Installer WireGuard depuis l'App Store (1.0.16+)
# 2. Ouvrir l'app > Ajouter un tunnel > Créer depuis QR code
# 3. Scanner le QR affiché par qrencode
# 4. Activer le toggle Status
# --- Android ---
# 1. Installer WireGuard 1.0.20260315 depuis le Play Store
# 2. Toucher le bouton + > Scanner depuis QR code
# 3. Donner les permissions VPN
# 4. Activer le toggle
```
Important : **supprimez le PNG temporaire** immédiatement après usage (`sudo shred -u /tmp/client1.png`) car il contient la clé privée du client en clair. Ne stockez jamais ces QR codes dans un gestionnaire de fichiers cloud non chiffré.

## Étape 9 : Script Bash pour automatiser l’ajout de clients

Pour onboarder rapidement de nouveaux utilisateurs, voici un script de production qui génère les clés, ajoute le peer au serveur, applique la configuration sans interruption et produit un fichier client prêt à l’emploi. Sauvegardez-le dans `/usr/local/sbin/wg-add-client.sh`.

```
#!/usr/bin/env bash
# wg-add-client.sh - Génère un nouveau client WireGuard
# Usage : sudo wg-add-client.sh <nom_client>
set -euo pipefail
CLIENT_NAME="${1:?Usage: $0 <nom_client>}"
WG_DIR="/etc/wireguard"
SERVER_PUB=$(cat "$WG_DIR/server_public.key")
PSK=$(cat "$WG_DIR/server_preshared.key")
ENDPOINT="vpn.exemple.fr:51820"
DNS="1.1.1.1, 9.9.9.9"
# Calcul de la prochaine IP libre dans 10.66.0.0/24
LAST_IP=$(grep -E "AllowedIPs = 10\.66\.0\." "$WG_DIR/wg0.conf" \
  | awk -F. '{print $4}' | awk -F/ '{print $1}' \
  | sort -n | tail -1)
NEXT_IP=$((LAST_IP + 1))
[[ $NEXT_IP -gt 254 ]] && { echo "Plage saturée"; exit 1; }
CLIENT_IP="10.66.0.$NEXT_IP"
# Génération des clés
umask 077
wg genkey | tee "$WG_DIR/${CLIENT_NAME}_private.key" \
  | wg pubkey > "$WG_DIR/${CLIENT_NAME}_public.key"
CLIENT_PRIV=$(cat "$WG_DIR/${CLIENT_NAME}_private.key")
CLIENT_PUB=$(cat "$WG_DIR/${CLIENT_NAME}_public.key")
# Ajout du peer côté serveur
cat >> "$WG_DIR/wg0.conf" <<EOF
[Peer]
# ${CLIENT_NAME}
PublicKey = $CLIENT_PUB
PresharedKey = $PSK
AllowedIPs = $CLIENT_IP/32
EOF
# Application sans interruption
wg syncconf wg0 <(wg-quick strip wg0)
# Fichier client
cat > "$WG_DIR/${CLIENT_NAME}.conf" <<EOF
[Interface]
PrivateKey = $CLIENT_PRIV
Address = $CLIENT_IP/32
DNS = $DNS
MTU = 1420
[Peer]
PublicKey = $SERVER_PUB
PresharedKey = $PSK
Endpoint = $ENDPOINT
AllowedIPs = 0.0.0.0/0, ::/0
PersistentKeepalive = 25
EOF
echo "Client $CLIENT_NAME créé avec IP $CLIENT_IP"
echo "Configuration : $WG_DIR/${CLIENT_NAME}.conf"
echo "QR code : qrencode -t ansiutf8 < $WG_DIR/${CLIENT_NAME}.conf"
```
Rendez le script exécutable : `sudo chmod 700 /usr/local/sbin/wg-add-client.sh`. Ensuite, créer un nouvel utilisateur prend trois secondes : `sudo wg-add-client.sh laptop-marie`. Vous obtenez un fichier prêt à scanner ou à transférer.

## Étape 10 : VPN site à site entre deux serveurs WireGuard

Au-delà du VPN d’accès distant, WireGuard excelle dans les **liaisons site à site** — typique des connexions entre deux datacenters, deux bureaux ou deux clouds (AWS ↔ OVH par exemple). Chaque site agit comme un peer dont l’`AllowedIPs` annonce le subnet local.

```
# --- Site A (Paris) - LAN 192.168.10.0/24, IP publique 203.0.113.10 ---
# /etc/wireguard/wg0.conf
[Interface]
Address = 10.99.0.1/30
ListenPort = 51820
PrivateKey = <PRIV_A>
PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
[Peer]
# Site B (Lyon)
PublicKey = <PUB_B>
Endpoint = 198.51.100.20:51820
AllowedIPs = 10.99.0.2/32, 192.168.20.0/24
PersistentKeepalive = 25
# --- Site B (Lyon) - LAN 192.168.20.0/24, IP publique 198.51.100.20 ---
[Interface]
Address = 10.99.0.2/30
ListenPort = 51820
PrivateKey = <PRIV_B>
[Peer]
# Site A (Paris)
PublicKey = <PUB_A>
Endpoint = 203.0.113.10:51820
AllowedIPs = 10.99.0.1/32, 192.168.10.0/24
PersistentKeepalive = 25
# Sur chaque routeur LAN, ajouter une route statique
ip route add 192.168.20.0/24 via <IP_LAN_serveurA>  # côté Paris
ip route add 192.168.10.0/24 via <IP_LAN_serveurB>  # côté Lyon
```
Cette topologie permet à un poste du LAN de Paris (192.168.10.42) de joindre transparemment un serveur du LAN de Lyon (192.168.20.50) à travers le tunnel chiffré. Pour aller plus loin, vous pouvez configurer un **routage dynamique BGP** via **BIRD** ou **FRRouting** par-dessus le tunnel WireGuard, scénario désormais utilisé par OVHcloud, Scaleway et Hetzner pour leurs offres VPC inter-régions.

## Étape 11 : Mesh peer-to-peer avec Tailscale, Headscale et NetBird

Configurer manuellement n × (n-1) tunnels devient vite ingérable. Pour un mesh de 50 machines, vous itérez sur 2 450 paires. Trois écosystèmes sur WireGuard ont émergé pour automatiser ce travail : **Tailscale** (SaaS, plan gratuit jusqu’à 100 machines), **Headscale** (serveur de contrôle Tailscale auto-hébergé) et **NetBird** (open source, SSO intégré).

