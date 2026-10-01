---
id: collect-261001-general-networking/general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes-5
title: "Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws", "compute", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes.md
source_anchor: ""
source_lines: [454, 575]
sha256: d0992a0f7c280ce4f6905cb0b3ca51a55d8a5182c2c166a06a3316e307184a60
---

# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13

- **« RTNETLINK answers: Operation not supported »** : kernel trop ancien (< 5.6) ou module wireguard non chargé. Solution :`modprobe wireguard` ou mise à jour du kernel.
- **« Handshake never completed »** : vérifiez que l’`Endpoint` et le`ListenPort` sont accessibles depuis le client (`nc -uvz vpn.exemple.fr 51820` ) et que les clés correspondent.
- **Trafic bloqué après handshake réussi** :`net.ipv4.ip_forward = 0` . Activez-le et rejouez le NAT.
- **« peer (xxx) – Receive: Invalid handshake initiation »** : désynchronisation horaire ou clés serveur/client inversées.
- **Vitesse anormalement basse (10 Mbit/s au lieu de 900)** : probable saturation CPU sur architecture ARM ancienne, ou MTU mal réglé. Forcez`MTU = 1280` et testez avec`iperf3` .
- **Le tunnel se ferme après quelques minutes** : NAT CGNAT ou pare-feu state-tracking trop agressif. Ajoutez`PersistentKeepalive = 25` côté client.
- **« Address already in use »** : une autre interface utilise déjà le subnet`10.66.0.0/24` . Choisissez un autre RFC1918 (`10.99.0.0/24` ,`172.30.0.0/24` ).
- **DNS ne fonctionne pas après connexion** :`resolvconf` non installé ou systemd-resolved en conflit. Installez`openresolv` et déclarez`DNS = 1.1.1.1` dans`[Interface]` .

## Comparatif WireGuard avec les écosystèmes mesh modernes

| Solution | Modèle | SSO | NAT traversal | Plan gratuit | Souveraineté UE | 
|---|---|---|---|---|---|
| WireGuard pur | Manuel | Aucun | Manuel (port forwarding) | Oui (auto-hébergé) | Totale | 
| Tailscale | SaaS | Google, MS, Okta, GitHub | DERP relays automatique | ≤ 100 machines, 3 users | Données contrôle hors UE | 
| Headscale | Auto-hébergé | OIDC standard | Compatible STUN | Illimité (open source) | Totale | 
| NetBird | SaaS + auto-hébergé | Keycloak, Zitadel, Auth0 | STUN/TURN intégré | ≤ 100 machines (cloud) | Cloud UE disponible | 
| Cilium WireGuard | Kubernetes CNI | RBAC K8s | Pod-to-pod direct | Open source | Totale | 
| Innernet | Auto-hébergé | Tokens d’invitation | NAT-PMP/PCP | Open source (Tonari) | Totale | 

Pour une PME française soumise au **RGPD** et soucieuse de souveraineté, le combo **Headscale + WireGuard** sur un VPS souverain (OVHcloud, Scaleway, Outscale) coche toutes les cases : open source, hébergé en France, auditable, et compatible avec les clients officiels Tailscale gratuits. Comptez environ **5 €/mois** pour le VPS et **0 €** de licence.

## Projet complet : VPN d’équipe avec 10 utilisateurs et accès LAN

Mettons en pratique tout le tutoriel avec un cas d’usage réaliste : votre startup de 10 personnes a besoin d’accéder à un serveur de fichiers (192.168.1.50), un GitLab interne (192.168.1.20) et un PostgreSQL (192.168.1.30) hébergés au bureau, depuis n’importe où. Architecture cible : un VPS OVHcloud à Roubaix (D2-2, 5 €/mois) sert de point d’entrée WireGuard, et un Raspberry Pi 5 dans le bureau joue le rôle de pont vers le LAN.

