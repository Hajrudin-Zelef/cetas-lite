---
id: collect-261001-general-networking/general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes-4
title: "Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["arr", "attention", "kill switch", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes.md
source_anchor: ""
source_lines: [350, 453]
sha256: a80b7cd167cebf464bc8043b991787eda9ea4e995668acabc4846d58a5c04c51
---

# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13

```
# --- Tailscale (SaaS, scénario le plus rapide) ---
curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up --ssh --advertise-routes=192.168.10.0/24
# Authentification via votre IdP : Google, Microsoft, GitHub, Okta
# --- Headscale (auto-hébergé, sans dépendance Tailscale Inc.) ---
wget https://github.com/juanfont/headscale/releases/latest/download/headscale_linux_amd64.deb
sudo dpkg -i headscale_linux_amd64.deb
sudo nano /etc/headscale/config.yaml
# server_url: https://hs.exemple.fr:8080
# listen_addr: 0.0.0.0:8080
# database:
#   type: sqlite
#   sqlite:
#     path: /var/lib/headscale/db.sqlite
# ip_prefixes: [10.64.0.0/10]
sudo systemctl enable --now headscale
sudo headscale users create equipe-tech
sudo headscale --user equipe-tech preauthkeys create --reusable --expiration 24h
# Sur chaque client (CLI Tailscale détourné vers Headscale)
sudo tailscale up --login-server https://hs.exemple.fr:8080 --authkey <preauth>
# --- NetBird (open source, SSO Keycloak/Auth0/Zitadel) ---
curl -fsSL https://pkgs.netbird.io/install.sh | sh
sudo netbird up --setup-key <cle_setup> --management-url https://app.netbird.io
```
Tailscale apporte aussi **MagicDNS** (résolution `nom-machine.tailnet.ts.net`), **ACL en JSON/HuJSON**, **Tailscale SSH** (sans clés sur les serveurs), et **Funnel** pour exposer un service privé en HTTPS public. Headscale réplique 95 % de ces fonctionnalités sans envoyer aucune métadonnée à un tiers — choix recommandé pour la souveraineté française et européenne.

## Étape 12 : Supervision, métriques et journalisation

WireGuard n’écrit pas de logs en standard : il pousse des compteurs minimalistes via `wg show`. Pour un usage production, exposez ces métriques à **Prometheus** via l’exporter `prometheus-wireguard-exporter` et tracez l’activité dans **Grafana**.

```
# Exporter Prometheus pour WireGuard
wget https://github.com/MindFlavor/prometheus_wireguard_exporter/releases/latest/download/prometheus_wireguard_exporter_linux_amd64
sudo mv prometheus_wireguard_exporter_linux_amd64 /usr/local/bin/wg_exporter
sudo chmod +x /usr/local/bin/wg_exporter
sudo tee /etc/systemd/system/wg-exporter.service >/dev/null <<EOF
[Unit]
Description=WireGuard Prometheus exporter
After=network.target
[Service]
ExecStart=/usr/local/bin/wg_exporter -a -n /etc/wireguard/wg0.conf
Restart=always
User=root
[Install]
WantedBy=multi-user.target
EOF
sudo systemctl daemon-reload
sudo systemctl enable --now wg-exporter
# Test
curl -s http://localhost:9586/metrics | grep wireguard_
# Scrape côté Prometheus (prometheus.yml)
- job_name: 'wireguard'
  static_configs:
    - targets: ['vpn.exemple.fr:9586']
# Activation des logs noyau (debug)
sudo modprobe wireguard
echo 'module wireguard +p' | sudo tee /sys/kernel/debug/dynamic_debug/control
sudo dmesg -w | grep wireguard
```
Dans Grafana, importez le dashboard officiel WireGuard (ID **12177**) qui affiche par peer : dernier handshake, octets reçus/envoyés, statut connecté/déconnecté, et alerte automatique si un peer reste silencieux plus de 3 minutes. Couplez ces alertes à **Alertmanager** pour des notifications Slack, e-mail ou PagerDuty.

## Tableau récapitulatif des commandes essentielles WireGuard

