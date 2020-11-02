# log function
log() {
    case $1 in
        "SUC")
            echo -ne "[\e[32mSUC\e[0m]"
            ;;
        "ERR")
            echo -ne "[\e[31mERR\e[0m]"
            ;;
        "INF")
            echo -ne "[\e[34mINF\e[0m]"
            ;;
        *)
            echo -ne "[$1]"
            ;;
    esac
    printf ": %s\n" "$2"
}