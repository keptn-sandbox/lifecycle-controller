#!/bin/bash

######################################################################################
#
# This script heavily uses Bash parameter expansion.
# To better understand what is going on here, you can read up on the topic here:
# https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html
#
######################################################################################

CHANGED_FILES=$1

if [ $# -ne 1 ]; then
  echo "Usage: $0 CHANGED_FILES"
  exit
fi

# initialize variables with false (make sure they are also set in needs.prepare_ci_run.outputs !!!)
BUILD_FUNCTIONS_RUNTIME_SVC=false
BUILD_LFC_SCHEDULER_SVC=false
BUILD_OPERATOR_SVC=false

artifacts=(
  "$FUNCTIONS_RUNTIME_SVC_ARTIFACT_PREFIX"
  "$LFC_SCHEDULER_SVC_ARTIFACT_PREFIX"
  "$OPERATOR_SVC_ARTIFACT_PREFIX"
)

echo "Changed files:"
echo "$CHANGED_FILES"
matrix_config='{"config":['
# shellcheck disable=SC2016
build_artifact_template='{"artifact":$artifact,"working-dir":$working_dir,"should-run":$should_run}'

# Add all changed artifacts to the build matrix
echo "Checking changed files against artifacts now"
echo "::group::Check output"
for changed_file in $CHANGED_FILES; do
  echo "Checking if $changed_file leads to a build..."

  for artifact in "${artifacts[@]}"; do
    # Prepare variables
    artifact_fullname="${artifact}_ARTIFACT"
    artifact_folder="${artifact}_FOLDER"
    should_build_artifact="BUILD_${artifact}"
    should_run="SHOULD_RUN_${artifact}"

    if [ "${!should_run}" != "false" ]; then
      should_run="true"
    else
      should_run="false"
    fi

    if [[ ( $changed_file == ${!artifact_folder}* ) && ( "${!should_build_artifact}" != 'true' ) ]]; then
      echo "Found changes in $artifact"
      # Set the artifact's should-build variable to true
      IFS= read -r "${should_build_artifact?}" <<< "true"

      # Render build matrix string for the current artifact
      artifact_config=$(jq -j -n \
        --arg artifact "${!artifact_fullname}" \
        --arg working_dir "${!artifact_folder}" \
        --arg should_run "${should_run}" \
        "$build_artifact_template"
      )

      # Add rendered string to matrix
      matrix_config="$matrix_config$artifact_config,"
    fi
  done
done

echo "Done checking changed files"
echo "Checking for build-everything build"

# If this is a build-everything build, also add all other unchanged artifacts to the build matrix
if [[ $BUILD_EVERYTHING == 'true' ]]; then
  for artifact in "${artifacts[@]}"; do
    # Prepare variables
    artifact_fullname="${artifact}_ARTIFACT"
    artifact_folder="${artifact}_FOLDER"
    should_build_artifact="BUILD_${artifact}"
    should_run="SHOULD_RUN_${artifact}"

    if [ "${!should_run}" != "false" ]; then
      should_run="true"
    else
      should_run="false"
    fi

    if [[ "${!should_build_artifact}" != 'true' ]]; then
      # Render build matrix string for the current artifact
      echo "Adding unchanged artifact $artifact to build matrix since build everything was requested"
      artifact_config=$(jq -j -n \
        --arg artifact "${!artifact_fullname}" \
        --arg working_dir "${!artifact_folder}" \
        --arg should_run "${should_run}" \
        "$build_artifact_template"
      )

      # Add rendered string to matrix
      matrix_config="$matrix_config$artifact_config,"
    fi
  done
fi
echo "::endgroup::"


# Terminate matrix JSON config and remove trailing comma
matrix_config="${matrix_config%,}]}"

# Escape newlines for multiline string support in GH actions
# Reference: https://github.community/t/set-output-truncates-multiline-strings/16852
matrix_config="${matrix_config//'%'/''}"
matrix_config="${matrix_config//$'\n'/''}"
matrix_config="${matrix_config//$'\r'/''}"
matrix_config="${matrix_config//$' '/''}"

echo "::group::Build Matrix"
echo "$matrix_config"
echo "::endgroup::"

# print job outputs (make sure they are also set in needs.prepare_ci_run.outputs !!!)
echo "::set-output name=BUILD_MATRIX::$matrix_config"
echo ""
echo "The following artifacts have changes and will be built fresh:"
echo "BUILD_FUNCTIONS_RUNTIME_SVC: $BUILD_FUNCTIONS_RUNTIME_SVC"
echo "BUILD_LFC_SCHEDULER_SVC: $BUILD_LFC_SCHEDULER_SVC"
echo "BUILD_OPERATOR_SVC: $BUILD_OPERATOR_SVC"

if [[ "$matrix_config" == '{"config":[]}' ]]; then
  echo "Build matrix is emtpy, setting output..."
  echo "::set-output name=BUILD_MATRIX_EMPTY::true"
else
  echo "Build matrix is NOT emtpy, setting output..."
  echo "::set-output name=BUILD_MATRIX_EMPTY::false"
fi
