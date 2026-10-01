---
id: collect-261001-ia-llm/ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026-2
title: ".github/copilot-instructions.md"
domain: ia-llm
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "agents", "mai"]
source: docs/RAG/collect-261001-ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026.md
source_anchor: ""
source_lines: [45, 102]
sha256: 621f1090d8f73bd339e1f747ee0ee971ce35e5b1d693a6ca39c360b1a8c745a8
---

# .github/copilot-instructions.md

Pour suivre ce tutoriel, le palier **Pro à 10 $/mois** suffit largement : GitHub avait fixé ce tarif dès avril 2025, aux côtés d’un palier Pro+ à 39 $/mois offrant 1 500 requêtes premium mensuelles, avant de confirmer en mars 2026 la disponibilité générale de l’Agent Mode sur VS Code et JetBrains avec 300 requêtes premium mensuelles pour Pro et des complétions illimitées, un quota depuis converti en crédits dans le nouveau système de facturation, mais qui donne une idée du volume disponible pour tester le projet complet présenté plus loin. Passez à Pro+ (39 $/mois, 1 500 requêtes premium selon la même grille de mars 2026, avec accès aux modèles de pointe) si vous voulez accéder aux modèles premium comme Opus pour des tâches de raisonnement plus complexes, ou au palier Max (100 $/mois, 200 $ de crédits IA inclus selon un état des lieux de Developers Digest publié en juin 2026) si votre volume de travail agentique dépasse une ou deux heures par jour. Les équipes qui gèrent plusieurs développeurs ont tout intérêt à regarder du côté de Business, facturé 19 $/utilisateur/mois avec 300 requêtes premium par utilisateur, un plafond déjà en place entre le 12 et le 19 mai 2025 lorsque GitHub avait lié ce même quota à l’Agent Mode, aux côtés d’un plafond de 1 000 requêtes premium mensuelles pour Enterprise, facturé 39 $/utilisateur/mois pour un accès complet. Ces deux paliers centralisent la facturation et ajoutent des contrôles d’administration au niveau de l’organisation. La grille tarifaire officielle reste la référence à consulter avant de vous engager, car ces paliers évoluent au rythme des annonces produit de GitHub.

Un point à anticiper si vous changez de palier en cours de mois : sur les paliers individuels, les crédits non consommés ne se reportent pas automatiquement d’un mois sur l’autre. Si une semaine chargée en refactoring ou en migration de dépendances s’annonce, mieux vaut passer au palier supérieur avant de démarrer plutôt que de vous retrouver bloqué en pleine session agentique, au milieu d’une série de modifications non committées.

## Étape 2 – Installer VS Code et les extensions Copilot

Téléchargez Visual Studio Code depuis le site officiel si ce n’est pas déjà fait, puis ouvrez le panneau Extensions avec `Ctrl+Shift+X` (ou `Cmd+Shift+X` sur Mac). Recherchez « GitHub Copilot » et installez l’extension proposée par GitHub : elle embarque automatiquement GitHub Copilot Chat, qui contient l’interface de l’Agent Mode. Redémarrez VS Code une fois l’installation terminée.

Vérifiez ensuite que votre version de l’éditeur est suffisamment récente. Ouvrez un terminal et lancez :

`code --version`
Vous devez obtenir une sortie proche de celle-ci, avec un numéro de version et un identifiant de build :

```
1.99.3
f97cb4a8d0d4c76f0f8f9e5c1a2b3d4e5f6a7b8c
x64
```
Si le numéro affiché est inférieur à 1.99, mettez à jour via `Aide > Rechercher les mises à jour` avant de continuer. Une version trop ancienne est la cause la plus fréquente d’un Agent Mode invisible dans le sélecteur de mode, un problème détaillé dans la section dépannage.

## Étape 3 – Connecter votre compte GitHub à VS Code

Ouvrez la palette de commandes avec `Ctrl+Shift+P` (`Cmd+Shift+P` sur Mac), tapez `GitHub Copilot: Sign In` et validez. VS Code ouvre votre navigateur par défaut pour finaliser l’authentification OAuth avec votre compte GitHub. Une fois l’autorisation confirmée dans le navigateur, revenez dans VS Code : une icône Copilot apparaît dans la barre d’état, en bas à droite de la fenêtre.

