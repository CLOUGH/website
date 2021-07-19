import 'bootstrap/dist/css/bootstrap.css'
import '../assets/styles/globals.css'
import type { AppProps } from 'next/app'
import Head from 'next/head';
import { AuthUserProvider } from '../context/AuthUserContext';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <AuthUserProvider>
      <Head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <Component {...pageProps} />
    </AuthUserProvider>

  )
}
export default MyApp
