---
id: collect-261001-cisco/cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes-3
title: "Dans chrome://extensions, pour chaque extension :"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-09-06", "2026-12-06"]
keywords: ["valuation"]
source: docs/RAG/collect-261001-cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes.md
source_anchor: ""
source_lines: [129, 208]
sha256: ee2643c2beb23a7231e66c5180d09d53592f2dc3fe6eaa0c20b91feb6d4ee9fd
---

# Dans chrome://extensions, pour chaque extension :

## Étape 8 : mettre en place une liste blanche pour un parc de postes

Pour une entreprise ou une administration, l’audit manuel poste par poste ne tient pas dans la durée. La recommandation qui s’impose depuis 2026 consiste à basculer d’une logique de liste noire (bloquer les extensions connues comme dangereuses) vers une logique de liste blanche (n’autoriser que les extensions explicitement validées). Chrome Enterprise Browser Cloud Management et Microsoft Edge Management Service permettent tous les deux de définir centralement quelles extensions sont autorisées sur l’ensemble d’un parc, et de bloquer automatiquement toute installation en dehors de cette liste.

Concrètement, la démarche recommandée est la suivante : inventorier toutes les extensions actuellement installées sur les postes qui accèdent à des données personnelles ou professionnelles sensibles, valider une liste restreinte avec l’équipe sécurité, puis pousser une stratégie de groupe qui bloque tout le reste. Chaque changement d’éditeur d’une extension déjà présente sur la liste blanche doit être traité comme un événement à réévaluer, et non comme une simple mise à jour de routine.

## Étape 9 : surveiller le comportement réseau des extensions

Un module compromis ne se contente généralement pas de lire des données localement : il les envoie quelque part. La surveillance des connexions réseau sortantes inhabituelles, via une solution comme Microsoft Defender sur les postes de travail, permet de repérer une extension qui communique avec un serveur externe alors que sa fonction annoncée ne le justifie pas. Sur un poste personnel, les pare-feu applicatifs intégrés à la plupart des suites de sécurité offrent une granularité suffisante pour identifier ce type de comportement.

Pour les extensions qui demandent la permission technique « declarativeNetRequest », un audit renforcé est recommandé : cette autorisation donne à l’extension la capacité de modifier ou de bloquer des requêtes réseau, ce qui peut être détourné pour rediriger du trafic ou injecter du contenu publicitaire malveillant sans que l’utilisateur ne s’en aperçoive.

## Étape 10 : vérifier régulièrement les extensions dopées à l’IA

Les extensions présentées comme « assistants IA », capables de résumer, traduire ou automatiser une navigation, méritent un traitement à part. Leur promesse fonctionnelle nécessite presque toujours un accès étendu au contenu des pages, ce qui rend l’évaluation du risque plus complexe qu’avec une extension classique. La règle pratique : n’installer ce type de module que depuis des éditeurs identifiables (grands éditeurs de navigateurs, entreprises connues du secteur de l’IA), vérifier la politique de traitement des données affichée sur la fiche du magasin d’extensions, et réévaluer périodiquement si l’éditeur reste le même.

Le cas du navigateur agentique Comet, documenté début 2026, illustre le risque spécifique de cette catégorie : une vulnérabilité d’injection de prompt pouvait détourner un gestionnaire de mots de passe intégré et exfiltrer des identifiants sans action explicite de l’utilisateur. La mise à jour vers les versions publiées après février 2026 corrige ce point précis, mais le principe reste valable pour tout navigateur ou extension combinant capacités d’IA générative et accès large au contenu des pages.

## Étape 11 : nettoyer les extensions abandonnées

Une extension qui n’a pas reçu de mise à jour depuis plus d’un an, même si elle n’a jamais montré de comportement suspect, représente un risque latent. Un éditeur qui abandonne un projet devient une cible facile pour un rachat discret, précisément le scénario documenté avec les 19 extensions « armées » signalées fin août 2026. La règle simple à appliquer : toute extension non mise à jour depuis 12 mois et non indispensable doit être désinstallée, quitte à la réinstaller plus tard si le besoin réapparaît et si l’extension a repris un développement actif.

