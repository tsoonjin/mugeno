# GraphQL

## Intro

- syntax to ask for data from server
- like a middleware, think smart endpoint vs dumb endpoint (REST)

## What is it GREAT for

- built to replace complicated recursive SQL join. solve N+1 problem
- solve under and over fetching of data. get exactly what you want
- aggregates data from different sources
- tailor to needs of different client instead of creating dedicated endpoint. Mobile don't need certain data compared to web
- no more API versioning
- self documenting


## What NOT to use for

- added complexity
- causes some overhead
- no support for HTTP caching
- security. crafting slow query can ddos endpoint
- lack of authentication flow

## Some issues

- bottleneck endpoint (not true since nothing to do with number of endpoints)
- monitoring
