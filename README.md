# licensd.in

### Notes
Create a copy of `config.sample.yml` in the current working directory as `config.yml` with the right values specified.

#### Sending a license_finder CSV report from CI
```
curl -X POST http://localhost:8080/payload \
  -H "X-Licensd-API-Token: abcdef12345" \
  -F "file=@license.csv" \
  -F "package_manager=maven"  \
  -F "version=${GO_PIPELINE_NUMBER}" \
  -F "matrix=all" \
  -H "Content-Type: multipart/form-data"
```

## License
https://www.apache.org/licenses/LICENSE-2.0
