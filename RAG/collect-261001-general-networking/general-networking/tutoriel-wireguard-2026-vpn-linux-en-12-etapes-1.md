---
id: collect-261001-general-networking/general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes-1
title: "Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["benchmark", "benchmarks", "cyber", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-wireguard-2026-vpn-linux-en-12-etapes.md
source_anchor: ""
source_lines: [1, 46]
sha256: 467d0a6721864e9e20cebb948c361fda4a05c1ae2171a4033cb90c0e868974f9
---

# Ubuntu 22.04 / 24.04 LTS et Debian 12 / 13

**WireGuard** s’est imposé en 2026 comme le standard de fait pour les VPN modernes, avec environ **4 000 lignes de code** contre plus de 100 000 pour OpenVPN et près de 600 000 pour IPsec/strongSwan. Intégré au noyau Linux depuis la version **5.6** (mars 2020) et toujours activement maintenu, ce protocole signé Jason A. Donenfeld utilise la cryptographie moderne — **ChaCha20-Poly1305**, **Curve25519**, **BLAKE2s** — orchestrée par le framework Noise. Résultat : selon un benchmark publié en mars 2026 par Zeroday Cyber Academy, un débit qui grimpe jusqu’à **8,7 Gbit/s** sur une liaison à 10 Gbit/s, une poignée de main inférieure à **1 ms** et une consommation batterie nettement inférieure à OpenVPN sur smartphone.

Ce tutoriel, mis à jour le **23 septembre 2026** — juste après la sortie de **WireGuard for Windows 1.1.1** le 20 septembre 2026 et le passage du paquet **wireguard-tools 1.0.20260223-2** en testing chez Debian — et structuré en **12 étapes** réalisables en environ **20 minutes**, vous guide pas à pas pour déployer un serveur WireGuard en production sur Linux, configurer plusieurs clients (Linux, Windows, macOS, iOS, Android), automatiser la gestion des clés, et industrialiser la solution avec des outils comme **Tailscale** ou **Headscale**. Vous obtiendrez à la fin un VPN site à site fonctionnel, un mesh peer-to-peer optionnel et un script complet pour générer des fichiers de configuration client en quelques secondes.

## Pourquoi choisir WireGuard en 2026 plutôt qu’OpenVPN ou IPsec

Le marché du VPN d’entreprise et personnel en 2026 est dominé par trois protocoles : **WireGuard**, **OpenVPN** et **IPsec/IKEv2**. WireGuard l’emporte sur trois axes critiques pour les équipes DevOps et SRE : simplicité opérationnelle, performances brutes et surface d’attaque minimale. Sa base de code, auditable par une seule personne en quelques jours, contraste avec les centaines de milliers de lignes des alternatives historiques, ce qui réduit drastiquement le risque de vulnérabilité critique.

Côté performances, WireGuard fonctionne en **espace noyau** sur Linux (module mainline depuis le kernel 5.6), évitant les coûteux changements de contexte qui pénalisent OpenVPN, lequel reste un démon en espace utilisateur. Les benchmarks publiés depuis 2021 par Phoronix, confirmés par un comparatif GeekSynapse de mars 2026 qui mesure un avantage de **2 à 4 fois** sur du matériel identique, et complétés par un nouveau comparatif publié en **juin 2026** qui chiffre l’écart à **3 fois plus rapide** qu’OpenVPN, montrent un écart net en débit comme en latence. Sur un VPS modeste (2 vCPU, 4 Go de RAM), WireGuard sature facilement une liaison **1 Gbit/s** — Zeroday Cyber Academy y a mesuré en mars 2026 un débit de **940 à 960 Mbit/s** — là où OpenVPN plafonne souvent autour de **250 Mbit/s**.

Le troisième argument décisif : la **mobilité native**. WireGuard maintient les sessions actives même quand le client change d’IP (Wi-Fi → 5G → Wi-Fi public). C’est ce qui a permis à **Tailscale**, mesh VPN bâti sur WireGuard, de séduire des dizaines de milliers d’entreprises depuis 2020 et de lever plus de 160 millions de dollars en série C en 2024 : une enquête homelab menée en 2025 et relayée par Phoronix en juin 2026 recense d’ailleurs **41 %** d’utilisateurs passant par Tailscale contre **28 %** configurant WireGuard directement en ligne de commande. **NetBird**, alternative open source, et **Headscale**, serveur de contrôle Tailscale auto-hébergé, élargissent l’écosystème en 2026.

