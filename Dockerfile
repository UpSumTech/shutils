FROM upsumtech/golang:onbuild-1.14.0
LABEL git_tag="$GIT_TAG" \
  build_time="$BUILD_TIME" \
  git_ref="$GIT_REF" \
  github_user="$GITHUB_USERNAME" \
  github_organization="$GITHUB_ORGANIZATION" \
  build_user="$BUILD_USER" \
  repo_name="$REPO_NAME"

ENV IMAGE_TAG="$GIT_TAG" \
  BUILD_TIMESTAMP="$BUILD_TIME" \
  BUILD_RUNNER="$BUILD_USER" \
  BUILD_CONTEXT_URL="$GIT_REF"
