import { useState } from 'react'
import {Toaster} from 'sonner'

import LogInForm from '../components/LogInForm.jsx'
import RegisterForm from '../components/RegisterForm.jsx'
import {FormContextProvider,useForm} from '../context/FormContext.jsx'

const FormContent = () => {
  const { activeTab,switchTab,handleSubmit} = useForm();
  return(<div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-blue-50 flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        <div className="bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden">
          {/* Header */}
          <div className="px-8 pt-8 pb-6">
            <div className="text-center">
              <h1 className="text-3xl font-bold text-gray-900 mb-2">Welcome</h1>
              <p className="text-gray-600">Sign in to your account or create a new one</p>
            </div>
          </div>

          {/* Tabs */}
          <div className="px-8">
            <div className="flex bg-gray-100 rounded-lg p-1 mb-8">
              <button
                onClick={() => switchTab('login')}
                className={`flex-1 py-2 px-4 rounded-md text-sm font-medium transition-all duration-200 ${
                  activeTab === 'login'
                    ? 'bg-white text-blue-600 shadow-sm'
                    : 'text-gray-600 hover:text-gray-900'
                }`}
              >
                Sign In
              </button>
              <button
                onClick={() => switchTab('register')}
                className={`flex-1 py-2 px-4 rounded-md text-sm font-medium transition-all duration-200 ${
                  activeTab === 'register'
                    ? 'bg-white text-blue-600 shadow-sm'
                    : 'text-gray-600 hover:text-gray-900'
                }`}
              >
                Sign Up
              </button>
            </div>
          </div>

          {/* Form */}
          <div className="px-8 pb-8">
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="animate-fade-in">
                {activeTab === 'register' ? 
                  <RegisterForm/> : 
                  <LogInForm/> }
                <button
                  type="submit"
                  className="w-full bg-blue-600 text-white py-3 px-4 rounded-lg font-medium hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all duration-200 transform hover:scale-[1.02] active:scale-[0.98] mt-4"
                >
                  {activeTab === 'login' ? 'Sign In' : 'Create Account'}
                </button>
              </div>
            </form>

          </div>
        </div>
      </div>
    </div>)
}
export default function AuthForm() {
  

  return (
    <FormContextProvider>
       <Toaster richColors closeButton/>
      <FormContent/>
    </FormContextProvider>
  )
}