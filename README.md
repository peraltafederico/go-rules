---
runme:
  id: 01JF5XA4NFQWD4NY65V7VFSEZC
  version: v3
---

### Go Rules Example Projecy

This project uses this tutorial as a base: https://docs.gorules.io/docs/fintech-company-analysis

Install dependencies

```sh {"id":"01JF5XEKYN7206XNA53TKWV695"}
go mod tidy
```

Configure your envs

```sh {"id":"01JF5XV2VJPX8Q46TE0VGAX27M"}
export ACCESS_TOKEN="The access token for the API"
export PROJECT_ID="The project ID"
export DOMAIN="The domain"
export DOCUMENT_PATH="The document path of your decision"
```

Run in watch mode

```sh {"id":"01JF5XFS7P6PTSB9WYFYDB9ZR9"}
air
```