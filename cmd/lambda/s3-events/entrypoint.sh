#!/bin/sh
set -e ; # Have script exit in the event of a failed command.

# checkBucketExists ($bucket)
# Check if the bucket exists, by using the exit code of `mc ls`
checkBucketExists() {
  BUCKET=$1
  CMD=$(/usr/bin/mc events list local/$BUCKET > /dev/null 2>&1)
  return $?
}

# createBucket ($bucket, $policy)
# Ensure bucket exists, purging if asked to
createBucket() {
  BUCKET=$1
  POLICY=$2

  # Purge the bucket, if set & exists
  # Since PURGE is user input, check explicitly for `true`
  if [ $PURGE_BUCKETS = true ]; then
    if checkBucketExists $BUCKET ; then
      echo "Purging bucket '$BUCKET'."
      set +e ; # don't exit if this fails
      /usr/bin/mc rm -r --force local/$BUCKET
      set -e ; # reset `e` as active
    else
      echo "Bucket '$BUCKET' does not exist, skipping purge."
    fi
  fi

  # Create the bucket if it does not exist
  if ! checkBucketExists $BUCKET ; then
    echo "Creating bucket '$BUCKET'"
    /usr/bin/mc mb local/$BUCKET
  else
    echo "Bucket '$BUCKET' already exists."
  fi

  # At this point, the bucket should exist, skip checking for existence
  # Set policy on the bucket
  echo "Setting policy of bucket '$BUCKET' to '$POLICY'."
  /usr/bin/mc policy $POLICY local/$BUCKET
}

mc config host list

# create buckets
# createBucket <bucket-name> <policy>
createBucket workshop public

# list buckets
mc ls local

# setup bucket notifications
if output=$(mc events list local/workshop)  &&  [ -z "$output" ]
then
  mc events add local/workshop arn:minio:sqs::workshop:webhook --suffix .json --events put
else
  echo "There is already an event defined for the workshop bucket:"
  echo "$output"
fi
