set -e
WRK=$(realpath $(dirname $(dirname "$0")))
CHROMIUM="$WRK/chromium"
export PATH="$WRK/depot_tools:$PATH"

cd "$CHROMIUM/src"
COUNT=$(git ls-files | wc -l)
SIZE=$(du -s -b "$CHROMIUM/src" | awk '{print $1}')
AVGSIZE=$(echo "$SIZE/$COUNT" | bc)

cd $WRK

AVGSIZER=$(numfmt --to=si --suffix=B --format="%9.2f" "$AVGSIZE")
SIZER=$(numfmt --to=iec --suffix=B --format="%9.2f" "$SIZE")
COUNTR=$(numfmt --to=iec --format="%9.2f" "$COUNT")

echo "Chromium's stats"
echo "Files: $COUNTR"
echo "Size: $SIZER"
echo "Average size: $AVGSIZER"

set +e