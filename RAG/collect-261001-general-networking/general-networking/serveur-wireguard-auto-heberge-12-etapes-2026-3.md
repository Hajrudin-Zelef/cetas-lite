---
id: collect-261001-general-networking/general-networking/serveur-wireguard-auto-heberge-12-etapes-2026-3
title: "Les blocs [Peer] des clients seront ajoutés ici à l'étape 7"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["benchmark", "diffusion", "open source"]
source: docs/RAG/collect-261001-general-networking/serveur-wireguard-auto-heberge-12-etapes-2026.md
source_anchor: ""
source_lines: [161, 206]
sha256: a9c715188efa443e897098dbf8a3997cb59992d8e7004c8b844401671720cdd0
---

# Les blocs [Peer] des clients seront ajoutés ici à l'étape 7

Le fichier `client1.conf` généré à l’étape 8 fonctionne de manière identique sur tous les systèmes, mais la procédure d’importation change selon la plateforme. Sur **Windows**, téléchargez l’application officielle depuis le site WireGuard — le pilote WireGuardNT est passé en version 0.11 et le client Windows en version 0.6 en avril 2026, avant que le paquet Windows n’atteigne officiellement la version 1.0 ce même mois, selon la liste de diffusion WireGuard —, ouvrez-la, cliquez sur « Importer un tunnel à partir d’un fichier » et sélectionnez votre `client1.conf`. L’activation se fait ensuite par un simple bouton bascule dans l’interface.

Sur **macOS**, la procédure est identique via l’application Mac App Store, avec la même option d’import de fichier. Sur **Linux**, pas besoin d’interface graphique : copiez le fichier dans `/etc/wireguard/wg0.conf` sur la machine cliente et démarrez le tunnel avec `wg-quick up wg0`, exactement comme pour un serveur. Sur **iOS** et **Android**, l’application officielle (disponible sur l’App Store et Google Play) propose un scan de QR code, ce qui évite de transférer manuellement le fichier de configuration. Côté Android, la build officielle en est à la version 1.0.20260315 depuis mars 2026 et nécessite Android 12 ou supérieur pour l’installation. Générez ce QR code depuis le serveur avec la commande suivante.

```
sudo apt install qrencode -y
qrencode -t ansiutf8 < client1.conf
```
Le QR code s'affiche directement dans le terminal SSH. Ouvrez l'application WireGuard sur votre téléphone, appuyez sur le bouton d'ajout de tunnel, choisissez « Scanner depuis un QR code », et pointez la caméra vers l'écran de votre terminal. La connexion s'établit en quelques secondes, sans jamais avoir à copier manuellement de clé privée sur l'appareil mobile.

## Combien coûte un VPN auto-hébergé par rapport à un abonnement commercial

La question du coût revient systématiquement dès qu'on compare l'auto-hébergement à un abonnement VPN classique. Un Raspberry Pi 4 ou 5 représente un investissement unique, sans coût récurrent au-delà de l'électricité consommée (quelques euros par an pour un appareil aussi peu gourmand). Un VPS d'entrée de gamme chez un hébergeur européen coûte généralement entre 3 et 6 euros par mois, soit un ordre de grandeur comparable à un abonnement VPN premium, mais avec un contrôle total sur l'infrastructure et sans limite de bande passante artificielle.

| Solution | Coût mensuel estimé | Contrôle des journaux | Nombre d'appareils | 
|---|---|---|---|
| Raspberry Pi 4/5 auto-hébergé | ~0,50-1 € (électricité) | Total, aucun tiers | Illimité (limité par la bande passante montante) | 
| VPS européen + WireGuard | 3-6 € | Total, aucun tiers | Illimité (limité par la bande passante montante) | 
| Abonnement VPN commercial premium | 5-12 € | Dépend de la politique du fournisseur | 5 à 10 appareils simultanés en général | 
| Tailscale (plan gratuit) | 0 € jusqu'à 3 utilisateurs / 100 appareils | Coordination via serveur tiers, trafic chiffré de bout en bout | Jusqu'à 100 appareils sur le plan gratuit | 

