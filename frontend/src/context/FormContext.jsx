import { createContext, useContext, useState} from "react";
import {toast} from 'sonner'

const FormContext = createContext();
import {useUser} from './UserContext.jsx'
import { useNavigate }  from 'react-router-dom'

export const useForm = () => {
    const context = useContext(FormContext);
    if(!context) {
        throw new Error("use Form must be used within a FormContextProvider");
    }
    return context;
}

export const FormContextProvider = ({children}) => {
    const { login,register } = useUser();
    const navigate = useNavigate();

    const [activeTab, setActiveTab] = useState('login');
    const [showPassword, setShowPassword] = useState(false)
    const [showConfirmPassword, setShowConfirmPassword] = useState(false)
    const [formData, setFormData] = useState({
        email: '',
        password: '',
        confirmPassword: '',
        fullName: ''
    })
    const [errors, setErrors] = useState({});

    const switchTab = (tab) => {
        setActiveTab(tab)
        setErrors({})
        setFormData({
        email: '',
        password: '',
        confirmPassword: '',
        fullName: '',
        agreeToTerms: false
        })
    }

    const handleInputChange = (e) => {
        const { name, value, type, checked } = e.target
        setFormData(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : value
        }))
        
        // Clear error when user starts typing
        if (errors[name]) {
            setErrors(prev => ({ ...prev, [name]: '' }))
        }
    }

    const validateForm = () => {
        const newErrors = {}
    
        if (!formData.email) {
            newErrors.email = 'Email is required'
        } 
        else if (!/\S+@\S+\.\S+/.test(formData.email)) {
            newErrors.email = 'Please enter a valid email'
        }
    
        if (!formData.password) {
            newErrors.password = 'Password is required'
        } 
        else if (formData.password.length < 6) {
            newErrors.password = 'Password must be at least 6 characters'
        }
    
        if (activeTab === 'register') {
            if (!formData.fullName) {
                newErrors.fullName = 'Full name is required'
            }
      
            if (!formData.confirmPassword) {
                newErrors.confirmPassword = 'Please confirm your password'
            } 
            else if (formData.password !== formData.confirmPassword) {
                newErrors.confirmPassword = 'Passwords do not match'
            }
        }
        return newErrors
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        const newErrors = validateForm()
    
        if (Object.keys(newErrors).length > 0) {
            setErrors(newErrors)
            return
        }
        toast.loading('Send Information');
        if (activeTab === 'register') {
            const val = await register(formData.fullName,formData.email,formData.password);
            if(val.success) {
                toast.success(`${val.message}`);
                switchTab('login');
                
            }
            else{
                toast.error(`Error: ${val.message}`);
            }
            return;
        }
        const val = await login(formData.email,formData.password);
        if(val.success) {
            navigate(`/`);
        }
        else{
            toast.error(`Error: ${val.message}`);
        }
    }

    return <FormContext.Provider value={{showPassword,setShowPassword,showConfirmPassword,setShowConfirmPassword,activeTab,switchTab,formData,errors,handleInputChange,validateForm,handleSubmit}}> {children} </FormContext.Provider>;
}