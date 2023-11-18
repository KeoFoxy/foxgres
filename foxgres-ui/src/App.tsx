import { Routes, Route } from 'react-router-dom';
import {
  Login,
  Home,
  Profile,
  ErrorPage,
} from './pages';

import {
  Layout,
} from './components';

import './App.css'

function App() {
  return (
    <Routes>
      <Route path='/' element={<Layout />}>
        {/* Login Page */}
        <Route path='login' element={<Login />} />

        {/* Main Pages */}
        {/* <Route element={<RequireAuth />}>   */}
          <Route path='home' element={<Home />} />
          <Route path='profile' element={<Profile />} />
        {/* </Route> */}

        {/* Error route */}
        <Route path='*' element={<ErrorPage />}/>
      </Route>
    </Routes>
  )
}

export default App
