// This context interacts with user information
import { createContext, useContext, useState} from "react";
import axios from 'axios'
import { useNavigate }  from 'react-router-dom'

export const UserContext = createContext();

export const useUser = () => {
    const context = useContext(UserContext);
    if(!context) {
        throw new Error("use User must be used within a UserContextProvider");
    }
    return context;
}

export const UserContextProvider = ({children}) => {
    // const navigate = useNavigate();
    const [user,setUser] = useState(null);
    const [loading, setLoading] = useState(false);
    const url = import.meta.env.VITE_BACKEND_URL;
    const navigate = useNavigate();
    // Get request that obtains the user's information and verifies that the user is logged in
    async function login(email,password) {
        setLoading(true);
        try {
            const response = await axios.post(`${url}/auth/login`, { email, password });
            sessionStorage.setItem("token",response.data.token);
            return { success: true};
        } catch (error) {
            console.error('Login error:', error);
            return { success: false, message: error.message};
        } finally {
            setLoading(false);
        }
    }
    
    async function register(name,email,password) {
        setLoading(true);
        try {
            const response = await axios.post(`${url}/user`, {"name":name, email, password });
            return {success: true, message: response.data.message}
            // sessionStorage.setItem("token",response.data.token);
        } catch (error) {
            console.error('Register error:', error.response.data.error);
            return { success: false, message: error.response.data.error};
        } finally {
            setLoading(false);
        }
    }

    // Get request that obtains the user's information and verifies that the user is logged in
    async function getProfile() {
        try {
            const strStorage = sessionStorage.getItem("token");
            // const {token} = JSON.parse(strStorage);
            const reqOptions = {
                url: `${url}/profile`,
                method: "GET",
                headers: {
                    "Authorization": `Bearer ${strStorage}`
                }
            }
            let response = await axios.request(reqOptions);
            setUser(response.data);
        } catch (error) {
            // If no response is obtained from the request, redirect to the login
            console.error('Get profile error');
            setUser(null);
            // sessionStorage.removeItem('token');
            navigate('/login');
        }
    }
    
    return <UserContext.Provider value={{ user, loading, login, register, getProfile }}> {children} </UserContext.Provider>;
}