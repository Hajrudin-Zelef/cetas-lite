---
id: collect-261001-general-networking/general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026-2
title: "Temps de requete affiche sur la ligne \"Query time\""
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["benchmarks", "diffusion", "open source"]
source: docs/RAG/collect-261001-general-networking/pi-hole-vs-adguard-vs-nextdns-300k-requetes-mois-2026.md
source_anchor: ""
source_lines: [33, 98]
sha256: 7a578b2ca4c19570e4aebc1410917c6b441b6c0d49b7b030fede4a8214d37bc4
---

# Temps de requete affiche sur la ligne "Query time"

NextDNS a été fondé en 2019 par deux ingénieurs français, Romain Cointepas et Olivier Poitrey, tous deux passés par Dailymotion, où Poitrey occupait le poste de directeur technique. Poitrey a ensuite dirigé l’ingénierie du réseau de diffusion de contenu Open Connect chez Netflix avant de lancer NextDNS. L’entité juridique, NextDNS Inc., est toutefois incorporée dans le Delaware, aux États-Unis, et non en France, un point à connaître pour qui cherche une solution strictement européenne au sens juridique. Les deux fondateurs sont par ailleurs à l’origine de dns0.eu, un résolveur DNS à but non lucratif basé en Europe, ce qui illustre leur implication continue dans l’écosystème DNS européen.

Sur le plan technique, NextDNS va au-delà des listes statiques en combinant plusieurs sources de blocage (listes communautaires, signatures de télémétrie, heuristiques anti-phishing) mises à jour en continu côté serveur. L’offre gratuite autorise 300 000 requêtes par mois et un nombre illimité d’appareils et de profils de configuration. Passé ce quota, NextDNS ne coupe pas la résolution : le service continue de répondre aux requêtes comme un DNS classique non filtrant jusqu’au mois suivant, un détail important qui évite toute coupure d’accès à internet. Les applications mobiles officielles pour iOS et Android, avec synchronisation de la configuration entre appareils, constituent un vrai avantage face à Pi-hole et AdGuard Home, qui ne proposent ni l’une ni l’autre en natif.

## Tableau comparatif complet : Pi-hole vs AdGuard Home vs NextDNS

Le tableau ci-dessous synthétise les critères techniques et pratiques qui distinguent les trois solutions à la mi-2026.

| Critère | Pi-hole | AdGuard Home | NextDNS | 
|---|---|---|---|
| Licence | EUPL 1.2 (open source) | GPLv3 (open source) | Propriétaire | 
| Modèle d’hébergement | Auto-hébergé uniquement | Auto-hébergé uniquement | Cloud uniquement | 
| Porteur du projet | Communauté open source | Société AdGuard | NextDNS Inc. (fondateurs français) | 
| Branche actuelle (mi-2026) | v6.x | v0.107.x stable / v0.108 bêta | Service cloud, pas de version | 
| Chiffrement DNS côté client | Non natif | Natif (DoH, DoT, DoQ, DNSCrypt) | Natif (DoH, DoT) | 
| Chiffrement DNS en amont | Nécessite Unbound ou cloudflared | Natif, activable sans module externe | Géré côté serveur | 
| Application mobile dédiée | Non | Non (interface web uniquement) | Oui (iOS et Android) | 
| Synchronisation multi-appareils | Non | Non | Oui | 
| Contrôle parental intégré | Manuel (listes et regex) | Natif (filtres, SafeSearch, horaires) | Natif (profils cloud) | 
| Règles par client | Par groupes | Par IP, CIDR ou MAC | Par appareil (profil cloud) | 
| Paquet natif pour routeur OpenWrt | Non officiel | Oui | Sans objet (cloud) | 
| Matériel requis | Oui (Raspberry Pi, VM, Docker) | Oui (mêmes plateformes) | Aucun | 
| Empreinte mémoire estimée | Environ 80 à 90 Mo (plus avec Unbound) | Environ 50 à 60 Mo | Nulle en local | 
| Traitement des données | 100 % local, aucune donnée ne sort | 100 % local, aucune donnée ne sort | Cloud, conforme RGPD, pas de journalisation par défaut | 

Ce tableau met en évidence l’arbitrage central du comparatif. Pi-hole et AdGuard Home offrent une souveraineté totale sur les données, puisqu’aucune requête DNS ne quitte le réseau local. NextDNS sacrifie une partie de cette souveraineté contre zéro maintenance et une couverture mobile native, tout en s’appuyant sur une politique de confidentialité conforme au RGPD et sur un traitement sans journalisation par défaut.

