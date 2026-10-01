---
id: collect-261001-cisco/cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes-2
title: "Dans chrome://extensions, pour chaque extension :"
domain: cisco
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["exploit", "valuation"]
source: docs/RAG/collect-261001-cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes.md
source_anchor: ""
source_lines: [56, 128]
sha256: b253b18119856842aed97f145bee2826b3ffab1de8dc5acef7fcf51eb7441590
---

# Dans chrome://extensions, pour chaque extension :

Pour objectiver cette évaluation sans tout faire manuellement, des outils dédiés existent. ExtensionPedia permet de rechercher n’importe quel module et d’obtenir une note de risque basée sur l’analyse des permissions, la réputation de l’éditeur et les comportements suspects observés. Pour un usage professionnel à plus grande échelle, CRXcavator et LayerX analysent en masse les extensions installées sur un parc de postes et signalent celles qui présentent un profil de risque élevé.

| Niveau de risque | Signaux typiques | Action recommandée | 
|---|---|---|
| Faible | Éditeur vérifié, mises à jour récentes, permissions limitées au strict nécessaire | Conserver, revérifier tous les 6 mois | 
| Modéré | Éditeur identifiable mais permissions larges par rapport à la fonction | Restreindre les permissions de site, surveiller | 
| Élevé | Éditeur non vérifié, mise à jour ancienne, accès à toutes les données | Désinstaller ou remplacer par une alternative | 
| Critique | Changement d’éditeur récent, comportement réseau inhabituel signalé | Désinstaller immédiatement, changer les mots de passe utilisés dans le navigateur | 
| Extension IA générique | Accès complet au contenu des pages pour fonctionner | N’installer que depuis des éditeurs reconnus, isoler dans un profil dédié | 

## Étape 3 : restreindre les permissions au lieu de tout supprimer

Supprimer une extension utile n’est pas toujours la meilleure option. Chrome, Edge et Firefox permettent de restreindre finement l’accès d’un module à certains sites plutôt qu’à l’ensemble du web. Dans Chrome ou Edge, ouvrez la page des extensions, cliquez sur « Détails » pour le module concerné, puis dans la section « Autorisations de site », remplacez « Sur tous les sites » par « Sur des sites spécifiques » ou « Au clic ». Cette dernière option est la plus restrictive : l’extension ne s’active que lorsque vous cliquez explicitement sur son icône.

```
# Dans chrome://extensions, pour chaque extension :
1. Cliquer sur "Détails"
2. Section "Autorisations de site"
3. Choisir :
   - "Au clic"                 -> le plus restrictif, recommandé par défaut
   - "Sur des sites spécifiques" -> pour les extensions utilisées sur un périmètre connu
   - "Sur tous les sites"      -> à réserver aux extensions de confiance absolue
                                  (gestionnaire de mots de passe vérifié, par exemple)
```
Cette approche « au clic » change concrètement l’exposition d’un poste de travail. Une extension de traduction, par exemple, n’a besoin d’accéder au contenu d’une page que lorsque vous lui demandez de traduire quelque chose : il n’y a aucune raison de lui laisser un accès permanent à votre messagerie professionnelle ou à votre espace bancaire en ligne.

## Étape 4 : désactiver définitivement le mode développeur

Le mode développeur, présent sur Chrome et Edge, permet d’installer des extensions non signées en dehors du Chrome Web Store ou du Microsoft Store, une pratique appelée sideloading. C’est une fonctionnalité légitime pour les développeurs qui testent leurs propres extensions, mais c’est aussi la méthode préférée des malwares pour s’injecter directement dans le navigateur sans passer par les contrôles du magasin officiel.

Si vous n’êtes pas développeur d’extensions, désactivez ce mode. Sur la page chrome://extensions ou edge://extensions, repérez le bouton bascule « Mode développeur » en haut à droite de l’écran et vérifiez qu’il est désactivé. Si une extension inconnue apparaît dans votre liste sans que vous vous souveniez de l’avoir installée depuis le magasin officiel, c’est un signal d’alarme immédiat : désinstallez-la et changez les mots de passe utilisés récemment dans ce navigateur.

