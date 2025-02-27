# adsys
Active Directory bridging toolset

[![Code quality](https://github.com/ubuntu/adsys/workflows/QA/badge.svg)](https://github.com/ubuntu/adsys/actions?query=workflow%3AQA)
[![Code coverage](https://codecov.io/gh/ubuntu/adsys/branch/master/graph/badge.svg)](https://codecov.io/gh/ubuntu/adsys)
[![Go Reference](https://pkg.go.dev/badge/github.com/ubuntu/adsys.svg)](https://pkg.go.dev/github.com/ubuntu/adsys)
[![Go Report Card](https://goreportcard.com/badge/ubuntu/adsys)](https://goreportcard.com/report/ubuntu/adsys)
[![License](https://img.shields.io/badge/License-GPL3.0-blue.svg)](https://github.com/ubuntu/adsys/blob/master/LICENSE)

## Usage

### User commands

#### adsysctl

AD integration client

##### Synopsis

Active Directory integration bridging toolset command line tool.

```
adsysctl COMMAND [flags]
```

##### Options

```
  -c, --config string   use a specific configuration file
  -h, --help            help for adsysctl
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl applied

Print last applied GPOs for current or given user/machine

##### Synopsis

Alias of "policy applied"

```
adsysctl applied [USER_NAME] [flags]
```

##### Options

```
  -a, --all        show overridden rules in each GPOs.
      --details    show applied rules in addition to GPOs.
  -h, --help       help for applied
  -m, --machine    show applied rules to the machine.
      --no-color   don't display colorized version.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl completion

Generate the autocompletion script for the specified shell

##### Synopsis

Generate the autocompletion script for adsysctl for the specified shell.
See each sub-command's help for details on how to use the generated script.


##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl completion bash

Generate the autocompletion script for bash

##### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(adsysctl completion bash)

To load completions for every new session, execute once:

###### Linux:

	adsysctl completion bash > /etc/bash_completion.d/adsysctl

###### macOS:

	adsysctl completion bash > $(brew --prefix)/etc/bash_completion.d/adsysctl

You will need to start a new shell for this setup to take effect.


```
adsysctl completion bash
```

##### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl completion fish

Generate the autocompletion script for fish

##### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	adsysctl completion fish | source

To load completions for every new session, execute once:

	adsysctl completion fish > ~/.config/fish/completions/adsysctl.fish

You will need to start a new shell for this setup to take effect.


```
adsysctl completion fish [flags]
```

##### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl completion powershell

Generate the autocompletion script for powershell

##### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	adsysctl completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
adsysctl completion powershell [flags]
```

##### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl completion zsh

Generate the autocompletion script for zsh

##### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(adsysctl completion zsh); compdef _adsysctl adsysctl

To load completions for every new session, execute once:

###### Linux:

	adsysctl completion zsh > "${fpath[1]}/_adsysctl"

###### macOS:

	adsysctl completion zsh > $(brew --prefix)/share/zsh/site-functions/_adsysctl

You will need to start a new shell for this setup to take effect.


```
adsysctl completion zsh [flags]
```

##### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl doc

Documentation

```
adsysctl doc [CHAPTER] [flags]
```

##### Options

```
  -d, --dest string     Write documentation file(s) to this directory.
  -f, --format string   Format type (markdown, raw or html). (default "markdown")
  -h, --help            help for doc
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy

Policy management

```
adsysctl policy COMMAND [flags]
```

##### Options

```
  -h, --help   help for policy
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy admx

Dump windows policy definitions

```
adsysctl policy admx lts-only|all [flags]
```

##### Options

```
      --distro string   distro for which to retrieve policy definition. (default "Ubuntu")
  -h, --help            help for admx
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy applied

Print last applied GPOs for current or given user/machine

```
adsysctl policy applied [USER_NAME] [flags]
```

##### Options

```
  -a, --all        show overridden rules in each GPOs.
      --details    show applied rules in addition to GPOs.
  -h, --help       help for applied
  -m, --machine    show applied rules to the machine.
      --no-color   don't display colorized version.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy update

Updates/Create a policy for current user or given user with its kerberos ticket

```
adsysctl policy update [USER_NAME KERBEROS_TICKET_PATH] [flags]
```

##### Options

```
  -a, --all       all updates the policy of the computer and all the logged in users. -m or USER_NAME/TICKET cannot be used with this option.
  -h, --help      help for update
  -m, --machine   machine updates the policy of the computer.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service

Service management

```
adsysctl service COMMAND [flags]
```

##### Options

```
  -h, --help   help for service
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service cat

Print service logs

```
adsysctl service cat [flags]
```

##### Options

```
  -h, --help   help for cat
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service status

Print service status

```
adsysctl service status [flags]
```

##### Options

```
  -h, --help   help for status
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service stop

Requests to stop the service once all connections are done

```
adsysctl service stop [flags]
```

##### Options

```
  -f, --force   force will shut it down immediately and drop existing connections.
  -h, --help    help for stop
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl update

Updates/Create a policy for current user or given user with its kerberos ticket

##### Synopsis

Alias of "policy update"

```
adsysctl update [USER_NAME KERBEROS_TICKET_PATH] [flags]
```

##### Options

```
  -a, --all       all updates the policy of the computer and all the logged in users. -m or USER_NAME/TICKET cannot be used with this option.
  -h, --help      help for update
  -m, --machine   machine updates the policy of the computer.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl version

Returns version of client and service

```
adsysctl version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd

AD integration daemon

##### Synopsis

Active Directory integration bridging toolset daemon.

```
adsysd COMMAND [flags]
```

##### Options

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
  -h, --help                    help for adsysd
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd completion

Generate the autocompletion script for the specified shell

##### Synopsis

Generate the autocompletion script for adsysd for the specified shell.
See each sub-command's help for details on how to use the generated script.


##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd completion bash

Generate the autocompletion script for bash

##### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(adsysd completion bash)

To load completions for every new session, execute once:

###### Linux:

	adsysd completion bash > /etc/bash_completion.d/adsysd

###### macOS:

	adsysd completion bash > $(brew --prefix)/etc/bash_completion.d/adsysd

You will need to start a new shell for this setup to take effect.


```
adsysd completion bash
```

##### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd completion fish

Generate the autocompletion script for fish

##### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	adsysd completion fish | source

To load completions for every new session, execute once:

	adsysd completion fish > ~/.config/fish/completions/adsysd.fish

You will need to start a new shell for this setup to take effect.


```
adsysd completion fish [flags]
```

##### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd completion powershell

Generate the autocompletion script for powershell

##### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	adsysd completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
adsysd completion powershell [flags]
```

##### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd completion zsh

Generate the autocompletion script for zsh

##### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(adsysd completion zsh); compdef _adsysd adsysd

To load completions for every new session, execute once:

###### Linux:

	adsysd completion zsh > "${fpath[1]}/_adsysd"

###### macOS:

	adsysd completion zsh > $(brew --prefix)/share/zsh/site-functions/_adsysd

You will need to start a new shell for this setup to take effect.


```
adsysd completion zsh [flags]
```

##### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd version

Returns version of service and exits

```
adsysd version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd

AD watch daemon

##### Synopsis

Watch directories for changes and bump the relevant GPT.ini versions.

```
adwatchd [COMMAND] [flags]
```

##### Options

```
  -c, --config string   use a specific configuration file
  -h, --help            help for adwatchd
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd completion

Generate the autocompletion script for the specified shell

##### Synopsis

Generate the autocompletion script for adwatchd for the specified shell.
See each sub-command's help for details on how to use the generated script.


##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd completion bash

Generate the autocompletion script for bash

##### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(adwatchd completion bash)

To load completions for every new session, execute once:

###### Linux:

	adwatchd completion bash > /etc/bash_completion.d/adwatchd

###### macOS:

	adwatchd completion bash > $(brew --prefix)/etc/bash_completion.d/adwatchd

You will need to start a new shell for this setup to take effect.


```
adwatchd completion bash
```

##### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd completion fish

Generate the autocompletion script for fish

##### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	adwatchd completion fish | source

To load completions for every new session, execute once:

	adwatchd completion fish > ~/.config/fish/completions/adwatchd.fish

You will need to start a new shell for this setup to take effect.


```
adwatchd completion fish [flags]
```

##### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd completion powershell

Generate the autocompletion script for powershell

##### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	adwatchd completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
adwatchd completion powershell [flags]
```

##### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd completion zsh

Generate the autocompletion script for zsh

##### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(adwatchd completion zsh); compdef _adwatchd adwatchd

To load completions for every new session, execute once:

###### Linux:

	adwatchd completion zsh > "${fpath[1]}/_adwatchd"

###### macOS:

	adwatchd completion zsh > $(brew --prefix)/share/zsh/site-functions/_adwatchd

You will need to start a new shell for this setup to take effect.


```
adwatchd completion zsh [flags]
```

##### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd run

Starts the directory watch loop

##### Synopsis

Can run as a service through the service manager or interactively as a standalone application.

The program will monitor the configured directories for changes and bump the appropriate GPT.ini versions anytime a change is detected.
If a GPT.ini file does not exist for a directory, a warning will be issued and the file will be created. If the GPT.ini file is incompatible or malformed, the program will report an error.


```
adwatchd run [flags]
```

##### Options

```
  -c, --config string    use a specific configuration file
  -d, --dirs directory   a directory to check for changes (can be specified multiple times)
  -f, --force            force the program to run even if another instance is already running
  -h, --help             help for run
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service

Manages the adwatchd service

##### Synopsis

The service command allows the user to interact with the adwatchd service. It can manage and query the service status, and also install and uninstall the service.

```
adwatchd service COMMAND [flags]
```

##### Options

```
  -h, --help   help for service
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service install

Installs the service

##### Synopsis

Installs the adwatchd service.

The service will be installed as a Windows service.


```
adwatchd service install [flags]
```

##### Options

```
  -c, --config string   use a specific configuration file
  -h, --help            help for install
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service restart

Restarts the service

##### Synopsis

Restarts the adwatchd service.

```
adwatchd service restart [flags]
```

##### Options

```
  -h, --help   help for restart
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service start

Starts the service

##### Synopsis

Starts the adwatchd service.

```
adwatchd service start [flags]
```

##### Options

```
  -h, --help   help for start
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service status

Returns service status

##### Synopsis

Returns the status of the adwatchd service.

```
adwatchd service status [flags]
```

##### Options

```
  -h, --help   help for status
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service stop

Stops the service

##### Synopsis

Stops the adwatchd service.

```
adwatchd service stop [flags]
```

##### Options

```
  -h, --help   help for stop
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd service uninstall

Uninstalls the service

##### Synopsis

Uninstalls the adwatchd service.

```
adwatchd service uninstall [flags]
```

##### Options

```
  -h, --help   help for uninstall
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adwatchd version

Returns version of service and exits

```
adwatchd version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

### Hidden commands

Those commands are hidden from help and should primarily be used by the system or for debugging.

#### adsysctl policy debug

Debug various policy infos

```
adsysctl policy debug [flags]
```

##### Options

```
  -h, --help   help for debug
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy debug gpolist-script

Write GPO list python embeeded script in current directory

```
adsysctl policy debug gpolist-script [flags]
```

##### Options

```
  -h, --help   help for gpolist-script
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd runscripts

Runs scripts in the given subdirectory

```
adsysd runscripts ORDER_FILE [flags]
```

##### Options

```
      --allow-order-missing   allow ORDER_FILE to be missing once the scripts are ready.
  -h, --help                  help for runscripts
```

##### Options inherited from parent commands

```
      --ad-backend string       Active Directory authentication backend (default "sssd")
      --cache-dir string        directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string           use a specific configuration file
      --run-dir string          directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string           socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
      --sssd.cache-dir string   SSSd cache directory (default "/var/lib/sss/db")
      --sssd.config string      SSSd config file path (default "/etc/sssd/sssd.conf")
  -t, --timeout int             time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count           issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

## Installing development versions

For every commit on the `main` branch of the `adsys` repository, the GitHub Actions CI builds a development version of the `adwatchd` project. This is *NOT* a stable version of the application and should not be used for production purposes. However, it may prove useful to preview features or bugfixes not yet available as part of a stable release.

To get access to the build artifact you need to be logged in on GitHub. Then, click on any passing run of the [QA workflow](https://github.com/ubuntu/adsys/actions/workflows/qa.yaml) that has the `Windows tests for adwatchd` job, and look for the `adwatchd_setup` file.
