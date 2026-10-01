---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026-2
title: "AGENTS.md"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "claude", "gemini", "mai"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [52, 126]
sha256: 22d5fca3196cda2837319e5741204449850a60b0164ec6d8c638ea469b16d8f1
---

# AGENTS.md

Connectez-vous avec un email, un compte Google ou un compte GitHub. L’interface reprend les repères classiques d’un éditeur de code : explorateur de fichiers à gauche, zone d’édition au centre, terminal intégré en bas. Trois éléments spécifiques à Cursor AI méritent d’être repérés dès cette étape : le panneau de chat contextuel (raccourci `Ctrl+L` ou `Cmd+L`), l’édition en ligne rapide (`Ctrl+K` ou `Cmd+K`) et le bouton de bascule entre mode Chat et mode Agent, situé dans la barre de saisie du panneau latéral.

Prenez le temps de parcourir le petit didacticiel intégré qui s’affiche au premier lancement : il présente le fonctionnement du Tab prédictif, capable de deviner la prochaine modification à plusieurs endroits du fichier et pas seulement à la position du curseur. C’est une fonction moins spectaculaire que le mode Agent, mais c’est souvent elle qui fait gagner le plus de temps au quotidien, sur des tâches répétitives comme le renommage cohérent d’une variable dans tout un fichier.

## Étapes 4 et 5 : créer un compte et choisir sa formule tarifaire

### Étape 4 – Créer un compte Cursor et lier votre dépôt

Une fois connecté dans l’application, un compte est automatiquement créé sur le tableau de bord en ligne de Cursor. Vérifiez votre adresse email, puis liez votre compte GitHub depuis les réglages si vous comptez utiliser les Cloud Agents ou la revue de code automatisée Bugbot. Cette liaison est aussi ce qui permettra, plus tard, de déclencher un agent directement depuis un commentaire sur une pull request.

### Étape 5 – Comparer les formules Hobby, Pro, Pro+, Ultra et Teams

Le tarif de Cursor AI en 2026 repose sur un système de crédits par formule plutôt que sur un nombre fixe de requêtes. Voici le détail publié sur la page officielle de tarification, recoupé avec l’analyse du cabinet CloudZero publiée en mai 2026.

| Formule | Prix mensuel | Prix annuel (mensualisé) | Pour qui | 
|---|---|---|---|
| Hobby | Gratuit | — | Découverte, usage occasionnel, sans carte bancaire | 
| Pro | 20 $ | ~16 $ | Développeur individuel avec un usage régulier | 
| Pro+ | 60 $ | ~48 $ | Usage intensif, gros volumes de requêtes agent | 
| Ultra | 200 $ | ~160 $ | Power users, pool de crédits élevé (environ 400 $ équivalent) | 
| Teams | 40 $/utilisateur | ~32 $/utilisateur | Équipes : Bugbot, SSO, facturation centralisée | 
| Enterprise | Sur devis | Sur devis | Grands comptes : SCIM, logs d’audit, usage mutualisé | 

La formule Hobby suffit pour tester l’outil et comprendre le mode Agent, mais ses limites de requêtes se font sentir dès qu’on travaille sur un vrai projet au quotidien. Pour un développeur solo qui utilise Cursor AI plusieurs heures par jour, la formule Pro à 20 $/mois reste le point d’entrée le plus raisonnable. La formule Teams ajoute surtout des fonctions de gouvernance : contrôle d’accès centralisé, mode confidentialité d’équipe et connexion SSO via SAML ou OIDC, des critères qui pèsent lourd pour toute entreprise européenne soumise à des audits de sécurité.

L’engagement annuel fait mécaniquement baisser la facture mensuelle affichée, mais il fige aussi votre budget sur douze mois dans un secteur où les tarifs bougent encore régulièrement, comme le montre plus loin l’exemple de Windsurf qui a changé de modèle tarifaire en mars 2026. Pour une équipe qui découvre tout juste Cursor AI, mieux vaut démarrer en mensuel le temps de mesurer la consommation réelle de crédits sur six à huit semaines avant de s’engager sur un an.

