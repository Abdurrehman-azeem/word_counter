if [[ "$#" -ne 2 ]]; then
    echo "Must provide flag, with filename."
fi

fl="$1"
file="$2"

if ! test -f "$file"; then
    echo "File $file, does not exist."
    echo "Exiting..."
    exit 1
fi

chmod +x main.go
go run main.go $fl $file; exit "$?"