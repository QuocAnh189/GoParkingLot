//hook
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

//redux
import { useSignUpMutation } from '@redux/services/auth'
import { useAppDispatch } from '@redux/hook'

//components
import { toast, ToastContainer } from 'react-toastify'
import Loading from '@components/common/Loading'

//asset
import logo from '@assets/logo.png'

//interfaces
import { SignUpPayload } from '@interfaces/auth'
import { setAuth } from '@redux/slices/auth'

//icon
import { EyeIcon, EyeOff } from 'lucide-react'

const initForm: SignUpPayload = {
  email: '',
  name: '',
  password: '',
  avatar: null,
}

const SignUp = () => {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  const [showPassword, setShowPassword] = useState<boolean>(false)
  const [form, setForm] = useState<SignUpPayload>(initForm)

  const [Register, { isLoading }] = useSignUpMutation()

  const handleChangeForm = (name: string, value: any) => {
    setForm((prev) => ({ ...prev, [name]: value }))
  }

  const handleRegister = async () => {
    const formData = new FormData()
    formData.append('email', form.email)
    formData.append('name', form.name)
    formData.append('password', form.password)
    formData.append('avatar', form.avatar)

    try {
      const result = await Register(formData).unwrap()

      if (result) {
        dispatch(setAuth(result))
        localStorage.setItem(
          'token',
          JSON.stringify({
            accessToken: result.accessToken,
            refreshToken: result.refreshToken,
          }),
        )
        localStorage.setItem('user', JSON.stringify(result.user))
        toast.success('Register successfully', {
          autoClose: 200,
        })
        navigate('/main')
      }
    } catch (e: any) {
      toast.error(e.data.message, {
        autoClose: 400,
      })
    }
  }

  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <div className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
          <img className="w-8 h-8 mr-2" src={logo} alt="logo" />
          Quản lý nhà xe
        </div>
        <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
              Đăng ký
            </h1>
            <div className="space-y-4 md:space-y-6">
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your Email</label>
                <input
                  value={form.email}
                  type="email"
                  name="email"
                  id="email"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="name@company.com"
                  onChange={(e) => {
                    handleChangeForm('email', e.target.value)
                  }}
                />
              </div>

              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">User Name</label>
                <input
                  value={form.name}
                  type="text"
                  name="username"
                  id="username"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="Enter Username"
                  onChange={(e) => {
                    handleChangeForm('name', e.target.value)
                  }}
                />
              </div>

              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                <div className="relative">
                  <input
                    value={form.password}
                    type={showPassword ? 'text' : 'password'}
                    name="password"
                    id="password"
                    onChange={(e) => {
                      handleChangeForm('password', e.target.value)
                    }}
                    placeholder="••••••••"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  />
                  <button
                    onClick={() => {
                      setShowPassword(!showPassword)
                    }}
                    className="absolute right-3 top-[50%] translate-y-[-50%] hover:cursor-pointer"
                  >
                    {showPassword ? <EyeIcon width={20} /> : <EyeOff width={20} />}
                  </button>
                </div>
              </div>

              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Avatar</label>
                <input
                  type="file"
                  name="avatar"
                  id="avatar"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  onChange={(e: any) => {
                    handleChangeForm('avatar', e.target.files[0])
                  }}
                />
              </div>

              <button
                disabled={isLoading}
                onClick={handleRegister}
                className="w-full text-white bg-blue-500 hover:bg-blue-400 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
              >
                {isLoading ? <Loading /> : 'Create an account'}
              </button>

              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                Already have an account?{' '}
                <button
                  onClick={() => {
                    navigate('/')
                  }}
                  className="font-medium text-primary-600 hover:underline dark:text-primary-500"
                >
                  Login here
                </button>
              </p>
            </div>
          </div>
        </div>
        <ToastContainer />
      </div>
    </section>
  )
}

export default SignUp
