# Changelog

## [1.0.2](https://github.com/lichb0rn/go-social-n/compare/v1.0.1...v1.0.2) (2025-05-22)


### Bug Fixes

* update workflow to fix version extraction and replacement ([d9a2527](https://github.com/lichb0rn/go-social-n/commit/d9a252795a6214795bc5da366af940c3635c305f))

## [1.0.1](https://github.com/lichb0rn/go-social-n/compare/v1.0.0...v1.0.1) (2025-05-22)


### Bug Fixes

* change deprecated set-output command in GitHub Actions ([afcee7c](https://github.com/lichb0rn/go-social-n/commit/afcee7c3de6d9ca5f7f6e00b2bb84a94287c31bc))

## 1.0.0 (2025-05-22)


### Features

* add basic auth ([76342fa](https://github.com/lichb0rn/go-social-n/commit/76342fa4eff3334b7cb4ad2b7d64859296ec4541))
* add followers and user feed ([116ca32](https://github.com/lichb0rn/go-social-n/commit/116ca326c1c9e367f59a5c6018d20c126e9ae591))
* add pagination ([c5352dd](https://github.com/lichb0rn/go-social-n/commit/c5352dd532c7812f2e3eb19fa8aca6d87b83c497))
* add rate limiter middleware and fixed window implementation ([b3e2025](https://github.com/lichb0rn/go-social-n/commit/b3e2025e18a948a7c9133b9dcf916ab543a855a4))
* add structured loggin ([8559806](https://github.com/lichb0rn/go-social-n/commit/855980604fe64cafff06910a4484115e910d0d78))
* add test utilities and mock implementations for user and cache stores ([bffe760](https://github.com/lichb0rn/go-social-n/commit/bffe760ab3db1a244ad83b60ed415d53e7342f77))
* add user invitation and activation ([7c9773b](https://github.com/lichb0rn/go-social-n/commit/7c9773b1c2a9e04d9330dbbcd685df0400678244))
* add user invitation email ([ae81ab7](https://github.com/lichb0rn/go-social-n/commit/ae81ab7e68c91470b7c32cc2d3b44eb900078e0e))
* implement basic server metrics ([1c062c2](https://github.com/lichb0rn/go-social-n/commit/1c062c2cd5ea9fae2a9a5a5af421a0f0f39ed19a))
* implement graceful shutdown for HTTP server ([0f9cc6a](https://github.com/lichb0rn/go-social-n/commit/0f9cc6ad597c367677d457bcf54d5e6a0effc9b5))
* implement JWT authentication and token generation ([5c9226a](https://github.com/lichb0rn/go-social-n/commit/5c9226a0c67941c06624bbeb1fba7cc442ee3b49))
* implement Redis caching for user data and update user retrieval logic ([bd48af9](https://github.com/lichb0rn/go-social-n/commit/bd48af9a6faa88613f712e39325c0143b70d7c5a))
* implement role-based access control and user role management ([3481691](https://github.com/lichb0rn/go-social-n/commit/348169108a8aa6e635ee77a55e72e88956d18a20))
* initialize web application with React, Vite, and TypeScript ([996856f](https://github.com/lichb0rn/go-social-n/commit/996856fb6f393e22f386e2d1f44c28cd16ba6fb6))
* project setup ([d05f551](https://github.com/lichb0rn/go-social-n/commit/d05f551c0554f3005ecf2d62d3da12450c547aef))


### Bug Fixes

* optimistic concurency problem when update post ([6652224](https://github.com/lichb0rn/go-social-n/commit/66522244c850350782915f6821d60d64b21abbe0))
* remove redundant RoleID field from User struct ([f5b0935](https://github.com/lichb0rn/go-social-n/commit/f5b09358a23091eb137cf9b1098d396934303231))
