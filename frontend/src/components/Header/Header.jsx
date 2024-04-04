import "./Header.css"
import capybara from "../../static/img/capybara-logo.png"
import { Link } from "react-router-dom";
import { useEffect, useState } from "react";

export default function Header() {
  const [navbar, setNavbar] = useState([]);
  const hasToken = localStorage.getItem("token")
  
  function removeToken() {
    localStorage.removeItem("token")
  }

  useEffect(() => {
    if (hasToken !== null) {
      setNavbar(
        <div className="nav-bar">
          <Link onClick={removeToken}  to="/">Log out</Link>
          <Link to="/posts">Feed</Link>
        </div>
      );
    } else {
      setNavbar(
        <div className="nav-bar">
          <Link to="/signup">Sign up</Link>
          <Link to="/login">Log in</Link>
        </div>
      );
    }
  }, [hasToken]);

  return (
  <div className="header">
    <div className="logo">
      <img src={capybara} alt="logo" />
      <Link to="/">Acebook</Link>
    </div>
    {navbar}
  </div>)

}