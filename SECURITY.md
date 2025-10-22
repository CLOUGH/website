# Security & DDoS Protection

This document outlines the security measures and DDoS protection implemented in the website.

## Rate Limiting

### Global Rate Limiting
- **Limit**: 100 requests/second per IP
- **Burst**: 200 requests
- **Purpose**: Prevent general abuse and excessive resource consumption
- **Applies to**: All endpoints

### Strict API Rate Limiting
- **Limit**: 10 requests/second per IP
- **Burst**: 20 requests
- **Purpose**: Protect API endpoints from abuse
- **Applies to**: `/api/*` endpoints, admin routes

### Contact Form Rate Limiting
- **Limit**: 1 request/second per IP
- **Burst**: 3 requests
- **Purpose**: Prevent spam and abuse of email functionality
- **Applies to**: `/contact` and `/api/contact` POST endpoints

## Response Caching

### Public Pages Cache
- **TTL**: 5 minutes
- **Applies to**: All GET requests (home page, public pages, API GET endpoints)
- **Purpose**: Reduce database load and improve response times
- **Headers**: 
  - `X-Cache: HIT` (cached response)
  - `X-Cache: MISS` (fresh response)
  - `Cache-Control: public, max-age=300`

## Spam Protection

### Contact Form Protection
1. **Honeypot Field**: Hidden field that bots fill but humans don't
2. **Timing Check**: Rejects submissions under 3 seconds
3. **Rate Limiting**: 1 request/minute per IP
4. **Content Analysis**:
   - Blocks excessive links (>3)
   - Detects spam keywords
   - Requires minimum 10 character messages
   - Validates name format

## Server Hardening

### Connection Timeouts
- **Read Timeout**: 10 seconds
- **Write Timeout**: 10 seconds
- **Idle Timeout**: 120 seconds
- **Max Header Size**: 1 MB

### Security Headers
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Referrer-Policy: strict-origin-when-cross-origin`
- `Strict-Transport-Security` (HTTPS only)
- `Content-Security-Policy` (restrictive CSP)
- `Permissions-Policy` (disable unnecessary features)

### CORS Protection
- Whitelist-based origin checking
- Only allows specific trusted domains
- Credentials support for authenticated requests

## Memory Management

### Automatic Cleanup
- **Rate Limiter**: Removes IP entries after 10 minutes of inactivity
- **Cache**: Automatically expires entries after TTL
- **Periodic Cleanup**: Runs every 5 minutes to prevent memory leaks

## Monitoring

### Response Headers
- `X-Cache`: Indicates cache hit/miss
- Custom headers for debugging (development only)

## Cost Optimization

1. **Caching reduces**:
   - Firebase/Firestore reads
   - CPU usage for template rendering
   - Network bandwidth

2. **Rate limiting prevents**:
   - Excessive API calls
   - Email sending abuse
   - Database query floods

3. **Connection limits prevent**:
   - Slowloris attacks
   - Connection exhaustion
   - Memory overflow

## Recommendations for Production

1. **Enable HTTPS** - Security headers work best with TLS
2. **Configure Environment Variables**:
   ```bash
   GIN_MODE=release
   PORT=8080
   ```

3. **Monitor Logs** - Watch for 429 (Too Many Requests) responses
4. **Adjust Limits** - Fine-tune rate limits based on traffic patterns
5. **Use CDN** - Add Cloudflare or similar for additional DDoS protection
6. **Database Backups** - Regular backups of Firebase data
7. **Alert System** - Set up alerts for unusual traffic patterns

## Testing DDoS Protection

### Test Rate Limiting
```bash
# Should eventually return 429 status
for i in {1..150}; do curl http://localhost:8080/; done
```

### Test Cache
```bash
# First request: X-Cache: MISS
curl -I http://localhost:8080/

# Second request (within 5min): X-Cache: HIT
curl -I http://localhost:8080/
```

### Test Contact Form Rate Limit
```bash
# Third request should fail
curl -X POST http://localhost:8080/api/contact \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@test.com","message":"Test message"}'
```

## Emergency Response

If under active DDoS attack:

1. **Temporarily reduce rate limits** in `main.go`
2. **Lower cache TTL** to reduce stale data
3. **Enable additional firewall rules** at infrastructure level
4. **Consider enabling Cloudflare** for automatic DDoS mitigation
5. **Monitor Firebase quotas** and upgrade if needed

## Cost Estimates

With these protections:
- **Caching**: ~70-90% reduction in database reads
- **Rate Limiting**: Blocks abusive traffic before processing
- **Connection Limits**: Prevents memory/CPU exhaustion

Expected savings: **$50-200/month** depending on attack volume.
