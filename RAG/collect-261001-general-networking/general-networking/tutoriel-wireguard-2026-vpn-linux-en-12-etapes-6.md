---
id: collect-261001-general-networking/general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes-6
title: "Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "cyber", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes.md
source_anchor: ""
source_lines: [576, 610]
sha256: 02ebf9e6e39457a7a457627ccbe62e98d454363d2bbd0dd033cbc48f7f0b272b
---

# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13

### WireGuard est-il vraiment plus rapide qu’OpenVPN ?

Oui, et la marge est significative. Sur un lien 1 Gbit/s entre deux VPS modernes, WireGuard atteint typiquement 940 à 960 Mbit/s contre 200-300 Mbit/s pour OpenVPN UDP — et sur un lien 10 Gbit/s, un benchmark Zeroday Cyber Academy de mars 2026 relève jusqu’à 8,7 Gbit/s. Un comparatif GeekSynapse publié le même mois confirme un avantage constant de 2 à 4 fois sur matériel identique. La raison : WireGuard tourne en espace noyau (sur Linux), évite les changements de contexte coûteux, et utilise ChaCha20-Poly1305, plus rapide qu’AES en software sur ARM. Sur smartphone, l’écart se traduit aussi par une meilleure autonomie batterie.

### Quelles distributions Linux supportent WireGuard nativement en 2026 ?

Toutes les distributions avec un noyau ≥ 5.6 incluent le module WireGuard. En 2026, cela couvre : Ubuntu 22.04 LTS / 24.04 LTS, Debian 12 (Bookworm) et 13 (Trixie), Rocky Linux 9, Alma Linux 9, RHEL 9, Fedora 40+, openSUSE Leap 15.5+, Arch Linux (rolling), Alpine 3.18+. Pour les plus anciennes (CentOS 7, Ubuntu 18.04), une migration s’impose : leurs cycles de support sont terminés ou imminents.

### WireGuard est-il sûr face aux ordinateurs quantiques ?

Curve25519 sera vulnérable à un ordinateur quantique disposant d’environ 4 000 qubits logiques stables — niveau pas encore atteint en 2026. Pour anticiper, ajoutez systématiquement une **PresharedKey** entre chaque paire : combinée via HKDF, elle ajoute une couche symétrique qu’un quantique ne peut pas casser même rétroactivement. C’est gratuit et déjà recommandé par l’ANSSI pour les usages sensibles.

### Quelle différence entre WireGuard, Tailscale et Headscale ?

WireGuard est le protocole de chiffrement et l’implémentation de bas niveau. Tailscale est une plate-forme commerciale qui simplifie le mesh, l’authentification SSO et la traversée NAT — son plan de contrôle est hébergé par Tailscale Inc. (USA). Headscale est une réimplémentation open source de ce plan de contrôle : vous gardez la simplicité opérationnelle de Tailscale tout en hébergeant tout vous-même, ce qui satisfait les exigences de souveraineté française et européenne.

### Peut-on utiliser WireGuard derrière un CGNAT mobile ou un hôtel ?

Oui, à condition que le client puisse atteindre le serveur. Avec un CGNAT côté client, ajoutez `PersistentKeepalive = 25` pour maintenir la traduction NAT active. Si le réseau hôte filtre l’UDP 51820, vous pouvez faire écouter le serveur sur UDP 53 ou 443 (ports rarement bloqués). Pour les réseaux qui ne tolèrent que TCP, utilisez **udp2raw** ou **wstunnel** en encapsulation.

### Combien de peers un serveur WireGuard peut-il gérer ?

Pratiquement illimité. WireGuard maintient une simple table de hachage en kernel space ; un VPS modeste (2 vCPU, 4 Go) gère sans problème 1 000+ peers actifs. La limite vient plutôt de la bande passante et du CPU pour le chiffrement : comptez 800-1 200 Mbit/s saturable par cœur ChaCha20. Pour aller au-delà, scalez horizontalement avec plusieurs serveurs WireGuard derrière un load balancer L4 (HAProxy UDP, ECMP).

### WireGuard est-il compatible avec IPv6 ?

Totalement. Vous pouvez utiliser WireGuard pour transporter du IPv4, du IPv6, ou les deux simultanément. La syntaxe est identique : déclarez `Address = 10.66.0.1/24, fd42::1/64` dans `[Interface]` et `AllowedIPs = 0.0.0.0/0, ::/0` dans `[Peer]`. Côté pare-feu, dupliquez les règles avec `ip6tables` ou utilisez `nftables` qui gère les deux familles dans une seule table.

### Existe-t-il une interface graphique pour gérer WireGuard ?

Oui, plusieurs projets matures : **wg-easy** (interface web Docker, idéale pour un usage perso), **WG-Dashboard**, **Firezone** (open source, SSO et politiques RBAC), et les solutions commerciales comme **Tailscale Admin Console** ou **NetBird Console**. Pour 1 à 10 utilisateurs, wg-easy lancé en un `docker run` suffit ; pour 50+ utilisateurs avec contrôle d’accès, Firezone ou NetBird sont préférables.

### Couverture associée

*Sources externes : Site officiel WireGuard, Quickstart WireGuard.com, Manuel wg(8) Linux, Manuel wg-quick(8), Tailscale. Tutoriel rédigé le 30 avril 2026, mis à jour le 23 septembre 2026 avec les dernières versions de wireguard-tools (1.0.20260223-2, Debian testing) et de WireGuard for Windows (1.1.1, 20 septembre 2026).*