Cliquez sur cette icône pour vérifier que votre abonnement est bien reconnu. Si vous venez de souscrire un palier payant, patientez une ou deux minutes : la synchronisation entre le compte GitHub et l’extension n’est pas toujours instantanée. Un compte reconnu mais sans abonnement actif affiche l’Agent Mode grisé dans le sélecteur de mode de l’étape suivante.

Si votre compte dépend d’une organisation sous licence Business ou Enterprise, un administrateur doit au préalable vous avoir attribué un siège Copilot : la connexion individuelle ne suffit pas à elle seule à activer l’accès. Dans ce cas, l’icône Copilot de la barre d’état affiche un message d’attente plutôt qu’une erreur, ce qui prête parfois à confusion la première fois qu’on le rencontre.

## Étape 4 – Activer l’Agent Mode et lancer votre premier prompt

Ouvrez la vue Chat avec `Ctrl+Shift+I` (`Cmd+Shift+I` sur Mac), ou en cliquant sur l’icône Copilot dans la barre latérale. En haut du panneau de discussion, un menu déroulant liste les modes disponibles : Ask, Edit, Agent, Plan, Cloud et Copilot CLI. Sélectionnez **Agent**. Le 11 mars 2026, GitHub a fait passer en disponibilité générale les agents personnalisés, les sous-agents et le mode Plan pour JetBrains, tout en annonçant le retrait progressif de l’entrée « Edit » historique du sélecteur au profit d’une bascule directe vers l’Agent Mode ou vers un agent personnalisé, une évolution à surveiller si vous découvrez ce menu pour la première fois. Le tutoriel officiel de l’équipe VS Code détaille chacun de ces modes si vous voulez explorer les options Plan et Cloud au-delà de ce guide.

Ouvrez un dossier de projet existant, ou créez-en un vide pour tester. Dans le champ de saisie, écrivez une instruction claire et contextualisée plutôt qu’une phrase vague. Voici un exemple de prompt correct pour un premier essai :

```
Crée un script Python qui lit un fichier CSV nommé ventes.csv
(colonnes : produit, quantite, prix_unitaire), calcule le chiffre
d'affaires total par produit, et affiche un tableau trié par
chiffre d'affaires décroissant. Ajoute une gestion d'erreur si
le fichier est absent.
```
L’agent commence par annoncer un plan en quelques phrases : créer le script, définir la fonction de lecture du CSV, agréger les montants, gérer le cas du fichier manquant. Il crée ensuite le fichier, écrit le code, puis propose d’exécuter le script dans le terminal intégré pour vérifier qu’il fonctionne. Une fenêtre de progression affiche chaque action au fur et à mesure, avec un résumé du type :

```
▸ Lecture du dossier de travail
▸ Création de ventes.py
▸ Écriture de la fonction analyser_ventes()
▸ Proposition de commande : python ventes.py
▸ En attente de votre validation pour exécuter la commande
```
C’est précisément à ce moment que le contrôle humain entre en jeu, décrit dans les deux étapes suivantes.

## Étape 5 – Comprendre le cycle plan, édition, exécution, vérification

L’Agent Mode fonctionne selon une boucle en quatre temps qui se répète jusqu’à ce que la tâche soit terminée ou que vous l’interrompiez. D’abord, l’agent **planifie** : il décompose votre demande en une suite d’actions concrètes et affiche généralement ce plan avant d’agir. Ensuite vient l’**édition** : il crée ou modifie les fichiers nécessaires, en s’appuyant sur le contexte du projet qu’il a lu au préalable (structure des dossiers, dépendances, conventions de nommage existantes).

Troisième temps, l’**exécution** : l’agent propose des commandes terminal pour tester ce qu’il vient d’écrire, qu’il s’agisse de lancer un script, d’installer une dépendance ou de faire tourner une suite de tests. Enfin, la **vérification** : il lit la sortie de la commande, détecte les erreurs éventuelles, et repart en phase d’édition pour les corriger. Ce cycle peut se répéter plusieurs fois sur une tâche complexe, ce qui explique pourquoi une session d’Agent Mode dure parfois plusieurs minutes pour une demande qui semblait simple au départ.

