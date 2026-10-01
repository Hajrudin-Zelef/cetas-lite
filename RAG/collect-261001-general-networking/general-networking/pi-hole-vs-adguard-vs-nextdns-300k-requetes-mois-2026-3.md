---
id: collect-261001-general-networking/general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026-3
title: "Temps de requete affiche sur la ligne \"Query time\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026.md
source_anchor: ""
source_lines: [99, 136]
sha256: 6057d52c3ccd0de480af4d46a21991bac9323152272685e32ba4992fd45e5bf3
---

# Temps de requete affiche sur la ligne "Query time"

## Listes de blocage et faux positifs : une gestion très différente

La qualité d’un bloqueur DNS ne se limite pas à son code. Elle dépend surtout de la qualité et de la fraîcheur des listes qu’il applique, et c’est justement là que les trois outils divergent le plus dans leur philosophie. Pi-hole ne fournit aucune liste par défaut au premier lancement : c’est à l’utilisateur d’ajouter manuellement les listes de son choix, les plus connues étant celles de StevenBlack, OISD ou HaGeZi, chacune avec un niveau d’agressivité différent. Cette liberté totale a un revers, un utilisateur pressé qui coche toutes les listes disponibles sans réfléchir se retrouve vite avec des faux positifs, des sites légitimes bloqués parce qu’ils partagent un sous-domaine ou une infrastructure CDN avec un domaine publicitaire.

AdGuard Home adopte une approche plus encadrée. L’outil propose dès l’installation une sélection de listes curatées et régulièrement testées par l’équipe du projet, réduisant le risque de faux positifs massifs pour un utilisateur qui ne touche à rien. La liste blanche (whitelist) et la liste noire personnalisée s’éditent depuis la même interface que les statistiques de blocage, ce qui rend le diagnostic d’un faux positif plus rapide : on voit immédiatement quelle règle a bloqué quel domaine et on peut créer une exception en un clic depuis le journal des requêtes.

NextDNS s’écarte le plus du modèle des listes statiques. Le service combine des listes communautaires classiques avec ses propres signaux de détection, mis à jour côté serveur sans que l’utilisateur ait à s’en soucier. L’avantage est la réactivité face à un nouveau domaine de phishing qui vient d’apparaître, la détection ne dépend pas de la prochaine mise à jour d’une liste statique téléchargée localement. L’inconvénient est la transparence : contrairement à une liste communautaire publique sur GitHub que n’importe qui peut auditer, une partie de la logique de détection de NextDNS reste propriétaire, ce qui demande un peu plus de confiance de la part de l’utilisateur exigeant sur ce point précis.

Dans les trois cas, corriger un faux positif suit la même logique de base : identifier le domaine bloqué à tort dans les journaux de requêtes, puis l’ajouter à une liste blanche locale qui prime sur les listes de blocage générales. AdGuard Home et NextDNS simplifient cette étape avec un bouton dédié directement dans l’historique des requêtes, quand Pi-hole demande de passer par l’onglet listes blanches et de coller le domaine manuellement.

## Chiffrement DNS : DoH, DoT et DoQ, lequel protège vraiment votre trafic

Sans chiffrement, chaque requête DNS circule en clair sur le réseau, ce qui permet à votre fournisseur d’accès, à l’opérateur du Wi-Fi public ou à tout intermédiaire réseau de voir exactement quels domaines vous consultez, et parfois de les modifier ou de les rediriger. DNS-over-HTTPS (DoH), DNS-over-TLS (DoT) et le plus récent DNS-over-QUIC (DoQ) chiffrent ce trafic pour empêcher cette visibilité.

AdGuard Home est la seule des trois solutions auto-hébergées à gérer ces trois protocoles nativement, aussi bien pour les appareils qui l’interrogent que pour ses propres requêtes vers les résolveurs en amont. Pi-hole, lui, chiffre uniquement si vous ajoutez Unbound (résolution récursive complète, sans dépendre d’un tiers) ou cloudflared (proxy vers un résolveur DoH externe comme Cloudflare). Cette étape supplémentaire n’est pas insurmontable mais elle ajoute de la complexité et un composant de plus à maintenir à jour pour rester protégé contre d’éventuelles vulnérabilités.

