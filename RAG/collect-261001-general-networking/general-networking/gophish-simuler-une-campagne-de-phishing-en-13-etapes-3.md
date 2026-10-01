---
id: collect-261001-general-networking/general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes-3
title: "1. Mise à jour système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["dpo"]
source: docs/RAG/collect-261001-general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes.md
source_anchor: ""
source_lines: [188, 255]
sha256: 2490139380c2dce68a2acdc017e109b4990059eda5ba4a3bb9062f1df87b9a5b
---

# 1. Mise à jour système

```
First Name,Last Name,Email,Position
Camille,Girard,[email protected],Responsable Comptabilité
Karim,Benali,[email protected],Technicien Support
Léa,Fontaine,[email protected],Chargée RH
```
Ne dépassez pas la portée définie dans votre autorisation écrite. Si l'accord signé couvre le service comptabilité, n'importez pas l'ensemble de l'annuaire : chaque extension de périmètre doit être validée séparément avec la direction et le DPO.

## Étape 11 : lancer la campagne

Dans Campaigns, créez une nouvelle campagne en assemblant les briques précédentes : nom, modèle d'email, page d'atterrissage, URL du serveur de phishing, profil d'envoi SMTP, et groupe(s) de cibles. Vous pouvez soit lancer immédiatement, soit programmer une date et heure de démarrage.

| Paramètre | Valeur recommandée | 
|---|---|
| Name | Nom explicite incluant la date, ex. « Sensibilisation T3-2026-Compta » | 
| URL | Domaine dédié à la simulation, jamais un sous-domaine du site officiel | 
| Launch Date | Hors période de forte charge (éviter fin de mois pour la compta, par exemple) | 
| Send Emails By | Étalé sur plusieurs heures plutôt qu'un envoi groupé instantané, pour éviter de déclencher les filtres antispam | 

Vérifiez une dernière fois chaque élément avant de cliquer sur Launch Campaign : une fois lancée, une campagne ne peut pas être interrompue proprement sans laisser certains emails déjà partis en file d'attente SMTP.

## Étape 12 : lire et interpréter les résultats

La page de résultats de chaque campagne affiche, pour chaque destinataire, une chronologie d'événements. GoPhish distingue quatre statuts principaux, chacun révélant un niveau de vigilance différent chez l'employé testé.

| Statut | Déclencheur technique | Ce que cela révèle | 
|---|---|---|
| Email Sent | L'email a été remis au serveur SMTP | Aucune information sur le comportement de la cible | 
| Email Opened | Chargement du pixel de tracking intégré | La cible a ouvert l'email (dépend du client de messagerie) | 
| Clicked Link | Clic sur le lien de tracking unique | La cible a cliqué sans vérifier la légitimité du lien | 
| Submitted Data | Soumission du formulaire sur la landing page | La cible a saisi des identifiants ou des données sensibles | 
| Email Reported | Signalement via le bouton de report ou transfert à la boîte dédiée | La cible a identifié la tentative et suivi la bonne procédure | 

Exportez les résultats au format CSV pour construire vos indicateurs de sensibilisation : taux de clic, taux de soumission de données, et surtout taux de signalement, qui est l'indicateur le plus révélateur de la maturité sécurité d'une équipe. Un taux de signalement en hausse d'une campagne à l'autre est souvent plus significatif qu'un taux de clic en baisse.

## Étape 13 : sécuriser durablement l'installation

Une fois la première campagne terminée, prenez le temps de durcir l'installation avant tout usage récurrent.

