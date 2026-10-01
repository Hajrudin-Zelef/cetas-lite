---
id: collect-261001-general-networking/general-networking/google-antigravity-tutoriel-en-12-etapes-2026-5
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "chatgpt", "claude", "compute", "copilot", "gemini", "mai", "opus 4"]
source: docs/RAG/collect-261001-general-networking/google-antigravity-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [327, 389]
sha256: ff523f2ee2db1736f01646c5f6d25aa06be50878937819bd534df5839f8edabb
---

# macOS (Homebrew)

| Problème | Cause probable | Solution | 
|---|---|---|
| L’application ne démarre pas sous Linux | Dépendances graphiques manquantes | Installer `libnss3` ,`libgbm1` ,`libasound2` via apt | 
| Impossible de se connecter avec Google | Compte Workspace bloqué par l’admin | Utiliser un compte Google personnel | 
| « Rate limit reached » | Quota gratuit épuisé | Attendre le rafraîchissement (~5 h) ou passer sur Flash / clé Anthropic | 
| L’agent perd le fil en cours de tâche | Contexte trop long, erreur mémoire | Découper la mission, revenir à un checkpoint | 
| Le sous-agent navigateur n’ouvre rien | Extension Chrome non installée | Installer l’extension et autoriser la connexion IDE ↔ navigateur | 
| SmartScreen bloque l’installeur (Windows) | Binaire d’aperçu non signé | « Informations complémentaires » > « Exécuter quand même » | 
| Les modèles Claude sont grisés | Clé API Anthropic absente | Ajouter la clé dans Settings > Model providers | 
| Les extensions VS Code manquent | Import non effectué au 1er lancement | Réimporter via la palette (Import VS Code settings) | 

Deux cas méritent un mot de plus. Pour l’erreur de **quota**, gardez en tête que les limites de l’aperçu gratuit ont été revues à la baisse plusieurs fois depuis novembre 2025 ; si vous travaillez intensivement, router Claude via votre propre clé API ou souscrire à un plan payant devient vite nécessaire. Pour la **perte de contexte**, la meilleure parade reste de découper : une mission bien bornée réussit là où une mission fourre-tout dérape. Enfin, si Antigravity se fige (UI freeze) sur une opération longue, redémarrez l’application – vos fichiers et checkpoints sont conservés sur le disque.

## Astuces avancées pour aller plus loin

Une fois à l’aise avec le flux de base, ces techniques vous feront gagner en efficacité et en sûreté.

- **Agents en parallèle.** Dans le Manager, lancez une mission « backend » et une mission « tests et documentation » simultanément, sur des espaces de travail distincts. Vous divisez le temps total sans mélanger les contextes.
- **Modèle par agent.** Assignez Gemini 3 Pro à l’agent d’architecture et Gemini 3 Flash aux agents d’exécution. Vous optimisez le rapport qualité/quota mission par mission.
- **Base de connaissances vivante.** Encouragez l’agent à consigner les décisions importantes (choix de bibliothèque, conventions d’API) dans la base de connaissances. Les missions suivantes en héritent, ce qui améliore la cohérence.
- **Durcir la sécurité.** Traitez tout contenu web comme non fiable. Ne connectez pas de secrets de production à un agent qui contrôle le navigateur. Révisez les diffs avant d’exécuter des commandes système sensibles, et privilégiez un environnement isolé (conteneur, machine virtuelle) pour les projets sensibles.
- **Gérer les crédits.** Depuis mars 2026, Google propose des crédits IA (25 $ pour 2 500 crédits, soit 0,01 $ le crédit) en complément des quotas. Surveillez votre consommation dans les réglages pour éviter les mauvaises surprises en fin de projet.

Enfin, gardez `AGENTS.md` comme document vivant. À mesure que le projet grossit, ajoutez-y les décisions structurantes (choix de base de données réel, conventions de nommage des routes, stratégie de tests). Un `AGENTS.md` bien tenu est le meilleur multiplicateur de qualité pour un IDE agentique – bien plus qu’un prompt parfait pour une mission isolée. Pour approfondir l’usage d’assistants en ligne de commande complémentaires, voyez notre tutoriel Gemini CLI.

## Tarifs, crédits et limites en 2026