## Étape 5 : activer les outils de scan intégrés au navigateur

Les navigateurs modernes intègrent désormais leurs propres outils de détection. Sur Windows 11, Microsoft Edge dispose d’une fonctionnalité appelée Browser Essentials, capable de scanner les extensions installées pour repérer celles qui ont un impact négatif documenté sur la sécurité ou les performances. Pour l’activer, ouvrez le panneau latéral d’Edge, sélectionnez Browser Essentials, puis lancez un scan de performance et de sécurité.

Sur Chrome, l’équivalent se trouve dans les paramètres de sécurité, sous « Confidentialité et sécurité » puis « Sécurité », où l’option de navigation sécurisée avancée (Enhanced Safe Browsing) permet à Google d’analyser en temps réel les pages visitées et, indirectement, de mieux détecter les comportements suspects liés aux extensions. Activez cette option si ce n’est pas déjà fait.

```
Chrome : Paramètres > Confidentialité et sécurité > Sécurité
         > Sécurité renforcée (Enhanced Safe Browsing) > Activer
Edge   : Panneau latéral > Browser Essentials > Lancer un scan
```
## Étape 6 : mettre à jour le navigateur immédiatement

Aucun réglage d’extension ne compense un navigateur non mis à jour. Le CERT-FR a confirmé début septembre 2026 que la faille CVE-2026-85046 affectant Google Chrome est activement exploitée dans la nature. Malwarebytes recommandait fin août 2026 de mettre à jour Chrome avant même de reprendre sa navigation, tant les correctifs récents couvraient des failles critiques — un conseil d’autant plus pertinent que Chrome 152, stable depuis le 25 août 2026, a déjà été remplacé par Chrome 153 le 8 septembre 2026, lequel avait introduit dès août 2026 la nouvelle API browser.publicSuffix selon Google Chrome Developers, signe d’un rythme de publication qui ne laisse plus de place à un navigateur laissé à l’abandon. Voici la procédure de vérification pour chaque navigateur :

```
Chrome  : Menu ⋮ > Aide > À propos de Google Chrome
          -> la mise à jour se télécharge automatiquement
          -> cliquer sur "Relancer" pour l'appliquer
Edge    : Menu ... > Aide et commentaires > À propos de Microsoft Edge
          -> vérifie et installe automatiquement les mises à jour
Firefox : Menu ≡ > Aide > À propos de Firefox
          -> "Rechercher des mises à jour" puis "Redémarrer pour mettre à jour"
```
Configurez également la mise à jour automatique quand elle est disponible plutôt que de compter sur une vérification manuelle régulière. La fenêtre entre la divulgation d’une faille et son exploitation active s’est nettement réduite ces derniers mois, et attendre plusieurs jours avant d’appliquer un correctif expose inutilement le poste de travail.

## Étape 7 : cloisonner les usages sensibles dans un profil dédié

Les trois navigateurs permettent de créer plusieurs profils indépendants, chacun avec sa propre liste d’extensions, ses propres cookies et son propre historique. C’est la parade la plus efficace contre le risque résiduel d’une extension compromise : créez un profil « Banque et administratif » totalement dépourvu d’extensions, réservé exclusivement aux opérations sensibles (accès bancaire, déclarations fiscales, portail employeur), et gardez vos extensions habituelles dans un profil « Navigation courante » séparé.

```
Chrome / Edge :
1. Cliquer sur l'icône de profil en haut à droite
2. "Ajouter un profil" ou "Ajouter un compte"
3. Choisir "Continuer sans compte" pour un profil local isolé
4. Ne jamais installer d'extension dans ce profil dédié
Firefox :
about:profiles > "Créer un nouveau profil"
```
Cette séparation limite mécaniquement l’impact d’une extension malveillante : même si un module compromis parvient à lire les cookies de session du profil de navigation courante, il n’aura aucun accès à la session ouverte dans le profil bancaire isolé, puisque les deux environnements ne partagent ni cookies ni stockage local.

