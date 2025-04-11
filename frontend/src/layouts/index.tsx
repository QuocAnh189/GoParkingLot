import { Outlet, useNavigate } from 'react-router-dom'

import routes from '@constants/routes'

//component
import Link from '@components/common/Link'
import AvatarUser from '@components/common/Avatar'
import Loading from '@components/common/Loading'
import { Button } from '@components/ui/button'

//redux
import { useSignOutMutation } from '@redux/services/auth'
import { useAppDispatch } from '@redux/hook'
import { setAuth } from '@redux/slices/auth'

const Layout = () => {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  const user = JSON.parse(localStorage.getItem('user')!)
  const [SignOut, { isLoading }] = useSignOutMutation()

  const handleLogout = async () => {
    const result = await SignOut()
    if (result) {
      console.log('result: ')
      dispatch(setAuth(null))
      localStorage.removeItem('token')
      navigate('/')
    }
  }

  return (
    <div className="h-screen">
      <header className="flex  items-center justify-between gap-6 py-4 px-4">
        <div className="flex items-center gap-6">
          {routes.map((route: any, index: number) => (
            <Link key={`link-${index}`} link={route} />
          ))}
        </div>

        <div className="flex items-center gap-6">
          <div className="flex items-center gap-2 hover:cursor-pointer">
            <AvatarUser avatar_url={user?.avatar_url} />
            <p>{user?.name}</p>
          </div>
          <Button onClick={handleLogout} className="hover:cursor-pointer">
            {isLoading ? <Loading /> : 'Logout'}
          </Button>
        </div>
      </header>
      <Outlet />
    </div>
  )
}

export default Layout