| Action | Commande | 
|---|---|
| Générer une clé privée | `wg genkey` | 
| Dériver la clé publique | `echo <priv> \| wg pubkey` | 
| Générer une PSK | `wg genpsk` | 
| Démarrer un tunnel | `wg-quick up wg0` | 
| Arrêter un tunnel | `wg-quick down wg0` | 
| Activer au boot | `systemctl enable wg-quick@wg0` | 
| Statut détaillé | `wg show` | 
| Statistiques transferts | `wg show all transfer` | 
| Recharge sans coupure | `wg syncconf wg0 <(wg-quick strip wg0)` | 
| Sauvegarder l’état actif | `wg-quick save wg0` | 
| Ajouter un peer à chaud | `wg set wg0 peer <PUB> allowed-ips 10.66.0.10/32` | 
| Retirer un peer à chaud | `wg set wg0 peer <PUB> remove` | 

## Pièges courants et erreurs à éviter avec WireGuard

- **Oublier d’activer `net.ipv4.ip_forward`** : le tunnel monte mais aucun trafic ne sort vers Internet. C’est l’erreur n°1 des débutants.
- **Confondre AllowedIPs côté serveur et côté client** : côté serveur, c’est l’IP du peer (`10.66.0.X/32` ) ; côté client, c’est ce que vous voulez router à travers le tunnel (`0.0.0.0/0` ou un subnet).
- **MTU laissé à 1500** : provoque de la fragmentation UDP et de la perte de paquets sur les liens PPPoE/4G. Forcez`MTU = 1420` , voire 1280 sur les liens problématiques.
- **Port UDP 51820 bloqué par l’ISP** : certains hôtels et opérateurs mobiles filtrent les ports inhabituels. Solution : faire écouter sur 53 ou 443 UDP, ou encapsuler dans**udp2raw** .
- **NAT côté CGNAT mobile** : sans`PersistentKeepalive = 25` , le tunnel meurt silencieusement après 2-5 minutes d’inactivité.
- **Clé privée commitée dans Git** : ajoutez`/etc/wireguard/*.key` à votre`.gitignore` et utilisez`git-crypt` ou**SOPS** +age pour les secrets versionnés.
- **Doubler les règles NAT entre `PostUp` et iptables-persistent** : cause des compteurs faux et des paquets droppés.
- **Mauvais ordre des règles iptables** : si`FORWARD` par défaut est`DROP` , vos règles`ACCEPT` doivent précéder le`REJECT` final.
- **DNS leak via le résolveur système** : forcez un DNS public (`1.1.1.1` ) ou interne dans la section`[Interface]` .
- **Heure système désynchronisée** : WireGuard rejette les handshakes dont le timestamp est aberrant. Activez`chronyd` ou`systemd-timesyncd` .

## Astuces avancées : chiffrement post-quantique, kill switch, multi-hop

Une fois la base maîtrisée, plusieurs raffinements méritent l’attention. **Chiffrement résistant au post-quantique** : ajoutez systématiquement une **PresharedKey** entre chaque paire de pairs. WireGuard la combine au handshake Curve25519 via HKDF, ce qui empêche un futur ordinateur quantique de déchiffrer les enregistrements capturés aujourd’hui (attaque « harvest now, decrypt later »). C’est gratuit, renouvelable et obligatoire pour les usages sensibles.

**Kill switch** : bloquez tout trafic sortant du client si le tunnel tombe, pour éviter une fuite d’IP réelle. Avec `iptables`, ajoutez `iptables -I OUTPUT ! -o wg0 -m mark ! --mark $(wg show wg0 fwmark) -m addrtype ! --dst-type LOCAL -j REJECT` dans `PostUp`. Ou utilisez la directive `SaveConfig = false` combinée à un namespace réseau dédié pour isoler complètement le processus VPN.

**Multi-hop (cascade)** : chaînez deux serveurs WireGuard pour augmenter l’anonymat. Le client se connecte au serveur A (`wg0`), qui route le trafic vers un second serveur B via une seconde interface (`wg1`). C’est l’architecture utilisée par **Mullvad** et **IVPN** pour leurs offres « dual-hop ». Ajoutez une politique de routage `ip rule` par marque de paquet pour aiguiller proprement entre les deux tunnels.

Enfin, pour les **conteneurs Docker et Kubernetes**, l’image officielle `linuxserver/wireguard` (basée sur Alpine) intègre `wireguard-tools` et permet de monter un tunnel par sidecar. Sur Kubernetes, les CNI **Cilium** et **Calico** supportent désormais WireGuard comme overlay chiffré entre nœuds — fonctionnalité GA depuis Cilium 1.14 (juin 2023).

## Dépannage : les 8 erreurs WireGuard les plus fréquentes

