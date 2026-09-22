---
id: briefing-general-tech-2026/06-consumer-devices/03-duo-specs
title: "Duo specifications in detail"
domain: consumer-devices
role: deep-dive
task: consumer
actors: ["Apple", "Samsung"]
dates: ["2026-09-10", "2026-10-16", "2026-10-23"]
keywords: ["crease", "foldable"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g07-3"
source_lines: [6818, 6939]
sha256: 7851ac54f883363f731a89c91641d4e267220d8f1dd63497322b6f20ec47f795
---

# Duo specifications in detail

<a id="g07-3"></a>
### 7.3 Duo specifications in detail

The iPhone Duo's specification sheet reads as a series of deliberate answers to the
foldable's known weaknesses — the crease, the weight, the water resistance, the
software — with one conspicuous absence that will generate its own debate. The
table below consolidates the confirmed specifications; the commentary that
follows addresses each choice in turn.

| Specification | iPhone Duo (confirmed) |
|---|---|
| External display | 5.4-inch Super Retina XDR, 120 Hz |
| Internal display | 7.6-inch foldable Super Retina XDR OLED, 120 Hz — "largest display ever on an iPhone" |
| Design language | "Passport" design: compact when closed, same aspect ratio on both screens |
| Chip | A20 Pro |
| Modem | Apple C2 |
| Chassis | Grade-5 titanium |
| Thickness | 5.2 mm open / 11.3 mm closed |
| Weight | 254 g |
| Biometrics | No Face ID — side Touch ID, plus Apple Watch unlock |
| Front camera (internal) | Under-display FaceTime camera |
| Stylus | Apple Pencil support (first Pencil-compatible iPhone) |
| Software | iOS 27, dual-app multitasking |
| Water resistance | IP68: 6 meters for 30 minutes per IEC 60529 |
| Availability | Preorders 16/10/2026, on sale 23/10/2026, 70+ countries (announced/planned) |

The displays are the heart of the product. The external 5.4-inch panel is a Super Retina XDR
display at 120 Hz — small by modern flagship standards, deliberately so, because it is meant
to be a phone screen in the literal sense: usable one-handed, compact, a cover display rather
than a shrunken tablet. The internal 7.6-inch foldable OLED, also 120 Hz Super Retina XDR, is
described by Apple as the largest display ever put on an iPhone, and the claim is literal: no
previous iPhone display has approached 7.6 inches diagonally. The "passport" design
description — compact when closed, with the same aspect ratio on both the external and
internal screens — is Apple's answer to one of the oldest foldable complaints: cover screens
that are too narrow to type on comfortably, and internal screens whose proportions feel alien
next to the phone in your pocket. Whether the execution matches the intent will be decided by
hands-on testing after October 23.

The silicon is conventionally Apple: an A20 Pro chip and the Apple C2 modem, the company's
second-generation in-house cellular modem. The chassis is grade-5 titanium — the same
aerospace-grade alloy Apple uses on its Pro phones — which makes the thickness figures more
impressive than they first appear. At 5.2 mm open and 11.3 mm closed, the Duo is
extraordinarily thin for a foldable; a closed thickness of 11.3 mm is only a few millimeters
beyond a conventional flagship. But titanium is dense, and thinness has a price in grams, as
section 7.4 will show.

The most debated specification is the one that is missing: there is no Face ID. Apple has
replaced it with a side-mounted Touch ID sensor, supplemented by Apple Watch-based unlocking.
This is a genuine regression in biometric convenience for users who have spent nearly a
decade with face unlock, and Apple has not publicly explained the engineering reason — though
the under-display FaceTime camera on the internal screen suggests the company judged that
foldable display layers and the TrueDepth sensor array do not yet coexist well. It is the
single most consequential compromise in the specification sheet, and it will be the detail
reviewers return to most often.

Two software-adjacent details complete the picture. The iPhone Duo is the first iPhone
compatible with the Apple Pencil — though with the precise caveats that section 7.4
dissects — and it runs iOS 27 with a dual-app multitasking mode designed for the unfolded
7.6-inch canvas. The multitasking implementation matters more than it sounds: a foldable
whose software treats the big screen as merely a larger phone screen fails the category's
central promise, and Apple's decision to ship a purpose-built dual-app mode at launch
suggests the company understands that.

Finally, water resistance. Apple rates the Duo at IP68: submersion to 6 meters for
30 minutes under IEC 60529, as reported via MacObserver on September 10, 2026. That is a
strong rating for any foldable — the hinge has historically been the weak point — and it
compares favorably with Samsung's IP48 on the Z Fold 8 (section 7.7). But Apple's standard
caveat applies and is worth quoting in spirit: water resistance is not a permanent condition
and degrades with normal wear. On a device whose defining feature is a mechanical hinge that
opens and closes hundreds of times a month, that caveat is more than boilerplate. It is the
honest fine print under the headline number.

Read as a whole, the specification sheet describes three engineering bets.

The first bet is on thinness over lightness. Grade-5 titanium plus a closed
thickness of 11.3 mm and an open thickness of 5.2 mm produce a device that is
remarkably slim for a foldable — but titanium is dense, and the bill comes
due at 254 grams. The tradeoff, stated plainly:

| Priority | Duo's choice | Cost |
|---|---|---|
| Thinness | 5.2 mm open / 11.3 mm closed | — |
| Rigidity and premium feel | Grade-5 titanium chassis | 254 g total weight |
| Weight | — | 53 g heavier than the Z Fold 8 (~26%) |

Apple has decided that a thin foldable that feels dense and rigid is
preferable to a light foldable that feels less substantial. It is a
characteristically Apple tradeoff — the company has always preferred mass
that reads as quality over lightness that reads as cheap — but it is also
the specification most likely to be weaponized against the Duo in
comparisons, because grams are the one dimension where no marketing copy
helps.

The second bet is on the "passport" aspect ratio. Foldables have historically
failed on the cover screen: too narrow for comfortable typing, too cramped
for real use, relegating the closed device to a notification pager. By giving
both screens the same aspect ratio and keeping the closed device compact,
Apple is arguing that the cover display should be a complete phone, not a
compromise — and that the internal display should feel like the same phone,
larger, rather than a different device with different proportions. If the
execution matches the description, this solves the category's oldest
usability complaint. If it does not, the Duo inherits it.

The third bet is the most visible compromise: biometrics. No Face ID, a
side-mounted Touch ID sensor, Apple Watch unlock as backup, and an
under-display FaceTime camera on the internal screen. The combination
suggests a single underlying judgment — that the TrueDepth array and a
folding display stack do not yet coexist to Apple's standard — expressed as
three separate product decisions. It is the specification sheet's honest
scar: every other choice on the page is about what Apple could do, and this
one is about what it could not. Reviewers will return to it, switchers will
ask about it, and Apple's answer — presumably, that Touch ID is fast,
familiar, and sufficient — will be tested against a decade of face-unlock
habit.

Two quieter details complete the picture. The Apple C2 modem — second-
generation in-house cellular silicon — signals the modem program has
graduated to Apple's most complex device. And iOS 27's dual-app
multitasking is the acknowledgment that a foldable whose software treats
the big screen as merely a larger phone screen fails the category's
central promise.

