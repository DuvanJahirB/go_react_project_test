import { useState } from 'react'
import AuthForm from './pages/AuthForm'
import Dashboard from './pages/Dashboard'
// import Dashboard from './pages/Dashboard'
import {Route,Routes} from 'react-router-dom'
import './index.css'
import {UserContextProvider} from './context/UserContext.jsx'

function App() {
  return (
    <UserContextProvider>
      <Routes>
        <Route path="/login" element={<AuthForm/>}></Route>
        <Route path="/" element={<Dashboard/>}></Route>
      </Routes>
    </UserContextProvider>
  )
}

export default App