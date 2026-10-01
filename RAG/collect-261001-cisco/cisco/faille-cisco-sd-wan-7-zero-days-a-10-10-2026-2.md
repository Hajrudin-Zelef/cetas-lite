---
id: collect-261001-cisco/cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026-2
title: "Commande observée par Mandiant (avril 2026)"
domain: cisco
role: reference
task: reference
actors: ["CISA", "Google"]
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026.md
source_anchor: ""
source_lines: [45, 100]
sha256: ee76b984cc60c658f65f6ba3a252708cabf8146618a2cedbdad04f3a916495f3
---

# Commande observée par Mandiant (avril 2026)

Comprendre pourquoi ces failles sont si dangereuses suppose de comprendre ce qu’est réellement un contrôleur SD-WAN. Dans une architecture Cisco Catalyst SD-WAN, le Manager (anciennement vManage) et le Controller constituent le plan de gestion et de contrôle : le « cerveau » qui décide comment le trafic circule entre le siège, les agences, les datacenters et le cloud. Compromettre ce cerveau, ce n’est pas pirater un serveur parmi d’autres – c’est prendre la main sur l’ensemble du système nerveux du réseau d’entreprise.

### Étape 1 : le pair malveillant et la manipulation de NETCONF

La première étape exploite le contournement d’authentification. En se faisant passer pour un pair légitime, l’attaquant fait apparaître un composant SD-WAN temporaire, contrôlé par lui, qui peut mener des actions de confiance au sein des plans de gestion et de contrôle. Via NETCONF, il peut alors lire ou modifier la configuration de n’importe quel équipement du fabric. Pour une entreprise multisite, cela signifie qu’un seul point d’entrée exposé sur Internet peut compromettre des dizaines, voire des centaines de sites distants en cascade.

### Étape 2 : le retour arrière logiciel et l’escalade vers root

Le compte obtenu via le contournement n’est pas root. Pour franchir ce dernier palier, les attaquants ont recours à une technique élégante et redoutable : le retour arrière logiciel. En s’appuyant sur le mécanisme de mise à jour intégré, ils déploient temporairement une version antérieure et vulnérable du logiciel, exploitent CVE-2022-20775 (une faille de traversée de chemin en ligne de commande, notée 7,8) pour obtenir root, puis restaurent la version d’origine. L’équipement paraît à jour, mais l’attaquant dispose désormais des privilèges les plus élevés. C’est cette manœuvre qui transforme un accès administrateur en contrôle total et furtif de la machine.

## La 7e faille : CVE-2026-20245 et l’exécution de commandes root

La vulnérabilité la plus récente, CVE-2026-20245, a été signalée à Cisco par Mandiant, la filiale de Google spécialisée dans la réponse aux incidents. D’après le Google Threat Intelligence Group, il s’agit d’une élévation de privilèges dans l’interface en ligne de commande du SD-WAN Manager : la fonction de téléversement de fichiers ne filtre pas correctement les données malveillantes. Un attaquant local authentifié peut ainsi exécuter des commandes arbitraires en tant que root.

Mandiant a documenté un cas d’exploitation réelle survenu dès avril 2026, sur l’infrastructure SD-WAN d’un fournisseur de services. L’attaquant a utilisé la commande `request tenant-upload` pour importer un fichier CSV piégé, créant un compte non autorisé baptisé `troot` doté de privilèges root :

```
# Commande observée par Mandiant (avril 2026)
request tenant-upload tenant-list /home/admin/evil_tenant.csv vpn 0
# Resultat : creation d'un compte "troot" avec privileges root
```
Le fait que CVE-2026-20245 exige un accès local authentifié n’est pas rassurant : il s’inscrit précisément dans la logique de chaîne décrite plus haut. Les contournements à 10/10 fournissent le pied dans la porte ; CVE-2026-20245 (ou CVE-2022-20775) fournit l’escalade vers root. Les chercheurs Chester Sng, Pete Boonyakarn, Logeswaran Nadarajan et Lukasz Lamparski, de Mandiant, ont mené l’investigation qui a permis d’identifier ce nouveau vecteur d’élévation.

## Un PoC public déclenche une ruée de dix groupes d’attaquants

