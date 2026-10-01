---
id: collect-261001-general-networking/general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026-2
title: "OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA"
domain: general-networking
role: reference
task: reference
actors: ["CISA", "Falcon", "Google", "Oracle"]
dates: []
keywords: ["arr", "cyber", "mai", "valuation"]
source: docs/RAG/collect-261001-general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026.md
source_anchor: ""
source_lines: [42, 90]
sha256: 9e4bf3922843cd69862de44957bba414103819e3f1877c1a65678475bf78428e
---

# OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA

La **cryptographie post-quantique** désigne des algorithmes exécutés sur des ordinateurs classiques mais conçus pour résister à un attaquant équipé d’une machine quantique. Contrairement à la cryptographie quantique (QKD), qui exige du matériel photonique dédié, la PQC est un remplacement logiciel : de nouveaux algorithmes mathématiques, fondés notamment sur les réseaux euclidiens (*lattices*) ou les fonctions de hachage, viennent se substituer à RSA et à l’ECC.

En août 2024, le NIST a finalisé les trois premiers standards, socle de toute la transition mondiale. Publié en mars 2025, le rapport NIST IR 8545 documente précisément le statut de la « quatrième phase » (Fourth Round) du concours de standardisation : il y confirme la sélection de HQC comme mécanisme d’encapsulation de clés de secours, fondé sur une famille mathématique différente pour ne pas mettre tous les œufs dans le même panier, et dresse la liste des cinq algorithmes désormais considérés comme prioritaires : ML-KEM, ML-DSA, SLH-DSA, FN-DSA et HQC. La dynamique ne s’est pas arrêtée là : en mai 2026, le rapport NIST IR 8610 a fait état de quatre schémas de signature supplémentaires ayant franchi la deuxième phase d’évaluation, preuve que le chantier de standardisation reste actif au-delà des quatre familles déjà figées. Le tableau suivant récapitule les briques standardisées et ce qu’elles remplacent.

| Standard | Nom (ex-candidat) | Fonction | Remplace | 
|---|---|---|---|
| FIPS 203 | ML-KEM (ex-CRYSTALS-Kyber) | Établissement de clés (KEM) | RSA, ECDH | 
| FIPS 204 | ML-DSA (ex-CRYSTALS-Dilithium) | Signature numérique | RSA, ECDSA | 
| FIPS 205 | SLH-DSA (ex-SPHINCS+) | Signature fondée sur le hachage | RSA, ECDSA | 
| FIPS 206 (en préparation) | FN-DSA (ex-Falcon) | Signature compacte | RSA, ECDSA | 
| Sélection 2025 | HQC | KEM de secours (code correcteur) | Complément de ML-KEM | 

Ces algorithmes sont désormais accessibles aux développeurs, et les adoptions concrètes se multiplient bien au-delà du seul monde de la sécurité. OpenSSL 3.5, publié en avril 2025, intègre nativement ML-KEM, ML-DSA et SLH-DSA, tandis que des navigateurs et CDN majeurs déploient déjà l’échange de clés hybride X25519MLKEM768 en TLS 1.3 ; Oracle a rejoint le mouvement en intégrant ce même échange de clés hybride post-quantique en TLS 1.3 à Java 27, disponible depuis septembre 2026. Le phénomène déborde même sur la blockchain : Algorand a fait passer ses comptes post-quantiques natifs en production sur son réseau principal dès août 2026 avec la version 5.0.0, sa feuille de route promettant en outre la prise en charge de la signature Falcon-1024 dès la publication du troisième trimestre 2026. Le NIST continue d’ailleurs d’affiner le mode d’emploi : le projet de norme SP 800-133 Rev. 3, publié en avril 2026, met à jour les règles de génération de clés pour ces quatre algorithmes FIPS post-quantiques, signe que la spécification technique elle-même est encore en train de se stabiliser. Vous pouvez le vérifier en quelques commandes.

```
# OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA
openssl list -kem-algorithms | grep -i mlkem
# Négocier un échange de clés hybride (X25519 + ML-KEM-768) en TLS 1.3
openssl s_client -connect exemple.fr:443 -groups X25519MLKEM768
# Générer une paire de clés de signature post-quantique (ML-DSA-65)
openssl genpkey -algorithm ML-DSA-65 -out cle_pq.pem
```
### Le pari de la cryptographie hybride