## Étapes 6 et 7 : configurer le modèle IA et les réglages de l’éditeur

### Étape 6 – Choisir son modèle IA par défaut

Ouvrez Settings puis l’onglet Models. Cursor AI propose plusieurs familles de modèles récentes (Claude, GPT et Gemini selon les intégrations disponibles au moment de votre installation), ainsi qu’un mode de sélection automatique qui choisit le modèle le plus adapté à la tâche. Vous pouvez désactiver les modèles que vous n’utilisez jamais pour simplifier le menu, et surtout ajouter votre propre clé API si vous préférez payer directement le fournisseur du modèle plutôt que de consommer les crédits inclus dans votre formule. Cette option de clé personnelle est particulièrement utile pour les équipes qui ont déjà un contrat entreprise avec un fournisseur de modèle et veulent éviter de payer deux fois.

Le choix du modèle n’est pas qu’une question de préférence personnelle, c’est aussi un arbitrage coût-qualité. Un modèle plus capable consomme davantage de crédits par requête, ce qui a du sens pour un refactor d’architecture complexe mais devient du gaspillage pour générer un test unitaire simple ou corriger une faute de frappe dans un message d’erreur. Le mode de sélection automatique tente de faire cet arbitrage à votre place, mais une équipe expérimentée gagne souvent à fixer manuellement un modèle économique pour les tâches routinières et à réserver le modèle premium aux vraies décisions d’architecture.

### Étape 7 – Personnaliser Tab, raccourcis et thème

Dans Settings > Editor, ajustez le comportement de l’autocomplétion Tab : délai avant suggestion, acceptation partielle mot par mot avec `Ctrl+→`, et portée du contexte pris en compte (fichier courant seulement, ou fichiers ouverts dans l’onglet). Comme Cursor AI reprend le moteur de VS Code, les raccourcis clavier personnalisés se gèrent de la même façon, via un fichier de configuration. Voici un exemple de personnalisation courante, qui réassigne l’acceptation d’une suggestion Agent à une touche dédiée :

```
[
  {
    "key": "ctrl+enter",
    "command": "composer.submitComposerPrompt",
    "when": "composerFocus"
  },
  {
    "key": "ctrl+shift+l",
    "command": "workbench.action.chat.open"
  }
]
```
Ce fichier se modifie depuis la palette de commandes, en cherchant “Preferences: Open Keyboard Shortcuts (JSON)”. Les noms de commandes exacts peuvent varier d’une version à l’autre : le plus fiable reste de chercher l’action voulue dans l’interface graphique, puis de récupérer son identifiant technique avant de l’éditer en JSON.

## Étapes 8 et 9 : mettre en place vos règles de projet avec .cursor/rules

C’est l’étape la plus déterminante de tout ce tutoriel. Les grands modèles de langage ne conservent aucune mémoire d’une complétion à l’autre : sans mécanisme de contexte persistant, vous répéteriez les mêmes instructions d’architecture à chaque nouvelle conversation. Les règles de projet résolvent ce problème.

### Étape 8 – Créer le dossier .cursor/rules et votre premier fichier .mdc

À la racine de votre projet, créez un dossier `.cursor/rules` puis un premier fichier à l’intérieur. Le point qui piège le plus de développeurs débutants sur Cursor AI, comme le précise la documentation officielle sur les règles : l’extension doit obligatoirement être `.mdc`, pas `.md`. Selon cette même documentation, **« un simple fichier .md placé dans .cursor/rules est ignoré par le système de règles car il ne contient pas l’en-tête nécessaire pour spécifier description, globs et alwaysApply »**. Voici un exemple de règle complète pour un projet React :

```
---
description: "Conventions de composants React"
globs: src/components/**/*.tsx
alwaysApply: false
---
Utilise des composants fonctionnels avec des hooks, jamais de classes.
Chaque composant reçoit ses props typées via une interface TypeScript dédiée.
Les appels réseau passent systématiquement par le client défini dans src/lib/api.ts.
N'ajoute jamais de style inline : utilise les classes Tailwind existantes.
```
Le même principe s’applique côté backend, avec un fichier ciblé sur un tout autre dossier et un tout autre langage :

