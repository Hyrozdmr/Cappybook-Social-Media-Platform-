import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
import "./SignupPage.css"; // Import CSS for the signup page
import { signup } from "../../services/authentication";

export const SignupPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      await signup(email, password);
      console.log("redirecting...:");
      navigate("/login");
    } catch (err) {
      console.error(err);
      navigate("/signup");
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
          <div className="signup-container">
            <form className="signup-form" onSubmit={handleSubmit}>
              <h1>Sign Up</h1>
              <h2>Create Your Account</h2>
              <input
                  id="email"
                  type="text"
                  value={email}
                  onChange={handleEmailChange}
                  placeholder="Email"
              />
              <input
                  id="password"
                  type="password"
                  value={password}
                  onChange={handlePasswordChange}
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
          <div className="login">
            <Link to="/login">Login</Link>
          </div>
        </div>
        <div className="right-half">
          <h3>Login Here</h3>
          <p>If you already have an account, login here:</p>
          <Link className="login-link" to="/login">Login</Link>
        </div>
      </div>
  );
};
