## <a name="resource-messages">Messages</a>

Stability: `prototype`

FIXME

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **created_at** | *date-time* | when messages was created | `"2015-01-01T12:00:00Z"` |
| **id** | *uuid* | unique identifier of messages | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **name** | *string* | unique name of messages | `"example"` |
| **updated_at** | *date-time* | when messages was updated | `"2015-01-01T12:00:00Z"` |

### Messages Create

Create a new messages.

```
POST /messagess
```


#### Curl Example

```bash
$ curl -n -X POST https://api.airshooter.com/messagess \
  -d '{
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Messages Delete

Delete an existing messages.

```
DELETE /messagess/{messages_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X DELETE https://api.airshooter.com/messagess/$MESSAGES_ID_OR_NAME \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Messages Info

Info for existing messages.

```
GET /messagess/{messages_id_or_name}
```


#### Curl Example

```bash
$ curl -n https://api.airshooter.com/messagess/$MESSAGES_ID_OR_NAME
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Messages List

List existing messagess.

```
GET /messagess
```


#### Curl Example

```bash
$ curl -n https://api.airshooter.com/messagess
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "created_at": "2015-01-01T12:00:00Z",
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "example",
    "updated_at": "2015-01-01T12:00:00Z"
  }
]
```

### Messages Update

Update an existing messages.

```
PATCH /messagess/{messages_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X PATCH https://api.airshooter.com/messagess/$MESSAGES_ID_OR_NAME \
  -d '{
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```


## <a name="resource-notifications">Notification API</a>

Stability: `prototype`

Managing notifications, unread/read histories.

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **created_at** | *date-time* | when notifications was created | `"2015-01-01T12:00:00Z"` |
| **id** | *uuid* | unique identifier of notifications | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **name** | *string* | unique name of notifications | `"example"` |
| **updated_at** | *date-time* | when notifications was updated | `"2015-01-01T12:00:00Z"` |

### Notification API Create

Create a new notifications.

```
POST /notificationss
```


#### Curl Example

```bash
$ curl -n -X POST https://api.airshooter.com/notificationss \
  -d '{
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Notification API Delete

Delete an existing notifications.

```
DELETE /notificationss/{notifications_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X DELETE https://api.airshooter.com/notificationss/$NOTIFICATIONS_ID_OR_NAME \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Notification API Info

Info for existing notifications.

```
GET /notificationss/{notifications_id_or_name}
```


#### Curl Example

```bash
$ curl -n https://api.airshooter.com/notificationss/$NOTIFICATIONS_ID_OR_NAME
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Notification API List

List existing notificationss.

```
GET /notificationss
```


#### Curl Example

```bash
$ curl -n https://api.airshooter.com/notificationss
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "created_at": "2015-01-01T12:00:00Z",
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "example",
    "updated_at": "2015-01-01T12:00:00Z"
  }
]
```

### Notification API Update

Update an existing notifications.

```
PATCH /notificationss/{notifications_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X PATCH https://api.airshooter.com/notificationss/$NOTIFICATIONS_ID_OR_NAME \
  -d '{
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```


## <a name="resource-struct">Struct</a>

Stability: `prototype`

FIXME

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **created_at** | *date-time* | when struct was created | `"2015-01-01T12:00:00Z"` |
| **id** | *uuid* | unique identifier of struct | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **name** | *string* | unique name of struct | `"example"` |
| **updated_at** | *date-time* | when struct was updated | `"2015-01-01T12:00:00Z"` |

### Struct Create

Create a new struct.

```
POST /structs
```


#### Curl Example

```bash
$ curl -n -X POST https://api.airshooter.com/structs \
  -d '{
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Struct Delete

Delete an existing struct.

```
DELETE /structs/{struct_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X DELETE https://api.airshooter.com/structs/$STRUCT_ID_OR_NAME \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Struct Info

Info for existing struct.

```
GET /structs/{struct_id_or_name}
```


#### Curl Example

```bash
$ curl -n https://api.airshooter.com/structs/$STRUCT_ID_OR_NAME
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```

### Struct List

List existing structs.

```
GET /structs
```


#### Curl Example

```bash
$ curl -n https://api.airshooter.com/structs
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "created_at": "2015-01-01T12:00:00Z",
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "example",
    "updated_at": "2015-01-01T12:00:00Z"
  }
]
```

### Struct Update

Update an existing struct.

```
PATCH /structs/{struct_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X PATCH https://api.airshooter.com/structs/$STRUCT_ID_OR_NAME \
  -d '{
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2015-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "updated_at": "2015-01-01T12:00:00Z"
}
```


