---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-1
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [1, 110]
sha256: fed6eec9e3acca2cca43ad053d1d669e3acab5592dab067043c37deab9cfcceb
---

# git version 2.47.1

GitLab CI/CD est devenu l’outil de référence pour automatiser les tests, le build et le déploiement d’applications, porté par une croissance commerciale solide : au deuxième trimestre de son exercice fiscal 2027 (clos en juillet 2026), GitLab a déclaré un chiffre d’affaires de 286,3 millions de dollars (+21 % sur un an), tiré notamment par les abonnements CI/CD, avec une base clients passée de 10 338 à 11 114 comptes (+8 %) selon le 10-Q déposé auprès de la SEC. Avec la version 18.10, publiée le 19 mars 2026, la plateforme intègre désormais l’auto-remédiation des vulnérabilités et l’IA GitLab Duo pour le débogage de pipelines — une IA que GitLab met gratuitement à disposition de 100 % de ses clients sans abonnement Duo séparé depuis janvier 2026, selon DevOps.com — ainsi que les inputs dynamiques pour jobs manuels (menus déroulants en cascade), apparus dès la version 18.7 en février 2026 ; les correctifs 18.10.1, 18.9.3 et 18.8.7, sortis le 11 mars 2026, ont par ailleurs consolidé la CI/CD sur les éditions Community et Enterprise. Ce tutoriel GitLab CI/CD vous guide pas à pas, de l’installation à un pipeline de production complet avec Docker, tests automatisés, SAST et déploiement multi-environnements.

Que vous migriez depuis Jenkins (utilisé par 28 % des organisations) ou GitHub Actions (33 %), ou que vous découvriez l’intégration continue — pratique que 55 % des développeurs déclarent utiliser régulièrement selon le State of Developer Ecosystem 2025 de JetBrains, relayé par Quash en mars 2026, où GitLab CI ne recueille pour l’instant que 19 % d’adoption —, ce guide couvre chaque étape avec des fichiers `.gitlab-ci.yml` fonctionnels, des exemples de sortie console et des solutions aux erreurs les plus fréquentes. À la fin de ce tutoriel, vous disposerez d’un pipeline GitLab CI/CD de production prêt à l’emploi.

## Prérequis et Versions Nécessaires

Avant de commencer ce tutoriel GitLab CI/CD, assurez-vous de disposer des éléments suivants. Chaque version a été testée et validée pour garantir la compatibilité avec les exemples de ce guide.

| Outil | Version Minimale | Version Recommandée | Rôle | 
|---|---|---|---|
| GitLab (SaaS ou Self-Managed) | 17.0 | 18.10+ | Plateforme CI/CD | 
| Git | 2.39 | 2.47+ | Contrôle de version | 
| Docker Engine | 24.0 | 27.0+ | Conteneurisation des builds | 
| Node.js (projet démo) | 20 LTS | 22 LTS | Runtime applicatif | 
| GitLab Runner | 17.0 | 18.10+ | Exécution des jobs | 
| Compte GitLab | Free | Premium ou Ultimate | Accès aux fonctionnalités avancées | 

Le tier Free de GitLab offre 400 minutes CI/CD par mois sans frais par utilisateur ainsi que 10 Gio de stockage, un forfait confirmé aussi bien par l’aperçu d’eesel.ai en septembre 2025 que par la revue WorkflowAutomation.net d’avril 2026. Pour des projets d’entreprise, le tier Premium — facturé 24 $/utilisateur/mois selon TechRepublic en novembre 2025, puis relevé à 29 $/utilisateur/mois d’après la page tarifs officielle de GitLab mise à jour en août 2026 et confirmé par la revue Q4 2025 d’eesel.ai — inclut 10 000 minutes CI/CD mensuelles ainsi que les merge request approvals et les environnements protégés, tandis qu’Ultimate (99 $/utilisateur/mois, un tarif encore décrit par TechRepublic en novembre 2025 puis par Octopus Deploy en septembre 2026) débloque 50 000 minutes CI/CD, le DAST, le dependency scanning et les politiques de conformité. Au-delà de ce quota, GitLab facture désormais 10 $ par tranche de 1 000 minutes supplémentaires depuis juin 2026 selon le blueprint UsagePricing : à titre d’exemple, une équipe de 25 utilisateurs sur Premium paie 725 $/mois de licences plus 400 $ pour 40 000 minutes additionnelles. Pour suivre ce tutoriel dans son intégralité, le tier Free suffit pour les 10 premières étapes.