Google Antigravity reste **gratuit pendant l’aperçu public** – 0 $ pour les particuliers dès le lancement de novembre 2025, selon le Google Developers Blog –, mais l’histoire de ses quotas est mouvementée. Le blog officiel d’antigravity.google précisait déjà, en novembre 2025, que les limites de l’aperçu pour Gemini 3/3.1 Pro se rafraîchissaient toutes les 5 heures ; depuis, ces limites ont été réduites à plusieurs reprises : d’après les rapports d’utilisateurs, le plafond initial d’environ 250 requêtes par jour a été ramené à une vingtaine dès décembre 2025, avant un passage à des limites hebdomadaires. En mars 2026, Google a introduit un système de crédits IA. À l’occasion de Google I/O 2026 (19-20 mai 2026), l’entreprise a lancé l’abonnement **Google AI Ultra** à 100 $/mois, pensé pour donner un accès prioritaire à Antigravity avec des limites d’usage cinq fois supérieures à celles du plan Pro, selon Google ; puis, en août 2026, un palier Ultra encore plus élevé est passé de 250 $ à 200 $/mois tout en offrant des limites vingt fois supérieures au plan Pro, d’après Emergent.

| Offre | Prix (indicatif) | Accès aux modèles | Limites | 
|---|---|---|---|
| Gratuit (aperçu) | 0 $ | Gamme complète | Quota compute, rafraîchi ~toutes les 5 h | 
| Google AI Pro | 20 $/mois | Gamme complète | Limites élevées, priorité | 
| Google AI Ultra | 100 $/mois | Gamme complète | 5× les limites du plan Pro | 
| Google AI Ultra Max | 200 $/mois | Gamme complète | 20× les limites Pro (baissé de 250 $ à 200 $ en août 2026, selon Emergent) | 
| Crédits IA | 25 $ / 2 500 crédits | Complément | 0,01 $ par crédit | 

Pour les utilisateurs européens, ces tarifs en dollars sont facturés à leur équivalent en euros par Google. Notre conseil pour l’aperçu gratuit : réservez-le à l’apprentissage et aux petits projets. Dès que vous travaillez quotidiennement avec des agents, soit vous routez Claude via votre propre clé Anthropic, soit vous passez sur un plan payant. Les allers-retours de quotas de fin 2025 ont montré que compter uniquement sur la gratuité pour un usage professionnel est risqué. Pour situer ces coûts face à d’autres assistants, notre comparatif Copilot vs ChatGPT donne des points de repère utiles.

## Foire aux questions

### Google Antigravity est-il vraiment gratuit ?

Oui, pendant l’aperçu public, avec un compte Google personnel et sans carte bancaire. Vous accédez à la gamme complète de modèles, mais dans les limites d’un quota qui se rafraîchit environ toutes les 5 heures. Ce quota a été réduit plusieurs fois depuis novembre 2025 ; pour un usage intensif, un plan payant ou une clé API Anthropic devient nécessaire.

### Quels modèles IA Antigravity prend-il en charge ?

D’après la documentation officielle et la fiche Wikipédia à jour en 2026 : Gemini 3 Pro (par défaut), Gemini 3 Flash, Claude Sonnet 4.6, Claude Opus 4.6 et GPT-OSS-120B. Vous pouvez assigner un modèle différent à chaque agent d’une même mission, et router Claude via votre propre compte Anthropic.

### Puis-je réutiliser mes extensions VS Code ?

Oui. Antigravity étant un fork profondément modifié de Visual Studio Code, vos extensions, thèmes et raccourcis sont importables au premier lancement. Si l’import n’a pas eu lieu, relancez-le via la palette de commandes (« Import VS Code settings »).

### Quelle différence entre l’Éditeur et le Manager ?

L’Éditeur est une IDE synchrone classique (complétion, commandes en ligne) pour le travail que vous supervisez activement. Le Manager orchestre des agents autonomes et asynchrones, en parallèle, sur plusieurs espaces de travail. On délègue le gros œuvre au Manager et on peaufine dans l’Éditeur.

### Antigravity fonctionne-t-il hors ligne ?

Non pour les fonctions IA : les agents dialoguent avec les serveurs de Google (ou d’Anthropic pour Claude). L’éditeur de base reste utilisable hors ligne comme n’importe quel VS Code, mais la complétion IA, les missions et le sous-agent navigateur exigent une connexion internet.

### Quels sont les risques de sécurité ?

