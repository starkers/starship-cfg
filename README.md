

Starship is great but it doesn't allow you its configuration to come from more than one file.
This is a problem if there are certain things in your configuration that you would not like to be public.

There does not seem to have an appetite to adress this:
- https://github.com/starship/starship/issues/5341

This tool will do this before starship loads. It works by:



1. reading all `*.toml` files from `${XDG_CONFIG_HOME}/starship.d/`  (typically `~/.config/starship.d`)
2. merging them
3. writing the result to `${XDG_CONFIG_HOME}/starship.toml`




# Installation

## via go
`go install github.com/starkers/starship-cfg`

## download a release

grab a release binary that is compatible with your platform and manually put it into your $PATH




# Usage

## fish


TLDR: before running  `starship init fish | source` , run `starship-cfg`

```fish
if status is-interactive

    if command -s starship-cfg > /dev/null
        starship-cfg
    else
        echo '#WARN: starship-cfg not installed'
    end

    if command -s starship > /dev/null
        starship init fish | source
    else
        echo '#WARN: starship not installed'
    end
end

```

Don't forget to move your config(s) into `~/.config/starship.d/` !



# TODO

- [ ] support native `$STARSHIP_CONFIG` [configuration](https://starship.rs/config/#config-file-location)
- [ ] debug mode?
