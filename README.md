AirShooter
====

High-performance Notification and Messaging API


## API

### Notification

```
POST /api/notifications/enqueue
PUT /api/notifications/publish
GET /api/notifications/subscriptions
PUT /api/notifications/subscribe

cron PUT /api/notifications/expire
```

### Messaging

```
POST /api/messages (pass the combo of `from` and `to`)
GET /api/messages (pass the combo of `from` and `to`)
GET /api/messages/:id
```

### Toolkit

| Framework | Echo |
| --- | --- |
| DB | Redis(or MySQL(Q4M)) |
| Validation | JSV |
| Auth | JWTAuth |

