AirShooter
====

High-performance Notification and Messaging API

## Preparation

```
# install depended packages
$ godep go install

# run with specified packages
$ godep go run server.go
```


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
| DB | MySQL |
| Validation | JSV |
| Auth | JWTAuth |