```
# Vérification manuelle, à répéter tous les 3 mois
Pour chaque extension installée :
  Si (dernière mise à jour > 12 mois) ET (usage non quotidien) :
      -> Désinstaller
  Si (changement d'éditeur détecté) :
      -> Désinstaller immédiatement + changer les mots de passe du navigateur
  Si (permissions > besoin réel de la fonction) :
      -> Restreindre à "Au clic"
```
## Étape 12 : documenter et planifier l’audit suivant

La sécurisation des extensions n’est pas une opération ponctuelle mais un cycle. Notez la date de votre audit, la liste des extensions conservées et leur niveau de risque évalué, puis fixez une date de rappel à trois mois. Pour un usage personnel, un simple rappel dans l’agenda suffit. Pour une organisation, ce cycle doit être intégré au processus de gestion des vulnérabilités déjà en place pour les autres composants du système d’information, au même titre que les correctifs serveurs ou les mises à jour d’antivirus.

## Le projet complet : script d’audit et checklist réutilisable

Voici un récapitulatif complet, sous forme de checklist, à suivre dans l’ordre pour un audit initial complet des extensions sur un poste Windows, macOS ou Linux :

```
CHECKLIST D'AUDIT DES EXTENSIONS DE NAVIGATEUR - 2026
[ ] 1. Ouvrir chrome://extensions, edge://extensions et about:addons
[ ] 2. Lister éditeur, date de MAJ et permissions de chaque extension
[ ] 3. Vérifier chaque extension sur ExtensionPedia (ou équivalent)
[ ] 4. Désinstaller les extensions non utilisées depuis 12+ mois
[ ] 5. Désinstaller les extensions à éditeur non vérifié + permissions larges
[ ] 6. Passer les extensions restantes en "Au clic" quand c'est possible
[ ] 7. Désactiver le mode développeur sur Chrome et Edge
[ ] 8. Activer Enhanced Safe Browsing (Chrome) et Browser Essentials (Edge)
[ ] 9. Mettre à jour le navigateur vers la dernière version stable
[ ] 10. Créer un profil dédié sans extension pour les usages sensibles
[ ] 11. Pour une entreprise : basculer sur une liste blanche centralisée
[ ] 12. Noter la date de l'audit et planifier le suivant à J+90
```
Ce projet, une fois mis en place, prend moins de 20 minutes à exécuter lors des audits suivants, contre 45 à 60 minutes pour le premier passage sur un poste chargé en extensions accumulées au fil des années.

## Résultats attendus après un audit complet

Sur un poste de travail type, comportant entre 15 et 25 extensions installées au fil du temps, un audit complet aboutit généralement à la suppression de 30 à 50 % des modules, soit parce qu’ils ne sont plus utilisés, soit parce que leurs permissions sont disproportionnées par rapport à leur fonction réelle. Un exemple de rapport final après audit :

```
Résumé de l'audit - Poste "PC-Bureau-01"
Date : 06/09/2026
Extensions initiales      : 21
Extensions supprimées     : 9
  - 6 non utilisées depuis 12+ mois
  - 2 éditeur non vérifié + permissions larges
  - 1 changement d'éditeur détecté (action critique)
Extensions conservées     : 12
  - 8 passées en mode "Au clic"
  - 4 conservées en accès complet (gestionnaires vérifiés)
Mode développeur          : Désactivé
Navigateur                : À jour (dernière version stable)
Prochain audit programmé  : 06/12/2026
```
## 5 pièges fréquents à éviter

**Premier piège : confondre popularité et sécurité.** Une extension téléchargée par des millions d’utilisateurs n’est pas automatiquement sûre. Le nombre d’installations reflète la notoriété au moment du lancement, pas le comportement actuel du code après un éventuel rachat de l’éditeur.

