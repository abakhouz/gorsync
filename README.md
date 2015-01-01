# gorsync

A small helper utility that takes away the need to remember exact rsync
commands to sync certain directories.

I have a lot of rsync-ed directories on my machine and need to get the right
command from my bash history every time. Instead this tool allows me to
add a simple configuration file `gorsync.yml|json|toml` to the parent directory
and run `gorsync`.

It's also just an experimental tool I wrote learning [Go](http://golang.org/).
So please don't use it in a production environment!

## Install

Download the latest version
[here](https://github.com/tombruijn/gorsync/releases).

Put the binary in directory that's in your $PATH.

See if it is availble by running `which gorsync`.

## Setup

Add a file to the directory you would normally execute the rsync command.
Name it `gorsync.yml|json|toml` and add the following config (in whatever
format you choose).

- options
    - an array of options to give to the rsync command.

      `"options": ["-a", "-v", "--delete", "--exclude=some/dir"]`

- directories
    - nested key value object with a `from` and `to` key.
    - `to` should be relative to the path your config file is in.

      ```json
      "directories": {
        "from": "/my/remote/dir",
        "to": "my_local_dir_name"
      }
      ```

### Example config file

```json
{
  "options": ["-a", "-v", "--delete"],
  "directories": {
    "from": "/my/remote/dir",
    "to": "my_local_dir_name"
  }
}
```

## Run

Just run: `gorsync` in a directory with a `gorsync.yml|json|tom` config file.

It's that easy!

## License

gorsync released under the MIT License. See the bundled LICENSE file for
details.
