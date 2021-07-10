// ./pages/api/login
import { setAuthCookies } from 'next-firebase-auth'
import initAuth from '../../lib/init-auth' // the module you created above

initAuth()

const handler = async (req: any, res: any) => {
  try {
    await setAuthCookies(req, res)
  } catch (e) {
    return res.status(500).json({ error: 'Unexpected error.' })
  }
  return res.status(200).json({ success: true })
}

export default handler