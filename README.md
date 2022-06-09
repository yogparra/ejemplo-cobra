
# ejemplo-grpc
    go mod init github.com/yogparra/ejemplo-cobra

# instalacion
    go get -u github.com/spf13/cobra@latest

 # ejemplo
    mkdir ejemplo-cobra
    cd ejemplo-cobra
    go mod init github.com/yogparra/ejemplo-cobra

 # inicializando el proyecto
    ejemplo-cobra/cobra-cli init

# add comandos al proyecto
    cobra-cli add info
    cobra-cli add scan
    cobra-cli add version
    cobra-cli add list


# ejecutable
    go install

# tip
    APPNAME Command Args --Flags or APPNAME Command --Flags Args (--flags, -f)

# ejemplo detallado de documentar
    hugo is the main command, used to build your Hugo site.

    Hugo is a Fast and Flexible Static Site Generator
    built with love by spf13 and friends in Go.

    Complete documentation is available at http://gohugo.io/.

    Usage:
    hugo [flags]
    hugo [command]

    Available Commands:
    config      Print the site configuration
    convert     Convert your content to different formats
    deploy      Deploy your site to a Cloud provider.
    env         Print Hugo version and environment info
    gen         A collection of several useful generators.
    help        Help about any command
    import      Import your site from others.
    list        Listing out various types of content
    mod         Various Hugo Modules helpers.
    new         Create new content for your site
    server      A high performance webserver
    version     Print the version number of Hugo

    Flags:
    -b, --baseURL string         hostname (and path) to the root, e.g. http://spf13.com/
    -D, --buildDrafts            include content marked as draft
    -E, --buildExpired           include expired content
    -t, --theme strings          themes to use (located in /themes/THEMENAME/)
        --themesDir string       filesystem path to themes directory
        --trace file             write trace to file (not useful in general)
    -v, --verbose                verbose output
        --verboseLog             verbose logging
    -w, --watch                  watch filesystem for changes and recreate as needed

    Additional help topics:
    hugo check   Contains some verification checks

    Use "hugo [command] --help" for more information about a command.

# docu
    [ejemplo-cobra] es el comando principal para escanear rutas
    
    [ejemplo-cobra] es un escaneador rapido y flexible contruido con cariño por guillermo

    La documentacion completa esta disponible en github.com/yogparra/ejemplo-cobra

    Usage:
    ejemplo-cobra [command]

    Available Commands:
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    info        A brief description of your command

    Flags:
    -h, --help   help for ejemplo-cobra

    Use "ejemplo-cobra [command] --help" for more information about a command.1

# rutas1
    [ejemplo-cobra] version     Print the version number of ejemplo-cobra
    v0.00.1

# rutas2
    [ejemplo-cobra] list        List the paths to scan