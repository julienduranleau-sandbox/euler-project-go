inotifywait -e close_write -m . |
while read -r directory events filename; do
  if [ "$filename" = $1 ]; then
    clear
    go run $1
  fi
done

