---
id: collect-261001-cisco/cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes-1
title: "Dans chrome://extensions, pour chaque extension :"
domain: cisco
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["exploit", "gemini"]
source: docs/RAG/collect-261001-cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes.md
source_anchor: ""
source_lines: [1, 55]
sha256: 08b145715efdf1af43840d8ebc6a9103f6ef32149a20344620ee713198b42348
---

# Dans chrome://extensions, pour chaque extension :

Une extension de navigateur qui change de propriétaire du jour au lendemain, un module d’IA qui lit vos cookies de session, un correctif Chrome publié en urgence parce qu’une faille est déjà exploitée : la liste des raisons de faire le ménage dans ses extensions n’a jamais été aussi longue en 2026. Selon une étude relayée en avril 2026, les extensions dopées à l’IA présentent 60 % de risques de vulnérabilité supplémentaires par rapport aux extensions classiques, et sont trois fois plus susceptibles d’accéder à des données sensibles comme les cookies de session ou les identifiants. Ce tutoriel vous montre, étape par étape, comment auditer, restreindre et sécuriser les extensions installées sur Chrome, Edge et Firefox, avec les bons réglages, les bons outils et les pièges à éviter.

L’actualité récente donne du poids à la démarche. Le CERT-FR a confirmé le 1er septembre 2026 que la vulnérabilité CVE-2026-85046 dans Google Chrome est activement exploitée. Quelques semaines plus tôt, Chrome 152, devenu la version stable de référence le 25 août 2026 selon Supercharge Browser, corrigeait 26 failles de sécurité, dont deux vulnérabilités critiques de type use-after-free (CVE-2026-84353 et CVE-2026-84352) permettant une évasion du bac à sable ; il a depuis cédé la place à Chrome 153, lancé le 8 septembre 2026 sur un nouveau cycle de publication accéléré à deux semaines. Et un rapport publié fin août 2026 signale que 19 extensions de navigateur ont été rachetées puis « armées » avec du code malveillant, touchant potentiellement des millions d’installations actives. Autant dire que remettre à plat sa liste d’extensions n’est plus une option pour qui navigue, télétravaille ou gère des comptes sensibles depuis son navigateur.

## Pourquoi les extensions de navigateur sont devenues une cible prioritaire

Une extension de navigateur tourne avec des privilèges bien plus larges qu’une application mobile classique. Elle peut lire le contenu de chaque page visitée, intercepter les requêtes réseau, injecter du JavaScript et, selon les autorisations accordées, accéder aux cookies de session qui maintiennent vos connexions à votre banque, votre messagerie ou votre espace professionnel. Ce niveau d’accès explique pourquoi les extensions sont devenues un vecteur d’attaque à part entière, documenté noir sur blanc par le CERT-FR et le CERT Santé dans leurs bulletins hebdomadaires.

Le problème s’est aggravé avec l’arrivée des extensions dites « IA », qui promettent de résumer une page, réécrire un e-mail ou automatiser une tâche de navigation. Pour fonctionner, ces modules demandent souvent un accès complet au contenu des pages, ce qui élargit mécaniquement leur surface d’attaque. Une analyse publiée en mars 2026 a mis en évidence une vulnérabilité (CVE-2026-0628) dans le panneau Gemini Live de Chrome, corrigée dans la version 143 de janvier 2026 : des extensions malveillantes pouvaient y injecter du code JavaScript pour obtenir une escalade de privilèges dans le navigateur. La même étude a signalé des failles d’injection de prompt dans le navigateur agentique Comet, capables de détourner des gestionnaires de mots de passe et d’exfiltrer des secrets.

Autre phénomène documenté : le rachat silencieux d’extensions populaires. Un développeur indépendant vend son module à un tiers qui, quelques semaines plus tard, pousse une mise à jour intégrant du code de collecte de données ou de détournement de trafic publicitaire. Un bulletin de fin août 2026 recense 19 extensions concernées par ce scénario, marquées « ACTION_REQUISE » par les outils de suivi spécialisés. Comme les mises à jour se font automatiquement et sans confirmation explicite de l’utilisateur, ce type de dérive passe totalement inaperçu tant qu’on ne fait pas d’audit régulier.

