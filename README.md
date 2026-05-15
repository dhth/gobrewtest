# gobrewtest

A demo project to test homebrew publishing using
[goreleaser](https://github.com/goreleaser/goreleaser).

🔐 Verifying release artifacts
---

In case you get the `gobrewtest` binary directly from a
[release](https://github.com/dhth/gobrewtest/releases), you may want to verify
its authenticity. Checksums are applied to all released artifacts, and the
resulting checksum file is signed using
[cosign](https://github.com/sigstore/cosign) (v3).

1. Get the checksum and cosign signature from the release:

    ```shell
    curl -sSLO https://github.com/dhth/gobrewtest/releases/download/vx.y.z/gobrewtest_x.y.z_checksums.txt
    curl -sSLO https://github.com/dhth/gobrewtest/releases/download/vx.y.z/gobrewtest_x.y.z_checksums.txt.sigstore.json
    ```

2. Verify the signature of the checksum file:

    ```shell
    cosign verify-blob \
        --bundle gobrewtest_x.y.z_checksums.txt.sigstore.json \
        --certificate-identity-regexp 'https://github\.com/dhth/gobrewtest/\.github/workflows/.+' \
        --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
        gobrewtest_x.y.z_checksums.txt
    ```

3. Download the compressed archive you want, and validate its checksum:

    ```shell
    curl -sSLO https://github.com/dhth/gobrewtest/releases/download/vx.y.z/gobrewtest_x.y.z_linux_amd64.tar.gz
    sha256sum --ignore-missing -c gobrewtest_x.y.z_checksums.txt
    ```

3. If checksum validation goes through, uncompress the archive:

    ```shell
    tar -xzf gobrewtest_x.y.z_linux_amd64.tar.gz
    ./gobrewtest -h
    ```