## Tarifs 2026 : gratuit, freemium ou abonnement

Contrairement à beaucoup d’outils de cybersécurité, aucune des trois solutions n’exige de payer pour un usage domestique de base. Les différences apparaissent surtout au-delà du foyer, sur la quantité de requêtes ou le nombre de comptes gérés.

| Offre | Prix | Limites | 
|---|---|---|
| Pi-hole (logiciel) | 0 €, gratuit à vie | Aucune limite de requêtes | 
| Matériel Pi-hole (Raspberry Pi 4/5 typique) | Environ 60 à 100 € une fois | Investissement matériel unique | 
| AdGuard Home (auto-hébergé) | 0 €, gratuit à vie | Aucune limite de requêtes | 
| AdGuard DNS (service cloud public, produit distinct) | 0 € pour les serveurs publics | Formules personnelles payantes disponibles, tarifs non publiés en clair | 
| NextDNS Gratuit | 0 € | 300 000 requêtes par mois, appareils illimités | 
| NextDNS Pro | 1,99 €/mois ou 19,90 €/an | Requêtes illimitées, usage personnel et familial | 
| NextDNS Business | 19,90 €/mois ou 199 €/an | Jusqu’à 50 employés, support par e-mail | 
| NextDNS Education | 19,90 €/mois ou 199 €/an | Jusqu’à 250 étudiants, support par e-mail | 

Le seul vrai coût d’entrée pour Pi-hole ou AdGuard Home est matériel, et souvent nul si vous recyclez un boîtier déjà présent chez vous : un NAS, un mini-PC ou même une machine virtuelle sur un serveur existant suffisent. NextDNS reste gratuit tant que le foyer ne dépasse pas 300 000 requêtes mensuelles, un plafond confortable pour un usage individuel mais qui se franchit vite avec plusieurs appareils actifs en permanence. L’abonnement Pro à 1,99 € par mois, détaillé sur la page tarifs officielle de NextDNS, lève cette limite pour le prix d’un café.

## Benchmarks : latence, RAM et charge CPU

Les trois outils appliquent le même principe de filtrage, ce qui limite les écarts de performance bruts quand ils utilisent des listes équivalentes. La documentation officielle des projets et les retours de la communauté self-hosting convergent toutefois sur quelques différences mesurables. Sur un Raspberry Pi 4, Pi-hole seul consomme environ 80 à 90 Mo de RAM, un chiffre qui grimpe sensiblement dès qu’on ajoute Unbound pour obtenir un chiffrement DNS complet. AdGuard Home, dont le chiffrement est intégré au binaire principal, tourne généralement autour de 50 à 60 Mo sur le même matériel. Ni l’un ni l’autre ne sollicite significativement le CPU en usage courant sur du matériel récent.

NextDNS élimine cette question puisque rien ne tourne localement, mais déplace le sujet vers la latence réseau. Le service s’appuie sur un réseau de points de présence répartis pour limiter le nombre de sauts avant résolution, avec des temps de réponse généralement dans la même fourchette qu’un résolveur local bien configuré pour un utilisateur en France ou en Europe de l’Ouest. La différence pratique se joue surtout hors du domicile : un résolveur auto-hébergé n’est directement joignable que sur le réseau local sauf configuration VPN supplémentaire, alors que NextDNS répond identiquement depuis n’importe quel réseau mobile ou Wi-Fi public.

Pour mesurer vous-même la latence de résolution d’un domaine bloqué ou autorisé, la commande `dig` reste l’outil de référence sous Linux et macOS :

```
dig @192.168.1.10 doubleclick.net +stats
# Temps de requete affiche sur la ligne "Query time"
# Comparez le meme domaine via un resolveur non filtrant
dig @1.1.1.1 doubleclick.net +stats
```
Pour tester une résolution chiffrée en DNS-over-HTTPS, un simple appel curl vers un point de terminaison JSON DoH permet de vérifier que le chiffrement amont fonctionne réellement plutôt que de se fier à une case cochée dans une interface :

```
curl -s -H 'accept: application/dns-json' \
  "https://dns.google/resolve?name=doubleclick.net&type=A"
```
En pratique, la différence de charge entre Pi-hole et AdGuard Home ne devient perceptible que sur du matériel très contraint, comme un Raspberry Pi Zero. Sur un Raspberry Pi 4 ou 5, une VM avec 1 Go de RAM ou un mini-PC, les deux solutions gèrent sans effort plusieurs centaines de requêtes par seconde, largement suffisant pour un foyer ou une petite structure.