Vérifiez vos installations avec ces commandes :

```
git --version
# git version 2.47.1
docker --version
# Docker version 27.5.1, build 9f9e405
node --version
# v22.14.0
gitlab-runner --version
# Version: 18.10.0
```
## Étape 1 : Créer le Projet et Initialiser le Dépôt

La première étape de tout pipeline GitLab CI/CD consiste à créer un projet sur GitLab et à y pousser votre code. GitLab organise les projets en groupes (équivalent des organisations GitHub), ce qui facilite la gestion des permissions et des variables partagées.

Connectez-vous à `gitlab.com` (ou votre instance self-managed), cliquez sur **New project**, puis **Create blank project**. Nommez-le `demo-cicd-2026`, choisissez la visibilité Private, et cochez **Initialize repository with a README**.

Clonez ensuite le dépôt et créez la structure du projet Node.js :

```
git clone [email protected]:votre-groupe/demo-cicd-2026.git
cd demo-cicd-2026
# Initialiser le projet Node.js
npm init -y
npm install express --save
npm install jest supertest --save-dev
# Créer la structure
mkdir -p src tests
```
Créez le fichier `src/app.js` avec une API Express minimaliste :

```
// src/app.js
const express = require('express');
const app = express();
app.use(express.json());
app.get('/health', (req, res) => {
  res.json({ status: 'ok', version: process.env.APP_VERSION || '1.0.0' });
});
app.get('/api/users', (req, res) => {
  res.json([
    { id: 1, name: 'Alice Martin' },
    { id: 2, name: 'Bob Dupont' }
  ]);
});
module.exports = app;
// Démarrage si exécuté directement
if (require.main === module) {
  const PORT = process.env.PORT || 3000;
  app.listen(PORT, () => console.log(`Serveur démarré sur le port ${PORT}`));
}
```
Ajoutez un fichier de test `tests/app.test.js` :

```
// tests/app.test.js
const request = require('supertest');
const app = require('../src/app');
describe('API Health Check', () => {
  test('GET /health retourne status ok', async () => {
    const res = await request(app).get('/health');
    expect(res.statusCode).toBe(200);
    expect(res.body.status).toBe('ok');
  });
  test('GET /api/users retourne la liste', async () => {
    const res = await request(app).get('/api/users');
    expect(res.statusCode).toBe(200);
    expect(res.body).toHaveLength(2);
  });
});
```
Mettez à jour le `package.json` pour ajouter les scripts de test :

```
{
  "scripts": {
    "start": "node src/app.js",
    "test": "jest --coverage --forceExit",
    "lint": "eslint src/ tests/"
  }
}
```
Poussez le tout vers GitLab avec `git add . && git commit -m "feat: initial project setup" && git push origin main`. Le projet est maintenant prêt à recevoir son premier pipeline CI/CD.

## Étape 2 : Comprendre l’Architecture de GitLab CI/CD

Avant de rédiger le fichier `.gitlab-ci.yml`, il est essentiel de comprendre comment GitLab CI/CD orchestre l’exécution. L’architecture repose sur trois composants principaux : le **serveur GitLab** qui parse le fichier YAML et planifie les jobs, les **runners** qui exécutent physiquement les commandes, et le **registre d’artefacts** qui stocke les résultats entre les stages.

Un pipeline GitLab CI/CD se décompose en **stages** (étapes séquentielles) contenant des **jobs** (tâches parallèles). Par défaut, tous les jobs d’un même stage s’exécutent en parallèle, puis le stage suivant démarre uniquement si tous les jobs précédents ont réussi. Ce modèle permet d’accélérer les pipelines tout en maintenant l’intégrité du processus.

