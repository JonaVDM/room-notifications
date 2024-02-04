# Room.nl notifications

Get a notification when a new room is posted on Room.nl. Makes use of a service
called `Pushover` to send the notifications.

## Running it

```bash
docker run -e PUSH_TOKEN [pushover_app_token] -e PUSH_USER [pushover_user/group_token] ghcr.io/jonavdm/room-notifications:latest
```