```
# --- VPS Roubaix (vpn.startup.fr) ---
# /etc/wireguard/wg0.conf
[Interface]
Address = 10.66.0.1/24
ListenPort = 51820
PrivateKey = <PRIV_VPS>
PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
# Pont LAN bureau (Raspberry Pi 5)
[Peer]
PublicKey = <PUB_PI>
PresharedKey = <PSK_PI>
AllowedIPs = 10.66.0.254/32, 192.168.1.0/24
PersistentKeepalive = 25
# 10 utilisateurs (générés via wg-add-client.sh)
[Peer]
# alice
PublicKey = ...
AllowedIPs = 10.66.0.2/32
[Peer]
# bob
PublicKey = ...
AllowedIPs = 10.66.0.3/32
# ... idem jusqu'à 10.66.0.11
# --- Raspberry Pi 5 bureau ---
# /etc/wireguard/wg0.conf
[Interface]
Address = 10.66.0.254/32
PrivateKey = <PRIV_PI>
PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT
[Peer]
PublicKey = <PUB_VPS>
PresharedKey = <PSK_PI>
Endpoint = vpn.startup.fr:51820
# Le Pi route le LAN bureau pour les utilisateurs
AllowedIPs = 10.66.0.0/24
PersistentKeepalive = 25
# --- Configuration utilisateur (alice.conf) ---
[Interface]
PrivateKey = <PRIV_ALICE>
Address = 10.66.0.2/32
DNS = 192.168.1.1, 1.1.1.1
MTU = 1420
[Peer]
PublicKey = <PUB_VPS>
PresharedKey = <PSK_ALICE>
Endpoint = vpn.startup.fr:51820
# Split tunnel : RFC1918 + 10.66.0.0/24
AllowedIPs = 10.66.0.0/24, 192.168.1.0/24
PersistentKeepalive = 25
```
Avec ce design, Alice peut `ssh [email protected]`, ouvrir `https://gitlab.startup.local` et monter le SMB `//192.168.1.50/projets` exactement comme si elle était au bureau. Coût mensuel total : **5 € VPS + 0 € licence + ~3 € électricité Pi**. Pour comparer, une licence **Cisco AnyConnect** ou **Fortinet FortiClient** équivalente coûte entre 50 et 150 € par utilisateur et par an.

## Sécurité durable : rotation des clés, audit et conformité

WireGuard ne propose pas nativement de mécanisme de rotation automatique. C’est à vous de scripter la régénération périodique. Voici un rythme recommandé pour un usage entreprise : **PSK tous les 90 jours**, **clés Curve25519 tous les 12 mois**, **audit complet annuel** avec inventaire des peers (CMDB).

```
# Rotation PSK (sans interruption, 1 peer à la fois)
NEW_PSK=$(wg genpsk)
echo "$NEW_PSK" | sudo tee /etc/wireguard/server_preshared.key
sudo wg set wg0 peer <PUB_PEER> preshared-key <(echo "$NEW_PSK")
# Mettre à jour le fichier client puis le redéployer
# Audit des peers actifs
sudo wg show wg0 latest-handshakes | awk -v now=$(date +%s) '{
  diff = now - $2;
  if (diff > 86400) print $1, "INACTIF", diff/3600, "h"
}'
# Révocation immédiate d'un peer
sudo wg set wg0 peer <PUB_REVOQUE> remove
sudo sed -i "/PublicKey = <PUB_REVOQUE>/,/^$/d" /etc/wireguard/wg0.conf
# Sauvegarde chiffrée des secrets (avec age)
sudo apt install -y age
age-keygen -o ~/age-key.txt
age -r <age_pubkey> -o /backup/wireguard-$(date +%F).tar.age \
  < <(sudo tar c /etc/wireguard/)
```
Pour la conformité **NIS2** et **DORA** entrées en application en 2025-2026, documentez votre configuration dans un registre (qui est connecté, depuis quelle IP source, vers quels actifs internes). WireGuard étant journalisé indirectement via le pare-feu, complétez avec des logs **nftables** exportés vers un SIEM (Elastic, Wazuh, Splunk).

## WireGuard sur cloud public : AWS, Azure, GCP, OVHcloud, Scaleway

Les hyperscalers et les fournisseurs européens proposent tous des images Linux récentes compatibles WireGuard sans configuration supplémentaire. Voici les particularités à connaître :

- **AWS EC2** : utiliser une AMI Ubuntu 24.04 ou Amazon Linux 2023. Désactiver le**Source/Destination Check** dans la console pour autoriser le forwarding. Ouvrir UDP 51820 dans le Security Group.
- **Azure VM** : image Ubuntu Server 24.04 LTS. Activer**IP Forwarding** sur la NIC associée. Network Security Group : règle inbound UDP 51820.
- **Google Cloud Compute Engine** : cocher « IP forwarding » au moment de la création. Firewall : allow-wireguard sur UDP 51820 + tag réseau.
- **OVHcloud Public Cloud** : instance D2-2 ou B3-8 avec image Debian 13. Pas de pare-feu opérateur par défaut, gérer via`ufw` directement.
- **Scaleway Instances** : PLAY2-MICRO suffit. Activer le DEV1 NIC, ouvrir UDP 51820 dans Security Group.
- **Hetzner Cloud** : CX22 (4,49 €/mois) avec Ubuntu 24.04. Configurer le Cloud Firewall depuis la console.

Pour des charges de travail multi-cloud, l’orchestration via **Terraform** du VPC peering au-dessus de WireGuard devient courante. Le module `terraform-aws-wireguard` de Cloudposse, et le `tailscale_acl` resource du provider Tailscale officiel, simplifient grandement le déploiement IaC.

## Foire aux questions sur WireGuard et le déploiement VPN moderne