## Prérequis avant de commencer

Ce tutoriel fonctionne sur les trois navigateurs les plus utilisés en France et en Europe. Vérifiez que vous disposez des versions suivantes ou plus récentes avant de commencer, car les pages de gestion des extensions et les fonctionnalités de sécurité décrites ci-dessous évoluent d’une version à l’autre — d’autant que Google est passé, avec Chrome 153 lancé le 8 septembre 2026, à un cycle de publication accéléré à deux semaines, ce qui resserre encore l’écart entre deux vérifications :

- Google Chrome 153 ou version ultérieure, stable depuis le 8 septembre 2026 (menu ⋮ puis Aide > À propos de Google Chrome pour vérifier)
- Microsoft Edge basé sur Chromium, version 140 ou ultérieure, avec Browser Essentials activé
- Mozilla Firefox 145 ou ultérieure
- Un compte administrateur sur la machine si vous gérez un parc d’entreprise via une stratégie de groupe
- 15 à 20 minutes pour l’audit initial, un peu plus si vous gérez plusieurs postes
- Optionnel : un compte Chrome Enterprise Browser Cloud Management ou Microsoft Edge Management Service si vous administrez plusieurs postes

Aucune extension supplémentaire n’est strictement nécessaire pour suivre ce guide, mais nous verrons plus loin des outils optionnels comme ExtensionPedia ou CRXcavator pour évaluer le niveau de risque d’un module avant de l’installer.

## Étape 1 : lister toutes les extensions installées

La première étape consiste à dresser un inventaire complet, navigateur par navigateur. C’est fastidieux mais indispensable : on ne peut pas sécuriser ce qu’on ne connaît pas. Ouvrez chaque page dédiée aux extensions dans votre navigateur :

```
Chrome  : chrome://extensions
Edge    : edge://extensions
Firefox : about:addons (rubrique Extensions)
```
Pour chaque module listé, notez trois informations : le nom de l’éditeur, la date de dernière mise à jour et les autorisations demandées. Sur Chrome et Edge, cliquez sur « Détails » pour chaque extension afin d’afficher ces éléments. Sur Firefox, un clic sur le nom du module ouvre sa fiche complète avec les permissions accordées.

Un exemple de sortie typique pour un poste de travail avec un usage courant :

```
Extension : Gestionnaire de mots de passe X
Éditeur   : Vérifié
Dernière MAJ : il y a 6 jours
Permissions : Lire et modifier toutes vos données sur tous les sites
Extension : Convertisseur PDF Rapide
Éditeur   : Non vérifié
Dernière MAJ : il y a 14 mois
Permissions : Lire et modifier toutes vos données sur tous les sites,
              Gérer vos téléchargements
```
Dans cet exemple, la seconde extension cumule deux signaux d’alerte : un éditeur non vérifié et une absence de mise à jour depuis plus d’un an, alors qu’elle réclame un accès complet à toutes les pages visitées. C’est exactement le profil que les attaquants recherchent pour racheter un module à l’abandon et y injecter du code malveillant.

## Étape 2 : évaluer le niveau de risque de chaque extension

Une fois l’inventaire terminé, il faut trier. Les critères à appliquer systématiquement sont les suivants : l’éditeur est-il identifiable et vérifié, l’extension est-elle mise à jour régulièrement, les autorisations demandées sont-elles proportionnées à la fonction annoncée, et le nombre d’utilisateurs actifs est-il cohérent avec l’ancienneté du module. Une extension à « fonction banale » (convertisseur, capture d’écran, thème visuel) qui réclame un accès à toutes les données de tous les sites doit systématiquement être questionnée : ce type de profil est régulièrement utilisé pour voler des portefeuilles de cryptomonnaies ou détourner des sessions de connexion.

