# Reference

## [Understanding and Preventing DoS in web applications](https://r2c.dev/blog/2020/understanding-and-preventing-dos-in-web-apps/)

- High leverage DOS vulnerabilities
    - Disk space: magnify and fill disk
    - Network bandwidth
    - CPU Utilization
    - Concurrency Limit
- Level of authentication required to trigger the exploit
    - if unauthenticated -> high risk -> eliminate
    - resource starvation -> rate limit
    - DDOS -> Cloudflare. Network rules to block