| Critère | WireGuard | OpenVPN | IPsec/IKEv2 | 
|---|---|---|---|
| Lignes de code (cœur) | ~4 000 | ~100 000 | ~600 000 | 
| Espace d’exécution (Linux) | Noyau (depuis 5.6) | Espace utilisateur | Noyau + utilisateur | 
| Débit typique (1 Gbit/s lien) | 950+ Mbit/s | 200-300 Mbit/s | 700-900 Mbit/s | 
| Temps de poignée de main | < 1 ms | 100-300 ms | 50-200 ms | 
| Port par défaut | UDP 51820 | UDP 1194 / TCP 443 | UDP 500/4500 | 
| Cryptographie | ChaCha20-Poly1305, Curve25519 | AES-256-GCM, RSA | AES-256-GCM, ECDSA | 
| Roaming natif (changement IP) | Oui | Non (reconnect) | MOBIKE (RFC 4555) | 
| Configuration moyenne | 10-15 lignes | 50-100 lignes | 100-200 lignes | 
| Audit sécurité réalisable | 1 personne / quelques jours | Équipe / plusieurs mois | Équipe / 6+ mois | 

## Prérequis matériels, logiciels et versions WireGuard 2026

Avant de plonger dans la configuration, vérifiez que votre environnement remplit les conditions suivantes. Ce tutoriel a été validé sur **Ubuntu Server 24.04 LTS** et **Debian 13 (Trixie)**, mais les commandes restent identiques sur Rocky Linux 9, Fedora 41 et Alma Linux 9 — soit, avec les autres distributions couvertes par ce guide, une dizaine d’environnements testés au total, dont au moins **8 distributions majeures** embarquant nativement le module noyau WireGuard (kernel ≥ 5.6) depuis avril 2026. Pour le kernel embarqué, le module WireGuard est désormais inclus en standard ; vous n’avez donc plus besoin du paquet `wireguard-dkms` historique. Côté outillage userspace, la version de référence reste **wireguard-tools 1.0.20260223** — répertoriée comme version stable par Wikipedia dès le 23 février 2026 — mais son empaquetage a continué d’évoluer : la révision **1.0.20260223-2** a été acceptée dans la branche *unstable* de Debian le 27 août 2026, puis promue en *testing* en septembre 2026 selon le suivi officiel des paquets Debian : assurez-vous que `wg --version` renvoie bien ce numéro ou une version plus récente avant de poursuivre.

- **Serveur** : VPS ou machine dédiée, minimum 1 vCPU et 1 Go de RAM (recommandé : 2 vCPU, 2 Go).
- **Système d’exploitation** : Linux avec kernel ≥ 5.6 (Ubuntu 22.04+, Debian 12+, RHEL 9+).
- **Paquet wireguard-tools** : version 1.0.20210914 ou supérieure (incluant`wg` et`wg-quick` ).
- **Client Android** : WireGuard 1.0.20260315 (Play Store, mars 2026).
- **Client iOS** : WireGuard 1.0.16+ (App Store).
- **Client Windows** : WireGuard for Windows 1.1.1+ (Windows 10/11/Server 2019/2022/2025), publié le 20 septembre 2026.
- **Client macOS** : WireGuard for macOS 1.0.16+ (Mac App Store, compatible macOS 11+).
- **Accès root ou sudo** : indispensable pour charger le module et configurer iptables/nftables.
- **IP publique fixe** (ou DDNS) sur le serveur, et**port UDP 51820** ouvert dans le pare-feu opérateur.
- **Forwarding IPv4/IPv6 activable** via sysctl.
- **Compétences : Linux niveau intermédiaire** , manipulation d’iptables/nftables, lecture de logs systemd.

Si vous travaillez sur un Raspberry Pi 5 (ARM64) ou un microserveur basse consommation, WireGuard reste parfaitement utilisable : le module noyau y est compilé et un Pi 5 atteint sans difficulté **700 à 900 Mbit/s** de débit chiffré, largement suffisant pour la plupart des tunnels résidentiels.

## Étape 1 : Installation de WireGuard sur Ubuntu 24.04 et Debian 13

Commencez par mettre à jour la base de paquets puis installez la suite `wireguard`. Sur les distributions modernes, ce méta-paquet tire automatiquement `wireguard-tools` et le module noyau (qui n’a plus besoin de DKMS depuis le kernel 5.6).