NextDNS chiffre par défaut toutes les requêtes entre vos appareils et son infrastructure via DoH ou DoT, ce qui est cohérent avec un modèle où la confiance se déplace du réseau local vers le fournisseur cloud. La question n’est alors plus « mes requêtes sont-elles chiffrées » mais « en qui ai-je confiance pour les traiter », un arbitrage propre à chaque utilisateur.

## RGPD et souveraineté des données : où vont vos requêtes DNS

Pour un lecteur français ou européen, la question de la localisation des données pèse autant que les fonctionnalités. Avec Pi-hole et AdGuard Home, la réponse est la plus simple possible : aucune requête ne sort du réseau local avant d’atteindre le résolveur en amont choisi par l’utilisateur. Aucun tiers, aucun hébergeur, aucune juridiction étrangère à considérer sur cette partie du flux.

NextDNS adopte une posture différente mais documentée. Le service cite explicitement le RGPD dans sa politique de confidentialité pour qualifier ses sous-traitants techniques, et précise que si aucune donnée n’est demandée spécifiquement par l’utilisateur, rien n’est journalisé. Quand la journalisation est activée pour bénéficier des statistiques et de l’historique, l’utilisateur garde un contrôle complet sur la durée de rétention et peut exporter ou supprimer ses données à tout moment. Cette politique se rapproche des standards attendus par la CNIL en matière de minimisation des données, même si l’entité légale NextDNS Inc. reste incorporée aux États-Unis et non dans l’Union européenne.

Ce contexte réglementaire prend un relief particulier en France depuis l’entrée en application de la directive NIS2, qui pousse les entités concernées à documenter précisément où transitent leurs flux réseau critiques. Notre article sur la situation de la France face à NIS2 détaille les obligations qui pèsent désormais sur les infrastructures numériques, y compris pour des briques a priori aussi anodines qu’un résolveur DNS. Les administrations et entreprises soumises à ces obligations, ou simplement les utilisateurs les plus attentifs à leur souveraineté numérique, peuvent consulter les recommandations générales de l’ANSSI sur le choix des prestataires numériques.

## 5 usages réels observés en 2026

- **Le foyer connecté avec Raspberry Pi dédié.** Un particulier installe Pi-hole sur un Raspberry Pi 4 branché en permanence, redirige le DHCP de sa box pour que tous les appareils du foyer passent par lui, et ajoute Unbound après quelques semaines pour chiffrer les requêtes sortantes.
- **Le passionné d’auto-hébergement (homelab).** Un utilisateur qui fait déjà tourner Nextcloud, Vaultwarden ou Jellyfin sur un serveur Proxmox ajoute AdGuard Home comme conteneur Docker supplémentaire, profitant du chiffrement natif sans composant tiers à gérer en plus de sa pile existante.
- **Le télétravailleur multi-sites.** Un consultant qui alterne domicile, espace de coworking et déplacements professionnels choisit NextDNS pour conserver la même politique de filtrage et de sécurité sur son ordinateur portable et son téléphone, quel que soit le réseau utilisé, sans dépendre d’une connexion VPN vers son domicile.
- **La petite structure avec un administrateur réseau.** Une TPE de moins de dix postes déploie AdGuard Home sur un mini-PC et configure des règles par adresse IP pour appliquer des listes différentes au réseau invité, au poste de la comptabilité et aux imprimantes connectées.
- **Le parent qui gère le contrôle parental à distance.** Un parent utilise l’application mobile NextDNS pour ajuster en temps réel les horaires d’accès et les catégories bloquées sur le téléphone de son enfant, changement appliqué immédiatement sans avoir à toucher à un boîtier physique resté à la maison.

## Quel outil choisir selon votre profil

Aucun des trois outils ne s’impose universellement. Voici les recommandations qui ressortent le plus nettement de ce comparatif selon votre situation.

