# Changelog

## [1.2.0](https://github.com/0xshiku/sogo/compare/v1.1.0...v1.2.0) (2024-11-15)


### Features

* adds dockerfile ([5f0850a](https://github.com/0xshiku/sogo/commit/5f0850a22310b417f35cc4bfa94b8aa9a0d7c1b7))

## [1.1.0](https://github.com/0xshiku/sogo/compare/v1.0.0...v1.1.0) (2024-11-15)


### Features

* adds script to update api version on main.go ([5767756](https://github.com/0xshiku/sogo/commit/57677561e461f99654645cae84b4e4ee690ea3c4))

## 1.0.0 (2024-11-15)


### Features

* Add User and Post data model with basic fetch and post query ([4530b07](https://github.com/0xshiku/sogo/commit/4530b072e5f55b1e794f1ab275b93c059be9bd8a))
* Adds comments to posts on the storage and api ([8baf2c8](https://github.com/0xshiku/sogo/commit/8baf2c8b1ba79fc144b594ac38c555d2bd323240))
* adds followers to database and endpoint ([1965ec1](https://github.com/0xshiku/sogo/commit/1965ec11d137d0f178498c916511d61cb14c95b2))
* adds release please script to workflows ([d3611b1](https://github.com/0xshiku/sogo/commit/d3611b1de096e849a231b902314ca88fe24b9612))
* authorization database setup ([f9ab9fd](https://github.com/0xshiku/sogo/commit/f9ab9fd66d51bd4b98dddf857ea1d2f3b65363e7))
* bootstrap frontend, create confirmation page UI, and bring basic cors to backend. ([c2f8e88](https://github.com/0xshiku/sogo/commit/c2f8e882848bb9d30212575c5f205a5153372bd2))
* caches with redis the user profile ([adc857f](https://github.com/0xshiku/sogo/commit/adc857fd10706620add5a4c30cab2f58b817cb3b))
* create basic authentication ([6b371d2](https://github.com/0xshiku/sogo/commit/6b371d2e409121f4bb39c89e46cec94e0f394b7d))
* Create filtering, sorting and pagination ([b7030c7](https://github.com/0xshiku/sogo/commit/b7030c7341f99cd6e5df1365f2d915665896f718))
* Create User feed algorithm ([1a95d85](https://github.com/0xshiku/sogo/commit/1a95d8505bf06131cd30484112caa17ee75604aa))
* creates the role precedence middleware ([74a7416](https://github.com/0xshiku/sogo/commit/74a741626a3d076ef36bf1bbb70b996f1b81e7f1))
* Creates user activation flow ([382f5fd](https://github.com/0xshiku/sogo/commit/382f5fdc29ddc096ffe20fba03eae4271011064a))
* Creates user registration foundations and sql transactions ([24e4775](https://github.com/0xshiku/sogo/commit/24e477523489c9055d7255df06a2936428500449))
* finishes database setup and runs first migrations ([5c53943](https://github.com/0xshiku/sogo/commit/5c539438ad5f60ab7c88cc1a04700c72de3d785f))
* gets a user by id by creating endpoint and sql query ([78b00d6](https://github.com/0xshiku/sogo/commit/78b00d64163e38d34168d6f07060ecb2e1315019))
* gracefully shutdown server ([5aa5e61](https://github.com/0xshiku/sogo/commit/5aa5e615dda243427392339dec3fbb0197cc7af3))
* Implement rate limiter with fixed window algorithm to the API ([f4be372](https://github.com/0xshiku/sogo/commit/f4be3728e2befde1af39cb9d0b6125e233b133d6))
* introduce basic metrics on the project ([0204c27](https://github.com/0xshiku/sogo/commit/0204c27a68f7fb49f27584d8111719a71f15bdb0))
* introduce github actions ([57c0a43](https://github.com/0xshiku/sogo/commit/57c0a4307caa0788a134b27acb0a0a092f257666))
* invalidate post cache ([847f82c](https://github.com/0xshiku/sogo/commit/847f82cf3ccb65ef125a996fb5ff26213bade6bc))
* make token generation ([9a0200a](https://github.com/0xshiku/sogo/commit/9a0200a574272411c70e2786a6fb49dad8b51244))
* marshalling JSON responses and creates and fetch users feed post ([6d7857a](https://github.com/0xshiku/sogo/commit/6d7857accf787394c2062b0872e9d47efc7c45d8))
* Sending the invitation email ([34929b3](https://github.com/0xshiku/sogo/commit/34929b38979cf9afac97a66505cec34421305042))
* Updating and deleting posts, and as well creating the middleware to get posts from context. ([e56b398](https://github.com/0xshiku/sogo/commit/e56b39818890b58c74437551bfd7bce4253c70d1))
* validate tokens ([85fdd52](https://github.com/0xshiku/sogo/commit/85fdd52fef17bdf78f4a2c6bf2a3a104ee710c3b))


### Bug Fixes

* comment conflict response error function. it is unused. ([be1aefc](https://github.com/0xshiku/sogo/commit/be1aefc553de6d9ac7a69ada0a61c8ed392a9331))
* fix CORS definitions on API ([b3cde4d](https://github.com/0xshiku/sogo/commit/b3cde4da82255b7f0788a670437114c08e29b207))
* fix the user invitation issue with the role insertion. ([311314f](https://github.com/0xshiku/sogo/commit/311314f9bff421dd1bad18e4f27d6c2edc03c95e))
* fixes unused function on users.go and passing MockUserStore without a pointer on mock of cache ([64860f7](https://github.com/0xshiku/sogo/commit/64860f793075a0e5bf573da213baf24edee711c2))
* missing hash password comparison on the token creation handler ([a1a3851](https://github.com/0xshiku/sogo/commit/a1a38517009ad2bfe9e3382b6f9cc2199a080791))
* remove log from env and test_concurrency script ([9cf32b1](https://github.com/0xshiku/sogo/commit/9cf32b1aff974cbe7d6090a491e6632f896a03f6))
