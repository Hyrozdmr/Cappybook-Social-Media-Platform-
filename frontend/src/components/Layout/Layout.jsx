import Header from '../Header/Header'
import Footer from '../Footer/Footer'
import { Outlet } from 'react-router-dom';
import "./Layout.css"

function Layout({ children }) {

  return (
    <>
    <div className='layout-container'>
      <Header/>
      {children}
      <Outlet/>
      <Footer />
    </div>
    </>
  )
}

export default Layout;