L’ANSSI ne demande pas de jeter la cryptographie classique : elle exige de la *combiner* avec la PQC. Cette approche hybride protège contre deux risques opposés. D’un côté, une faille encore inconnue dans un algorithme post-quantique récent ; de l’autre, l’arrivée de l’ordinateur quantique. En chaînant X25519 et ML-KEM, un attaquant devrait casser les deux mécanismes pour réussir. Le document « Agreed Cryptographic Mechanisms v2 » de l’ECCG (mai 2025) va jusqu’à préciser que les mécanismes fondés sur les réseaux LWE et MLWE ne devraient pas être utilisés seuls. L’agence étend cette prudence aux signatures, jugées plus récentes et moins éprouvées que l’établissement de clés. Le NIST a d’ailleurs illustré cette même prudence en avril 2026 : le projet de norme SP 800-230 a introduit de nouveaux jeux de paramètres SLH-DSA taillés pour des cas de signature à usage limité, une façon de donner plus de marge de manœuvre aux implémenteurs les plus exposés.

## La qualification ANSSI, sésame du marché public français

Pour comprendre la brutalité potentielle de la mesure, il faut saisir ce qu’est la **certification ANSSI**. L’agence délivre plusieurs niveaux d’assurance : la Certification de Sécurité de Premier Niveau (CSPN), les Critères Communs (CC), et surtout la qualification (élémentaire, standard, renforcée), qui est une recommandation officielle d’usage par l’État. Sans qualification, un HSM, une passerelle VPN ou une solution de messagerie chiffrée est de facto exclu des appels d’offres des ministères et des opérateurs d’importance vitale (OIV).

Or un cycle de qualification dure généralement de 12 à 18 mois. Un produit qui entre en évaluation à la mi-2026 dispose donc d’une fenêtre serrée pour décrocher son label avant que la porte ne se referme en 2027. Les éditeurs qui n’ont pas encore intégré ces algorithmes post-quantiques à leur feuille de route se retrouvent, mécaniquement, en retard d’un cycle. Cette contrainte temporelle transforme l’annonce de juin 2026 en signal d’alarme : ce n’est pas 2030 qu’il faut regarder, mais l’année qui vient.

La contrainte s’articule par ailleurs avec la directive NIS2, dont la transposition française élargit considérablement le périmètre des entités « essentielles » et « importantes » tenues à un haut niveau de sécurité. Pour ces milliers d’organisations, la conformité passera de plus en plus par des briques technologiques elles-mêmes qualifiées – et donc, demain, post-quantiques.

## Une convergence internationale : NSA, UE, NCSC

La France n’invente pas un calendrier isolé. Les grandes autorités cyber occidentales alignent leurs échéances autour d’un même horizon 2030-2035, avec des jalons intermédiaires en 2027-2028. Le tableau ci-dessous compare les feuilles de route des principaux régulateurs.

| Autorité | Jalon intermédiaire | Cible finale | Portée | 
|---|---|---|---|
| ANSSI (France) | Arrêt des certifications sans PQC dès 2027 | Achats « quantum-safe » d’ici 2030 | Produits qualifiés / État & OIV | 
| NSA (États-Unis, CNSA 2.0) | Prise en charge exigée au 1er janv. 2027 | Usage exclusif d’ici 2030-2033 | Systèmes de sécurité nationale | 
| Union européenne | Feuilles de route nationales fin 2026 | Haut risque en 2030, complet en 2035 | Infrastructures critiques | 
| NCSC (Royaume-Uni) | Inventaire et priorisation d’ici 2028 | Migration complète d’ici 2035 | Secteurs critiques & grandes entreprises | 

Cette convergence n’est pas fortuite, et les hyperscalers l’ont déjà anticipée : Google Cloud a publié en août 2026 une feuille de route confirmant que ses points de terminaison API proposent désormais un échange de clés quantum-safe, avec une compatibilité PQC complète visée pour 2029. Elle crée un effet de cliquet : dès lors que la France, les États-Unis et les principaux fournisseurs cloud exigent ou proposent la PQC à partir de 2027, aucun grand éditeur international ne peut se permettre de proposer deux gammes de produits. La cryptographie résistante au quantique devient, de fait, le nouveau standard mondial par défaut.

## L’Europe et sa feuille de route coordonnée

