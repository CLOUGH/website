
module.exports = {
    async redirects() {
        if (process.env.MAINTENANCE_MODE==='true') {
            return [{
                source: '/((?!\\w+).*$)',
                destination: '/maintenance',
                permanent: false,
            }]
        }
        return [];
    },
}
