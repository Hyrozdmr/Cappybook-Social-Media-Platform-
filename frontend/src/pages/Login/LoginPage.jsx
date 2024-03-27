import { useState } from "react";
import {Link, useNavigate} from "react-router-dom";
import "./LoginPage.css";
import { login } from "../../services/authentication";

export const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const token = await login(email, password);
      localStorage.setItem("token", token);
      navigate("/posts");
    } catch (err) {
      console.error(err);
      navigate("/login");
    }
  };

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  return (
      <div className="home">
        <div className="left-half">
          <div className="login-container">
            <form className="login-form" onSubmit={handleSubmit}>
              <h1>Login</h1>
              <h2>Already got an account?</h2>
              <input
                  id="email"
                  type="text"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="Email"
              />
              <input
                  id="password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  placeholder="Password"
              />
              <input
                  role="submit-button"
                  id="submit"
                  type="submit"
                  value="Submit"
              />
            </form>
          </div>
        </div>
        <div className="right-half">
          <h3>New Here?</h3>
          <p>Sign up and start sharing moments with your friends today! Join our community and explore endless possibilities together.</p>
          <Link className="signup" to="/signup">Sign Up</Link>
        </div>
      </div>
  );
};
