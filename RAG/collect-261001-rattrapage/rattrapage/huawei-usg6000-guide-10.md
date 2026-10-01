---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-10
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1404, 1557]
sha256: f99b6f1cebdfc7f777dcf1a2b720afc8b7c3869932f5650c0d422bb921c348bd
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

Stratégie terrain :
- **LAN→Internet** : IPS en mode **block** sur les signatures critiques/hautes, **alert** sur les moyennes (pour éviter les faux positifs bloquants au début).
- **untrust→DMZ** : IPS en **block** strict — c'est ta vitrine, elle est scannée en permanence.
- **Phase pilote** : 1 à 2 semaines en **alert only**, analyse des logs, puis bascule en block. Ça évite le « l'IPS a coupé la prod » du lundi matin.

## 78. Filtrage URL : catégories et listes

Base de données : **130+ catégories, 500M+ d'URL** (cloud Huawei, mise à jour auto).

```huawei
system-view
[USG] profile type url-filter name url-profile-lan
# Exemple de logique : bloquer les catégories à risque, alerter le reste
[USG-profile-url-filter-url-profile-lan] quit
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] profile url-profile-lan
[USG-policy-security-rule-LAN-web-out] quit
[USG-policy-security] quit
save
```

Catégories à bloquer en entreprise (à adapter à ta charte informatique) :
- **Bloquer** : malware, phishing, pornographie, jeux d'argent, P2P/torrents, proxys d'anonymisation.
- **Limiter** (quota temps ou alerte) : réseaux sociaux, streaming vidéo, shopping.
- **Listes blanches/noires personnalisées** : toujours prévoir une whitelist pour les faux positifs (le site d'un fournisseur classé « parking de domaine » par erreur, ça arrive).

⚠️ HTTPS : sans **inspection SSL**, le filtrage URL ne voit que le **nom de domaine** (via SNI), pas l'URL complète. `https://malin.example.com/virus.exe` sera vu comme `malin.example.com`.

## 79. Contrôle applicatif : bloquer par application, pas par port

6000+ applications reconnues (signatures + comportement, pas les ports).

```huawei
system-view
[USG] profile type app name app-profile-lan
[USG-profile-app-app-profile-lan] quit
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] profile app-profile-lan
[USG-policy-security-rule-LAN-web-out] quit
# Ou directement dans la règle (filtrage par application, sans profil) :
[USG-policy-security] rule name bloquer-p2p-strict
[USG-policy-security-rule-bloquer-p2p-strict] source-zone trust
[USG-policy-security-rule-bloquer-p2p-strict] destination-zone untrust
[USG-policy-security-rule-bloquer-p2p-strict] application p2p
[USG-policy-security-rule-bloquer-p2p-strict] action deny
[USG-policy-security-rule-bloquer-p2p-strict] quit
[USG-policy-security] quit
save
```

Cas d'usage : bloquer **BitTorrent, Tor, applications de contournement**, limiter **YouTube/Netflix** en bande passante (avec la gestion de bande passante, section 80), autoriser **WhatsApp** mais pas les appels vidéo, etc.

## 80. Gestion de bande passante par application/utilisateur

```huawei
system-view
# Limiter le streaming à 50 Mbit/s au total sur la règle LAN→Internet :
[USG] traffic-policy
[USG-policy-traffic] profile bandwidth-profile-streaming
[USG-traffic-profile-bandwidth-profile-streaming] bandwidth maximum 50000  # kbit/s
[USG-traffic-profile-bandwidth-profile-streaming] quit
[USG-policy-traffic] rule name limiter-streaming
[USG-policy-traffic-rule-limiter-streaming] source-zone trust
[USG-policy-traffic-rule-limiter-streaming] destination-zone untrust
[USG-policy-traffic-rule-limiter-streaming] application streaming-media
[USG-policy-traffic-rule-limiter-streaming] action bandwidth profile bandwidth-profile-streaming
[USG-policy-traffic-rule-limiter-streaming] quit
[USG-policy-traffic] quit
save
```

💡 **Garantir** la bande passante de la VoIP/appli métier (minimum garanti) est souvent plus utile que de limiter le reste. Les deux se combinent.

## 81. Anti-spam : protéger la messagerie

L'anti-spam USG filtre les mails entrants (SMTP) : RBL (listes noires), analyse de contenu, anti-phishing.

```huawei
system-view
[USG] profile type antispam name antispam-profile-mail
[USG-profile-antispam-antispam-profile-mail] quit
[USG] security-policy
[USG-policy-security] rule name mail-entrant
[USG-policy-security-rule-mail-entrant] source-zone untrust
[USG-policy-security-rule-mail-entrant] destination-zone dmz
[USG-policy-security-rule-mail-entrant] destination-address 172.16.1.25 32  # serveur mail
[USG-policy-security-rule-mail-entrant] service smtp
[USG-policy-security-rule-mail-entrant] profile antispam-profile-mail
[USG-policy-security-rule-mail-entrant] action permit
[USG-policy-security-rule-mail-entrant] quit
[USG-policy-security] quit
save
```

⚠️ L'anti-spam du firewall est un **complément**, pas un remplaçant d'une vraie passerelle mail (qui gère aussi l'anti-phishing avancé, le sandboxing de pièces jointes, la continuité). Si ton flux mail est critique, envisage une solution dédiée.

## 82. Inspection SSL/TLS : puissance et précautions

```huawei
# Principe : l'USG se place en "man-in-the-middle" légitime avec un certificat CA interne
# déployé sur les postes (via GPO/AD).
# Activer l'inspection sur une règle (syntaxe à vérifier selon version) :
[USG] security-policy
[USG-policy-security] rule name LAN-web-inspecte
[USG-policy-security-rule-LAN-web-inspecte] source-zone trust
[USG-policy-security-rule-LAN-web-inspecte] destination-zone untrust
[USG-policy-security-rule-LAN-web-inspecte] service https
[USG-policy-security-rule-LAN-web-inspecte] profile av-profile-lan
[USG-policy-security-rule-LAN-web-inspecte] action permit
# + activer le déchiffrement SSL dans les paramètres de la règle / du profil
[USG-policy-security-rule-LAN-web-inspecte] quit
[USG-policy-security] quit
save
```

⚠️ **Triple avertissement :**
1. **Légal** : en France/Europe, l'inspection du trafic chiffré des salariés exige **information des représentants du personnel / charte informatique**, et des **exclusions** (banque, santé...). Valide avec ta direction/juriste.
2. **Technique** : déploie la **CA interne sur tous les postes** (sinon alertes certificat partout) ; certaines applis (bancaires, avec certificate pinning) **cassent** — prévois une liste d'exclusion.
3. **Performance** : c'est la fonction la plus gourmande (section 73). Ne l'active que là où c'est justifié.

## 83. Filtrage de fichiers : bloquer les types dangereux

```huawei
system-view
[USG] profile type file-filter name file-profile-lan
# Bloquer : .exe, .bat, .ps1, .vbs, .js en téléchargement web/mail (selon version)
[USG-profile-file-filter-file-profile-lan] quit
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] profile file-profile-lan
[USG-policy-security-rule-LAN-web-out] quit
[USG-policy-security] quit
save
```

💡 Moins violent que l'AV seul : le file-filter bloque **par extension/type**, même sans signature. Parfait contre les `.exe` « cadeaux » des campagnes de phishing.

## 84. Politique de contenu : assembler les profils par zone (matrice recommandée)

| Flux | AV | IPS | URL | APP | File-filter | Inspection SSL |
|------|----|-----|-----|-----|-------------|----------------|
| trust→untrust (LAN) | ✅ block | ✅ block/alert | ✅ | ✅ | ✅ | ⚠️ ciblé |
| GUEST→untrust | ✅ block | ✅ block | ✅ strict | ✅ strict | ✅ | ❌ |
| untrust→dmz | ✅ block | ✅ block strict | ❌ | ❌ | ✅ | ❌ |
| dmz→untrust | ❌ | ✅ alert | ❌ | ❌ | ❌ | ❌ |
| trust→dmz | ❌ | ✅ alert | ❌ | ❌ | ❌ | ❌ |
| VPN→trust | ✅ block | ✅ block | ❌ | ❌ | ❌ | ❌ |

Légende : ✅ recommandé, ⚠️ selon contexte légal/technique, ❌ inutile (coût perf sans bénéfice).

## 85. Tester l'UTM : le fichier EICAR et les sites de test

