---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-11
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [539, 549]
sha256: 08762fe75500f246c227658ec330e635e6c35c9d22c803f6bcaddb3d33d1df25
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

The RADIUS server may have RADIUS attributes with the same attribute IDs and names as but different encapsulation formats or contents from those on the device. In this case, you can configure the RADIUS attribute disablement function to disable such attributes. The device then does not parse these attributes after receiving them from the RADIUS server, and does not encapsulate these attributes into RADIUS packets to be sent to the server.
Currently, RADIUS attributes supported by Huawei(with the attribute names and IDs supported by Huawei) in a sent or received packet can be disabled on a device.
RADIUS attribute translation is used for achieve compatibility between RADIUS attributes defined by different vendors. For example, a Huawei device delivers the priority of an administrator using the Huawei proprietary attribute Exec-Privilege (26-29), whereas another vendor's NAS and the RADIUS server deliver this priority using the Login-service (15) attribute. In a scenario where the Huawei device and another vendor's NAS share one RADIUS server, users want the Huawei device to be compatible with the Login-service (15) attribute. After RADIUS attribute translation is configured on the Huawei device, the device automatically processes the Login-service (15) attribute in a received RADIUS authentication response packet as the Exec-Privilege (26-29) attribute.
The device supports translation between RADIUS attributes supported and unsupported by Huawei. Table 1-16 describes the translation modes.
The device can translate a RADIUS attribute of another vendor only if the length of the Type field in the attribute is 1 octet.
The device can translate the RADIUS attribute only when the type of the source RADIUS attribute is the same as that of the destination RADIUS attribute. For example, the types of NAS-Identifier and NAS-Port-Id attributes are string, and they can be translated into each other. The types of NAS-Identifier and NAS-Port attributes are string and integer respectively, they cannot be translated into each other.
| Whether the Source RADIUS Attribute Is Supported by Huawei | Whether the Destination RADIUS Attribute Is Supported by Huawei | Supported Translation Direction | Configuration Command (RADIUS Server Template View) | 
|---|---|---|---|
| Supported | Supported | Transmit and receive directions | radius-attribute translate src-attribute-name dest-attribute-name { receive \| send \| access-accept \| access-request \| account-request \| account-response } * | 
| Supported | Not supported | Transmit direction | radius-attribute translate extend src-attribute-name vendor-specific dest-vendor-id dest-sub-id { access-request \| account-request } * | 
| Not supported | Supported | Receive direction | radius-attribute translate extend vendor-specific src-vendor-id src-sub-id dest-attribute-name { access-accept \| account-response } * |
