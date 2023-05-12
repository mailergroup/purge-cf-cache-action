# GitHub Action to Purge Cloudflare Cache  🗑️ 


## Usage

All sensitive variables should be [set as encrypted secrets](https://help.github.com/en/articles/virtual-environments-for-github-actions#creating-and-using-secrets-encrypted-variables) in the action's configuration.

### `workflow.yml` Example

#### Run using CloudFlare auth token
```yaml
    - name: Purge cache
      uses: mailergroup/purge-cf-cache-action@v0.0.2
      with:
        cf_token: ${{ secrets.CF_TOKEN }}
        cf_zone_name: ${{ secrets.CF_ZONE_NAME }}
```

### Purging specific files

To purge only specific files, you can pass an array of **fully qualified URLs** via a fourth environment variable named `CF_PURGE_URLS`.

```yaml
cf_purge_urls: '["https://aorfanos.com/styles.css","https://aorfanos.com/style.css"]'
```

### Purging by host

To purge only specific files, you can pass an array of **fully qualified URLs** via a fourth environment variable named `CF_PURGE_URLS`.

```yaml
cf_purge_hosts: "https://mailergroup.com,https://foo.mailergroup.com"
```
