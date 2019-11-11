#!/bin/bash

NS=$1

if [ -z "$NS" ]; then
    echo "Specify a namespace argument"
    exit 1
fi

if [ -z "$2" ]; then
    echo "No enterprise license set"
    exit 0
fi

case $(uname) in
    Darwin)
        LICENSE=$(echo -n "$2" | base64 -b 0)
        ;;
    *)
        LICENSE=$(echo -n "$2" | base64 -w 0)
        ;;
esac

kubectl apply -f - <<EOF
apiVersion: v1
data:
  token: ${LICENSE}
kind: Secret
metadata:
  name: arangodb-jenkins-license-key
  namespace: ${NS}
type: Opaque
EOF
