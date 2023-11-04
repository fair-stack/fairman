---
title: t1 v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# t1

> v1.0.0

Base URLs:

# Authentication

# Base

## GET Obtain unique identification uuid

GET /v1/base/uuid

> Return Example

> 200 Response

```json
"string"
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|{"success":true,"data":{},"msg":"Successfully obtained"}|string|

# SysApi

## POST Create Foundationapi

POST /v2/api

> Return Example

> 200 Response

```json
"string"
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|{"success":true,"data":{},"msg":"Created successfully"}|string|

# Docker Base

## POST Upload files

POST /v2/base/file

Upload files

> Body Request parameters

```yaml
file: string

```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|object| no |none|
|» file|body|string(binary)| yes |file|

> Return Example

> 200 Response

```json
{
  "code": 0,
  "data": "string",
  "message": "string"
}
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[response.Response](#schemaresponse.response)|

# docker

## POST Testing a Single Environment

POST /v2/endpoint/one/test

Testing a Single Environment

> Body Request parameters

```json
{
  "TLSConfig": {
    "TLSCACertPath": "string",
    "TLSCertPath": "string",
    "TLSKeyPath": "string",
    "TLSType": 1
  },
  "connected": true,
  "createAt": "string",
  "envURL": "string",
  "id": "string",
  "isSocket": true,
  "isTLS": true,
  "k8sConfig": "string",
  "name": "string",
  "password": "string",
  "securitySettings": {
    "allowBindMountsForRegularUsers": false,
    "allowContainerCapabilitiesForRegularUsers": true,
    "allowDeviceMappingForRegularUsers": true,
    "allowHostNamespaceForRegularUsers": true,
    "allowPrivilegedModeForRegularUsers": false,
    "allowStackManagementForRegularUsers": true,
    "allowSysctlSettingForRegularUsers": true,
    "allowVolumeBrowserForRegularUsers": true,
    "enableHostManagementFeatures": true
  },
  "status": 1,
  "templates": [
    {
      "category": [
        "string"
      ],
      "config": {
        "createAt": "string",
        "dockerConfig": {},
        "dockerImage": "string",
        "dockerRegistry": null,
        "fileSize": 0,
        "id": "string",
        "isDel": true,
        "k8SConfig": null,
        "k8SContent": "string",
        "k8SRegistry": [
          null
        ],
        "manual": null,
        "platform": 0,
        "swarmConfig": {},
        "swarmContent": "string",
        "swarmRegistry": [
          null
        ],
        "templateId": "string",
        "updateAt": "string",
        "userId": "string",
        "version": "string",
        "versionContent": "string"
      },
      "copyright": "string",
      "createAt": "string",
      "devLanguage": [
        "string"
      ],
      "downloads": 0,
      "features": "string",
      "id": "string",
      "isDel": true,
      "logo": "string",
      "name": "string",
      "nameZh": "string",
      "notice": {
        "host": "string",
        "method": "string",
        "query": [
          null
        ],
        "uri": "string"
      },
      "poster": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "preview": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "relationSoft": [
        {
          "category": [
            "string"
          ],
          "config": null,
          "copyright": "string",
          "createAt": "string",
          "devLanguage": [
            "string"
          ],
          "downloads": 0,
          "features": "string",
          "id": "string",
          "isDel": true,
          "logo": "string",
          "name": "string",
          "nameZh": "string",
          "notice": null,
          "poster": [
            {}
          ],
          "preview": [
            {}
          ],
          "relationSoft": [
            {}
          ],
          "relationSoftIds": [
            "string"
          ],
          "score": 0,
          "state": 0,
          "supportMode": null,
          "sysLanguage": [
            "string"
          ],
          "team": "string",
          "teamIcon": "string",
          "updateAt": "string",
          "userId": "string"
        }
      ],
      "relationSoftIds": [
        "string"
      ],
      "score": 0,
      "state": 0,
      "supportMode": {
        "docker": true,
        "swarm": true
      },
      "sysLanguage": [
        "string"
      ],
      "team": "string",
      "teamIcon": "string",
      "updateAt": "string",
      "userId": "string"
    }
  ],
  "type": 0,
  "uniqueId": "string",
  "updateAt": "string",
  "userId": "string",
  "username": "string"
}
```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|[docker.DocEndpoint](#schemadocker.docendpoint)| no |none|

> Return Example

> 200 Response

```json
{
  "TLSConfig": {
    "TLSCACertPath": "string",
    "TLSCertPath": "string",
    "TLSKeyPath": "string",
    "TLSType": 1
  },
  "connected": true,
  "createAt": "string",
  "envURL": "string",
  "id": "string",
  "isSocket": true,
  "isTLS": true,
  "k8sConfig": "string",
  "name": "string",
  "password": "string",
  "securitySettings": {
    "allowBindMountsForRegularUsers": false,
    "allowContainerCapabilitiesForRegularUsers": true,
    "allowDeviceMappingForRegularUsers": true,
    "allowHostNamespaceForRegularUsers": true,
    "allowPrivilegedModeForRegularUsers": false,
    "allowStackManagementForRegularUsers": true,
    "allowSysctlSettingForRegularUsers": true,
    "allowVolumeBrowserForRegularUsers": true,
    "enableHostManagementFeatures": true
  },
  "status": 1,
  "templates": [
    {
      "category": [
        "string"
      ],
      "config": {
        "createAt": "string",
        "dockerConfig": {},
        "dockerImage": "string",
        "dockerRegistry": null,
        "fileSize": 0,
        "id": "string",
        "isDel": true,
        "k8SConfig": null,
        "k8SContent": "string",
        "k8SRegistry": [
          null
        ],
        "manual": null,
        "platform": 0,
        "swarmConfig": {},
        "swarmContent": "string",
        "swarmRegistry": [
          null
        ],
        "templateId": "string",
        "updateAt": "string",
        "userId": "string",
        "version": "string",
        "versionContent": "string"
      },
      "copyright": "string",
      "createAt": "string",
      "devLanguage": [
        "string"
      ],
      "downloads": 0,
      "features": "string",
      "id": "string",
      "isDel": true,
      "logo": "string",
      "name": "string",
      "nameZh": "string",
      "notice": {
        "host": "string",
        "method": "string",
        "query": [
          null
        ],
        "uri": "string"
      },
      "poster": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "preview": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "relationSoft": [
        {
          "category": [
            "string"
          ],
          "config": null,
          "copyright": "string",
          "createAt": "string",
          "devLanguage": [
            "string"
          ],
          "downloads": 0,
          "features": "string",
          "id": "string",
          "isDel": true,
          "logo": "string",
          "name": "string",
          "nameZh": "string",
          "notice": null,
          "poster": [
            {}
          ],
          "preview": [
            {}
          ],
          "relationSoft": [
            {}
          ],
          "relationSoftIds": [
            "string"
          ],
          "score": 0,
          "state": 0,
          "supportMode": null,
          "sysLanguage": [
            "string"
          ],
          "team": "string",
          "teamIcon": "string",
          "updateAt": "string",
          "userId": "string"
        }
      ],
      "relationSoftIds": [
        "string"
      ],
      "score": 0,
      "state": 0,
      "supportMode": {
        "docker": true,
        "swarm": true
      },
      "sysLanguage": [
        "string"
      ],
      "team": "string",
      "teamIcon": "string",
      "updateAt": "string",
      "userId": "string"
    }
  ],
  "type": 0,
  "uniqueId": "string",
  "updateAt": "string",
  "userId": "string",
  "username": "string"
}
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[docker.DocEndpoint](#schemadocker.docendpoint)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[response.Response](#schemaresponse.response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[response.Response](#schemaresponse.response)|

## POST Testing the environment for local storage

POST /v2/endpoint/test

Testing the environment for local storage

> Body Request parameters

```json
{
  "TLSConfig": {
    "TLSCACertPath": "string",
    "TLSCertPath": "string",
    "TLSKeyPath": "string",
    "TLSType": 1
  },
  "connected": true,
  "createAt": "string",
  "envURL": "string",
  "id": "string",
  "isSocket": true,
  "isTLS": true,
  "k8sConfig": "string",
  "name": "string",
  "password": "string",
  "securitySettings": {
    "allowBindMountsForRegularUsers": false,
    "allowContainerCapabilitiesForRegularUsers": true,
    "allowDeviceMappingForRegularUsers": true,
    "allowHostNamespaceForRegularUsers": true,
    "allowPrivilegedModeForRegularUsers": false,
    "allowStackManagementForRegularUsers": true,
    "allowSysctlSettingForRegularUsers": true,
    "allowVolumeBrowserForRegularUsers": true,
    "enableHostManagementFeatures": true
  },
  "status": 1,
  "templates": [
    {
      "category": [
        "string"
      ],
      "config": {
        "createAt": "string",
        "dockerConfig": {},
        "dockerImage": "string",
        "dockerRegistry": null,
        "fileSize": 0,
        "id": "string",
        "isDel": true,
        "k8SConfig": null,
        "k8SContent": "string",
        "k8SRegistry": [
          null
        ],
        "manual": null,
        "platform": 0,
        "swarmConfig": {},
        "swarmContent": "string",
        "swarmRegistry": [
          null
        ],
        "templateId": "string",
        "updateAt": "string",
        "userId": "string",
        "version": "string",
        "versionContent": "string"
      },
      "copyright": "string",
      "createAt": "string",
      "devLanguage": [
        "string"
      ],
      "downloads": 0,
      "features": "string",
      "id": "string",
      "isDel": true,
      "logo": "string",
      "name": "string",
      "nameZh": "string",
      "notice": {
        "host": "string",
        "method": "string",
        "query": [
          null
        ],
        "uri": "string"
      },
      "poster": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "preview": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "relationSoft": [
        {
          "category": [
            "string"
          ],
          "config": null,
          "copyright": "string",
          "createAt": "string",
          "devLanguage": [
            "string"
          ],
          "downloads": 0,
          "features": "string",
          "id": "string",
          "isDel": true,
          "logo": "string",
          "name": "string",
          "nameZh": "string",
          "notice": null,
          "poster": [
            {}
          ],
          "preview": [
            {}
          ],
          "relationSoft": [
            {}
          ],
          "relationSoftIds": [
            "string"
          ],
          "score": 0,
          "state": 0,
          "supportMode": null,
          "sysLanguage": [
            "string"
          ],
          "team": "string",
          "teamIcon": "string",
          "updateAt": "string",
          "userId": "string"
        }
      ],
      "relationSoftIds": [
        "string"
      ],
      "score": 0,
      "state": 0,
      "supportMode": {
        "docker": true,
        "swarm": true
      },
      "sysLanguage": [
        "string"
      ],
      "team": "string",
      "teamIcon": "string",
      "updateAt": "string",
      "userId": "string"
    }
  ],
  "type": 0,
  "uniqueId": "string",
  "updateAt": "string",
  "userId": "string",
  "username": "string"
}
```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|[docker.DocEndpoint](#schemadocker.docendpoint)| no |none|

> Return Example

> 200 Response

```json
{
  "TLSConfig": {
    "TLSCACertPath": "string",
    "TLSCertPath": "string",
    "TLSKeyPath": "string",
    "TLSType": 1
  },
  "connected": true,
  "createAt": "string",
  "envURL": "string",
  "id": "string",
  "isSocket": true,
  "isTLS": true,
  "k8sConfig": "string",
  "name": "string",
  "password": "string",
  "securitySettings": {
    "allowBindMountsForRegularUsers": false,
    "allowContainerCapabilitiesForRegularUsers": true,
    "allowDeviceMappingForRegularUsers": true,
    "allowHostNamespaceForRegularUsers": true,
    "allowPrivilegedModeForRegularUsers": false,
    "allowStackManagementForRegularUsers": true,
    "allowSysctlSettingForRegularUsers": true,
    "allowVolumeBrowserForRegularUsers": true,
    "enableHostManagementFeatures": true
  },
  "status": 1,
  "templates": [
    {
      "category": [
        "string"
      ],
      "config": {
        "createAt": "string",
        "dockerConfig": {},
        "dockerImage": "string",
        "dockerRegistry": null,
        "fileSize": 0,
        "id": "string",
        "isDel": true,
        "k8SConfig": null,
        "k8SContent": "string",
        "k8SRegistry": [
          null
        ],
        "manual": null,
        "platform": 0,
        "swarmConfig": {},
        "swarmContent": "string",
        "swarmRegistry": [
          null
        ],
        "templateId": "string",
        "updateAt": "string",
        "userId": "string",
        "version": "string",
        "versionContent": "string"
      },
      "copyright": "string",
      "createAt": "string",
      "devLanguage": [
        "string"
      ],
      "downloads": 0,
      "features": "string",
      "id": "string",
      "isDel": true,
      "logo": "string",
      "name": "string",
      "nameZh": "string",
      "notice": {
        "host": "string",
        "method": "string",
        "query": [
          null
        ],
        "uri": "string"
      },
      "poster": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "preview": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "relationSoft": [
        {
          "category": [
            "string"
          ],
          "config": null,
          "copyright": "string",
          "createAt": "string",
          "devLanguage": [
            "string"
          ],
          "downloads": 0,
          "features": "string",
          "id": "string",
          "isDel": true,
          "logo": "string",
          "name": "string",
          "nameZh": "string",
          "notice": null,
          "poster": [
            {}
          ],
          "preview": [
            {}
          ],
          "relationSoft": [
            {}
          ],
          "relationSoftIds": [
            "string"
          ],
          "score": 0,
          "state": 0,
          "supportMode": null,
          "sysLanguage": [
            "string"
          ],
          "team": "string",
          "teamIcon": "string",
          "updateAt": "string",
          "userId": "string"
        }
      ],
      "relationSoftIds": [
        "string"
      ],
      "score": 0,
      "state": 0,
      "supportMode": {
        "docker": true,
        "swarm": true
      },
      "sysLanguage": [
        "string"
      ],
      "team": "string",
      "teamIcon": "string",
      "updateAt": "string",
      "userId": "string"
    }
  ],
  "type": 0,
  "uniqueId": "string",
  "updateAt": "string",
  "userId": "string",
  "username": "string"
}
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[docker.DocEndpoint](#schemadocker.docendpoint)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[response.Response](#schemaresponse.response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[response.Response](#schemaresponse.response)|

# kubernetes

## GET k8sPodjournal

GET /v2/kubernetes/pod/log

k8sPodjournal

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|namespace|query|string| no |namespace|
|podName|query|string| no |podName|
|containerName|query|string| no |containerName|
|tailLines|query|string| no |tailLines|

> Return Example

> 200 Response

```json
{}
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|Inline|

### Return Data Structure

Status code **200**

*[Unified backend return results]<[Implementation of Excel T() function
<p/>
If the argument is a text or error value it is returned unmodified.  All other argument types
cause an empty string result.  If the argument is an area, the first (top-left) cell is used
(regardless of the coordinates of the evaluating formula cell).]>*

|name|name|name|name|name|name|
|---|---|---|---|---|---|

## GET k8sPodterminal

GET /v2/kubernetes/pod/terminal

k8sPodterminal

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|namespace|query|string| no |namespace|
|podName|query|string| no |podName|
|containerName|query|string| no |containerName|

> Return Example

> 200 Response

```json
{}
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|Inline|

### Return Data Structure

Status code **200**

*[Unified backend return results]<[Implementation of Excel T() function
<p/>
If the argument is a text or error value it is returned unmodified.  All other argument types
cause an empty string result.  If the argument is an area, the first (top-left) cell is used
(regardless of the coordinates of the evaluating formula cell).]>*

|name|name|name|name|name|name|
|---|---|---|---|---|---|

## GET obtaink8sPodobtain

GET /v2/kubernetes/pods

obtaink8sPodobtain

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|namespace|query|string| no |namespace|

> Return Example

> 200 Response

```json
{}
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|Inline|

### Return Data Structure

Status code **200**

*[Unified backend return results]<[Implementation of Excel T() function
<p/>
If the argument is a text or error value it is returned unmodified.  All other argument types
cause an empty string result.  If the argument is an area, the first (top-left) cell is used
(regardless of the coordinates of the evaluating formula cell).]>*

|name|name|name|name|name|name|
|---|---|---|---|---|---|

# SoftwareApi

## POST Create softwareApi

POST /v2/software

> Body Request parameters

```json
{
  "createAt": "string",
  "dockerYml": {
    "networks": {
      "property1": {
        "external": true
      },
      "property2": {
        "external": true
      }
    },
    "secrets": {
      "property1": {
        "external": "string",
        "file": "string",
        "name": "string"
      },
      "property2": {
        "external": "string",
        "file": "string",
        "name": "string"
      }
    },
    "services": {
      "property1": {
        "build": {},
        "capAdd": [
          "string"
        ],
        "capDrop": [
          "string"
        ],
        "cgroupParent": "string",
        "command": [
          "string"
        ],
        "containerName": "string",
        "dependsOn": [
          "string"
        ],
        "deploy": {},
        "domainName": "string",
        "envFile": [
          "string"
        ],
        "environment": [
          {
            "desc": null,
            "key": null,
            "val": null
          }
        ],
        "hostName": "string",
        "image": {},
        "ipc": "string",
        "macAddress": "string",
        "networkMode": "string",
        "networks": {
          "property1": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          },
          "property2": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          }
        },
        "ports": [
          {
            "desc": null,
            "mode": null,
            "protocol": null,
            "published": null,
            "target": null
          }
        ],
        "privileged": true,
        "proflies": [
          "string"
        ],
        "readOnly": true,
        "restart": "string",
        "secrets": [
          {
            "gid": null,
            "mode": null,
            "source": null,
            "target": null,
            "uid": null
          }
        ],
        "shmSize": "string",
        "stdinOpen": true,
        "tty": true,
        "user": "string",
        "volumes": [
          {
            "desc": null,
            "readOnly": null,
            "source": null,
            "target": null,
            "type": null
          }
        ],
        "workingDir": "string"
      },
      "property2": {
        "build": {},
        "capAdd": [
          "string"
        ],
        "capDrop": [
          "string"
        ],
        "cgroupParent": "string",
        "command": [
          "string"
        ],
        "containerName": "string",
        "dependsOn": [
          "string"
        ],
        "deploy": {},
        "domainName": "string",
        "envFile": [
          "string"
        ],
        "environment": [
          {
            "desc": null,
            "key": null,
            "val": null
          }
        ],
        "hostName": "string",
        "image": {},
        "ipc": "string",
        "macAddress": "string",
        "networkMode": "string",
        "networks": {
          "property1": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          },
          "property2": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          }
        },
        "ports": [
          {
            "desc": null,
            "mode": null,
            "protocol": null,
            "published": null,
            "target": null
          }
        ],
        "privileged": true,
        "proflies": [
          "string"
        ],
        "readOnly": true,
        "restart": "string",
        "secrets": [
          {
            "gid": null,
            "mode": null,
            "source": null,
            "target": null,
            "uid": null
          }
        ],
        "shmSize": "string",
        "stdinOpen": true,
        "tty": true,
        "user": "string",
        "volumes": [
          {
            "desc": null,
            "readOnly": null,
            "source": null,
            "target": null,
            "type": null
          }
        ],
        "workingDir": "string"
      }
    },
    "version": "string",
    "volumes": {
      "property1": {
        "external": true
      },
      "property2": {
        "external": true
      }
    }
  },
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": null,
      "TLSCertPath": null,
      "TLSKeyPath": null,
      "TLSType": null
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": null,
      "allowContainerCapabilitiesForRegularUsers": null,
      "allowDeviceMappingForRegularUsers": null,
      "allowHostNamespaceForRegularUsers": null,
      "allowPrivilegedModeForRegularUsers": null,
      "allowStackManagementForRegularUsers": null,
      "allowSysctlSettingForRegularUsers": null,
      "allowVolumeBrowserForRegularUsers": null,
      "enableHostManagementFeatures": null
    },
    "status": 1,
    "templates": [
      {
        "category": [
          null
        ],
        "config": null,
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          null
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": null,
        "poster": [
          null
        ],
        "preview": [
          null
        ],
        "relationSoft": [
          null
        ],
        "relationSoftIds": [
          null
        ],
        "score": 0,
        "state": 0,
        "supportMode": null,
        "sysLanguage": [
          null
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "entryPoint": "string",
  "env": [
    {
      "name": "name",
      "value": "value"
    }
  ],
  "homeHost": "string",
  "homePort": "string",
  "id": "string",
  "name": "string",
  "projectPath": "string",
  "status": "string",
  "template": {
    "category": [
      "string"
    ],
    "config": {
      "createAt": null,
      "dockerConfig": null,
      "dockerImage": null,
      "dockerRegistry": null,
      "fileSize": null,
      "id": null,
      "isDel": null,
      "k8SConfig": null,
      "k8SContent": null,
      "k8SRegistry": null,
      "manual": null,
      "platform": null,
      "swarmConfig": null,
      "swarmContent": null,
      "swarmRegistry": null,
      "templateId": null,
      "updateAt": null,
      "userId": null,
      "version": null,
      "versionContent": null
    },
    "copyright": "string",
    "createAt": "string",
    "devLanguage": [
      "string"
    ],
    "downloads": 0,
    "features": "string",
    "id": "string",
    "isDel": true,
    "logo": "string",
    "name": "string",
    "nameZh": "string",
    "notice": {
      "host": null,
      "method": null,
      "query": null,
      "uri": null
    },
    "poster": [
      {
        "createAt": "string",
        "id": "string",
        "name": "string",
        "originalName": "string",
        "size": 0,
        "suffix": "string",
        "updateAt": "string",
        "url": "string"
      }
    ],
    "preview": [
      {
        "createAt": "string",
        "id": "string",
        "name": "string",
        "originalName": "string",
        "size": 0,
        "suffix": "string",
        "updateAt": "string",
        "url": "string"
      }
    ],
    "relationSoft": [
      {
        "category": [
          null
        ],
        "config": null,
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          null
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": null,
        "poster": [
          null
        ],
        "preview": [
          null
        ],
        "relationSoft": [
          null
        ],
        "relationSoftIds": [
          null
        ],
        "score": 0,
        "state": 0,
        "supportMode": null,
        "sysLanguage": [
          null
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "relationSoftIds": [
      "string"
    ],
    "score": 0,
    "state": 0,
    "supportMode": {
      "docker": null,
      "swarm": null
    },
    "sysLanguage": [
      "string"
    ],
    "team": "string",
    "teamIcon": "string",
    "updateAt": "string",
    "userId": "string"
  },
  "updateAt": "string",
  "userId": "string",
  "userIp": "string",
  "userIp2Region": "string"
}
```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|[docker.DocSoftware](#schemadocker.docsoftware)| no |none|

> Return Example

> 200 Response

```json
"string"
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|{"code":0,"data":{},"msg":"Created successfully"}|string|

## DELETE Remove softwareApi

DELETE /v2/software

> Body Request parameters

```json
{
  "id": "string"
}
```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|[request.DelSoftwareReq](#schemarequest.delsoftwarereq)| no |none|

> Return Example

> 200 Response

```json
"string"
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|{"code":0,"data":{},"msg":"Successfully deleted"}|string|

## POST executeApi

POST /v2/software/start

> Body Request parameters

```json
{
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": "string",
      "TLSCertPath": "string",
      "TLSKeyPath": "string",
      "TLSType": 1
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": false,
      "allowContainerCapabilitiesForRegularUsers": true,
      "allowDeviceMappingForRegularUsers": true,
      "allowHostNamespaceForRegularUsers": true,
      "allowPrivilegedModeForRegularUsers": false,
      "allowStackManagementForRegularUsers": true,
      "allowSysctlSettingForRegularUsers": true,
      "allowVolumeBrowserForRegularUsers": true,
      "enableHostManagementFeatures": true
    },
    "status": 1,
    "templates": [
      {
        "category": [
          "string"
        ],
        "config": {},
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          "string"
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": {},
        "poster": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "preview": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "relationSoft": [
          {
            "category": null,
            "config": null,
            "copyright": null,
            "createAt": null,
            "devLanguage": null,
            "downloads": null,
            "features": null,
            "id": null,
            "isDel": null,
            "logo": null,
            "name": null,
            "nameZh": null,
            "notice": null,
            "poster": null,
            "preview": null,
            "relationSoft": null,
            "relationSoftIds": null,
            "score": null,
            "state": null,
            "supportMode": null,
            "sysLanguage": null,
            "team": null,
            "teamIcon": null,
            "updateAt": null,
            "userId": null
          }
        ],
        "relationSoftIds": [
          "string"
        ],
        "score": 0,
        "state": 0,
        "supportMode": {},
        "sysLanguage": [
          "string"
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "endpointId": "string",
  "name": "string",
  "projectPath": "string"
}
```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|[request.StartSoftwareRequest](#schemarequest.startsoftwarerequest)| no |none|

> Return Example

> 200 Response

```json
"string"
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|{"code":0,"data":{},"msg":"Successfully started"}|string|

## POST Stop softwareApi

POST /v2/software/stop

> Body Request parameters

```json
{
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": "string",
      "TLSCertPath": "string",
      "TLSKeyPath": "string",
      "TLSType": 1
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": false,
      "allowContainerCapabilitiesForRegularUsers": true,
      "allowDeviceMappingForRegularUsers": true,
      "allowHostNamespaceForRegularUsers": true,
      "allowPrivilegedModeForRegularUsers": false,
      "allowStackManagementForRegularUsers": true,
      "allowSysctlSettingForRegularUsers": true,
      "allowVolumeBrowserForRegularUsers": true,
      "enableHostManagementFeatures": true
    },
    "status": 1,
    "templates": [
      {
        "category": [
          "string"
        ],
        "config": {},
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          "string"
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": {},
        "poster": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "preview": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "relationSoft": [
          {
            "category": null,
            "config": null,
            "copyright": null,
            "createAt": null,
            "devLanguage": null,
            "downloads": null,
            "features": null,
            "id": null,
            "isDel": null,
            "logo": null,
            "name": null,
            "nameZh": null,
            "notice": null,
            "poster": null,
            "preview": null,
            "relationSoft": null,
            "relationSoftIds": null,
            "score": null,
            "state": null,
            "supportMode": null,
            "sysLanguage": null,
            "team": null,
            "teamIcon": null,
            "updateAt": null,
            "userId": null
          }
        ],
        "relationSoftIds": [
          "string"
        ],
        "score": 0,
        "state": 0,
        "supportMode": {},
        "sysLanguage": [
          "string"
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "endpointId": "string",
  "name": "string",
  "projectPath": "string"
}
```

### Request parameters

|name|name|name|name|name|
|---|---|---|---|---|
|body|body|[request.StopSoftwareRequest](#schemarequest.stopsoftwarerequest)| no |none|

> Return Example

> 200 Response

```json
"string"
```

### Return Results

|Status code|Status code|Status code|Status code|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|{"code":0,"data":{},"msg":"Stopped successfully"}|string|

# data model

<h2 id="tocS_v1.WindowsSecurityContextOptions">v1.WindowsSecurityContextOptions</h2>

<a id="schemav1.windowssecuritycontextoptions"></a>
<a id="schema_v1.WindowsSecurityContextOptions"></a>
<a id="tocSv1.windowssecuritycontextoptions"></a>
<a id="tocsv1.windowssecuritycontextoptions"></a>

```json
{
  "gmsaCredentialSpec": "string",
  "gmsaCredentialSpecName": "string",
  "hostProcess": true,
  "runAsUserName": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|gmsaCredentialSpec|string|false|none||GMSACredentialSpec is where the GMSA admission webhook<br />(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the<br />GMSA credential spec named by the GMSACredentialSpecName field.<br />+optional|
|gmsaCredentialSpecName|string|false|none||GMSACredentialSpecName is the name of the GMSA credential spec to use.<br />+optional|
|hostProcess|boolean|false|none||HostProcess determines if a container should be run as a 'Host Process' container.<br />This field is alpha-level and will only be honored by components that enable the<br />WindowsHostProcessContainers feature flag. Setting this field without the feature<br />flag will result in errors when validating the Pod. All of a Pod's containers must<br />have the same effective HostProcess value (it is not allowed to have a mix of HostProcess<br />containers and non-HostProcess containers).  In addition, if HostProcess is true<br />then HostNetwork must also be set to true.<br />+optional|
|runAsUserName|string|false|none||The UserName in Windows to run the entrypoint of the container process.<br />Defaults to the user specified in image metadata if unspecified.<br />May also be set in PodSecurityContext. If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|

<h2 id="tocS_v1.WeightedPodAffinityTerm">v1.WeightedPodAffinityTerm</h2>

<a id="schemav1.weightedpodaffinityterm"></a>
<a id="schema_v1.WeightedPodAffinityTerm"></a>
<a id="tocSv1.weightedpodaffinityterm"></a>
<a id="tocsv1.weightedpodaffinityterm"></a>

```json
{
  "podAffinityTerm": {
    "labelSelector": {
      "matchExpressions": null,
      "matchLabels": null
    },
    "namespaceSelector": {
      "matchExpressions": null,
      "matchLabels": null
    },
    "namespaces": [
      "string"
    ],
    "topologyKey": "string"
  },
  "weight": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|podAffinityTerm|[v1.PodAffinityTerm](#schemav1.podaffinityterm)|false|none||Required. A pod affinity term, associated with the corresponding weight.|
|weight|integer|false|none||weight associated with matching the corresponding podAffinityTerm,<br />in the range 1-100.|

<h2 id="tocS_v1.VsphereVirtualDiskVolumeSource">v1.VsphereVirtualDiskVolumeSource</h2>

<a id="schemav1.vspherevirtualdiskvolumesource"></a>
<a id="schema_v1.VsphereVirtualDiskVolumeSource"></a>
<a id="tocSv1.vspherevirtualdiskvolumesource"></a>
<a id="tocsv1.vspherevirtualdiskvolumesource"></a>

```json
{
  "fsType": "string",
  "storagePolicyID": "string",
  "storagePolicyName": "string",
  "volumePath": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />+optional|
|storagePolicyID|string|false|none||Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.<br />+optional|
|storagePolicyName|string|false|none||Storage Policy Based Management (SPBM) profile name.<br />+optional|
|volumePath|string|false|none||Path that identifies vSphere volume vmdk|

<h2 id="tocS_v1.VolumeProjection">v1.VolumeProjection</h2>

<a id="schemav1.volumeprojection"></a>
<a id="schema_v1.VolumeProjection"></a>
<a id="tocSv1.volumeprojection"></a>
<a id="tocsv1.volumeprojection"></a>

```json
{
  "configMap": {
    "items": [
      {
        "key": "string",
        "mode": 0,
        "path": "string"
      }
    ],
    "name": "string",
    "optional": true
  },
  "downwardAPI": {
    "items": [
      {
        "fieldRef": null,
        "mode": 0,
        "path": "string",
        "resourceFieldRef": null
      }
    ]
  },
  "secret": {
    "items": [
      {
        "key": "string",
        "mode": 0,
        "path": "string"
      }
    ],
    "name": "string",
    "optional": true
  },
  "serviceAccountToken": {
    "audience": "string",
    "expirationSeconds": 0,
    "path": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|configMap|[v1.ConfigMapProjection](#schemav1.configmapprojection)|false|none||information about the configMap data to project<br />+optional|
|downwardAPI|[v1.DownwardAPIProjection](#schemav1.downwardapiprojection)|false|none||information about the downwardAPI data to project<br />+optional|
|secret|[v1.SecretProjection](#schemav1.secretprojection)|false|none||information about the secret data to project<br />+optional|
|serviceAccountToken|[v1.ServiceAccountTokenProjection](#schemav1.serviceaccounttokenprojection)|false|none||information about the serviceAccountToken data to project<br />+optional|

<h2 id="tocS_v1.VolumeMount">v1.VolumeMount</h2>

<a id="schemav1.volumemount"></a>
<a id="schema_v1.VolumeMount"></a>
<a id="tocSv1.volumemount"></a>
<a id="tocsv1.volumemount"></a>

```json
{
  "mountPath": "string",
  "mountPropagation": "None",
  "name": "string",
  "readOnly": true,
  "subPath": "string",
  "subPathExpr": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|mountPath|string|false|none||Path within the container at which the volume should be mounted.  Must<br />not contain ':'.|
|mountPropagation|[v1.MountPropagationMode](#schemav1.mountpropagationmode)|false|none||mountPropagation determines how mounts are propagated from the host<br />to container and the other way around.<br />When not set, MountPropagationNone is used.<br />This field is beta in 1.10.<br />+optional|
|name|string|false|none||This must match the Name of a Volume.|
|readOnly|boolean|false|none||Mounted read-only if true, read-write otherwise (false or unspecified).<br />Defaults to false.<br />+optional|
|subPath|string|false|none||Path within the volume from which the container's volume should be mounted.<br />Defaults to "" (volume's root).<br />+optional|
|subPathExpr|string|false|none||Expanded path within the volume from which the container's volume should be mounted.<br />Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.<br />Defaults to "" (volume's root).<br />SubPathExpr and SubPath are mutually exclusive.<br />+optional|

<h2 id="tocS_v1.VolumeDevice">v1.VolumeDevice</h2>

<a id="schemav1.volumedevice"></a>
<a id="schema_v1.VolumeDevice"></a>
<a id="tocSv1.volumedevice"></a>
<a id="tocsv1.volumedevice"></a>

```json
{
  "devicePath": "string",
  "name": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|devicePath|string|false|none||devicePath is the path inside of the container that the device will be mapped to.|
|name|string|false|none||name must match the name of a persistentVolumeClaim in the pod|

<h2 id="tocS_v1.Volume">v1.Volume</h2>

<a id="schemav1.volume"></a>
<a id="schema_v1.Volume"></a>
<a id="tocSv1.volume"></a>
<a id="tocsv1.volume"></a>

```json
{
  "awsElasticBlockStore": {
    "fsType": "string",
    "partition": 0,
    "readOnly": true,
    "volumeID": "string"
  },
  "azureDisk": {
    "cachingMode": "None",
    "diskName": "string",
    "diskURI": "string",
    "fsType": "string",
    "kind": "Shared",
    "readOnly": true
  },
  "azureFile": {
    "readOnly": true,
    "secretName": "string",
    "shareName": "string"
  },
  "cephfs": {
    "monitors": [
      "string"
    ],
    "path": "string",
    "readOnly": true,
    "secretFile": "string",
    "secretRef": {
      "name": null
    },
    "user": "string"
  },
  "cinder": {
    "fsType": "string",
    "readOnly": true,
    "secretRef": {
      "name": null
    },
    "volumeID": "string"
  },
  "configMap": {
    "defaultMode": 0,
    "items": [
      {
        "key": "string",
        "mode": 0,
        "path": "string"
      }
    ],
    "name": "string",
    "optional": true
  },
  "csi": {
    "driver": "string",
    "fsType": "string",
    "nodePublishSecretRef": {
      "name": null
    },
    "readOnly": true,
    "volumeAttributes": {
      "property1": "string",
      "property2": "string"
    }
  },
  "downwardAPI": {
    "defaultMode": 0,
    "items": [
      {
        "fieldRef": null,
        "mode": 0,
        "path": "string",
        "resourceFieldRef": null
      }
    ]
  },
  "emptyDir": {
    "medium": "",
    "sizeLimit": {
      "Format": null
    }
  },
  "ephemeral": {
    "volumeClaimTemplate": {
      "metadata": null,
      "spec": null
    }
  },
  "fc": {
    "fsType": "string",
    "lun": 0,
    "readOnly": true,
    "targetWWNs": [
      "string"
    ],
    "wwids": [
      "string"
    ]
  },
  "flexVolume": {
    "driver": "string",
    "fsType": "string",
    "options": {
      "property1": "string",
      "property2": "string"
    },
    "readOnly": true,
    "secretRef": {
      "name": null
    }
  },
  "flocker": {
    "datasetName": "string",
    "datasetUUID": "string"
  },
  "gcePersistentDisk": {
    "fsType": "string",
    "partition": 0,
    "pdName": "string",
    "readOnly": true
  },
  "gitRepo": {
    "directory": "string",
    "repository": "string",
    "revision": "string"
  },
  "glusterfs": {
    "endpoints": "string",
    "path": "string",
    "readOnly": true
  },
  "hostPath": {
    "path": "string",
    "type": ""
  },
  "iscsi": {
    "chapAuthDiscovery": true,
    "chapAuthSession": true,
    "fsType": "string",
    "initiatorName": "string",
    "iqn": "string",
    "iscsiInterface": "string",
    "lun": 0,
    "portals": [
      "string"
    ],
    "readOnly": true,
    "secretRef": {
      "name": null
    },
    "targetPortal": "string"
  },
  "name": "string",
  "nfs": {
    "path": "string",
    "readOnly": true,
    "server": "string"
  },
  "persistentVolumeClaim": {
    "claimName": "string",
    "readOnly": true
  },
  "photonPersistentDisk": {
    "fsType": "string",
    "pdID": "string"
  },
  "portworxVolume": {
    "fsType": "string",
    "readOnly": true,
    "volumeID": "string"
  },
  "projected": {
    "defaultMode": 0,
    "sources": [
      {
        "configMap": null,
        "downwardAPI": null,
        "secret": null,
        "serviceAccountToken": null
      }
    ]
  },
  "quobyte": {
    "group": "string",
    "readOnly": true,
    "registry": "string",
    "tenant": "string",
    "user": "string",
    "volume": "string"
  },
  "rbd": {
    "fsType": "string",
    "image": "string",
    "keyring": "string",
    "monitors": [
      "string"
    ],
    "pool": "string",
    "readOnly": true,
    "secretRef": {
      "name": null
    },
    "user": "string"
  },
  "scaleIO": {
    "fsType": "string",
    "gateway": "string",
    "protectionDomain": "string",
    "readOnly": true,
    "secretRef": {
      "name": null
    },
    "sslEnabled": true,
    "storageMode": "string",
    "storagePool": "string",
    "system": "string",
    "volumeName": "string"
  },
  "secret": {
    "defaultMode": 0,
    "items": [
      {
        "key": "string",
        "mode": 0,
        "path": "string"
      }
    ],
    "optional": true,
    "secretName": "string"
  },
  "storageos": {
    "fsType": "string",
    "readOnly": true,
    "secretRef": {
      "name": null
    },
    "volumeName": "string",
    "volumeNamespace": "string"
  },
  "vsphereVolume": {
    "fsType": "string",
    "storagePolicyID": "string",
    "storagePolicyName": "string",
    "volumePath": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|awsElasticBlockStore|[v1.AWSElasticBlockStoreVolumeSource](#schemav1.awselasticblockstorevolumesource)|false|none||AWSElasticBlockStore represents an AWS Disk resource that is attached to a<br />kubelet's host machine and then exposed to the pod.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore<br />+optional|
|azureDisk|[v1.AzureDiskVolumeSource](#schemav1.azurediskvolumesource)|false|none||AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br />+optional|
|azureFile|[v1.AzureFileVolumeSource](#schemav1.azurefilevolumesource)|false|none||AzureFile represents an Azure File Service mount on the host and bind mount to the pod.<br />+optional|
|cephfs|[v1.CephFSVolumeSource](#schemav1.cephfsvolumesource)|false|none||CephFS represents a Ceph FS mount on the host that shares a pod's lifetime<br />+optional|
|cinder|[v1.CinderVolumeSource](#schemav1.cindervolumesource)|false|none||Cinder represents a cinder volume attached and mounted on kubelets host machine.<br />More info: https://examples.k8s.io/mysql-cinder-pd/README.md<br />+optional|
|configMap|[v1.ConfigMapVolumeSource](#schemav1.configmapvolumesource)|false|none||ConfigMap represents a configMap that should populate this volume<br />+optional|
|csi|[v1.CSIVolumeSource](#schemav1.csivolumesource)|false|none||CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).<br />+optional|
|downwardAPI|[v1.DownwardAPIVolumeSource](#schemav1.downwardapivolumesource)|false|none||DownwardAPI represents downward API about the pod that should populate this volume<br />+optional|
|emptyDir|[v1.EmptyDirVolumeSource](#schemav1.emptydirvolumesource)|false|none||EmptyDir represents a temporary directory that shares a pod's lifetime.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir<br />+optional|
|ephemeral|[v1.EphemeralVolumeSource](#schemav1.ephemeralvolumesource)|false|none||Ephemeral represents a volume that is handled by a cluster storage driver.<br />The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br />and deleted when the pod is removed.<br /><br />Use this if:<br />a) the volume is only needed while the pod runs,<br />b) features of normal volumes like restoring from snapshot or capacity<br />   tracking are needed,<br />c) the storage driver is specified through a storage class, and<br />d) the storage driver supports dynamic volume provisioning through<br />   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br />   information on the connection between this volume type<br />   and PersistentVolumeClaim).<br /><br />Use PersistentVolumeClaim or one of the vendor-specific<br />APIs for volumes that persist for longer than the lifecycle<br />of an individual pod.<br /><br />Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br />be used that way - see the documentation of the driver for<br />more information.<br /><br />A pod can use both types of ephemeral volumes and<br />persistent volumes at the same time.<br /><br />This is a beta feature and only available when the GenericEphemeralVolume<br />feature gate is enabled.<br /><br />+optional|
|fc|[v1.FCVolumeSource](#schemav1.fcvolumesource)|false|none||FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.<br />+optional|
|flexVolume|[v1.FlexVolumeSource](#schemav1.flexvolumesource)|false|none||FlexVolume represents a generic volume resource that is<br />provisioned/attached using an exec based plugin.<br />+optional|
|flocker|[v1.FlockerVolumeSource](#schemav1.flockervolumesource)|false|none||Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running<br />+optional|
|gcePersistentDisk|[v1.GCEPersistentDiskVolumeSource](#schemav1.gcepersistentdiskvolumesource)|false|none||GCEPersistentDisk represents a GCE Disk resource that is attached to a<br />kubelet's host machine and then exposed to the pod.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br />+optional|
|gitRepo|[v1.GitRepoVolumeSource](#schemav1.gitrepovolumesource)|false|none||GitRepo represents a git repository at a particular revision.<br />DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an<br />EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br />into the Pod's container.<br />+optional|
|glusterfs|[v1.GlusterfsVolumeSource](#schemav1.glusterfsvolumesource)|false|none||Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br />More info: https://examples.k8s.io/volumes/glusterfs/README.md<br />+optional|
|hostPath|[v1.HostPathVolumeSource](#schemav1.hostpathvolumesource)|false|none||HostPath represents a pre-existing file or directory on the host<br />machine that is directly exposed to the container. This is generally<br />used for system agents or other privileged things that are allowed<br />to see the host machine. Most containers will NOT need this.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath<br />---<br />TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not<br />mount host directories as read/write.<br />+optional|
|iscsi|[v1.ISCSIVolumeSource](#schemav1.iscsivolumesource)|false|none||ISCSI represents an ISCSI Disk resource that is attached to a<br />kubelet's host machine and then exposed to the pod.<br />More info: https://examples.k8s.io/volumes/iscsi/README.md<br />+optional|
|name|string|false|none||Volume's name.<br />Must be a DNS_LABEL and unique within the pod.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names|
|nfs|[v1.NFSVolumeSource](#schemav1.nfsvolumesource)|false|none||NFS represents an NFS mount on the host that shares a pod's lifetime<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs<br />+optional|
|persistentVolumeClaim|[v1.PersistentVolumeClaimVolumeSource](#schemav1.persistentvolumeclaimvolumesource)|false|none||PersistentVolumeClaimVolumeSource represents a reference to a<br />PersistentVolumeClaim in the same namespace.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims<br />+optional|
|photonPersistentDisk|[v1.PhotonPersistentDiskVolumeSource](#schemav1.photonpersistentdiskvolumesource)|false|none||PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine|
|portworxVolume|[v1.PortworxVolumeSource](#schemav1.portworxvolumesource)|false|none||PortworxVolume represents a portworx volume attached and mounted on kubelets host machine<br />+optional|
|projected|[v1.ProjectedVolumeSource](#schemav1.projectedvolumesource)|false|none||Items for all in one resources secrets, configmaps, and downward API|
|quobyte|[v1.QuobyteVolumeSource](#schemav1.quobytevolumesource)|false|none||Quobyte represents a Quobyte mount on the host that shares a pod's lifetime<br />+optional|
|rbd|[v1.RBDVolumeSource](#schemav1.rbdvolumesource)|false|none||RBD represents a Rados Block Device mount on the host that shares a pod's lifetime.<br />More info: https://examples.k8s.io/volumes/rbd/README.md<br />+optional|
|scaleIO|[v1.ScaleIOVolumeSource](#schemav1.scaleiovolumesource)|false|none||ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br />+optional|
|secret|[v1.SecretVolumeSource](#schemav1.secretvolumesource)|false|none||Secret represents a secret that should populate this volume.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#secret<br />+optional|
|storageos|[v1.StorageOSVolumeSource](#schemav1.storageosvolumesource)|false|none||StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br />+optional|
|vsphereVolume|[v1.VsphereVirtualDiskVolumeSource](#schemav1.vspherevirtualdiskvolumesource)|false|none||VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine<br />+optional|

<h2 id="tocS_v1.UnsatisfiableConstraintAction">v1.UnsatisfiableConstraintAction</h2>

<a id="schemav1.unsatisfiableconstraintaction"></a>
<a id="schema_v1.UnsatisfiableConstraintAction"></a>
<a id="tocSv1.unsatisfiableconstraintaction"></a>
<a id="tocsv1.unsatisfiableconstraintaction"></a>

```json
"DoNotSchedule"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|DoNotSchedule|
|*anonymous*|ScheduleAnyway|

<h2 id="tocS_v1.URIScheme">v1.URIScheme</h2>

<a id="schemav1.urischeme"></a>
<a id="schema_v1.URIScheme"></a>
<a id="tocSv1.urischeme"></a>
<a id="tocsv1.urischeme"></a>

```json
"HTTP"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|HTTP|
|*anonymous*|HTTPS|

<h2 id="tocS_v1.TypedLocalObjectReference">v1.TypedLocalObjectReference</h2>

<a id="schemav1.typedlocalobjectreference"></a>
<a id="schema_v1.TypedLocalObjectReference"></a>
<a id="tocSv1.typedlocalobjectreference"></a>
<a id="tocsv1.typedlocalobjectreference"></a>

```json
{
  "apiGroup": "string",
  "kind": "string",
  "name": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiGroup|string|false|none||APIGroup is the group for the resource being referenced.<br />If APIGroup is not specified, the specified Kind must be in the core API group.<br />For any other third-party types, APIGroup is required.<br />+optional|
|kind|string|false|none||Kind is the type of resource being referenced|
|name|string|false|none||Name is the name of resource being referenced|

<h2 id="tocS_v1.TopologySpreadConstraint">v1.TopologySpreadConstraint</h2>

<a id="schemav1.topologyspreadconstraint"></a>
<a id="schema_v1.TopologySpreadConstraint"></a>
<a id="tocSv1.topologyspreadconstraint"></a>
<a id="tocsv1.topologyspreadconstraint"></a>

```json
{
  "labelSelector": {
    "matchExpressions": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ],
    "matchLabels": {
      "property1": "string",
      "property2": "string"
    }
  },
  "maxSkew": 0,
  "topologyKey": "string",
  "whenUnsatisfiable": "DoNotSchedule"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|labelSelector|[v1.LabelSelector](#schemav1.labelselector)|false|none||LabelSelector is used to find matching pods.<br />Pods that match this label selector are counted to determine the number of pods<br />in their corresponding topology domain.<br />+optional|
|maxSkew|integer|false|none||MaxSkew describes the degree to which pods may be unevenly distributed.<br />When `whenUnsatisfiable=DoNotSchedule`, it is the maximum permitted difference<br />between the number of matching pods in the target topology and the global minimum.<br />For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same<br />labelSelector spread as 1/1/0:<br />+-------+-------+-------+<br />| zone1 | zone2 | zone3 |<br />+-------+-------+-------+<br />|   P   |   P   |       |<br />+-------+-------+-------+<br />- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1;<br />scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2)<br />violate MaxSkew(1).<br />- if MaxSkew is 2, incoming pod can be scheduled onto any zone.<br />When `whenUnsatisfiable=ScheduleAnyway`, it is used to give higher precedence<br />to topologies that satisfy it.<br />It's a required field. Default value is 1 and 0 is not allowed.|
|topologyKey|string|false|none||TopologyKey is the key of node labels. Nodes that have a label with this key<br />and identical values are considered to be in the same topology.<br />We consider each <key, value> as a "bucket", and try to put balanced number<br />of pods into each bucket.<br />It's a required field.|
|whenUnsatisfiable|[v1.UnsatisfiableConstraintAction](#schemav1.unsatisfiableconstraintaction)|false|none||WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy<br />the spread constraint.<br />- DoNotSchedule (default) tells the scheduler not to schedule it.<br />- ScheduleAnyway tells the scheduler to schedule the pod in any location,<br />  but giving higher precedence to topologies that would help reduce the<br />  skew.<br />A constraint is considered "Unsatisfiable" for an incoming pod<br />if and only if every possible node assigment for that pod would violate<br />"MaxSkew" on some topology.<br />For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same<br />labelSelector spread as 3/1/1:<br />+-------+-------+-------+<br />| zone1 | zone2 | zone3 |<br />+-------+-------+-------+<br />| P P P |   P   |   P   |<br />+-------+-------+-------+<br />If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled<br />to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies<br />MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler<br />won't make it *more* imbalanced.<br />It's a required field.|

<h2 id="tocS_v1.TolerationOperator">v1.TolerationOperator</h2>

<a id="schemav1.tolerationoperator"></a>
<a id="schema_v1.TolerationOperator"></a>
<a id="tocSv1.tolerationoperator"></a>
<a id="tocsv1.tolerationoperator"></a>

```json
"Exists"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Exists|
|*anonymous*|Equal|

<h2 id="tocS_v1.Toleration">v1.Toleration</h2>

<a id="schemav1.toleration"></a>
<a id="schema_v1.Toleration"></a>
<a id="tocSv1.toleration"></a>
<a id="tocsv1.toleration"></a>

```json
{
  "effect": "NoSchedule",
  "key": "string",
  "operator": "Exists",
  "tolerationSeconds": 0,
  "value": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|effect|[v1.TaintEffect](#schemav1.tainteffect)|false|none||Effect indicates the taint effect to match. Empty means match all taint effects.<br />When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.<br />+optional|
|key|string|false|none||Key is the taint key that the toleration applies to. Empty means match all taint keys.<br />If the key is empty, operator must be Exists; this combination means to match all values and all keys.<br />+optional|
|operator|[v1.TolerationOperator](#schemav1.tolerationoperator)|false|none||Operator represents a key's relationship to the value.<br />Valid operators are Exists and Equal. Defaults to Equal.<br />Exists is equivalent to wildcard for value, so that a pod can<br />tolerate all taints of a particular category.<br />+optional|
|tolerationSeconds|integer|false|none||TolerationSeconds represents the period of time the toleration (which must be<br />of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,<br />it is not set, which means tolerate the taint forever (do not evict). Zero and<br />negative values will be treated as 0 (evict immediately) by the system.<br />+optional|
|value|string|false|none||Value is the taint value the toleration matches to.<br />If the operator is Exists, the value should be empty, otherwise just a regular string.<br />+optional|

<h2 id="tocS_v1.TerminationMessagePolicy">v1.TerminationMessagePolicy</h2>

<a id="schemav1.terminationmessagepolicy"></a>
<a id="schema_v1.TerminationMessagePolicy"></a>
<a id="tocSv1.terminationmessagepolicy"></a>
<a id="tocsv1.terminationmessagepolicy"></a>

```json
"File"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|File|
|*anonymous*|FallbackToLogsOnError|

<h2 id="tocS_v1.TaintEffect">v1.TaintEffect</h2>

<a id="schemav1.tainteffect"></a>
<a id="schema_v1.TaintEffect"></a>
<a id="tocSv1.tainteffect"></a>
<a id="tocsv1.tainteffect"></a>

```json
"NoSchedule"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|NoSchedule|
|*anonymous*|PreferNoSchedule|
|*anonymous*|NoExecute|

<h2 id="tocS_v1.TCPSocketAction">v1.TCPSocketAction</h2>

<a id="schemav1.tcpsocketaction"></a>
<a id="schema_v1.TCPSocketAction"></a>
<a id="tocSv1.tcpsocketaction"></a>
<a id="tocsv1.tcpsocketaction"></a>

```json
{
  "host": "string",
  "port": {
    "intVal": 0,
    "strVal": "string",
    "type": 0
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|host|string|false|none||Optional: Host name to connect to, defaults to the pod IP.<br />+optional|
|port|[intstr.IntOrString](#schemaintstr.intorstring)|false|none||Number or name of the port to access on the container.<br />Number must be in the range 1 to 65535.<br />Name must be an IANA_SVC_NAME.|

<h2 id="tocS_v1.Sysctl">v1.Sysctl</h2>

<a id="schemav1.sysctl"></a>
<a id="schema_v1.Sysctl"></a>
<a id="tocSv1.sysctl"></a>
<a id="tocsv1.sysctl"></a>

```json
{
  "name": "string",
  "value": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Name of a property to set|
|value|string|false|none||Value of a property to set|

<h2 id="tocS_v1.StorageOSVolumeSource">v1.StorageOSVolumeSource</h2>

<a id="schemav1.storageosvolumesource"></a>
<a id="schema_v1.StorageOSVolumeSource"></a>
<a id="tocSv1.storageosvolumesource"></a>
<a id="tocsv1.storageosvolumesource"></a>

```json
{
  "fsType": "string",
  "readOnly": true,
  "secretRef": {
    "name": "string"
  },
  "volumeName": "string",
  "volumeNamespace": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />+optional|
|readOnly|boolean|false|none||Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||SecretRef specifies the secret to use for obtaining the StorageOS API<br />credentials.  If not specified, default values will be attempted.<br />+optional|
|volumeName|string|false|none||VolumeName is the human-readable name of the StorageOS volume.  Volume<br />names are only unique within a namespace.|
|volumeNamespace|string|false|none||VolumeNamespace specifies the scope of the volume within StorageOS.  If no<br />namespace is specified then the Pod's namespace will be used.  This allows the<br />Kubernetes name scoping to be mirrored within StorageOS for tighter integration.<br />Set VolumeName to any name to override the default behaviour.<br />Set to "default" if you are not using namespaces within StorageOS.<br />Namespaces that do not pre-exist within StorageOS will be created.<br />+optional|

<h2 id="tocS_v1.StorageMedium">v1.StorageMedium</h2>

<a id="schemav1.storagemedium"></a>
<a id="schema_v1.StorageMedium"></a>
<a id="tocSv1.storagemedium"></a>
<a id="tocsv1.storagemedium"></a>

```json
""

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*||
|*anonymous*|Memory|
|*anonymous*|HugePages|
|*anonymous*|HugePages-|

<h2 id="tocS_v1.SessionAffinityConfig">v1.SessionAffinityConfig</h2>

<a id="schemav1.sessionaffinityconfig"></a>
<a id="schema_v1.SessionAffinityConfig"></a>
<a id="tocSv1.sessionaffinityconfig"></a>
<a id="tocsv1.sessionaffinityconfig"></a>

```json
{
  "clientIP": {
    "timeoutSeconds": 0
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|clientIP|[v1.ClientIPConfig](#schemav1.clientipconfig)|false|none||clientIP contains the configurations of Client IP based session affinity.<br />+optional|

<h2 id="tocS_v1.ServiceType">v1.ServiceType</h2>

<a id="schemav1.servicetype"></a>
<a id="schema_v1.ServiceType"></a>
<a id="tocSv1.servicetype"></a>
<a id="tocsv1.servicetype"></a>

```json
"ClusterIP"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|ClusterIP|
|*anonymous*|NodePort|
|*anonymous*|LoadBalancer|
|*anonymous*|ExternalName|

<h2 id="tocS_v1.ServiceStatus">v1.ServiceStatus</h2>

<a id="schemav1.servicestatus"></a>
<a id="schema_v1.ServiceStatus"></a>
<a id="tocSv1.servicestatus"></a>
<a id="tocsv1.servicestatus"></a>

```json
{
  "conditions": [
    {
      "lastTransitionTime": "string",
      "message": "string",
      "observedGeneration": 0,
      "reason": "string",
      "status": "True",
      "type": "string"
    }
  ],
  "loadBalancer": {
    "ingress": [
      {
        "hostname": "string",
        "ip": "string",
        "ports": [
          null
        ]
      }
    ]
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|conditions|[[v1.Condition](#schemav1.condition)]|false|none||Current service state<br />+optional<br />+patchMergeKey=type<br />+patchStrategy=merge<br />+listType=map<br />+listMapKey=type|
|loadBalancer|[v1.LoadBalancerStatus](#schemav1.loadbalancerstatus)|false|none||LoadBalancer contains the current status of the load-balancer,<br />if one is present.<br />+optional|

<h2 id="tocS_v1.ServiceSpec">v1.ServiceSpec</h2>

<a id="schemav1.servicespec"></a>
<a id="schema_v1.ServiceSpec"></a>
<a id="tocSv1.servicespec"></a>
<a id="tocsv1.servicespec"></a>

```json
{
  "allocateLoadBalancerNodePorts": true,
  "clusterIP": "string",
  "clusterIPs": [
    "string"
  ],
  "externalIPs": [
    "string"
  ],
  "externalName": "string",
  "externalTrafficPolicy": "Local",
  "healthCheckNodePort": 0,
  "internalTrafficPolicy": "Cluster",
  "ipFamilies": [
    "IPv4"
  ],
  "ipFamilyPolicy": "SingleStack",
  "loadBalancerClass": "string",
  "loadBalancerIP": "string",
  "loadBalancerSourceRanges": [
    "string"
  ],
  "ports": [
    {
      "appProtocol": "string",
      "name": "string",
      "nodePort": 0,
      "port": 0,
      "protocol": "TCP",
      "targetPort": {
        "intVal": 0,
        "strVal": "string",
        "type": "["
      }
    }
  ],
  "publishNotReadyAddresses": true,
  "selector": {
    "property1": "string",
    "property2": "string"
  },
  "sessionAffinity": "ClientIP",
  "sessionAffinityConfig": {
    "clientIP": {
      "timeoutSeconds": null
    }
  },
  "type": "ClusterIP"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|allocateLoadBalancerNodePorts|boolean|false|none||allocateLoadBalancerNodePorts defines if NodePorts will be automatically<br />allocated for services with type LoadBalancer.  Default is "true". It<br />may be set to "false" if the cluster load-balancer does not rely on<br />NodePorts.  If the caller requests specific NodePorts (by specifying a<br />value), those requests will be respected, regardless of this field.<br />This field may only be set for services with type LoadBalancer and will<br />be cleared if the type is changed to any other type.<br />This field is beta-level and is only honored by servers that enable the ServiceLBNodePortControl feature.<br />+featureGate=ServiceLBNodePortControl<br />+optional|
|clusterIP|string|false|none||clusterIP is the IP address of the service and is usually assigned<br />randomly. If an address is specified manually, is in-range (as per<br />system configuration), and is not in use, it will be allocated to the<br />service; otherwise creation of the service will fail. This field may not<br />be changed through updates unless the type field is also being changed<br />to ExternalName (which requires this field to be blank) or the type<br />field is being changed from ExternalName (in which case this field may<br />optionally be specified, as describe above).  Valid values are "None",<br />empty string (""), or a valid IP address. Setting this to "None" makes a<br />"headless service" (no virtual IP), which is useful when direct endpoint<br />connections are preferred and proxying is not required.  Only applies to<br />types ClusterIP, NodePort, and LoadBalancer. If this field is specified<br />when creating a Service of type ExternalName, creation will fail. This<br />field will be wiped when updating a Service to type ExternalName.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies<br />+optional|
|clusterIPs|[string]|false|none||ClusterIPs is a list of IP addresses assigned to this service, and are<br />usually assigned randomly.  If an address is specified manually, is<br />in-range (as per system configuration), and is not in use, it will be<br />allocated to the service; otherwise creation of the service will fail.<br />This field may not be changed through updates unless the type field is<br />also being changed to ExternalName (which requires this field to be<br />empty) or the type field is being changed from ExternalName (in which<br />case this field may optionally be specified, as describe above).  Valid<br />values are "None", empty string (""), or a valid IP address.  Setting<br />this to "None" makes a "headless service" (no virtual IP), which is<br />useful when direct endpoint connections are preferred and proxying is<br />not required.  Only applies to types ClusterIP, NodePort, and<br />LoadBalancer. If this field is specified when creating a Service of type<br />ExternalName, creation will fail. This field will be wiped when updating<br />a Service to type ExternalName.  If this field is not specified, it will<br />be initialized from the clusterIP field.  If this field is specified,<br />clients must ensure that clusterIPs[0] and clusterIP have the same<br />value.<br /><br />Unless the "IPv6DualStack" feature gate is enabled, this field is<br />limited to one value, which must be the same as the clusterIP field.  If<br />the feature gate is enabled, this field may hold a maximum of two<br />entries (dual-stack IPs, in either order).  These IPs must correspond to<br />the values of the ipFamilies field. Both clusterIPs and ipFamilies are<br />governed by the ipFamilyPolicy field.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies<br />+listType=atomic<br />+optional|
|externalIPs|[string]|false|none||externalIPs is a list of IP addresses for which nodes in the cluster<br />will also accept traffic for this service.  These IPs are not managed by<br />Kubernetes.  The user is responsible for ensuring that traffic arrives<br />at a node with this IP.  A common example is external load-balancers<br />that are not part of the Kubernetes system.<br />+optional|
|externalName|string|false|none||externalName is the external reference that discovery mechanisms will<br />return as an alias for this service (e.g. a DNS CNAME record). No<br />proxying will be involved.  Must be a lowercase RFC-1123 hostname<br />(https://tools.ietf.org/html/rfc1123) and requires `type` to be "ExternalName".<br />+optional|
|externalTrafficPolicy|[v1.ServiceExternalTrafficPolicyType](#schemav1.serviceexternaltrafficpolicytype)|false|none||externalTrafficPolicy denotes if this Service desires to route external<br />traffic to node-local or cluster-wide endpoints. "Local" preserves the<br />client source IP and avoids a second hop for LoadBalancer and Nodeport<br />type services, but risks potentially imbalanced traffic spreading.<br />"Cluster" obscures the client source IP and may cause a second hop to<br />another node, but should have good overall load-spreading.<br />+optional|
|healthCheckNodePort|integer|false|none||healthCheckNodePort specifies the healthcheck nodePort for the service.<br />This only applies when type is set to LoadBalancer and<br />externalTrafficPolicy is set to Local. If a value is specified, is<br />in-range, and is not in use, it will be used.  If not specified, a value<br />will be automatically allocated.  External systems (e.g. load-balancers)<br />can use this port to determine if a given node holds endpoints for this<br />service or not.  If this field is specified when creating a Service<br />which does not need it, creation will fail. This field will be wiped<br />when updating a Service to no longer need it (e.g. changing type).<br />+optional|
|internalTrafficPolicy|[v1.ServiceInternalTrafficPolicyType](#schemav1.serviceinternaltrafficpolicytype)|false|none||InternalTrafficPolicy specifies if the cluster internal traffic<br />should be routed to all endpoints or node-local endpoints only.<br />"Cluster" routes internal traffic to a Service to all endpoints.<br />"Local" routes traffic to node-local endpoints only, traffic is<br />dropped if no node-local endpoints are ready.<br />The default value is "Cluster".<br />+featureGate=ServiceInternalTrafficPolicy<br />+optional|
|ipFamilies|[[v1.IPFamily](#schemav1.ipfamily)]|false|none||IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this<br />service, and is gated by the "IPv6DualStack" feature gate.  This field<br />is usually assigned automatically based on cluster configuration and the<br />ipFamilyPolicy field. If this field is specified manually, the requested<br />family is available in the cluster, and ipFamilyPolicy allows it, it<br />will be used; otherwise creation of the service will fail.  This field<br />is conditionally mutable: it allows for adding or removing a secondary<br />IP family, but it does not allow changing the primary IP family of the<br />Service.  Valid values are "IPv4" and "IPv6".  This field only applies<br />to Services of types ClusterIP, NodePort, and LoadBalancer, and does<br />apply to "headless" services.  This field will be wiped when updating a<br />Service to type ExternalName.<br /><br />This field may hold a maximum of two entries (dual-stack families, in<br />either order).  These families must correspond to the values of the<br />clusterIPs field, if specified. Both clusterIPs and ipFamilies are<br />governed by the ipFamilyPolicy field.<br />+listType=atomic<br />+optional|
|ipFamilyPolicy|[v1.IPFamilyPolicyType](#schemav1.ipfamilypolicytype)|false|none||IPFamilyPolicy represents the dual-stack-ness requested or required by<br />this Service, and is gated by the "IPv6DualStack" feature gate.  If<br />there is no value provided, then this field will be set to SingleStack.<br />Services can be "SingleStack" (a single IP family), "PreferDualStack"<br />(two IP families on dual-stack configured clusters or a single IP family<br />on single-stack clusters), or "RequireDualStack" (two IP families on<br />dual-stack configured clusters, otherwise fail). The ipFamilies and<br />clusterIPs fields depend on the value of this field.  This field will be<br />wiped when updating a service to type ExternalName.<br />+optional|
|loadBalancerClass|string|false|none||loadBalancerClass is the class of the load balancer implementation this Service belongs to.<br />If specified, the value of this field must be a label-style identifier, with an optional prefix,<br />e.g. "internal-vip" or "example.com/internal-vip". Unprefixed names are reserved for end-users.<br />This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load<br />balancer implementation is used, today this is typically done through the cloud provider integration,<br />but should apply for any default implementation. If set, it is assumed that a load balancer<br />implementation is watching for Services with a matching class. Any default load balancer<br />implementation (e.g. cloud providers) should ignore Services that set this field.<br />This field can only be set when creating or updating a Service to type 'LoadBalancer'.<br />Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.<br />+featureGate=LoadBalancerClass<br />+optional|
|loadBalancerIP|string|false|none||Only applies to Service Type: LoadBalancer<br />LoadBalancer will get created with the IP specified in this field.<br />This feature depends on whether the underlying cloud-provider supports specifying<br />the loadBalancerIP when a load balancer is created.<br />This field will be ignored if the cloud-provider does not support the feature.<br />+optional|
|loadBalancerSourceRanges|[string]|false|none||If specified and supported by the platform, this will restrict traffic through the cloud-provider<br />load-balancer will be restricted to the specified client IPs. This field will be ignored if the<br />cloud-provider does not support the feature."<br />More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/<br />+optional|
|ports|[[v1.ServicePort](#schemav1.serviceport)]|false|none||The list of ports that are exposed by this service.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies<br />+patchMergeKey=port<br />+patchStrategy=merge<br />+listType=map<br />+listMapKey=port<br />+listMapKey=protocol|
|publishNotReadyAddresses|boolean|false|none||publishNotReadyAddresses indicates that any agent which deals with endpoints for this<br />Service should disregard any indications of ready/not-ready.<br />The primary use case for setting this field is for a StatefulSet's Headless Service to<br />propagate SRV DNS records for its Pods for the purpose of peer discovery.<br />The Kubernetes controllers that generate Endpoints and EndpointSlice resources for<br />Services interpret this to mean that all endpoints are considered "ready" even if the<br />Pods themselves are not. Agents which consume only Kubernetes generated endpoints<br />through the Endpoints or EndpointSlice resources can safely assume this behavior.<br />+optional|
|selector|object|false|none||Route service traffic to pods with label keys and values matching this<br />selector. If empty or not present, the service is assumed to have an<br />external process managing its endpoints, which Kubernetes will not<br />modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.<br />Ignored if type is ExternalName.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/<br />+optional<br />+mapType=atomic|
|» **additionalProperties**|string|false|none||none|
|sessionAffinity|[v1.ServiceAffinity](#schemav1.serviceaffinity)|false|none||Supports "ClientIP" and "None". Used to maintain session affinity.<br />Enable client IP based session affinity.<br />Must be ClientIP or None.<br />Defaults to None.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies<br />+optional|
|sessionAffinityConfig|[v1.SessionAffinityConfig](#schemav1.sessionaffinityconfig)|false|none||sessionAffinityConfig contains the configurations of session affinity.<br />+optional|
|type|[v1.ServiceType](#schemav1.servicetype)|false|none||type determines how the Service is exposed. Defaults to ClusterIP. Valid<br />options are ExternalName, ClusterIP, NodePort, and LoadBalancer.<br />"ClusterIP" allocates a cluster-internal IP address for load-balancing<br />to endpoints. Endpoints are determined by the selector or if that is not<br />specified, by manual construction of an Endpoints object or<br />EndpointSlice objects. If clusterIP is "None", no virtual IP is<br />allocated and the endpoints are published as a set of endpoints rather<br />than a virtual IP.<br />"NodePort" builds on ClusterIP and allocates a port on every node which<br />routes to the same endpoints as the clusterIP.<br />"LoadBalancer" builds on NodePort and creates an external load-balancer<br />(if supported in the current cloud) which routes to the same endpoints<br />as the clusterIP.<br />"ExternalName" aliases this service to the specified externalName.<br />Several other fields do not apply to ExternalName services.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types<br />+optional|

<h2 id="tocS_v1.ServicePort">v1.ServicePort</h2>

<a id="schemav1.serviceport"></a>
<a id="schema_v1.ServicePort"></a>
<a id="tocSv1.serviceport"></a>
<a id="tocsv1.serviceport"></a>

```json
{
  "appProtocol": "string",
  "name": "string",
  "nodePort": 0,
  "port": 0,
  "protocol": "TCP",
  "targetPort": {
    "intVal": 0,
    "strVal": "string",
    "type": 0
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|appProtocol|string|false|none||The application protocol for this port.<br />This field follows standard Kubernetes label syntax.<br />Un-prefixed names are reserved for IANA standard service names (as per<br />RFC-6335 and http://www.iana.org/assignments/service-names).<br />Non-standard protocols should use prefixed names such as<br />mycompany.com/my-custom-protocol.<br />+optional|
|name|string|false|none||The name of this port within the service. This must be a DNS_LABEL.<br />All ports within a ServiceSpec must have unique names. When considering<br />the endpoints for a Service, this must match the 'name' field in the<br />EndpointPort.<br />Optional if only one ServicePort is defined on this service.<br />+optional|
|nodePort|integer|false|none||The port on each node on which this service is exposed when type is<br />NodePort or LoadBalancer.  Usually assigned by the system. If a value is<br />specified, in-range, and not in use it will be used, otherwise the<br />operation will fail.  If not specified, a port will be allocated if this<br />Service requires one.  If this field is specified when creating a<br />Service which does not need it, creation will fail. This field will be<br />wiped when updating a Service to no longer need it (e.g. changing type<br />from NodePort to ClusterIP).<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport<br />+optional|
|port|integer|false|none||The port that will be exposed by this service.|
|protocol|[v1.Protocol](#schemav1.protocol)|false|none||The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".<br />Default is TCP.<br />+default="TCP"<br />+optional|
|targetPort|[intstr.IntOrString](#schemaintstr.intorstring)|false|none||Number or name of the port to access on the pods targeted by the service.<br />Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.<br />If this is a string, it will be looked up as a named port in the<br />target Pod's container ports. If this is not specified, the value<br />of the 'port' field is used (an identity map).<br />This field is ignored for services with clusterIP=None, and should be<br />omitted or set equal to the 'port' field.<br />More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service<br />+optional|

<h2 id="tocS_v1.ServiceInternalTrafficPolicyType">v1.ServiceInternalTrafficPolicyType</h2>

<a id="schemav1.serviceinternaltrafficpolicytype"></a>
<a id="schema_v1.ServiceInternalTrafficPolicyType"></a>
<a id="tocSv1.serviceinternaltrafficpolicytype"></a>
<a id="tocsv1.serviceinternaltrafficpolicytype"></a>

```json
"Cluster"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Cluster|
|*anonymous*|Local|

<h2 id="tocS_v1.ServiceExternalTrafficPolicyType">v1.ServiceExternalTrafficPolicyType</h2>

<a id="schemav1.serviceexternaltrafficpolicytype"></a>
<a id="schema_v1.ServiceExternalTrafficPolicyType"></a>
<a id="tocSv1.serviceexternaltrafficpolicytype"></a>
<a id="tocsv1.serviceexternaltrafficpolicytype"></a>

```json
"Local"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Local|
|*anonymous*|Cluster|

<h2 id="tocS_v1.ServiceAffinity">v1.ServiceAffinity</h2>

<a id="schemav1.serviceaffinity"></a>
<a id="schema_v1.ServiceAffinity"></a>
<a id="tocSv1.serviceaffinity"></a>
<a id="tocsv1.serviceaffinity"></a>

```json
"ClientIP"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|ClientIP|
|*anonymous*|None|

<h2 id="tocS_v1.ServiceAccountTokenProjection">v1.ServiceAccountTokenProjection</h2>

<a id="schemav1.serviceaccounttokenprojection"></a>
<a id="schema_v1.ServiceAccountTokenProjection"></a>
<a id="tocSv1.serviceaccounttokenprojection"></a>
<a id="tocsv1.serviceaccounttokenprojection"></a>

```json
{
  "audience": "string",
  "expirationSeconds": 0,
  "path": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|audience|string|false|none||Audience is the intended audience of the token. A recipient of a token<br />must identify itself with an identifier specified in the audience of the<br />token, and otherwise should reject the token. The audience defaults to the<br />identifier of the apiserver.<br />+optional|
|expirationSeconds|integer|false|none||ExpirationSeconds is the requested duration of validity of the service<br />account token. As the token approaches expiration, the kubelet volume<br />plugin will proactively rotate the service account token. The kubelet will<br />start trying to rotate the token if the token is older than 80 percent of<br />its time to live or if the token is older than 24 hours.Defaults to 1 hour<br />and must be at least 10 minutes.<br />+optional|
|path|string|false|none||Path is the path relative to the mount point of the file to project the<br />token into.|

<h2 id="tocS_v1.Service">v1.Service</h2>

<a id="schemav1.service"></a>
<a id="schema_v1.Service"></a>
<a id="tocSv1.service"></a>
<a id="tocsv1.service"></a>

```json
{
  "apiVersion": "string",
  "kind": "string",
  "metadata": {
    "annotations": {
      "property1": "string",
      "property2": "string"
    },
    "clusterName": "string",
    "creationTimestamp": "string",
    "deletionGracePeriodSeconds": 0,
    "deletionTimestamp": "string",
    "finalizers": [
      "string"
    ],
    "generateName": "string",
    "generation": 0,
    "labels": {
      "property1": "string",
      "property2": "string"
    },
    "managedFields": [
      {
        "apiVersion": "string",
        "fieldsType": "string",
        "fieldsV1": null,
        "manager": "string",
        "operation": null,
        "subresource": "string",
        "time": "string"
      }
    ],
    "name": "string",
    "namespace": "string",
    "ownerReferences": [
      {
        "apiVersion": "string",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "string",
        "name": "string",
        "uid": "string"
      }
    ],
    "resourceVersion": "string",
    "selfLink": "string",
    "uid": "string"
  },
  "spec": {
    "allocateLoadBalancerNodePorts": true,
    "clusterIP": "string",
    "clusterIPs": [
      "string"
    ],
    "externalIPs": [
      "string"
    ],
    "externalName": "string",
    "externalTrafficPolicy": "Local",
    "healthCheckNodePort": 0,
    "internalTrafficPolicy": "Cluster",
    "ipFamilies": [
      "IPv4"
    ],
    "ipFamilyPolicy": "SingleStack",
    "loadBalancerClass": "string",
    "loadBalancerIP": "string",
    "loadBalancerSourceRanges": [
      "string"
    ],
    "ports": [
      {
        "appProtocol": "string",
        "name": "string",
        "nodePort": 0,
        "port": 0,
        "protocol": null,
        "targetPort": null
      }
    ],
    "publishNotReadyAddresses": true,
    "selector": {
      "property1": "string",
      "property2": "string"
    },
    "sessionAffinity": "ClientIP",
    "sessionAffinityConfig": {
      "clientIP": null
    },
    "type": "ClusterIP"
  },
  "status": {
    "conditions": [
      {
        "lastTransitionTime": "string",
        "message": "string",
        "observedGeneration": 0,
        "reason": "string",
        "status": null,
        "type": "string"
      }
    ],
    "loadBalancer": {
      "ingress": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiVersion|string|false|none||APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources<br />+optional|
|kind|string|false|none||Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br />+optional|
|metadata|[v1.ObjectMeta](#schemav1.objectmeta)|false|none||Standard object's metadata.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata<br />+optional|
|spec|[v1.ServiceSpec](#schemav1.servicespec)|false|none||Spec defines the behavior of a service.<br />https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status<br />+optional|
|status|[v1.ServiceStatus](#schemav1.servicestatus)|false|none||Most recently observed status of the service.<br />Populated by the system.<br />Read-only.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status<br />+optional|

<h2 id="tocS_v1.SecurityContext">v1.SecurityContext</h2>

<a id="schemav1.securitycontext"></a>
<a id="schema_v1.SecurityContext"></a>
<a id="tocSv1.securitycontext"></a>
<a id="tocsv1.securitycontext"></a>

```json
{
  "allowPrivilegeEscalation": true,
  "capabilities": {
    "add": [
      "string"
    ],
    "drop": [
      "string"
    ]
  },
  "privileged": true,
  "procMount": "Default",
  "readOnlyRootFilesystem": true,
  "runAsGroup": 0,
  "runAsNonRoot": true,
  "runAsUser": 0,
  "seLinuxOptions": {
    "level": "string",
    "role": "string",
    "type": "string",
    "user": "string"
  },
  "seccompProfile": {
    "localhostProfile": "string",
    "type": "Unconfined"
  },
  "windowsOptions": {
    "gmsaCredentialSpec": "string",
    "gmsaCredentialSpecName": "string",
    "hostProcess": true,
    "runAsUserName": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|allowPrivilegeEscalation|boolean|false|none||AllowPrivilegeEscalation controls whether a process can gain more<br />privileges than its parent process. This bool directly controls if<br />the no_new_privs flag will be set on the container process.<br />AllowPrivilegeEscalation is true always when the container is:<br />1) run as Privileged<br />2) has CAP_SYS_ADMIN<br />+optional|
|capabilities|[v1.Capabilities](#schemav1.capabilities)|false|none||The capabilities to add/drop when running containers.<br />Defaults to the default set of capabilities granted by the container runtime.<br />+optional|
|privileged|boolean|false|none||Run container in privileged mode.<br />Processes in privileged containers are essentially equivalent to root on the host.<br />Defaults to false.<br />+optional|
|procMount|[v1.ProcMountType](#schemav1.procmounttype)|false|none||procMount denotes the type of proc mount to use for the containers.<br />The default is DefaultProcMount which uses the container runtime defaults for<br />readonly paths and masked paths.<br />This requires the ProcMountType feature flag to be enabled.<br />+optional|
|readOnlyRootFilesystem|boolean|false|none||Whether this container has a read-only root filesystem.<br />Default is false.<br />+optional|
|runAsGroup|integer|false|none||The GID to run the entrypoint of the container process.<br />Uses runtime default if unset.<br />May also be set in PodSecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|
|runAsNonRoot|boolean|false|none||Indicates that the container must run as a non-root user.<br />If true, the Kubelet will validate the image at runtime to ensure that it<br />does not run as UID 0 (root) and fail to start the container if it does.<br />If unset or false, no such validation will be performed.<br />May also be set in PodSecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|
|runAsUser|integer|false|none||The UID to run the entrypoint of the container process.<br />Defaults to user specified in image metadata if unspecified.<br />May also be set in PodSecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|
|seLinuxOptions|[v1.SELinuxOptions](#schemav1.selinuxoptions)|false|none||The SELinux context to be applied to the container.<br />If unspecified, the container runtime will allocate a random SELinux context for each<br />container.  May also be set in PodSecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|
|seccompProfile|[v1.SeccompProfile](#schemav1.seccompprofile)|false|none||The seccomp options to use by this container. If seccomp options are<br />provided at both the pod & container level, the container options<br />override the pod options.<br />+optional|
|windowsOptions|[v1.WindowsSecurityContextOptions](#schemav1.windowssecuritycontextoptions)|false|none||The Windows specific settings applied to all containers.<br />If unspecified, the options from the PodSecurityContext will be used.<br />If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|

<h2 id="tocS_v1.SecretVolumeSource">v1.SecretVolumeSource</h2>

<a id="schemav1.secretvolumesource"></a>
<a id="schema_v1.SecretVolumeSource"></a>
<a id="tocSv1.secretvolumesource"></a>
<a id="tocsv1.secretvolumesource"></a>

```json
{
  "defaultMode": 0,
  "items": [
    {
      "key": "string",
      "mode": 0,
      "path": "string"
    }
  ],
  "optional": true,
  "secretName": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|defaultMode|integer|false|none||Optional: mode bits used to set permissions on created files by default.<br />Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br />YAML accepts both octal and decimal values, JSON requires decimal values<br />for mode bits. Defaults to 0644.<br />Directories within the path are not affected by this setting.<br />This might be in conflict with other options that affect the file<br />mode, like fsGroup, and the result can be other mode bits set.<br />+optional|
|items|[[v1.KeyToPath](#schemav1.keytopath)]|false|none||If unspecified, each key-value pair in the Data field of the referenced<br />Secret will be projected into the volume as a file whose name is the<br />key and content is the value. If specified, the listed keys will be<br />projected into the specified paths, and unlisted keys will not be<br />present. If a key is specified which is not present in the Secret,<br />the volume setup will error unless it is marked optional. Paths must be<br />relative and may not contain the '..' path or start with '..'.<br />+optional|
|optional|boolean|false|none||Specify whether the Secret or its keys must be defined<br />+optional|
|secretName|string|false|none||Name of the secret in the pod's namespace to use.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#secret<br />+optional|

<h2 id="tocS_v1.SecretProjection">v1.SecretProjection</h2>

<a id="schemav1.secretprojection"></a>
<a id="schema_v1.SecretProjection"></a>
<a id="tocSv1.secretprojection"></a>
<a id="tocsv1.secretprojection"></a>

```json
{
  "items": [
    {
      "key": "string",
      "mode": 0,
      "path": "string"
    }
  ],
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|items|[[v1.KeyToPath](#schemav1.keytopath)]|false|none||If unspecified, each key-value pair in the Data field of the referenced<br />Secret will be projected into the volume as a file whose name is the<br />key and content is the value. If specified, the listed keys will be<br />projected into the specified paths, and unlisted keys will not be<br />present. If a key is specified which is not present in the Secret,<br />the volume setup will error unless it is marked optional. Paths must be<br />relative and may not contain the '..' path or start with '..'.<br />+optional|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the Secret or its key must be defined<br />+optional|

<h2 id="tocS_v1.SecretKeySelector">v1.SecretKeySelector</h2>

<a id="schemav1.secretkeyselector"></a>
<a id="schema_v1.SecretKeySelector"></a>
<a id="tocSv1.secretkeyselector"></a>
<a id="tocsv1.secretkeyselector"></a>

```json
{
  "key": "string",
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|key|string|false|none||The key of the secret to select from.  Must be a valid secret key.|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the Secret or its key must be defined<br />+optional|

<h2 id="tocS_v1.SecretEnvSource">v1.SecretEnvSource</h2>

<a id="schemav1.secretenvsource"></a>
<a id="schema_v1.SecretEnvSource"></a>
<a id="tocSv1.secretenvsource"></a>
<a id="tocsv1.secretenvsource"></a>

```json
{
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the Secret must be defined<br />+optional|

<h2 id="tocS_v1.SeccompProfileType">v1.SeccompProfileType</h2>

<a id="schemav1.seccompprofiletype"></a>
<a id="schema_v1.SeccompProfileType"></a>
<a id="tocSv1.seccompprofiletype"></a>
<a id="tocsv1.seccompprofiletype"></a>

```json
"Unconfined"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Unconfined|
|*anonymous*|RuntimeDefault|
|*anonymous*|Localhost|

<h2 id="tocS_v1.SeccompProfile">v1.SeccompProfile</h2>

<a id="schemav1.seccompprofile"></a>
<a id="schema_v1.SeccompProfile"></a>
<a id="tocSv1.seccompprofile"></a>
<a id="tocsv1.seccompprofile"></a>

```json
{
  "localhostProfile": "string",
  "type": "Unconfined"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|localhostProfile|string|false|none||localhostProfile indicates a profile defined in a file on the node should be used.<br />The profile must be preconfigured on the node to work.<br />Must be a descending path, relative to the kubelet's configured seccomp profile location.<br />Must only be set if type is "Localhost".<br />+optional|
|type|[v1.SeccompProfileType](#schemav1.seccompprofiletype)|false|none||type indicates which kind of seccomp profile will be applied.<br />Valid options are:<br /><br />Localhost - a profile defined in a file on the node should be used.<br />RuntimeDefault - the container runtime default profile should be used.<br />Unconfined - no profile should be applied.<br />+unionDiscriminator|

<h2 id="tocS_v1.ScaleIOVolumeSource">v1.ScaleIOVolumeSource</h2>

<a id="schemav1.scaleiovolumesource"></a>
<a id="schema_v1.ScaleIOVolumeSource"></a>
<a id="tocSv1.scaleiovolumesource"></a>
<a id="tocsv1.scaleiovolumesource"></a>

```json
{
  "fsType": "string",
  "gateway": "string",
  "protectionDomain": "string",
  "readOnly": true,
  "secretRef": {
    "name": "string"
  },
  "sslEnabled": true,
  "storageMode": "string",
  "storagePool": "string",
  "system": "string",
  "volumeName": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs".<br />Default is "xfs".<br />+optional|
|gateway|string|false|none||The host address of the ScaleIO API Gateway.|
|protectionDomain|string|false|none||The name of the ScaleIO Protection Domain for the configured storage.<br />+optional|
|readOnly|boolean|false|none||Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||SecretRef references to the secret for ScaleIO user and other<br />sensitive information. If this is not provided, Login operation will fail.|
|sslEnabled|boolean|false|none||Flag to enable/disable SSL communication with Gateway, default false<br />+optional|
|storageMode|string|false|none||Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.<br />Default is ThinProvisioned.<br />+optional|
|storagePool|string|false|none||The ScaleIO Storage Pool associated with the protection domain.<br />+optional|
|system|string|false|none||The name of the storage system as configured in ScaleIO.|
|volumeName|string|false|none||The name of a volume already created in the ScaleIO system<br />that is associated with this volume source.|

<h2 id="tocS_v1.SELinuxOptions">v1.SELinuxOptions</h2>

<a id="schemav1.selinuxoptions"></a>
<a id="schema_v1.SELinuxOptions"></a>
<a id="tocSv1.selinuxoptions"></a>
<a id="tocsv1.selinuxoptions"></a>

```json
{
  "level": "string",
  "role": "string",
  "type": "string",
  "user": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|level|string|false|none||Level is SELinux level label that applies to the container.<br />+optional|
|role|string|false|none||Role is a SELinux role label that applies to the container.<br />+optional|
|type|string|false|none||Type is a SELinux type label that applies to the container.<br />+optional|
|user|string|false|none||User is a SELinux user label that applies to the container.<br />+optional|

<h2 id="tocS_v1.RollingUpdateDeployment">v1.RollingUpdateDeployment</h2>

<a id="schemav1.rollingupdatedeployment"></a>
<a id="schema_v1.RollingUpdateDeployment"></a>
<a id="tocSv1.rollingupdatedeployment"></a>
<a id="tocsv1.rollingupdatedeployment"></a>

```json
{
  "maxSurge": {
    "intVal": 0,
    "strVal": "string",
    "type": 0
  },
  "maxUnavailable": {
    "intVal": 0,
    "strVal": "string",
    "type": 0
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|maxSurge|[intstr.IntOrString](#schemaintstr.intorstring)|false|none||The maximum number of pods that can be scheduled above the desired number of<br />pods.<br />Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).<br />This can not be 0 if MaxUnavailable is 0.<br />Absolute number is calculated from percentage by rounding up.<br />Defaults to 25%.<br />Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when<br />the rolling update starts, such that the total number of old and new pods do not exceed<br />130% of desired pods. Once old pods have been killed,<br />new ReplicaSet can be scaled up further, ensuring that total number of pods running<br />at any time during the update is at most 130% of desired pods.<br />+optional|
|maxUnavailable|[intstr.IntOrString](#schemaintstr.intorstring)|false|none||The maximum number of pods that can be unavailable during the update.<br />Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).<br />Absolute number is calculated from percentage by rounding down.<br />This can not be 0 if MaxSurge is 0.<br />Defaults to 25%.<br />Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods<br />immediately when the rolling update starts. Once new pods are ready, old ReplicaSet<br />can be scaled down further, followed by scaling up the new ReplicaSet, ensuring<br />that the total number of pods available at all times during the update is at<br />least 70% of desired pods.<br />+optional|

<h2 id="tocS_v1.RestartPolicy">v1.RestartPolicy</h2>

<a id="schemav1.restartpolicy"></a>
<a id="schema_v1.RestartPolicy"></a>
<a id="tocSv1.restartpolicy"></a>
<a id="tocsv1.restartpolicy"></a>

```json
"Always"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Always|
|*anonymous*|OnFailure|
|*anonymous*|Never|

<h2 id="tocS_v1.ResourceRequirements">v1.ResourceRequirements</h2>

<a id="schemav1.resourcerequirements"></a>
<a id="schema_v1.ResourceRequirements"></a>
<a id="tocSv1.resourcerequirements"></a>
<a id="tocsv1.resourcerequirements"></a>

```json
{
  "limits": {
    "property1": {
      "Format": "DecimalExponent"
    },
    "property2": {
      "Format": "DecimalExponent"
    }
  },
  "requests": {
    "property1": {
      "Format": "DecimalExponent"
    },
    "property2": {
      "Format": "DecimalExponent"
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|limits|[v1.ResourceList](#schemav1.resourcelist)|false|none||Limits describes the maximum amount of compute resources allowed.<br />More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/<br />+optional|
|requests|[v1.ResourceList](#schemav1.resourcelist)|false|none||Requests describes the minimum amount of compute resources required.<br />If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,<br />otherwise to an implementation-defined value.<br />More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/<br />+optional|

<h2 id="tocS_v1.ResourceList">v1.ResourceList</h2>

<a id="schemav1.resourcelist"></a>
<a id="schema_v1.ResourceList"></a>
<a id="tocSv1.resourcelist"></a>
<a id="tocsv1.resourcelist"></a>

```json
{
  "property1": {
    "Format": "DecimalExponent"
  },
  "property2": {
    "Format": "DecimalExponent"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|**additionalProperties**|[resource.Quantity](#schemaresource.quantity)|false|none||none|

<h2 id="tocS_v1.ResourceFieldSelector">v1.ResourceFieldSelector</h2>

<a id="schemav1.resourcefieldselector"></a>
<a id="schema_v1.ResourceFieldSelector"></a>
<a id="tocSv1.resourcefieldselector"></a>
<a id="tocsv1.resourcefieldselector"></a>

```json
{
  "containerName": "string",
  "divisor": {
    "Format": "DecimalExponent"
  },
  "resource": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|containerName|string|false|none||Container name: required for volumes, optional for env vars<br />+optional|
|divisor|[resource.Quantity](#schemaresource.quantity)|false|none||Specifies the output format of the exposed resources, defaults to "1"<br />+optional|
|resource|string|false|none||Required: resource to select|

<h2 id="tocS_v1.RBDVolumeSource">v1.RBDVolumeSource</h2>

<a id="schemav1.rbdvolumesource"></a>
<a id="schema_v1.RBDVolumeSource"></a>
<a id="tocSv1.rbdvolumesource"></a>
<a id="tocsv1.rbdvolumesource"></a>

```json
{
  "fsType": "string",
  "image": "string",
  "keyring": "string",
  "monitors": [
    "string"
  ],
  "pool": "string",
  "readOnly": true,
  "secretRef": {
    "name": "string"
  },
  "user": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type of the volume that you want to mount.<br />Tip: Ensure that the filesystem type is supported by the host operating system.<br />Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd<br />TODO: how do we prevent errors in the filesystem from compromising the machine<br />+optional|
|image|string|false|none||The rados image name.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it|
|keyring|string|false|none||Keyring is the path to key ring for RBDUser.<br />Default is /etc/ceph/keyring.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br />+optional|
|monitors|[string]|false|none||A collection of Ceph monitors.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it|
|pool|string|false|none||The rados pool name.<br />Default is rbd.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br />+optional|
|readOnly|boolean|false|none||ReadOnly here will force the ReadOnly setting in VolumeMounts.<br />Defaults to false.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||SecretRef is name of the authentication secret for RBDUser. If provided<br />overrides keyring.<br />Default is nil.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br />+optional|
|user|string|false|none||The rados user name.<br />Default is admin.<br />More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br />+optional|

<h2 id="tocS_v1.QuobyteVolumeSource">v1.QuobyteVolumeSource</h2>

<a id="schemav1.quobytevolumesource"></a>
<a id="schema_v1.QuobyteVolumeSource"></a>
<a id="tocSv1.quobytevolumesource"></a>
<a id="tocsv1.quobytevolumesource"></a>

```json
{
  "group": "string",
  "readOnly": true,
  "registry": "string",
  "tenant": "string",
  "user": "string",
  "volume": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|group|string|false|none||Group to map volume access to<br />Default is no group<br />+optional|
|readOnly|boolean|false|none||ReadOnly here will force the Quobyte volume to be mounted with read-only permissions.<br />Defaults to false.<br />+optional|
|registry|string|false|none||Registry represents a single or multiple Quobyte Registry services<br />specified as a string as host:port pair (multiple entries are separated with commas)<br />which acts as the central registry for volumes|
|tenant|string|false|none||Tenant owning the given Quobyte volume in the Backend<br />Used with dynamically provisioned Quobyte volumes, value is set by the plugin<br />+optional|
|user|string|false|none||User to map volume access to<br />Defaults to serivceaccount user<br />+optional|
|volume|string|false|none||Volume is a string that references an already created Quobyte volume by name.|

<h2 id="tocS_v1.PullPolicy">v1.PullPolicy</h2>

<a id="schemav1.pullpolicy"></a>
<a id="schema_v1.PullPolicy"></a>
<a id="tocSv1.pullpolicy"></a>
<a id="tocsv1.pullpolicy"></a>

```json
"Always"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Always|
|*anonymous*|Never|
|*anonymous*|IfNotPresent|

<h2 id="tocS_v1.Protocol">v1.Protocol</h2>

<a id="schemav1.protocol"></a>
<a id="schema_v1.Protocol"></a>
<a id="tocSv1.protocol"></a>
<a id="tocsv1.protocol"></a>

```json
"TCP"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|TCP|
|*anonymous*|UDP|
|*anonymous*|SCTP|

<h2 id="tocS_v1.ProjectedVolumeSource">v1.ProjectedVolumeSource</h2>

<a id="schemav1.projectedvolumesource"></a>
<a id="schema_v1.ProjectedVolumeSource"></a>
<a id="tocSv1.projectedvolumesource"></a>
<a id="tocsv1.projectedvolumesource"></a>

```json
{
  "defaultMode": 0,
  "sources": [
    {
      "configMap": {
        "items": [
          null
        ],
        "name": "string",
        "optional": true
      },
      "downwardAPI": {
        "items": [
          null
        ]
      },
      "secret": {
        "items": [
          null
        ],
        "name": "string",
        "optional": true
      },
      "serviceAccountToken": {
        "audience": "string",
        "expirationSeconds": 0,
        "path": "string"
      }
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|defaultMode|integer|false|none||Mode bits used to set permissions on created files by default.<br />Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br />YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br />Directories within the path are not affected by this setting.<br />This might be in conflict with other options that affect the file<br />mode, like fsGroup, and the result can be other mode bits set.<br />+optional|
|sources|[[v1.VolumeProjection](#schemav1.volumeprojection)]|false|none||list of volume projections<br />+optional|

<h2 id="tocS_v1.ProcMountType">v1.ProcMountType</h2>

<a id="schemav1.procmounttype"></a>
<a id="schema_v1.ProcMountType"></a>
<a id="tocSv1.procmounttype"></a>
<a id="tocsv1.procmounttype"></a>

```json
"Default"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Default|
|*anonymous*|Unmasked|

<h2 id="tocS_v1.Probe">v1.Probe</h2>

<a id="schemav1.probe"></a>
<a id="schema_v1.Probe"></a>
<a id="tocSv1.probe"></a>
<a id="tocsv1.probe"></a>

```json
{
  "exec": {
    "command": [
      "string"
    ]
  },
  "failureThreshold": 0,
  "httpGet": {
    "host": "string",
    "httpHeaders": [
      {
        "name": "string",
        "value": "string"
      }
    ],
    "path": "string",
    "port": {
      "intVal": null,
      "strVal": null,
      "type": null
    },
    "scheme": "HTTP"
  },
  "initialDelaySeconds": 0,
  "periodSeconds": 0,
  "successThreshold": 0,
  "tcpSocket": {
    "host": "string",
    "port": {
      "intVal": null,
      "strVal": null,
      "type": null
    }
  },
  "terminationGracePeriodSeconds": 0,
  "timeoutSeconds": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|exec|[v1.ExecAction](#schemav1.execaction)|false|none||One and only one of the following should be specified.<br />Exec specifies the action to take.<br />+optional|
|failureThreshold|integer|false|none||Minimum consecutive failures for the probe to be considered failed after having succeeded.<br />Defaults to 3. Minimum value is 1.<br />+optional|
|httpGet|[v1.HTTPGetAction](#schemav1.httpgetaction)|false|none||HTTPGet specifies the http request to perform.<br />+optional|
|initialDelaySeconds|integer|false|none||Number of seconds after the container has started before liveness probes are initiated.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes<br />+optional|
|periodSeconds|integer|false|none||How often (in seconds) to perform the probe.<br />Default to 10 seconds. Minimum value is 1.<br />+optional|
|successThreshold|integer|false|none||Minimum consecutive successes for the probe to be considered successful after having failed.<br />Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.<br />+optional|
|tcpSocket|[v1.TCPSocketAction](#schemav1.tcpsocketaction)|false|none||TCPSocket specifies an action involving a TCP port.<br />TCP hooks not yet supported<br />TODO: implement a realistic TCP lifecycle hook<br />+optional|
|terminationGracePeriodSeconds|integer|false|none||Optional duration in seconds the pod needs to terminate gracefully upon probe failure.<br />The grace period is the duration in seconds after the processes running in the pod are sent<br />a termination signal and the time when the processes are forcibly halted with a kill signal.<br />Set this value longer than the expected cleanup time for your process.<br />If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this<br />value overrides the value provided by the pod spec.<br />Value must be non-negative integer. The value zero indicates stop immediately via<br />the kill signal (no opportunity to shut down).<br />This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.<br />Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.<br />+optional|
|timeoutSeconds|integer|false|none||Number of seconds after which the probe times out.<br />Defaults to 1 second. Minimum value is 1.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes<br />+optional|

<h2 id="tocS_v1.PreferredSchedulingTerm">v1.PreferredSchedulingTerm</h2>

<a id="schemav1.preferredschedulingterm"></a>
<a id="schema_v1.PreferredSchedulingTerm"></a>
<a id="tocSv1.preferredschedulingterm"></a>
<a id="tocsv1.preferredschedulingterm"></a>

```json
{
  "preference": {
    "matchExpressions": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ],
    "matchFields": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ]
  },
  "weight": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|preference|[v1.NodeSelectorTerm](#schemav1.nodeselectorterm)|false|none||A node selector term, associated with the corresponding weight.|
|weight|integer|false|none||Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.|

<h2 id="tocS_v1.PreemptionPolicy">v1.PreemptionPolicy</h2>

<a id="schemav1.preemptionpolicy"></a>
<a id="schema_v1.PreemptionPolicy"></a>
<a id="tocSv1.preemptionpolicy"></a>
<a id="tocsv1.preemptionpolicy"></a>

```json
"PreemptLowerPriority"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|PreemptLowerPriority|
|*anonymous*|Never|

<h2 id="tocS_v1.PortworxVolumeSource">v1.PortworxVolumeSource</h2>

<a id="schemav1.portworxvolumesource"></a>
<a id="schema_v1.PortworxVolumeSource"></a>
<a id="tocSv1.portworxvolumesource"></a>
<a id="tocsv1.portworxvolumesource"></a>

```json
{
  "fsType": "string",
  "readOnly": true,
  "volumeID": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||FSType represents the filesystem type to mount<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.|
|readOnly|boolean|false|none||Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|
|volumeID|string|false|none||VolumeID uniquely identifies a Portworx volume|

<h2 id="tocS_v1.PortStatus">v1.PortStatus</h2>

<a id="schemav1.portstatus"></a>
<a id="schema_v1.PortStatus"></a>
<a id="tocSv1.portstatus"></a>
<a id="tocsv1.portstatus"></a>

```json
{
  "error": "string",
  "port": 0,
  "protocol": "TCP"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|error|string|false|none||Error is to record the problem with the service port<br />The format of the error shall comply with the following rules:<br />- built-in error values shall be specified in this file and those shall use<br />  CamelCase names<br />- cloud provider specific error values must have names that comply with the<br />  format foo.example.com/CamelCase.<br />---<br />The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)<br />+optional<br />+kubebuilder:validation:Required<br />+kubebuilder:validation:Pattern=`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`<br />+kubebuilder:validation:MaxLength=316|
|port|integer|false|none||Port is the port number of the service port of which status is recorded here|
|protocol|[v1.Protocol](#schemav1.protocol)|false|none||Protocol is the protocol of the service port of which status is recorded here<br />The supported values are: "TCP", "UDP", "SCTP"|

<h2 id="tocS_v1.PodTemplateSpec">v1.PodTemplateSpec</h2>

<a id="schemav1.podtemplatespec"></a>
<a id="schema_v1.PodTemplateSpec"></a>
<a id="tocSv1.podtemplatespec"></a>
<a id="tocsv1.podtemplatespec"></a>

```json
{
  "metadata": {
    "annotations": {
      "property1": "string",
      "property2": "string"
    },
    "clusterName": "string",
    "creationTimestamp": "string",
    "deletionGracePeriodSeconds": 0,
    "deletionTimestamp": "string",
    "finalizers": [
      "string"
    ],
    "generateName": "string",
    "generation": 0,
    "labels": {
      "property1": "string",
      "property2": "string"
    },
    "managedFields": [
      {
        "apiVersion": "string",
        "fieldsType": "string",
        "fieldsV1": null,
        "manager": "string",
        "operation": null,
        "subresource": "string",
        "time": "string"
      }
    ],
    "name": "string",
    "namespace": "string",
    "ownerReferences": [
      {
        "apiVersion": "string",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "string",
        "name": "string",
        "uid": "string"
      }
    ],
    "resourceVersion": "string",
    "selfLink": "string",
    "uid": "string"
  },
  "spec": {
    "activeDeadlineSeconds": 0,
    "affinity": {
      "nodeAffinity": null,
      "podAffinity": null,
      "podAntiAffinity": null
    },
    "automountServiceAccountToken": true,
    "containers": [
      {
        "args": [
          null
        ],
        "command": [
          null
        ],
        "env": [
          null
        ],
        "envFrom": [
          null
        ],
        "image": "string",
        "imagePullPolicy": null,
        "lifecycle": null,
        "livenessProbe": null,
        "name": "string",
        "ports": [
          null
        ],
        "readinessProbe": null,
        "resources": null,
        "securityContext": null,
        "startupProbe": null,
        "stdin": true,
        "stdinOnce": true,
        "terminationMessagePath": "string",
        "terminationMessagePolicy": null,
        "tty": true,
        "volumeDevices": [
          null
        ],
        "volumeMounts": [
          null
        ],
        "workingDir": "string"
      }
    ],
    "dnsConfig": {
      "nameservers": null,
      "options": null,
      "searches": null
    },
    "dnsPolicy": "ClusterFirstWithHostNet",
    "enableServiceLinks": true,
    "ephemeralContainers": [
      {
        "args": [
          null
        ],
        "command": [
          null
        ],
        "env": [
          null
        ],
        "envFrom": [
          null
        ],
        "image": "string",
        "imagePullPolicy": null,
        "lifecycle": null,
        "livenessProbe": null,
        "name": "string",
        "ports": [
          null
        ],
        "readinessProbe": null,
        "resources": null,
        "securityContext": null,
        "startupProbe": null,
        "stdin": true,
        "stdinOnce": true,
        "targetContainerName": "string",
        "terminationMessagePath": "string",
        "terminationMessagePolicy": null,
        "tty": true,
        "volumeDevices": [
          null
        ],
        "volumeMounts": [
          null
        ],
        "workingDir": "string"
      }
    ],
    "hostAliases": [
      {
        "hostnames": [
          null
        ],
        "ip": "string"
      }
    ],
    "hostIPC": true,
    "hostNetwork": true,
    "hostPID": true,
    "hostname": "string",
    "imagePullSecrets": [
      {
        "name": "string"
      }
    ],
    "initContainers": [
      {
        "args": [
          null
        ],
        "command": [
          null
        ],
        "env": [
          null
        ],
        "envFrom": [
          null
        ],
        "image": "string",
        "imagePullPolicy": null,
        "lifecycle": null,
        "livenessProbe": null,
        "name": "string",
        "ports": [
          null
        ],
        "readinessProbe": null,
        "resources": null,
        "securityContext": null,
        "startupProbe": null,
        "stdin": true,
        "stdinOnce": true,
        "terminationMessagePath": "string",
        "terminationMessagePolicy": null,
        "tty": true,
        "volumeDevices": [
          null
        ],
        "volumeMounts": [
          null
        ],
        "workingDir": "string"
      }
    ],
    "nodeName": "string",
    "nodeSelector": {
      "property1": "string",
      "property2": "string"
    },
    "overhead": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    },
    "preemptionPolicy": "PreemptLowerPriority",
    "priority": 0,
    "priorityClassName": "string",
    "readinessGates": [
      {
        "conditionType": null
      }
    ],
    "restartPolicy": "Always",
    "runtimeClassName": "string",
    "schedulerName": "string",
    "securityContext": {
      "fsGroup": null,
      "fsGroupChangePolicy": null,
      "runAsGroup": null,
      "runAsNonRoot": null,
      "runAsUser": null,
      "seLinuxOptions": null,
      "seccompProfile": null,
      "supplementalGroups": null,
      "sysctls": null,
      "windowsOptions": null
    },
    "serviceAccount": "string",
    "serviceAccountName": "string",
    "setHostnameAsFQDN": true,
    "shareProcessNamespace": true,
    "subdomain": "string",
    "terminationGracePeriodSeconds": 0,
    "tolerations": [
      {
        "effect": null,
        "key": "string",
        "operator": null,
        "tolerationSeconds": 0,
        "value": "string"
      }
    ],
    "topologySpreadConstraints": [
      {
        "labelSelector": null,
        "maxSkew": 0,
        "topologyKey": "string",
        "whenUnsatisfiable": null
      }
    ],
    "volumes": [
      {
        "awsElasticBlockStore": null,
        "azureDisk": null,
        "azureFile": null,
        "cephfs": null,
        "cinder": null,
        "configMap": null,
        "csi": null,
        "downwardAPI": null,
        "emptyDir": null,
        "ephemeral": null,
        "fc": null,
        "flexVolume": null,
        "flocker": null,
        "gcePersistentDisk": null,
        "gitRepo": null,
        "glusterfs": null,
        "hostPath": null,
        "iscsi": null,
        "name": "string",
        "nfs": null,
        "persistentVolumeClaim": null,
        "photonPersistentDisk": null,
        "portworxVolume": null,
        "projected": null,
        "quobyte": null,
        "rbd": null,
        "scaleIO": null,
        "secret": null,
        "storageos": null,
        "vsphereVolume": null
      }
    ]
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|metadata|[v1.ObjectMeta](#schemav1.objectmeta)|false|none||Standard object's metadata.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata<br />+optional|
|spec|[v1.PodSpec](#schemav1.podspec)|false|none||Specification of the desired behavior of the pod.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status<br />+optional|

<h2 id="tocS_v1.PodSpec">v1.PodSpec</h2>

<a id="schemav1.podspec"></a>
<a id="schema_v1.PodSpec"></a>
<a id="tocSv1.podspec"></a>
<a id="tocsv1.podspec"></a>

```json
{
  "activeDeadlineSeconds": 0,
  "affinity": {
    "nodeAffinity": {
      "preferredDuringSchedulingIgnoredDuringExecution": null,
      "requiredDuringSchedulingIgnoredDuringExecution": null
    },
    "podAffinity": {
      "preferredDuringSchedulingIgnoredDuringExecution": null,
      "requiredDuringSchedulingIgnoredDuringExecution": null
    },
    "podAntiAffinity": {
      "preferredDuringSchedulingIgnoredDuringExecution": null,
      "requiredDuringSchedulingIgnoredDuringExecution": null
    }
  },
  "automountServiceAccountToken": true,
  "containers": [
    {
      "args": [
        "string"
      ],
      "command": [
        "string"
      ],
      "env": [
        {
          "name": "string",
          "value": "string",
          "valueFrom": null
        }
      ],
      "envFrom": [
        {
          "configMapRef": null,
          "prefix": "string",
          "secretRef": null
        }
      ],
      "image": "string",
      "imagePullPolicy": "Always",
      "lifecycle": {
        "postStart": null,
        "preStop": null
      },
      "livenessProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "name": "string",
      "ports": [
        {
          "containerPort": 0,
          "hostIP": "string",
          "hostPort": 0,
          "name": "string",
          "protocol": null
        }
      ],
      "readinessProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "resources": {
        "limits": null,
        "requests": null
      },
      "securityContext": {
        "allowPrivilegeEscalation": true,
        "capabilities": null,
        "privileged": true,
        "procMount": null,
        "readOnlyRootFilesystem": true,
        "runAsGroup": 0,
        "runAsNonRoot": true,
        "runAsUser": 0,
        "seLinuxOptions": null,
        "seccompProfile": null,
        "windowsOptions": null
      },
      "startupProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "stdin": true,
      "stdinOnce": true,
      "terminationMessagePath": "string",
      "terminationMessagePolicy": "File",
      "tty": true,
      "volumeDevices": [
        {
          "devicePath": "string",
          "name": "string"
        }
      ],
      "volumeMounts": [
        {
          "mountPath": "string",
          "mountPropagation": null,
          "name": "string",
          "readOnly": true,
          "subPath": "string",
          "subPathExpr": "string"
        }
      ],
      "workingDir": "string"
    }
  ],
  "dnsConfig": {
    "nameservers": [
      "string"
    ],
    "options": [
      {
        "name": "string",
        "value": "string"
      }
    ],
    "searches": [
      "string"
    ]
  },
  "dnsPolicy": "ClusterFirstWithHostNet",
  "enableServiceLinks": true,
  "ephemeralContainers": [
    {
      "args": [
        "string"
      ],
      "command": [
        "string"
      ],
      "env": [
        {
          "name": "string",
          "value": "string",
          "valueFrom": null
        }
      ],
      "envFrom": [
        {
          "configMapRef": null,
          "prefix": "string",
          "secretRef": null
        }
      ],
      "image": "string",
      "imagePullPolicy": "Always",
      "lifecycle": {
        "postStart": null,
        "preStop": null
      },
      "livenessProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "name": "string",
      "ports": [
        {
          "containerPort": 0,
          "hostIP": "string",
          "hostPort": 0,
          "name": "string",
          "protocol": null
        }
      ],
      "readinessProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "resources": {
        "limits": null,
        "requests": null
      },
      "securityContext": {
        "allowPrivilegeEscalation": true,
        "capabilities": null,
        "privileged": true,
        "procMount": null,
        "readOnlyRootFilesystem": true,
        "runAsGroup": 0,
        "runAsNonRoot": true,
        "runAsUser": 0,
        "seLinuxOptions": null,
        "seccompProfile": null,
        "windowsOptions": null
      },
      "startupProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "stdin": true,
      "stdinOnce": true,
      "targetContainerName": "string",
      "terminationMessagePath": "string",
      "terminationMessagePolicy": "File",
      "tty": true,
      "volumeDevices": [
        {
          "devicePath": "string",
          "name": "string"
        }
      ],
      "volumeMounts": [
        {
          "mountPath": "string",
          "mountPropagation": null,
          "name": "string",
          "readOnly": true,
          "subPath": "string",
          "subPathExpr": "string"
        }
      ],
      "workingDir": "string"
    }
  ],
  "hostAliases": [
    {
      "hostnames": [
        "string"
      ],
      "ip": "string"
    }
  ],
  "hostIPC": true,
  "hostNetwork": true,
  "hostPID": true,
  "hostname": "string",
  "imagePullSecrets": [
    {
      "name": "string"
    }
  ],
  "initContainers": [
    {
      "args": [
        "string"
      ],
      "command": [
        "string"
      ],
      "env": [
        {
          "name": "string",
          "value": "string",
          "valueFrom": null
        }
      ],
      "envFrom": [
        {
          "configMapRef": null,
          "prefix": "string",
          "secretRef": null
        }
      ],
      "image": "string",
      "imagePullPolicy": "Always",
      "lifecycle": {
        "postStart": null,
        "preStop": null
      },
      "livenessProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "name": "string",
      "ports": [
        {
          "containerPort": 0,
          "hostIP": "string",
          "hostPort": 0,
          "name": "string",
          "protocol": null
        }
      ],
      "readinessProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "resources": {
        "limits": null,
        "requests": null
      },
      "securityContext": {
        "allowPrivilegeEscalation": true,
        "capabilities": null,
        "privileged": true,
        "procMount": null,
        "readOnlyRootFilesystem": true,
        "runAsGroup": 0,
        "runAsNonRoot": true,
        "runAsUser": 0,
        "seLinuxOptions": null,
        "seccompProfile": null,
        "windowsOptions": null
      },
      "startupProbe": {
        "exec": null,
        "failureThreshold": 0,
        "httpGet": null,
        "initialDelaySeconds": 0,
        "periodSeconds": 0,
        "successThreshold": 0,
        "tcpSocket": null,
        "terminationGracePeriodSeconds": 0,
        "timeoutSeconds": 0
      },
      "stdin": true,
      "stdinOnce": true,
      "terminationMessagePath": "string",
      "terminationMessagePolicy": "File",
      "tty": true,
      "volumeDevices": [
        {
          "devicePath": "string",
          "name": "string"
        }
      ],
      "volumeMounts": [
        {
          "mountPath": "string",
          "mountPropagation": null,
          "name": "string",
          "readOnly": true,
          "subPath": "string",
          "subPathExpr": "string"
        }
      ],
      "workingDir": "string"
    }
  ],
  "nodeName": "string",
  "nodeSelector": {
    "property1": "string",
    "property2": "string"
  },
  "overhead": {
    "property1": {
      "Format": "DecimalExponent"
    },
    "property2": {
      "Format": "DecimalExponent"
    }
  },
  "preemptionPolicy": "PreemptLowerPriority",
  "priority": 0,
  "priorityClassName": "string",
  "readinessGates": [
    {
      "conditionType": "ContainersReady"
    }
  ],
  "restartPolicy": "Always",
  "runtimeClassName": "string",
  "schedulerName": "string",
  "securityContext": {
    "fsGroup": 0,
    "fsGroupChangePolicy": "OnRootMismatch",
    "runAsGroup": 0,
    "runAsNonRoot": true,
    "runAsUser": 0,
    "seLinuxOptions": {
      "level": null,
      "role": null,
      "type": null,
      "user": null
    },
    "seccompProfile": {
      "localhostProfile": null,
      "type": null
    },
    "supplementalGroups": [
      0
    ],
    "sysctls": [
      {
        "name": "string",
        "value": "string"
      }
    ],
    "windowsOptions": {
      "gmsaCredentialSpec": null,
      "gmsaCredentialSpecName": null,
      "hostProcess": null,
      "runAsUserName": null
    }
  },
  "serviceAccount": "string",
  "serviceAccountName": "string",
  "setHostnameAsFQDN": true,
  "shareProcessNamespace": true,
  "subdomain": "string",
  "terminationGracePeriodSeconds": 0,
  "tolerations": [
    {
      "effect": "NoSchedule",
      "key": "string",
      "operator": "Exists",
      "tolerationSeconds": 0,
      "value": "string"
    }
  ],
  "topologySpreadConstraints": [
    {
      "labelSelector": {
        "matchExpressions": [
          null
        ],
        "matchLabels": {}
      },
      "maxSkew": 0,
      "topologyKey": "string",
      "whenUnsatisfiable": "DoNotSchedule"
    }
  ],
  "volumes": [
    {
      "awsElasticBlockStore": {
        "fsType": "string",
        "partition": 0,
        "readOnly": true,
        "volumeID": "string"
      },
      "azureDisk": {
        "cachingMode": null,
        "diskName": "string",
        "diskURI": "string",
        "fsType": "string",
        "kind": null,
        "readOnly": true
      },
      "azureFile": {
        "readOnly": true,
        "secretName": "string",
        "shareName": "string"
      },
      "cephfs": {
        "monitors": [
          null
        ],
        "path": "string",
        "readOnly": true,
        "secretFile": "string",
        "secretRef": null,
        "user": "string"
      },
      "cinder": {
        "fsType": "string",
        "readOnly": true,
        "secretRef": null,
        "volumeID": "string"
      },
      "configMap": {
        "defaultMode": 0,
        "items": [
          null
        ],
        "name": "string",
        "optional": true
      },
      "csi": {
        "driver": "string",
        "fsType": "string",
        "nodePublishSecretRef": null,
        "readOnly": true,
        "volumeAttributes": {}
      },
      "downwardAPI": {
        "defaultMode": 0,
        "items": [
          null
        ]
      },
      "emptyDir": {
        "medium": null,
        "sizeLimit": null
      },
      "ephemeral": {
        "volumeClaimTemplate": null
      },
      "fc": {
        "fsType": "string",
        "lun": 0,
        "readOnly": true,
        "targetWWNs": [
          null
        ],
        "wwids": [
          null
        ]
      },
      "flexVolume": {
        "driver": "string",
        "fsType": "string",
        "options": {},
        "readOnly": true,
        "secretRef": null
      },
      "flocker": {
        "datasetName": "string",
        "datasetUUID": "string"
      },
      "gcePersistentDisk": {
        "fsType": "string",
        "partition": 0,
        "pdName": "string",
        "readOnly": true
      },
      "gitRepo": {
        "directory": "string",
        "repository": "string",
        "revision": "string"
      },
      "glusterfs": {
        "endpoints": "string",
        "path": "string",
        "readOnly": true
      },
      "hostPath": {
        "path": "string",
        "type": null
      },
      "iscsi": {
        "chapAuthDiscovery": true,
        "chapAuthSession": true,
        "fsType": "string",
        "initiatorName": "string",
        "iqn": "string",
        "iscsiInterface": "string",
        "lun": 0,
        "portals": [
          null
        ],
        "readOnly": true,
        "secretRef": null,
        "targetPortal": "string"
      },
      "name": "string",
      "nfs": {
        "path": "string",
        "readOnly": true,
        "server": "string"
      },
      "persistentVolumeClaim": {
        "claimName": "string",
        "readOnly": true
      },
      "photonPersistentDisk": {
        "fsType": "string",
        "pdID": "string"
      },
      "portworxVolume": {
        "fsType": "string",
        "readOnly": true,
        "volumeID": "string"
      },
      "projected": {
        "defaultMode": 0,
        "sources": [
          null
        ]
      },
      "quobyte": {
        "group": "string",
        "readOnly": true,
        "registry": "string",
        "tenant": "string",
        "user": "string",
        "volume": "string"
      },
      "rbd": {
        "fsType": "string",
        "image": "string",
        "keyring": "string",
        "monitors": [
          null
        ],
        "pool": "string",
        "readOnly": true,
        "secretRef": null,
        "user": "string"
      },
      "scaleIO": {
        "fsType": "string",
        "gateway": "string",
        "protectionDomain": "string",
        "readOnly": true,
        "secretRef": null,
        "sslEnabled": true,
        "storageMode": "string",
        "storagePool": "string",
        "system": "string",
        "volumeName": "string"
      },
      "secret": {
        "defaultMode": 0,
        "items": [
          null
        ],
        "optional": true,
        "secretName": "string"
      },
      "storageos": {
        "fsType": "string",
        "readOnly": true,
        "secretRef": null,
        "volumeName": "string",
        "volumeNamespace": "string"
      },
      "vsphereVolume": {
        "fsType": "string",
        "storagePolicyID": "string",
        "storagePolicyName": "string",
        "volumePath": "string"
      }
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|activeDeadlineSeconds|integer|false|none||Optional duration in seconds the pod may be active on the node relative to<br />StartTime before the system will actively try to mark it failed and kill associated containers.<br />Value must be a positive integer.<br />+optional|
|affinity|[v1.Affinity](#schemav1.affinity)|false|none||If specified, the pod's scheduling constraints<br />+optional|
|automountServiceAccountToken|boolean|false|none||AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.<br />+optional|
|containers|[[v1.Container](#schemav1.container)]|false|none||List of containers belonging to the pod.<br />Containers cannot currently be added or removed.<br />There must be at least one container in a Pod.<br />Cannot be updated.<br />+patchMergeKey=name<br />+patchStrategy=merge|
|dnsConfig|[v1.PodDNSConfig](#schemav1.poddnsconfig)|false|none||Specifies the DNS parameters of a pod.<br />Parameters specified here will be merged to the generated DNS<br />configuration based on DNSPolicy.<br />+optional|
|dnsPolicy|[v1.DNSPolicy](#schemav1.dnspolicy)|false|none||Set DNS policy for the pod.<br />Defaults to "ClusterFirst".<br />Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.<br />DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.<br />To have DNS options set along with hostNetwork, you have to specify DNS policy<br />explicitly to 'ClusterFirstWithHostNet'.<br />+optional|
|enableServiceLinks|boolean|false|none||EnableServiceLinks indicates whether information about services should be injected into pod's<br />environment variables, matching the syntax of Docker links.<br />Optional: Defaults to true.<br />+optional|
|ephemeralContainers|[[v1.EphemeralContainer](#schemav1.ephemeralcontainer)]|false|none||List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing<br />pod to perform user-initiated actions such as debugging. This list cannot be specified when<br />creating a pod, and it cannot be modified by updating the pod spec. In order to add an<br />ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource.<br />This field is alpha-level and is only honored by servers that enable the EphemeralContainers feature.<br />+optional<br />+patchMergeKey=name<br />+patchStrategy=merge|
|hostAliases|[[v1.HostAlias](#schemav1.hostalias)]|false|none||HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts<br />file if specified. This is only valid for non-hostNetwork pods.<br />+optional<br />+patchMergeKey=ip<br />+patchStrategy=merge|
|hostIPC|boolean|false|none||Use the host's ipc namespace.<br />Optional: Default to false.<br />+k8s:conversion-gen=false<br />+optional|
|hostNetwork|boolean|false|none||Host networking requested for this pod. Use the host's network namespace.<br />If this option is set, the ports that will be used must be specified.<br />Default to false.<br />+k8s:conversion-gen=false<br />+optional|
|hostPID|boolean|false|none||Use the host's pid namespace.<br />Optional: Default to false.<br />+k8s:conversion-gen=false<br />+optional|
|hostname|string|false|none||Specifies the hostname of the Pod<br />If not specified, the pod's hostname will be set to a system-defined value.<br />+optional|
|imagePullSecrets|[[v1.LocalObjectReference](#schemav1.localobjectreference)]|false|none||ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.<br />If specified, these secrets will be passed to individual puller implementations for them to use. For example,<br />in the case of docker, only DockerConfig type secrets are honored.<br />More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod<br />+optional<br />+patchMergeKey=name<br />+patchStrategy=merge|
|initContainers|[[v1.Container](#schemav1.container)]|false|none||List of initialization containers belonging to the pod.<br />Init containers are executed in order prior to containers being started. If any<br />init container fails, the pod is considered to have failed and is handled according<br />to its restartPolicy. The name for an init container or normal container must be<br />unique among all containers.<br />Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.<br />The resourceRequirements of an init container are taken into account during scheduling<br />by finding the highest request/limit for each resource type, and then using the max of<br />of that value or the sum of the normal containers. Limits are applied to init containers<br />in a similar fashion.<br />Init containers cannot currently be added or removed.<br />Cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/<br />+patchMergeKey=name<br />+patchStrategy=merge|
|nodeName|string|false|none||NodeName is a request to schedule this pod onto a specific node. If it is non-empty,<br />the scheduler simply schedules this pod onto that node, assuming that it fits resource<br />requirements.<br />+optional|
|nodeSelector|object|false|none||NodeSelector is a selector which must be true for the pod to fit on a node.<br />Selector which must match a node's labels for the pod to be scheduled on that node.<br />More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/<br />+optional<br />+mapType=atomic|
|» **additionalProperties**|string|false|none||none|
|overhead|[v1.ResourceList](#schemav1.resourcelist)|false|none||Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.<br />This field will be autopopulated at admission time by the RuntimeClass admission controller. If<br />the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests.<br />The RuntimeClass admission controller will reject Pod create requests which have the overhead already<br />set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value<br />defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero.<br />More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md<br />This field is beta-level as of Kubernetes v1.18, and is only honored by servers that enable the PodOverhead feature.<br />+optional|
|preemptionPolicy|[v1.PreemptionPolicy](#schemav1.preemptionpolicy)|false|none||PreemptionPolicy is the Policy for preempting pods with lower priority.<br />One of Never, PreemptLowerPriority.<br />Defaults to PreemptLowerPriority if unset.<br />This field is beta-level, gated by the NonPreemptingPriority feature-gate.<br />+optional|
|priority|integer|false|none||The priority value. Various system components use this field to find the<br />priority of the pod. When Priority Admission Controller is enabled, it<br />prevents users from setting this field. The admission controller populates<br />this field from PriorityClassName.<br />The higher the value, the higher the priority.<br />+optional|
|priorityClassName|string|false|none||If specified, indicates the pod's priority. "system-node-critical" and<br />"system-cluster-critical" are two special keywords which indicate the<br />highest priorities with the former being the highest priority. Any other<br />name must be defined by creating a PriorityClass object with that name.<br />If not specified, the pod priority will be default or zero if there is no<br />default.<br />+optional|
|readinessGates|[[v1.PodReadinessGate](#schemav1.podreadinessgate)]|false|none||If specified, all readiness gates will be evaluated for pod readiness.<br />A pod is ready when all its containers are ready AND<br />all conditions specified in the readiness gates have status equal to "True"<br />More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates<br />+optional|
|restartPolicy|[v1.RestartPolicy](#schemav1.restartpolicy)|false|none||Restart policy for all containers within the pod.<br />One of Always, OnFailure, Never.<br />Default to Always.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy<br />+optional|
|runtimeClassName|string|false|none||RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used<br />to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.<br />If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an<br />empty definition that uses the default runtime handler.<br />More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class<br />This is a beta feature as of Kubernetes v1.14.<br />+optional|
|schedulerName|string|false|none||If specified, the pod will be dispatched by specified scheduler.<br />If not specified, the pod will be dispatched by default scheduler.<br />+optional|
|securityContext|[v1.PodSecurityContext](#schemav1.podsecuritycontext)|false|none||SecurityContext holds pod-level security attributes and common container settings.<br />Optional: Defaults to empty.  See type description for default values of each field.<br />+optional|
|serviceAccount|string|false|none||DeprecatedServiceAccount is a depreciated alias for ServiceAccountName.<br />Deprecated: Use serviceAccountName instead.<br />+k8s:conversion-gen=false<br />+optional|
|serviceAccountName|string|false|none||ServiceAccountName is the name of the ServiceAccount to use to run this pod.<br />More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/<br />+optional|
|setHostnameAsFQDN|boolean|false|none||If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default).<br />In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).<br />In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN.<br />If a pod does not have FQDN, this has no effect.<br />Default to false.<br />+optional|
|shareProcessNamespace|boolean|false|none||Share a single process namespace between all of the containers in a pod.<br />When this is set containers will be able to view and signal processes from other containers<br />in the same pod, and the first process in each container will not be assigned PID 1.<br />HostPID and ShareProcessNamespace cannot both be set.<br />Optional: Default to false.<br />+k8s:conversion-gen=false<br />+optional|
|subdomain|string|false|none||If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".<br />If not specified, the pod will not have a domainname at all.<br />+optional|
|terminationGracePeriodSeconds|integer|false|none||Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.<br />Value must be non-negative integer. The value zero indicates stop immediately via<br />the kill signal (no opportunity to shut down).<br />If this value is nil, the default grace period will be used instead.<br />The grace period is the duration in seconds after the processes running in the pod are sent<br />a termination signal and the time when the processes are forcibly halted with a kill signal.<br />Set this value longer than the expected cleanup time for your process.<br />Defaults to 30 seconds.<br />+optional|
|tolerations|[[v1.Toleration](#schemav1.toleration)]|false|none||If specified, the pod's tolerations.<br />+optional|
|topologySpreadConstraints|[[v1.TopologySpreadConstraint](#schemav1.topologyspreadconstraint)]|false|none||TopologySpreadConstraints describes how a group of pods ought to spread across topology<br />domains. Scheduler will schedule pods in a way which abides by the constraints.<br />All topologySpreadConstraints are ANDed.<br />+optional<br />+patchMergeKey=topologyKey<br />+patchStrategy=merge<br />+listType=map<br />+listMapKey=topologyKey<br />+listMapKey=whenUnsatisfiable|
|volumes|[[v1.Volume](#schemav1.volume)]|false|none||List of volumes that can be mounted by containers belonging to the pod.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes<br />+optional<br />+patchMergeKey=name<br />+patchStrategy=merge,retainKeys|

<h2 id="tocS_v1.PodSecurityContext">v1.PodSecurityContext</h2>

<a id="schemav1.podsecuritycontext"></a>
<a id="schema_v1.PodSecurityContext"></a>
<a id="tocSv1.podsecuritycontext"></a>
<a id="tocsv1.podsecuritycontext"></a>

```json
{
  "fsGroup": 0,
  "fsGroupChangePolicy": "OnRootMismatch",
  "runAsGroup": 0,
  "runAsNonRoot": true,
  "runAsUser": 0,
  "seLinuxOptions": {
    "level": "string",
    "role": "string",
    "type": "string",
    "user": "string"
  },
  "seccompProfile": {
    "localhostProfile": "string",
    "type": "Unconfined"
  },
  "supplementalGroups": [
    0
  ],
  "sysctls": [
    {
      "name": "string",
      "value": "string"
    }
  ],
  "windowsOptions": {
    "gmsaCredentialSpec": "string",
    "gmsaCredentialSpecName": "string",
    "hostProcess": true,
    "runAsUserName": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsGroup|integer|false|none||A special supplemental group that applies to all containers in a pod.<br />Some volume types allow the Kubelet to change the ownership of that volume<br />to be owned by the pod:<br /><br />1. The owning GID will be the FSGroup<br />2. The setgid bit is set (new files created in the volume will be owned by FSGroup)<br />3. The permission bits are OR'd with rw-rw----<br /><br />If unset, the Kubelet will not modify the ownership and permissions of any volume.<br />+optional|
|fsGroupChangePolicy|[v1.PodFSGroupChangePolicy](#schemav1.podfsgroupchangepolicy)|false|none||fsGroupChangePolicy defines behavior of changing ownership and permission of the volume<br />before being exposed inside Pod. This field will only apply to<br />volume types which support fsGroup based ownership(and permissions).<br />It will have no effect on ephemeral volume types such as: secret, configmaps<br />and emptydir.<br />Valid values are "OnRootMismatch" and "Always". If not specified, "Always" is used.<br />+optional|
|runAsGroup|integer|false|none||The GID to run the entrypoint of the container process.<br />Uses runtime default if unset.<br />May also be set in SecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence<br />for that container.<br />+optional|
|runAsNonRoot|boolean|false|none||Indicates that the container must run as a non-root user.<br />If true, the Kubelet will validate the image at runtime to ensure that it<br />does not run as UID 0 (root) and fail to start the container if it does.<br />If unset or false, no such validation will be performed.<br />May also be set in SecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|
|runAsUser|integer|false|none||The UID to run the entrypoint of the container process.<br />Defaults to user specified in image metadata if unspecified.<br />May also be set in SecurityContext.  If set in both SecurityContext and<br />PodSecurityContext, the value specified in SecurityContext takes precedence<br />for that container.<br />+optional|
|seLinuxOptions|[v1.SELinuxOptions](#schemav1.selinuxoptions)|false|none||The SELinux context to be applied to all containers.<br />If unspecified, the container runtime will allocate a random SELinux context for each<br />container.  May also be set in SecurityContext.  If set in<br />both SecurityContext and PodSecurityContext, the value specified in SecurityContext<br />takes precedence for that container.<br />+optional|
|seccompProfile|[v1.SeccompProfile](#schemav1.seccompprofile)|false|none||The seccomp options to use by the containers in this pod.<br />+optional|
|supplementalGroups|[integer]|false|none||A list of groups applied to the first process run in each container, in addition<br />to the container's primary GID.  If unspecified, no groups will be added to<br />any container.<br />+optional|
|sysctls|[[v1.Sysctl](#schemav1.sysctl)]|false|none||Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported<br />sysctls (by the container runtime) might fail to launch.<br />+optional|
|windowsOptions|[v1.WindowsSecurityContextOptions](#schemav1.windowssecuritycontextoptions)|false|none||The Windows specific settings applied to all containers.<br />If unspecified, the options within a container's SecurityContext will be used.<br />If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.<br />+optional|

<h2 id="tocS_v1.PodReadinessGate">v1.PodReadinessGate</h2>

<a id="schemav1.podreadinessgate"></a>
<a id="schema_v1.PodReadinessGate"></a>
<a id="tocSv1.podreadinessgate"></a>
<a id="tocsv1.podreadinessgate"></a>

```json
{
  "conditionType": "ContainersReady"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|conditionType|[v1.PodConditionType](#schemav1.podconditiontype)|false|none||ConditionType refers to a condition in the pod's condition list with matching type.|

<h2 id="tocS_v1.PodFSGroupChangePolicy">v1.PodFSGroupChangePolicy</h2>

<a id="schemav1.podfsgroupchangepolicy"></a>
<a id="schema_v1.PodFSGroupChangePolicy"></a>
<a id="tocSv1.podfsgroupchangepolicy"></a>
<a id="tocsv1.podfsgroupchangepolicy"></a>

```json
"OnRootMismatch"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|OnRootMismatch|
|*anonymous*|Always|

<h2 id="tocS_v1.PodDNSConfigOption">v1.PodDNSConfigOption</h2>

<a id="schemav1.poddnsconfigoption"></a>
<a id="schema_v1.PodDNSConfigOption"></a>
<a id="tocSv1.poddnsconfigoption"></a>
<a id="tocsv1.poddnsconfigoption"></a>

```json
{
  "name": "string",
  "value": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Required.|
|value|string|false|none||+optional|

<h2 id="tocS_v1.PodDNSConfig">v1.PodDNSConfig</h2>

<a id="schemav1.poddnsconfig"></a>
<a id="schema_v1.PodDNSConfig"></a>
<a id="tocSv1.poddnsconfig"></a>
<a id="tocsv1.poddnsconfig"></a>

```json
{
  "nameservers": [
    "string"
  ],
  "options": [
    {
      "name": "string",
      "value": "string"
    }
  ],
  "searches": [
    "string"
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|nameservers|[string]|false|none||A list of DNS name server IP addresses.<br />This will be appended to the base nameservers generated from DNSPolicy.<br />Duplicated nameservers will be removed.<br />+optional|
|options|[[v1.PodDNSConfigOption](#schemav1.poddnsconfigoption)]|false|none||A list of DNS resolver options.<br />This will be merged with the base options generated from DNSPolicy.<br />Duplicated entries will be removed. Resolution options given in Options<br />will override those that appear in the base DNSPolicy.<br />+optional|
|searches|[string]|false|none||A list of DNS search domains for host-name lookup.<br />This will be appended to the base search paths generated from DNSPolicy.<br />Duplicated search paths will be removed.<br />+optional|

<h2 id="tocS_v1.PodConditionType">v1.PodConditionType</h2>

<a id="schemav1.podconditiontype"></a>
<a id="schema_v1.PodConditionType"></a>
<a id="tocSv1.podconditiontype"></a>
<a id="tocsv1.podconditiontype"></a>

```json
"ContainersReady"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|ContainersReady|
|*anonymous*|Initialized|
|*anonymous*|Ready|
|*anonymous*|PodScheduled|

<h2 id="tocS_v1.PodAntiAffinity">v1.PodAntiAffinity</h2>

<a id="schemav1.podantiaffinity"></a>
<a id="schema_v1.PodAntiAffinity"></a>
<a id="tocSv1.podantiaffinity"></a>
<a id="tocsv1.podantiaffinity"></a>

```json
{
  "preferredDuringSchedulingIgnoredDuringExecution": [
    {
      "podAffinityTerm": {
        "labelSelector": null,
        "namespaceSelector": null,
        "namespaces": [
          null
        ],
        "topologyKey": "string"
      },
      "weight": 0
    }
  ],
  "requiredDuringSchedulingIgnoredDuringExecution": [
    {
      "labelSelector": {
        "matchExpressions": [
          null
        ],
        "matchLabels": {}
      },
      "namespaceSelector": {
        "matchExpressions": [
          null
        ],
        "matchLabels": {}
      },
      "namespaces": [
        "string"
      ],
      "topologyKey": "string"
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|preferredDuringSchedulingIgnoredDuringExecution|[[v1.WeightedPodAffinityTerm](#schemav1.weightedpodaffinityterm)]|false|none||The scheduler will prefer to schedule pods to nodes that satisfy<br />the anti-affinity expressions specified by this field, but it may choose<br />a node that violates one or more of the expressions. The node that is<br />most preferred is the one with the greatest sum of weights, i.e.<br />for each node that meets all of the scheduling requirements (resource<br />request, requiredDuringScheduling anti-affinity expressions, etc.),<br />compute a sum by iterating through the elements of this field and adding<br />"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br />node(s) with the highest sum are the most preferred.<br />+optional|
|requiredDuringSchedulingIgnoredDuringExecution|[[v1.PodAffinityTerm](#schemav1.podaffinityterm)]|false|none||If the anti-affinity requirements specified by this field are not met at<br />scheduling time, the pod will not be scheduled onto the node.<br />If the anti-affinity requirements specified by this field cease to be met<br />at some point during pod execution (e.g. due to a pod label update), the<br />system may or may not try to eventually evict the pod from its node.<br />When there are multiple elements, the lists of nodes corresponding to each<br />podAffinityTerm are intersected, i.e. all terms must be satisfied.<br />+optional|

<h2 id="tocS_v1.PodAffinityTerm">v1.PodAffinityTerm</h2>

<a id="schemav1.podaffinityterm"></a>
<a id="schema_v1.PodAffinityTerm"></a>
<a id="tocSv1.podaffinityterm"></a>
<a id="tocsv1.podaffinityterm"></a>

```json
{
  "labelSelector": {
    "matchExpressions": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ],
    "matchLabels": {
      "property1": "string",
      "property2": "string"
    }
  },
  "namespaceSelector": {
    "matchExpressions": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ],
    "matchLabels": {
      "property1": "string",
      "property2": "string"
    }
  },
  "namespaces": [
    "string"
  ],
  "topologyKey": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|labelSelector|[v1.LabelSelector](#schemav1.labelselector)|false|none||A label query over a set of resources, in this case pods.<br />+optional|
|namespaceSelector|[v1.LabelSelector](#schemav1.labelselector)|false|none||A label query over the set of namespaces that the term applies to.<br />The term is applied to the union of the namespaces selected by this field<br />and the ones listed in the namespaces field.<br />null selector and null or empty namespaces list means "this pod's namespace".<br />An empty selector ({}) matches all namespaces.<br />This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.<br />+optional|
|namespaces|[string]|false|none||namespaces specifies a static list of namespace names that the term applies to.<br />The term is applied to the union of the namespaces listed in this field<br />and the ones selected by namespaceSelector.<br />null or empty namespaces list and null namespaceSelector means "this pod's namespace"<br />+optional|
|topologyKey|string|false|none||This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br />the labelSelector in the specified namespaces, where co-located is defined as running on a node<br />whose value of the label with key topologyKey matches that of any node on which any of the<br />selected pods is running.<br />Empty topologyKey is not allowed.|

<h2 id="tocS_v1.PodAffinity">v1.PodAffinity</h2>

<a id="schemav1.podaffinity"></a>
<a id="schema_v1.PodAffinity"></a>
<a id="tocSv1.podaffinity"></a>
<a id="tocsv1.podaffinity"></a>

```json
{
  "preferredDuringSchedulingIgnoredDuringExecution": [
    {
      "podAffinityTerm": {
        "labelSelector": null,
        "namespaceSelector": null,
        "namespaces": [
          null
        ],
        "topologyKey": "string"
      },
      "weight": 0
    }
  ],
  "requiredDuringSchedulingIgnoredDuringExecution": [
    {
      "labelSelector": {
        "matchExpressions": [
          null
        ],
        "matchLabels": {}
      },
      "namespaceSelector": {
        "matchExpressions": [
          null
        ],
        "matchLabels": {}
      },
      "namespaces": [
        "string"
      ],
      "topologyKey": "string"
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|preferredDuringSchedulingIgnoredDuringExecution|[[v1.WeightedPodAffinityTerm](#schemav1.weightedpodaffinityterm)]|false|none||The scheduler will prefer to schedule pods to nodes that satisfy<br />the affinity expressions specified by this field, but it may choose<br />a node that violates one or more of the expressions. The node that is<br />most preferred is the one with the greatest sum of weights, i.e.<br />for each node that meets all of the scheduling requirements (resource<br />request, requiredDuringScheduling affinity expressions, etc.),<br />compute a sum by iterating through the elements of this field and adding<br />"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br />node(s) with the highest sum are the most preferred.<br />+optional|
|requiredDuringSchedulingIgnoredDuringExecution|[[v1.PodAffinityTerm](#schemav1.podaffinityterm)]|false|none||If the affinity requirements specified by this field are not met at<br />scheduling time, the pod will not be scheduled onto the node.<br />If the affinity requirements specified by this field cease to be met<br />at some point during pod execution (e.g. due to a pod label update), the<br />system may or may not try to eventually evict the pod from its node.<br />When there are multiple elements, the lists of nodes corresponding to each<br />podAffinityTerm are intersected, i.e. all terms must be satisfied.<br />+optional|

<h2 id="tocS_v1.PhotonPersistentDiskVolumeSource">v1.PhotonPersistentDiskVolumeSource</h2>

<a id="schemav1.photonpersistentdiskvolumesource"></a>
<a id="schema_v1.PhotonPersistentDiskVolumeSource"></a>
<a id="tocSv1.photonpersistentdiskvolumesource"></a>
<a id="tocsv1.photonpersistentdiskvolumesource"></a>

```json
{
  "fsType": "string",
  "pdID": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.|
|pdID|string|false|none||ID that identifies Photon Controller persistent disk|

<h2 id="tocS_v1.PersistentVolumeMode">v1.PersistentVolumeMode</h2>

<a id="schemav1.persistentvolumemode"></a>
<a id="schema_v1.PersistentVolumeMode"></a>
<a id="tocSv1.persistentvolumemode"></a>
<a id="tocsv1.persistentvolumemode"></a>

```json
"Block"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Block|
|*anonymous*|Filesystem|

<h2 id="tocS_v1.PersistentVolumeClaimVolumeSource">v1.PersistentVolumeClaimVolumeSource</h2>

<a id="schemav1.persistentvolumeclaimvolumesource"></a>
<a id="schema_v1.PersistentVolumeClaimVolumeSource"></a>
<a id="tocSv1.persistentvolumeclaimvolumesource"></a>
<a id="tocsv1.persistentvolumeclaimvolumesource"></a>

```json
{
  "claimName": "string",
  "readOnly": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|claimName|string|false|none||ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims|
|readOnly|boolean|false|none||Will force the ReadOnly setting in VolumeMounts.<br />Default false.<br />+optional|

<h2 id="tocS_v1.PersistentVolumeClaimTemplate">v1.PersistentVolumeClaimTemplate</h2>

<a id="schemav1.persistentvolumeclaimtemplate"></a>
<a id="schema_v1.PersistentVolumeClaimTemplate"></a>
<a id="tocSv1.persistentvolumeclaimtemplate"></a>
<a id="tocsv1.persistentvolumeclaimtemplate"></a>

```json
{
  "metadata": {
    "annotations": {
      "property1": "string",
      "property2": "string"
    },
    "clusterName": "string",
    "creationTimestamp": "string",
    "deletionGracePeriodSeconds": 0,
    "deletionTimestamp": "string",
    "finalizers": [
      "string"
    ],
    "generateName": "string",
    "generation": 0,
    "labels": {
      "property1": "string",
      "property2": "string"
    },
    "managedFields": [
      {
        "apiVersion": "string",
        "fieldsType": "string",
        "fieldsV1": null,
        "manager": "string",
        "operation": null,
        "subresource": "string",
        "time": "string"
      }
    ],
    "name": "string",
    "namespace": "string",
    "ownerReferences": [
      {
        "apiVersion": "string",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "string",
        "name": "string",
        "uid": "string"
      }
    ],
    "resourceVersion": "string",
    "selfLink": "string",
    "uid": "string"
  },
  "spec": {
    "accessModes": [
      "ReadWriteOnce"
    ],
    "dataSource": {
      "apiGroup": null,
      "kind": null,
      "name": null
    },
    "dataSourceRef": {
      "apiGroup": null,
      "kind": null,
      "name": null
    },
    "resources": {
      "limits": null,
      "requests": null
    },
    "selector": {
      "matchExpressions": null,
      "matchLabels": null
    },
    "storageClassName": "string",
    "volumeMode": "Block",
    "volumeName": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|metadata|[v1.ObjectMeta](#schemav1.objectmeta)|false|none||May contain labels and annotations that will be copied into the PVC<br />when creating it. No other fields are allowed and will be rejected during<br />validation.<br /><br />+optional|
|spec|[v1.PersistentVolumeClaimSpec](#schemav1.persistentvolumeclaimspec)|false|none||The specification for the PersistentVolumeClaim. The entire content is<br />copied unchanged into the PVC that gets created from this<br />template. The same fields as in a PersistentVolumeClaim<br />are also valid here.|

<h2 id="tocS_v1.PersistentVolumeClaimStatus">v1.PersistentVolumeClaimStatus</h2>

<a id="schemav1.persistentvolumeclaimstatus"></a>
<a id="schema_v1.PersistentVolumeClaimStatus"></a>
<a id="tocSv1.persistentvolumeclaimstatus"></a>
<a id="tocsv1.persistentvolumeclaimstatus"></a>

```json
{
  "accessModes": [
    "ReadWriteOnce"
  ],
  "capacity": {
    "property1": {
      "Format": "DecimalExponent"
    },
    "property2": {
      "Format": "DecimalExponent"
    }
  },
  "conditions": [
    {
      "lastProbeTime": "string",
      "lastTransitionTime": "string",
      "message": "string",
      "reason": "string",
      "status": "True",
      "type": "Resizing"
    }
  ],
  "phase": "Pending"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|accessModes|[[v1.PersistentVolumeAccessMode](#schemav1.persistentvolumeaccessmode)]|false|none||AccessModes contains the actual access modes the volume backing the PVC has.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1<br />+optional|
|capacity|[v1.ResourceList](#schemav1.resourcelist)|false|none||Represents the actual resources of the underlying volume.<br />+optional|
|conditions|[[v1.PersistentVolumeClaimCondition](#schemav1.persistentvolumeclaimcondition)]|false|none||Current Condition of persistent volume claim. If underlying persistent volume is being<br />resized then the Condition will be set to 'ResizeStarted'.<br />+optional<br />+patchMergeKey=type<br />+patchStrategy=merge|
|phase|[v1.PersistentVolumeClaimPhase](#schemav1.persistentvolumeclaimphase)|false|none||Phase represents the current phase of PersistentVolumeClaim.<br />+optional|

<h2 id="tocS_v1.PersistentVolumeClaimSpec">v1.PersistentVolumeClaimSpec</h2>

<a id="schemav1.persistentvolumeclaimspec"></a>
<a id="schema_v1.PersistentVolumeClaimSpec"></a>
<a id="tocSv1.persistentvolumeclaimspec"></a>
<a id="tocsv1.persistentvolumeclaimspec"></a>

```json
{
  "accessModes": [
    "ReadWriteOnce"
  ],
  "dataSource": {
    "apiGroup": "string",
    "kind": "string",
    "name": "string"
  },
  "dataSourceRef": {
    "apiGroup": "string",
    "kind": "string",
    "name": "string"
  },
  "resources": {
    "limits": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    },
    "requests": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    }
  },
  "selector": {
    "matchExpressions": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ],
    "matchLabels": {
      "property1": "string",
      "property2": "string"
    }
  },
  "storageClassName": "string",
  "volumeMode": "Block",
  "volumeName": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|accessModes|[[v1.PersistentVolumeAccessMode](#schemav1.persistentvolumeaccessmode)]|false|none||AccessModes contains the desired access modes the volume should have.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1<br />+optional|
|dataSource|[v1.TypedLocalObjectReference](#schemav1.typedlocalobjectreference)|false|none||This field can be used to specify either:<br />* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br />* An existing PVC (PersistentVolumeClaim)<br />If the provisioner or an external controller can support the specified data source,<br />it will create a new volume based on the contents of the specified data source.<br />If the AnyVolumeDataSource feature gate is enabled, this field will always have<br />the same contents as the DataSourceRef field.<br />+optional|
|dataSourceRef|[v1.TypedLocalObjectReference](#schemav1.typedlocalobjectreference)|false|none||Specifies the object from which to populate the volume with data, if a non-empty<br />volume is desired. This may be any local object from a non-empty API group (non<br />core object) or a PersistentVolumeClaim object.<br />When this field is specified, volume binding will only succeed if the type of<br />the specified object matches some installed volume populator or dynamic<br />provisioner.<br />This field will replace the functionality of the DataSource field and as such<br />if both fields are non-empty, they must have the same value. For backwards<br />compatibility, both fields (DataSource and DataSourceRef) will be set to the same<br />value automatically if one of them is empty and the other is non-empty.<br />There are two important differences between DataSource and DataSourceRef:<br />* While DataSource only allows two specific types of objects, DataSourceRef<br />  allows any non-core object, as well as PersistentVolumeClaim objects.<br />* While DataSource ignores disallowed values (dropping them), DataSourceRef<br />  preserves all values, and generates an error if a disallowed value is<br />  specified.<br />(Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br />+optional|
|resources|[v1.ResourceRequirements](#schemav1.resourcerequirements)|false|none||Resources represents the minimum resources the volume should have.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources<br />+optional|
|selector|[v1.LabelSelector](#schemav1.labelselector)|false|none||A label query over volumes to consider for binding.<br />+optional|
|storageClassName|string|false|none||Name of the StorageClass required by the claim.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1<br />+optional|
|volumeMode|[v1.PersistentVolumeMode](#schemav1.persistentvolumemode)|false|none||volumeMode defines what type of volume is required by the claim.<br />Value of Filesystem is implied when not included in claim spec.<br />+optional|
|volumeName|string|false|none||VolumeName is the binding reference to the PersistentVolume backing this claim.<br />+optional|

<h2 id="tocS_v1.PersistentVolumeClaimPhase">v1.PersistentVolumeClaimPhase</h2>

<a id="schemav1.persistentvolumeclaimphase"></a>
<a id="schema_v1.PersistentVolumeClaimPhase"></a>
<a id="tocSv1.persistentvolumeclaimphase"></a>
<a id="tocsv1.persistentvolumeclaimphase"></a>

```json
"Pending"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Pending|
|*anonymous*|Bound|
|*anonymous*|Lost|

<h2 id="tocS_v1.PersistentVolumeClaimConditionType">v1.PersistentVolumeClaimConditionType</h2>

<a id="schemav1.persistentvolumeclaimconditiontype"></a>
<a id="schema_v1.PersistentVolumeClaimConditionType"></a>
<a id="tocSv1.persistentvolumeclaimconditiontype"></a>
<a id="tocsv1.persistentvolumeclaimconditiontype"></a>

```json
"Resizing"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Resizing|
|*anonymous*|FileSystemResizePending|

<h2 id="tocS_v1.PersistentVolumeClaimCondition">v1.PersistentVolumeClaimCondition</h2>

<a id="schemav1.persistentvolumeclaimcondition"></a>
<a id="schema_v1.PersistentVolumeClaimCondition"></a>
<a id="tocSv1.persistentvolumeclaimcondition"></a>
<a id="tocsv1.persistentvolumeclaimcondition"></a>

```json
{
  "lastProbeTime": "string",
  "lastTransitionTime": "string",
  "message": "string",
  "reason": "string",
  "status": "True",
  "type": "Resizing"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|lastProbeTime|string|false|none||Last time we probed the condition.<br />+optional|
|lastTransitionTime|string|false|none||Last time the condition transitioned from one status to another.<br />+optional|
|message|string|false|none||Human-readable message indicating details about last transition.<br />+optional|
|reason|string|false|none||Unique, this should be a short, machine understandable string that gives the reason<br />for condition's last transition. If it reports "ResizeStarted" that means the underlying<br />persistent volume is being resized.<br />+optional|
|status|[k8s_io_api_core_v1.ConditionStatus](#schemak8s_io_api_core_v1.conditionstatus)|false|none||none|
|type|[v1.PersistentVolumeClaimConditionType](#schemav1.persistentvolumeclaimconditiontype)|false|none||none|

<h2 id="tocS_v1.PersistentVolumeClaim">v1.PersistentVolumeClaim</h2>

<a id="schemav1.persistentvolumeclaim"></a>
<a id="schema_v1.PersistentVolumeClaim"></a>
<a id="tocSv1.persistentvolumeclaim"></a>
<a id="tocsv1.persistentvolumeclaim"></a>

```json
{
  "apiVersion": "string",
  "kind": "string",
  "metadata": {
    "annotations": {
      "property1": "string",
      "property2": "string"
    },
    "clusterName": "string",
    "creationTimestamp": "string",
    "deletionGracePeriodSeconds": 0,
    "deletionTimestamp": "string",
    "finalizers": [
      "string"
    ],
    "generateName": "string",
    "generation": 0,
    "labels": {
      "property1": "string",
      "property2": "string"
    },
    "managedFields": [
      {
        "apiVersion": "string",
        "fieldsType": "string",
        "fieldsV1": null,
        "manager": "string",
        "operation": null,
        "subresource": "string",
        "time": "string"
      }
    ],
    "name": "string",
    "namespace": "string",
    "ownerReferences": [
      {
        "apiVersion": "string",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "string",
        "name": "string",
        "uid": "string"
      }
    ],
    "resourceVersion": "string",
    "selfLink": "string",
    "uid": "string"
  },
  "spec": {
    "accessModes": [
      "ReadWriteOnce"
    ],
    "dataSource": {
      "apiGroup": null,
      "kind": null,
      "name": null
    },
    "dataSourceRef": {
      "apiGroup": null,
      "kind": null,
      "name": null
    },
    "resources": {
      "limits": null,
      "requests": null
    },
    "selector": {
      "matchExpressions": null,
      "matchLabels": null
    },
    "storageClassName": "string",
    "volumeMode": "Block",
    "volumeName": "string"
  },
  "status": {
    "accessModes": [
      "ReadWriteOnce"
    ],
    "capacity": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    },
    "conditions": [
      {
        "lastProbeTime": "string",
        "lastTransitionTime": "string",
        "message": "string",
        "reason": "string",
        "status": "[",
        "type": "["
      }
    ],
    "phase": "Pending"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiVersion|string|false|none||APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources<br />+optional|
|kind|string|false|none||Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br />+optional|
|metadata|[v1.ObjectMeta](#schemav1.objectmeta)|false|none||Standard object's metadata.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata<br />+optional|
|spec|[v1.PersistentVolumeClaimSpec](#schemav1.persistentvolumeclaimspec)|false|none||Spec defines the desired characteristics of a volume requested by a pod author.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims<br />+optional|
|status|[v1.PersistentVolumeClaimStatus](#schemav1.persistentvolumeclaimstatus)|false|none||Status represents the current information/status of a persistent volume claim.<br />Read-only.<br />More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims<br />+optional|

<h2 id="tocS_v1.PersistentVolumeAccessMode">v1.PersistentVolumeAccessMode</h2>

<a id="schemav1.persistentvolumeaccessmode"></a>
<a id="schema_v1.PersistentVolumeAccessMode"></a>
<a id="tocSv1.persistentvolumeaccessmode"></a>
<a id="tocsv1.persistentvolumeaccessmode"></a>

```json
"ReadWriteOnce"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|ReadWriteOnce|
|*anonymous*|ReadOnlyMany|
|*anonymous*|ReadWriteMany|
|*anonymous*|ReadWriteOncePod|

<h2 id="tocS_v1.OwnerReference">v1.OwnerReference</h2>

<a id="schemav1.ownerreference"></a>
<a id="schema_v1.OwnerReference"></a>
<a id="tocSv1.ownerreference"></a>
<a id="tocsv1.ownerreference"></a>

```json
{
  "apiVersion": "string",
  "blockOwnerDeletion": true,
  "controller": true,
  "kind": "string",
  "name": "string",
  "uid": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiVersion|string|false|none||API version of the referent.|
|blockOwnerDeletion|boolean|false|none||If true, AND if the owner has the "foregroundDeletion" finalizer, then<br />the owner cannot be deleted from the key-value store until this<br />reference is removed.<br />Defaults to false.<br />To set this field, a user needs "delete" permission of the owner,<br />otherwise 422 (Unprocessable Entity) will be returned.<br />+optional|
|controller|boolean|false|none||If true, this reference points to the managing controller.<br />+optional|
|kind|string|false|none||Kind of the referent.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds|
|name|string|false|none||Name of the referent.<br />More info: http://kubernetes.io/docs/user-guide/identifiers#names|
|uid|string|false|none||UID of the referent.<br />More info: http://kubernetes.io/docs/user-guide/identifiers#uids|

<h2 id="tocS_v1.ObjectMeta">v1.ObjectMeta</h2>

<a id="schemav1.objectmeta"></a>
<a id="schema_v1.ObjectMeta"></a>
<a id="tocSv1.objectmeta"></a>
<a id="tocsv1.objectmeta"></a>

```json
{
  "annotations": {
    "property1": "string",
    "property2": "string"
  },
  "clusterName": "string",
  "creationTimestamp": "string",
  "deletionGracePeriodSeconds": 0,
  "deletionTimestamp": "string",
  "finalizers": [
    "string"
  ],
  "generateName": "string",
  "generation": 0,
  "labels": {
    "property1": "string",
    "property2": "string"
  },
  "managedFields": [
    {
      "apiVersion": "string",
      "fieldsType": "string",
      "fieldsV1": {},
      "manager": "string",
      "operation": "Apply",
      "subresource": "string",
      "time": "string"
    }
  ],
  "name": "string",
  "namespace": "string",
  "ownerReferences": [
    {
      "apiVersion": "string",
      "blockOwnerDeletion": true,
      "controller": true,
      "kind": "string",
      "name": "string",
      "uid": "string"
    }
  ],
  "resourceVersion": "string",
  "selfLink": "string",
  "uid": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|annotations|object|false|none||Annotations is an unstructured key value map stored with a resource that may be<br />set by external tools to store and retrieve arbitrary metadata. They are not<br />queryable and should be preserved when modifying objects.<br />More info: http://kubernetes.io/docs/user-guide/annotations<br />+optional|
|» **additionalProperties**|string|false|none||none|
|clusterName|string|false|none||The name of the cluster which the object belongs to.<br />This is used to distinguish resources with same name and namespace in different clusters.<br />This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.<br />+optional|
|creationTimestamp|string|false|none||CreationTimestamp is a timestamp representing the server time when this object was<br />created. It is not guaranteed to be set in happens-before order across separate operations.<br />Clients may not set this value. It is represented in RFC3339 form and is in UTC.<br /><br />Populated by the system.<br />Read-only.<br />Null for lists.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata<br />+optional|
|deletionGracePeriodSeconds|integer|false|none||Number of seconds allowed for this object to gracefully terminate before<br />it will be removed from the system. Only set when deletionTimestamp is also set.<br />May only be shortened.<br />Read-only.<br />+optional|
|deletionTimestamp|string|false|none||DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This<br />field is set by the server when a graceful deletion is requested by the user, and is not<br />directly settable by a client. The resource is expected to be deleted (no longer visible<br />from resource lists, and not reachable by name) after the time in this field, once the<br />finalizers list is empty. As long as the finalizers list contains items, deletion is blocked.<br />Once the deletionTimestamp is set, this value may not be unset or be set further into the<br />future, although it may be shortened or the resource may be deleted prior to this time.<br />For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react<br />by sending a graceful termination signal to the containers in the pod. After that 30 seconds,<br />the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup,<br />remove the pod from the API. In the presence of network partitions, this object may still<br />exist after this timestamp, until an administrator or automated process can determine the<br />resource is fully terminated.<br />If not set, graceful deletion of the object has not been requested.<br /><br />Populated by the system when a graceful deletion is requested.<br />Read-only.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata<br />+optional|
|finalizers|[string]|false|none||Must be empty before the object is deleted from the registry. Each entry<br />is an identifier for the responsible component that will remove the entry<br />from the list. If the deletionTimestamp of the object is non-nil, entries<br />in this list can only be removed.<br />Finalizers may be processed and removed in any order.  Order is NOT enforced<br />because it introduces significant risk of stuck finalizers.<br />finalizers is a shared field, any actor with permission can reorder it.<br />If the finalizer list is processed in order, then this can lead to a situation<br />in which the component responsible for the first finalizer in the list is<br />waiting for a signal (field value, external system, or other) produced by a<br />component responsible for a finalizer later in the list, resulting in a deadlock.<br />Without enforced ordering finalizers are free to order amongst themselves and<br />are not vulnerable to ordering changes in the list.<br />+optional<br />+patchStrategy=merge|
|generateName|string|false|none||GenerateName is an optional prefix, used by the server, to generate a unique<br />name ONLY IF the Name field has not been provided.<br />If this field is used, the name returned to the client will be different<br />than the name passed. This value will also be combined with a unique suffix.<br />The provided value has the same validation rules as the Name field,<br />and may be truncated by the length of the suffix required to make the value<br />unique on the server.<br /><br />If this field is specified and the generated name exists, the server will<br />NOT return a 409 - instead, it will either return 201 Created or 500 with Reason<br />ServerTimeout indicating a unique name could not be found in the time allotted, and the client<br />should retry (optionally after the time indicated in the Retry-After header).<br /><br />Applied only if Name is not specified.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency<br />+optional|
|generation|integer|false|none||A sequence number representing a specific generation of the desired state.<br />Populated by the system. Read-only.<br />+optional|
|labels|object|false|none||Map of string keys and values that can be used to organize and categorize<br />(scope and select) objects. May match selectors of replication controllers<br />and services.<br />More info: http://kubernetes.io/docs/user-guide/labels<br />+optional|
|» **additionalProperties**|string|false|none||none|
|managedFields|[[v1.ManagedFieldsEntry](#schemav1.managedfieldsentry)]|false|none||ManagedFields maps workflow-id and version to the set of fields<br />that are managed by that workflow. This is mostly for internal<br />housekeeping, and users typically shouldn't need to set or<br />understand this field. A workflow can be the user's name, a<br />controller's name, or the name of a specific apply path like<br />"ci-cd". The set of fields is always in the version that the<br />workflow used when modifying the object.<br /><br />+optional|
|name|string|false|none||Name must be unique within a namespace. Is required when creating resources, although<br />some resources may allow a client to request the generation of an appropriate name<br />automatically. Name is primarily intended for creation idempotence and configuration<br />definition.<br />Cannot be updated.<br />More info: http://kubernetes.io/docs/user-guide/identifiers#names<br />+optional|
|namespace|string|false|none||Namespace defines the space within which each name must be unique. An empty namespace is<br />equivalent to the "default" namespace, but "default" is the canonical representation.<br />Not all objects are required to be scoped to a namespace - the value of this field for<br />those objects will be empty.<br /><br />Must be a DNS_LABEL.<br />Cannot be updated.<br />More info: http://kubernetes.io/docs/user-guide/namespaces<br />+optional|
|ownerReferences|[[v1.OwnerReference](#schemav1.ownerreference)]|false|none||List of objects depended by this object. If ALL objects in the list have<br />been deleted, this object will be garbage collected. If this object is managed by a controller,<br />then an entry in this list will point to this controller, with the controller field set to true.<br />There cannot be more than one managing controller.<br />+optional<br />+patchMergeKey=uid<br />+patchStrategy=merge|
|resourceVersion|string|false|none||An opaque value that represents the internal version of this object that can<br />be used by clients to determine when objects have changed. May be used for optimistic<br />concurrency, change detection, and the watch operation on a resource or set of resources.<br />Clients must treat these values as opaque and passed unmodified back to the server.<br />They may only be valid for a particular resource or set of resources.<br /><br />Populated by the system.<br />Read-only.<br />Value must be treated as opaque by clients and .<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency<br />+optional|
|selfLink|string|false|none||SelfLink is a URL representing this object.<br />Populated by the system.<br />Read-only.<br /><br />DEPRECATED<br />Kubernetes will stop propagating this field in 1.20 release and the field is planned<br />to be removed in 1.21 release.<br />+optional|
|uid|string|false|none||UID is the unique in time and space value for this object. It is typically generated by<br />the server on successful creation of a resource and is not allowed to change on PUT<br />operations.<br /><br />Populated by the system.<br />Read-only.<br />More info: http://kubernetes.io/docs/user-guide/identifiers#uids<br />+optional|

<h2 id="tocS_v1.ObjectFieldSelector">v1.ObjectFieldSelector</h2>

<a id="schemav1.objectfieldselector"></a>
<a id="schema_v1.ObjectFieldSelector"></a>
<a id="tocSv1.objectfieldselector"></a>
<a id="tocsv1.objectfieldselector"></a>

```json
{
  "apiVersion": "string",
  "fieldPath": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiVersion|string|false|none||Version of the schema the FieldPath is written in terms of, defaults to "v1".<br />+optional|
|fieldPath|string|false|none||Path of the field to select in the specified API version.|

<h2 id="tocS_v1.NodeSelectorTerm">v1.NodeSelectorTerm</h2>

<a id="schemav1.nodeselectorterm"></a>
<a id="schema_v1.NodeSelectorTerm"></a>
<a id="tocSv1.nodeselectorterm"></a>
<a id="tocsv1.nodeselectorterm"></a>

```json
{
  "matchExpressions": [
    {
      "key": "string",
      "operator": "In",
      "values": [
        "string"
      ]
    }
  ],
  "matchFields": [
    {
      "key": "string",
      "operator": "In",
      "values": [
        "string"
      ]
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|matchExpressions|[[v1.NodeSelectorRequirement](#schemav1.nodeselectorrequirement)]|false|none||A list of node selector requirements by node's labels.<br />+optional|
|matchFields|[[v1.NodeSelectorRequirement](#schemav1.nodeselectorrequirement)]|false|none||A list of node selector requirements by node's fields.<br />+optional|

<h2 id="tocS_v1.NodeSelectorRequirement">v1.NodeSelectorRequirement</h2>

<a id="schemav1.nodeselectorrequirement"></a>
<a id="schema_v1.NodeSelectorRequirement"></a>
<a id="tocSv1.nodeselectorrequirement"></a>
<a id="tocsv1.nodeselectorrequirement"></a>

```json
{
  "key": "string",
  "operator": "In",
  "values": [
    "string"
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|key|string|false|none||The label key that the selector applies to.|
|operator|[v1.NodeSelectorOperator](#schemav1.nodeselectoroperator)|false|none||Represents a key's relationship to a set of values.<br />Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.|
|values|[string]|false|none||An array of string values. If the operator is In or NotIn,<br />the values array must be non-empty. If the operator is Exists or DoesNotExist,<br />the values array must be empty. If the operator is Gt or Lt, the values<br />array must have a single element, which will be interpreted as an integer.<br />This array is replaced during a strategic merge patch.<br />+optional|

<h2 id="tocS_v1.NodeSelectorOperator">v1.NodeSelectorOperator</h2>

<a id="schemav1.nodeselectoroperator"></a>
<a id="schema_v1.NodeSelectorOperator"></a>
<a id="tocSv1.nodeselectoroperator"></a>
<a id="tocsv1.nodeselectoroperator"></a>

```json
"In"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|In|
|*anonymous*|NotIn|
|*anonymous*|Exists|
|*anonymous*|DoesNotExist|
|*anonymous*|Gt|
|*anonymous*|Lt|

<h2 id="tocS_v1.NodeSelector">v1.NodeSelector</h2>

<a id="schemav1.nodeselector"></a>
<a id="schema_v1.NodeSelector"></a>
<a id="tocSv1.nodeselector"></a>
<a id="tocsv1.nodeselector"></a>

```json
{
  "nodeSelectorTerms": [
    {
      "matchExpressions": [
        {
          "key": "string",
          "operator": null,
          "values": [
            "string"
          ]
        }
      ],
      "matchFields": [
        {
          "key": "string",
          "operator": null,
          "values": [
            "string"
          ]
        }
      ]
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|nodeSelectorTerms|[[v1.NodeSelectorTerm](#schemav1.nodeselectorterm)]|false|none||Required. A list of node selector terms. The terms are ORed.|

<h2 id="tocS_v1.NodeAffinity">v1.NodeAffinity</h2>

<a id="schemav1.nodeaffinity"></a>
<a id="schema_v1.NodeAffinity"></a>
<a id="tocSv1.nodeaffinity"></a>
<a id="tocsv1.nodeaffinity"></a>

```json
{
  "preferredDuringSchedulingIgnoredDuringExecution": [
    {
      "preference": {
        "matchExpressions": [
          null
        ],
        "matchFields": [
          null
        ]
      },
      "weight": 0
    }
  ],
  "requiredDuringSchedulingIgnoredDuringExecution": {
    "nodeSelectorTerms": [
      {
        "matchExpressions": [
          null
        ],
        "matchFields": [
          null
        ]
      }
    ]
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|preferredDuringSchedulingIgnoredDuringExecution|[[v1.PreferredSchedulingTerm](#schemav1.preferredschedulingterm)]|false|none||The scheduler will prefer to schedule pods to nodes that satisfy<br />the affinity expressions specified by this field, but it may choose<br />a node that violates one or more of the expressions. The node that is<br />most preferred is the one with the greatest sum of weights, i.e.<br />for each node that meets all of the scheduling requirements (resource<br />request, requiredDuringScheduling affinity expressions, etc.),<br />compute a sum by iterating through the elements of this field and adding<br />"weight" to the sum if the node matches the corresponding matchExpressions; the<br />node(s) with the highest sum are the most preferred.<br />+optional|
|requiredDuringSchedulingIgnoredDuringExecution|[v1.NodeSelector](#schemav1.nodeselector)|false|none||If the affinity requirements specified by this field are not met at<br />scheduling time, the pod will not be scheduled onto the node.<br />If the affinity requirements specified by this field cease to be met<br />at some point during pod execution (e.g. due to an update), the system<br />may or may not try to eventually evict the pod from its node.<br />+optional|

<h2 id="tocS_v1.NFSVolumeSource">v1.NFSVolumeSource</h2>

<a id="schemav1.nfsvolumesource"></a>
<a id="schema_v1.NFSVolumeSource"></a>
<a id="tocSv1.nfsvolumesource"></a>
<a id="tocsv1.nfsvolumesource"></a>

```json
{
  "path": "string",
  "readOnly": true,
  "server": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|path|string|false|none||Path that is exported by the NFS server.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs|
|readOnly|boolean|false|none||ReadOnly here will force<br />the NFS export to be mounted with read-only permissions.<br />Defaults to false.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs<br />+optional|
|server|string|false|none||Server is the hostname or IP address of the NFS server.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs|

<h2 id="tocS_v1.MountPropagationMode">v1.MountPropagationMode</h2>

<a id="schemav1.mountpropagationmode"></a>
<a id="schema_v1.MountPropagationMode"></a>
<a id="tocSv1.mountpropagationmode"></a>
<a id="tocsv1.mountpropagationmode"></a>

```json
"None"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|None|
|*anonymous*|HostToContainer|
|*anonymous*|Bidirectional|

<h2 id="tocS_v1.ManagedFieldsOperationType">v1.ManagedFieldsOperationType</h2>

<a id="schemav1.managedfieldsoperationtype"></a>
<a id="schema_v1.ManagedFieldsOperationType"></a>
<a id="tocSv1.managedfieldsoperationtype"></a>
<a id="tocsv1.managedfieldsoperationtype"></a>

```json
"Apply"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Apply|
|*anonymous*|Update|

<h2 id="tocS_v1.ManagedFieldsEntry">v1.ManagedFieldsEntry</h2>

<a id="schemav1.managedfieldsentry"></a>
<a id="schema_v1.ManagedFieldsEntry"></a>
<a id="tocSv1.managedfieldsentry"></a>
<a id="tocsv1.managedfieldsentry"></a>

```json
{
  "apiVersion": "string",
  "fieldsType": "string",
  "fieldsV1": {},
  "manager": "string",
  "operation": "Apply",
  "subresource": "string",
  "time": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiVersion|string|false|none||APIVersion defines the version of this resource that this field set<br />applies to. The format is "group/version" just like the top-level<br />APIVersion field. It is necessary to track the version of a field<br />set because it cannot be automatically converted.|
|fieldsType|string|false|none||FieldsType is the discriminator for the different fields format and version.<br />There is currently only one possible value: "FieldsV1"|
|fieldsV1|[v1.FieldsV1](#schemav1.fieldsv1)|false|none||FieldsV1 holds the first JSON version format as described in the "FieldsV1" type.<br />+optional|
|manager|string|false|none||Manager is an identifier of the workflow managing these fields.|
|operation|[v1.ManagedFieldsOperationType](#schemav1.managedfieldsoperationtype)|false|none||Operation is the type of operation which lead to this ManagedFieldsEntry being created.<br />The only valid values for this field are 'Apply' and 'Update'.|
|subresource|string|false|none||Subresource is the name of the subresource used to update that object, or<br />empty string if the object was updated through the main resource. The<br />value of this field is used to distinguish between managers, even if they<br />share the same name. For example, a status update will be distinct from a<br />regular update using the same manager name.<br />Note that the APIVersion field is not related to the Subresource field and<br />it always corresponds to the version of the main resource.|
|time|string|false|none||Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'<br />+optional|

<h2 id="tocS_v1.LocalObjectReference">v1.LocalObjectReference</h2>

<a id="schemav1.localobjectreference"></a>
<a id="schema_v1.LocalObjectReference"></a>
<a id="tocSv1.localobjectreference"></a>
<a id="tocsv1.localobjectreference"></a>

```json
{
  "name": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|

<h2 id="tocS_v1.LoadBalancerStatus">v1.LoadBalancerStatus</h2>

<a id="schemav1.loadbalancerstatus"></a>
<a id="schema_v1.LoadBalancerStatus"></a>
<a id="tocSv1.loadbalancerstatus"></a>
<a id="tocsv1.loadbalancerstatus"></a>

```json
{
  "ingress": [
    {
      "hostname": "string",
      "ip": "string",
      "ports": [
        {
          "error": "string",
          "port": 0,
          "protocol": null
        }
      ]
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|ingress|[[v1.LoadBalancerIngress](#schemav1.loadbalanceringress)]|false|none||Ingress is a list containing ingress points for the load-balancer.<br />Traffic intended for the service should be sent to these ingress points.<br />+optional|

<h2 id="tocS_v1.LoadBalancerIngress">v1.LoadBalancerIngress</h2>

<a id="schemav1.loadbalanceringress"></a>
<a id="schema_v1.LoadBalancerIngress"></a>
<a id="tocSv1.loadbalanceringress"></a>
<a id="tocsv1.loadbalanceringress"></a>

```json
{
  "hostname": "string",
  "ip": "string",
  "ports": [
    {
      "error": "string",
      "port": 0,
      "protocol": "TCP"
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|hostname|string|false|none||Hostname is set for load-balancer ingress points that are DNS based<br />(typically AWS load-balancers)<br />+optional|
|ip|string|false|none||IP is set for load-balancer ingress points that are IP based<br />(typically GCE or OpenStack load-balancers)<br />+optional|
|ports|[[v1.PortStatus](#schemav1.portstatus)]|false|none||Ports is a list of records of service ports<br />If used, every port defined in the service should have an entry in it<br />+listType=atomic<br />+optional|

<h2 id="tocS_v1.Lifecycle">v1.Lifecycle</h2>

<a id="schemav1.lifecycle"></a>
<a id="schema_v1.Lifecycle"></a>
<a id="tocSv1.lifecycle"></a>
<a id="tocsv1.lifecycle"></a>

```json
{
  "postStart": {
    "exec": {
      "command": null
    },
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "tcpSocket": {
      "host": null,
      "port": null
    }
  },
  "preStop": {
    "exec": {
      "command": null
    },
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "tcpSocket": {
      "host": null,
      "port": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|postStart|[v1.Handler](#schemav1.handler)|false|none||PostStart is called immediately after a container is created. If the handler fails,<br />the container is terminated and restarted according to its restart policy.<br />Other management of the container blocks until the hook completes.<br />More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks<br />+optional|
|preStop|[v1.Handler](#schemav1.handler)|false|none||PreStop is called immediately before a container is terminated due to an<br />API request or management event such as liveness/startup probe failure,<br />preemption, resource contention, etc. The handler is not called if the<br />container crashes or exits. The reason for termination is passed to the<br />handler. The Pod's termination grace period countdown begins before the<br />PreStop hooked is executed. Regardless of the outcome of the handler, the<br />container will eventually terminate within the Pod's termination grace<br />period. Other management of the container blocks until the hook completes<br />or until the termination grace period is reached.<br />More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks<br />+optional|

<h2 id="tocS_v1.LabelSelectorRequirement">v1.LabelSelectorRequirement</h2>

<a id="schemav1.labelselectorrequirement"></a>
<a id="schema_v1.LabelSelectorRequirement"></a>
<a id="tocSv1.labelselectorrequirement"></a>
<a id="tocsv1.labelselectorrequirement"></a>

```json
{
  "key": "string",
  "operator": "In",
  "values": [
    "string"
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|key|string|false|none||key is the label key that the selector applies to.<br />+patchMergeKey=key<br />+patchStrategy=merge|
|operator|[v1.LabelSelectorOperator](#schemav1.labelselectoroperator)|false|none||operator represents a key's relationship to a set of values.<br />Valid operators are In, NotIn, Exists and DoesNotExist.|
|values|[string]|false|none||values is an array of string values. If the operator is In or NotIn,<br />the values array must be non-empty. If the operator is Exists or DoesNotExist,<br />the values array must be empty. This array is replaced during a strategic<br />merge patch.<br />+optional|

<h2 id="tocS_v1.LabelSelectorOperator">v1.LabelSelectorOperator</h2>

<a id="schemav1.labelselectoroperator"></a>
<a id="schema_v1.LabelSelectorOperator"></a>
<a id="tocSv1.labelselectoroperator"></a>
<a id="tocsv1.labelselectoroperator"></a>

```json
"In"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|In|
|*anonymous*|NotIn|
|*anonymous*|Exists|
|*anonymous*|DoesNotExist|

<h2 id="tocS_v1.LabelSelector">v1.LabelSelector</h2>

<a id="schemav1.labelselector"></a>
<a id="schema_v1.LabelSelector"></a>
<a id="tocSv1.labelselector"></a>
<a id="tocsv1.labelselector"></a>

```json
{
  "matchExpressions": [
    {
      "key": "string",
      "operator": "In",
      "values": [
        "string"
      ]
    }
  ],
  "matchLabels": {
    "property1": "string",
    "property2": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|matchExpressions|[[v1.LabelSelectorRequirement](#schemav1.labelselectorrequirement)]|false|none||matchExpressions is a list of label selector requirements. The requirements are ANDed.<br />+optional|
|matchLabels|object|false|none||matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels<br />map is equivalent to an element of matchExpressions, whose key field is "key", the<br />operator is "In", and the values array contains only "value". The requirements are ANDed.<br />+optional|
|» **additionalProperties**|string|false|none||none|

<h2 id="tocS_v1.KeyToPath">v1.KeyToPath</h2>

<a id="schemav1.keytopath"></a>
<a id="schema_v1.KeyToPath"></a>
<a id="tocSv1.keytopath"></a>
<a id="tocsv1.keytopath"></a>

```json
{
  "key": "string",
  "mode": 0,
  "path": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|key|string|false|none||The key to project.|
|mode|integer|false|none||Optional: mode bits used to set permissions on this file.<br />Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br />YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br />If not specified, the volume defaultMode will be used.<br />This might be in conflict with other options that affect the file<br />mode, like fsGroup, and the result can be other mode bits set.<br />+optional|
|path|string|false|none||The relative path of the file to map the key to.<br />May not be an absolute path.<br />May not contain the path element '..'.<br />May not start with the string '..'.|

<h2 id="tocS_v1.ISCSIVolumeSource">v1.ISCSIVolumeSource</h2>

<a id="schemav1.iscsivolumesource"></a>
<a id="schema_v1.ISCSIVolumeSource"></a>
<a id="tocSv1.iscsivolumesource"></a>
<a id="tocsv1.iscsivolumesource"></a>

```json
{
  "chapAuthDiscovery": true,
  "chapAuthSession": true,
  "fsType": "string",
  "initiatorName": "string",
  "iqn": "string",
  "iscsiInterface": "string",
  "lun": 0,
  "portals": [
    "string"
  ],
  "readOnly": true,
  "secretRef": {
    "name": "string"
  },
  "targetPortal": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|chapAuthDiscovery|boolean|false|none||whether support iSCSI Discovery CHAP authentication<br />+optional|
|chapAuthSession|boolean|false|none||whether support iSCSI Session CHAP authentication<br />+optional|
|fsType|string|false|none||Filesystem type of the volume that you want to mount.<br />Tip: Ensure that the filesystem type is supported by the host operating system.<br />Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi<br />TODO: how do we prevent errors in the filesystem from compromising the machine<br />+optional|
|initiatorName|string|false|none||Custom iSCSI Initiator Name.<br />If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<br /><target portal>:<volume name> will be created for the connection.<br />+optional|
|iqn|string|false|none||Target iSCSI Qualified Name.|
|iscsiInterface|string|false|none||iSCSI Interface Name that uses an iSCSI transport.<br />Defaults to 'default' (tcp).<br />+optional|
|lun|integer|false|none||iSCSI Target Lun number.|
|portals|[string]|false|none||iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port<br />is other than default (typically TCP ports 860 and 3260).<br />+optional|
|readOnly|boolean|false|none||ReadOnly here will force the ReadOnly setting in VolumeMounts.<br />Defaults to false.<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||CHAP Secret for iSCSI target and initiator authentication<br />+optional|
|targetPortal|string|false|none||iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port<br />is other than default (typically TCP ports 860 and 3260).|

<h2 id="tocS_v1.IPFamilyPolicyType">v1.IPFamilyPolicyType</h2>

<a id="schemav1.ipfamilypolicytype"></a>
<a id="schema_v1.IPFamilyPolicyType"></a>
<a id="tocSv1.ipfamilypolicytype"></a>
<a id="tocsv1.ipfamilypolicytype"></a>

```json
"SingleStack"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|SingleStack|
|*anonymous*|PreferDualStack|
|*anonymous*|RequireDualStack|

<h2 id="tocS_v1.IPFamily">v1.IPFamily</h2>

<a id="schemav1.ipfamily"></a>
<a id="schema_v1.IPFamily"></a>
<a id="tocSv1.ipfamily"></a>
<a id="tocsv1.ipfamily"></a>

```json
"IPv4"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|IPv4|
|*anonymous*|IPv6|

<h2 id="tocS_v1.HostPathVolumeSource">v1.HostPathVolumeSource</h2>

<a id="schemav1.hostpathvolumesource"></a>
<a id="schema_v1.HostPathVolumeSource"></a>
<a id="tocSv1.hostpathvolumesource"></a>
<a id="tocsv1.hostpathvolumesource"></a>

```json
{
  "path": "string",
  "type": ""
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|path|string|false|none||Path of the directory on the host.<br />If the path is a symlink, it will follow the link to the real path.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath|
|type|[v1.HostPathType](#schemav1.hostpathtype)|false|none||Type for HostPath Volume<br />Defaults to ""<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath<br />+optional|

<h2 id="tocS_v1.HostPathType">v1.HostPathType</h2>

<a id="schemav1.hostpathtype"></a>
<a id="schema_v1.HostPathType"></a>
<a id="tocSv1.hostpathtype"></a>
<a id="tocsv1.hostpathtype"></a>

```json
""

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*||
|*anonymous*|DirectoryOrCreate|
|*anonymous*|Directory|
|*anonymous*|FileOrCreate|
|*anonymous*|File|
|*anonymous*|Socket|
|*anonymous*|CharDevice|
|*anonymous*|BlockDevice|

<h2 id="tocS_v1.HostAlias">v1.HostAlias</h2>

<a id="schemav1.hostalias"></a>
<a id="schema_v1.HostAlias"></a>
<a id="tocSv1.hostalias"></a>
<a id="tocsv1.hostalias"></a>

```json
{
  "hostnames": [
    "string"
  ],
  "ip": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|hostnames|[string]|false|none||Hostnames for the above IP address.|
|ip|string|false|none||IP address of the host file entry.|

<h2 id="tocS_v1.Handler">v1.Handler</h2>

<a id="schemav1.handler"></a>
<a id="schema_v1.Handler"></a>
<a id="tocSv1.handler"></a>
<a id="tocsv1.handler"></a>

```json
{
  "exec": {
    "command": [
      "string"
    ]
  },
  "httpGet": {
    "host": "string",
    "httpHeaders": [
      {
        "name": "string",
        "value": "string"
      }
    ],
    "path": "string",
    "port": {
      "intVal": null,
      "strVal": null,
      "type": null
    },
    "scheme": "HTTP"
  },
  "tcpSocket": {
    "host": "string",
    "port": {
      "intVal": null,
      "strVal": null,
      "type": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|exec|[v1.ExecAction](#schemav1.execaction)|false|none||One and only one of the following should be specified.<br />Exec specifies the action to take.<br />+optional|
|httpGet|[v1.HTTPGetAction](#schemav1.httpgetaction)|false|none||HTTPGet specifies the http request to perform.<br />+optional|
|tcpSocket|[v1.TCPSocketAction](#schemav1.tcpsocketaction)|false|none||TCPSocket specifies an action involving a TCP port.<br />TCP hooks not yet supported<br />TODO: implement a realistic TCP lifecycle hook<br />+optional|

<h2 id="tocS_v1.HTTPHeader">v1.HTTPHeader</h2>

<a id="schemav1.httpheader"></a>
<a id="schema_v1.HTTPHeader"></a>
<a id="tocSv1.httpheader"></a>
<a id="tocsv1.httpheader"></a>

```json
{
  "name": "string",
  "value": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||The header field name|
|value|string|false|none||The header field value|

<h2 id="tocS_v1.HTTPGetAction">v1.HTTPGetAction</h2>

<a id="schemav1.httpgetaction"></a>
<a id="schema_v1.HTTPGetAction"></a>
<a id="tocSv1.httpgetaction"></a>
<a id="tocsv1.httpgetaction"></a>

```json
{
  "host": "string",
  "httpHeaders": [
    {
      "name": "string",
      "value": "string"
    }
  ],
  "path": "string",
  "port": {
    "intVal": 0,
    "strVal": "string",
    "type": 0
  },
  "scheme": "HTTP"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|host|string|false|none||Host name to connect to, defaults to the pod IP. You probably want to set<br />"Host" in httpHeaders instead.<br />+optional|
|httpHeaders|[[v1.HTTPHeader](#schemav1.httpheader)]|false|none||Custom headers to set in the request. HTTP allows repeated headers.<br />+optional|
|path|string|false|none||Path to access on the HTTP server.<br />+optional|
|port|[intstr.IntOrString](#schemaintstr.intorstring)|false|none||Name or number of the port to access on the container.<br />Number must be in the range 1 to 65535.<br />Name must be an IANA_SVC_NAME.|
|scheme|[v1.URIScheme](#schemav1.urischeme)|false|none||Scheme to use for connecting to the host.<br />Defaults to HTTP.<br />+optional|

<h2 id="tocS_v1.GlusterfsVolumeSource">v1.GlusterfsVolumeSource</h2>

<a id="schemav1.glusterfsvolumesource"></a>
<a id="schema_v1.GlusterfsVolumeSource"></a>
<a id="tocSv1.glusterfsvolumesource"></a>
<a id="tocsv1.glusterfsvolumesource"></a>

```json
{
  "endpoints": "string",
  "path": "string",
  "readOnly": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|endpoints|string|false|none||EndpointsName is the endpoint name that details Glusterfs topology.<br />More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod|
|path|string|false|none||Path is the Glusterfs volume path.<br />More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod|
|readOnly|boolean|false|none||ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions.<br />Defaults to false.<br />More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod<br />+optional|

<h2 id="tocS_v1.GitRepoVolumeSource">v1.GitRepoVolumeSource</h2>

<a id="schemav1.gitrepovolumesource"></a>
<a id="schema_v1.GitRepoVolumeSource"></a>
<a id="tocSv1.gitrepovolumesource"></a>
<a id="tocsv1.gitrepovolumesource"></a>

```json
{
  "directory": "string",
  "repository": "string",
  "revision": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|directory|string|false|none||Target directory name.<br />Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the<br />git repository.  Otherwise, if specified, the volume will contain the git repository in<br />the subdirectory with the given name.<br />+optional|
|repository|string|false|none||Repository URL|
|revision|string|false|none||Commit hash for the specified revision.<br />+optional|

<h2 id="tocS_v1.GCEPersistentDiskVolumeSource">v1.GCEPersistentDiskVolumeSource</h2>

<a id="schemav1.gcepersistentdiskvolumesource"></a>
<a id="schema_v1.GCEPersistentDiskVolumeSource"></a>
<a id="tocSv1.gcepersistentdiskvolumesource"></a>
<a id="tocsv1.gcepersistentdiskvolumesource"></a>

```json
{
  "fsType": "string",
  "partition": 0,
  "pdName": "string",
  "readOnly": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type of the volume that you want to mount.<br />Tip: Ensure that the filesystem type is supported by the host operating system.<br />Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br />TODO: how do we prevent errors in the filesystem from compromising the machine<br />+optional|
|partition|integer|false|none||The partition in the volume that you want to mount.<br />If omitted, the default is to mount by volume name.<br />Examples: For volume /dev/sda1, you specify the partition as "1".<br />Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br />+optional|
|pdName|string|false|none||Unique name of the PD resource in GCE. Used to identify the disk in GCE.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk|
|readOnly|boolean|false|none||ReadOnly here will force the ReadOnly setting in VolumeMounts.<br />Defaults to false.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br />+optional|

<h2 id="tocS_v1.FlockerVolumeSource">v1.FlockerVolumeSource</h2>

<a id="schemav1.flockervolumesource"></a>
<a id="schema_v1.FlockerVolumeSource"></a>
<a id="tocSv1.flockervolumesource"></a>
<a id="tocsv1.flockervolumesource"></a>

```json
{
  "datasetName": "string",
  "datasetUUID": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|datasetName|string|false|none||Name of the dataset stored as metadata -> name on the dataset for Flocker<br />should be considered as deprecated<br />+optional|
|datasetUUID|string|false|none||UUID of the dataset. This is unique identifier of a Flocker dataset<br />+optional|

<h2 id="tocS_v1.FlexVolumeSource">v1.FlexVolumeSource</h2>

<a id="schemav1.flexvolumesource"></a>
<a id="schema_v1.FlexVolumeSource"></a>
<a id="tocSv1.flexvolumesource"></a>
<a id="tocsv1.flexvolumesource"></a>

```json
{
  "driver": "string",
  "fsType": "string",
  "options": {
    "property1": "string",
    "property2": "string"
  },
  "readOnly": true,
  "secretRef": {
    "name": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|driver|string|false|none||Driver is the name of the driver to use for this volume.|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.<br />+optional|
|options|object|false|none||Optional: Extra command options if any.<br />+optional|
|» **additionalProperties**|string|false|none||none|
|readOnly|boolean|false|none||Optional: Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||Optional: SecretRef is reference to the secret object containing<br />sensitive information to pass to the plugin scripts. This may be<br />empty if no secret object is specified. If the secret object<br />contains more than one secret, all secrets are passed to the plugin<br />scripts.<br />+optional|

<h2 id="tocS_v1.FieldsV1">v1.FieldsV1</h2>

<a id="schemav1.fieldsv1"></a>
<a id="schema_v1.FieldsV1"></a>
<a id="tocSv1.fieldsv1"></a>
<a id="tocsv1.fieldsv1"></a>

```json
{}

```

### attribute

*None*

<h2 id="tocS_v1.FCVolumeSource">v1.FCVolumeSource</h2>

<a id="schemav1.fcvolumesource"></a>
<a id="schema_v1.FCVolumeSource"></a>
<a id="tocSv1.fcvolumesource"></a>
<a id="tocsv1.fcvolumesource"></a>

```json
{
  "fsType": "string",
  "lun": 0,
  "readOnly": true,
  "targetWWNs": [
    "string"
  ],
  "wwids": [
    "string"
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />TODO: how do we prevent errors in the filesystem from compromising the machine<br />+optional|
|lun|integer|false|none||Optional: FC target lun number<br />+optional|
|readOnly|boolean|false|none||Optional: Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|
|targetWWNs|[string]|false|none||Optional: FC target worldwide names (WWNs)<br />+optional|
|wwids|[string]|false|none||Optional: FC volume world wide identifiers (wwids)<br />Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.<br />+optional|

<h2 id="tocS_v1.ExecAction">v1.ExecAction</h2>

<a id="schemav1.execaction"></a>
<a id="schema_v1.ExecAction"></a>
<a id="tocSv1.execaction"></a>
<a id="tocsv1.execaction"></a>

```json
{
  "command": [
    "string"
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|command|[string]|false|none||Command is the command line to execute inside the container, the working directory for the<br />command  is root ('/') in the container's filesystem. The command is simply exec'd, it is<br />not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use<br />a shell, you need to explicitly call out to that shell.<br />Exit status of 0 is treated as live/healthy and non-zero is unhealthy.<br />+optional|

<h2 id="tocS_v1.EphemeralVolumeSource">v1.EphemeralVolumeSource</h2>

<a id="schemav1.ephemeralvolumesource"></a>
<a id="schema_v1.EphemeralVolumeSource"></a>
<a id="tocSv1.ephemeralvolumesource"></a>
<a id="tocsv1.ephemeralvolumesource"></a>

```json
{
  "volumeClaimTemplate": {
    "metadata": {
      "annotations": null,
      "clusterName": null,
      "creationTimestamp": null,
      "deletionGracePeriodSeconds": null,
      "deletionTimestamp": null,
      "finalizers": null,
      "generateName": null,
      "generation": null,
      "labels": null,
      "managedFields": null,
      "name": null,
      "namespace": null,
      "ownerReferences": null,
      "resourceVersion": null,
      "selfLink": null,
      "uid": null
    },
    "spec": {
      "accessModes": null,
      "dataSource": null,
      "dataSourceRef": null,
      "resources": null,
      "selector": null,
      "storageClassName": null,
      "volumeMode": null,
      "volumeName": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|volumeClaimTemplate|[v1.PersistentVolumeClaimTemplate](#schemav1.persistentvolumeclaimtemplate)|false|none||Will be used to create a stand-alone PVC to provision the volume.<br />The pod in which this EphemeralVolumeSource is embedded will be the<br />owner of the PVC, i.e. the PVC will be deleted together with the<br />pod.  The name of the PVC will be `<pod name>-<volume name>` where<br />`<volume name>` is the name from the `PodSpec.Volumes` array<br />entry. Pod validation will reject the pod if the concatenated name<br />is not valid for a PVC (for example, too long).<br /><br />An existing PVC with that name that is not owned by the pod<br />will *not* be used for the pod to avoid using an unrelated<br />volume by mistake. Starting the pod is then blocked until<br />the unrelated PVC is removed. If such a pre-created PVC is<br />meant to be used by the pod, the PVC has to updated with an<br />owner reference to the pod once the pod exists. Normally<br />this should not be necessary, but it may be useful when<br />manually reconstructing a broken cluster.<br /><br />This field is read-only and no changes will be made by Kubernetes<br />to the PVC after it has been created.<br /><br />Required, must not be nil.|

<h2 id="tocS_v1.EphemeralContainer">v1.EphemeralContainer</h2>

<a id="schemav1.ephemeralcontainer"></a>
<a id="schema_v1.EphemeralContainer"></a>
<a id="tocSv1.ephemeralcontainer"></a>
<a id="tocsv1.ephemeralcontainer"></a>

```json
{
  "args": [
    "string"
  ],
  "command": [
    "string"
  ],
  "env": [
    {
      "name": "string",
      "value": "string",
      "valueFrom": {
        "configMapKeyRef": null,
        "fieldRef": null,
        "resourceFieldRef": null,
        "secretKeyRef": null
      }
    }
  ],
  "envFrom": [
    {
      "configMapRef": {
        "name": "string",
        "optional": true
      },
      "prefix": "string",
      "secretRef": {
        "name": "string",
        "optional": true
      }
    }
  ],
  "image": "string",
  "imagePullPolicy": "Always",
  "lifecycle": {
    "postStart": {
      "exec": null,
      "httpGet": null,
      "tcpSocket": null
    },
    "preStop": {
      "exec": null,
      "httpGet": null,
      "tcpSocket": null
    }
  },
  "livenessProbe": {
    "exec": {
      "command": null
    },
    "failureThreshold": 0,
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "initialDelaySeconds": 0,
    "periodSeconds": 0,
    "successThreshold": 0,
    "tcpSocket": {
      "host": null,
      "port": null
    },
    "terminationGracePeriodSeconds": 0,
    "timeoutSeconds": 0
  },
  "name": "string",
  "ports": [
    {
      "containerPort": 0,
      "hostIP": "string",
      "hostPort": 0,
      "name": "string",
      "protocol": "TCP"
    }
  ],
  "readinessProbe": {
    "exec": {
      "command": null
    },
    "failureThreshold": 0,
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "initialDelaySeconds": 0,
    "periodSeconds": 0,
    "successThreshold": 0,
    "tcpSocket": {
      "host": null,
      "port": null
    },
    "terminationGracePeriodSeconds": 0,
    "timeoutSeconds": 0
  },
  "resources": {
    "limits": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    },
    "requests": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    }
  },
  "securityContext": {
    "allowPrivilegeEscalation": true,
    "capabilities": {
      "add": null,
      "drop": null
    },
    "privileged": true,
    "procMount": "Default",
    "readOnlyRootFilesystem": true,
    "runAsGroup": 0,
    "runAsNonRoot": true,
    "runAsUser": 0,
    "seLinuxOptions": {
      "level": null,
      "role": null,
      "type": null,
      "user": null
    },
    "seccompProfile": {
      "localhostProfile": null,
      "type": null
    },
    "windowsOptions": {
      "gmsaCredentialSpec": null,
      "gmsaCredentialSpecName": null,
      "hostProcess": null,
      "runAsUserName": null
    }
  },
  "startupProbe": {
    "exec": {
      "command": null
    },
    "failureThreshold": 0,
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "initialDelaySeconds": 0,
    "periodSeconds": 0,
    "successThreshold": 0,
    "tcpSocket": {
      "host": null,
      "port": null
    },
    "terminationGracePeriodSeconds": 0,
    "timeoutSeconds": 0
  },
  "stdin": true,
  "stdinOnce": true,
  "targetContainerName": "string",
  "terminationMessagePath": "string",
  "terminationMessagePolicy": "File",
  "tty": true,
  "volumeDevices": [
    {
      "devicePath": "string",
      "name": "string"
    }
  ],
  "volumeMounts": [
    {
      "mountPath": "string",
      "mountPropagation": "None",
      "name": "string",
      "readOnly": true,
      "subPath": "string",
      "subPathExpr": "string"
    }
  ],
  "workingDir": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|args|[string]|false|none||Arguments to the entrypoint.<br />The docker image's CMD is used if this is not provided.<br />Variable references $(VAR_NAME) are expanded using the container's environment. If a variable<br />cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced<br />to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will<br />produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless<br />of whether the variable exists or not. Cannot be updated.<br />More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell<br />+optional|
|command|[string]|false|none||Entrypoint array. Not executed within a shell.<br />The docker image's ENTRYPOINT is used if this is not provided.<br />Variable references $(VAR_NAME) are expanded using the container's environment. If a variable<br />cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced<br />to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will<br />produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless<br />of whether the variable exists or not. Cannot be updated.<br />More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell<br />+optional|
|env|[[v1.EnvVar](#schemav1.envvar)]|false|none||List of environment variables to set in the container.<br />Cannot be updated.<br />+optional<br />+patchMergeKey=name<br />+patchStrategy=merge|
|envFrom|[[v1.EnvFromSource](#schemav1.envfromsource)]|false|none||List of sources to populate environment variables in the container.<br />The keys defined within a source must be a C_IDENTIFIER. All invalid keys<br />will be reported as an event when the container is starting. When a key exists in multiple<br />sources, the value associated with the last source will take precedence.<br />Values defined by an Env with a duplicate key will take precedence.<br />Cannot be updated.<br />+optional|
|image|string|false|none||Docker image name.<br />More info: https://kubernetes.io/docs/concepts/containers/images|
|imagePullPolicy|[v1.PullPolicy](#schemav1.pullpolicy)|false|none||Image pull policy.<br />One of Always, Never, IfNotPresent.<br />Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.<br />Cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/containers/images#updating-images<br />+optional|
|lifecycle|[v1.Lifecycle](#schemav1.lifecycle)|false|none||Lifecycle is not allowed for ephemeral containers.<br />+optional|
|livenessProbe|[v1.Probe](#schemav1.probe)|false|none||Probes are not allowed for ephemeral containers.<br />+optional|
|name|string|false|none||Name of the ephemeral container specified as a DNS_LABEL.<br />This name must be unique among all containers, init containers and ephemeral containers.|
|ports|[[v1.ContainerPort](#schemav1.containerport)]|false|none||Ports are not allowed for ephemeral containers.|
|readinessProbe|[v1.Probe](#schemav1.probe)|false|none||Probes are not allowed for ephemeral containers.<br />+optional|
|resources|[v1.ResourceRequirements](#schemav1.resourcerequirements)|false|none||Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources<br />already allocated to the pod.<br />+optional|
|securityContext|[v1.SecurityContext](#schemav1.securitycontext)|false|none||Optional: SecurityContext defines the security options the ephemeral container should be run with.<br />If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.<br />+optional|
|startupProbe|[v1.Probe](#schemav1.probe)|false|none||Probes are not allowed for ephemeral containers.<br />+optional|
|stdin|boolean|false|none||Whether this container should allocate a buffer for stdin in the container runtime. If this<br />is not set, reads from stdin in the container will always result in EOF.<br />Default is false.<br />+optional|
|stdinOnce|boolean|false|none||Whether the container runtime should close the stdin channel after it has been opened by<br />a single attach. When stdin is true the stdin stream will remain open across multiple attach<br />sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the<br />first client attaches to stdin, and then remains open and accepts data until the client disconnects,<br />at which time stdin is closed and remains closed until the container is restarted. If this<br />flag is false, a container processes that reads from stdin will never receive an EOF.<br />Default is false<br />+optional|
|targetContainerName|string|false|none||If set, the name of the container from PodSpec that this ephemeral container targets.<br />The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container.<br />If not set then the ephemeral container is run in whatever namespaces are shared<br />for the pod. Note that the container runtime must support this feature.<br />+optional|
|terminationMessagePath|string|false|none||Optional: Path at which the file to which the container's termination message<br />will be written is mounted into the container's filesystem.<br />Message written is intended to be brief final status, such as an assertion failure message.<br />Will be truncated by the node if greater than 4096 bytes. The total message length across<br />all containers will be limited to 12kb.<br />Defaults to /dev/termination-log.<br />Cannot be updated.<br />+optional|
|terminationMessagePolicy|[v1.TerminationMessagePolicy](#schemav1.terminationmessagepolicy)|false|none||Indicate how the termination message should be populated. File will use the contents of<br />terminationMessagePath to populate the container status message on both success and failure.<br />FallbackToLogsOnError will use the last chunk of container log output if the termination<br />message file is empty and the container exited with an error.<br />The log output is limited to 2048 bytes or 80 lines, whichever is smaller.<br />Defaults to File.<br />Cannot be updated.<br />+optional|
|tty|boolean|false|none||Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.<br />Default is false.<br />+optional|
|volumeDevices|[[v1.VolumeDevice](#schemav1.volumedevice)]|false|none||volumeDevices is the list of block devices to be used by the container.<br />+patchMergeKey=devicePath<br />+patchStrategy=merge<br />+optional|
|volumeMounts|[[v1.VolumeMount](#schemav1.volumemount)]|false|none||Pod volumes to mount into the container's filesystem.<br />Cannot be updated.<br />+optional<br />+patchMergeKey=mountPath<br />+patchStrategy=merge|
|workingDir|string|false|none||Container's working directory.<br />If not specified, the container runtime's default will be used, which<br />might be configured in the container image.<br />Cannot be updated.<br />+optional|

<h2 id="tocS_v1.EnvVarSource">v1.EnvVarSource</h2>

<a id="schemav1.envvarsource"></a>
<a id="schema_v1.EnvVarSource"></a>
<a id="tocSv1.envvarsource"></a>
<a id="tocsv1.envvarsource"></a>

```json
{
  "configMapKeyRef": {
    "key": "string",
    "name": "string",
    "optional": true
  },
  "fieldRef": {
    "apiVersion": "string",
    "fieldPath": "string"
  },
  "resourceFieldRef": {
    "containerName": "string",
    "divisor": {
      "Format": null
    },
    "resource": "string"
  },
  "secretKeyRef": {
    "key": "string",
    "name": "string",
    "optional": true
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|configMapKeyRef|[v1.ConfigMapKeySelector](#schemav1.configmapkeyselector)|false|none||Selects a key of a ConfigMap.<br />+optional|
|fieldRef|[v1.ObjectFieldSelector](#schemav1.objectfieldselector)|false|none||Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`,<br />spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.<br />+optional|
|resourceFieldRef|[v1.ResourceFieldSelector](#schemav1.resourcefieldselector)|false|none||Selects a resource of the container: only resources limits and requests<br />(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.<br />+optional|
|secretKeyRef|[v1.SecretKeySelector](#schemav1.secretkeyselector)|false|none||Selects a key of a secret in the pod's namespace<br />+optional|

<h2 id="tocS_v1.EnvVar">v1.EnvVar</h2>

<a id="schemav1.envvar"></a>
<a id="schema_v1.EnvVar"></a>
<a id="tocSv1.envvar"></a>
<a id="tocsv1.envvar"></a>

```json
{
  "name": "string",
  "value": "string",
  "valueFrom": {
    "configMapKeyRef": {
      "key": null,
      "name": null,
      "optional": null
    },
    "fieldRef": {
      "apiVersion": null,
      "fieldPath": null
    },
    "resourceFieldRef": {
      "containerName": null,
      "divisor": null,
      "resource": null
    },
    "secretKeyRef": {
      "key": null,
      "name": null,
      "optional": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Name of the environment variable. Must be a C_IDENTIFIER.|
|value|string|false|none||Variable references $(VAR_NAME) are expanded<br />using the previously defined environment variables in the container and<br />any service environment variables. If a variable cannot be resolved,<br />the reference in the input string will be unchanged. Double $$ are reduced<br />to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.<br />"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".<br />Escaped references will never be expanded, regardless of whether the variable<br />exists or not.<br />Defaults to "".<br />+optional|
|valueFrom|[v1.EnvVarSource](#schemav1.envvarsource)|false|none||Source for the environment variable's value. Cannot be used if value is not empty.<br />+optional|

<h2 id="tocS_v1.EnvFromSource">v1.EnvFromSource</h2>

<a id="schemav1.envfromsource"></a>
<a id="schema_v1.EnvFromSource"></a>
<a id="tocSv1.envfromsource"></a>
<a id="tocsv1.envfromsource"></a>

```json
{
  "configMapRef": {
    "name": "string",
    "optional": true
  },
  "prefix": "string",
  "secretRef": {
    "name": "string",
    "optional": true
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|configMapRef|[v1.ConfigMapEnvSource](#schemav1.configmapenvsource)|false|none||The ConfigMap to select from<br />+optional|
|prefix|string|false|none||An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.<br />+optional|
|secretRef|[v1.SecretEnvSource](#schemav1.secretenvsource)|false|none||The Secret to select from<br />+optional|

<h2 id="tocS_v1.EmptyDirVolumeSource">v1.EmptyDirVolumeSource</h2>

<a id="schemav1.emptydirvolumesource"></a>
<a id="schema_v1.EmptyDirVolumeSource"></a>
<a id="tocSv1.emptydirvolumesource"></a>
<a id="tocsv1.emptydirvolumesource"></a>

```json
{
  "medium": "",
  "sizeLimit": {
    "Format": "DecimalExponent"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|medium|[v1.StorageMedium](#schemav1.storagemedium)|false|none||What type of storage medium should back this directory.<br />The default is "" which means to use the node's default medium.<br />Must be an empty string (default) or Memory.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir<br />+optional|
|sizeLimit|[resource.Quantity](#schemaresource.quantity)|false|none||Total amount of local storage required for this EmptyDir volume.<br />The size limit is also applicable for memory medium.<br />The maximum usage on memory medium EmptyDir would be the minimum value between<br />the SizeLimit specified here and the sum of memory limits of all containers in a pod.<br />The default is nil which means that the limit is undefined.<br />More info: http://kubernetes.io/docs/user-guide/volumes#emptydir<br />+optional|

<h2 id="tocS_v1.DownwardAPIVolumeSource">v1.DownwardAPIVolumeSource</h2>

<a id="schemav1.downwardapivolumesource"></a>
<a id="schema_v1.DownwardAPIVolumeSource"></a>
<a id="tocSv1.downwardapivolumesource"></a>
<a id="tocsv1.downwardapivolumesource"></a>

```json
{
  "defaultMode": 0,
  "items": [
    {
      "fieldRef": {
        "apiVersion": "string",
        "fieldPath": "string"
      },
      "mode": 0,
      "path": "string",
      "resourceFieldRef": {
        "containerName": "string",
        "divisor": null,
        "resource": "string"
      }
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|defaultMode|integer|false|none||Optional: mode bits to use on created files by default. Must be a<br />Optional: mode bits used to set permissions on created files by default.<br />Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br />YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br />Defaults to 0644.<br />Directories within the path are not affected by this setting.<br />This might be in conflict with other options that affect the file<br />mode, like fsGroup, and the result can be other mode bits set.<br />+optional|
|items|[[v1.DownwardAPIVolumeFile](#schemav1.downwardapivolumefile)]|false|none||Items is a list of downward API volume file<br />+optional|

<h2 id="tocS_v1.DownwardAPIVolumeFile">v1.DownwardAPIVolumeFile</h2>

<a id="schemav1.downwardapivolumefile"></a>
<a id="schema_v1.DownwardAPIVolumeFile"></a>
<a id="tocSv1.downwardapivolumefile"></a>
<a id="tocsv1.downwardapivolumefile"></a>

```json
{
  "fieldRef": {
    "apiVersion": "string",
    "fieldPath": "string"
  },
  "mode": 0,
  "path": "string",
  "resourceFieldRef": {
    "containerName": "string",
    "divisor": {
      "Format": null
    },
    "resource": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fieldRef|[v1.ObjectFieldSelector](#schemav1.objectfieldselector)|false|none||Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.<br />+optional|
|mode|integer|false|none||Optional: mode bits used to set permissions on this file, must be an octal value<br />between 0000 and 0777 or a decimal value between 0 and 511.<br />YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br />If not specified, the volume defaultMode will be used.<br />This might be in conflict with other options that affect the file<br />mode, like fsGroup, and the result can be other mode bits set.<br />+optional|
|path|string|false|none||Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'|
|resourceFieldRef|[v1.ResourceFieldSelector](#schemav1.resourcefieldselector)|false|none||Selects a resource of the container: only resources limits and requests<br />(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.<br />+optional|

<h2 id="tocS_v1.DownwardAPIProjection">v1.DownwardAPIProjection</h2>

<a id="schemav1.downwardapiprojection"></a>
<a id="schema_v1.DownwardAPIProjection"></a>
<a id="tocSv1.downwardapiprojection"></a>
<a id="tocsv1.downwardapiprojection"></a>

```json
{
  "items": [
    {
      "fieldRef": {
        "apiVersion": "string",
        "fieldPath": "string"
      },
      "mode": 0,
      "path": "string",
      "resourceFieldRef": {
        "containerName": "string",
        "divisor": null,
        "resource": "string"
      }
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|items|[[v1.DownwardAPIVolumeFile](#schemav1.downwardapivolumefile)]|false|none||Items is a list of DownwardAPIVolume file<br />+optional|

<h2 id="tocS_v1.DeploymentStrategyType">v1.DeploymentStrategyType</h2>

<a id="schemav1.deploymentstrategytype"></a>
<a id="schema_v1.DeploymentStrategyType"></a>
<a id="tocSv1.deploymentstrategytype"></a>
<a id="tocsv1.deploymentstrategytype"></a>

```json
"Recreate"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Recreate|
|*anonymous*|RollingUpdate|

<h2 id="tocS_v1.DeploymentStrategy">v1.DeploymentStrategy</h2>

<a id="schemav1.deploymentstrategy"></a>
<a id="schema_v1.DeploymentStrategy"></a>
<a id="tocSv1.deploymentstrategy"></a>
<a id="tocsv1.deploymentstrategy"></a>

```json
{
  "rollingUpdate": {
    "maxSurge": {
      "intVal": null,
      "strVal": null,
      "type": null
    },
    "maxUnavailable": {
      "intVal": null,
      "strVal": null,
      "type": null
    }
  },
  "type": "Recreate"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|rollingUpdate|[v1.RollingUpdateDeployment](#schemav1.rollingupdatedeployment)|false|none||Rolling update config params. Present only if DeploymentStrategyType =<br />RollingUpdate.<br />---<br />TODO: Update this to follow our convention for oneOf, whatever we decide it<br />to be.<br />+optional|
|type|[v1.DeploymentStrategyType](#schemav1.deploymentstrategytype)|false|none||Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.<br />+optional|

<h2 id="tocS_v1.DeploymentStatus">v1.DeploymentStatus</h2>

<a id="schemav1.deploymentstatus"></a>
<a id="schema_v1.DeploymentStatus"></a>
<a id="tocSv1.deploymentstatus"></a>
<a id="tocsv1.deploymentstatus"></a>

```json
{
  "availableReplicas": 0,
  "collisionCount": 0,
  "conditions": [
    {
      "lastTransitionTime": "string",
      "lastUpdateTime": "string",
      "message": "string",
      "reason": "string",
      "status": "True",
      "type": "Available"
    }
  ],
  "observedGeneration": 0,
  "readyReplicas": 0,
  "replicas": 0,
  "unavailableReplicas": 0,
  "updatedReplicas": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|availableReplicas|integer|false|none||Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.<br />+optional|
|collisionCount|integer|false|none||Count of hash collisions for the Deployment. The Deployment controller uses this<br />field as a collision avoidance mechanism when it needs to create the name for the<br />newest ReplicaSet.<br />+optional|
|conditions|[[v1.DeploymentCondition](#schemav1.deploymentcondition)]|false|none||Represents the latest available observations of a deployment's current state.<br />+patchMergeKey=type<br />+patchStrategy=merge|
|observedGeneration|integer|false|none||The generation observed by the deployment controller.<br />+optional|
|readyReplicas|integer|false|none||Total number of ready pods targeted by this deployment.<br />+optional|
|replicas|integer|false|none||Total number of non-terminated pods targeted by this deployment (their labels match the selector).<br />+optional|
|unavailableReplicas|integer|false|none||Total number of unavailable pods targeted by this deployment. This is the total number of<br />pods that are still required for the deployment to have 100% available capacity. They may<br />either be pods that are running but not yet available or pods that still have not been created.<br />+optional|
|updatedReplicas|integer|false|none||Total number of non-terminated pods targeted by this deployment that have the desired template spec.<br />+optional|

<h2 id="tocS_v1.DeploymentSpec">v1.DeploymentSpec</h2>

<a id="schemav1.deploymentspec"></a>
<a id="schema_v1.DeploymentSpec"></a>
<a id="tocSv1.deploymentspec"></a>
<a id="tocsv1.deploymentspec"></a>

```json
{
  "minReadySeconds": 0,
  "paused": true,
  "progressDeadlineSeconds": 0,
  "replicas": 0,
  "revisionHistoryLimit": 0,
  "selector": {
    "matchExpressions": [
      {
        "key": "string",
        "operator": null,
        "values": [
          null
        ]
      }
    ],
    "matchLabels": {
      "property1": "string",
      "property2": "string"
    }
  },
  "strategy": {
    "rollingUpdate": {
      "maxSurge": null,
      "maxUnavailable": null
    },
    "type": "Recreate"
  },
  "template": {
    "metadata": {
      "annotations": null,
      "clusterName": null,
      "creationTimestamp": null,
      "deletionGracePeriodSeconds": null,
      "deletionTimestamp": null,
      "finalizers": null,
      "generateName": null,
      "generation": null,
      "labels": null,
      "managedFields": null,
      "name": null,
      "namespace": null,
      "ownerReferences": null,
      "resourceVersion": null,
      "selfLink": null,
      "uid": null
    },
    "spec": {
      "activeDeadlineSeconds": null,
      "affinity": null,
      "automountServiceAccountToken": null,
      "containers": null,
      "dnsConfig": null,
      "dnsPolicy": null,
      "enableServiceLinks": null,
      "ephemeralContainers": null,
      "hostAliases": null,
      "hostIPC": null,
      "hostNetwork": null,
      "hostPID": null,
      "hostname": null,
      "imagePullSecrets": null,
      "initContainers": null,
      "nodeName": null,
      "nodeSelector": null,
      "overhead": null,
      "preemptionPolicy": null,
      "priority": null,
      "priorityClassName": null,
      "readinessGates": null,
      "restartPolicy": null,
      "runtimeClassName": null,
      "schedulerName": null,
      "securityContext": null,
      "serviceAccount": null,
      "serviceAccountName": null,
      "setHostnameAsFQDN": null,
      "shareProcessNamespace": null,
      "subdomain": null,
      "terminationGracePeriodSeconds": null,
      "tolerations": null,
      "topologySpreadConstraints": null,
      "volumes": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|minReadySeconds|integer|false|none||Minimum number of seconds for which a newly created pod should be ready<br />without any of its container crashing, for it to be considered available.<br />Defaults to 0 (pod will be considered available as soon as it is ready)<br />+optional|
|paused|boolean|false|none||Indicates that the deployment is paused.<br />+optional|
|progressDeadlineSeconds|integer|false|none||The maximum time in seconds for a deployment to make progress before it<br />is considered to be failed. The deployment controller will continue to<br />process failed deployments and a condition with a ProgressDeadlineExceeded<br />reason will be surfaced in the deployment status. Note that progress will<br />not be estimated during the time a deployment is paused. Defaults to 600s.|
|replicas|integer|false|none||Number of desired pods. This is a pointer to distinguish between explicit<br />zero and not specified. Defaults to 1.<br />+optional|
|revisionHistoryLimit|integer|false|none||The number of old ReplicaSets to retain to allow rollback.<br />This is a pointer to distinguish between explicit zero and not specified.<br />Defaults to 10.<br />+optional|
|selector|[v1.LabelSelector](#schemav1.labelselector)|false|none||Label selector for pods. Existing ReplicaSets whose pods are<br />selected by this will be the ones affected by this deployment.<br />It must match the pod template's labels.|
|strategy|[v1.DeploymentStrategy](#schemav1.deploymentstrategy)|false|none||The deployment strategy to use to replace existing pods with new ones.<br />+optional<br />+patchStrategy=retainKeys|
|template|[v1.PodTemplateSpec](#schemav1.podtemplatespec)|false|none||Template describes the pods that will be created.|

<h2 id="tocS_v1.DeploymentConditionType">v1.DeploymentConditionType</h2>

<a id="schemav1.deploymentconditiontype"></a>
<a id="schema_v1.DeploymentConditionType"></a>
<a id="tocSv1.deploymentconditiontype"></a>
<a id="tocsv1.deploymentconditiontype"></a>

```json
"Available"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Available|
|*anonymous*|Progressing|
|*anonymous*|ReplicaFailure|

<h2 id="tocS_v1.DeploymentCondition">v1.DeploymentCondition</h2>

<a id="schemav1.deploymentcondition"></a>
<a id="schema_v1.DeploymentCondition"></a>
<a id="tocSv1.deploymentcondition"></a>
<a id="tocsv1.deploymentcondition"></a>

```json
{
  "lastTransitionTime": "string",
  "lastUpdateTime": "string",
  "message": "string",
  "reason": "string",
  "status": "True",
  "type": "Available"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|lastTransitionTime|string|false|none||Last time the condition transitioned from one status to another.|
|lastUpdateTime|string|false|none||The last time this condition was updated.|
|message|string|false|none||A human readable message indicating details about the transition.|
|reason|string|false|none||The reason for the condition's last transition.|
|status|[k8s_io_api_core_v1.ConditionStatus](#schemak8s_io_api_core_v1.conditionstatus)|false|none||Status of the condition, one of True, False, Unknown.|
|type|[v1.DeploymentConditionType](#schemav1.deploymentconditiontype)|false|none||Type of deployment condition.|

<h2 id="tocS_v1.Deployment">v1.Deployment</h2>

<a id="schemav1.deployment"></a>
<a id="schema_v1.Deployment"></a>
<a id="tocSv1.deployment"></a>
<a id="tocsv1.deployment"></a>

```json
{
  "apiVersion": "string",
  "kind": "string",
  "metadata": {
    "annotations": {
      "property1": "string",
      "property2": "string"
    },
    "clusterName": "string",
    "creationTimestamp": "string",
    "deletionGracePeriodSeconds": 0,
    "deletionTimestamp": "string",
    "finalizers": [
      "string"
    ],
    "generateName": "string",
    "generation": 0,
    "labels": {
      "property1": "string",
      "property2": "string"
    },
    "managedFields": [
      {
        "apiVersion": "string",
        "fieldsType": "string",
        "fieldsV1": null,
        "manager": "string",
        "operation": null,
        "subresource": "string",
        "time": "string"
      }
    ],
    "name": "string",
    "namespace": "string",
    "ownerReferences": [
      {
        "apiVersion": "string",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "string",
        "name": "string",
        "uid": "string"
      }
    ],
    "resourceVersion": "string",
    "selfLink": "string",
    "uid": "string"
  },
  "spec": {
    "minReadySeconds": 0,
    "paused": true,
    "progressDeadlineSeconds": 0,
    "replicas": 0,
    "revisionHistoryLimit": 0,
    "selector": {
      "matchExpressions": null,
      "matchLabels": null
    },
    "strategy": {
      "rollingUpdate": null,
      "type": null
    },
    "template": {
      "metadata": null,
      "spec": null
    }
  },
  "status": {
    "availableReplicas": 0,
    "collisionCount": 0,
    "conditions": [
      {
        "lastTransitionTime": "string",
        "lastUpdateTime": "string",
        "message": "string",
        "reason": "string",
        "status": null,
        "type": null
      }
    ],
    "observedGeneration": 0,
    "readyReplicas": 0,
    "replicas": 0,
    "unavailableReplicas": 0,
    "updatedReplicas": 0
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|apiVersion|string|false|none||APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources<br />+optional|
|kind|string|false|none||Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br />+optional|
|metadata|[v1.ObjectMeta](#schemav1.objectmeta)|false|none||Standard object's metadata.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata<br />+optional|
|spec|[v1.DeploymentSpec](#schemav1.deploymentspec)|false|none||Specification of the desired behavior of the Deployment.<br />+optional|
|status|[v1.DeploymentStatus](#schemav1.deploymentstatus)|false|none||Most recently observed status of the Deployment.<br />+optional|

<h2 id="tocS_v1.DNSPolicy">v1.DNSPolicy</h2>

<a id="schemav1.dnspolicy"></a>
<a id="schema_v1.DNSPolicy"></a>
<a id="tocSv1.dnspolicy"></a>
<a id="tocsv1.dnspolicy"></a>

```json
"ClusterFirstWithHostNet"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|ClusterFirstWithHostNet|
|*anonymous*|ClusterFirst|
|*anonymous*|Default|
|*anonymous*|None|

<h2 id="tocS_v1.ContainerPort">v1.ContainerPort</h2>

<a id="schemav1.containerport"></a>
<a id="schema_v1.ContainerPort"></a>
<a id="tocSv1.containerport"></a>
<a id="tocsv1.containerport"></a>

```json
{
  "containerPort": 0,
  "hostIP": "string",
  "hostPort": 0,
  "name": "string",
  "protocol": "TCP"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|containerPort|integer|false|none||Number of port to expose on the pod's IP address.<br />This must be a valid port number, 0 < x < 65536.|
|hostIP|string|false|none||What host IP to bind the external port to.<br />+optional|
|hostPort|integer|false|none||Number of port to expose on the host.<br />If specified, this must be a valid port number, 0 < x < 65536.<br />If HostNetwork is specified, this must match ContainerPort.<br />Most containers do not need this.<br />+optional|
|name|string|false|none||If specified, this must be an IANA_SVC_NAME and unique within the pod. Each<br />named port in a pod must have a unique name. Name for the port that can be<br />referred to by services.<br />+optional|
|protocol|[v1.Protocol](#schemav1.protocol)|false|none||Protocol for port. Must be UDP, TCP, or SCTP.<br />Defaults to "TCP".<br />+optional<br />+default="TCP"|

<h2 id="tocS_v1.Container">v1.Container</h2>

<a id="schemav1.container"></a>
<a id="schema_v1.Container"></a>
<a id="tocSv1.container"></a>
<a id="tocsv1.container"></a>

```json
{
  "args": [
    "string"
  ],
  "command": [
    "string"
  ],
  "env": [
    {
      "name": "string",
      "value": "string",
      "valueFrom": {
        "configMapKeyRef": null,
        "fieldRef": null,
        "resourceFieldRef": null,
        "secretKeyRef": null
      }
    }
  ],
  "envFrom": [
    {
      "configMapRef": {
        "name": "string",
        "optional": true
      },
      "prefix": "string",
      "secretRef": {
        "name": "string",
        "optional": true
      }
    }
  ],
  "image": "string",
  "imagePullPolicy": "Always",
  "lifecycle": {
    "postStart": {
      "exec": null,
      "httpGet": null,
      "tcpSocket": null
    },
    "preStop": {
      "exec": null,
      "httpGet": null,
      "tcpSocket": null
    }
  },
  "livenessProbe": {
    "exec": {
      "command": null
    },
    "failureThreshold": 0,
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "initialDelaySeconds": 0,
    "periodSeconds": 0,
    "successThreshold": 0,
    "tcpSocket": {
      "host": null,
      "port": null
    },
    "terminationGracePeriodSeconds": 0,
    "timeoutSeconds": 0
  },
  "name": "string",
  "ports": [
    {
      "containerPort": 0,
      "hostIP": "string",
      "hostPort": 0,
      "name": "string",
      "protocol": "TCP"
    }
  ],
  "readinessProbe": {
    "exec": {
      "command": null
    },
    "failureThreshold": 0,
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "initialDelaySeconds": 0,
    "periodSeconds": 0,
    "successThreshold": 0,
    "tcpSocket": {
      "host": null,
      "port": null
    },
    "terminationGracePeriodSeconds": 0,
    "timeoutSeconds": 0
  },
  "resources": {
    "limits": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    },
    "requests": {
      "property1": {
        "Format": "DecimalExponent"
      },
      "property2": {
        "Format": "DecimalExponent"
      }
    }
  },
  "securityContext": {
    "allowPrivilegeEscalation": true,
    "capabilities": {
      "add": null,
      "drop": null
    },
    "privileged": true,
    "procMount": "Default",
    "readOnlyRootFilesystem": true,
    "runAsGroup": 0,
    "runAsNonRoot": true,
    "runAsUser": 0,
    "seLinuxOptions": {
      "level": null,
      "role": null,
      "type": null,
      "user": null
    },
    "seccompProfile": {
      "localhostProfile": null,
      "type": null
    },
    "windowsOptions": {
      "gmsaCredentialSpec": null,
      "gmsaCredentialSpecName": null,
      "hostProcess": null,
      "runAsUserName": null
    }
  },
  "startupProbe": {
    "exec": {
      "command": null
    },
    "failureThreshold": 0,
    "httpGet": {
      "host": null,
      "httpHeaders": null,
      "path": null,
      "port": null,
      "scheme": null
    },
    "initialDelaySeconds": 0,
    "periodSeconds": 0,
    "successThreshold": 0,
    "tcpSocket": {
      "host": null,
      "port": null
    },
    "terminationGracePeriodSeconds": 0,
    "timeoutSeconds": 0
  },
  "stdin": true,
  "stdinOnce": true,
  "terminationMessagePath": "string",
  "terminationMessagePolicy": "File",
  "tty": true,
  "volumeDevices": [
    {
      "devicePath": "string",
      "name": "string"
    }
  ],
  "volumeMounts": [
    {
      "mountPath": "string",
      "mountPropagation": "None",
      "name": "string",
      "readOnly": true,
      "subPath": "string",
      "subPathExpr": "string"
    }
  ],
  "workingDir": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|args|[string]|false|none||Arguments to the entrypoint.<br />The docker image's CMD is used if this is not provided.<br />Variable references $(VAR_NAME) are expanded using the container's environment. If a variable<br />cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced<br />to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will<br />produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless<br />of whether the variable exists or not. Cannot be updated.<br />More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell<br />+optional|
|command|[string]|false|none||Entrypoint array. Not executed within a shell.<br />The docker image's ENTRYPOINT is used if this is not provided.<br />Variable references $(VAR_NAME) are expanded using the container's environment. If a variable<br />cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced<br />to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will<br />produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless<br />of whether the variable exists or not. Cannot be updated.<br />More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell<br />+optional|
|env|[[v1.EnvVar](#schemav1.envvar)]|false|none||List of environment variables to set in the container.<br />Cannot be updated.<br />+optional<br />+patchMergeKey=name<br />+patchStrategy=merge|
|envFrom|[[v1.EnvFromSource](#schemav1.envfromsource)]|false|none||List of sources to populate environment variables in the container.<br />The keys defined within a source must be a C_IDENTIFIER. All invalid keys<br />will be reported as an event when the container is starting. When a key exists in multiple<br />sources, the value associated with the last source will take precedence.<br />Values defined by an Env with a duplicate key will take precedence.<br />Cannot be updated.<br />+optional|
|image|string|false|none||Docker image name.<br />More info: https://kubernetes.io/docs/concepts/containers/images<br />This field is optional to allow higher level config management to default or override<br />container images in workload controllers like Deployments and StatefulSets.<br />+optional|
|imagePullPolicy|[v1.PullPolicy](#schemav1.pullpolicy)|false|none||Image pull policy.<br />One of Always, Never, IfNotPresent.<br />Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.<br />Cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/containers/images#updating-images<br />+optional|
|lifecycle|[v1.Lifecycle](#schemav1.lifecycle)|false|none||Actions that the management system should take in response to container lifecycle events.<br />Cannot be updated.<br />+optional|
|livenessProbe|[v1.Probe](#schemav1.probe)|false|none||Periodic probe of container liveness.<br />Container will be restarted if the probe fails.<br />Cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes<br />+optional|
|name|string|false|none||Name of the container specified as a DNS_LABEL.<br />Each container in a pod must have a unique name (DNS_LABEL).<br />Cannot be updated.|
|ports|[[v1.ContainerPort](#schemav1.containerport)]|false|none||List of ports to expose from the container. Exposing a port here gives<br />the system additional information about the network connections a<br />container uses, but is primarily informational. Not specifying a port here<br />DOES NOT prevent that port from being exposed. Any port which is<br />listening on the default "0.0.0.0" address inside a container will be<br />accessible from the network.<br />Cannot be updated.<br />+optional<br />+patchMergeKey=containerPort<br />+patchStrategy=merge<br />+listType=map<br />+listMapKey=containerPort<br />+listMapKey=protocol|
|readinessProbe|[v1.Probe](#schemav1.probe)|false|none||Periodic probe of container service readiness.<br />Container will be removed from service endpoints if the probe fails.<br />Cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes<br />+optional|
|resources|[v1.ResourceRequirements](#schemav1.resourcerequirements)|false|none||Compute Resources required by this container.<br />Cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/<br />+optional|
|securityContext|[v1.SecurityContext](#schemav1.securitycontext)|false|none||SecurityContext defines the security options the container should be run with.<br />If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.<br />More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/<br />+optional|
|startupProbe|[v1.Probe](#schemav1.probe)|false|none||StartupProbe indicates that the Pod has successfully initialized.<br />If specified, no other probes are executed until this completes successfully.<br />If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.<br />This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,<br />when it might take a long time to load data or warm a cache, than during steady-state operation.<br />This cannot be updated.<br />More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes<br />+optional|
|stdin|boolean|false|none||Whether this container should allocate a buffer for stdin in the container runtime. If this<br />is not set, reads from stdin in the container will always result in EOF.<br />Default is false.<br />+optional|
|stdinOnce|boolean|false|none||Whether the container runtime should close the stdin channel after it has been opened by<br />a single attach. When stdin is true the stdin stream will remain open across multiple attach<br />sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the<br />first client attaches to stdin, and then remains open and accepts data until the client disconnects,<br />at which time stdin is closed and remains closed until the container is restarted. If this<br />flag is false, a container processes that reads from stdin will never receive an EOF.<br />Default is false<br />+optional|
|terminationMessagePath|string|false|none||Optional: Path at which the file to which the container's termination message<br />will be written is mounted into the container's filesystem.<br />Message written is intended to be brief final status, such as an assertion failure message.<br />Will be truncated by the node if greater than 4096 bytes. The total message length across<br />all containers will be limited to 12kb.<br />Defaults to /dev/termination-log.<br />Cannot be updated.<br />+optional|
|terminationMessagePolicy|[v1.TerminationMessagePolicy](#schemav1.terminationmessagepolicy)|false|none||Indicate how the termination message should be populated. File will use the contents of<br />terminationMessagePath to populate the container status message on both success and failure.<br />FallbackToLogsOnError will use the last chunk of container log output if the termination<br />message file is empty and the container exited with an error.<br />The log output is limited to 2048 bytes or 80 lines, whichever is smaller.<br />Defaults to File.<br />Cannot be updated.<br />+optional|
|tty|boolean|false|none||Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.<br />Default is false.<br />+optional|
|volumeDevices|[[v1.VolumeDevice](#schemav1.volumedevice)]|false|none||volumeDevices is the list of block devices to be used by the container.<br />+patchMergeKey=devicePath<br />+patchStrategy=merge<br />+optional|
|volumeMounts|[[v1.VolumeMount](#schemav1.volumemount)]|false|none||Pod volumes to mount into the container's filesystem.<br />Cannot be updated.<br />+optional<br />+patchMergeKey=mountPath<br />+patchStrategy=merge|
|workingDir|string|false|none||Container's working directory.<br />If not specified, the container runtime's default will be used, which<br />might be configured in the container image.<br />Cannot be updated.<br />+optional|

<h2 id="tocS_v1.ConfigMapVolumeSource">v1.ConfigMapVolumeSource</h2>

<a id="schemav1.configmapvolumesource"></a>
<a id="schema_v1.ConfigMapVolumeSource"></a>
<a id="tocSv1.configmapvolumesource"></a>
<a id="tocsv1.configmapvolumesource"></a>

```json
{
  "defaultMode": 0,
  "items": [
    {
      "key": "string",
      "mode": 0,
      "path": "string"
    }
  ],
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|defaultMode|integer|false|none||Optional: mode bits used to set permissions on created files by default.<br />Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br />YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br />Defaults to 0644.<br />Directories within the path are not affected by this setting.<br />This might be in conflict with other options that affect the file<br />mode, like fsGroup, and the result can be other mode bits set.<br />+optional|
|items|[[v1.KeyToPath](#schemav1.keytopath)]|false|none||If unspecified, each key-value pair in the Data field of the referenced<br />ConfigMap will be projected into the volume as a file whose name is the<br />key and content is the value. If specified, the listed keys will be<br />projected into the specified paths, and unlisted keys will not be<br />present. If a key is specified which is not present in the ConfigMap,<br />the volume setup will error unless it is marked optional. Paths must be<br />relative and may not contain the '..' path or start with '..'.<br />+optional|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the ConfigMap or its keys must be defined<br />+optional|

<h2 id="tocS_v1.ConfigMapProjection">v1.ConfigMapProjection</h2>

<a id="schemav1.configmapprojection"></a>
<a id="schema_v1.ConfigMapProjection"></a>
<a id="tocSv1.configmapprojection"></a>
<a id="tocsv1.configmapprojection"></a>

```json
{
  "items": [
    {
      "key": "string",
      "mode": 0,
      "path": "string"
    }
  ],
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|items|[[v1.KeyToPath](#schemav1.keytopath)]|false|none||If unspecified, each key-value pair in the Data field of the referenced<br />ConfigMap will be projected into the volume as a file whose name is the<br />key and content is the value. If specified, the listed keys will be<br />projected into the specified paths, and unlisted keys will not be<br />present. If a key is specified which is not present in the ConfigMap,<br />the volume setup will error unless it is marked optional. Paths must be<br />relative and may not contain the '..' path or start with '..'.<br />+optional|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the ConfigMap or its keys must be defined<br />+optional|

<h2 id="tocS_v1.ConfigMapKeySelector">v1.ConfigMapKeySelector</h2>

<a id="schemav1.configmapkeyselector"></a>
<a id="schema_v1.ConfigMapKeySelector"></a>
<a id="tocSv1.configmapkeyselector"></a>
<a id="tocsv1.configmapkeyselector"></a>

```json
{
  "key": "string",
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|key|string|false|none||The key to select.|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the ConfigMap or its key must be defined<br />+optional|

<h2 id="tocS_v1.ConfigMapEnvSource">v1.ConfigMapEnvSource</h2>

<a id="schemav1.configmapenvsource"></a>
<a id="schema_v1.ConfigMapEnvSource"></a>
<a id="tocSv1.configmapenvsource"></a>
<a id="tocsv1.configmapenvsource"></a>

```json
{
  "name": "string",
  "optional": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Name of the referent.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />TODO: Add other useful fields. apiVersion, kind, uid?<br />+optional|
|optional|boolean|false|none||Specify whether the ConfigMap must be defined<br />+optional|

<h2 id="tocS_v1.Condition">v1.Condition</h2>

<a id="schemav1.condition"></a>
<a id="schema_v1.Condition"></a>
<a id="tocSv1.condition"></a>
<a id="tocsv1.condition"></a>

```json
{
  "lastTransitionTime": "string",
  "message": "string",
  "observedGeneration": 0,
  "reason": "string",
  "status": "True",
  "type": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|lastTransitionTime|string|false|none||lastTransitionTime is the last time the condition transitioned from one status to another.<br />This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br />+required<br />+kubebuilder:validation:Required<br />+kubebuilder:validation:Type=string<br />+kubebuilder:validation:Format=date-time|
|message|string|false|none||message is a human readable message indicating details about the transition.<br />This may be an empty string.<br />+required<br />+kubebuilder:validation:Required<br />+kubebuilder:validation:MaxLength=32768|
|observedGeneration|integer|false|none||observedGeneration represents the .metadata.generation that the condition was set based upon.<br />For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date<br />with respect to the current state of the instance.<br />+optional<br />+kubebuilder:validation:Minimum=0|
|reason|string|false|none||reason contains a programmatic identifier indicating the reason for the condition's last transition.<br />Producers of specific condition types may define expected values and meanings for this field,<br />and whether the values are considered a guaranteed API.<br />The value should be a CamelCase string.<br />This field may not be empty.<br />+required<br />+kubebuilder:validation:Required<br />+kubebuilder:validation:MaxLength=1024<br />+kubebuilder:validation:MinLength=1<br />+kubebuilder:validation:Pattern=`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`|
|status|[k8s_io_apimachinery_pkg_apis_meta_v1.ConditionStatus](#schemak8s_io_apimachinery_pkg_apis_meta_v1.conditionstatus)|false|none||status of the condition, one of True, False, Unknown.<br />+required<br />+kubebuilder:validation:Required<br />+kubebuilder:validation:Enum=True;False;Unknown|
|type|string|false|none||type of condition in CamelCase or in foo.example.com/CamelCase.<br />---<br />Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be<br />useful (see .node.status.conditions), the ability to deconflict is important.<br />The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)<br />+required<br />+kubebuilder:validation:Required<br />+kubebuilder:validation:Pattern=`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`<br />+kubebuilder:validation:MaxLength=316|

<h2 id="tocS_v1.ClientIPConfig">v1.ClientIPConfig</h2>

<a id="schemav1.clientipconfig"></a>
<a id="schema_v1.ClientIPConfig"></a>
<a id="tocSv1.clientipconfig"></a>
<a id="tocsv1.clientipconfig"></a>

```json
{
  "timeoutSeconds": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|timeoutSeconds|integer|false|none||timeoutSeconds specifies the seconds of ClientIP type session sticky time.<br />The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP".<br />Default value is 10800(for 3 hours).<br />+optional|

<h2 id="tocS_v1.CinderVolumeSource">v1.CinderVolumeSource</h2>

<a id="schemav1.cindervolumesource"></a>
<a id="schema_v1.CinderVolumeSource"></a>
<a id="tocSv1.cindervolumesource"></a>
<a id="tocsv1.cindervolumesource"></a>

```json
{
  "fsType": "string",
  "readOnly": true,
  "secretRef": {
    "name": "string"
  },
  "volumeID": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />More info: https://examples.k8s.io/mysql-cinder-pd/README.md<br />+optional|
|readOnly|boolean|false|none||Optional: Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />More info: https://examples.k8s.io/mysql-cinder-pd/README.md<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||Optional: points to a secret object containing parameters used to connect<br />to OpenStack.<br />+optional|
|volumeID|string|false|none||volume id used to identify the volume in cinder.<br />More info: https://examples.k8s.io/mysql-cinder-pd/README.md|

<h2 id="tocS_v1.CephFSVolumeSource">v1.CephFSVolumeSource</h2>

<a id="schemav1.cephfsvolumesource"></a>
<a id="schema_v1.CephFSVolumeSource"></a>
<a id="tocSv1.cephfsvolumesource"></a>
<a id="tocsv1.cephfsvolumesource"></a>

```json
{
  "monitors": [
    "string"
  ],
  "path": "string",
  "readOnly": true,
  "secretFile": "string",
  "secretRef": {
    "name": "string"
  },
  "user": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|monitors|[string]|false|none||Required: Monitors is a collection of Ceph monitors<br />More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it|
|path|string|false|none||Optional: Used as the mounted root, rather than the full Ceph tree, default is /<br />+optional|
|readOnly|boolean|false|none||Optional: Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it<br />+optional|
|secretFile|string|false|none||Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret<br />More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it<br />+optional|
|secretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||Optional: SecretRef is reference to the authentication secret for User, default is empty.<br />More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it<br />+optional|
|user|string|false|none||Optional: User is the rados user name, default is admin<br />More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it<br />+optional|

<h2 id="tocS_v1.Capabilities">v1.Capabilities</h2>

<a id="schemav1.capabilities"></a>
<a id="schema_v1.Capabilities"></a>
<a id="tocSv1.capabilities"></a>
<a id="tocsv1.capabilities"></a>

```json
{
  "add": [
    "string"
  ],
  "drop": [
    "string"
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|add|[string]|false|none||Added capabilities<br />+optional|
|drop|[string]|false|none||Removed capabilities<br />+optional|

<h2 id="tocS_v1.CSIVolumeSource">v1.CSIVolumeSource</h2>

<a id="schemav1.csivolumesource"></a>
<a id="schema_v1.CSIVolumeSource"></a>
<a id="tocSv1.csivolumesource"></a>
<a id="tocsv1.csivolumesource"></a>

```json
{
  "driver": "string",
  "fsType": "string",
  "nodePublishSecretRef": {
    "name": "string"
  },
  "readOnly": true,
  "volumeAttributes": {
    "property1": "string",
    "property2": "string"
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|driver|string|false|none||Driver is the name of the CSI driver that handles this volume.<br />Consult with your admin for the correct name as registered in the cluster.|
|fsType|string|false|none||Filesystem type to mount. Ex. "ext4", "xfs", "ntfs".<br />If not provided, the empty value is passed to the associated CSI driver<br />which will determine the default filesystem to apply.<br />+optional|
|nodePublishSecretRef|[v1.LocalObjectReference](#schemav1.localobjectreference)|false|none||NodePublishSecretRef is a reference to the secret object containing<br />sensitive information to pass to the CSI driver to complete the CSI<br />NodePublishVolume and NodeUnpublishVolume calls.<br />This field is optional, and  may be empty if no secret is required. If the<br />secret object contains more than one secret, all secret references are passed.<br />+optional|
|readOnly|boolean|false|none||Specifies a read-only configuration for the volume.<br />Defaults to false (read/write).<br />+optional|
|volumeAttributes|object|false|none||VolumeAttributes stores driver-specific properties that are passed to the CSI<br />driver. Consult your driver's documentation for supported values.<br />+optional|
|» **additionalProperties**|string|false|none||none|

<h2 id="tocS_v1.AzureFileVolumeSource">v1.AzureFileVolumeSource</h2>

<a id="schemav1.azurefilevolumesource"></a>
<a id="schema_v1.AzureFileVolumeSource"></a>
<a id="tocSv1.azurefilevolumesource"></a>
<a id="tocsv1.azurefilevolumesource"></a>

```json
{
  "readOnly": true,
  "secretName": "string",
  "shareName": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|readOnly|boolean|false|none||Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|
|secretName|string|false|none||the name of secret that contains Azure Storage Account Name and Key|
|shareName|string|false|none||Share Name|

<h2 id="tocS_v1.AzureDiskVolumeSource">v1.AzureDiskVolumeSource</h2>

<a id="schemav1.azurediskvolumesource"></a>
<a id="schema_v1.AzureDiskVolumeSource"></a>
<a id="tocSv1.azurediskvolumesource"></a>
<a id="tocsv1.azurediskvolumesource"></a>

```json
{
  "cachingMode": "None",
  "diskName": "string",
  "diskURI": "string",
  "fsType": "string",
  "kind": "Shared",
  "readOnly": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|cachingMode|[v1.AzureDataDiskCachingMode](#schemav1.azuredatadiskcachingmode)|false|none||Host Caching mode: None, Read Only, Read Write.<br />+optional|
|diskName|string|false|none||The Name of the data disk in the blob storage|
|diskURI|string|false|none||The URI the data disk in the blob storage|
|fsType|string|false|none||Filesystem type to mount.<br />Must be a filesystem type supported by the host operating system.<br />Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />+optional|
|kind|[v1.AzureDataDiskKind](#schemav1.azuredatadiskkind)|false|none||Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared|
|readOnly|boolean|false|none||Defaults to false (read/write). ReadOnly here will force<br />the ReadOnly setting in VolumeMounts.<br />+optional|

<h2 id="tocS_v1.AzureDataDiskKind">v1.AzureDataDiskKind</h2>

<a id="schemav1.azuredatadiskkind"></a>
<a id="schema_v1.AzureDataDiskKind"></a>
<a id="tocSv1.azuredatadiskkind"></a>
<a id="tocsv1.azuredatadiskkind"></a>

```json
"Shared"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|Shared|
|*anonymous*|Dedicated|
|*anonymous*|Managed|

<h2 id="tocS_v1.AzureDataDiskCachingMode">v1.AzureDataDiskCachingMode</h2>

<a id="schemav1.azuredatadiskcachingmode"></a>
<a id="schema_v1.AzureDataDiskCachingMode"></a>
<a id="tocSv1.azuredatadiskcachingmode"></a>
<a id="tocsv1.azuredatadiskcachingmode"></a>

```json
"None"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|None|
|*anonymous*|ReadOnly|
|*anonymous*|ReadWrite|

<h2 id="tocS_v1.Affinity">v1.Affinity</h2>

<a id="schemav1.affinity"></a>
<a id="schema_v1.Affinity"></a>
<a id="tocSv1.affinity"></a>
<a id="tocsv1.affinity"></a>

```json
{
  "nodeAffinity": {
    "preferredDuringSchedulingIgnoredDuringExecution": [
      {
        "preference": null,
        "weight": 0
      }
    ],
    "requiredDuringSchedulingIgnoredDuringExecution": {
      "nodeSelectorTerms": null
    }
  },
  "podAffinity": {
    "preferredDuringSchedulingIgnoredDuringExecution": [
      {
        "podAffinityTerm": null,
        "weight": 0
      }
    ],
    "requiredDuringSchedulingIgnoredDuringExecution": [
      {
        "labelSelector": null,
        "namespaceSelector": null,
        "namespaces": [
          null
        ],
        "topologyKey": "string"
      }
    ]
  },
  "podAntiAffinity": {
    "preferredDuringSchedulingIgnoredDuringExecution": [
      {
        "podAffinityTerm": null,
        "weight": 0
      }
    ],
    "requiredDuringSchedulingIgnoredDuringExecution": [
      {
        "labelSelector": null,
        "namespaceSelector": null,
        "namespaces": [
          null
        ],
        "topologyKey": "string"
      }
    ]
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|nodeAffinity|[v1.NodeAffinity](#schemav1.nodeaffinity)|false|none||Describes node affinity scheduling rules for the pod.<br />+optional|
|podAffinity|[v1.PodAffinity](#schemav1.podaffinity)|false|none||Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).<br />+optional|
|podAntiAffinity|[v1.PodAntiAffinity](#schemav1.podantiaffinity)|false|none||Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).<br />+optional|

<h2 id="tocS_v1.AWSElasticBlockStoreVolumeSource">v1.AWSElasticBlockStoreVolumeSource</h2>

<a id="schemav1.awselasticblockstorevolumesource"></a>
<a id="schema_v1.AWSElasticBlockStoreVolumeSource"></a>
<a id="tocSv1.awselasticblockstorevolumesource"></a>
<a id="tocsv1.awselasticblockstorevolumesource"></a>

```json
{
  "fsType": "string",
  "partition": 0,
  "readOnly": true,
  "volumeID": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|fsType|string|false|none||Filesystem type of the volume that you want to mount.<br />Tip: Ensure that the filesystem type is supported by the host operating system.<br />Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore<br />TODO: how do we prevent errors in the filesystem from compromising the machine<br />+optional|
|partition|integer|false|none||The partition in the volume that you want to mount.<br />If omitted, the default is to mount by volume name.<br />Examples: For volume /dev/sda1, you specify the partition as "1".<br />Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br />+optional|
|readOnly|boolean|false|none||Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".<br />If omitted, the default is "false".<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore<br />+optional|
|volumeID|string|false|none||Unique ID of the persistent disk resource in AWS (Amazon EBS volume).<br />More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore|

<h2 id="tocS_system.SysBaseFiles">system.SysBaseFiles</h2>

<a id="schemasystem.sysbasefiles"></a>
<a id="schema_system.SysBaseFiles"></a>
<a id="tocSsystem.sysbasefiles"></a>
<a id="tocssystem.sysbasefiles"></a>

```json
{
  "createAt": "string",
  "id": "string",
  "name": "string",
  "originalName": "string",
  "size": 0,
  "suffix": "string",
  "updateAt": "string",
  "url": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|createAt|string|false|none||none|
|id|string|false|none||Default parameters _id Default parameters|
|name|string|false|none||file name|
|originalName|string|false|none||original file name|
|size|integer|false|none||file size|
|suffix|string|false|none||file extension|
|updateAt|string|false|none||none|
|url|string|false|none||url address|

<h2 id="tocS_response.Response">response.Response</h2>

<a id="schemaresponse.response"></a>
<a id="schema_response.Response"></a>
<a id="tocSresponse.response"></a>
<a id="tocsresponse.response"></a>

```json
{
  "code": 0,
  "data": "string",
  "message": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|code|integer|false|none||none|
|data|string|false|none||none|
|message|string|false|none||none|

<h2 id="tocS_resource.Quantity">resource.Quantity</h2>

<a id="schemaresource.quantity"></a>
<a id="schema_resource.Quantity"></a>
<a id="tocSresource.quantity"></a>
<a id="tocsresource.quantity"></a>

```json
{
  "Format": "DecimalExponent"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|Format|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|Format|DecimalExponent|
|Format|BinarySI|
|Format|DecimalSI|

<h2 id="tocS_request.UpdateSoftwareRequest">request.UpdateSoftwareRequest</h2>

<a id="schemarequest.updatesoftwarerequest"></a>
<a id="schema_request.UpdateSoftwareRequest"></a>
<a id="tocSrequest.updatesoftwarerequest"></a>
<a id="tocsrequest.updatesoftwarerequest"></a>

```json
{
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": "string",
      "TLSCertPath": "string",
      "TLSKeyPath": "string",
      "TLSType": 1
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": false,
      "allowContainerCapabilitiesForRegularUsers": true,
      "allowDeviceMappingForRegularUsers": true,
      "allowHostNamespaceForRegularUsers": true,
      "allowPrivilegedModeForRegularUsers": false,
      "allowStackManagementForRegularUsers": true,
      "allowSysctlSettingForRegularUsers": true,
      "allowVolumeBrowserForRegularUsers": true,
      "enableHostManagementFeatures": true
    },
    "status": 1,
    "templates": [
      {
        "category": [
          "string"
        ],
        "config": {},
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          "string"
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": {},
        "poster": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "preview": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "relationSoft": [
          {
            "category": null,
            "config": null,
            "copyright": null,
            "createAt": null,
            "devLanguage": null,
            "downloads": null,
            "features": null,
            "id": null,
            "isDel": null,
            "logo": null,
            "name": null,
            "nameZh": null,
            "notice": null,
            "poster": null,
            "preview": null,
            "relationSoft": null,
            "relationSoftIds": null,
            "score": null,
            "state": null,
            "supportMode": null,
            "sysLanguage": null,
            "team": null,
            "teamIcon": null,
            "updateAt": null,
            "userId": null
          }
        ],
        "relationSoftIds": [
          "string"
        ],
        "score": 0,
        "state": 0,
        "supportMode": {},
        "sysLanguage": [
          "string"
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "endpointId": "string",
  "kubernetesConfig": {
    "deployment": {
      "apiVersion": "string",
      "kind": "string",
      "metadata": {
        "annotations": null,
        "clusterName": null,
        "creationTimestamp": null,
        "deletionGracePeriodSeconds": null,
        "deletionTimestamp": null,
        "finalizers": null,
        "generateName": null,
        "generation": null,
        "labels": null,
        "managedFields": null,
        "name": null,
        "namespace": null,
        "ownerReferences": null,
        "resourceVersion": null,
        "selfLink": null,
        "uid": null
      },
      "spec": {
        "minReadySeconds": null,
        "paused": null,
        "progressDeadlineSeconds": null,
        "replicas": null,
        "revisionHistoryLimit": null,
        "selector": null,
        "strategy": null,
        "template": null
      },
      "status": {
        "availableReplicas": null,
        "collisionCount": null,
        "conditions": null,
        "observedGeneration": null,
        "readyReplicas": null,
        "replicas": null,
        "unavailableReplicas": null,
        "updatedReplicas": null
      }
    },
    "persistentVolumeClaim": {
      "apiVersion": "string",
      "kind": "string",
      "metadata": {
        "annotations": null,
        "clusterName": null,
        "creationTimestamp": null,
        "deletionGracePeriodSeconds": null,
        "deletionTimestamp": null,
        "finalizers": null,
        "generateName": null,
        "generation": null,
        "labels": null,
        "managedFields": null,
        "name": null,
        "namespace": null,
        "ownerReferences": null,
        "resourceVersion": null,
        "selfLink": null,
        "uid": null
      },
      "spec": {
        "accessModes": null,
        "dataSource": null,
        "dataSourceRef": null,
        "resources": null,
        "selector": null,
        "storageClassName": null,
        "volumeMode": null,
        "volumeName": null
      },
      "status": {
        "accessModes": null,
        "capacity": null,
        "conditions": null,
        "phase": null
      }
    },
    "service": {
      "apiVersion": "string",
      "kind": "string",
      "metadata": {
        "annotations": null,
        "clusterName": null,
        "creationTimestamp": null,
        "deletionGracePeriodSeconds": null,
        "deletionTimestamp": null,
        "finalizers": null,
        "generateName": null,
        "generation": null,
        "labels": null,
        "managedFields": null,
        "name": null,
        "namespace": null,
        "ownerReferences": null,
        "resourceVersion": null,
        "selfLink": null,
        "uid": null
      },
      "spec": {
        "allocateLoadBalancerNodePorts": null,
        "clusterIP": null,
        "clusterIPs": null,
        "externalIPs": null,
        "externalName": null,
        "externalTrafficPolicy": null,
        "healthCheckNodePort": null,
        "internalTrafficPolicy": null,
        "ipFamilies": null,
        "ipFamilyPolicy": null,
        "loadBalancerClass": null,
        "loadBalancerIP": null,
        "loadBalancerSourceRanges": null,
        "ports": null,
        "publishNotReadyAddresses": null,
        "selector": null,
        "sessionAffinity": null,
        "sessionAffinityConfig": null,
        "type": null
      },
      "status": {
        "conditions": null,
        "loadBalancer": null
      }
    }
  },
  "name": "string",
  "newConfig": {
    "property1": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    },
    "property2": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    }
  },
  "newContent": "string",
  "oldConfig": {
    "property1": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    },
    "property2": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    }
  },
  "oldContent": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|endpoint|[docker.DocEndpoint](#schemadocker.docendpoint)|false|none||none|
|endpointId|string|false|none||none|
|kubernetesConfig|[docker.KubernetesConfig](#schemadocker.kubernetesconfig)|false|none||none|
|name|string|false|none||none|
|newConfig|object|false|none||none|
|» **additionalProperties**|[docker.TemplateConfig](#schemadocker.templateconfig)|false|none||none|
|newContent|string|false|none||none|
|oldConfig|object|false|none||none|
|» **additionalProperties**|[docker.TemplateConfig](#schemadocker.templateconfig)|false|none||none|
|oldContent|string|false|none||none|

<h2 id="tocS_request.StopSoftwareRequest">request.StopSoftwareRequest</h2>

<a id="schemarequest.stopsoftwarerequest"></a>
<a id="schema_request.StopSoftwareRequest"></a>
<a id="tocSrequest.stopsoftwarerequest"></a>
<a id="tocsrequest.stopsoftwarerequest"></a>

```json
{
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": "string",
      "TLSCertPath": "string",
      "TLSKeyPath": "string",
      "TLSType": 1
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": false,
      "allowContainerCapabilitiesForRegularUsers": true,
      "allowDeviceMappingForRegularUsers": true,
      "allowHostNamespaceForRegularUsers": true,
      "allowPrivilegedModeForRegularUsers": false,
      "allowStackManagementForRegularUsers": true,
      "allowSysctlSettingForRegularUsers": true,
      "allowVolumeBrowserForRegularUsers": true,
      "enableHostManagementFeatures": true
    },
    "status": 1,
    "templates": [
      {
        "category": [
          "string"
        ],
        "config": {},
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          "string"
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": {},
        "poster": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "preview": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "relationSoft": [
          {
            "category": null,
            "config": null,
            "copyright": null,
            "createAt": null,
            "devLanguage": null,
            "downloads": null,
            "features": null,
            "id": null,
            "isDel": null,
            "logo": null,
            "name": null,
            "nameZh": null,
            "notice": null,
            "poster": null,
            "preview": null,
            "relationSoft": null,
            "relationSoftIds": null,
            "score": null,
            "state": null,
            "supportMode": null,
            "sysLanguage": null,
            "team": null,
            "teamIcon": null,
            "updateAt": null,
            "userId": null
          }
        ],
        "relationSoftIds": [
          "string"
        ],
        "score": 0,
        "state": 0,
        "supportMode": {},
        "sysLanguage": [
          "string"
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "endpointId": "string",
  "name": "string",
  "projectPath": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|endpoint|[docker.DocEndpoint](#schemadocker.docendpoint)|false|none||none|
|endpointId|string|false|none||none|
|name|string|false|none||none|
|projectPath|string|false|none||none|

<h2 id="tocS_request.StartSoftwareRequest">request.StartSoftwareRequest</h2>

<a id="schemarequest.startsoftwarerequest"></a>
<a id="schema_request.StartSoftwareRequest"></a>
<a id="tocSrequest.startsoftwarerequest"></a>
<a id="tocsrequest.startsoftwarerequest"></a>

```json
{
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": "string",
      "TLSCertPath": "string",
      "TLSKeyPath": "string",
      "TLSType": 1
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": false,
      "allowContainerCapabilitiesForRegularUsers": true,
      "allowDeviceMappingForRegularUsers": true,
      "allowHostNamespaceForRegularUsers": true,
      "allowPrivilegedModeForRegularUsers": false,
      "allowStackManagementForRegularUsers": true,
      "allowSysctlSettingForRegularUsers": true,
      "allowVolumeBrowserForRegularUsers": true,
      "enableHostManagementFeatures": true
    },
    "status": 1,
    "templates": [
      {
        "category": [
          "string"
        ],
        "config": {},
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          "string"
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": {},
        "poster": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "preview": [
          {
            "createAt": null,
            "id": null,
            "name": null,
            "originalName": null,
            "size": null,
            "suffix": null,
            "updateAt": null,
            "url": null
          }
        ],
        "relationSoft": [
          {
            "category": null,
            "config": null,
            "copyright": null,
            "createAt": null,
            "devLanguage": null,
            "downloads": null,
            "features": null,
            "id": null,
            "isDel": null,
            "logo": null,
            "name": null,
            "nameZh": null,
            "notice": null,
            "poster": null,
            "preview": null,
            "relationSoft": null,
            "relationSoftIds": null,
            "score": null,
            "state": null,
            "supportMode": null,
            "sysLanguage": null,
            "team": null,
            "teamIcon": null,
            "updateAt": null,
            "userId": null
          }
        ],
        "relationSoftIds": [
          "string"
        ],
        "score": 0,
        "state": 0,
        "supportMode": {},
        "sysLanguage": [
          "string"
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "endpointId": "string",
  "name": "string",
  "projectPath": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|endpoint|[docker.DocEndpoint](#schemadocker.docendpoint)|false|none||none|
|endpointId|string|false|none||none|
|name|string|false|none||none|
|projectPath|string|false|none||none|

<h2 id="tocS_request.DelSoftwareReq">request.DelSoftwareReq</h2>

<a id="schemarequest.delsoftwarereq"></a>
<a id="schema_request.DelSoftwareReq"></a>
<a id="tocSrequest.delsoftwarereq"></a>
<a id="tocsrequest.delsoftwarereq"></a>

```json
{
  "id": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|id|string|false|none||none|

<h2 id="tocS_k8s_io_apimachinery_pkg_apis_meta_v1.ConditionStatus">k8s_io_apimachinery_pkg_apis_meta_v1.ConditionStatus</h2>

<a id="schemak8s_io_apimachinery_pkg_apis_meta_v1.conditionstatus"></a>
<a id="schema_k8s_io_apimachinery_pkg_apis_meta_v1.ConditionStatus"></a>
<a id="tocSk8s_io_apimachinery_pkg_apis_meta_v1.conditionstatus"></a>
<a id="tocsk8s_io_apimachinery_pkg_apis_meta_v1.conditionstatus"></a>

```json
"True"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|True|
|*anonymous*|False|
|*anonymous*|Unknown|

<h2 id="tocS_k8s_io_api_core_v1.ConditionStatus">k8s_io_api_core_v1.ConditionStatus</h2>

<a id="schemak8s_io_api_core_v1.conditionstatus"></a>
<a id="schema_k8s_io_api_core_v1.ConditionStatus"></a>
<a id="tocSk8s_io_api_core_v1.conditionstatus"></a>
<a id="tocsk8s_io_api_core_v1.conditionstatus"></a>

```json
"True"

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|string|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|True|
|*anonymous*|False|
|*anonymous*|Unknown|

<h2 id="tocS_intstr.Type">intstr.Type</h2>

<a id="schemaintstr.type"></a>
<a id="schema_intstr.Type"></a>
<a id="tocSintstr.type"></a>
<a id="tocsintstr.type"></a>

```json
0

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|*anonymous*|integer|false|none||none|

#### enum 

|attribute|attribute|
|---|---|
|*anonymous*|0|
|*anonymous*|1|

<h2 id="tocS_intstr.IntOrString">intstr.IntOrString</h2>

<a id="schemaintstr.intorstring"></a>
<a id="schema_intstr.IntOrString"></a>
<a id="tocSintstr.intorstring"></a>
<a id="tocsintstr.intorstring"></a>

```json
{
  "intVal": 0,
  "strVal": "string",
  "type": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|intVal|integer|false|none||none|
|strVal|string|false|none||none|
|type|[intstr.Type](#schemaintstr.type)|false|none||none|

<h2 id="tocS_docker.YmlSecret">docker.YmlSecret</h2>

<a id="schemadocker.ymlsecret"></a>
<a id="schema_docker.YmlSecret"></a>
<a id="tocSdocker.ymlsecret"></a>
<a id="tocsdocker.ymlsecret"></a>

```json
{
  "external": "string",
  "file": "string",
  "name": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|external|string|false|none||Does it already exist，Does it already exist。|
|file|string|false|none||File Path|
|name|string|false|none||(v3.5+) name|

<h2 id="tocS_docker.Yml">docker.Yml</h2>

<a id="schemadocker.yml"></a>
<a id="schema_docker.Yml"></a>
<a id="tocSdocker.yml"></a>
<a id="tocsdocker.yml"></a>

```json
{
  "networks": {
    "property1": {
      "external": true
    },
    "property2": {
      "external": true
    }
  },
  "secrets": {
    "property1": {
      "external": "string",
      "file": "string",
      "name": "string"
    },
    "property2": {
      "external": "string",
      "file": "string",
      "name": "string"
    }
  },
  "services": {
    "property1": {
      "build": {
        "args": {},
        "cacheFrom": [
          null
        ],
        "context": "string",
        "dockerfile": "string",
        "labels": {},
        "shmSize": "string",
        "target": "string"
      },
      "capAdd": [
        "string"
      ],
      "capDrop": [
        "string"
      ],
      "cgroupParent": "string",
      "command": [
        "string"
      ],
      "containerName": "string",
      "dependsOn": [
        "string"
      ],
      "deploy": {
        "endpointMode": "string",
        "labels": [
          null
        ],
        "maxReplicasPerNode": 0,
        "mode": "string",
        "placement": [
          null
        ],
        "replicas": 0,
        "resources": null,
        "restartPolicy": null,
        "rollbackConfig": null,
        "updateConfig": null
      },
      "domainName": "string",
      "envFile": [
        "string"
      ],
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostName": "string",
      "image": {
        "name": "string",
        "tag": "string"
      },
      "ipc": "string",
      "macAddress": "string",
      "networkMode": "string",
      "networks": {
        "property1": {
          "aliases": [
            "string"
          ],
          "ipv4Address": "string",
          "ipv6Address": "string"
        },
        "property2": {
          "aliases": [
            "string"
          ],
          "ipv4Address": "string",
          "ipv6Address": "string"
        }
      },
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "privileged": true,
      "proflies": [
        "string"
      ],
      "readOnly": true,
      "restart": "string",
      "secrets": [
        {
          "gid": "string",
          "mode": "string",
          "source": "string",
          "target": "string",
          "uid": "string"
        }
      ],
      "shmSize": "string",
      "stdinOpen": true,
      "tty": true,
      "user": "string",
      "volumes": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ],
      "workingDir": "string"
    },
    "property2": {
      "build": {
        "args": {},
        "cacheFrom": [
          null
        ],
        "context": "string",
        "dockerfile": "string",
        "labels": {},
        "shmSize": "string",
        "target": "string"
      },
      "capAdd": [
        "string"
      ],
      "capDrop": [
        "string"
      ],
      "cgroupParent": "string",
      "command": [
        "string"
      ],
      "containerName": "string",
      "dependsOn": [
        "string"
      ],
      "deploy": {
        "endpointMode": "string",
        "labels": [
          null
        ],
        "maxReplicasPerNode": 0,
        "mode": "string",
        "placement": [
          null
        ],
        "replicas": 0,
        "resources": null,
        "restartPolicy": null,
        "rollbackConfig": null,
        "updateConfig": null
      },
      "domainName": "string",
      "envFile": [
        "string"
      ],
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostName": "string",
      "image": {
        "name": "string",
        "tag": "string"
      },
      "ipc": "string",
      "macAddress": "string",
      "networkMode": "string",
      "networks": {
        "property1": {
          "aliases": [
            "string"
          ],
          "ipv4Address": "string",
          "ipv6Address": "string"
        },
        "property2": {
          "aliases": [
            "string"
          ],
          "ipv4Address": "string",
          "ipv6Address": "string"
        }
      },
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "privileged": true,
      "proflies": [
        "string"
      ],
      "readOnly": true,
      "restart": "string",
      "secrets": [
        {
          "gid": "string",
          "mode": "string",
          "source": "string",
          "target": "string",
          "uid": "string"
        }
      ],
      "shmSize": "string",
      "stdinOpen": true,
      "tty": true,
      "user": "string",
      "volumes": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ],
      "workingDir": "string"
    }
  },
  "version": "string",
  "volumes": {
    "property1": {
      "external": true
    },
    "property2": {
      "external": true
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|networks|object|false|none||network configuration|
|» **additionalProperties**|[docker.Network](#schemadocker.network)|false|none||none|
|secrets|object|false|none||secret key|
|» **additionalProperties**|[docker.YmlSecret](#schemadocker.ymlsecret)|false|none||none|
|services|object|false|none||Service configuration|
|» **additionalProperties**|[docker.Service](#schemadocker.service)|false|none||none|
|version|string|false|none||Version number|
|volumes|object|false|none||Mount Volume Configuration|
|» **additionalProperties**|[docker.Volume](#schemadocker.volume)|false|none||none|

<h2 id="tocS_docker.VolumeMap">docker.VolumeMap</h2>

<a id="schemadocker.volumemap"></a>
<a id="schema_docker.VolumeMap"></a>
<a id="tocSdocker.volumemap"></a>
<a id="tocsdocker.volumemap"></a>

```json
{
  "desc": "string",
  "readOnly": true,
  "source": "string",
  "target": "string",
  "type": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|desc|string|false|none||describe|
|readOnly|boolean|false|none||Read Only|
|source|string|false|none||Path to external host|
|target|string|false|none||Path to internal container|
|type|string|false|none||type bind, volume, tmpfs|

<h2 id="tocS_docker.Volume">docker.Volume</h2>

<a id="schemadocker.volume"></a>
<a id="schema_docker.Volume"></a>
<a id="tocSdocker.volume"></a>
<a id="tocsdocker.volume"></a>

```json
{
  "external": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|external|boolean|false|none||TODO driver<br />TODO driver_opts|

<h2 id="tocS_docker.UpdateConfig">docker.UpdateConfig</h2>

<a id="schemadocker.updateconfig"></a>
<a id="schema_docker.UpdateConfig"></a>
<a id="tocSdocker.updateconfig"></a>
<a id="tocsdocker.updateconfig"></a>

```json
{
  "delay": 0,
  "failureAction": "string",
  "maxFailureRatio": 0,
  "monitor": 0,
  "order": "string",
  "parallelism": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|delay|integer|false|none||delay Service update delay Service update delay，Service update delay 0 Service update delay 1 000 000 000|
|failureAction|string|false|none||failure_action Service update failure behavior Service update failure behavior，Service update failure behavior continue、pause、rollback|
|maxFailureRatio|number|false|none||max_failure_ratio Service update failure rate Service update failure rate，Service update failure rate 0.000 Service update failure rate 1.000|
|monitor|integer|false|none||monitor Service update monitoring Service update monitoring，Service update monitoring 0 Service update monitoring 1 000 000 000|
|order|string|false|none||order Service update order Service update order，Service update order start-first、stop-first|
|parallelism|integer|false|none||parallelism Service update parallelism Service update parallelism，Service update parallelism 0 Service update parallelism 1 000|

<h2 id="tocS_docker.TemplateConfig">docker.TemplateConfig</h2>

<a id="schemadocker.templateconfig"></a>
<a id="schema_docker.TemplateConfig"></a>
<a id="tocSdocker.templateconfig"></a>
<a id="tocsdocker.templateconfig"></a>

```json
{
  "environment": [
    {
      "desc": "string",
      "key": "string",
      "val": "string"
    }
  ],
  "hostname": "string",
  "ports": [
    {
      "desc": "string",
      "mode": "string",
      "protocol": "string",
      "published": "string",
      "target": "string"
    }
  ],
  "volume": [
    {
      "desc": "string",
      "readOnly": true,
      "source": "string",
      "target": "string",
      "type": "string"
    }
  ]
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|environment|[[docker.Environment](#schemadocker.environment)]|false|none||none|
|hostname|string|false|none||none|
|ports|[[docker.Port](#schemadocker.port)]|false|none||none|
|volume|[[docker.VolumeMap](#schemadocker.volumemap)]|false|none||none|

<h2 id="tocS_docker.TLSConfiguration">docker.TLSConfiguration</h2>

<a id="schemadocker.tlsconfiguration"></a>
<a id="schema_docker.TLSConfiguration"></a>
<a id="tocSdocker.tlsconfiguration"></a>
<a id="tocsdocker.tlsconfiguration"></a>

```json
{
  "TLSCACertPath": "string",
  "TLSCertPath": "string",
  "TLSKeyPath": "string",
  "TLSType": 1
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|TLSCACertPath|string|false|none||TLS CAThe path to the certificate file // ca.pem|
|TLSCertPath|string|false|none||TLSThe path to the client certificate file // cert.pem|
|TLSKeyPath|string|false|none||TLSThe path to the client key file // key.pem|
|TLSType|integer|false|none||TLSType TLStype|

<h2 id="tocS_docker.Service">docker.Service</h2>

<a id="schemadocker.service"></a>
<a id="schema_docker.Service"></a>
<a id="tocSdocker.service"></a>
<a id="tocsdocker.service"></a>

```json
{
  "build": {
    "args": {},
    "cacheFrom": [
      {
        "name": "string",
        "tag": "string"
      }
    ],
    "context": "string",
    "dockerfile": "string",
    "labels": {},
    "shmSize": "string",
    "target": "string"
  },
  "capAdd": [
    "string"
  ],
  "capDrop": [
    "string"
  ],
  "cgroupParent": "string",
  "command": [
    "string"
  ],
  "containerName": "string",
  "dependsOn": [
    "string"
  ],
  "deploy": {
    "endpointMode": "string",
    "labels": [
      "string"
    ],
    "maxReplicasPerNode": 0,
    "mode": "string",
    "placement": [
      {
        "property1": "string",
        "property2": "string"
      }
    ],
    "replicas": 0,
    "resources": {
      "cpus": null,
      "memory": null
    },
    "restartPolicy": {
      "condition": null
    },
    "rollbackConfig": {
      "delay": null,
      "failureAction": null,
      "maxFailureRatio": null,
      "monitor": null,
      "order": null,
      "parallelism": null
    },
    "updateConfig": {
      "delay": null,
      "failureAction": null,
      "maxFailureRatio": null,
      "monitor": null,
      "order": null,
      "parallelism": null
    }
  },
  "domainName": "string",
  "envFile": [
    "string"
  ],
  "environment": [
    {
      "desc": "string",
      "key": "string",
      "val": "string"
    }
  ],
  "hostName": "string",
  "image": {
    "name": "string",
    "tag": "string"
  },
  "ipc": "string",
  "macAddress": "string",
  "networkMode": "string",
  "networks": {
    "property1": {
      "aliases": [
        "string"
      ],
      "ipv4Address": "string",
      "ipv6Address": "string"
    },
    "property2": {
      "aliases": [
        "string"
      ],
      "ipv4Address": "string",
      "ipv6Address": "string"
    }
  },
  "ports": [
    {
      "desc": "string",
      "mode": "string",
      "protocol": "string",
      "published": "string",
      "target": "string"
    }
  ],
  "privileged": true,
  "proflies": [
    "string"
  ],
  "readOnly": true,
  "restart": "string",
  "secrets": [
    {
      "gid": "string",
      "mode": "string",
      "source": "string",
      "target": "string",
      "uid": "string"
    }
  ],
  "shmSize": "string",
  "stdinOpen": true,
  "tty": true,
  "user": "string",
  "volumes": [
    {
      "desc": "string",
      "readOnly": true,
      "source": "string",
      "target": "string",
      "type": "string"
    }
  ],
  "workingDir": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|build|[docker.Build](#schemadocker.build)|false|none||Configuration items during construction|
|capAdd|[string]|false|none||Added container functionality|
|capDrop|[string]|false|none||Added container functionality|
|cgroupParent|string|false|none||Optional Parent Control Group|
|command|[string]|false|none||Default Command|
|containerName|string|false|none||TODO configs|
|dependsOn|[string]|false|none||TODO credential_spec|
|deploy|[docker.Deploy](#schemadocker.deploy)|false|none||TODO deploy|
|domainName|string|false|none||domain name|
|envFile|[string]|false|none||TODO devices<br />TODO dns<br />TODO dns_search<br />TODO entrypoint|
|environment|[[docker.Environment](#schemadocker.environment)]|false|none||environment variable|
|hostName|string|false|none||host name|
|image|[docker.Image](#schemadocker.image)|false|none||TODO expose<br />TODO external_links<br />TODO extra_hosts<br />TODO healthcheck|
|ipc|string|false|none||ipc|
|macAddress|string|false|none||macaddress|
|networkMode|string|false|none||TODO init<br />TODO isolation<br />TODO labels<br />TODO links<br />TODO logging|
|networks|object|false|none||Joined network|
|» **additionalProperties**|[docker.NetworkMap](#schemadocker.networkmap)|false|none||none|
|ports|[[docker.Port](#schemadocker.port)]|false|none||TODO pid|
|privileged|boolean|false|none||privilege|
|proflies|[string]|false|none||Service configuration file|
|readOnly|boolean|false|none||read-only|
|restart|string|false|none||Restart strategy|
|secrets|[[docker.Secret](#schemadocker.secret)]|false|none||secret key|
|shmSize|string|false|none||Shared Memory Size|
|stdinOpen|boolean|false|none||Standard input|
|tty|boolean|false|none||Is it enabledtty|
|user|string|false|none||Each of these is a single value, analogous to its docker run counterpart. Note that mac_address is a legacy option.|
|volumes|[[docker.VolumeMap](#schemadocker.volumemap)]|false|none||TODO security_opt<br />TODO stop_grace_period<br />TODO stop_signal<br />TODO sysctls<br />TODO tmpfs<br />TODO ulimits<br />TODO userns_mode|
|workingDir|string|false|none||working directory|

<h2 id="tocS_docker.Secret">docker.Secret</h2>

<a id="schemadocker.secret"></a>
<a id="schema_docker.Secret"></a>
<a id="tocSdocker.secret"></a>
<a id="tocsdocker.secret"></a>

```json
{
  "gid": "string",
  "mode": "string",
  "source": "string",
  "target": "string",
  "uid": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|gid|string|false|none||fileGID|
|mode|string|false|none||file right|
|source|string|false|none||name|
|target|string|false|none||file name|
|uid|string|false|none||fileUID|

<h2 id="tocS_docker.RollbackConfig">docker.RollbackConfig</h2>

<a id="schemadocker.rollbackconfig"></a>
<a id="schema_docker.RollbackConfig"></a>
<a id="tocSdocker.rollbackconfig"></a>
<a id="tocsdocker.rollbackconfig"></a>

```json
{
  "delay": 0,
  "failureAction": "string",
  "maxFailureRatio": 0,
  "monitor": 0,
  "order": "string",
  "parallelism": 0
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|delay|integer|false|none||delay Service Rollback Delay Service Rollback Delay，Service Rollback Delay 0 Service Rollback Delay 1 000 000 000|
|failureAction|string|false|none||failure_action Service rollback failure behavior Service rollback failure behavior，Service rollback failure behavior pause、continue|
|maxFailureRatio|number|false|none||max_failure_ratio Service rollback failure rate Service rollback failure rate，Service rollback failure rate 0.000 Service rollback failure rate 1.000|
|monitor|integer|false|none||monitor Service Rollback Monitoring Service Rollback Monitoring，Service Rollback Monitoring 0 Service Rollback Monitoring 1 000 000 000|
|order|string|false|none||order Service Rollback Order Service Rollback Order，Service Rollback Order start-first、stop-first|
|parallelism|integer|false|none||parallelism Number of service rollback parallels Number of service rollback parallels，Number of service rollback parallels 0 Number of service rollback parallels 1 000 000 000|

<h2 id="tocS_docker.RestartPolicy">docker.RestartPolicy</h2>

<a id="schemadocker.restartpolicy"></a>
<a id="schema_docker.RestartPolicy"></a>
<a id="tocSdocker.restartpolicy"></a>
<a id="tocsdocker.restartpolicy"></a>

```json
{
  "condition": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|condition|string|false|none||condition Service restart conditions Service restart conditions，Service restart conditions none、on-failure、any|

<h2 id="tocS_docker.Resources">docker.Resources</h2>

<a id="schemadocker.resources"></a>
<a id="schema_docker.Resources"></a>
<a id="tocSdocker.resources"></a>
<a id="tocsdocker.resources"></a>

```json
{
  "cpus": "string",
  "memory": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|cpus|string|false|none||cpus service CPU service service CPU service，service 0.000 service 1 000.000|
|memory|string|false|none||memory Service Memory Configuration Service Memory Configuration，Service Memory Configuration 4M Service Memory Configuration 1 000G|

<h2 id="tocS_docker.Port">docker.Port</h2>

<a id="schemadocker.port"></a>
<a id="schema_docker.Port"></a>
<a id="tocSdocker.port"></a>
<a id="tocsdocker.port"></a>

```json
{
  "desc": "string",
  "mode": "string",
  "protocol": "string",
  "published": "string",
  "target": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|desc|string|false|none||describe|
|mode|string|false|none||none|
|protocol|string|false|none||none|
|published|string|false|none||none|
|target|string|false|none||none|

<h2 id="tocS_docker.Pair">docker.Pair</h2>

<a id="schemadocker.pair"></a>
<a id="schema_docker.Pair"></a>
<a id="tocSdocker.pair"></a>
<a id="tocsdocker.pair"></a>

```json
{
  "name": "name",
  "value": "value"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||none|
|value|string|false|none||none|

<h2 id="tocS_docker.Notice">docker.Notice</h2>

<a id="schemadocker.notice"></a>
<a id="schema_docker.Notice"></a>
<a id="tocSdocker.notice"></a>
<a id="tocsdocker.notice"></a>

```json
{
  "host": "string",
  "method": "string",
  "query": [
    "string"
  ],
  "uri": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|host|string|false|none||ip/domain name|
|method|string|false|none||Request Type Request TypeGET、POST|
|query|[string]|false|none||parameter parameter version/details|
|uri|string|false|none||route|

<h2 id="tocS_docker.NetworkMap">docker.NetworkMap</h2>

<a id="schemadocker.networkmap"></a>
<a id="schema_docker.NetworkMap"></a>
<a id="tocSdocker.networkmap"></a>
<a id="tocsdocker.networkmap"></a>

```json
{
  "aliases": [
    "string"
  ],
  "ipv4Address": "string",
  "ipv6Address": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|aliases|[string]|false|none||Alias List|
|ipv4Address|string|false|none||ipv4address|
|ipv6Address|string|false|none||ipv6address|

<h2 id="tocS_docker.Network">docker.Network</h2>

<a id="schemadocker.network"></a>
<a id="schema_docker.Network"></a>
<a id="tocSdocker.network"></a>
<a id="tocsdocker.network"></a>

```json
{
  "external": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|external|boolean|false|none||TODO driver<br />TODO driver_opts<br />TODO attachable<br />TODO enable_ipv6<br />TODO ipam<br />TODO internal<br />TODO labels|

<h2 id="tocS_docker.Mode">docker.Mode</h2>

<a id="schemadocker.mode"></a>
<a id="schema_docker.Mode"></a>
<a id="tocSdocker.mode"></a>
<a id="tocsdocker.mode"></a>

```json
{
  "docker": true,
  "swarm": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|docker|boolean|false|none||Dockermode|
|swarm|boolean|false|none||Swarmmode|

<h2 id="tocS_docker.KubernetesConfig">docker.KubernetesConfig</h2>

<a id="schemadocker.kubernetesconfig"></a>
<a id="schema_docker.KubernetesConfig"></a>
<a id="tocSdocker.kubernetesconfig"></a>
<a id="tocsdocker.kubernetesconfig"></a>

```json
{
  "deployment": {
    "apiVersion": "string",
    "kind": "string",
    "metadata": {
      "annotations": {
        "property1": "string",
        "property2": "string"
      },
      "clusterName": "string",
      "creationTimestamp": "string",
      "deletionGracePeriodSeconds": 0,
      "deletionTimestamp": "string",
      "finalizers": [
        "string"
      ],
      "generateName": "string",
      "generation": 0,
      "labels": {
        "property1": "string",
        "property2": "string"
      },
      "managedFields": [
        {}
      ],
      "name": "string",
      "namespace": "string",
      "ownerReferences": [
        {}
      ],
      "resourceVersion": "string",
      "selfLink": "string",
      "uid": "string"
    },
    "spec": {
      "minReadySeconds": 0,
      "paused": true,
      "progressDeadlineSeconds": 0,
      "replicas": 0,
      "revisionHistoryLimit": 0,
      "selector": null,
      "strategy": null,
      "template": null
    },
    "status": {
      "availableReplicas": 0,
      "collisionCount": 0,
      "conditions": [
        {}
      ],
      "observedGeneration": 0,
      "readyReplicas": 0,
      "replicas": 0,
      "unavailableReplicas": 0,
      "updatedReplicas": 0
    }
  },
  "persistentVolumeClaim": {
    "apiVersion": "string",
    "kind": "string",
    "metadata": {
      "annotations": {
        "property1": "string",
        "property2": "string"
      },
      "clusterName": "string",
      "creationTimestamp": "string",
      "deletionGracePeriodSeconds": 0,
      "deletionTimestamp": "string",
      "finalizers": [
        "string"
      ],
      "generateName": "string",
      "generation": 0,
      "labels": {
        "property1": "string",
        "property2": "string"
      },
      "managedFields": [
        {}
      ],
      "name": "string",
      "namespace": "string",
      "ownerReferences": [
        {}
      ],
      "resourceVersion": "string",
      "selfLink": "string",
      "uid": "string"
    },
    "spec": {
      "accessModes": [
        "["
      ],
      "dataSource": null,
      "dataSourceRef": null,
      "resources": null,
      "selector": null,
      "storageClassName": "string",
      "volumeMode": null,
      "volumeName": "string"
    },
    "status": {
      "accessModes": [
        "["
      ],
      "capacity": null,
      "conditions": [
        {}
      ],
      "phase": null
    }
  },
  "service": {
    "apiVersion": "string",
    "kind": "string",
    "metadata": {
      "annotations": {
        "property1": "string",
        "property2": "string"
      },
      "clusterName": "string",
      "creationTimestamp": "string",
      "deletionGracePeriodSeconds": 0,
      "deletionTimestamp": "string",
      "finalizers": [
        "string"
      ],
      "generateName": "string",
      "generation": 0,
      "labels": {
        "property1": "string",
        "property2": "string"
      },
      "managedFields": [
        {}
      ],
      "name": "string",
      "namespace": "string",
      "ownerReferences": [
        {}
      ],
      "resourceVersion": "string",
      "selfLink": "string",
      "uid": "string"
    },
    "spec": {
      "allocateLoadBalancerNodePorts": true,
      "clusterIP": "string",
      "clusterIPs": [
        "string"
      ],
      "externalIPs": [
        "string"
      ],
      "externalName": "string",
      "externalTrafficPolicy": null,
      "healthCheckNodePort": 0,
      "internalTrafficPolicy": null,
      "ipFamilies": [
        "["
      ],
      "ipFamilyPolicy": null,
      "loadBalancerClass": "string",
      "loadBalancerIP": "string",
      "loadBalancerSourceRanges": [
        "string"
      ],
      "ports": [
        {}
      ],
      "publishNotReadyAddresses": true,
      "selector": {
        "property1": "string",
        "property2": "string"
      },
      "sessionAffinity": null,
      "sessionAffinityConfig": null,
      "type": null
    },
    "status": {
      "conditions": [
        {}
      ],
      "loadBalancer": null
    }
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|deployment|[v1.Deployment](#schemav1.deployment)|false|none||none|
|persistentVolumeClaim|[v1.PersistentVolumeClaim](#schemav1.persistentvolumeclaim)|false|none||none|
|service|[v1.Service](#schemav1.service)|false|none||none|

<h2 id="tocS_docker.Image">docker.Image</h2>

<a id="schemadocker.image"></a>
<a id="schema_docker.Image"></a>
<a id="tocSdocker.image"></a>
<a id="tocsdocker.image"></a>

```json
{
  "name": "string",
  "tag": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|name|string|false|none||Mirror Name|
|tag|string|false|none||Mirror label|

<h2 id="tocS_docker.GitlabRegistryData">docker.GitlabRegistryData</h2>

<a id="schemadocker.gitlabregistrydata"></a>
<a id="schema_docker.GitlabRegistryData"></a>
<a id="tocSdocker.gitlabregistrydata"></a>
<a id="tocsdocker.gitlabregistrydata"></a>

```json
{
  "instanceURL": "string",
  "projectID": 0,
  "projectPath": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|instanceURL|string|false|none||none|
|projectID|integer|false|none||none|
|projectPath|string|false|none||none|

<h2 id="tocS_docker.Environment">docker.Environment</h2>

<a id="schemadocker.environment"></a>
<a id="schema_docker.Environment"></a>
<a id="tocSdocker.environment"></a>
<a id="tocsdocker.environment"></a>

```json
{
  "desc": "string",
  "key": "string",
  "val": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|desc|string|false|none||none|
|key|string|false|none||none|
|val|string|false|none||none|

<h2 id="tocS_docker.EndpointSecuritySettings">docker.EndpointSecuritySettings</h2>

<a id="schemadocker.endpointsecuritysettings"></a>
<a id="schema_docker.EndpointSecuritySettings"></a>
<a id="tocSdocker.endpointsecuritysettings"></a>
<a id="tocsdocker.endpointsecuritysettings"></a>

```json
{
  "allowBindMountsForRegularUsers": false,
  "allowContainerCapabilitiesForRegularUsers": true,
  "allowDeviceMappingForRegularUsers": true,
  "allowHostNamespaceForRegularUsers": true,
  "allowPrivilegedModeForRegularUsers": false,
  "allowStackManagementForRegularUsers": true,
  "allowSysctlSettingForRegularUsers": true,
  "allowVolumeBrowserForRegularUsers": true,
  "enableHostManagementFeatures": true
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|allowBindMountsForRegularUsers|boolean|false|none||Whether non-administrator should be able to use bind mounts when creating containers<br />Should non administrators be able to use binding mount when creating containers|
|allowContainerCapabilitiesForRegularUsers|boolean|false|none||Whether non-administrator should be able to use container capabilities<br />Should non administrators be able to use container functionality|
|allowDeviceMappingForRegularUsers|boolean|false|none||Whether non-administrator should be able to use device mapping<br />Should non administrators be able to use device mapping|
|allowHostNamespaceForRegularUsers|boolean|false|none||Whether non-administrator should be able to use the host pid<br />Should non administrators be able to use the hostpid|
|allowPrivilegedModeForRegularUsers|boolean|false|none||Whether non-administrator should be able to use privileged mode when creating containers<br />Should non administrators be able to use privileged mode when creating containers|
|allowStackManagementForRegularUsers|boolean|false|none||Whether non-administrator should be able to manage stacks<br />Should non administrators be able to manage the stack|
|allowSysctlSettingForRegularUsers|boolean|false|none||Whether non-administrator should be able to use sysctl settings<br />Should non administrators be able to usesysctlShould non administrators be able to use|
|allowVolumeBrowserForRegularUsers|boolean|false|none||Whether non-administrator should be able to browse volumes<br />Should non administrators be able to browse volumes|
|enableHostManagementFeatures|boolean|false|none||Whether host management features are enabled<br />Enable host management function|

<h2 id="tocS_docker.DocTemplateConfig">docker.DocTemplateConfig</h2>

<a id="schemadocker.doctemplateconfig"></a>
<a id="schema_docker.DocTemplateConfig"></a>
<a id="tocSdocker.doctemplateconfig"></a>
<a id="tocsdocker.doctemplateconfig"></a>

```json
{
  "createAt": "string",
  "dockerConfig": {
    "property1": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    },
    "property2": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    }
  },
  "dockerImage": "string",
  "dockerRegistry": {
    "AccessToken": "string",
    "AccessTokenExpiry": 0,
    "authentication": true,
    "baseURL": "string",
    "createAt": "string",
    "gitlab": {
      "instanceURL": null,
      "projectID": null,
      "projectPath": null
    },
    "id": "string",
    "name": "string",
    "password": "string",
    "registryType": 1,
    "updateAt": "string",
    "url": "string",
    "userId": "string",
    "username": "string"
  },
  "fileSize": 0,
  "id": "string",
  "isDel": true,
  "k8SConfig": {
    "deployment": {
      "apiVersion": "string",
      "kind": "string",
      "metadata": null,
      "spec": null,
      "status": null
    },
    "persistentVolumeClaim": {
      "apiVersion": "string",
      "kind": "string",
      "metadata": null,
      "spec": null,
      "status": null
    },
    "service": {
      "apiVersion": "string",
      "kind": "string",
      "metadata": null,
      "spec": null,
      "status": null
    }
  },
  "k8SContent": "string",
  "k8SRegistry": [
    {
      "AccessToken": "string",
      "AccessTokenExpiry": 0,
      "authentication": true,
      "baseURL": "string",
      "createAt": "string",
      "gitlab": {
        "instanceURL": "string",
        "projectID": 0,
        "projectPath": "string"
      },
      "id": "string",
      "name": "string",
      "password": "string",
      "registryType": 1,
      "updateAt": "string",
      "url": "string",
      "userId": "string",
      "username": "string"
    }
  ],
  "manual": {
    "createAt": "string",
    "id": "string",
    "name": "string",
    "originalName": "string",
    "size": 0,
    "suffix": "string",
    "updateAt": "string",
    "url": "string"
  },
  "platform": 0,
  "swarmConfig": {
    "property1": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    },
    "property2": {
      "environment": [
        {
          "desc": "string",
          "key": "string",
          "val": "string"
        }
      ],
      "hostname": "string",
      "ports": [
        {
          "desc": "string",
          "mode": "string",
          "protocol": "string",
          "published": "string",
          "target": "string"
        }
      ],
      "volume": [
        {
          "desc": "string",
          "readOnly": true,
          "source": "string",
          "target": "string",
          "type": "string"
        }
      ]
    }
  },
  "swarmContent": "string",
  "swarmRegistry": [
    {
      "AccessToken": "string",
      "AccessTokenExpiry": 0,
      "authentication": true,
      "baseURL": "string",
      "createAt": "string",
      "gitlab": {
        "instanceURL": "string",
        "projectID": 0,
        "projectPath": "string"
      },
      "id": "string",
      "name": "string",
      "password": "string",
      "registryType": 1,
      "updateAt": "string",
      "url": "string",
      "userId": "string",
      "username": "string"
    }
  ],
  "templateId": "string",
  "updateAt": "string",
  "userId": "string",
  "version": "string",
  "versionContent": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|createAt|string|false|none||none|
|dockerConfig|object|false|none||Dockerconfiguration information configuration informationdefault|
|» **additionalProperties**|[docker.TemplateConfig](#schemadocker.templateconfig)|false|none||none|
|dockerImage|string|false|none||DockerImage image|
|dockerRegistry|[docker.DocRegistry](#schemadocker.docregistry)|false|none||warehouseID|
|fileSize|number|false|none||Software size|
|id|string|false|none||Default parameters _id Default parameters|
|isDel|boolean|false|none||Delete or not|
|k8SConfig|[docker.KubernetesConfig](#schemadocker.kubernetesconfig)|false|none||Swarmreplace content replace content|
|k8SContent|string|false|none||SwarmFile Content|
|k8SRegistry|[[docker.DocRegistry](#schemadocker.docregistry)]|false|none||SwarmwarehouseID|
|manual|[system.SysBaseFiles](#schemasystem.sysbasefiles)|false|none||User Manual|
|platform|integer|false|none||(Reserved)Reserved|
|swarmConfig|object|false|none||Swarmreplace content replace content|
|» **additionalProperties**|[docker.TemplateConfig](#schemadocker.templateconfig)|false|none||none|
|swarmContent|string|false|none||SwarmFile Content|
|swarmRegistry|[[docker.DocRegistry](#schemadocker.docregistry)]|false|none||SwarmwarehouseID|
|templateId|string|false|none||TemplateID|
|updateAt|string|false|none||none|
|userId|string|false|none||userID|
|version|string|false|none||Version number|
|versionContent|string|false|none||Version Content|

<h2 id="tocS_docker.DocTemplate">docker.DocTemplate</h2>

<a id="schemadocker.doctemplate"></a>
<a id="schema_docker.DocTemplate"></a>
<a id="tocSdocker.doctemplate"></a>
<a id="tocsdocker.doctemplate"></a>

```json
{
  "category": [
    "string"
  ],
  "config": {
    "createAt": "string",
    "dockerConfig": {
      "property1": {
        "environment": [
          null
        ],
        "hostname": "string",
        "ports": [
          null
        ],
        "volume": [
          null
        ]
      },
      "property2": {
        "environment": [
          null
        ],
        "hostname": "string",
        "ports": [
          null
        ],
        "volume": [
          null
        ]
      }
    },
    "dockerImage": "string",
    "dockerRegistry": {
      "AccessToken": null,
      "AccessTokenExpiry": null,
      "authentication": null,
      "baseURL": null,
      "createAt": null,
      "gitlab": null,
      "id": null,
      "name": null,
      "password": null,
      "registryType": null,
      "updateAt": null,
      "url": null,
      "userId": null,
      "username": null
    },
    "fileSize": 0,
    "id": "string",
    "isDel": true,
    "k8SConfig": {
      "deployment": null,
      "persistentVolumeClaim": null,
      "service": null
    },
    "k8SContent": "string",
    "k8SRegistry": [
      {
        "AccessToken": "string",
        "AccessTokenExpiry": 0,
        "authentication": true,
        "baseURL": "string",
        "createAt": "string",
        "gitlab": null,
        "id": "string",
        "name": "string",
        "password": "string",
        "registryType": "[",
        "updateAt": "string",
        "url": "string",
        "userId": "string",
        "username": "string"
      }
    ],
    "manual": {
      "createAt": null,
      "id": null,
      "name": null,
      "originalName": null,
      "size": null,
      "suffix": null,
      "updateAt": null,
      "url": null
    },
    "platform": 0,
    "swarmConfig": {
      "property1": {
        "environment": [
          null
        ],
        "hostname": "string",
        "ports": [
          null
        ],
        "volume": [
          null
        ]
      },
      "property2": {
        "environment": [
          null
        ],
        "hostname": "string",
        "ports": [
          null
        ],
        "volume": [
          null
        ]
      }
    },
    "swarmContent": "string",
    "swarmRegistry": [
      {
        "AccessToken": "string",
        "AccessTokenExpiry": 0,
        "authentication": true,
        "baseURL": "string",
        "createAt": "string",
        "gitlab": null,
        "id": "string",
        "name": "string",
        "password": "string",
        "registryType": "[",
        "updateAt": "string",
        "url": "string",
        "userId": "string",
        "username": "string"
      }
    ],
    "templateId": "string",
    "updateAt": "string",
    "userId": "string",
    "version": "string",
    "versionContent": "string"
  },
  "copyright": "string",
  "createAt": "string",
  "devLanguage": [
    "string"
  ],
  "downloads": 0,
  "features": "string",
  "id": "string",
  "isDel": true,
  "logo": "string",
  "name": "string",
  "nameZh": "string",
  "notice": {
    "host": "string",
    "method": "string",
    "query": [
      "string"
    ],
    "uri": "string"
  },
  "poster": [
    {
      "createAt": "string",
      "id": "string",
      "name": "string",
      "originalName": "string",
      "size": 0,
      "suffix": "string",
      "updateAt": "string",
      "url": "string"
    }
  ],
  "preview": [
    {
      "createAt": "string",
      "id": "string",
      "name": "string",
      "originalName": "string",
      "size": 0,
      "suffix": "string",
      "updateAt": "string",
      "url": "string"
    }
  ],
  "relationSoft": [
    {
      "category": [
        "string"
      ],
      "config": {
        "createAt": "string",
        "dockerConfig": {},
        "dockerImage": "string",
        "dockerRegistry": null,
        "fileSize": 0,
        "id": "string",
        "isDel": true,
        "k8SConfig": null,
        "k8SContent": "string",
        "k8SRegistry": [
          null
        ],
        "manual": null,
        "platform": 0,
        "swarmConfig": {},
        "swarmContent": "string",
        "swarmRegistry": [
          null
        ],
        "templateId": "string",
        "updateAt": "string",
        "userId": "string",
        "version": "string",
        "versionContent": "string"
      },
      "copyright": "string",
      "createAt": "string",
      "devLanguage": [
        "string"
      ],
      "downloads": 0,
      "features": "string",
      "id": "string",
      "isDel": true,
      "logo": "string",
      "name": "string",
      "nameZh": "string",
      "notice": {
        "host": "string",
        "method": "string",
        "query": [
          null
        ],
        "uri": "string"
      },
      "poster": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "preview": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "relationSoft": [
        {
          "category": [
            "string"
          ],
          "config": null,
          "copyright": "string",
          "createAt": "string",
          "devLanguage": [
            "string"
          ],
          "downloads": 0,
          "features": "string",
          "id": "string",
          "isDel": true,
          "logo": "string",
          "name": "string",
          "nameZh": "string",
          "notice": null,
          "poster": [
            {}
          ],
          "preview": [
            {}
          ],
          "relationSoft": [
            {}
          ],
          "relationSoftIds": [
            "string"
          ],
          "score": 0,
          "state": 0,
          "supportMode": null,
          "sysLanguage": [
            "string"
          ],
          "team": "string",
          "teamIcon": "string",
          "updateAt": "string",
          "userId": "string"
        }
      ],
      "relationSoftIds": [
        "string"
      ],
      "score": 0,
      "state": 0,
      "supportMode": {
        "docker": true,
        "swarm": true
      },
      "sysLanguage": [
        "string"
      ],
      "team": "string",
      "teamIcon": "string",
      "updateAt": "string",
      "userId": "string"
    }
  ],
  "relationSoftIds": [
    "string"
  ],
  "score": 0,
  "state": 0,
  "supportMode": {
    "docker": true,
    "swarm": true
  },
  "sysLanguage": [
    "string"
  ],
  "team": "string",
  "teamIcon": "string",
  "updateAt": "string",
  "userId": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|category|[string]|false|none||Template classification|
|config|[docker.DocTemplateConfig](#schemadocker.doctemplateconfig)|false|none||configuration information configuration information|
|copyright|string|false|none||Copyright Notice|
|createAt|string|false|none||none|
|devLanguage|[string]|false|none||development language|
|downloads|integer|false|none||Number of downloads|
|features|string|false|none||Template Properties|
|id|string|false|none||Default parameters _id Default parameters|
|isDel|boolean|false|none||Delete or not|
|logo|string|true|none||softwarelogo softwarehttpsoftware|
|name|string|true|none||Software Name|
|nameZh|string|false|none||Software name in Chinese|
|notice|[docker.Notice](#schemadocker.notice)|false|none||Message notification|
|poster|[[system.SysBaseFiles](#schemasystem.sysbasefiles)]|false|none||poster|
|preview|[[system.SysBaseFiles](#schemasystem.sysbasefiles)]|false|none||preview|
|relationSoft|[[docker.DocTemplate](#schemadocker.doctemplate)]|false|none||Related software|
|relationSoftIds|[string]|false|none||Related software Related software|
|score|number|false|none||Average score|
|state|integer|false|none||Software status 1. Software status 2. Software status 3. Software status 4. Software status|
|supportMode|[docker.Mode](#schemadocker.mode)|false|none||Software supported modes|
|sysLanguage|[string]|false|none||System Language|
|team|string|false|none||the development team|
|teamIcon|string|false|none||the development teamIcon|
|updateAt|string|false|none||none|
|userId|string|false|none||userID|

<h2 id="tocS_docker.DocSoftware">docker.DocSoftware</h2>

<a id="schemadocker.docsoftware"></a>
<a id="schema_docker.DocSoftware"></a>
<a id="tocSdocker.docsoftware"></a>
<a id="tocsdocker.docsoftware"></a>

```json
{
  "createAt": "string",
  "dockerYml": {
    "networks": {
      "property1": {
        "external": true
      },
      "property2": {
        "external": true
      }
    },
    "secrets": {
      "property1": {
        "external": "string",
        "file": "string",
        "name": "string"
      },
      "property2": {
        "external": "string",
        "file": "string",
        "name": "string"
      }
    },
    "services": {
      "property1": {
        "build": {},
        "capAdd": [
          "string"
        ],
        "capDrop": [
          "string"
        ],
        "cgroupParent": "string",
        "command": [
          "string"
        ],
        "containerName": "string",
        "dependsOn": [
          "string"
        ],
        "deploy": {},
        "domainName": "string",
        "envFile": [
          "string"
        ],
        "environment": [
          {
            "desc": null,
            "key": null,
            "val": null
          }
        ],
        "hostName": "string",
        "image": {},
        "ipc": "string",
        "macAddress": "string",
        "networkMode": "string",
        "networks": {
          "property1": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          },
          "property2": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          }
        },
        "ports": [
          {
            "desc": null,
            "mode": null,
            "protocol": null,
            "published": null,
            "target": null
          }
        ],
        "privileged": true,
        "proflies": [
          "string"
        ],
        "readOnly": true,
        "restart": "string",
        "secrets": [
          {
            "gid": null,
            "mode": null,
            "source": null,
            "target": null,
            "uid": null
          }
        ],
        "shmSize": "string",
        "stdinOpen": true,
        "tty": true,
        "user": "string",
        "volumes": [
          {
            "desc": null,
            "readOnly": null,
            "source": null,
            "target": null,
            "type": null
          }
        ],
        "workingDir": "string"
      },
      "property2": {
        "build": {},
        "capAdd": [
          "string"
        ],
        "capDrop": [
          "string"
        ],
        "cgroupParent": "string",
        "command": [
          "string"
        ],
        "containerName": "string",
        "dependsOn": [
          "string"
        ],
        "deploy": {},
        "domainName": "string",
        "envFile": [
          "string"
        ],
        "environment": [
          {
            "desc": null,
            "key": null,
            "val": null
          }
        ],
        "hostName": "string",
        "image": {},
        "ipc": "string",
        "macAddress": "string",
        "networkMode": "string",
        "networks": {
          "property1": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          },
          "property2": {
            "aliases": null,
            "ipv4Address": null,
            "ipv6Address": null
          }
        },
        "ports": [
          {
            "desc": null,
            "mode": null,
            "protocol": null,
            "published": null,
            "target": null
          }
        ],
        "privileged": true,
        "proflies": [
          "string"
        ],
        "readOnly": true,
        "restart": "string",
        "secrets": [
          {
            "gid": null,
            "mode": null,
            "source": null,
            "target": null,
            "uid": null
          }
        ],
        "shmSize": "string",
        "stdinOpen": true,
        "tty": true,
        "user": "string",
        "volumes": [
          {
            "desc": null,
            "readOnly": null,
            "source": null,
            "target": null,
            "type": null
          }
        ],
        "workingDir": "string"
      }
    },
    "version": "string",
    "volumes": {
      "property1": {
        "external": true
      },
      "property2": {
        "external": true
      }
    }
  },
  "endpoint": {
    "TLSConfig": {
      "TLSCACertPath": null,
      "TLSCertPath": null,
      "TLSKeyPath": null,
      "TLSType": null
    },
    "connected": true,
    "createAt": "string",
    "envURL": "string",
    "id": "string",
    "isSocket": true,
    "isTLS": true,
    "k8sConfig": "string",
    "name": "string",
    "password": "string",
    "securitySettings": {
      "allowBindMountsForRegularUsers": null,
      "allowContainerCapabilitiesForRegularUsers": null,
      "allowDeviceMappingForRegularUsers": null,
      "allowHostNamespaceForRegularUsers": null,
      "allowPrivilegedModeForRegularUsers": null,
      "allowStackManagementForRegularUsers": null,
      "allowSysctlSettingForRegularUsers": null,
      "allowVolumeBrowserForRegularUsers": null,
      "enableHostManagementFeatures": null
    },
    "status": 1,
    "templates": [
      {
        "category": [
          null
        ],
        "config": null,
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          null
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": null,
        "poster": [
          null
        ],
        "preview": [
          null
        ],
        "relationSoft": [
          null
        ],
        "relationSoftIds": [
          null
        ],
        "score": 0,
        "state": 0,
        "supportMode": null,
        "sysLanguage": [
          null
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "type": 0,
    "uniqueId": "string",
    "updateAt": "string",
    "userId": "string",
    "username": "string"
  },
  "entryPoint": "string",
  "env": [
    {
      "name": "name",
      "value": "value"
    }
  ],
  "homeHost": "string",
  "homePort": "string",
  "id": "string",
  "name": "string",
  "projectPath": "string",
  "status": "string",
  "template": {
    "category": [
      "string"
    ],
    "config": {
      "createAt": null,
      "dockerConfig": null,
      "dockerImage": null,
      "dockerRegistry": null,
      "fileSize": null,
      "id": null,
      "isDel": null,
      "k8SConfig": null,
      "k8SContent": null,
      "k8SRegistry": null,
      "manual": null,
      "platform": null,
      "swarmConfig": null,
      "swarmContent": null,
      "swarmRegistry": null,
      "templateId": null,
      "updateAt": null,
      "userId": null,
      "version": null,
      "versionContent": null
    },
    "copyright": "string",
    "createAt": "string",
    "devLanguage": [
      "string"
    ],
    "downloads": 0,
    "features": "string",
    "id": "string",
    "isDel": true,
    "logo": "string",
    "name": "string",
    "nameZh": "string",
    "notice": {
      "host": null,
      "method": null,
      "query": null,
      "uri": null
    },
    "poster": [
      {
        "createAt": "string",
        "id": "string",
        "name": "string",
        "originalName": "string",
        "size": 0,
        "suffix": "string",
        "updateAt": "string",
        "url": "string"
      }
    ],
    "preview": [
      {
        "createAt": "string",
        "id": "string",
        "name": "string",
        "originalName": "string",
        "size": 0,
        "suffix": "string",
        "updateAt": "string",
        "url": "string"
      }
    ],
    "relationSoft": [
      {
        "category": [
          null
        ],
        "config": null,
        "copyright": "string",
        "createAt": "string",
        "devLanguage": [
          null
        ],
        "downloads": 0,
        "features": "string",
        "id": "string",
        "isDel": true,
        "logo": "string",
        "name": "string",
        "nameZh": "string",
        "notice": null,
        "poster": [
          null
        ],
        "preview": [
          null
        ],
        "relationSoft": [
          null
        ],
        "relationSoftIds": [
          null
        ],
        "score": 0,
        "state": 0,
        "supportMode": null,
        "sysLanguage": [
          null
        ],
        "team": "string",
        "teamIcon": "string",
        "updateAt": "string",
        "userId": "string"
      }
    ],
    "relationSoftIds": [
      "string"
    ],
    "score": 0,
    "state": 0,
    "supportMode": {
      "docker": null,
      "swarm": null
    },
    "sysLanguage": [
      "string"
    ],
    "team": "string",
    "teamIcon": "string",
    "updateAt": "string",
    "userId": "string"
  },
  "updateAt": "string",
  "userId": "string",
  "userIp": "string",
  "userIp2Region": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|createAt|string|false|none||none|
|dockerYml|[docker.Yml](#schemadocker.yml)|false|none||none|
|endpoint|[docker.DocEndpoint](#schemadocker.docendpoint)|false|none||environment|
|entryPoint|string|false|none||none|
|env|[[docker.Pair](#schemadocker.pair)]|false|none||none|
|homeHost|string|false|none||Home page address|
|homePort|string|false|none||Homepage Port|
|id|string|false|none||Default parameters _id Default parameters|
|name|string|false|none||The name of the software installation|
|projectPath|string|false|none||The following are all temporary configurations|
|status|string|false|none||running state|
|template|[docker.DocTemplate](#schemadocker.doctemplate)|false|none||Template|
|updateAt|string|false|none||none|
|userId|string|false|none||userID|
|userIp|string|false|none||userIP|
|userIp2Region|string|false|none||userIP2REGIN|

<h2 id="tocS_docker.DocRegistry">docker.DocRegistry</h2>

<a id="schemadocker.docregistry"></a>
<a id="schema_docker.DocRegistry"></a>
<a id="tocSdocker.docregistry"></a>
<a id="tocsdocker.docregistry"></a>

```json
{
  "AccessToken": "string",
  "AccessTokenExpiry": 0,
  "authentication": true,
  "baseURL": "string",
  "createAt": "string",
  "gitlab": {
    "instanceURL": "string",
    "projectID": 0,
    "projectPath": "string"
  },
  "id": "string",
  "name": "string",
  "password": "string",
  "registryType": 1,
  "updateAt": "string",
  "url": "string",
  "userId": "string",
  "username": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|AccessToken|string|false|none||Store temporary access token(Store temporary access token)|
|AccessTokenExpiry|integer|false|none||none|
|authentication|boolean|false|none||Authentication Whether to enable authentication|
|baseURL|string|false|none||BaseURL in order toProGetin order to|
|createAt|string|false|none||none|
|gitlab|[docker.GitlabRegistryData](#schemadocker.gitlabregistrydata)|false|none||Additional fields for other warehouse types need to be stored in the middle(Additional fields for other warehouse types need to be stored in the middle，Additional fields for other warehouse types need to be stored in the middle3，6Additional fields for other warehouse types need to be stored in the middle)|
|id|string|false|none||Default parameters _id Default parameters|
|name|string|false|none||Name Warehouse alias Warehouse alias|
|password|string|false|none||Password password|
|registryType|integer|false|none||Type Warehouse type<br />Registry Type. Valid values are:<br />	1 (Quay.io), // Not yet developed quay.io Not yet developed<br />	2 (Azure container registry), // Not yet developed<br />	3 (custom registry), // Self built warehouse, Self built warehouse<br />	4 (Gitlab registry), // Not yet developed gitlab Not yet developed<br />	5 (ProGet registry), // Prepare for development<br />	6 (DockerHub) // docker hub<br />	7 (ECR) // Not yet developed aws ECR Not yet developed|
|updateAt|string|false|none||none|
|url|string|false|none||URL or IP Warehouse address Warehouse address Warehouse address|
|userId|string|false|none||Userid Reserved Usersid，Reserved Users|
|username|string|false|none||Username user name|

#### enum 

|attribute|attribute|
|---|---|
|registryType|1|
|registryType|2|
|registryType|3|
|registryType|4|
|registryType|5|
|registryType|6|
|registryType|7|

<h2 id="tocS_docker.DocEndpoint">docker.DocEndpoint</h2>

<a id="schemadocker.docendpoint"></a>
<a id="schema_docker.DocEndpoint"></a>
<a id="tocSdocker.docendpoint"></a>
<a id="tocsdocker.docendpoint"></a>

```json
{
  "TLSConfig": {
    "TLSCACertPath": "string",
    "TLSCertPath": "string",
    "TLSKeyPath": "string",
    "TLSType": 1
  },
  "connected": true,
  "createAt": "string",
  "envURL": "string",
  "id": "string",
  "isSocket": true,
  "isTLS": true,
  "k8sConfig": "string",
  "name": "string",
  "password": "string",
  "securitySettings": {
    "allowBindMountsForRegularUsers": false,
    "allowContainerCapabilitiesForRegularUsers": true,
    "allowDeviceMappingForRegularUsers": true,
    "allowHostNamespaceForRegularUsers": true,
    "allowPrivilegedModeForRegularUsers": false,
    "allowStackManagementForRegularUsers": true,
    "allowSysctlSettingForRegularUsers": true,
    "allowVolumeBrowserForRegularUsers": true,
    "enableHostManagementFeatures": true
  },
  "status": 1,
  "templates": [
    {
      "category": [
        "string"
      ],
      "config": {
        "createAt": "string",
        "dockerConfig": {},
        "dockerImage": "string",
        "dockerRegistry": null,
        "fileSize": 0,
        "id": "string",
        "isDel": true,
        "k8SConfig": null,
        "k8SContent": "string",
        "k8SRegistry": [
          null
        ],
        "manual": null,
        "platform": 0,
        "swarmConfig": {},
        "swarmContent": "string",
        "swarmRegistry": [
          null
        ],
        "templateId": "string",
        "updateAt": "string",
        "userId": "string",
        "version": "string",
        "versionContent": "string"
      },
      "copyright": "string",
      "createAt": "string",
      "devLanguage": [
        "string"
      ],
      "downloads": 0,
      "features": "string",
      "id": "string",
      "isDel": true,
      "logo": "string",
      "name": "string",
      "nameZh": "string",
      "notice": {
        "host": "string",
        "method": "string",
        "query": [
          null
        ],
        "uri": "string"
      },
      "poster": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "preview": [
        {
          "createAt": "string",
          "id": "string",
          "name": "string",
          "originalName": "string",
          "size": 0,
          "suffix": "string",
          "updateAt": "string",
          "url": "string"
        }
      ],
      "relationSoft": [
        {
          "category": [
            "string"
          ],
          "config": null,
          "copyright": "string",
          "createAt": "string",
          "devLanguage": [
            "string"
          ],
          "downloads": 0,
          "features": "string",
          "id": "string",
          "isDel": true,
          "logo": "string",
          "name": "string",
          "nameZh": "string",
          "notice": null,
          "poster": [
            {}
          ],
          "preview": [
            {}
          ],
          "relationSoft": [
            {}
          ],
          "relationSoftIds": [
            "string"
          ],
          "score": 0,
          "state": 0,
          "supportMode": null,
          "sysLanguage": [
            "string"
          ],
          "team": "string",
          "teamIcon": "string",
          "updateAt": "string",
          "userId": "string"
        }
      ],
      "relationSoftIds": [
        "string"
      ],
      "score": 0,
      "state": 0,
      "supportMode": {
        "docker": true,
        "swarm": true
      },
      "sysLanguage": [
        "string"
      ],
      "team": "string",
      "teamIcon": "string",
      "updateAt": "string",
      "userId": "string"
    }
  ],
  "type": 0,
  "uniqueId": "string",
  "updateAt": "string",
  "userId": "string",
  "username": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|TLSConfig|[docker.TLSConfiguration](#schemadocker.tlsconfiguration)|false|none||TLSConfig TLS allocation|
|connected|boolean|false|none||Connection status|
|createAt|string|false|none||none|
|envURL|string|false|none||URL or IP Related to this environment（Related to this environment）Related to this environmentDockerRelated to this environmentURLRelated to this environmentIPRelated to this environment|
|id|string|false|none||Default parameters _id Default parameters|
|isSocket|boolean|false|none||Whether to useSocket|
|isTLS|boolean|false|none||Whether to useTLS|
|k8sConfig|string|false|none||K8sConfig Kubernetes allocation|
|name|string|false|none||Environment(Endpoint) name|
|password|string|false|none||Server password|
|securitySettings|[docker.EndpointSecuritySettings](#schemadocker.endpointsecuritysettings)|false|none||Environment(Endpoint) Specific security settings|
|status|integer|false|none||Environmental status（Environmental status）(1 - up, 2 - down)|
|templates|[[docker.DocTemplate](#schemadocker.doctemplate)]|false|none||All templates // All templates All templates|
|type|integer|false|none||Environment(Endpoint) 1. remotedocker api|
|uniqueId|string|false|none||Server Unique Identification|
|updateAt|string|false|none||none|
|userId|string|false|none||userID|
|username|string|false|none||Server username|

<h2 id="tocS_docker.Deploy">docker.Deploy</h2>

<a id="schemadocker.deploy"></a>
<a id="schema_docker.Deploy"></a>
<a id="tocSdocker.deploy"></a>
<a id="tocsdocker.deploy"></a>

```json
{
  "endpointMode": "string",
  "labels": [
    "string"
  ],
  "maxReplicasPerNode": 0,
  "mode": "string",
  "placement": [
    {
      "property1": "string",
      "property2": "string"
    }
  ],
  "replicas": 0,
  "resources": {
    "cpus": "string",
    "memory": "string"
  },
  "restartPolicy": {
    "condition": "string"
  },
  "rollbackConfig": {
    "delay": 0,
    "failureAction": "string",
    "maxFailureRatio": 0,
    "monitor": 0,
    "order": "string",
    "parallelism": 0
  },
  "updateConfig": {
    "delay": 0,
    "failureAction": "string",
    "maxFailureRatio": 0,
    "monitor": 0,
    "order": "string",
    "parallelism": 0
  }
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|endpointMode|string|false|none||endpoint_mode EndPoint Model  EndPoint Model ，EndPoint Model  vip EndPoint Model  dnsrr|
|labels|[string]|false|none||labels Service Tag|
|maxReplicasPerNode|integer|false|none||max_replicas_per_node Service deployment strategy Service deployment strategy，Service deployment strategy max_replicas_per_node Service deployment strategy max_replicas_per_node|
|mode|string|false|none||mode Service mode Service mode，Service mode replicated Service mode global|
|placement|[object]|false|none||placement Service deployment strategy Service deployment strategy，Service deployment strategy max_replicas_per_node Service deployment strategy max_replicas_per_node|
|» **additionalProperties**|string|false|none||none|
|replicas|integer|false|none||replicas Number of service replicas Number of service replicas，Number of service replicas 0 Number of service replicas 1 000 000 000|
|resources|[docker.Resources](#schemadocker.resources)|false|none||resources Service resource allocation Service resource allocation|
|restartPolicy|[docker.RestartPolicy](#schemadocker.restartpolicy)|false|none||restart_policy Service restart strategy Service restart strategy|
|rollbackConfig|[docker.RollbackConfig](#schemadocker.rollbackconfig)|false|none||rollback_config Service Rollback Configuration Service Rollback Configuration|
|updateConfig|[docker.UpdateConfig](#schemadocker.updateconfig)|false|none||update_config Service Update Configuration Service Update Configuration|

<h2 id="tocS_docker.Build">docker.Build</h2>

<a id="schemadocker.build"></a>
<a id="schema_docker.Build"></a>
<a id="tocSdocker.build"></a>
<a id="tocsdocker.build"></a>

```json
{
  "args": {},
  "cacheFrom": [
    {
      "name": "string",
      "tag": "string"
    }
  ],
  "context": "string",
  "dockerfile": "string",
  "labels": {},
  "shmSize": "string",
  "target": "string"
}

```

### attribute

|name|name|name|name|name|name|
|---|---|---|---|---|---|
|args|object|false|none||DockerfileDefined inARGDefined in|
|cacheFrom|[[docker.Image](#schemadocker.image)]|false|none||(v3.2+) Cached Mirror List|
|context|string|false|none||A containingDockerfileA containing，A containingGitA containingURL|
|dockerfile|string|false|none||AlternativeDockerfile|
|labels|object|false|none||(v3.3+) The metadata information of the target image|
|shmSize|string|false|none||none|
|target|string|false|none||(v3.4+) structureDockerfilestructureStage|

