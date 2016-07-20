#!/bin/sh
#
# Generate test coverage statistics for Go packages.
#

set -e

output() {
  printf "\033[32m"
  echo $1
  printf "\033[0m"

  if [ "$2" = 1 ]; then
    exit 1
  fi
}

workdir=".cover"
coverage_report="$workdir/coverage.txt"
coverage_xml_report="$workdir/coverage.xml"
junit_report="$workdir/junit.txt"
junit_xml_report="$workdir/report.xml"
lint_report="$workdir/lint.txt"
vet_report="$workdir/vet.txt"
packages=$(go list ./... | grep -v vendor)

test -d $workdir || mkdir -p $workdir

show_help() {
cat << EOF
Generate test coverage statistics for Go packages.

  testing [set|count|atomic]     Run go testing for all packages
  coverage                       Generate coverage report for all packages
  junit                          Generate coverage xml report for junit plugin
  lint                           Generate Lint report for all packages
  vet                            Generate Vet report for all packages
EOF
}

testing() {
  test -f ${junit_report} && rm -f ${junit_report}
  coverage_mode=$@
  output "Running ${coverage_mode} mode for coverage."
  for pkg in $packages; do
    f="$workdir/$(echo $pkg | tr / -).cover"
    output "Testing coverage report for ${pkg}"
    go test -v -cover -coverprofile=${f} -covermode=${coverage_mode} $pkg | tee -a ${junit_report}
  done

  output "Convert all packages coverage report to $coverage_report"
  echo "mode: $coverage_mode" > "$coverage_report"
  grep -h -v "^mode:" "$workdir"/*.cover >> "$coverage_report"
}

generate_cover_report() {
  gocov convert ${coverage_report} | gocov-xml > ${coverage_xml_report}
}

generate_junit_report() {
  cat $junit_report | go-junit-report > ${junit_xml_report}
}

generate_lint_report() {
  for pkg in $packages; do
    output "Go Lint report for ${pkg}"
    golint ${pkg} | tee -a ${lint_report}
  done
}

generate_vet_report() {
  for pkg in $packages; do
    output "Go Vet report for ${pkg}"
    go vet -n -x ${pkg} | tee -a ${vet_report}
  done
}

case "$1" in
  "")
    show_help ;;
  testing)
    mode="set"
    test -z $2 || mode=$2
    testing $mode ;;
  coverage)
    generate_cover_report ;;
  junit)
    generate_junit_report ;;
  lint)
    generate_lint_report ;;
  vet)
    generate_vet_report ;;
  *)
    show_help ;;
esac
