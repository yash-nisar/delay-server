# App to test timeout scenarios

The primary purpose of this readme to build an app to test route timeout with envoy.

## Build and push image to docker registry

To build and push a mutliplatform image:

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t yourrepo.azurecr.io/external/delayserver:e2e \
  --push \
  .
```

## Testing

Make sure your endpoint has the format `/delay/x` where x is the time in seconds to delay the response.
