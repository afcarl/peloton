#!/usr/bin/env bash
# Takes an list of images (built locally) like uber/peloton:foo123, and tags and
# pushes to some registries.
#
# Usage:
# REGISTRY=[docker-registry] ./docker-push

[[ -z "${REGISTRY}" ]] && echo "No REGISTRY set" && exit 1;

[[ $(uname) == Darwin || -n ${JENKINS_HOME} ]] && docker_cmd='docker' || docker_cmd='sudo docker'

for image in "${@}" ; do
  echo "Pushing $image to ${REGISTRY}..."

  new_image="vendor/peloton"
  # pull version from the image, assume latest if not present
  ver="${image##*:}"
  if [[ "$ver" != "$image" ]] ; then
    version="${ver:-latest}"
  else
    version="latest"
  fi

  push_target="${REGISTRY}/${new_image}:${version}"

  ${docker_cmd} tag "${image}" "${push_target}";
  if ${docker_cmd} push "${push_target}"; then
    echo "The image can now be pulled from docker registry at ${REGISTRY}"
  else
    echo "Failed to push image to registry"
    exit 1
  fi
done
