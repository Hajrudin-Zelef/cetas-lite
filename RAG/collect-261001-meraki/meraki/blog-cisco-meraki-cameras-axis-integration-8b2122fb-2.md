---
id: collect-261001-meraki/meraki/blog-cisco-meraki-cameras-axis-integration-8b2122fb-2
title: "blog-cisco-meraki-cameras-axis-integration-8b2122fb"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-meraki/blog-cisco-meraki-cameras-axis-integration-8b2122fb.md
source_anchor: ""
source_lines: [43, 110]
sha256: 3c6cd2e1168d05d35a2f5fed3d79f85fa2987f640f171777953bd0be10dcf2bc
---

# blog-cisco-meraki-cameras-axis-integration-8b2122fb

La sécurité ne se résume pas non plus à la présence du mot « cloud ». Comptes nominatifs, authentification multifacteur, droits d’export limités et retrait des accès inutiles restent des exigences de recette. Faites confirmer l’hébergement, les conditions de traitement des images et les engagements contractuels applicables. Les 30 jours inclus sont une caractéristique commerciale, pas une durée de conservation à appliquer automatiquement.
Comment se passe l’enrôlement ?
Le parcours documenté repose sur des comptes administrateurs Meraki et Axis, une licence, un modèle compatible et un accès Internet autorisé vers les services concernés. Le minimum général indiqué est AXIS OS 9.8, avec AXIS OS 11 ou ultérieur recommandé ; il faut confirmer la version requise pour le modèle et le niveau retenus.
Après activation de l’intégration et autorisation OAuth, trois méthodes sont proposées : découverte via des switches Cisco administrés dans Dashboard, numéro de série et clé propriétaire OAK, ou récupération d’appareils déjà dans Axis Cloud Connect. La méthode OAK/série nécessite aussi l’activation d’O3C ou l’action « one-click » sur l’équipement.
Un switch Cisco facilite donc la découverte, mais n’est pas une condition universelle de raccordement. Pour un parc hétérogène, commencez par distinguer les caméras automatiquement découvertes de celles à enrôler autrement.
Pour le pilote, préparez un inventaire avec la référence exacte, le firmware, le site, le mode d’alimentation, le stockage et les dépendances applicatives. Traitez les clés OAK et les identifiants locaux comme des secrets, pas comme des colonnes anodines d’un tableur partagé. Enfin, planifiez les mises à jour : une évolution d’AXIS OS peut nécessiter un redémarrage et interrompre momentanément la vidéo.
Quels produits Axis sont compatibles avec Meraki ?
La matrice Cisco Meraki, consultée le 18 septembre 2026, comprend les 90 références ci-dessous. La liste est regroupée pour faciliter la lecture, sans étendre la compatibilité à d’autres variantes d’une même famille. Toutes les références portent la marque AXIS.
87 références compatibles Essentials et Advantage
| Groupe de références | Modèles explicitement listés | 
|---|---|
| M42 | M4227-LVE, M4228-LVE | 
| M43 | M4317-PLR, M4317-PLVE, M4318-PLR, M4318-PLVE, M4327-P, M4328-P | 
| P1387 | P1387, P1387-B, P1387-BE, P1387-LE | 
| P1388 | P1388, P1388-B, P1388-BE, P1388-LE | 
| P14 | P1467-LE, P1468-LE, P1468-XLE | 
| P15 | P1518-E, P1518-LE | 
| P326 | P3267-LV, P3267-LVE, P3268-LV, P3268-LVE, P3268-SLVE | 
| P327 | P3275-LV, P3275-LVE, P3277-LV, P3277-LVE, P3278-LV, P3278-LVE | 
| P328 | P3285-LV, P3285-LVE, P3287-LV, P3287-LVE, P3288-LV, P3288-LVE | 
| P37 et P38 | P3735-PLE, P3737-PLE, P3738-PLE, P3747-PLVE, P3748-PLVE, P3827-PVE | 
| P47 | P4705-PLVE, P4707-PLVE, P4708-PLVE | 
| P91 | P9117-PV | 
| Q1656 | Q1656, Q1656-B, Q1656-BE, Q1656-BLE, Q1656-DLE, Q1656-LE | 
| Q1686 et Q17 | Q1686-DLE, Q1728, Q1728-LE | 
| Q18 | Q1800-LE, Q1800-LE-3, Q1805-LE, Q1806-LE, Q1808-LE, Q1809-LE | 
| Q19 | Q1961-TE, Q1961-XTE, Q1971-E, Q1972-E | 
| Q21 | Q2101-TE, Q2111-E, Q2112-E | 
| Q353 | Q3536-LVE, Q3538-LVE, Q3538-SLVE | 
| Q354 et Q355 | Q3546-LVE, Q3548-LVE, Q3556-LVE, Q3558-LVE | 
| Q36 et Q38 | Q3626-VE, Q3628-VE, Q3839-PVE | 
| Q48 et Q60 | Q4809-PVE, Q6020-E | 
| Q63 | Q6300-E, Q6355-LE, Q6358-LE | 
| Q93 et XFQ | Q9307-LV, XFQ1656 | 
Trois références compatibles Essentials uniquement
| Modèle | Essentials | Advantage | 
|---|---|---|
| P1245 Mk II | Oui | Non | 
| P1265 Mk II | Oui | Non | 
| P1275 Mk II | Oui | Non | 
Ces trois modèles utilisent le SoC CV25 ; les 87 autres références de la matrice reposent sur ARTPEC-8 ou ARTPEC-9, avec parfois deux processeurs. Ce constat n’est pas une règle de compatibilité à extrapoler : un modèle absent ne devient pas éligible parce qu’il utilise une puce similaire.
Même précaution pour les usages spécialisés. Une référence thermique, panoramique ou radar-vidéo présente dans la liste n’implique pas que toutes ses données et fonctions soient exposées dans Vision. Demandez une démonstration des fonctions nécessaires sur la référence exacte, puis consignez le résultat.
Licences : les points à verrouiller avant le devis
Chaque appareil Axis administré exige une licence Cisco spécifique, distincte des licences MV. Les modes co-term, Enterprise Agreement et Subscription sont prévus. Surtout, une organisation ne peut avoir qu’un seul niveau Axis à la fois ; un changement en cours de période nécessite de passer par l’équipe commerciale, sans bascule autonome garantie. Ces conditions figurent dans le guide d’intégration et sa FAQ.
Cela mérite une attention particulière si votre inventaire mélange les trois modèles limités à Essentials et des caméras destinées à Advantage. Ne construisez pas le budget sur l’hypothèse d’un choix libre caméra par caméra dans la même organisation.
Demandez une proposition qui précise le niveau, la durée, les quantités, le renouvellement et les conditions de sortie. Faites également confirmer la disponibilité régionale du service et les modalités de stockage. Nous ne donnons pas ici de tarif : les sources consultées ne permettent pas d’établir un prix public applicable à un projet français.
Plan d’action recommandé : prouver l’usage avant la généralisation
Chez BoucheCousue, nous recommandons de qualifier un petit échantillon représentatif avant de décider pour tout le parc. Voici une grille de recette, à adapter au besoin réel.
| Test proposé | Preuve attendue | 
|---|---|
| Caméra fixe, panoramique ou PTZ représentative du parc | Direct exploitable et réglages utiles accessibles | 
| Recherche puis export d’une séquence avec Advantage | Fichier réellement récupéré et relu, avec horodatage et droits vérifiés | 
| Coexistence du VMS avec Essentials | Usages habituels du logiciel existant toujours opérationnels | 
| Coupure Internet contrôlée puis rétablissement | Comportement mesuré de l’enregistrement, de la consultation et des alertes | 
| Maintenance et redémarrage | Interruption mesurée, reprise vérifiée et responsable identifié | 
| Départ d’un administrateur | Accès révoqué sans perte de maîtrise des organisations | 
| Sortie de la solution | Procédure documentée pour les équipements, les archives et les licences | 
N’attribuez pas un gain de temps au seul rapprochement des interfaces. Comparez le temps nécessaire pour résoudre une panne ou retrouver une séquence avant et après le pilote. C’est ce résultat qui permet de décider si la convergence simplifie vraiment le travail.
FAQ
Une caméra Axis compatible devient-elle une caméra MV ?
Non. Elle reste un équipement Axis intégré à l’environnement Meraki, avec son matériel, son firmware et ses capacités propres. L’intégration ne démontre pas une identité fonctionnelle avec une MV.
Peut-on conserver son enregistreur actuel ?
Essentials permet la coexistence avec le VMS local. Advantage impose Meraki comme VMS exclusif pour les appareils concernés. Il faut donc examiner les archives et les intégrations avant de retirer l’ancien système.
Faut-il remplacer les switches non-Cisco ?
Pas systématiquement. L’enrôlement par numéro de série et OAK couvre notamment les appareils raccordés à d’autres switches. Le réseau doit toutefois fournir l’alimentation et la connectivité nécessaires.
Un modèle compatible garantit-il toutes les fonctions vidéo ?
Non. La liste valide une compatibilité par niveau ; elle ne détaille pas chaque fonction spécialisée. La recette doit rester centrée sur les usages attendus et la référence exacte.
Notre lecture : une vraie intégration, avec un choix d’architecture
