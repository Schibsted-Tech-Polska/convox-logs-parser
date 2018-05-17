# Convox Logs Parser

This tool uses `convox logs` command and transforms its output if the known application log format is discovered

The initial idea was to use it for applications which are sending out on standard output a JSON formatted logs.
This format is useful when using some log collector later on (like Datadog Logs, ELK stack) so you will not end up with 
multiple log entries when Stack Trace is thrown. Thanks to using JSON log format we can parse all logs more efficiently,
but we still need a way to read them on console in human-friendly form.

## Supported Formats

### Standard log output
This format is eventually displaying the logs as they came from `convox logs`. This is a default behaviour for data
which is not recognized by this application

### Log4j2 JSON layout
If output of the application is in JSON format and contains specific pre-defined fields, then `clp` (Convox Log Parser)
displays those lines like normally Log4j2 would do it without JSON layout.

Currently only pre-defined fields are supported, though it is planned to allow user set the output format. 

## References
[convox/rack](https://github.com/convox/rack) - as this application is used as a log parser, the log source is worth to be mentioned.

## How to install?

You should go to ***releases*** tab, select latest version and at ***Assets*** part - download proper version for your OS, 

If you want to download latest version for your architecture use following command (where ARCH can be one of `linux`, `darwin` or `windows`).
```
ARCH="linux"; curl -L $(curl -s https://api.github.com/repos/Schibsted-Tech-Polska/convox-logs-parser/releases/latest | jq -r ".assets[] | select(.name == \"clp-$ARCH\") | .browser_download_url") -o /usr/local/bin/clp; chmod +x /usr/local/bin/clp
```

Above command downloads `convox-logs-parser` to file `/usr/local/bin/clp` and adds to it execute permissions.
Requires commands to exist: `curl`, `jq`, `chmod`.
