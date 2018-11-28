
* Enable Cloud KMS
* Add permission to Cloud Build Service Account to use Cloud KMS
* Get github.com Access Token

```
# Create github.com setting
cat <<EOL >.netrc
machine github.com
    login nirasan
    password [GITHUB.COM ACCESS TOKEN]
EOL

# Create github.com known_hosts
ssh-keyscan -t rsa github.com > known_hosts

# Create Cloud KMS keyring
gcloud kms keyrings create gcb-keyring --location=global

# Create Cloud KMS key
gcloud kms keys create github-key \
--location=global --keyring=gcb-keyring \
--purpose=encryption

# Encrypt github.com setting
gcloud kms encrypt --plaintext-file=./.netrc \
--ciphertext-file=./.netrc.enc \
--location=global --keyring=gcb-keyring --key=github-key

# Upload encrypted github.com private key
gsutil mb gs://gcb-examples/
gsutil cp ./.netrc.enc gs://gcb-examples/
```
