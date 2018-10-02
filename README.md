# hof

CLI tool for interacting with Hofstadter Studios.

_you will need to have a Hofstadter Studios Project_

### Install

__Stable__

Install the appropriate release from
the [Hof GitHub releases page](https://github.com/hofstadter-io/hof/releases).

For MacOS, you can also `brew install hofstadter-io/homebrew-tap/hof`.

__Unstable__

```
go get github.com/hofstadter-io/hof
```

### Setup

```
mkdir my-project
cd my-project
hof init my-project
```

Now put your username and apikey in the `hof.yaml`

### Pull Project

```
hof pull
```

You should see the files which were retrieved.

You can also visit https://my-app-name.live.hofstadter.io

### Edit Designs

You can now edit and create more designs.
Pages, custom UI components, and serverless functions
are coming soon.

### Push Project Updates

After changing your design you can
push to the sync server and update
your live instance.

```
hof push
```

The updates should reload your local browser
automatically, within ~10 seconds.