Si UAT-8616 est l’acteur le plus discret, il est loin d’être le seul. En mars 2026, un code de démonstration (proof of concept, ou PoC) a été publié par ZeroZenX Labs. Sa mise en ligne a provoqué un effet d’aubaine immédiat : selon Tenable, au moins dix groupes d’attaquants supplémentaires se sont mis à exploiter la chaîne CVE-2026-20133 / CVE-2026-20128 / CVE-2026-20122 dans les jours qui ont suivi.

La mécanique de cette chaîne, documentée par VulnCheck, illustre comment trois failles de gravité moyenne, combinées, deviennent une arme d’accès initial :

- **CVE-2026-20133** (divulgation d’informations, non authentifiée) fait fuiter le fichier d’identifiants DCA ;
- **CVE-2026-20128** (lecture de fichier authentifiée) lit le secret utilisateur DCA ainsi exposé ;
- **CVE-2026-20122** (écrasement de fichier authentifié) téléverse un webshell via l’API.

À la différence d’UAT-8616, ces groupes opportunistes déploient des charges bruyantes : webshells, frameworks de red team détournés, mineurs de cryptomonnaie et voleurs d’identifiants. VulnCheck note d’ailleurs qu’un PoC largement partagé était mal attribué – il prétendait exploiter CVE-2026-20127 alors qu’il chaînait en réalité les trois autres failles. Cette confusion illustre un risque bien connu : la divulgation publique d’un code d’exploitation démocratise l’attaque et abaisse drastiquement le niveau de compétence requis pour compromettre un équipement exposé.

## Combien d’équipements exposés ? Les chiffres de Censys et VulnCheck

La question que se posent tous les RSSI européens est simple : combien d’équipements sont réellement atteignables depuis Internet ? Les moteurs de recherche d’exposition apportent des réponses convergentes, mais nuancées. Selon VulnCheck, Censys et Shodan dénombrent entre 450 et 550 instances de SD-WAN Manager exposées, tandis que ZoomEye en compte environ 275 et FOFA plus de 1 000. Un relevé Censys de mai 2026 portait le total à près de 2 000 équipements Catalyst SD-WAN visibles en ligne, plans de gestion et de contrôle confondus.

| Estimations d’exposition sur Internet des équipements Cisco Catalyst SD-WAN. Sources : VulnCheck, Censys. |  |  | 
|---|---|---|
| Source | Instances exposées | Périmètre | 
|---|---|---|
| Censys / Shodan | 450 – 550 | SD-WAN Manager | 
| ZoomEye | ~275 | SD-WAN Manager | 
| FOFA | > 1 000 | SD-WAN Manager | 
| Censys (mai 2026) | ~2 000 | Catalyst SD-WAN (global) | 

Ces chiffres, en apparence modestes, sont trompeurs. Le plan de gestion d’un SD-WAN n’est pas censé être exposé sur Internet : chaque instance visible représente une organisation entière – et l’ensemble de ses sites distants – à un pas d’une compromission totale. L’avis Censys sur CVE-2026-20127 souligne par ailleurs qu’une part significative des hôtes exposés laisse aussi ouverts les ports SSH (22) ou NETCONF (830), élargissant d’autant la surface d’attaque. Si la majorité des instances se concentrent aux États-Unis, l’Europe en héberge une proportion loin d’être négligeable, notamment chez les opérateurs télécoms et les grandes entreprises multisites.

## Directive d’urgence de la CISA et réponse de Cisco

La gravité de la situation a poussé l’agence américaine de cybersécurité, la CISA, à publier la directive d’urgence 26-03 le 14 mai 2026, imposant aux agences fédérales de remédier à CVE-2026-20182 dans des délais très courts, avec une échéance fixée au 17 mai. Les directives d’urgence sont rares : la CISA n’en émet que lorsqu’une menace présente un risque « inacceptable » pour les réseaux fédéraux. Chaque CVE de la campagne a par ailleurs été inscrit au catalogue KEV, ce qui, aux États-Unis, transforme le correctif en obligation légale pour les administrations.

Côté éditeur, Cisco a publié des correctifs pour l’ensemble des branches maintenues. Le tableau suivant reprend, pour la faille CVE-2026-20182, les premières versions corrigées telles que listées dans l’avis officiel. Les organisations exécutant une version antérieure à la 20.9 doivent migrer, et trois branches (20.11, 20.13 et 20.14) ont atteint leur fin de maintenance logicielle – un facteur aggravant pour les parcs vieillissants.

