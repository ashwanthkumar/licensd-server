[![Build Status](https://travis-ci.org/ashwanthkumar/licensd-server.svg?branch=master)](https://travis-ci.org/ashwanthkumar/licensd-server)
# licensd.in
LicensenD stores all the licenses of your dependencies used within the project and helps you make actionable insights on them.

## Formats Supported
- [LicenseFinder](https://github.com/pivotal/LicenseFinder)
- [sbt-license-report](https://github.com/sbt/sbt-license-report)

### Notes
Create a copy of `config.sample.yml` in the current working directory as `config.yml` with the right values specified.

#### Sending a CSV report from CI
For `license_finder` report use `file-format=license_finder` and for `sbt-license-report` use `file-format=sbt`.

```
curl -X POST http://localhost:8080/payload \
  -H "X-Licensd-API-Token: abcdef12345" \
  -F "license-file=@license.csv" \
  -F "project-type=maven"  \
  -F "build-version=${GO_PIPELINE_NUMBER}" \
  -F "build-url=${GO_SERVER_URL}/pipelines/${GO_PIPELINE_NAME}/${GO_PIPELINE_COUNTER}/${GO_STAGE_NAME}/${GO_STAGE_COUNTER}" \
  -F "matrix=all" \
  -F "file-format=sbt" \
  -H "Content-Type: multipart/form-data"
```

## License
https://www.apache.org/licenses/LICENSE-2.0
