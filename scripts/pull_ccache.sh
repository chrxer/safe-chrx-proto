# only use for local testing
# shouldn't run as sudo

set -e
BUCKET_NAME=amzn-s3-chrxer-bucket-v1

TMP=.
CCACHE_DIR=$HOME/.cache/ccache/

if aws s3 ls "s3://$BUCKET_NAME/ccache.tar.gz"; then
        echo "Fetching ccache from S3..."
        aws s3 cp "s3://$BUCKET_NAME/ccache.tar.gz" "$TMP/ccache.tar.gz"

        echo "Extracting ccache..."
        mkdir -p $CCACHE_DIR
        tar -xzf "ccache.tar.gz" -C "$CCACHE_DIR" --strip-components=2
        rm ccache.tar.gz
fi

