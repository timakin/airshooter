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
POST /api/messages
GET /api/messages/:id
```
