gonit
=====

*simple stand-alone web-api to provide the ability to create bare git repository on a server without admin access*

We're hosting our git repository on Windows servers, getting people admin access to create repositories seemed overkill and I wanted to play with Go anyway. It provides a simple RESTful API that returns all repositories/directories in the specified root directory for GET and calls `git init --bare` in response to a PUT. It's protected by a very simple "check the X-auth header for the correct word" mechanism.

Usage
-----

`gonit` expects a `config.json` in the working directory that looks like this:

    {
        "WebDir": "./web",
        "Port": ":9090",
        "RepositoryDir": "/tmp/repos"
    }

`WebDir`s content is provided at `/web` - I created a small [AngularJS](http://angularjs.org) app that talks to the API.
`Port` is .. well, the port.
`RepositoryDir` is your base git repository directory.

Just start the app and that's it! I use [winsw](https://github.com/kohsuke/winsw) to wrap `gonit` as a Windows service, it consumes about 6mb for me.

Build
-----

[Nothing special here](http://golang.org/doc/code.html), just `go install github.com/mathiasringhof/gonit' should do the trick.

TODO
----
* make `WebDir` optional
* add config option for the authentication password and make that optional as well