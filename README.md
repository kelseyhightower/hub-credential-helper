# hub-credential-helper

The `hub-credential-helper` implements the [git credentials API](https://git-scm.com/docs/git-credential) and provides usernames and passwords based on [hub](https://hub.github.com) configuration files.

## Install

## Usage

Generate a personal access [GitHub API token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line).

Create a hub configuration file and store it in the default configuration path:

```
cat <<EOF > ~/.config/hub
github.com:
  - protocol: https
    user: <github-username>
    oauth_token: <token>
EOF
``` 

> Set the HUB_CONFIG env var if you plan to store the hub configuration file in a different path.

Configure git to use the `hub-credential-helper` for the `https://github.com` URL:

```
git config --global credential.https://github.com.helper /usr/local/bin/hub-credential-helper
```

Test `hub-credential-helper`:

```
git credential fill
```

Stdin:
```
host=github.com
protocol=https

```

> Enter each line followed by a newline at the prompt. End the session with an addition newline.

Stdout:

```
protocol=https
host=github.com
username=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
password=x-oauth-basic
```