L'écart de coût n'est donc pas toujours spectaculaire en valeur absolue. Ce qui change réellement, c'est la maîtrise de l'infrastructure : avec un serveur auto-hébergé, vous savez exactement quelles données transitent, où elles sont stockées, et qui y a accès. C'est cet argument de souveraineté numérique, plus que l'économie financière, qui motive la majorité des utilisateurs techniques à franchir le pas de l'auto-hébergement.

## WireGuard vs OpenVPN vs IPsec : quel protocole choisir

WireGuard n'est pas le seul protocole VPN disponible, et il ne remplace pas nécessairement OpenVPN dans tous les scénarios. Le tableau ci-dessous synthétise les écarts de performance mesurés dans plusieurs études indépendantes publiées entre 2025 et 2026.

| Critère | WireGuard | OpenVPN (AES-256-GCM) | IPsec | 
|---|---|---|---|
| Débit sur liaison gigabit | ~940 Mb/s down / 920 Mb/s up | ~620 Mb/s down / 590 Mb/s up | Environ 30 % plus lent que WireGuard hors virtualisation | 
| Utilisation de la bande passante | 95-98 % | 65-80 % | Variable selon implémentation | 
| Lignes de code du projet | ~4 000 | ~70 000+ | Dépend fortement de la pile (StrongSwan, etc.) | 
| Intégration noyau Linux | Native depuis le noyau 5.6 | Espace utilisateur (module TUN/TAP) | Native (mais configuration complexe) | 
| Complexité de configuration | Faible (clés Curve25519) | Élevée (certificats X.509, PKI) | Élevée (IKE, associations de sécurité) | 
| Consommation CPU | ~3 % en charge | ~17 % en charge | Modérée à élevée | 

Dans un contexte cloud (Azure notamment), certains tests notent que OpenVPN et WireGuard obtiennent des débits comparables, autour de 280 à 290 Mb/s, et qu'OpenVPN se comporte parfois mieux en conditions de forte latence réseau. Un autre benchmark, mené par ProxyPoland en avril 2026, situe WireGuard à 479 Mb/s dans ses propres conditions de test, un rappel que les chiffres varient fortement selon l'infrastructure réseau et la méthodologie retenue. Cette nuance compte si votre serveur est très éloigné géographiquement de vos clients habituels. Pour un usage domestique classique en France ou en Europe, l'écart de performance reste en faveur de WireGuard dans la quasi-totalité des scénarios testés.

## WireGuard face aux solutions mesh : Tailscale, Headscale et Netbird

Si la configuration manuelle décrite plus haut vous paraît fastidieuse à répéter pour chaque appareil, plusieurs outils construisent une couche de gestion par-dessus WireGuard. Tailscale automatise l'échange de clés et la découverte des pairs via un service cloud, ce qui simplifie énormément l'ajout d'appareils mais réintroduit une dépendance à un tiers. Headscale est l'implémentation open source auto-hébergeable du serveur de coordination de Tailscale, pensée précisément pour ceux qui veulent garder le contrôle total. Netbird suit une logique similaire, avec une interface de gestion des pairs et des règles d'accès en équipe.

Pour un usage individuel avec deux ou trois appareils, la configuration manuelle de ce tutoriel reste largement suffisante et évite d'ajouter une dépendance logicielle supplémentaire. Pour une petite équipe ou un homelab avec de nombreux appareils qui rejoignent et quittent le réseau régulièrement, un outil comme `wg-easy` (interface web pour gérer les pairs) ou Headscale devient rapidement plus pratique à maintenir qu'un fichier `wg0.conf` édité à la main. Ces outils restent tous des surcouches à l'implémentation WireGuard du noyau : ils ne remplacent pas le protocole lui-même, ils en simplifient l'administration.

## 5 erreurs fréquentes à éviter lors de l'installation

