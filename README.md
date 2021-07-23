# dockerhub-image-keeper

イメージを指定するとDockerHubから1ヶ月以上pullされていないタグをpullして自動的にrmします

```
 ❯ go run . -image plusminusio/mastodon
pull targets:
  3.4.0
  3.3.0-arm64
  3.3-arm64
  3.3.0
  3.3

pull plusminusio/mastodon:3.4.0 ( 1 / 5 )
  3.4.0: Pulling from plusminusio/mastodon
  345e3491a907: Pulling fs layer
  57671312ef6f: Pulling fs layer
  5e9250ddb7d0: Pulling fs layer
  a930d0d87cfa: Pulling fs layer
  b2d83542d209: Pulling fs layer
  74e1c5f2ad45: Pulling fs layer
  50063ff985ae: Pulling fs layer
  eb7a00eb5951: Pulling fs layer
  90cba382b0b4: Pulling fs layer
  d157dbad3cd9: Pulling fs layer
  74e1c5f2ad45: Waiting
  50063ff985ae: Waiting
  a930d0d87cfa: Waiting
  eb7a00eb5951: Waiting
  b2d83542d209: Waiting
  90cba382b0b4: Waiting
  d157dbad3cd9: Waiting
  5e9250ddb7d0: Download complete
  57671312ef6f: Verifying Checksum
  57671312ef6f: Download complete
  345e3491a907: Verifying Checksum
  345e3491a907: Download complete
  345e3491a907: Pull complete
  57671312ef6f: Pull complete
  5e9250ddb7d0: Pull complete
  a930d0d87cfa: Verifying Checksum
  a930d0d87cfa: Download complete
  74e1c5f2ad45: Download complete
  a930d0d87cfa: Pull complete
  b2d83542d209: Verifying Checksum
  b2d83542d209: Download complete
  b2d83542d209: Pull complete
  74e1c5f2ad45: Pull complete
  eb7a00eb5951: Verifying Checksum
  eb7a00eb5951: Download complete
  50063ff985ae: Verifying Checksum
  50063ff985ae: Download complete
  90cba382b0b4: Verifying Checksum
  90cba382b0b4: Download complete
  d157dbad3cd9: Verifying Checksum
  d157dbad3cd9: Download complete
  50063ff985ae: Pull complete
  eb7a00eb5951: Pull complete
  90cba382b0b4: Pull complete
  d157dbad3cd9: Pull complete
  Digest: sha256:bf8291fb8832f03e90b53aa83d41e559a9e61584c1102aba021258bb40019134
  Status: Downloaded newer image for plusminusio/mastodon:3.4.0
  docker.io/plusminusio/mastodon:3.4.0

rm plusminusio/mastodon:3.4.0 ( 1 / 5 )
  Untagged: plusminusio/mastodon:3.4.0
  Untagged: plusminusio/mastodon@sha256:bf8291fb8832f03e90b53aa83d41e559a9e61584c1102aba021258bb40019134
  Deleted: sha256:a5e280581b8261ff9fc8e104bf9c9469530238b8ec28264c3fc1078bac7d9c3d
  Deleted: sha256:4032ccbea18bbd49cfd7df6ece4ab48562ac23a26db98529c29755946f3ec01f
  Deleted: sha256:6f09048131bd2100c76c83fd402613948b8346bf74f5892c534748d08c1d1560
  Deleted: sha256:03bf964d61a1026ffac322aa615bac7e84b7bcb55c7626670fdea69ecdb831ce
  Deleted: sha256:6175061ec9abeb3966fb6ff9708def8c965fcea1037c37bf1cbc0b304e1a60e1
  Deleted: sha256:af03e60d797562d2bcfb3305d1653c2595feb3d776f1b3a06560cbca4b0869a8
  Deleted: sha256:3d86b9dfb4a112ff0d7579eab1c6cfcc29dc7ffd2edfa5244195f6ea430f6349
  Deleted: sha256:39b006d4590b696ad37a872f736a337d53285d13916d2c9f44b8f1e5fec38946
  Deleted: sha256:3dd8c8d4fd5b59d543c8f75a67cdfaab30aef5a6d99aea3fe74d8cc69d4e7bf2
  Deleted: sha256:8d8dceacec7085abcab1f93ac1128765bc6cf0caac334c821e01546bd96eb741
  Deleted: sha256:ccdbb80308cc5ef43b605ac28fac29c6a597f89f5a169bbedbb8dec29c987439

pull plusminusio/mastodon:3.3.0-arm64 ( 2 / 5 )

  3.3.0-arm64: Pulling from plusminusio/mastodon
  a970164f39c1: Pulling fs layer
  e9c66f1fb5a2: Pulling fs layer
  94362ba2c285: Pulling fs layer

  ...
```
