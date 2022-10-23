# Real World base Golang For TiDB and TiDB Cloud


Thanks [xesina/golang-echo-realworld-example-app](https://github.com/xesina/golang-echo-realworld-example-app) project offer so great example.

This Project is created for Ti-Click project join PingCAP Hackathon 2022.

1. Ready the TiDB Cluster

You can create the TiDB Cluster on [TiDB Cloud](https://tidbcloud.com/) which offer free TiDB Cluster.

2. Access Gitpod
Access [Real World base Golang link](https://gitpod.io/#/https://github.com/ti-click/golang-tidb-realworld-example-app) and open the Gitpod online IDE.

Or you can install [Gitpod Chrome extension](https://chrome.google.com/webstore/detail/gitpod-always-ready-to-co/dodmmooeoklaejobgleioelladacbeki)  in advance, so you can just click the `Gitpod` button on the project's home page.

3. Experience Real World base Golang For TiDB

* Wait for all tasks to complete automatically.
  * <img width="500" alt="スクリーンショット 2022-10-23 14 09 58" src="https://user-images.githubusercontent.com/689799/197375028-a258ab31-0c1e-4b15-88bc-434b274a41e6.png">
  * <img width="500" alt="スクリーンショット 2022-10-23 14 09 43" src="https://user-images.githubusercontent.com/689799/197375030-129c1478-d16d-4eca-b308-6155ac7d13fc.png">
* Click the `frontend` link in the `PORTS` tab
  * <img width="500" alt="スクリーンショット 2022-10-23 14 10 14" src="https://user-images.githubusercontent.com/689799/197375041-14212d99-89fb-4ed1-a63c-fd8a846beb05.png">


4. Experience Real World base Golang For TiDB Cloud

* 4.1 Add Gitpod ip into TiDB Cloud access white list

  You can get the ip by the command in a Gitpod terminal below. Please add it into your TiDB Cloud security.

  > curl https://ipinfo.io/ip

* 4.2 You can restart the webserver 

  * 4.2.1 Click the `backend: go` terminal and switch to the tab
    * <img width="500" alt="スクリーンショット 2022-10-23 14 09 58" src="https://user-images.githubusercontent.com/689799/197375028-a258ab31-0c1e-4b15-88bc-434b274a41e6.png">
  * 4.2.2 Press ctrl+c to stop the golang backend server from running
  * rebuild data and start api server by commands below
  ```
  DB_USER=root DB_HOST={TiDBCloudHost} DB_PORT=4000 DB_NAME=test DB_PASS={TiDBCloudDBPassword} go run main.go
  ```

* 4.4 Click the `frontend` link in the `PORTS` tab , and experience the realworld site



## [xesina/golang-echo-realworld-example-app](https://github.com/xesina/golang-echo-realworld-example-app) README.md


# ![RealWorld Example App](logo.png)

> ### Golang/Echo codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.

### [Demo](https://github.com/gothinkster/realworld)&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)

[![Build Status](https://travis-ci.org/xesina/golang-echo-realworld-example-app.svg?branch=master)](https://travis-ci.org/xesina/golang-echo-realworld-example-app)

This codebase was created to demonstrate a fully fledged fullstack application built with **Golang/Echo** including CRUD operations, authentication, routing, pagination, and more.

## Getting started

### Install Golang (go1.11+)

Please check the official golang installation guide before you start. [Official Documentation](https://golang.org/doc/install)
Also make sure you have installed a go1.11+ version.

### Environment Config

make sure your ~/.*shrc have those variable:

```bash
➜  echo $GOPATH
/Users/xesina/go
➜  echo $GOROOT
/usr/local/go/
➜  echo $PATH
...:/usr/local/go/bin:/Users/xesina/test//bin:/usr/local/go/bin
```

For more info and detailed instructions please check this guide: [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Clone the repository

Clone this repository:

```bash
➜ git clone https://github.com/ti-click/golang-tidb-realworld-example-app.git
```

Or simply use the following command which will handle cloning the repo:

```bash
➜ go get -u -v github.com/ti-click/golang-tidb-realworld-example-app
```

Switch to the repo folder

```bash
➜ cd $GOPATH/src/github.com/ti-click/golang-tidb-realworld-example-app
```

### Install dependencies

```bash
➜ go mod download
```

### Run

```bash
➜ go run main.go
```

### Build

```bash
➜ go build
```

### Tests

```bash
➜ go test ./...
```
