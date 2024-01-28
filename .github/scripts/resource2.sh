#!/bin/bash

# Keptn Helm Testing
#
# This script supports the comparison of standard values and expected templated results to helm chart
# it is used to make sure changes to the chart are intentional and produce expected outcomes

echo "copying manifests"

dir="lifecycle-operator/config/crd/bases"
helm_dir="lifecycle-operator/chart/templates"

pickElement=(keptnappcontext keptnappcreationrequest)

n=10
truncate -s 0 $helm_dir/keptnapp-crd.yaml
cat $dir/lifecycle.keptn.sh_keptnapps.yaml >> $helm_dir/keptnapp-crd.yaml
# Loop through each file in the directory
for file in "$dir"/* ; do
    # Extract the basename of the file
    filename=$(basename "$file" .yaml)
    echo "Processing file: $filename"

    # Loop through each element in the pickElement array
    for element in "${pickElement[@]}"; do
        echo "Checking element: $element"

        # Check if the element is present in the filename
        if echo "$filename" | grep -q "$element" ; then
            echo "Match found: $element in $filename"

            # Loop through files in the helm directory
            for crds in "$helm_dir"/* ; do
                # Extract the basename of the helm file
                helm_filename=$(basename "$crds" .yaml)

                echo "Checking helm file: $helm_filename"

                # Check if the element is present in the helm filename
                if echo "$helm_filename" | grep -q "$element" ; then

                    echo "Match found: $element in $helm_filename"
                      truncate -s 0 "$crds"
                      ((n=n-1))
                    # Concatenate the content of the file into the helm file
                    cat "$file" >> "$crds"
                    break
                fi
            done
        fi
    done
done

for file in "$helm_dir"/*; do
    filename=$(basename "$file" .yaml)
    echo "Processing file: $filename.yaml"
    # Rest of your script...
done

for file in "$helm_dir"/*; do
    filename=$(basename "$file" .yaml)
        if [[ $filename == k* ]] ; then
                sed -i '/controller-gen.kubebuilder.io\/version: v0.14.0/a\
    {{- with .Values.global.caInjectionAnnotations }}\
    {{- toYaml . | nindent 4 }}\
    {{- end }}\
    {{- include "common.annotations" ( dict "context" . ) }}\
  labels:\
    app.kubernetes.io/part-of: keptn\
    crdGroup: lifecycle.keptn.sh\
    keptn.sh/inject-cert: "true"\
    {{- include "common.labels.standard" ( dict "context" . ) | nindent 4 }' "$helm_dir/$filename.yaml"
        fi

            if [[ "$filename" == "keptnappcontext-crd" ]] ; then
            sed -i "s|{{- with .Values.global.caInjectionAnnotations }}|cert-manager.io/inject-ca-from: '{{ .Release.Namespace }}/keptn-certs'|g" "$helm_dir/$filename.yaml"
            sed -i "/{{- toYaml \. \| nindent 4 }}/d" "$helm_dir/$filename.yaml"
            sed -i "/{{- end }}/d" "$helm_dir/$filename.yaml"
            fi
done