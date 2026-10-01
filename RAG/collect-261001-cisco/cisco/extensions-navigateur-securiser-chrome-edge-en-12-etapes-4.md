---
id: collect-261001-cisco/cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes-4
title: "Dans chrome://extensions, pour chaque extension :"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes.md
source_anchor: ""
source_lines: [209, 258]
sha256: 843d86196121b408d2f87307ea31b9200f0afc24495c81a7db1c89b391c3a686
---

# Dans chrome://extensions, pour chaque extension :

**Deuxième piège : ignorer les mises à jour silencieuses.** Les extensions se mettent à jour automatiquement sans demander de confirmation explicite à chaque changement de permissions, contrairement aux applications mobiles. Un module qui demandait un accès limité à son installation peut réclamer un accès complet six mois plus tard sans qu’aucune alerte visible ne s’affiche.

**Troisième piège : laisser le mode développeur activé « au cas où ».** Ce réglage, souvent activé une fois pour tester une extension personnelle puis oublié, reste une porte ouverte au sideloading de code malveillant pendant des mois, voire des années.

**Quatrième piège : traiter toutes les extensions IA de la même façon.** Un module IA d’un éditeur reconnu et transparent sur sa politique de données n’a pas le même profil de risque qu’une extension IA obscure sortie de nulle part avec des avis achetés. Le nom « IA » dans la description ne dit rien du niveau de sécurité réel.

**Cinquième piège : ne pas revenir sur l’audit après l’avoir fait une fois.** La sécurité des extensions se dégrade avec le temps : nouvelles failles découvertes, changements d’éditeur, extensions oubliées qui continuent de tourner en arrière-plan. Un audit unique sans suivi perd son intérêt en quelques mois.

## Comparatif des fonctionnalités de sécurité par navigateur

| Fonctionnalité | Chrome | Edge | Firefox | 
|---|---|---|---|
| Restriction de permission “Au clic” | Oui | Oui | Partielle (via extensions tierces) | 
| Scan intégré des extensions | Enhanced Safe Browsing | Browser Essentials | Non natif | 
| Gestion centralisée entreprise | Chrome Enterprise Browser Cloud Management | Microsoft Edge Management Service | Firefox Enterprise Policies | 
| Profils multiples isolés | Oui | Oui | Oui (about:profiles) | 
| Blocage du mode développeur | Via stratégie de groupe | Via stratégie de groupe | Via about:config | 

## Conseils avancés pour aller plus loin

Pour les utilisateurs avancés et les équipes IT, quelques réglages supplémentaires renforcent nettement la posture de sécurité. Sur Chrome et Edge, la stratégie de groupe ExtensionInstallBlocklist permet de bloquer explicitement des identifiants d’extensions connus comme malveillants, en complément de la liste blanche. Sur Firefox, l’option de télémétrie des modules complémentaires (about:config, clé extensions.webextensions.restrictedDomains) permet d’empêcher certaines extensions de s’exécuter sur des domaines sensibles définis manuellement, comme les portails bancaires ou les intranets d’entreprise.

Autre pratique utile : conserver une machine virtuelle ou un conteneur de navigation jetable pour tester une nouvelle extension avant de l’installer sur le poste principal. Cette approche, courante dans les équipes de sécurité, permet d’observer le comportement réseau d’un module inconnu pendant 48 heures sans exposer les données réelles de l’utilisateur.

Enfin, pour les organisations soumises à la directive NIS2 ou traitant des données personnelles au sens du RGPD, l’audit des extensions de navigateur mérite d’être formalisé dans la politique de sécurité des systèmes d’information, avec une fréquence de révision alignée sur celle des autres inventaires logiciels du parc.

## Dépannage : 8 problèmes courants et leurs solutions

**1. Une extension refuse de se désinstaller.** Certaines extensions poussées par une stratégie de groupe d’entreprise ne peuvent pas être supprimées manuellement. Vérifiez dans chrome://policy si l’extension est listée comme forcée, auquel cas la demande de suppression doit passer par l’administrateur IT.

**2. Le mode “Au clic” casse le fonctionnement d’une extension.** Certains modules ont besoin d’un accès permanent pour fonctionner correctement (bloqueurs de publicité en temps réel, par exemple). Dans ce cas, préférez « Sur des sites spécifiques » plutôt que « Sur tous les sites », en listant uniquement les domaines réellement utilisés.

**3. ExtensionPedia ou un outil similaire ne trouve pas une extension récente.** Les bases de données de risque mettent parfois plusieurs semaines à référencer les nouvelles extensions. En l’absence de données, appliquez le principe de précaution : limitez les permissions au strict minimum en attendant.

**4. Le mode développeur se réactive tout seul après une mise à jour du navigateur.** Certaines mises à jour majeures réinitialisent des paramètres avancés. Reprenez l’habitude de vérifier ce réglage après chaque mise à jour importante de Chrome ou Edge.

**5. Impossible de créer un second profil de navigateur.** Sur les postes gérés en entreprise, la création de profils supplémentaires peut être bloquée par une stratégie de groupe. Contactez le support IT pour obtenir un profil isolé dédié aux usages sensibles.

**6. Une extension légitime déclenche une alerte de sécurité après une mise à jour.** Vérifiez d’abord si l’éditeur a changé sur la fiche du magasin d’extensions. Si l’éditeur reste identique et que l’alerte semble être un faux positif, consultez les avis récents d’autres utilisateurs avant de réinstaller.

**7. Le navigateur refuse d’installer la dernière mise à jour de sécurité.** Un espace disque insuffisant ou une stratégie de groupe obsolète bloquent parfois les mises à jour automatiques. Vérifiez l’espace disponible et, sur un poste d’entreprise, signalez le blocage à l’équipe IT sans attendre.

**8. Impossible de savoir si une extension a réellement changé d’éditeur.** Comparez le nom du développeur affiché sur la fiche actuelle du magasin avec une capture d’écran ou une note prise lors de l’installation initiale. En l’absence de trace, désinstallez puis réinstallez l’extension depuis zéro pour repartir sur une base de permissions à jour et consciente.

## Cas concret : comment une extension légitime bascule dans la malveillance

Le scénario mérite d’être détaillé, car il explique pourquoi un audit ponctuel ne suffit jamais. Tout commence généralement par un développeur indépendant qui publie une extension utile, gratuite, sans intention malveillante : un bloqueur de fenêtres pop-up, un outil de capture d’écran ou un gestionnaire d’onglets. L’extension gagne en popularité sur plusieurs mois, parfois plusieurs années, accumulant des centaines de milliers d’installations actives. Puis, faute de temps ou de motivation pour continuer à maintenir le projet, le développeur accepte une offre de rachat d’un tiers, souvent présenté comme une simple société de maintenance logicielle.

Le nouvel éditeur publie alors une mise à jour, en apparence anodine, qui ajoute une nouvelle permission au manifeste de l’extension : accès aux requêtes réseau, lecture du contenu de toutes les pages, ou injection de scripts supplémentaires. Cette mise à jour s’installe automatiquement chez tous les utilisateurs existants, sans qu’aucune fenêtre de confirmation explicite n’apparaisse dans la majorité des cas, puisque l’utilisateur a déjà accordé une confiance initiale à l’extension. C’est précisément ce mécanisme qui a été documenté fin août 2026 pour les 19 extensions rachetées puis armées : le changement d’éditeur est passé inaperçu pendant plusieurs semaines avant d’être signalé par des chercheurs en sécurité.