- Placez un certificat TLS valide (Let's Encrypt) sur le serveur de phishing si vous voulez tester la réaction des employés face à un cadenas HTTPS trompeur
- Limitez l'accès à l'interface admin par IP ou VPN, jamais en accès public direct
- Changez le mot de passe admin par défaut immédiatement après le premier login
- Sauvegardez régulièrement le fichier gophish.db, qui contient l'historique complet des campagnes
- Purgez les données sensibles (identifiants saisis) une fois l'exploitation pédagogique terminée, conformément au principe de minimisation RGPD

## Les 5 pièges les plus fréquents

**1. Lancer une campagne sans autorisation écrite.** C'est la faute la plus grave. Même en interne, sans document signé par la direction et le DPO, vous vous exposez personnellement à des poursuites pour intrusion dans un système de traitement automatisé de données. Un simple accord verbal du manager d'équipe ne suffit pas : le document doit être signé par une personne habilitée à engager la responsabilité de l'organisation, et conservé aussi longtemps que la politique de sécurité de l'entreprise l'exige.

**2. Exposer l'interface admin sur Internet.** Le port 3333 ouvert en 0.0.0.0 est régulièrement scanné par des robots. Combiné à une vulnérabilité comme CVE-2025-70963 (voir section suivante), cela peut exposer la clé API de vos comptes administrateurs à n'importe quel script s'exécutant dans le navigateur. Un attaquant qui met la main sur cette clé API obtient un accès équivalent à celui d'un administrateur légitime, avec la possibilité de créer ses propres campagnes ou de récupérer l'historique complet des tests déjà menés dans l'organisation.

**3. Envoyer tous les emails d'un coup.** Un envoi massif instantané déclenche presque systématiquement les filtres antispam et les systèmes de détection d'anomalies de la messagerie de l'entreprise, ce qui fausse totalement les résultats de la campagne. Un pic soudain de connexions sortantes vers un même relais SMTP est aussi un signal classique que surveillent les solutions de détection d'anomalies réseau, ce qui peut déclencher une alerte interne inutile pendant votre propre exercice.

**4. Oublier de configurer SPF/DKIM sur le domaine de simulation.** Sans ces enregistrements DNS, les emails de test finiront en spam avant même d'atteindre la boîte de réception, rendant le test inutile. Pire, un domaine sans SPF/DKIM correctement configuré peut voir sa réputation dégradée auprès des grands fournisseurs de messagerie, ce qui compliquera l'envoi de futures campagnes légitimes depuis la même infrastructure.

**5. Réutiliser les résultats individuels à des fins disciplinaires.** Sanctionner nommément un employé qui a cliqué détruit la confiance nécessaire à toute démarche de sensibilisation future et pose un problème direct de conformité RGPD sur la finalité du traitement. Une fois qu'un salarié apprend qu'un test de phishing a servi de base à un avertissement, l'ensemble de l'équipe se met sur la défensive et les campagnes suivantes perdent toute valeur pédagogique, les employés cherchant alors à deviner les tests plutôt qu'à développer un réflexe de vigilance sincère.

## Vulnérabilités connues de GoPhish à surveiller

Comme tout logiciel, GoPhish a fait l'objet de plusieurs avis de sécurité qu'il faut connaître avant de le déployer en production.

| Référence | Type | Impact | Statut | 
|---|---|---|---|
| CVE-2026-39904 | Déni de service | Un utilisateur authentifié peut épuiser la mémoire du serveur via une pièce jointe Office piégée dans un modèle d'email | Aucun correctif officiel signalé à la date de l'avis ; limitez la taille des pièces jointes et l'accès à la création de modèles | 
| CVE-2025-70963 | Contrôle d'accès incorrect | La clé API de chaque utilisateur est exposée en clair dans le HTML/JS du tableau de bord admin, exploitable via un script malveillant dans le navigateur | Limitez strictement l'accès à l'interface admin, tournez les clés API régulièrement | 
| CVE-2020-24710 / 24708 / 24709 | Cross-Site Scripting stocké | Injection de script via des champs de modèle ou la fonction d'imitation | Corrigé dans les versions ultérieures à 0.10.1 | 
| CVE-2020-24713 | Expiration de session insuffisante | Le cookie de session restait utilisable après déconnexion | Corrigé | 

