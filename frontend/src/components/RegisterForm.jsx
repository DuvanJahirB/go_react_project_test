import { useEffect, useState } from 'react';
import { useForm } from '../context/FormContext';
import { EyeIcon, EyeSlashIcon } from '../components/Icons'

export default function RegisterForm() {
    const { formData,errors,handleInputChange,showPassword,setShowPassword,showConfirmPassword,setShowConfirmPassword} = useForm();

    return (
        <>
            <div>
                <label htmlFor="fullName" className="block text-sm font-medium text-gray-700 mb-1">Full Name</label>
                <input
                    type="text"
                    id="fullName"
                    name="fullName"
                    value={formData.fullName}
                    onChange={handleInputChange}
                    className={`w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 ${
                        errors.fullName ? 'border-red-500' : 'border-gray-300'
                    }`}
                    placeholder="Enter your full name"
                />
                {errors.fullName && (<p className="mt-1 text-sm text-red-600">{errors.fullName}</p>)}
            </div>
            <div>
                <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">Email Address</label>
                <input
                    type="email"
                    id="email"
                    name="email"
                    value={formData.email}
                    onChange={handleInputChange}
                    className={`w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 ${errors.email ? 'border-red-500' : 'border-gray-300'
                        }`}
                    placeholder="Enter your email"
                />
                {errors.email && (<p className="mt-1 text-sm text-red-600">{errors.email}</p>)}
            </div>
            <div>
                <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">Password</label>
                <div className="relative">
                    <input
                        type={showPassword ? 'text' : 'password'}
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleInputChange}
                        className={`w-full px-4 py-3 pr-12 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 ${errors.password ? 'border-red-500' : 'border-gray-300'
                            }`}
                        placeholder="Enter your password"
                    />
                    <button
                        type="button"
                        onClick={() => setShowPassword(!showPassword)}
                        className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-500 hover:text-gray-700 transition-colors duration-200"
                    >
                        {showPassword ? <EyeSlashIcon /> : <EyeIcon />}
                    </button>
                </div>
                {errors.password && (<p className="mt-1 text-sm text-red-600">{errors.password}</p>)}
            </div>
            <div className="animate-slide-in">
                <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-700 mb-1">Confirm Password</label>
                <div className="relative">
                    <input
                        type={showConfirmPassword ? 'text' : 'password'}
                        id="confirmPassword"
                        name="confirmPassword"
                        value={formData.confirmPassword}
                        onChange={handleInputChange}
                        className={`w-full px-4 py-3 pr-12 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 ${errors.confirmPassword ? 'border-red-500' : 'border-gray-300'}`}
                        placeholder="Confirm your password"
                    />
                    <button
                        type="button"
                        onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                        className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-500 hover:text-gray-700 transition-colors duration-200"
                    >
                        {showConfirmPassword ? <EyeSlashIcon /> : <EyeIcon />}
                    </button>
                </div>
                {/* {errors.confirmPassword && (
                    <p className="mt-1 text-sm text-red-600">{errors.confirmPassword}</p>
                )} */}
            </div>
        </>
    )
}